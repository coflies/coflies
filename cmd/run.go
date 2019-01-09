package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "running code",
	Long:  "Running code.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("run called.", development)
	},
}

func init() {
	runnerCmd.AddCommand(runCmd)
	// flags and configuration setting for run command
}
