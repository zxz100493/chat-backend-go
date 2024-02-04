package server

import (
	"chat-go/config"
	"chat-go/infrastructurre/log"
	"chat-go/route"
	"fmt"
	"net/http"
	"time"
)

const readTimeout = 10 * time.Second

const maxHeaderBytes = 1 << 20

func StartServer() {
	config.Init("./")
	cfg := config.Instance

	log.InitLog("Debug", cfg.Log.Path)

	fmt.Printf("cfg:%#+v", cfg)

	log.Logger.Info("start server", log.String("start", "start web sever..."))

	r := route.NewRoute()

	s := &http.Server{
		Addr:           cfg.AppPort,
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
