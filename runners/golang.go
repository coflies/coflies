package runners

import (
	"errors"

	common "github.com/coflies/coflies/common"
	log "github.com/sirupsen/logrus"
)

// Golang runner
type Golang struct {
	Project common.ProjectData
}

// Start the runner
func (r Golang) Start() error {
	log.Info("Starting the golang runner")
	// TODO getting project information
	// TODO generate/update implement code base on CodeData
	// TODO generate/update test code base on CodeData
	// TODO prepare/loading test data base on TestData
	// TODO run go test and getting output
	// TODO create ResultData base on output
	// TODO
	return errors.New("Error")
}

// Wait until the runner finished and return data
func (r Golang) Wait() (common.ResultData, error) {
	log.Info("Waiting until runner completed...")
	return common.ResultData{}, nil
}

// Cancel runner if it is runnint
func (r Golang) Cancel() bool {
	return true
}

// IsRunning check runner running or not
func (r Golang) IsRunning() bool {
	return false
}
