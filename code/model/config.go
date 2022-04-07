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
func NewTMPServiceConfig() *ServiceConfig {
	return &ServiceConfig{
		JWTSecure:       []byte(`secure`),
		SaltSize:        16,
		TokenExpireTime: 1,
		HashCost:        10,
	}
}
func NewTMPCacheConfig() *CacheConfig {
	return &CacheConfig{
		MaxReTry: 2,
	}
}
