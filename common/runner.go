package common

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"os/exec"
	"time"

	log "github.com/sirupsen/logrus"
)

// Runner The main entry point of the cmd
type Runable interface {
	Start() error
	// Wait util it's done
	Wait() (ResultData, error)
	// Verify process is running or not
	IsRunning() bool
}

type Runner struct {
	Lang LanguageData

	Project        ProjectData
	StandardOutput *bytes.Buffer
	ErrorOutput    *bytes.Buffer
	cmd            *exec.Cmd
}

func (r *Runner) wireOutput() (io.ReadCloser, io.ReadCloser, error) {
	stdout, err := r.cmd.StdoutPipe()
	if err != nil {
		return nil, nil, err
	}
	stderr, err := r.cmd.StderrPipe()
	if err != nil {
		return nil, nil, err
	}

	return stdout, stderr, nil
}

func (r *Runner) bindOutput(stdout io.ReadCloser, stderr io.ReadCloser) {
	r.StandardOutput = new(bytes.Buffer)
	r.ErrorOutput = new(bytes.Buffer)

	r.StandardOutput.ReadFrom(stdout)
	r.ErrorOutput.ReadFrom(stderr)
}

func compileSource(r Runner) error {
	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Millisecond)
	defer cancel()

	log.Infof("%v", r.Lang.CompileArgs)
	compileCmd := exec.CommandContext(ctx, r.Lang.CompilerName, r.Lang.CompileArgs...)
	compileCmd.Dir = r.Project.Workspace
	stdout, err := compileCmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
		return err
	}
	stderr, err := compileCmd.StderrPipe()
	if err != nil {
		return fmt.Errorf("compile result: %s", r.ErrorOutput.String())
	}
	if err = compileCmd.Start(); err != nil {
		return fmt.Errorf("compile result: %s", r.ErrorOutput.String())
	}

	r.bindOutput(stdout, stderr)
	if err = compileCmd.Wait(); err != nil {
		return fmt.Errorf("compile result: %s", r.ErrorOutput.String())
	}

	log.Infof("compile result: %s", r.StandardOutput.String())
	return nil
}

func (r *Runner) Start(args []string) error {
	if r.Lang.Type == "compiler" {
		log.Infof("compiling code before run...")
		// we need compile first
		if err := compileSource(*r); err != nil {
			return err
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Millisecond)
	defer cancel()
	r.cmd = exec.CommandContext(ctx, r.Project.Workspace+"/"+r.Lang.Exec, args...)
	r.cmd.Dir = r.Project.Workspace
	stdout, stderr, err := r.wireOutput()
	if err != nil {
		log.Fatal(err)
		return err
	}

	if err := r.cmd.Start(); err != nil {
		log.Fatalf("Error start command. %v", err)
		return err
	}

	r.bindOutput(stdout, stderr)

	return nil
}

func (r *Runner) Wait() (ResultData, error) {
	//
	if r.cmd == nil {
		return ResultData{}, errors.New("runner stopped")
	}

	if err := r.cmd.Wait(); err != nil {
		return ResultData{
			Stderr: r.ErrorOutput.String(),
		}, err
	}

	return ResultData{
		Stdout: r.StandardOutput.String(),
		Stderr: r.ErrorOutput.String(),
	}, nil
}
