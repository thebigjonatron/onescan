package main

import (
	"onescan/scans"
	"time"
)

func main() {
	s := new(scans.Synscan)
	mySlice := make([]string, 3)
	s.StartScan(mySlice, "192.168.1.1.", time.Duration(1)*2)
}
