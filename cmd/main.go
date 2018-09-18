package coflies

import "errors"

// Runner The main entry point of the cmd
type Runner interface {
	Start() error
	// Wait util it's done
	Wait() (ResultData, error)
	// Cancel/kill process if running
	Cancel() bool
	// Verify process is running or not
	IsRunning() bool
}

// Configuration properties needed for running
type Configuration struct {
	// Language - this will be set as default in each specific docker container
	Lang LanguageProperties
	// Project - this will be set as default in each specific docker container
	Project ProjectProperties
	// User input
	Implementation CodeData
	// User input
	Testing CodeData
	// User input
	TestData TestData
}

// LanguageProperties ...
type LanguageProperties struct {
	// Standard name of language. Ex: java, c, c++, golang, kotlin, python2, python3
	Name string
	// Language Implementation Version
	Version string
}

// ProjectProperties ...
type ProjectProperties struct{}

// CodeData properties
type CodeData struct{}

// TestData properties
type TestData struct{}

type ResultData struct{}

// MakeRunner Base on configuration return a specific runner
func MakeRunner(conf Configuration) (Runner, error) {
	// TODO configuration embedded into runner when initialization
	return new(golangRunner), nil
}

// Sample for Runner
type golangRunner struct {
}

func (r golangRunner) Start() error {
	return errors.New("Error")
}

func (r golangRunner) Wait() (ResultData, error) {
	return ResultData{}, nil
}

func (r golangRunner) Cancel() bool {
	return true
}

func (r golangRunner) IsRunning() bool {
	return false
}
