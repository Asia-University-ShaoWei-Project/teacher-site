package model

import "time"

type ServiceConfig struct {
	JWTSecure       []byte
	SaltSize        int
	TokenExpireTime time.Duration
	HashCost        int
}
type CacheConfig struct {
	MaxReTry int
}

// todo: use os.Getenv()
func NewMockServiceConfig() *ServiceConfig {
	return &ServiceConfig{
		JWTSecure:       []byte(`secure`),
		SaltSize:        16,
		TokenExpireTime: 1,
		HashCost:        10,
	}
}
func NewMockCacheConfig() *CacheConfig {
	return &CacheConfig{
		MaxReTry: 2,
	}
}
