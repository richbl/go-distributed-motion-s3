package client

import (
	"fmt"
	"go_server/dms_libs"
	"net"
	"time"
)

// Initialize comment
func Initialize(ServerIP string, ServerPort int, entryPointRoutine func(string)) {
	dmslibs.PrintFuncName()
	startClient(ServerIP, ServerPort, entryPointRoutine)
}

// startClient comment
func startClient(ServerIP string, ServerPort int, entryPointRoutine func(string)) {

	for {
		dmslibs.PrintFuncName()
		conn, err := net.Dial("tcp", ServerIP+":"+fmt.Sprint(ServerPort))

		if err != nil {
			// server not found, sleep and try again
			dmslibs.Info.Println(err.Error())
		} else {
			// server connection established
			defer conn.Close()
			go processClientRequest(conn, entryPointRoutine)
		}

		time.Sleep(CheckInterval * time.Second)
	}

}

// processClientRequest comment
func processClientRequest(conn net.Conn, entryPointRoutine func(string)) {

	dmslibs.PrintFuncName()

	buf := make([]byte, 256)
	n, err := conn.Read(buf)

	if err != nil {
		dmslibs.Info.Println(err.Error())
	} else {
		entryPointRoutine(string(buf[:n]))
	}

	conn.Close()
}
