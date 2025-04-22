package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
	"github.com/sashabaranov/go-openai"
	"github.com/spf13/cobra"
)

var Model 		string
var Temperature string

var RootCmd = &cobra.Command{
	Use: "cli-ai",
	Args: cobra.MinimumNArgs(1),
	Run: func (cmd *cobra.Command,args []string)  {
		prompt := strings.Join(args," ")
		fmt.Fprintf(cmd.OutOrStdout(),"Generating response for prompt %s with temperature %s and model %s. \n",prompt,Temperature,Model)

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
				Content: prompt,
			},
		},
		MaxTokens: 100,
		},
	)
		if err != nil {
			fmt.Printf("ChatCompletion error: %v\n", err)
		}
		fmt.Println(resp.Choices[0].Message.Content)
	},
}

func init(){
	RootCmd.Flags().StringVarP(&Model,"model","m","","Chat GPT AI model")
	RootCmd.Flags().StringVarP(&Temperature,"temperature","t","","Chat GPT temperature")
	RootCmd.MarkFlagRequired("model")
	RootCmd.MarkFlagRequired("temperature")
}


func Execute()error{
if err := RootCmd.Execute();err != nil{
	fmt.Println(err)
		os.Exit(1)
		return err
}
return nil
}