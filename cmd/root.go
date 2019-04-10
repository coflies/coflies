package cmd

import (
	"os"

	homedir "github.com/mitchellh/go-homedir"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

//
var cfgFile string
var development bool

// The Response value from Execute
type Response struct {
	Result string
	Err    error
	Cmd    *cobra.Command
}

var runnerCmd = &cobra.Command{
	Use:   "coflies",
	Short: "Run all kind of code",
	Long: `A Fast and Flexible code runner built with love
  by tntvu and friends in Go.`,
}

func Execute() {
	if err := runnerCmd.Execute(); err != nil {
		log.Fatalf("%v", err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	runnerCmd.PersistentFlags().StringVar(
		&cfgFile,
		"config",
		"",
		"config file (default is $HOME/.coflies/config.yaml",
	)
	runnerCmd.PersistentFlags().BoolVarP(&development, "development", "d", false, "Enable development mode. Default is false")
	viper.BindPFlag("development", runnerCmd.PersistentFlags().Lookup("development"))
	viper.SetDefault("development", false)
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			log.Fatalf("%v", err)
			os.Exit(1)
		}
		viper.AddConfigPath(home + "/.coflies")
		viper.SetConfigName("config")
	}
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err == nil {
		if viper.GetBool("development") {
			log.Infof("Using config file: %s", viper.ConfigFileUsed())
		}
	}
}
