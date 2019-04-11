package runners

import (
	"github.com/coflies/coflies/common"
	log "github.com/sirupsen/logrus"
)

type cpp struct {
	instance *common.Runner
	args     []string
	// many more for c++ specific
}

// Start the runner
func (r cpp) Start() error {
	log.Info("Starting the " + r.instance.Lang.Name + " runner")
	r.args = []string{}
	args := append(r.args, r.instance.Project.Args...)
	return r.instance.Start(args)
}

// Wait until the runner finished and return data
func (r cpp) Wait() (common.ResultData, error) {
	return r.instance.Wait()
}

// IsRunning check runner running or not
func (r cpp) IsRunning() bool {
	return false
}
