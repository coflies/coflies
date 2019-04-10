package runners

import (
	"bytes"
	"context"
	"errors"
	"io"
	"os/exec"
	"time"

	"github.com/coflies/coflies/common"
	log "github.com/sirupsen/logrus"
)

type cpp struct {
	instance *common.Runner
	// many more
	args []string
}

func (r *cpp) initOutputBuffers(stdout io.ReadCloser, stderr io.ReadCloser) {
	r.instance.StandardOutput = new(bytes.Buffer)
	r.instance.StandardOutput.ReadFrom(stdout)

	r.instance.ErrorOutput = new(bytes.Buffer)
	r.instance.ErrorOutput.ReadFrom(stderr)
}

// Start the Cpp
func (r cpp) Start() error {
	log.Info("Starting the " + r.instance.Lang.Name + " runner")
	// // TODO getting project information
	// // TODO generate/update implement code base on CodeData
	// // TODO generate/update test code base on CodeData
	// // TODO prepare/loading test data base on TestData
	// // TODO run go test with cancel timeout and getting output
	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Millisecond)
	defer cancel()

	r.args = []string{"--version"}
	args := append(r.args, r.instance.Project.Args...)
	r.instance.Cmd = exec.CommandContext(ctx, r.instance.Lang.CompilerName, args...)
	stdout, stderr, err := WireOutput(r.instance.Cmd)
	if err != nil {
		log.Fatal(err)
		return err
	}

	if err := r.instance.Cmd.Start(); err != nil {
		log.Fatalf("Error start command. %v", err)
		return err
	}

	r.initOutputBuffers(stdout, stderr)

	return nil
}

// Wait until the runner finished and return data
func (r cpp) Wait() (common.ResultData, error) {
	//
	if r.instance.Cmd == nil {
		return common.ResultData{}, errors.New("The command is not existed")
	}

	if err := r.instance.Cmd.Wait(); err != nil {
		log.Fatalf("Error waiting command. %v", err)
		return common.ResultData{}, err
	}

	// TODO resultHandler for parsing/beauty/validate output
	result := common.ResultData{
		Stdout: r.instance.StandardOutput.String(),
		Stderr: r.instance.ErrorOutput.String(),
	}

	return result, nil
}

// IsRunning check runner running or not
func (r cpp) IsRunning() bool {
	return false
}
