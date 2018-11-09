package common

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
