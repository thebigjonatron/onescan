package main

import (
	"fmt"
	"log"
	"net"
	"onescan/scans"
	"time"
)

func main() {
	s := new(scans.Synscan)

	for i := 1; i <= 100; i++ {
		ip := net.ParseIP("192.168.1.1")
		port := i
		conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", ip, port), 2*time.Second)
		if err != nil {
			log.Printf("Port is closed or filtered", port)
		} else {
			if _, err := conn.Write([]byte{}); err != nil {
				fmt.Printf("Error sending packet", err)
			}
			s.CloseConnection()
			fmt.Printf("Port %d is open\n", port)
		}
	}

}
