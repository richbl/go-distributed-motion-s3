package main

import (
	"go_server/dms_client"
)

func main() {
	client.Initialize(client.ServerIP, client.ServerPort, client.ProcessMotionState)
}
