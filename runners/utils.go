package runners

import (
	"errors"
	"io"
	"os/exec"

	common "github.com/coflies/coflies/common"
)

func WireOutput(cmd *exec.Cmd) (io.ReadCloser, io.ReadCloser, error) {
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

func MakeRunner(lang common.LanguageData) (common.Runable, error) {
	instance := &common.Runner{
		Lang: lang,
	}
	switch lang.Name {
	case "cpp":
		return cpp{instance: instance}, nil
	case "go":
		return Golang{instance: instance}, nil
	case "c":
		return c{instance: instance}, nil
	default:
		return nil, errors.New("language runner not implemented\nplease contact developer for more information")
	}
}
