package cmd

import (
	"os"
	"path/filepath"
	"pod-lifecycle-entrypoint/cle"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	// jww "github.com/spf13/jwalterweatherman"
)

var config string
var subCmd string
var subCmdArgs []string

var rootCmd = &cobra.Command{
	Use:   "cle",
	Short: "Container lifecycle entrypoint",
	Long:  "Sends probes according to the underlying cmd lifecycle state",
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("config: ", viper.GetString("config"))
		log.Info("cmd: ", viper.GetString("cmd"))
		log.Info("cmdARGS: ", viper.GetStringSlice("args"))

		cle.ExecCMD(
			viper.GetString("cmd"),
			viper.GetStringSlice("args"),
		)
	},
}

func init() {
	// //very useful to troubleshooting stuff with cobra+viper
	// jww.SetStdoutThreshold(jww.LevelDebug)

	viper.SetEnvPrefix("CLE")
	viper.AutomaticEnv()

	//viper root flags
	rootCmd.Flags().StringVar(
		&config,
		"config",
		".cle_config",
		"config file (default is \"./.cle_config.yaml\")",
	)

	rootCmd.Flags().StringVarP(
		&subCmd,
		"cmd",
		"",
		"",
		"Underlying cmd to exec",
	)

	rootCmd.Flags().StringSliceVarP(
		&subCmdArgs,
		"args",
		"",
		[]string{},
		"Underlying cmd args",
	)

	//viper binds
	viper.BindPFlag("cmd", rootCmd.Flags().Lookup("cmd"))
	viper.BindPFlag("args", rootCmd.Flags().Lookup("args"))
	viper.BindPFlag("config", rootCmd.Flags().Lookup("config"))

	//setup configuration file
	viper.AddConfigPath("/")
	viper.AddConfigPath("/etc/cle")
	if pwd, err := os.Getwd(); err == nil {
		viper.AddConfigPath(pwd)
	}
	cobra.OnInitialize(func() {
		configFile := viper.GetString("config")
		extension := filepath.Ext(configFile)
		if extension != "" {
			configFile = configFile[0 : len(configFile)-len(extension)]
			viper.Set("config", configFile) //fixes the value to future calls
		}
		viper.SetConfigName(viper.GetString("config"))
		viper.ReadInConfig()
	})
}

//RunCli Cobras entrypoint
func RunCli() {
	if err := rootCmd.Execute(); err != nil {
		log.Error(err)
		os.Exit(1)
	}
}
