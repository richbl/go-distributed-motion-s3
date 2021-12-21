// Package main dms3mail initializes a dms3mail device component
//
package main

import "github.com/richbl/go-distributed-motion-s3/dms3mail"

func main() {
	dms3mail.Init("/etc/distributed-motion-s3")
}
