package agent

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/playwright-community/playwright-go"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/sashabaranov/go-openai"
)
type ActionType string
type ElementType string

const ( 
	ActionClick ActionType = "click"
 	ActionNav   ActionType = "navigate"
 	ActionFill  ActionType = "fill"
)

const (
	ElementButton	   ElementType = "button"
	ElementTextInput   ElementType = "text-input"
	ElementURLBar	   ElementType = "url"
)

type FindSpec struct {
	Type        ElementType `json:"type"`
	ElementId string      `json:"elementId,omitempty"`
	FillValue   string      `json:"fillValue,omitempty"`
}

type Action struct {
	ActionType 			 ActionType `json:"action"`
	ActDescription       FindSpec   `json:"actDescription"`
}

type TaskStep struct {
	Description string   `json:"description"`
	Task      Action   `json:"task"`
}

type Agent struct {
	llm 		   *openai.Client
	pw 			   *playwright.Playwright
	browser  	   playwright.Browser
	browserContext playwright.BrowserContext
	browserPage    playwright.Page
	state	       *AgentState
	metrics    	   *AgentMetrics  
}
type AgentState struct {
	CurrentURL        string
    PageTitle         string
    FormFields        map[string]string
    NavigationHistory []string
	Retry			  int
}
type AgentResponse struct{
	response []TaskStep 
}

type AgentMetrics struct {
    TasksCompleted    prometheus.Counter
    TasksFailed       prometheus.Counter
    RetryCount        prometheus.Counter
    ExecutionDuration prometheus.Histogram
    BrowserActions    *prometheus.CounterVec
}
func NewAgentMetrics() *AgentMetrics {
    return &AgentMetrics{
        TasksCompleted: promauto.NewCounter(prometheus.CounterOpts{
            Name: "agent_tasks_completed_total",
            Help: "Total number of tasks successfully completed",
        }),
        TasksFailed: promauto.NewCounter(prometheus.CounterOpts{
            Name: "agent_tasks_failed_total",
            Help: "Total number of failed tasks",
        }),
        RetryCount: promauto.NewCounter(prometheus.CounterOpts{
            Name: "agent_task_retries_total",
            Help: "Total number of task retries",
        }),
        ExecutionDuration: promauto.NewHistogram(prometheus.HistogramOpts{
            Name:    "agent_task_duration_seconds",
            Help:    "Duration of task execution in seconds",
            Buckets: prometheus.DefBuckets,
        }),
        BrowserActions: promauto.NewCounterVec(prometheus.CounterOpts{
            Name: "agent_browser_actions_total",
            Help: "Count of browser actions by type",
        }, []string{"action_type"}),
    }
}
func NewAgent()*Agent{
	agent := Agent{}
	godotenv.Load()
	apiKey := os.Getenv("API_KEY")
	llm := openai.NewClient(apiKey)
	agentState := AgentState{
		CurrentURL: "",
		PageTitle: "",
		FormFields: map[string]string{},
		Retry: 0,
	}
	metrics := NewAgentMetrics()
	agent.llm = llm
	agent.state = &agentState
	agent.metrics = metrics
	return &agent
}


func (a *Agent)ExecuteTask(userPrompt string){
	timer := prometheus.NewTimer(a.metrics.ExecutionDuration)
    defer timer.ObserveDuration()
	aiResponse,err := a.createTaskPlan(userPrompt)
	if err != nil{
		a.metrics.TasksFailed.Inc()
		log.Println(err)
		return
	}
	
	for _,response := range aiResponse.response{
		a.metrics.BrowserActions.WithLabelValues(string(response.Task.ActionType)).Inc()
		a.executeBrowserAction(response)
	}

	state, _ := json.MarshalIndent(a.state, "", "  ")
	fmt.Println(string(state))
	a.metrics.TasksCompleted.Inc()
}

func (a *Agent)executeBrowserAction(response TaskStep){
	switch response.Task.ActionType{
	case ActionNav :
		a.launchBrowser()
		a.browserPage.Goto(response.Task.ActDescription.FillValue)
		a.state.CurrentURL = response.Task.ActDescription.FillValue
		heading,err := a.browserPage.Locator("#heading").InnerText()
		if err != nil{
			fmt.Println(err)
		}
		a.state.PageTitle = heading
	case ActionFill:
		a.browserPage.Locator(fmt.Sprintf("#%s",response.Task.ActDescription.ElementId)).Fill(response.Task.ActDescription.FillValue)
		a.state.FormFields[response.Task.ActDescription.ElementId]=response.Task.ActDescription.FillValue
	case ActionClick:
		a.browserPage.Locator(fmt.Sprintf("#%s",response.Task.ActDescription.ElementId)).Click()	

    err := a.browserPage.WaitForURL("**",playwright.PageWaitForURLOptions{
		Timeout:   playwright.Float(5000), 
        WaitUntil: playwright.WaitUntilStateNetworkidle,
	})
    if err != nil {
        // return fmt.Errorf("waiting for navigation failed: %w", err)
    }
    newURL := a.browserPage.URL()
	newPath := strings.Split(newURL, a.state.CurrentURL)[1]
	a.state.NavigationHistory = append(a.state.NavigationHistory, newPath)
	a.state.CurrentURL = newURL
	}
}

