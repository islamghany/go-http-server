package handlers

import (
	"fmt"
	"httpserver/internal/models"
	"httpserver/internal/services"
	"httpserver/internal/web"
	"httpserver/pkg/logger"
	"net/http"
)

type UserHandler struct {
	Logger      *logger.Logger
	WebApp      *web.WebApp
	userService *services.UserService
}

func NewUserHandler(
	logger *logger.Logger,
	webApp *web.WebApp,
	userService *services.UserService,
) *UserHandler {
	return &UserHandler{
		WebApp:      webApp,
		Logger:      logger,
		userService: userService,
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
	var nu models.CreateUserParams
	if err := web.Decode(w, r, &nu); err != nil {
		return web.NewError(err, http.StatusBadRequest)
	}

	n, err := h.userService.CreateUser(r.Context(), nu.Email, []byte(nu.Password))
	if err != nil {
		return web.NewError(err, http.StatusInternalServerError)
	}

	fmt.Println(n)

	return web.Response(w, r, http.StatusCreated, nu)
}
