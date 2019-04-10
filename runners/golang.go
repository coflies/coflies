package runners

import (
	"bytes"
	"errors"
	"io"

	"github.com/coflies/coflies/common"
	log "github.com/sirupsen/logrus"
)

type Golang struct {
	instance *common.Runner
	// many more
}

func (r Golang) initOutputBuffers(stdout io.ReadCloser, stderr io.ReadCloser) {
	r.instance.StandardOutput = new(bytes.Buffer)
	r.instance.StandardOutput.ReadFrom(stdout)

	r.instance.ErrorOutput = new(bytes.Buffer)
	r.instance.ErrorOutput.ReadFrom(stderr)
}

// Start the Cpp
func (r Golang) Start() error {
	log.Info("we not implement this language")

	return errors.New("Not yet implemented")
	// log.Info("Starting the " + r.instance.Lang.Name + " runner")
	// // TODO getting project information
	// // TODO generate/update implement code base on CodeData
	// // TODO generate/update test code base on CodeData
	// // TODO prepare/loading test data base on TestData
	// // TODO run go test with cancel timeout and getting output
	// ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Millisecond)
	// defer cancel()

	// r.instance.Cmd = exec.CommandContext(ctx, r.instance.Lang.Name, r.instance.Project.Args...)
	// stdout, stderr, err := WireOutput(r.instance.Cmd)
	// if err != nil {
	// 	log.Fatal(err)
	// 	return err
	// }

	// if err := r.instance.Cmd.Start(); err != nil {
	// 	log.Fatal(err)
	// 	return err
	// }

	// r.initOutputBuffers(stdout, stderr)

	// return nil
}

// Wait until the runner finished and return data
func (r Golang) Wait() (common.ResultData, error) {
	log.Info("we not implement this language")

	return common.ResultData{}, errors.New("Not yet implemented")
	// log.Info("Waiting until runner completed...")
	// if err := r.instance.Cmd.Wait(); err != nil {
	// 	log.Fatal(err)
	// 	return common.ResultData{}, err
	// }

	// log.Info("Completed.")
	// log.Info("Standard output: ", r.instance.StandardOutput.String())
	// log.Info("Error output: ", r.instance.ErrorOutput.String())

	// // TODO resultHandler for parsing/beauty/validate output
	// return common.ResultData{}, nil
}

// IsRunning check runner running or not
func (r Golang) IsRunning() bool {
	return false
}
