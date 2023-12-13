package config

import "os"

const HTTPAddr = ":8180"

type Config struct {
	HTTP HTTPConfig
	DB   DBConfig
}

type DBConfig struct {
	User     string
	Password string
	Host     string
	Port     int
	Dbname   string
	Sslmode  string
}

type ContextKey string

type HTTPConfig struct {
	HostAddr   string
	ContextKey ContextKey
}

func GetConfig() *Config {

	return &Config{

		DB: DBConfig{
			Dbname:   "url",
			User:     "user",
			Password: "user",
			Host:     "postgres", //"localhost",
			Port:     5432,
			Sslmode:  "",
		},

		HTTP: HTTPConfig{
			HostAddr:   os.Getenv("HOST_ADDR"),
			ContextKey: "History",
		},
	}

}
