package cmd

import (
	"os"
	"syscall"

	log "github.com/sirupsen/logrus"
)

func execCMD() {
	binary := os.Args[1]
	args := os.Args[2:]
	env := os.Environ()

	execErr := syscall.Exec(binary, args, env)

	//cleanup sidecars

	if execErr != nil {
		log.Fatal(execErr)
	}
}
