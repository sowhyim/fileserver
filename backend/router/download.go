package router

import (
	"fileserver/handler"
	"net/http"
)

func download() {
	http.Handle("/download", handler.DownloadWrapper(handler.Download))
}
