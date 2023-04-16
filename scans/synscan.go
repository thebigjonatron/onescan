package scans

import (
	"log"
	"net"
	"os"
	"syscall"
	"time"
)

type SynScan struct {
}

/*
*
Syn scan specific. Sends RST packet to close connection.
*/
func (scan *SynScan) closeConnection(connection *net.TCPConn) {
	err := connection.SetLinger(0)
	if err != nil {
		log.Printf("Can't set linger")
	}
	err = connection.Close()
	if err != nil {
		log.Printf("Can't close connection")
	}
}

func (scan *SynScan) scan(ip *net.IP, port string) {
	address := net.JoinHostPort(ip.String(), port)
	connection, err := net.DialTimeout("tcp", address, 2*time.Second)
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

func (scan *SynScan) Start(utils UtilsArp) {
	ipNet := net.ParseIP(utils.Ip)
	for _, port := range utils.Ports {
		go scan.scan(&ipNet, port)
	}
}
