package main

import (
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"

	"goshop/restful/config"
	"goshop/restful/routers"
)

func main() {
	s := &http.Server{
		Addr:           ":" + config.ServerConf.Port,
		Handler:        routers.Routers(),
		ReadTimeout:    120 * time.Second,
		WriteTimeout:   120 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Info("start server  on ", config.ServerConf.Port)

	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
