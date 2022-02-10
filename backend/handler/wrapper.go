package handler

import (
	"fileserver/model"
	"net/http"
)

func Authentication(h func(*http.Request, *model.Userinfo) *model.ResponseCode) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		// _, ok := utils.GetFormString(r, "token")
		// if !ok {
		// 	w.Write(model.ResponseCodeMissingParam.ToJsonBytes())
		// 	return
		// }

		// TODO verify

		res := h(r, &model.Userinfo{StoragePath: "/home/xx/storage/test"})
		w.Write(res.ToJsonBytes())
	}
}

func DefaultWrapper(h func(r *http.Request) *model.ResponseCode) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res := h(r)
		w.Write(res.ToJsonBytes())
	}
}
