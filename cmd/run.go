package cmd

import (
	"os"
	"os/exec"
	"strconv"

	"github.com/coflies/coflies/common"
	"github.com/coflies/coflies/runners"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/matishsiao/goInfo"
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
			log.Info("Language: ", language)
			log.Info("Version: ", version)

			log.Info("1. checking and verify system compabilities...")
			log.Info("  - verify os and arch information...")
			gi := goInfo.GetInfo()
			log.Info("    OS: " + gi.GoOS + ", Platform: " + gi.Platform + ", CPUs: " + strconv.Itoa(gi.CPUs))
			log.Info("  - initialize the runner for os and language")
			//
			langData, err := common.MakeLanguage("/home/tntvu/.coflies", language, version)
			if err != nil {
				log.Fatal(err)
				os.Exit(1)
			}
			log.Info("  - preparing workspace ")
			log.Info("    -- create folders IF neccessary") // design adaptable/extendable folders structure for all languages we will support
			projectData, err := common.MakeProject("/home/tntvu/repo", language)
			if err != nil {
				log.Fatal(err)
				os.Exit(1)
			}
			log.Info("    -- download/extract/copy necessary 3rds IF neccessary")                                // TODO
			log.Info("    -- create/copy/move project files, test files, data files into folders IF neccessary") // should reuse frameworks of old codewars
			log.Info("2. create/copy/move code files folders")                                                   // generate code files based on submission - old? or let user submit full source file
			config := common.Configuration{
				Lang:           langData,
				Project:        projectData,
				Implementation: common.CodeData{},
				Testing:        common.CodeData{},
				TestData:       common.TestData{},
			}
			codeRunner, err := runners.MakeRunner(config)
			if err != nil {
				log.Fatal(err)
				os.Exit(1)
			}
			log.Info("3. run code files") // multiple threads? async mode?
			if err = codeRunner.Start(); err != nil {
				log.Fatal(err)
				os.Exit(1)
			}
			log.Info("4. received and parse results") // multiple threads? async mode?
			results, err := codeRunner.Wait()
			if err != nil {
				log.Fatal(err)
				os.Exit(1)
			}
			log.Info("5. print out results") // multiple threads? async mode?
			log.Info(results)
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
