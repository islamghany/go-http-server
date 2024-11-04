package api

import (
	"httpserver/internal/config"
	db "httpserver/internal/db/sqlc"
	"httpserver/internal/handlers"
	"httpserver/internal/logger"
	"httpserver/internal/middleware"
	"httpserver/internal/services"
	"httpserver/internal/web"
	"net/http"
)

func NewServer(
	logger *logger.Logger,
	cfg *config.Config,
	store *db.Queries,

) http.Handler {

	app := web.NewApp(middleware.Logger(logger), middleware.Error())

	// Creating the services
	userService := services.NewUserService(store)

	// Register the routes
	handlers.NewUserHandler(logger, app, userService).RegisterRoutes()

	return app
}
