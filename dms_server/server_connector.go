package server

import (
	"fmt"
	"go_server/dms_libs"
	"net"
)

// Initialize comment
func Initialize(ServerPort int, entryPointRoutine func()) {
	startServer(ServerPort, entryPointRoutine)
}

// startServer comment
func startServer(ServerPort int, entryPointRoutine func()) {
	listener, error := net.Listen("tcp", ":"+fmt.Sprint(ServerPort))

	if error != nil {
		dmslibs.Error.Fatalln(error.Error())
	}

	defer listener.Close()
	serverLoop(listener, entryPointRoutine)
}

// serverLoop comment
func serverLoop(listener net.Listener, entryPointRoutine func()) {
	for {
		conn, err := listener.Accept()

		if err != nil {
			dmslibs.Error.Fatalln(err.Error())
		}

		fmt.Println("--> OPEN connection from: ", conn.RemoteAddr().String())
		go processClientRequest(conn, entryPointRoutine)
	}
}

// processClientRequest comment
func processClientRequest(conn net.Conn, entryPointRoutine func()) {

	dmslibs.PrintFuncName()

	// TODO

	_, err := conn.Write([]byte("disable"))

	if err != nil {
		dmslibs.Info.Println(err.Error())
	} else {
		entryPointRoutine()
	}

	fmt.Println("<-- CLOSE connection from: ", conn.RemoteAddr().String())
	conn.Close()
}
