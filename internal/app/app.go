package app

import (
	"fmt"
	"net"
	"net/http"
	"strconv"
)

type App struct {
	addr string
	mux  *http.ServeMux
}

func NewApp() App {
	return App{mux: http.NewServeMux()}
}

func (a *App) Run(ip string, port int) error {
	addr := net.JoinHostPort(ip, strconv.Itoa(port))
	fmt.Printf("Server starting at http://%s\n", addr)
	return http.ListenAndServe(addr, a.mux)
}
