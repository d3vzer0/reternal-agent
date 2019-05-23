// +build linux darwin

package modules

import (
	"bufio"
	"fmt"
	"os"
	"syscall"
)

func CreateNamedPipe(input string) string {
	err := syscall.Mkfifo(input, 0666)
	if err != nil {
		result := "Make named pipe file error:"
		return result
	}
	fmt.Println("open a named pipe file for read.")
	file, err := os.OpenFile(input, os.O_CREATE, os.ModeNamedPipe)
	if err != nil {
		result := "Open named pipe file error:"
		return result
	}
	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadBytes('\n')
		if err == nil {
			fmt.Print("load string:" + string(line))
		}
	}
}
