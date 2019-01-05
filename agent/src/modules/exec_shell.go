package modules

import (
	"bytes"
	"os/exec"
	"strings"
)

func ExecShell(input string) string {

	args := strings.Fields(input)
	result := "No input specified"

	if len(args) > 0 {
		command, parameters := args[0], args[1:]
		cmd := exec.Command(command, parameters...)
		var stdout, stderr bytes.Buffer
		cmd.Stdout = &stdout
		cmd.Stderr = &stderr
		err := cmd.Run()
		if err != nil {
			result = string(stderr.Bytes())
		} else {
			result = string(stdout.Bytes())
		}
	}
	return result
}
