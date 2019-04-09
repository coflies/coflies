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
	switch lang.Name {
	case "cpp":
		return Cpp{
			instance: common.Runner{
				Lang: lang,
			},
		}, nil
	case "go":
		return Golang{
			instance: common.Runner{
				Lang: lang,
			},
		}, nil
	default:
		return nil, errors.New("Language not supported")
	}
}