func(a *Agent) createTaskPlan(userPrompt string) (*AgentResponse,error){
	steps := []TaskStep{
		{
			Description: "Visit Site",
			Task: Action{
				ActionType: ActionNav,
				ActDescription: FindSpec{
					Type: ElementURLBar,
					FillValue: "http://localhost:5173",
				},
			},
		},
		{
			Description: "Fill Name",
			Task: Action{
				ActionType: ActionFill,
				ActDescription: FindSpec{
					Type: ElementTextInput,
					ElementId: "name",
					FillValue: "Arvind Khoisnam",
				},
			},
		},
		{
			Description: "Fill Email",
			Task: Action{
				ActionType: ActionFill,
				ActDescription: FindSpec{
					Type: ElementTextInput,
					ElementId: "email",
					FillValue: "arvind@example.com",
				},
			},
		},
		{
			Description: "Submit",
			Task: Action{
				ActionType: ActionClick,
				ActDescription: FindSpec{
					Type: ElementButton,
					ElementId: "button",
				},
			},
		},
	}
	exampleJSON, _ := json.MarshalIndent(steps, "", "  ")
	systemPrompt := fmt.Sprintf(`
	You are a well informed and a knowledgeable agent who can understand tasks and give a comprehensive response.
	For a given prompt you are expected to return a response strictly in the following format:
		- description: string describing the step
		- task:
  			- action: "%s" or "%s" or "%s"
  			- actDescription: (for fill/click/navigate)
    			- type: "%s" or "%s" or "%s"
				- elementId : name or email or contact or hobbies or message or image or button 
  				- fillValue: string  with the value to be filled. Omit incase of button click.
				  
	Here is an example: %s

	INSTRUCTIONS:
		1. Be precise in element identification
		2. Include necessary waits between steps
		3. Don't include browser launch
		4. Handle form submissions explicitly
		5. DO NOT include '*/' in the response
	`,ActionClick,ActionNav,ActionFill,ElementButton,ElementURLBar,ElementTextInput,string(exampleJSON))

	resp,err := a.llm.CreateCompletion(context.Background(),openai.CompletionRequest{
		Model:     openai.GPT3Dot5TurboInstruct, 
		Prompt:    systemPrompt + userPrompt,
		MaxTokens: 1000,
		Temperature: 0.3,
	})

	if err != nil{
		fmt.Printf("Completion error: %v\n", err)
	}

	aiResponse := AgentResponse{}
	err = json.Unmarshal([]byte(resp.Choices[0].Text),&aiResponse.response)
	if err != nil{
		fmt.Println(err)
		if a.state.Retry >= 3 {
			return nil, errors.New("an error occurred, retry exceeded, please try again")
		}
		a.metrics.RetryCount.Inc()
		fmt.Println("Retrying again...")
		a.state.Retry = a.state.Retry + 1
		fmt.Printf("retry for %s: %d. \n",strings.Split(strings.Split(userPrompt, "email")[0], "name")[1],a.state.Retry)
		return a.createTaskPlan(userPrompt)
	}	
	return &aiResponse,nil
}
func (a *Agent)launchBrowser(){
	pw,err := playwright.Run()
	if err != nil {
		log.Fatal(err)
	}
	a.pw = pw

	browser,err := a.pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(false),
	})
	if err != nil {
		pw.Stop()
		log.Fatal(err)
	}
	a.browser = browser

	context,err := a.browser.NewContext()
	if err != nil {
        browser.Close()
        pw.Stop()
		log.Fatal(err)
    }
    a.browserContext = context

	page,err := a.browserContext.NewPage()
	if err != nil {
		context.Close()
        browser.Close()
        pw.Stop()
		log.Fatal(err)
    }

	a.browserPage = page
}

func (a *Agent) Close() error {
	var errs []error
	if a.browserPage != nil {
        if err := a.browserPage.Close(); err != nil {
            errs = append(errs, fmt.Errorf("page close error: %w", err))
        }
    }
	if a.browserContext != nil {
        if err := a.browserContext.Close(); err != nil {
            errs = append(errs, fmt.Errorf("context close error: %w", err))
        }
    }
	if a.browser != nil {
        if err := a.browser.Close(); err != nil {
            errs = append(errs, fmt.Errorf("browser close error: %w", err))
        }
    }
	if a.pw != nil {
        if err := a.pw.Stop(); err != nil {
            errs = append(errs, fmt.Errorf("pw stop error: %w", err))
        }
    }
	if len(errs) > 0 {
        return fmt.Errorf("errors during cleanup: %v", errs)
    }
    return nil
}