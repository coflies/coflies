package main

import (
	"fmt"

	cmd "github.com/coflies/coflies/cmd"
	"github.com/coflies/coflies/runners"
)

func main() {
	// TODO agruments parsing
	// TODO    - language initialize
	// TODO    - project initialize
	// TODO    - code data initialize
	// TODO    - test data initialize
	// TODO configuration initialize
	config := cmd.Configuration{
		cmd.LanguageProperties{"golang", "1.11"},
		cmd.ProjectProperties{},
		cmd.CodeData{},
		cmd.CodeData{},
		cmd.TestData{},
	}

	// TODO runner initialize
	runner, _ := MakeRunner(config)
	runner.Start()
	result, _ := runner.Wait()

	// hahaha
	fmt.Println(result)
}

func printUsage() {
	fmt.Println(`Code Sandbox CLI Runner
	Usage: coflies`)
}

// MakeRunner Base on configuration return a specific runner
func MakeRunner(conf cmd.Configuration) (cmd.Runner, error) {
	// TODO configuration embedded into runner when initialization
	// TODO create runner base on conf.Lang
	return new(runners.Golang), nil
}
