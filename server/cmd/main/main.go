package main

import (
	"fileserver/router"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("starting...")
	router.Router()

	fmt.Println("started, listen port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
