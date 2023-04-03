package scans

import (
	"log"
	"net"
	"os"
	"syscall"
	"time"
)

type Synscan struct {
}

/*
*
Syn scan specific. Sends RST packet to close connection.
*/
func (scan *Synscan) closeConnection(connection *net.TCPConn) {
	err := connection.SetLinger(0)
	if err != nil {
		log.Printf("Can't set linger")
	}
	err = connection.Close()
	if err != nil {
		log.Printf("Can't close connection")
	}
}

func (scan *Synscan) Start(ports []string, ip string, timeout time.Duration) {
	ipNet := net.ParseIP(ip)
	for _, value := range ports {
		address := net.JoinHostPort(ipNet.String(), value)
		connection, err := net.DialTimeout("tcp", address, timeout)
		if err != nil {
			// Check for rst return
			if rstError, ok := err.(*net.OpError).Err.(*os.SyscallError); ok && rstError.Err == syscall.ECONNREFUSED {
				log.Printf("%s is closed", address)
				// Checks for timeout error
			} else if timeoutError, ok := err.(net.Error); ok && timeoutError.Timeout() {
				log.Printf("%s is filtered or unreachable", address)
			} else {
				log.Printf("%s", err)
			}
		} else {
			log.Printf("%s is open", address)
			scan.closeConnection(connection.(*net.TCPConn))
		}
	}
}
