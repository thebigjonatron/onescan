# Onescan
onescan is a network scanning tool written in Go that can perform two types of scans: ARP scanning and SYN scanning. By default, onescan will run both, starting with the ARP scan to discover the network and then performing the SYN scan to look for open ports on all discovered machines.

## Install
```
git clone https://github.com/vanderpluijmg/onescan
cd onescan
go build ./main.go
go run ./main.go
```

## Usage
To run onescan, simply execute the following command:

```
go run onescan.go <interface>
```
Where <interface> is the interface connected to the network you want to discover.

By default, onescan will run both ARP and SYN scans. To run only one of the scans, use the following commands:

```
go run onescan.go --arp <interface>    # run only ARP scan
go run onescan.go --syn <IP range>    # run only SYN scan
```
ARP scanning is used to discover hosts on the network. This scan sends ARP requests to all IP addresses on the network and listens for responses. The responses are used to identify hosts that are online and active.

SYN scanning is used to scan for open ports on a target host. This scan sends SYN packets to a range of ports on the target host and listens for responses. If a response is received, the port is marked as open.



