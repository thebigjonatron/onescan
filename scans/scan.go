package scans

type Scanner interface {
	Start(utils Utils)
}

type Utils struct {
	Ip      string
	Mask    string
	Intface string
	Ssid    string
	Ports   []string
}
