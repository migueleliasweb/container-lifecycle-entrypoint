package cmd

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
		log.Info("cmd: ", subCmd)
		log.Info("cmdARGS: ", viper.GetStringSlice("args"))
	},
}

func init() {
	viper.SetEnvPrefix("CLE")
	viper.AutomaticEnv()

	//root flags

	rootCmd.Flags().StringSliceVarP(
		&subCmdArgs,
		"args",
		"",
		[]string{},
		"Underlying cmd args",
	)

	rootCmd.Flags().StringVar(
		&config,
		"config",
		"",
		"config file (default is \"./.cle.yaml\")",
	)

	//viper binds
	viper.BindPFlag("args", rootCmd.Flags().Lookup("args"))
	viper.BindPFlag("config", rootCmd.Flags().Lookup("config"))

	fmt.Println("FMT: ", viper.GetString("config"))
}

//Run Cobras entrypoint
func Run() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
