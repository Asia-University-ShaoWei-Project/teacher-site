package model

import "time"

type Config struct {
	JWTSecure       []byte
	SaltSize        int
	TokenExpireTime time.Duration
	HashCost        int
}

// todo: use os.Getenv()
func NewTMPConfig() *Config {
	return &Config{
		JWTSecure:       []byte(`secure`),
		SaltSize:        16,
		TokenExpireTime: 1,
		HashCost:        10,
	}
}
