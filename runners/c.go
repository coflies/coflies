package runners

import (
	"github.com/coflies/coflies/common"
	log "github.com/sirupsen/logrus"
)

type c struct {
	instance *common.Runner
	args     []string
	// many more for c - specific
}

// Start the runner
func (r c) Start() error {
	log.Info("Starting the " + r.instance.Lang.Name + " runner")
	r.args = []string{}
	args := append(r.args, r.instance.Project.Args...)
	return r.instance.Start(args)
}

// Wait until the runner finished and return data
func (r c) Wait() (common.ResultData, error) {
	return r.instance.Wait()
}

// IsRunning check runner running or not
func (r c) IsRunning() bool {
	return false
}
