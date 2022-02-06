package main

import (
	"flag"
	"net/http"
	"os"
	"time"

	"workflou.com/auth/pkg/application"
)

func main() {
	cfg := application.Config{
		InfoLogOutput:  os.Stdout,
		ErrorLogOutput: os.Stdout,
		AuthSecret:     os.Getenv("AUTH_SECRET"),
		RefreshSecret:  os.Getenv("REFRESH_SECRET"),
		Dsn:            os.Getenv("DB_DSN"),
		Driver:         os.Getenv("DB_DRIVER"),
	}

	flag.StringVar(&cfg.Addr, "addr", ":4000", "Server address")
	flag.StringVar(&cfg.ConnMaxIdleTime, "conn-max-idle-time", "15m", "Database connection max idle time")
	flag.IntVar(&cfg.MaxOpenConns, "max-open-conns", 50, "Max open database connections")
	flag.IntVar(&cfg.MaxIdleConns, "max-idle-conns", 50, "Max idle database connections")
	flag.Parse()

	app := application.New(cfg)

	srv := http.Server{
		Addr:         cfg.Addr,
		Handler:      app.Handler(),
		ErrorLog:     app.ErrorLog,
		IdleTimeout:  time.Minute,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 30,
	}

	app.InfoLog.Printf("Listening on address %s\n", cfg.Addr)
	app.ErrorLog.Print(srv.ListenAndServe())
}
