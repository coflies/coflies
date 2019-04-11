package common

import "errors"

func (l *LanguageData) DownloadLink(os string, arch string, ext string) string {
	//
	switch l.Name {
	// usually we don't support c/cpp installation - user should download and install by themselves
	case "cpp":
		return ""
	case "go":
		return "go" + l.Version + "." + os + "-" + arch + "." + ext
	default:
		return ""
	}
}

func (l *LanguageData) SupportedVersion() ([]string, error) {
	// usually we don't support installation - user should download and install by themselves
	switch l.Name {
	case "cpp", "c":
		return []string{"7.0.1"}, nil
	case "go":
		return []string{"1.12.1"}, nil
	default:
		return nil, errors.New("Not supported")
	}
}
