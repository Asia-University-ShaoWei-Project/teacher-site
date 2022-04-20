package config

import "time"

// todo: use viper
type Config struct {
	Server *Server
	Secure *Secure
	Jwt    *Jwt
	Redis  *Redis
	DB     *DB
}
type Server struct {
	Addr               string
	StaticRelativePath string
	StaticRootPath     string
	TemplatePath       string
}

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

func New() *Config {
	return &Config{
		Server: newServer(),
		Secure: newSecure(),
		Jwt:    newJwt(),
		Redis:  newRedis(),
		DB:     newDB(),
	}
}
func newServer() *Server {
	return &Server{
		Addr:               ":80",
		StaticRelativePath: "/static",
		StaticRootPath:     "./static",
		TemplatePath:       "templates/*",
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
		TokenExpireTime: 2,
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
