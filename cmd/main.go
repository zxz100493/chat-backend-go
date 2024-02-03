package main

import (
	"chat-go/infrastructurre/log"
	"chat-go/route"
	"net/http"
	"time"
)

const readTimeout = 10 * time.Second

const maxHeaderBytes = 1 << 20

func main() {
	log.InitLog("Debug")
	log.Logger.Info("start server", log.String("start", "start web sever..."))

	r := route.NewRoute()

	s := &http.Server{
		Addr:           ":8866",
		Handler:        r,
		ReadTimeout:    readTimeout,
		WriteTimeout:   readTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}
	err := s.ListenAndServe()

	if err != nil {
		log.Logger.Info("start server error", log.String("error", err.Error()))
	}
}
