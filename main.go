package main

import (
	"fmt"
	"onescan/scans"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	fmt.Println(args)
	fmt.Println(len(args))
	s := new(scans.SynScan)
	ports := make([]string, 100)
	for i := 0; i < len(ports); i++ {
		ports[i] = strconv.Itoa(i)
	}
	s.Start(ports, "192.168.1.1")
	//s := new(scans.ArpScan)
	//s.Start()
}
