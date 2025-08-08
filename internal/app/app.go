package app

import (
	"database/sql"
	"fmt"
	"net"
	"net/http"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

type App struct {
	addr string
	mux  *http.ServeMux
	db   *sql.DB
}

func NewApp() *App {
	return &App{mux: http.NewServeMux()}
}

func (a *App) Run(ip string, port int) error {
	var err error
	a.db, err = sql.Open("sqlite3", "db/main.db")
	if err != nil {
		return err
	}
	defer a.db.Close()
	a.addr = net.JoinHostPort(ip, strconv.Itoa(port))
	fmt.Printf("Server starting at http://%s\n", a.addr)
	return http.ListenAndServe(a.addr, a.mux)
}

func (a *App) Addr() string {
	return a.addr
}

func (a *App) AddHandlerFunc(path string, f http.HandlerFunc) {
	a.mux.HandleFunc(path, f)
}
