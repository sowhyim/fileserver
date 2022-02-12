package router

import (
	"fileserver/handler"
	"net/http"
)

func upload() {
	http.Handle("/create", handler.UploadWrapper(handler.CreateUploadTask))
	http.Handle("/upload", handler.UploadWrapper(handler.DealUploadTask))
	http.Handle("/merge", handler.UploadWrapper(handler.MergeUploadTask))
}
