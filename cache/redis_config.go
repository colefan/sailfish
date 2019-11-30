package cache

import (
	"errors"
	"github.com/colefan/config"
)

type RedisConfig struct {
	conf        *config.IniConfig
	MaxIdle     int
	MaxActive   int
	IdleTimeout int
	Address     string
	DbNum       int
	Password    string
}

func NewRedisConfig(configFile string) (*RedisConfig, error) {
	c := &RedisConfig{}
	c.conf = config.NewIniConfig()
	if err := c.conf.Parse(configFile); err != nil {
		return nil, err
	}
	var err error
	if c.MaxIdle, err = c.conf.Int("MaxIdle"); err != nil {
		return nil, err
	}
	if c.MaxActive, err = c.conf.Int("MaxActive"); err != nil {
		return nil, err
	}
	if c.IdleTimeout, err = c.conf.Int("IdleTimeout"); err != nil {
		return nil, err
	}
	c.Address = c.conf.String("Address")
	if len(c.Address) == 0 {
		return nil, errors.New("redis address is empty")
	}
	if c.DbNum, err = c.conf.Int("DbNum"); err != nil {
		return nil, err
	}
	c.Password = c.conf.String("Password")
	return c, nil
}
