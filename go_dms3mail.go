// Package main go_dms3mail initializes a dms3mail device component
//
package main

import (
	"go-distributed-motion-s3/dms3mail"
)

func main() {
	dms3mail.Init("/etc/distributed-motion-s3")
}
