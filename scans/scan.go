package scans

import (
	"net"
)

type Scanner interface {
	CloseConnection(conn net.Conn)
}

func StartScanner(scanner Scanner, ports []int, ip net.IP) {

}
