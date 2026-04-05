package config

import "time"

type Config struct {
	Env      string        `yaml:"env"`
	TokenTTL time.Duration `yaml:"token_ttl"`
	Database Database      `yaml:"database"`
	GRPC     GRPC          `yaml:"grpc"`
}

type Database struct {
	Hostname string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

type GRPC struct {
	Port    int           `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}
