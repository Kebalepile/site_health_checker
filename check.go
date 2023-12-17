package main

import (
	"fmt"
	"net"
	"time"
)

func Check(destination string, port string) string {
	addr := fmt.Sprintf("%v:%v", destination, port)
	timeout := time.Duration(20 * time.Second)
	conn, err := net.DialTimeout("tcp", addr, timeout)

	var status string

	if err != nil {
		status = fmt.Sprintf("[DOWN] %v is offline,\n Error: %v", destination, err)
	} else {
		status = fmt.Sprintf("[UP] %v is online,\n From: %v To: %v", destination, conn.LocalAddr(), conn.RemoteAddr())
	}
	return status
}
