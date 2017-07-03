package client

import (
	"fmt"
	"go_server/libs"
	"net"
	"os"
	"time"
)

// Initialize comment
func Initialize(ServerIP string, ServerPort int, entryPointRoutine func(string)) {
	libconfig.PrintFunctionName()
	startClient(ServerIP, ServerPort, entryPointRoutine)
}

// startClient comment
func startClient(ServerIP string, ServerPort int, entryPointRoutine func(string)) {

	for {
		libconfig.PrintFunctionName()
		conn, error := net.Dial("tcp", ServerIP+":"+fmt.Sprint(ServerPort))

		if error != nil {
			fmt.Println("Error listening:", error.Error())
			os.Exit(1)
		}

		defer conn.Close()
		go processClientRequest(conn, entryPointRoutine)
		time.Sleep(CheckInterval * time.Second)
	}

}

// processClientRequest comment
func processClientRequest(conn net.Conn, entryPointRoutine func(string)) {

	libconfig.PrintFunctionName()

	buf := make([]byte, 256)
	n, err := conn.Read(buf)

	if err != nil {
		fmt.Println("ERROR processing client request:", err.Error())
	} else {
		entryPointRoutine(string(buf[:n]))
	}

	conn.Close()
}
