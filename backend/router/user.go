package router

import (
	"fileserver/handler"
	"net/http"
)

func user() {
	http.Handle("/regestry", handler.LanderWrapper(handler.Regestry))
	http.Handle("/login", handler.LanderWrapper(handler.Login))
	http.Handle("/check_alive", handler.LanderWrapper(handler.CheckAlive))
}
