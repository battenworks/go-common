package command

import (
	"github.com/battenworks/go-common/console"
	"os/exec"
	"strings"
)

type CommandExecutor struct{}

// Execute executes the given command with the given arguments
func (ce CommandExecutor) Execute(cmdName string, cmdArgs []string) ([]byte, error) {
	cmd := exec.Command(cmdName, cmdArgs...)
	out, err := cmd.CombinedOutput()

	console.Yellow(strings.Join(cmd.Args, " "))

	if err != nil {
		console.Red(" error\n")
		return out, err
	}

	console.Green(" done\n")

	return out, nil
}
