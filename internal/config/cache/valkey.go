package cache

import "os"

type ValkeyConf struct {
	Host string
	Port string
}

func NewValkeyConf() *ValkeyConf {
	return &ValkeyConf{
		Host: os.Getenv("VALKEY_HOST"),
		Port: os.Getenv("VALKEY_PORT"),
	}
}
