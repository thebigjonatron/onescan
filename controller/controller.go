package controller

import (
	"fmt"
	"onescan/scans"
	"strconv"
)

func Start(args []string) {
	fmt.Println(args)
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
	findMode(args[0]).Start(myUtils)
}

// Default mode auto
func findMode(arg string) scans.Scanner {
	switch arg {
	case "syn":
		return new(scans.SynScan)
	case "arp":
		return new(scans.ArpScan)
	default:
		return nil
	}
}

// Define if it's a range or not
func addressResolver(arg string) {

}

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
