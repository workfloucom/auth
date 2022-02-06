package test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"gorm.io/gorm"
	"workflou.com/auth/internal/application"
)

var (
	app    *application.Application
	srv    *httptest.Server
	client *http.Client
)

func TestMain(m *testing.M) {
	app = application.New(application.Config{
		InfoLogOutput:  io.Discard,
		ErrorLogOutput: io.Discard,
		Driver:         "sqlite",
		Dsn:            "file::memory:?cache=shared",
		Env:            "test",
		AuthSecret:     "AUTH_SECRET",
		RefreshSecret:  "REFRESH_SECRET",
	})

	app.DB.Migrate()

	srv = httptest.NewServer(app.Handler())
	defer srv.Close()

	client = srv.Client()
	os.Exit(m.Run())
}

func Create(model interface{}) *gorm.DB {
	return app.DB.Connection.Create(model)
}

func Get(url string) (*http.Response, error) {
	return client.Get(srv.URL + url)
}

func Post(url string, body io.Reader) (*http.Response, error) {
	return client.Post(srv.URL+url, "application/json", body)
}

func Patch(url string, body io.Reader) (*http.Response, error) {
	r, err := http.NewRequest(http.MethodPatch, srv.URL+url, body)

	if err != nil {
		return nil, err
	}

	return client.Do(r)
}

func Delete(url string) (*http.Response, error) {
	r, err := http.NewRequest(http.MethodDelete, srv.URL+url, nil)

	if err != nil {
		return nil, err
	}

	return client.Do(r)
}
