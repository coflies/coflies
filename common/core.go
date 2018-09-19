package common

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
	Lang LanguageData
	// Project - this will be set as default in each specific docker container
	Project ProjectData
	// User input
	Implementation CodeData
	// User input
	Testing CodeData
	// User input
	TestData TestData
}

// LanguageData ...
type LanguageData struct {
	// Standard name of language. Ex: java, c, c++, golang, kotlin, python2, python3
	Name string
	// Language Implementation Version
	Version string
}

// ProjectData store properties of the coflies project
type ProjectData struct{}

// CodeData store code data of the coflies project
type CodeData struct{}

// TestData store test data of the coflies project
type TestData struct{}

// ResultData store result of the coflies project
type ResultData struct{}
