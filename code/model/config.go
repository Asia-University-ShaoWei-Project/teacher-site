package model

type Config struct {
	JWTSecure      []byte
	PasswordSecure string
	HashCost       int
}
