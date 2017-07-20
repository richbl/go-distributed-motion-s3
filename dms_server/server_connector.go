package server

import (
	"fmt"
	"go_server/dms_libs"
	"net"
	"strconv"
)

// StartServer starts the TCP server
func StartServer(ServerPort int) {

	if listener, error := net.Listen("tcp", ":"+fmt.Sprint(ServerPort)); error != nil {
		dmslibs.LogFatal(error.Error())
	} else {
		defer listener.Close()
		serverLoop(listener)
	}

}

// serverLoop starts a loop to listen for clients, spawning a separate processing thread on
// client connect
func serverLoop(listener net.Listener) {

	for {

		if conn, err := listener.Accept(); err != nil {
			dmslibs.LogFatal(err.Error())
		} else {
			dmslibs.LogInfo("open connection from:" + conn.RemoteAddr().String())
			go processClientRequest(conn)
		}

	}

}

// processClientRequest passes motion detector application state to client listeners based on logic
// found in entryPointRoutine()
func processClientRequest(conn net.Conn) {
	dmslibs.LogDebug(dmslibs.GetFunctionName())
	state := DetermineMotionDetectorState()

	if _, err := conn.Write([]byte(strconv.Itoa(int(state)))); err != nil {
		dmslibs.LogInfo(err.Error())
	}

	dmslibs.LogInfo("close connection from:" + conn.RemoteAddr().String())
	conn.Close()
}
