package handlers

import (
	"httpserver/internal/models"
	"httpserver/internal/web"
	"log"
	"net/http"
)

type UserHandler struct {
	Logger *log.Logger
	WebApp *web.WebApp
}

func NewUserHandler(
	logger *log.Logger,
	webApp *web.WebApp,
) *UserHandler {
	return &UserHandler{
		WebApp: webApp,
		Logger: logger,
	}
}

func (h *UserHandler) RegisterRoutes() {
	v1 := "v1"
	h.WebApp.Handle(http.MethodGet, v1, "/users", h.GetUsers)
	h.WebApp.Handle(http.MethodGet, v1, "/users/:id", h.GetUser)
	h.WebApp.Handle(http.MethodPost, v1, "/users", h.CreateUser)
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) error {
	s := []string{"user1", "user2", "user3"}
	return web.Response(w, r, http.StatusOK, s)
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) error {
	userId, err := web.GetParamUUID(r, "id")
	if err != nil {
		return web.NewError(err, http.StatusBadRequest)
	}
	return web.Response(w, r, http.StatusOK, userId)
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) error {
	var nu models.User
	if err := web.Decode(w, r, &nu); err != nil {
		return web.NewError(err, http.StatusBadRequest)
	}
	return web.Response(w, r, http.StatusCreated, nu)
}
