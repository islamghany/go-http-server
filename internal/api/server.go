package api

import (
	"httpserver/internal/config"
	"httpserver/internal/handlers"
	"httpserver/internal/middleware"
	"httpserver/internal/web"
	"log"
	"net/http"
)

func NewServer(
	logger *log.Logger,
	cfg *config.Config,

) http.Handler {

	app := web.NewApp(middleware.Error())

	// Register the routes
	handlers.NewUserHandler(logger, app).RegisterRoutes()

	return app
}
