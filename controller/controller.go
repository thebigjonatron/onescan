package controller

import (
	"flag"
	"fmt"
	"onescan/scans"
	"strconv"
)

//Possible ways to start the program :
//	- onescan syn 10.10.10.10/32 12,14,15
//	- onescan 10.10.10.10
//	- onescan 10.10.10.10
//	- onescan 10.10.10.10

func Start(args []string) {
	fmt.Println(args)
	findMode(args[0]).Start(createUtils())
	ipAddress := flag.String("ip", "", "the IP address to scan (e.g. 10.10.10.10/32)")
	ports := flag.String("ports", "", "the comma-separated list of ports to scan (e.g. 12,14,15)")
	flag.Parse()
	fmt.Printf("IP address: %s\n", *ipAddress)
	fmt.Printf("Ports: %s\n", *ports)
	//portsSlice := parsePorts(*ports)ports)

}

func parseFlags() {

}

// Default mode auto
func findMode(arg string) scans.Scanner {
	switch arg {
	case "syn":
		return new(scans.SynScan)
	case "arp":
		return new(scans.ArpScan)
	default:
		return new(scans.DefaultScan) //Default scan
	}
}

// Define if it's a range or not
func addressResolver(arg string) {

}

func createUtils() scans.Utils {
	ports := make([]string, 100)
	for i := 0; i < len(ports); i++ {
		ports[i] = strconv.Itoa(i)
	}
	myUtils := scans.Utils{
		Ip:      "192.168.0.1",
		Mask:    "255.255.255.0",
		Intface: "eth0",
		Ssid:    "MyWifiNetwork",
		Ports:   ports,
	}
	return myUtils
}

/*
	func parsePorts(ports string) []int {
		portsSlice := []int{}
		for _, port := range split(ports, ',') {
			portsSlice = append(portsSlice, parseInt(port))
		}
		return portsSlice
	}
*/
func defaults() scans.Utils {
	ports := make([]string, 100)
	myUtils := scans.Utils{
		Ip:      "192.168.0.1",
		Mask:    "255.255.255.0",
		Intface: "eth0",
		Ssid:    "MyWifiNetwork",
		Ports:   ports,
	}
	return myUtils
}
