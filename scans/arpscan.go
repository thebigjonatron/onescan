package scans

import (
	"bytes"
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"log"
	"net"
	"strconv"
	"strings"
	"time"
)

type ArpScan struct {
}

func (arpScan *ArpScan) Start() {
	scan(getLocalInterface("enp58s0u1u4"))
}

func scan(intface *net.Interface) {
	handle, err := pcap.OpenLive(intface.Name, 65536, true, pcap.BlockForever)
	if err != nil {
		panic(err)
	}
	defer handle.Close()
	stop := make(chan struct{})
	go readARPFromHandle(handle, intface, stop)
	defer close(stop)
	writeARPToHandle(handle, intface, stop)
	time.Sleep(10 * time.Second) //Waits 10 secs and stop main exec thread
}

func readARPFromHandle(handle *pcap.Handle, intface *net.Interface, stop chan struct{}) {
	src := gopacket.NewPacketSource(handle, layers.LayerTypeEthernet)
	in := src.Packets()
	for {
		var packet gopacket.Packet
		select {
		case <-stop:
			return
		case packet = <-in:
			arpLayer := packet.Layer(layers.LayerTypeARP)
			if arpLayer == nil {
				continue
			}
			arp := arpLayer.(*layers.ARP)
			if arp.Operation != layers.ARPReply || bytes.Equal([]byte(intface.HardwareAddr), arp.SourceHwAddress) {
				// This is a packet I sent.
				continue
			}
			fmt.Printf("IP %v is active \n", net.IP(arp.SourceProtAddress))
		}
	}
}

func writeARPToHandle(handle *pcap.Handle, intface *net.Interface, stop chan struct{}) {
	intfaceAddr := getNetwork(intface)
	buffer := gopacket.NewSerializeBuffer()
	fmt.Println([]byte(net.ParseIP(intfaceAddr[0]).To4()))
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
		SourceProtAddress: []byte(net.ParseIP(intfaceAddr[0]).To4()),
		DstHwAddress:      []byte{0, 0, 0, 0, 0, 0},
	}
	for _, ip := range getSubnetRange(intfaceAddr) {
		arp.DstProtAddress = ip.To4()
		err := gopacket.SerializeLayers(buffer, options, &eth, &arp)
		if err != nil {
			fmt.Println("c", err)
			return
		}
		if err := handle.WritePacketData(buffer.Bytes()); err != nil {
			fmt.Println("Can't write to handle", err)
		}
	}
}

func getSubnetRange(network []string) []net.IP {
	var ips []net.IP
	i, err := strconv.Atoi(network[1])
	if err != nil {
		fmt.Println("Can't convert mask", err)
	}
	mask := net.CIDRMask(i, 32)
	networkAddr := net.ParseIP(network[0]).Mask(mask)
	broadcast := make(net.IP, len(networkAddr))
	for i := 0; i < len(networkAddr); i++ {
		broadcast[i] = networkAddr[i] | ^mask[i]
	}
	for ip := inc(networkAddr); !ip.Equal(broadcast); ip = inc(ip) {
		ips = append(ips, ip)
	}
	return ips
}

func inc(ip net.IP) net.IP {
	incIP := make(net.IP, len(ip))
	copy(incIP, ip)
	for i := len(incIP) - 1; i >= 0; i-- {
		(incIP)[i]++
		if (incIP)[i] > 0 {
			return incIP
		}
	}
	return nil
}

func getNetwork(intface *net.Interface) []string {
	addrs, err := intface.Addrs()
	if err != nil {
		fmt.Println("No address for interface")
	}
	return strings.Split(addrs[0].String(), "/")
}

func getLocalInterface(name string) *net.Interface {
	intfaces, err := net.Interfaces()
	if err != nil {
		log.Printf("Cannot get interfaces %s", err)
		return nil
	}
	for _, intface := range intfaces {
		if intface.Name == name {
			return &intface // Or find default with the ip range of the network
		}
	}
	return nil
}
