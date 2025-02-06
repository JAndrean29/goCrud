package config

import (
	"time"
)

type DatabaseConfig struct {
	Host           string // host (e.g. localhost) or absolute path to unix domain socket directory (e.g. /private/tmp)
	Port           uint16
	Database       string
	User           string
	Password       string
	ConnectTimeout time.Duration
}
