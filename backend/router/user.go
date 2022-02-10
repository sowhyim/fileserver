package router

import (
	"fileserver/handler"
	"net/http"
)

func user() {
	http.Handle("/regestry", handler.DefaultWrapper(handler.Regestry))
	http.Handle("/login", handler.DefaultWrapper(handler.Login))
	http.Handle("/check_alive", handler.DefaultWrapper(handler.CheckAlive))
}
