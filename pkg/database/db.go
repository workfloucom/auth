package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"workflou.com/auth/pkg/user"
)

type DB struct {
	Connection *gorm.DB
	Config     Config
}

func New(cfg Config) *DB {
	var d gorm.Dialector

	switch cfg.Driver {
	case "sqlite":
		d = sqlite.Open(cfg.Dsn)
	case "mysql":
		d = mysql.Open(cfg.Dsn)
	}

	conn, err := gorm.Open(d, &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return &DB{
		Connection: conn,
		Config:     cfg,
	}
}

func (db *DB) Migrate() {
	db.Connection.AutoMigrate(&user.User{})
}
