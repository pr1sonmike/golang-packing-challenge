package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/pr1sonmike/golang-packing-challenge/config"
	"github.com/pr1sonmike/golang-packing-challenge/internal/handlers"
	"github.com/pr1sonmike/golang-packing-challenge/internal/service"
	"github.com/pr1sonmike/golang-packing-challenge/internal/utils"
)

func main() {
	utils.InitLogger()

	cfg, err := config.LoadConfig()
	if err != nil {
		utils.Logger.Fatalf("Failed to load config: %v", err)
	}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Validator = utils.NewValidator()

	packsService := service.NewPacksService()
	packHandler := handlers.NewPacksHandler(packsService)

	e.Static("/static", "static")
	e.File("/", "templates/index.html")
	e.POST("/packs", packHandler.CalculatePacks)

	go func() {
		if err := e.Start(cfg.Server.Port); err != nil && !errors.Is(err, http.ErrServerClosed) {
			utils.Logger.Fatalf("Shutting down the server: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	utils.Logger.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		utils.Logger.Fatal(err)
	}
}
