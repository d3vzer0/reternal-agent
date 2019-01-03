package modules

import (
	"fmt"
	"os/exec"
	"strings"
)

func ExecShell(input string) string {
	args := strings.Fields(input)
	arguments := strings.Join(args[1:], " ")
	output, err := exec.Command(args[0], arguments).Output()
	if err != nil {
		error := fmt.Sprint(err)
		return string(error)
	} else {
		return string(output)
	}
}
