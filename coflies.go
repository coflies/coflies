package main

import (
	"fmt"

	coflies "github.com/coflies/coflies/cmd"
)

func main() {
	// TODO agruments parsing
	// TODO    - language initialize
	// TODO    - project initialize
	// TODO    - code data initialize
	// TODO    - test data initialize
	// TODO configuration initialize
	config := coflies.Configuration{
		coflies.LanguageProperties{"golang", "1.11"},
		coflies.ProjectProperties{},
		coflies.CodeData{},
		coflies.CodeData{},
		coflies.TestData{},
	}

	// TODO runner initialize
	runner, _ := coflies.MakeRunner(config)
	runner.Start()
	result, _ := runner.Wait()

	// hahaha
	fmt.Println(result)
}

func printUsage() {
	fmt.Println(`Code Sandbox CLI Runner
	Usage: coflies`)
}
