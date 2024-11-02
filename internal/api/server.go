package api

import (
	"httpserver/internal/config"
	db "httpserver/internal/db/sqlc"
	"httpserver/internal/handlers"
	"httpserver/internal/middleware"
	"httpserver/internal/services"
	"httpserver/internal/web"
	"log"
	"net/http"
)

func NewServer(
	logger *log.Logger,
	cfg *config.Config,
	store *db.Queries,

) http.Handler {

	app := web.NewApp(middleware.Error())

	// Creating the services
	userService := services.NewUserService(store)

	// Register the routes
	handlers.NewUserHandler(logger, app, userService).RegisterRoutes()

	return app
}
