package application

import "io"

type Config struct {
	Env             string
	Addr            string
	Driver          string
	Dsn             string
	AuthSecret      string
	RefreshSecret   string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxIdleTime string
	InfoLogOutput   io.Writer
	ErrorLogOutput  io.Writer
}
