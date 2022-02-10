package router

import (
	"fileserver/handler"
	"net/http"
)

func upload() {
	http.Handle("/upload", handler.Authentication(handler.DealUploadTask))
	http.Handle("/merge", handler.Authentication(handler.MergeUploadTask))
}
