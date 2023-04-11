package controller

import (
	"fmt"
	"onescan/scans"
	"strconv"
)

func Start(args []string) {
	fmt.Println(args)
	s := new(scans.SynScan)
	ports := make([]string, 100)
	for i := 0; i < len(ports); i++ {
		ports[i] = strconv.Itoa(i)
	}
	s.Start(ports, "192.168.1.1")
	//s := new(scans.ArpScan)
	//s.Start()
}

// Default mode auto
func mode(arg string) scans.Scanner {
	switch arg {
	case "syn":
		return new(scans.SynScan)
	case "arp":
		return new(scans.SynScan)
	default:
		return new(scans.SynScan)
	}
}

func addressResolve(arg string) {

}
