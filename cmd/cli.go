package cmd

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/altsrc"
	"gopkg.in/urfave/cli.v1"
)

//ConfigureCli Configures CLI flags
func ConfigureCli() {

	app := cli.NewApp()

	app.Name = "Pod lifecycle entrypoint"
	app.Description = "Sends probes according to the underlying cmd lifecycle"
	app.Version = "v0.1"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "pl-probe-name",
			Value:  "HTTPProbe",
			Usage:  "Probe to use when doing healthchecks",
			EnvVar: "PL_PROBE_NAME",
		},
		cli.StringFlag{
			Name:   "pl-cmd",
			Usage:  "Underlying command to call after the heathcheck passes",
			EnvVar: "PL_CMD",
		},
		cli.StringSliceFlag{
			Name:   "pl-args",
			Usage:  "Args to pass to the underlying cmd",
			EnvVar: "PL_ARGS",
		},
		cli.StringFlag{
			Name:   "config",
			Value:  ".pl_config.yaml",
			Usage:  "Path to config",
			EnvVar: "PL_CONFIG",
		},
	}

	app.Before = altsrc.InitInputSourceWithContext(
		app.Flags,
		altsrc.NewYamlSourceFromFlagFunc("config"),
	)

	app.Action = Run

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

//Run Runs the app
func Run(c *cli.Context) error {
	log.Println(c.StringSlice("pl-args"))
	return nil
}
