package config

type Server struct {
	Addr string `env:"HTTP_ADDR" env-default:"0.0.0.0"`
	Port string `env:"HTTP_PORT" env-default:"5000"`
}
