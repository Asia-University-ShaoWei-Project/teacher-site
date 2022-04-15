package model

import "time"

type ServiceConfig struct {
	JwtSecure       []byte
	SaltSize        int
	TokenExpireTime time.Duration
	HashCost        int
}

// todo: use os.Getenv()
func NewMockServiceConfig() *ServiceConfig {
	return &ServiceConfig{
		JwtSecure:       []byte(`secure`),
		SaltSize:        16,
		TokenExpireTime: 1,
		HashCost:        10,
	}
}

type CacheConfig struct {
	MaxReTry int
}

func NewMockCacheConfig() *CacheConfig {
	return &CacheConfig{
		MaxReTry: 2,
	}
}
