package scans

import (
	"log"
	"net"
	"time"
)

type Synscan struct {
}

/*
*
Syn scan specific. Sends RST packet to close connection.
*/
func (scan *Synscan) closeConnection(conn *net.TCPConn) {
	err := conn.SetLinger(0)
	if err != nil {
		log.Printf("Can't set linger")
	}
	err = conn.Close()
	if err != nil {
		log.Printf("Can't close connection")
	}
}

func (scan *Synscan) StartScan(ports []string, ip string, timeout time.Duration) {
	ipNet := net.ParseIP(ip)
	for _, value := range ports {
		address := net.JoinHostPort(string(ipNet), value)
		conn, err := net.DialTimeout("tcp", address, timeout)
		if err != nil {
			log.Printf("%s is closed or filtered", address)
		} else {
			/**if _, err := conn.Write([]byte{}); err != nil {
				fmt.Printf("Error sending packet", err)
			}
			fmt.Printf("Port %d is open\n", port)**/
			log.Printf("Connection is open")
			scan.closeConnection(conn.(*net.TCPConn))
		}
	}
}
