package main

import (
	"chat-go/route"
	"net/http"
	"time"
)

func main() {
	r := route.NewRoute()
	s := &http.Server{
		Addr:           ":8000",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
