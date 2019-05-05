package main

import (
	"pod-lifecycle-entrypoint/cmd"

	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.DebugLevel)
}

func main() {
	cmd.ConfigureCli()
}
