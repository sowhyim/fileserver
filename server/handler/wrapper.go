package handler

import (
	"fileserver/model"
	"net/http"
)

func UploadWrapper(h func(*http.Request, *model.Userinfo) *model.ResponseCode) http.HandlerFunc {
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

func DownloadWrapper(h func(r *http.Request) (map[string]string, []byte, *model.ResponseCode)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		headers, data, resCode := h(r)
		if resCode.Code != 0 {
			w.Write(resCode.ToJsonBytes())
			return
		}
		for i := range headers {
			w.Header().Set(i, headers[i])
		}
		w.Write(data)
	}
}

func LanderWrapper(h func(r *http.Request) *model.ResponseCode) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res := h(r)
		w.Write(res.ToJsonBytes())
	}
}
