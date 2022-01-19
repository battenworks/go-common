package command

import (
	"os/exec"
	"strings"

	"github.com/battenworks/go-common/console"
)

type CommandExecutor struct{}

// Execute executes the given command with the given arguments
func (ce CommandExecutor) Execute(cmdName string, cmdArgs ...string) ([]byte, error) {
	cmd := exec.Command(cmdName, cmdArgs...)
	out, err := cmd.CombinedOutput()

	console.Out("command: ")
	console.Yellow(strings.Join(cmd.Args, " ") + " ")

	if err != nil {
		console.Redln("error")
		return out, err
	}

	console.Greenln("success")

	return out, nil
}
