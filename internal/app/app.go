package app

import (
	"fmt"
	"html/template"
	"net"
	"net/http"
	"strconv"

	"github.com/jmoiron/sqlx"

	_ "github.com/mattn/go-sqlite3"
)

type App struct {
	mux       *http.ServeMux
	db        *sqlx.DB
	templates map[string]*template.Template
}

func NewApp() *App {
	return &App{
		mux:       http.NewServeMux(),
		templates: make(map[string]*template.Template),
	}
}

func (a *App) Run(ip string, port int) error {
	var err error
	a.db, err = sqlx.Open("sqlite3", "db/main.db")
	if err != nil {
		return err
	}
	defer a.db.Close()

	if err := a.db.Ping(); err != nil {
		return fmt.Errorf("DB connection failed: %w", err)
	}

	fs := http.FileServer(http.Dir("static"))
	a.mux.Handle("/static/", http.StripPrefix("/static/",
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// w.Header().Set("Cache-Control", "max-age=86400")
			fs.ServeHTTP(w, r)
		})),
	)

	addr := net.JoinHostPort(ip, strconv.Itoa(port))
	fmt.Printf("Server starting at http://%s\n", addr)
	return http.ListenAndServe(addr, a.mux)
}

func (a *App) AddHandlerFunc(path string, f http.HandlerFunc) {
	a.mux.HandleFunc(path, f)
}

func (a *App) DB() *sqlx.DB {
	return a.db
}

func (a *App) LoadTemplate(s string) (*template.Template, error) {
	// t, isExists := a.templates[s]
	// if isExists {
	// 	return t, nil
	// }
	t, err := template.ParseFiles("templates/" + s)
	if err != nil {
		return nil, err
	}
	// a.templates[s] = t
	return t, nil
}
