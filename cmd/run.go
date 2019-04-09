package cmd

import (
	"fmt"

	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "running code",
	Long:  "Compile and/or Running your code and testcases using specific language and version.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("run called.", viper.GetBool("development"))
	},
}

var language = ""
var version = ""

func init() {
	runnerCmd.AddCommand(runCmd)
	// flags and configuration setting for run command
	runCmd.Flags().StringVarP(&language, "language", "l", "", "The programing language")
	runCmd.Flags().StringVarP(&version, "version", "v", "", "The programing language version. For getting supported versions of languages, please use the info command")

	runCmd.MarkFlagRequired("language")
}
