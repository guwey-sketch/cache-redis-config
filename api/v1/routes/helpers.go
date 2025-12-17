package cache_redis_config

import (
	"fmt"
	"strings"
)

// RedisConfig holds the configuration for Redis
type RedisConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
	DB   int    `yaml:"db"`
}

// NewRedisConfig returns a new RedisConfig instance
func NewRedisConfig(host string, port int, db int) *RedisConfig {
	return &RedisConfig{
		Host: host,
		Port: port,
		DB:   db,
	}
}

// String returns a string representation of the RedisConfig instance
func (r *RedisConfig) String() string {
	return fmt.Sprintf("redis://%s:%d/%d", r.Host, r.Port, r.DB)
}

// IsValid checks if the RedisConfig instance is valid
func (r *RedisConfig) IsValid() error {
	if!strings.HasPrefix(r.Host, "localhost") &&!strings.HasPrefix(r.Host, "127.0.0.1") {
		return fmt.Errorf("invalid host: %s", r.Host)
	}
	if r.Port <= 0 {
		return fmt.Errorf("invalid port: %d", r.Port)
	}
	if r.DB <= 0 {
		return fmt.Errorf("invalid db: %d", r.DB)
	}
	return nil
}