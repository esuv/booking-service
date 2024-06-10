package app

import (
	"booking-service/internal/config"
	"booking-service/internal/logger"
	"booking-service/internal/repository"
	"booking-service/internal/service"
	"booking-service/internal/transport/rest/handler"
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run(configDir string) {
	//config
	cfg := config.MustLoad(configDir)

	//init logger
	var defaultLogger = logger.NewLogger()

	//clean architecture: handler -> service -> repository

	//init repository
	bookingRepository := repository.NewBookingInMemoryRepo(defaultLogger)

	//init service
	bookingService := service.NewBookingService(bookingRepository, defaultLogger)

	//init handler
	bookingHandler := handler.NewBookingHandler(&bookingService, defaultLogger)

	// create http router
	mux := http.NewServeMux()
	mux.HandleFunc("/orders", bookingHandler.Post)

	srv := &http.Server{
		Addr:    cfg.Http.HTTPAddr,
		Handler: mux,
	}

	// listen to OS signals and gracefully shutdown HTTP server
	stopped := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
		<-sigint
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			defaultLogger.LogErrorf("HTTP Server Shutdown Error:%v", err)
		}
		close(stopped)
	}()

	defaultLogger.LogInfo(fmt.Sprintf("Starting HTTP server on %s", cfg.Http.HTTPAddr))

	// start server
	if err := srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		defaultLogger.LogErrorf("HTTP server ListenAndServe Error: %v", err)
	}

	<-stopped

	defaultLogger.LogInfo("Have a nice day!")
}
