package main

import "onescan/scans"

func main() {
	/*s := new(scans.SynScan)
	ports := make([]string, 100)
	for i := 0; i < len(ports); i++ {
		ports[i] = strconv.Itoa(i)
	}
	s.Start(ports, "192.168.1.1", time.Second)*/
	s := new(scans.ArpScan)
	s.Start()
}
