package scans

import (
	"time"
)

type Scanner interface {
	Start(ports []string, ip string, timeout time.Duration)
}
