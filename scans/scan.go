package scans

type Scanner interface {
	Start(ports []string, ip string)
}
