package command

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/google/shlex"
	"github.com/joho/godotenv"
	"github.com/sashabaranov/go-openai"
	"github.com/spf13/cobra"
)

var Model string
var Temperature string
var Prompt string

var supportedModels [3]string = [3]string{"gpt-3.5-turbo","gpt-4","gpt-4-turbo"}

var RootCmd = &cobra.Command{
	Use:   "cli-ai",
	Short: "A CLI AI TOOL",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println()
		fmt.Println("Welcome to cli-ai 👋🤖.")
		fmt.Println("---- IMPORTANT INFO ----")
		fmt.Println("⚠️  Use command 'gen' to generate an ai response.")
		fmt.Println("⚠️  Desired OpenAI model is a required flag. -m GPT 3.0")
		fmt.Println("✅  Currently supported models",supportedModels)
		fmt.Println("⚠️  Desired OpenAI temperature is a required flag. -t 0.7")
		fmt.Println("⚠️  Your prompt is a required flag. -p Please explain Quantum computing")
		fmt.Println("⚠️  Final command should look like gen -t 0.7 -m gpt-4 Please explain Quantum computing")
		fmt.Println()
		fmt.Println("⛔️ Type 'exit' to quit.")
		fmt.Println()

		reader := bufio.NewReader(os.Stdin)
		
		for {
			fmt.Print("cli-ai 🤖: ")
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)
			if input == "exit"{
				break
			}
			// args := strings.Fields(input)
			args, err := shlex.Split(input)
			if err != nil {
				fmt.Println("❌ Error parsing input:", err)
				continue
			}
			if len(args) == 0 {
				continue
			}
			if args[0] != "gen"{
				fmt.Println("❌❌WRONG INPUT. PLEASE READ THE INSTRUCTIONS CAREFULLY.❌❌")
				continue
			}
			if args[0] == "gen"{
				newGenerate := *generate
				newGenerate.RunE = func(cmd *cobra.Command, args []string) error {

					if exists := includes(supportedModels,Model); !exists {
						return fmt.Errorf("❌❌WRONG INPUT: %s model is not supported.❌❌",Model)
					}

					err := godotenv.Load()
					if err != nil {
						log.Fatal("Error loading .env file")
					}
					temperature,err := strconv.ParseFloat(Temperature,64)

					if err != nil{
						fmt.Println(err)
					}
					apiKey := os.Getenv("API_KEY")
					client := openai.NewClient(apiKey)
					resp, err := client.CreateChatCompletion(
					context.Background(),
					openai.ChatCompletionRequest{
					Model: Model,
					Temperature: float32(temperature),
					Messages: []openai.ChatCompletionMessage{
						{
							Role:    openai.ChatMessageRoleUser,
							Content: Prompt,
						},
					},
					MaxTokens: 100,
				},
			)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		// return 
	}

				fmt.Println(resp.Choices[0].Message.Content)
					return nil
				}
				
				 // Execute the command properly
				 newCmd := &cobra.Command{}
				 newCmd.AddCommand(&newGenerate)
				 newCmd.SetArgs(args)
				 newCmd.Execute()
			}
		}
	  },
}


func init(){
	RootCmd.AddCommand(generate)
	generate.Flags().StringVarP(&Model,"model","m","","Your OpenAI desired model.")
	generate.Flags().StringVarP(&Temperature,"temp","t","","Your OpenAI desired temperature.")
	generate.Flags().StringVarP(&Prompt,"prompt","p","","Your prompt.")
	generate.MarkFlagRequired("model")
	generate.MarkFlagRequired("temp")
	generate.MarkFlagRequired("prompt")
}


func Execute()error {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
		return err
	}
	return nil
}

var generate = &cobra.Command{
Use: "gen",
}

func includes(slice [3]string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}