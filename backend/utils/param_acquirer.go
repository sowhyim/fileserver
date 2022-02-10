package utils

import (
	"net/http"
	"strconv"
)

func GetFormString(r *http.Request, key string) (string, bool) {
	s := r.Form.Get(key)
	if s == "" {
		return s, false
	}
	return s, true
}

func GetFormInt(r *http.Request, key string) (int, bool) {
	s := r.Form.Get(key)
	if s == "" {
		return 0, false
	}

	if i, err := strconv.Atoi(s); err == nil {
		return i, true
	}
	return 0, false
}

func GetFormInt64(r *http.Request, key string) (int64, bool) {
	s := r.Form.Get(key)
	if s == "" {
		return 0, false
	}

	if i64, err := strconv.ParseInt(s, 10, 64); err == nil {
		return i64, true
	}
	return 0, false
}
