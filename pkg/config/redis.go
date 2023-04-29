package config

type Redis struct {
	Host     string `env:"RDS_HOST" env-default:"127.0.0.1:6379"`
	User     string `env:"RDS_USER" env-default:""`
	Password string `env:"RDS_PASSWORD" env-default:""`
	DB       int    `env:"RDS_DB" env-default:""`
}
