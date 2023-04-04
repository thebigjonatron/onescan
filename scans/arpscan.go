package scans

import (
	"log"
	"net"
)

type ArpScan struct {
}

func (scan *ArpScan) Start() {
	getLocalInterfaceMAC("192.168.1.1")
}

func arpPacket() {

}

func getLocalInterfaceMAC(ip string) *net.Interface {
	// Should find a way to know which non-loopback interface we want
	intfaces, err := net.Interfaces()
	if err != nil {
		log.Printf("Cannot get interfaces %s", err)
		return nil
	}
	for _, intface := range intfaces {
		if intface.Flags&net.FlagUp != 0 &&
			intface.Flags&net.FlagBroadcast != 0 &&
			intface.Flags&net.FlagMulticast != 0 &&
			intface.Flags&net.FlagLoopback == 0 &&
			intface.HardwareAddr != nil {
			return &intface //Find first up interface, maybe not the one we want ?? Should be able to select interface
			// Or find default with the ip range of the network
		}
	}
	return nil
}
