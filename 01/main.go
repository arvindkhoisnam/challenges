package main

import (
	"fmt"

	"github.com/arvindkhoisnam/challanges/01/command"
)



func Greet()string{
	return "Hello World"
}
func main(){
	if err := command.Execute(); err != nil {
		fmt.Println(err)
	}
}


