package config

import (
	"errors"
	"flag"
	"log"
	"net/url"
	"os"
	"strconv"
)

func getEnv(key string, defaultVal string) string {
	if envVal, ok := os.LookupEnv(key); ok {
		return envVal
	}
	return defaultVal
}

func getEnvInt64(key string, defaultVal int64) int64 {
	if envVal, ok := os.LookupEnv(key); ok {
		envInt64, err := strconv.ParseInt(envVal, 10, 64)
		if err == nil {
			return envInt64
		}
	}
	return defaultVal
}

type Config struct {
	Host    string
	Port    int64
	Timeout int64
}

func New() (*Config, error) {
	var (
		host    = flag.String("host", getEnv("HOST", "https://127.0.0.1"), "URL of the service")
		port    = flag.Int64("port", getEnvInt64("PORT", 8080), "Port")
		timeout = flag.Int64("timeout", getEnvInt64("TIMEOUT", 120), "Connection Timeout")
	)
	flag.Parse()

	conf := &Config{
		Host:    *host,
		Port:    *port,
		Timeout: *timeout,
	}

	isValid := conf.Validate()
	if isValid {
		return conf, nil
	}
	return conf, errors.New("Malformed Config")
}

func (c Config) Validate() bool {
	if c.Timeout == 0 {
		log.Fatal("Timeout should be greater than zero")
		return false
	}
	if c.Port != 80 && c.Port != 8080 {
		log.Fatal("Acceptable ports are 80 and 8080. Provided:", c.Port)
		return false
	}
	_, err := url.ParseRequestURI(c.Host)
	if err != nil {
		log.Fatal("Invalid host:", c.Host)
		return false
	}
	return true
}
