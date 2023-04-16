package controller

import (
	"flag"
	"fmt"
	"log"
	"net"
	"onescan/scans"
	"strconv"
	"strings"
)

func Start(args []string) {
	portsFlag := flag.String("p", "", "Comma seperated list of ports or range or ports. Separate with , and indicate range with -")
	intFlag := flag.String("h", "localhost", "host name or IP address")
	flag.Parse()
	ports, err := parsePorts(*portsFlag)

	if err != nil {
	}

	findMode(args[0]).Start(createUtils(ports, *intFlag))
}

func parsePorts(portsStr string) ([]int, error) {
	var ports []int
	for _, portStr := range strings.Split(portsStr, ",") {
		portRange := strings.Split(portStr, "-")
		if len(portRange) == 2 {
			p1 := validatePort(portRange[0])
			p2 := validatePort(portRange[1])
			min, max := minmax(p1, p2)
			if p1 != 0 || p2 != 0 {
				for i := min; i <= max; i++ {
					fmt.Println(i)
					ports = append(ports, i)
				}
			}
		} else {
			ports = append(ports, validatePort(portStr))
		}
	}
	return ports, nil
}

func minmax(a int, b int) (int, int) {
	var max, min int
	if a > b {
		max = a
		min = b
	} else {
		max = b
		min = a
	}
	return min, max
}

func validatePort(portStr string) int {
	port, err := strconv.Atoi(portStr)
	if err != nil || port < 1 || port > 65535 {
		log.Println("non valid port")
		return 0
	}
	return port
}

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

func createUtils(ports []int, intfaceStr string) scans.UtilsArp {
	intface, err := net.InterfaceByName(intfaceStr)
	if err != nil {

	}
	myUtils := scans.UtilsArp{
		Intface: intface,
		Ports:   ports,
	}
	return myUtils
}
