package main

import (
	"fmt"

	"github.com/coflies/coflies/cmd"

	common "github.com/coflies/coflies/common"
	"github.com/coflies/coflies/runners"
	log "github.com/sirupsen/logrus"
)

func main() {
	cmd.Execute()
	// TODO additional HTTP server instead of cmd mode
	// // TODO use https/2 receive/return output
	// defer func() {
	// 	log.Info("Coflies shutdown.")
	// }()

	// log.Info("Starting Coflies server")

	// // TODO query available languages on this host
	// log.Info("Querying available languages")

	// // TODO getting languages and runners on ex-server if neccessary.

	//
	// TODO agruments parsing
	// TODO    - language initialize
	// TODO    - project initialize
	// TODO    - code data initialize
	// TODO    - test data initialize
	// TODO configuration initialize
	config := common.Configuration{
		Lang:           common.LanguageData{Name: "golang", Version: "1.11"},
		Project:        common.ProjectData{},
		Implementation: common.CodeData{},
		Testing:        common.CodeData{},
		TestData:       common.TestData{},
	}

	// TODO runner initialize
	runner, _ := MakeRunner(config)
	runner.Start()
	result, err := runner.Wait()
	if err != nil {
		log.Fatal(err)
		fmt.Println("Runner error. Details: ", err.Error())
		return
	}

	// hahaha
	fmt.Println(result)
}

func printUsage() {
	fmt.Println(`Code Sandbox CLI Runner
	Usage: coflies`)
}

// MakeRunner Base on configuration return a specific runner
func MakeRunner(conf common.Configuration) (common.Runner, error) {
	// TODO configuration embedded into runner when initialization
	// TODO create runner base on conf.Lang
	return new(runners.Golang), nil
}
