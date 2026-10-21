package config

import (
	"fmt"

	"github.com/redis/go-redis/v9"
)

type RedisConfig struct {
	Host     string
	Username string
	Password string
	Port     string
}

func NewRedis(host, username, password, port string) *RedisConfig {
	return &RedisConfig{
		Host:     host,
		Username: username,
		Password: password,
		Port:     port,
	}

}

func (rc *RedisConfig) Connect() *redis.Client {
	return redis.NewClient(&redis.Options{
		Username: rc.Username,
		Password: rc.Password,
		Addr:     fmt.Sprintf("%s:%s", rc.Host, rc.Port),
	})

}
