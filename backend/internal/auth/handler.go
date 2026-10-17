package auth

import (
	"errors"
	"log/slog"
	"net/http"
	"time"

	"github.com/ITU-BeeHub/BeeHub-website/backend/internal/httpx"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
	log     *slog.Logger
}

func NewHandler(service *Service, log *slog.Logger) *Handler {
	return &Handler{service: service, log: log}
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expiresAt"`
}

// Login exchanges admin credentials for a session token.
//
// @Summary  Admin login
// @Tags     Auth
// @Accept   json
// @Produce  json
// @Param    credentials body LoginRequest true "Admin credentials"
// @Success  200 {object} LoginResponse
// @Failure  400 {object} httpx.ErrorResponse
// @Failure  401 {object} httpx.ErrorResponse
// @Failure  429 {object} httpx.ErrorResponse
// @Router   /admin/login [post]
func (h *Handler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		httpx.Error(c, http.StatusBadRequest, "username and password are required")
		return
	}

	session, err := h.service.Login(c.Request.Context(), req.Username, req.Password)
	switch {
	case errors.Is(err, ErrInvalidCredentials):
		h.log.Warn("admin login failed", "username", req.Username, "ip", c.ClientIP())
		httpx.Error(c, http.StatusUnauthorized, "invalid username or password")
		return
	case err != nil:
		h.log.Error("admin login", "error", err)
		httpx.Error(c, http.StatusInternalServerError, "internal error")
		return
	}

	h.log.Info("admin logged in", "username", req.Username, "ip", c.ClientIP())
	c.JSON(http.StatusOK, LoginResponse{Token: session.Token, ExpiresAt: session.ExpiresAt})
}
