package api

import (
	handlersV1 "httpserver/api/v1"
	"httpserver/internal/config"
	db "httpserver/internal/db/sqlc"
	"httpserver/internal/middleware"
	"httpserver/internal/services"
	"httpserver/internal/web"
	"httpserver/pkg/logger"
	"net/http"
)

func NewServer(
	logger *logger.Logger,
	cfg *config.Config,
	store *db.Store,

) http.Handler {

	app := web.NewApp(middleware.Logger(logger), middleware.Error(logger), middleware.Panic())

	// Creating the services
	userService := services.NewUserService(store)

	// Register the routes
	handlersV1.NewUserHandler(logger, app, userService).RegisterRoutes()

	return app
}
