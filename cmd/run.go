package cmd

import (
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
			langData := common.MakeLanguage("/home/tntvu/.coflies", language, version)
			config := common.Configuration{
				Lang:           langData,
				Project:        common.ProjectData{},
				Implementation: common.CodeData{},
				Testing:        common.CodeData{},
				TestData:       common.TestData{},
			}
			codeRunner, _ := runners.MakeRunner(config.Lang)
			// log.Info("  - *(optional - not in priority) preparing language framework IF neccessary...")
			// log.Info("    -- download if not existed")
			// log.Info("        download: " + config.Lang.DownloadLink(gi.GoOS, gi.Platform, "tar.gz"))
			// log.Info("      --- extract if downloaded")
			// log.Info("      --- configure language")
			log.Info("  - preparing workspace ")
			log.Info("    -- create folders IF neccessary")                                                      // design adaptable/extendable folders structure for all languages we will support
			log.Info("    -- download/extract/copy necessary 3rds IF neccessary")                                // TODO
			log.Info("    -- create/copy/move project files, test files, data files into folders IF neccessary") // should reuse frameworks of old codewars
			log.Info("2. create/copy/move code files folders")                                                   // generate code files based on submission - old? or let user submit full source file
			log.Info("3. run code files")                                                                        // multiple threads? async mode?
			codeRunner.Start()
			log.Info("4. received and parse results") // multiple threads? async mode?
			results, _ := codeRunner.Wait()
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
