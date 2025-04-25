package routeHandlers

import (
	"context"
	"fmt"
	"os"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sashabaranov/go-openai"
)

 type Prompt struct{
	Prompt 		string 	`json:"prompt" binding:"required"`
	Temperature float32 `json:"temperature"`
 }

var (
	ResponseModels []string = []string{}
	modelMutex sync.Mutex
)

func HealthCheck(ctx *gin.Context)  {
	ctx.JSON(200,gin.H{"message":"Healthy Server."})
}

func Models(ctx *gin.Context){
	modelMutex.Lock()
	defer modelMutex.Unlock()

	if len(ResponseModels) > 0{
		ctx.JSON(200,gin.H{"data":ResponseModels})
		return
	}
	godotenv.Load()
	apiKey := os.Getenv("API_KEY")
	client := openai.NewClient(apiKey)

	models,err := client.ListModels(context.Background())
	if err != nil {
		fmt.Println(err)
		ctx.JSON(500,gin.H{"message":"Could not get OpenAI models."})
	}

	for _,model := range models.Models{
		ResponseModels = append(ResponseModels,model.ID)
	}

	ctx.JSON(200,gin.H{"data":ResponseModels})
}

func Completion(ctx *gin.Context){
	godotenv.Load()
	apiKey := os.Getenv("API_KEY")

	reqBody := &Prompt{}

	if err := ctx.ShouldBindJSON(reqBody); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid request", "details": err.Error()})
		return
	}

	client := openai.NewClient(apiKey)

	resp, err := client.CreateCompletion(
		context.Background(),
		openai.CompletionRequest{
			Model:     openai.GPT3Dot5TurboInstruct, 
			Prompt:    reqBody.Prompt,
			MaxTokens: 100,
			Temperature: reqBody.Temperature,
		},
	)
	
	if err != nil {
		fmt.Printf("Completion error: %v\n", err)
		return
	}
	
	ctx.JSON(200,gin.H{"data":resp.Choices[0].Text})
}

var conversationHistory = []openai.ChatCompletionMessage{}

func Chat(ctx *gin.Context) {
	fmt.Println(conversationHistory)
	godotenv.Load()
	apiKey := os.Getenv("API_KEY")

	client := openai.NewClient(apiKey)
	reqBody := &Prompt{}

	if err := ctx.ShouldBindJSON(reqBody); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid request", "details": err.Error()})
		return
	}

	conversationHistory = append(conversationHistory, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: reqBody.Prompt,
	})

	resp, err := client.CreateChatCompletion(context.Background(), openai.ChatCompletionRequest{
		Model:       openai.GPT3Dot5Turbo,
		Temperature: 0.7,
		MaxTokens:   400,
		Messages:    conversationHistory,
	})

	if err != nil {
		fmt.Printf("Completion error: %v\n", err)
		ctx.JSON(500, gin.H{"error": "Failed to get response from OpenAI"})
		return
	}

	assistantMsg := resp.Choices[0].Message

	conversationHistory = append(conversationHistory, assistantMsg)

	ctx.JSON(200, gin.H{"data": assistantMsg.Content})
}
