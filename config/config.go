package config

import (
	"os"
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
	Addr string
	Path Path

	SlidePathFormat    string
	HomeworkPathFormat string
	MaxMultipartMemory int64
}
type Path struct {
	StaticRelative string
	StaticRoot     string
	Template       string
	Document       string
}

func (p *Path) GetDocPath() string {
	return p.StaticRelative + p.Document
}

type Limit struct {
	TeacherListPageCount int
}
type Secure struct {
	Salt              []byte
	SaltSize          int ``
	HashCost          int
	SessionSecret     []byte
	CookieTokenMaxAge int
}
type Jwt struct {
	Secret          []byte
	TokenExpireTime time.Duration
}

type Redis struct {
	MaxReTry           int
	Addr               string
	Password           string
	Database           int
	LastModifiedExpire time.Duration
}
type DB struct {
	Filename string
	Database string
	User     string
	Password string
	Port     string
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
		Addr: ":" + os.Getenv("PORT"),
		Path: Path{
			StaticRelative: "/static",
			StaticRoot:     "./static",
			Template:       "templates/*",
			Document:       "/doc",
		},
		SlidePathFormat:    `static/doc/%s/slide/%s`,
		HomeworkPathFormat: `static/doc/%s/homework/%s`,
		MaxMultipartMemory: 8 << 20,
	}
}
func newDB() *DB {
	return &DB{
		Filename: "sqlite.db",
		Database: "postgres",
		User:     "postgres",
		Password: "postgres",
		Port:     "5432",
	}
}

func newLimit() *Limit {
	return &Limit{
		TeacherListPageCount: 10,
	}
}
func newSecure() *Secure {
	return &Secure{
		SaltSize:          16,
		HashCost:          10,
		SessionSecret:     []byte("secret"),
		CookieTokenMaxAge: (20 * int(time.Minute.Seconds())),
	}
}
func newJwt() *Jwt {
	return &Jwt{
		Secret: []byte(`secure`),
		// minute
		TokenExpireTime: 5,
	}
}
func newRedis() *Redis {
	return &Redis{
		MaxReTry:           2,
		Addr:               ":6379",
		Password:           "",
		Database:           0,
		LastModifiedExpire: 5 * time.Minute,
	}
}
