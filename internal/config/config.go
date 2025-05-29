package config

type Config struct {
	Auth
	Postgres
	Redis
}

type Auth struct {
	Secret string
}

type Postgres struct {
	DSN string
}

type Redis struct {
	Port     string
	Password string
}
