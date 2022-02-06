package database

type Config struct {
	Env             string
	Driver          string
	Dsn             string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxIdleTime string
}
