package controller

import (
	"fmt"
	"onescan/scans"
	"strconv"
)

func Start(args []string) {
	fmt.Println(args)
	myUtils := scans.Utils{
		Ip:      "192.168.0.1",
		Mask:    "255.255.255.0",
		Intface: "eth0",
		Ssid:    "MyWifiNetwork",
	}
	findMode(args[0]).Start(myUtils)

	s := new(scans.SynScan)
	ports := make([]string, 100)
	for i := 0; i < len(ports); i++ {
		ports[i] = strconv.Itoa(i)
	}
	s.Start()
	//s := new(scans.ArpScan)
	//s.Start()
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

func addressResolve(arg string) {

}
