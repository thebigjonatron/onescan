package scans

import (
	"log"
	"net"
)

type Synscan struct {
	conn *net.TCPConn
}

func (scan *Synscan) CloseConnection() {
	err := (*scan.conn).SetLinger(0)
	if err != nil {
		log.Printf("Can't set linger")
	}
	err = (*scan.conn).Close()
	if err != nil {
		log.Printf("Can't close connection")
	}
}
