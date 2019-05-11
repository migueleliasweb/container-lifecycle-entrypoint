package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	// jww "github.com/spf13/jwalterweatherman"
)

//ConfigureCli2 Configures CLI flags
// func ConfigureCli2() {

// 	app := cli.NewApp()

// 	app.Name = "Pod lifecycle entrypoint"
// 	app.Description = "Sends probes according to the underlying cmd lifecycle"
// 	app.Version = "v0.1"

// 	app.Flags = []cli.Flag{
// 		cli.StringFlag{
// 			Name:   "pl-probe-name",
// 			Value:  "HTTPProbe",
// 			Usage:  "Probe to use when doing healthchecks",
// 			EnvVar: "PL_PROBE_NAME",
// 		},
// 		cli.StringFlag{
// 			Name:   "pl-cmd",
// 			Usage:  "Underlying command to call after the heathcheck passes",
// 			EnvVar: "PL_CMD",
// 		},
// 		cli.StringSliceFlag{
// 			Name:   "pl-args",
// 			Usage:  "Args to pass to the underlying cmd",
// 			EnvVar: "PL_ARGS",
// 		},
// 		cli.StringFlag{
// 			Name:   "config",
// 			Value:  ".pl_config.yaml",
// 			Usage:  "Path to config",
// 			EnvVar: "PL_CONFIG",
// 		},
// 	}

// 	app.Before = altsrc.InitInputSourceWithContext(
// 		app.Flags,
// 		altsrc.NewYamlSourceFromFlagFunc("config"),
// 	)
// }

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

//Run Cobras entrypoint
func Run() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
