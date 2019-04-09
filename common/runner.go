package common

import (
	"bytes"
	"os/exec"
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
	Cmd            *exec.Cmd
}
