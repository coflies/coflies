package runners

import (
	"bytes"
	"io"
	"os/exec"

	common "github.com/coflies/coflies/common"
	log "github.com/sirupsen/logrus"
)

// Golang runner
type Golang struct {
	Project        common.ProjectData
	StandardOutput *bytes.Buffer
	ErrorOutput    *bytes.Buffer
	Cmd            *exec.Cmd
}

// Start the runner
func (r *Golang) Start() error {
	log.Info("Starting the golang runner")
	// TODO getting project information
	// TODO generate/update implement code base on CodeData
	// TODO generate/update test code base on CodeData
	// TODO prepare/loading test data base on TestData
	// TODO run go test and getting output
	r.Cmd = exec.Command("go", "test")
	stdout, stderr, err := wireOutput(r.Cmd)
	if err != nil {
		log.Fatal(err)
		return err
	}

	if err := r.Cmd.Start(); err != nil {
		log.Fatal(err)
		return err
	}

	r.initOutputBuffers(stdout, stderr)

	return nil
}

// Wait until the runner finished and return data
func (r *Golang) Wait() (common.ResultData, error) {
	log.Info("Waiting until runner completed...")
	if err := r.Cmd.Wait(); err != nil {
		log.Fatal(err)
		return common.ResultData{}, err
	}

	log.Info("Completed.")
	log.Info("Standard output: ", r.StandardOutput.String())
	log.Info("Error output: ", r.ErrorOutput.String())

	// TODO resultHandler for parsing/beauty/validate output

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

func wireOutput(cmd *exec.Cmd) (io.ReadCloser, io.ReadCloser, error) {
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, nil, err
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return nil, nil, err
	}

	return stdout, stderr, nil
}

func (r *Golang) initOutputBuffers(stdout io.ReadCloser, stderr io.ReadCloser) {
	r.StandardOutput = new(bytes.Buffer)
	r.StandardOutput.ReadFrom(stdout)

	r.ErrorOutput = new(bytes.Buffer)
	r.ErrorOutput.ReadFrom(stderr)
}
