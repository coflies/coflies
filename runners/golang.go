package runners

import (
	"github.com/coflies/coflies/common"
	log "github.com/sirupsen/logrus"
)

type golang struct {
	instance *common.Runner
	args     []string
	// many more for golang specific
}

// Start the runner
func (r golang) Start() error {
	log.Info("Starting the " + r.instance.Lang.Name + " runner")

	// copy/generate source file
	// testing
	err := CopyFileContents("projects/"+r.instance.Lang.Name+"/coflies_impl.go", r.instance.Project.Workspace+"/coflies.go")
	if err != nil {
		log.Fatalf("error generate code file - %v", err)
		return err
	}

	args := append(r.args, r.instance.Project.Args...)
	return r.instance.Start(args)
}

// Wait until the runner finished and return data
func (r golang) Wait() (common.ResultData, error) {
	return r.instance.Wait()
}

// IsRunning check runner running or not
func (r golang) IsRunning() bool {
	return false
}
