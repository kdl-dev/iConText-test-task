package config

type Postgres struct {
	Role     string `env:"PG_ROLE" env-default:"postgres"`
	Password string `env:"PG_PASSWORD" env-default:""`
	Host     string `env:"PG_HOST" env-default:"127.0.0.1"`
	Port     string `env:"PG_PORT" env-default:"5432"`
	DBName   string `env:"PG_DBNAME" env-default:""`
	SSLMode  string `env:"PG_SSLMODE" env-default:"disabled"`
}
