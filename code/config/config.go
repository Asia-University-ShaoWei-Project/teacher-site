package config

import "time"

// todo: use viper

type Secure struct {
	SaltSize int
	HashCost int
}
type Jwt struct {
	Secure          []byte
	TokenExpireTime time.Duration
}
type Redis struct {
	MaxReTry int
	Addr     string
	Password string
	Database int
}
type DB struct {
}
type Config struct {
	Secure *Secure
	Jwt    *Jwt
	Redis  *Redis
	DB     *DB
}

func New() *Config {
	return &Config{
		Secure: newSecure(),
		Jwt:    newJwt(),
		Redis:  newRedis(),
		DB:     newDB(),
	}
}
func newSecure() *Secure {
	return &Secure{
		SaltSize: 16,
		HashCost: 10,
	}
}
func newJwt() *Jwt {
	return &Jwt{
		Secure:          []byte(`secure`),
		TokenExpireTime: 1,
	}
}
func newRedis() *Redis {
	return &Redis{
		MaxReTry: 2,
		Addr:     ":6379",
		Password: "",
		Database: 0,
	}
}
func newDB() *DB {
	return &DB{}
}
