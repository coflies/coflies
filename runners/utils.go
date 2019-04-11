package runners

import (
	"errors"
	"io"
	"os"

	common "github.com/coflies/coflies/common"
)

// copyFileContents copies the contents of the file named src to the file named
// by dst. The file will be created if it does not already exist. If the
// destination file exists, all it's contents will be replaced by the contents
// of the source file.
func CopyFileContents(src, dst string) (err error) {
	in, err := os.Open(src)
	if err != nil {
		return
	}
	defer in.Close()
	out, err := os.Create(dst)
	if err != nil {
		return
	}
	defer func() {
		cerr := out.Close()
		if err == nil {
			err = cerr
		}
	}()
	if _, err = io.Copy(out, in); err != nil {
		return
	}
	err = out.Sync()
	return
}

// MakeRunner create wrapper type for common Runner instance
//            this is useful when we need add more logic into
//            specific language
func MakeRunner(config common.Configuration) (common.Runable, error) {
	instance := &common.Runner{
		Lang:    config.Lang,
		Project: config.Project,
	}
	switch config.Lang.Name {
	case "cpp":
		return cpp{instance: instance}, nil
	case "go":
		return golang{instance: instance}, nil
	case "c":
		return c{instance: instance}, nil
	default:
		return nil, errors.New("language runner not implemented\nplease contact developer for more information")
	}
}
