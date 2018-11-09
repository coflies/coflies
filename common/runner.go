package common

// Runner The main entry point of the cmd
type Runner interface {
	Start() error
	// Wait util it's done
	Wait() (ResultData, error)
	// Verify process is running or not
	IsRunning() bool
}
