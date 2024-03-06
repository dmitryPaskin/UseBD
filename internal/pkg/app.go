package app

import (
	"BDServer/internal/controller"
	"BDServer/internal/pkg/db"
	"BDServer/internal/pkg/router"
	"BDServer/internal/repository"
	"BDServer/internal/service"
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func StartServer() {
	dataBase, err := dataBase.New()
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}
	defer dataBase.DB.Close()

	userRepository := repository.New(dataBase.DB)
	userService := service.New(userRepository)
	userController := controller.New(userService)

	router := router.New(userController)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      router.Mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	sigChan := make(chan os.Signal, 1)
	defer close(sigChan)
	signal.Notify(sigChan, syscall.SIGINT)

	go func() {
		log.Println("Starting server...")
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Server error: %v", err)
		}
	}()

	<-sigChan
	log.Println("Server stopping...")
	stopCTX, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	if err := server.Shutdown(stopCTX); err != nil {
		log.Fatalf("Server shutdown error: %v", err)
	}
}
