package config

type Database struct {
	Host string
	Port int
	User string
	Pass string
	DB   string
}

type Config struct {
	Database
}
