package runners

import (
	"errors"

	common "github.com/coflies/coflies/common"
)

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
