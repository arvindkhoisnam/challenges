## Challenge 1: Getting Started with Go and OpenAI

This challenge has two approaches. The first approach is the one given in the original challenge and the second approach is a more advanced version.

## Requirements

1. You will need your own OpenAI API key. If you do not have it already, get it from [text](https://platform.openai.com/account/api-keys).
2. Create a .env file in the root folder. For reference, a .env.example file is given.

## First approach

- By default the main package is configured to run the first approach.
- You will need to run the command "go run main.go".
- Pass in the following flags
  -m or --model and choose a model of your choice gpt-3.5-turbo, gpt-4, gpt-4-turbo
  -t or --temperature and choose and choose a number between 0.1 to 0.9
- And finally pass in your prompt within double inverted commas.
- Final input should look like: go run main.go -t 0.4 -m gpt-4 "Tell me about quantum computing"

## Second approach

- This is more like an interactive app.
- The app does not exit out, instead it keeps on listening for prompts.
- Steps :
  1. Go to main.go and uncomment //"github.com/arvindkhoisnam/challanges/01/command"
  2. In the main func change cmd.Execute() to command.Execute()
  3. Run go run main.go
  4. Follow instructions.
