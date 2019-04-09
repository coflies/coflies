package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

func isCommandAvailable(name string) bool {
	cmd := exec.Command("/bin/sh", "-c", "command -v "+name)
	if err := cmd.Run(); err != nil {
		return false
	}
	return true
}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "running code",
	Long:  "Compile and/or Running your code and testcases using specific language and version.",
	Run: func(cmd *cobra.Command, args []string) {
		// Info - only print when debugging/development
		if viper.GetBool("development") {
			fmt.Println("Language: ", viper.GetString("language"))
			fmt.Println("Version: ", viper.GetString("version"))

			fmt.Println("checking and verify system compabilities...")
			fmt.Println("- verify os and arch information...")
			fmt.Println("- preparing language framework IF neccessary...")
			fmt.Println("-- download if not existed")
			fmt.Println("--- extract if downloaded")
			fmt.Println("--- configure language")
			fmt.Println("- preparing workspace ")
			fmt.Println("-- create folders IF neccessary")                                                      // design adaptable/extendable folders structure for all languages we will support
			fmt.Println("-- download/extract/copy necessary 3rds IF neccessary")                                // TODO
			fmt.Println("-- create/copy/move project files, test files, data files into folders IF neccessary") // should reuse frameworks of old codewars
			fmt.Println("create/copy/move code files folders")                                                  // generate code files based on submission - old? or let user submit full source file
			fmt.Println("run code files")                                                                       // multiple threads? async mode?
			fmt.Println("process and parse results")                                                            // multiple threads? async mode?
		}
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
