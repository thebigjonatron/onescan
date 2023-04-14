package scans

type Scanner interface {
	Start(utils Utils)
}

type DefaultScan struct {
}

func (*DefaultScan) Start(utils Utils) {
	//Run default scan. Arp scan and syn scan
}

type Utils struct {
	Ip      string
	Mask    string
	Intface string
	Ssid    string
	Ports   []string
}
