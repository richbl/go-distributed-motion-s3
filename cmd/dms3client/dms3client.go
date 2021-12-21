// Package main dms3client initializes a dms3client device component
//
package main

import "github.com/richbl/go-distributed-motion-s3/dms3client"

func main() {
	dms3client.Init("/etc/distributed-motion-s3")
}
