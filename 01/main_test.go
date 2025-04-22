package main

import (
	"bytes"
	"testing"

	"github.com/arvindkhoisnam/challanges/01/cmd"
)


func TestCLI(t *testing.T){
	var stdout bytes.Buffer

	cmd := cmd.RootCmd
	cmd.SetOut(&stdout)
	cmd.SetArgs([]string{"--model", "gpt-4","-t","0.7","Explain quantum computing."})
	err := cmd.Execute()
	if err != nil {
		t.Errorf("Unexpected error %v",err)
	}

	response := stdout.String()
	expected := "Generating response for prompt Explain quantum computing. with temperature 0.7 and model gpt-4. \n"

	if  response != expected {
		t.Errorf("Test failed. Expected output %q does not match actual response %q",expected,response)
	}
}
