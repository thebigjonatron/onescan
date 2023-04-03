package scans

import "net"

type Scanner interface {
	CloseConnection()
	StartScanner(ports []int, ip net.IP)
	StopScanner()
}
