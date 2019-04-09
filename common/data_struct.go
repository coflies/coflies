package common

type LanguageType int

// Language Type
const (
	COMPILER    LanguageType = iota
	INTERPRETER LanguageType = iota
)

// LanguageData ...
type LanguageData struct {
	// Standard name of language. Ex: java, c, c++, golang, kotlin, python2, python3
	Name string
	// Implementation Version
	Version string
	// relative directory install path
	Path string
	// Execution name of exec binary. Ex: java
	ExecName string
	// Compile name of compile binary. Ex: javac
	CompileName string
	// Type: compiler (need compile then run) | interpreter (don't need compile)
	Type LanguageType
}

// ProjectData store properties of the coflies project
type ProjectData struct {
	Args []string
	// working space path
	Workspace string
	//
}

// CodeData store code data of the coflies project
type CodeData struct{}

// TestData store test data of the coflies project
type TestData struct{}

// ResultData store result of the coflies project
type ResultData struct{}
