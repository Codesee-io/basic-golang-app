package main

import (
	"context"
	"net/http"

	"github.com/codesee/basic-golang-app/pkg/logger"
	"github.com/codesee/basic-golang-app/pkg/server"
	"github.com/codesee/basic-golang-app/pkg/signals"
)

func main() {
	log := logger.New()
	port := 2345

	srv, err := server.New(port)
	if err != nil {
		log.Err(err).Fatal("server error")
	}

	graceful := signals.Setup()

	go func() {
		log.Info("server started", logger.Data{"port": port})
		err := srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Err(err).Fatal("server stopped")
		}
		log.Info("server stopped")
	}()

	<-graceful
	log.Info("starting graceful shutdown")
	ctx := context.Background()

	err = srv.Shutdown(ctx)
	if err != nil {
		log.Err(err).Error("server shutdown error")
	}
	log.Info("server shutdown")
}
