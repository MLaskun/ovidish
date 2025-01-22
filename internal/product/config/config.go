package config

type Config struct {
	Address  string
	Database Database
}

type Database struct {
	Dsn string
}
