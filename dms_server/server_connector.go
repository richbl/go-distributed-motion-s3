package server

import (
	"fmt"
	"go_server/libs"
	"net"
	"os"
)

// Initialize comment
func Initialize(ServerPort int, entryPointRoutine func()) {
	startServer(ServerPort, entryPointRoutine)
}

// startServer comment
func startServer(ServerPort int, entryPointRoutine func()) {
	listener, error := net.Listen("tcp", ":"+fmt.Sprint(ServerPort))

	if error != nil {
		fmt.Println("Error listening:", error.Error())
		os.Exit(1)
	}

	defer listener.Close()
	serverLoop(listener, entryPointRoutine)
}

// serverLoop comment
func serverLoop(listener net.Listener, entryPointRoutine func()) {
	for {
		conn, err := listener.Accept()

		if err != nil {
			fmt.Println("Error accepting connection from: ", conn.RemoteAddr().String(), err.Error())
			os.Exit(1)
		}

		fmt.Println("--> OPEN connection from: ", conn.RemoteAddr().String())
		go processClientRequest(conn, entryPointRoutine)
	}
}

// processClientRequest comment
func processClientRequest(conn net.Conn, entryPointRoutine func()) {

	libconfig.PrintFunctionName()

	// TODO
	_, err := conn.Write([]byte("disable"))

	if err != nil {
		fmt.Println("ERROR processing client request:", err.Error())
	} else {
		entryPointRoutine()
	}

	fmt.Println("<-- CLOSE connection from: ", conn.RemoteAddr().String())
	conn.Close()
}
