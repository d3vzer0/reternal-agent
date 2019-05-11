// +build linux darwin

package modules

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func SocketForward(input string) string {
	ln, err := net.Listen("tcp", ":445")
	if err != nil {
		result := "Unable to open socket on port 445"
		return result
	}
	conn, err := ln.Accept()

	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message Received:", string(message))
		newmessage := strings.ToUpper(message)
		conn.Write([]byte(newmessage + "\n"))
	}
}
