package scans

import (
	"log"
	"net"
)

type ArpScan struct {
}

func Start() {

}

func arpPacket() {

}

func localMAC() *[]net.Interface {
	intface, err := net.Interfaces()
	if err != nil {
		log.Printf("Cannot get interfaces %s", err)
		return nil
	}
	return &intface
}
