package cle

import (
	"os"
	"os/exec"
	"syscall"

	log "github.com/sirupsen/logrus"
)

//ExecCMD Executes the underlying `cmd` with `args`
func ExecCMD(
	cmd string,
	args []string,
) error {
	initialCharacters := make(map[string]bool)
	initialCharacters["/"] = true
	initialCharacters["."] = true

	if _, found := initialCharacters[cmd[0:1]]; !found {
		foundCmd, err := exec.LookPath(cmd)
		if err != nil {
			log.Fatalf("Could not find executable '%s' in $PATH", cmd)
		}

		log.Printf("Found '%s' in '%s'", cmd, foundCmd)
		cmd = foundCmd
	}

	execErr := syscall.Exec(
		cmd,
		args,
		os.Environ(),
	)

	return execErr
}
