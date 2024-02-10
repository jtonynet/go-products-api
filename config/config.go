package config

type API struct {
	Port int
}
type Database struct {
	Host string
	Port int
	User string
	Pass string
	DB   string
}

type Config struct {
	API
	Database
}
