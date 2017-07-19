package server

import (
	"fmt"
	"go_server/dms_libs"
	"net"
	"strconv"
)

// StartServer starts the TCP server
func StartServer(ServerPort int) {
	listener, error := net.Listen("tcp", ":"+fmt.Sprint(ServerPort))

	if error != nil {
		dmslibs.LogFatal(error.Error())
	}

	defer listener.Close()
	serverLoop(listener)
}

// serverLoop starts a loop to listen for clients, spawning a separate processing thread on client connect
func serverLoop(listener net.Listener) {

	for {
		conn, err := listener.Accept()

		if err != nil {
			dmslibs.LogFatal(err.Error())
		}

		dmslibs.LogInfo("open connection from:" + conn.RemoteAddr().String())
		go processClientRequest(conn)
	}

}

// processClientRequest passes motion detector application state to client listeners based on logic found in entryPointRoutine()
func processClientRequest(conn net.Conn) {
	dmslibs.LogDebug(dmslibs.GetFunctionName())

	state := DetermineMotionDetectorState()
	_, err := conn.Write([]byte(strconv.Itoa(int(state))))

	if err != nil {
		dmslibs.LogInfo(err.Error())
	}

	dmslibs.LogInfo("close connection from:" + conn.RemoteAddr().String())
	conn.Close()
}
