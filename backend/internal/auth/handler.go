package auth

import (
	"net/http"

	"github.com/ITU-BeeHub/BeeHub-website/backend/pkg/models"
	"github.com/gin-gonic/gin"
)

// AdminLoginRequest, giriş isteğini temsil eder
type AdminLoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// AdminHandler, HTTP isteklerini işler
type AdminHandler struct {
	service *AdminService
}

func NewAdminHandler(service *AdminService) *AdminHandler {
	return &AdminHandler{service: service}
}

// LoginHandler, admin kullanıcılarının giriş yapmasına olanak tanır.
// @Summary Admin login
// @Description Authenticates an admin user and returns a JWT token.
// @Tags Auth
// @Accept json
// @Produce json
// @Param admin body auth.AdminLoginRequest true "Admin login credentials"
// @Success 200 {object} models.TokenResponse "JWT token"
// @Failure 400 {object} models.ErrorResponse "Invalid request"
// @Failure 401 {object} models.ErrorResponse "Invalid username or password"
// @Router /admin/login [post]
func (h *AdminHandler) LoginHandler(c *gin.Context) {
	var req AdminLoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid request"})
		return
	}

	// Servis katmanına yönlendirme yap
	token, err := h.service.AuthenticateAdmin(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{Error: "Invalid username or password"})
		return
	}

	// Başarılı giriş durumunda JWT token döner
	c.JSON(http.StatusOK, models.TokenResponse{Token: token})
}
