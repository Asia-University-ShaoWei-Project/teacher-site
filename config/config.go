package config

import (
	"time"
)

// todo: use viper
type Config struct {
	Server *Server
	Limit  *Limit
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
type Limit struct {
	TeacherListPageCount int
}
type Secure struct {
	Salt          []byte
	SaltSize      int ``
	HashCost      int
	SessionSecret []byte
}
type Jwt struct {
	Secret          []byte
	TokenExpireTime time.Duration
}

type Redis struct {
	MaxReTry int
	Addr     string
	Password string
	Database int
}
type DB struct {
	Filename string
}

func New() *Config {
	return &Config{
		Server: newServer(),
		Limit:  newLimit(),
		Secure: newSecure(),
		Jwt:    newJwt(),
		Redis:  newRedis(),
		DB:     newDB(),
	}
}
func newServer() *Server {
	return &Server{
		Addr: ":80",
		// Addr:               ":" + os.Getenv("PORT"),
		StaticRelativePath: "/static",
		StaticRootPath:     "./static",
		TemplatePath:       "templates/*",
	}
}
func newLimit() *Limit {
	return &Limit{
		TeacherListPageCount: 10,
	}
}
func newSecure() *Secure {
	return &Secure{
		SaltSize:      16,
		HashCost:      10,
		SessionSecret: []byte("secret"),
	}
}
func newJwt() *Jwt {
	return &Jwt{
		Secret:          []byte(`secure`),
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
	return &DB{
		Filename: "sqlite.db",
	}
}
