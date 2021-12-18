package main

import (
	"card/config"
	"card/routes"
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
)

func main() {
	config := config.LoadConfig()

	if config.Port == "" {
		config.Port = "8000"
	}

	e := echo.New()

	// Log all requests
	e.Use(echomiddleware.Logger())

	// Recover on panic
	e.Use(echomiddleware.Recover())

	routes.Run(e)

	go func() {
		e.Logger.Fatal(e.Start(fmt.Sprintf("0.0.0.0:%s", config.Port)))
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}

}
