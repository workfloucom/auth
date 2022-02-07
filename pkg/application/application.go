package application

import (
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"workflou.com/auth/internal/orm"
	"workflou.com/auth/pkg/database"
	"workflou.com/auth/pkg/link"
	"workflou.com/auth/pkg/validation"
)

type Application struct {
	InfoLog       *log.Logger
	ErrorLog      *log.Logger
	DB            *database.DB
	Validate      *validator.Validate
	AuthSecret    string
	RefreshSecret string
}

func New(cfg Config) *Application {
	return &Application{
		AuthSecret:    cfg.AuthSecret,
		RefreshSecret: cfg.RefreshSecret,
		InfoLog:       log.New(cfg.InfoLogOutput, "INFO\t", log.LstdFlags),
		ErrorLog:      log.New(cfg.ErrorLogOutput, "ERROR\t", log.LstdFlags),
		Validate:      validation.New(),
		DB: database.New(database.Config{
			Env:             cfg.Env,
			Driver:          cfg.Driver,
			Dsn:             cfg.Dsn,
			MaxOpenConns:    cfg.MaxOpenConns,
			MaxIdleConns:    cfg.MaxIdleConns,
			ConnMaxIdleTime: cfg.ConnMaxIdleTime,
		}),
	}
}

func (app *Application) Handler() http.Handler {
	users := &orm.Users{DB: app.DB.Connection}

	r := mux.NewRouter()
	r.Handle("/link", link.Create{
		Validate: *app.Validate,
		Users:    users,
	}).Methods(http.MethodPost)

	ar := r.NewRoute().Subrouter()
	ar.Use(app.Authenticated)

	return r
}

func (app *Application) Authenticated(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// tokenHeader := r.Header.Get("Authorization")

		// if tokenHeader == "" {
		// 	w.WriteHeader(http.StatusUnauthorized)
		// 	return
		// }

		// token, err := jwt.Parse(tokenHeader, func(t *jwt.Token) (interface{}, error) {
		// 	if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
		// 		return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		// 	}

		// 	return []byte(app.JwtSecret), nil
		// })

		// if err != nil {
		// 	w.WriteHeader(http.StatusUnauthorized)
		// 	return
		// }

		// claims, ok := token.Claims.(jwt.MapClaims)

		// if !ok || !token.Valid {
		// 	w.WriteHeader(http.StatusUnauthorized)
		// 	return
		// }

		// sub := claims["sub"].(float64)

		// users := &gormdb.UserRepository{DB: app.DB.Connection}
		// u, err := users.FindByID(uint(sub))

		// if err != nil {
		// 	w.WriteHeader(http.StatusUnauthorized)
		// 	return
		// }

		// ctx := context.WithValue(r.Context(), auth.UserContextKey, u)

		// next.ServeHTTP(w, r.WithContext(ctx))
		next.ServeHTTP(w, r)
	})
}
