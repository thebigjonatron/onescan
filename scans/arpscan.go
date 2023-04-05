package scans

import (
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"log"
	"net"
)

// TODO
// Better error management
// Better way to select interface (with ip, by name ?
// Multi-thread bois :}

type ArpScan struct {
}

func (arpScan *ArpScan) Start() {
	scan(getLocalInterface("enp58s0u1u4"))
}

// Must scan given subnet
func scan(intface *net.Interface) {
	// Open pcap handle to read and write arp requests
	handle, err := pcap.OpenLive(intface.Name, 65536, true, pcap.BlockForever)
	if err != nil {
		panic(err)
	}
	defer handle.Close()
	go writeARPToHandle(handle, intface)
	//go readARPFromHandle(handle)
}

func readARPFromHandle(handle *pcap.Handle) {

}

func writeARPToHandle(handle *pcap.Handle, intface *net.Interface) {
	intfaceAddr := getIP(intface)
	buffer := gopacket.NewSerializeBuffer()
	options := gopacket.SerializeOptions{
		FixLengths:       true,
		ComputeChecksums: true,
	}
	eth := layers.Ethernet{
		SrcMAC:       intface.HardwareAddr,
		DstMAC:       net.HardwareAddr{0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
		EthernetType: layers.EthernetTypeARP,
	}
	arp := layers.ARP{
		AddrType:          layers.LinkTypeEthernet,
		Protocol:          layers.EthernetTypeIPv4,
		HwAddressSize:     6,
		ProtAddressSize:   4,
		Operation:         layers.ARPRequest,
		SourceHwAddress:   []byte(intface.HardwareAddr),
		SourceProtAddress: []byte(intfaceAddr),
		DstHwAddress:      []byte{0, 0, 0, 0, 0, 0},
	}

	for _, ip := range getIPRange(intfaceAddr) {
		arp.DstProtAddress = []byte(ip)
		err := gopacket.SerializeLayers(buffer, options, &eth, &arp)
		if err != nil {
			return
		}
		if err := handle.WritePacketData(buffer.Bytes()); err != nil {
			fmt.Println("Can't write to handle")
		}
	}
}

func getIPRange(ip net.IP) []net.IP {

}

// Redo to get correct ip and assure it's valid.
func getIP(intface *net.Interface) net.IP {
	addrs, err := intface.Addrs()
	if err != nil {
		fmt.Println("no add")
	}
	for _, a := range addrs {
		println(a.Network(), a.String())
		return net.ParseIP(a.String())
	}
	return nil
}

func getLocalInterface(name string) *net.Interface {
	// Should find a way to know which non-loopback interface we want with ip ?
	intfaces, err := net.Interfaces()
	if err != nil {
		log.Printf("Cannot get interfaces %s", err)
		return nil
	}

	for _, intface := range intfaces {
		if intface.Name == name {
			fmt.Println(intface.Name)
			return &intface //Find first up interface, maybe not the one we want ?? Should be able to select interface
			// Or find default with the ip range of the network
		}
	}
	return nil
}
