package scans

import (
	"fmt"
	"log"
	"net"
)

type ArpScan struct {
}

func (scan *ArpScan) Start() {
	localMAC()
}

func arpPacket() {

}

func localMAC() *[]net.Interface {
	// Should find a way to know which non-loopback interface we want
	intface, err := net.Interfaces()
	if err != nil {
		log.Printf("Cannot get interfaces %s", err)
		return nil
	}
	fmt.Println(intface)
	return &intface
}
