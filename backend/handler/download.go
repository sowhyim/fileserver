package handler

import (
	"fileserver/model"
	"fmt"
	"net/http"
)

func Download(r *http.Request) (map[string]string, []byte, *model.ResponseCode) {
	// TODO find file and make sure OK

	var header = map[string]string{
		"Content-Type":              "application/octet-stream",
		"Content-Disposition":       fmt.Sprintf("attachment;filename=%s", ""),
		"Content-Transfer-Encoding": "binary",
		"Expires":                   "0",
	}
	return header, nil, model.ResponseCodeOK
}

// TODO MultipartDownload
