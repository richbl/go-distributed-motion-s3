// Package main go_dms3server initializes a dms3server device component
//
package main

import (
	"go-distributed-motion-s3/dms3server"
)

func main() {
	dms3server.Init("/etc/distributed-motion-s3")
}
