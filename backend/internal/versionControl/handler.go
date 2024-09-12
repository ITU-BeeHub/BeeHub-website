package versionControl

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	service *Service
	logger  *logrus.Logger
}

func NewHandler(service *Service, logger *logrus.Logger) *Handler {
	return &Handler{service: service, logger: logger}
}

// @Summary Get version information
// @Description Returns the current version of the BeeHub Web API
// @Tags Version
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]string
// @Router /version [get]
func (h *Handler) GetVersion(c *gin.Context) {
	h.logger.Info("Version request received")

	// Service katmanından versiyon bilgisini al
	versionInfo, err := h.service.GetVersionInfo()
	if err != nil {
		h.logger.Error("Error while fetching version info: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch version info"})
		return
	}

	// Versiyon bilgisini döndür
	c.JSON(http.StatusOK, versionInfo)
}
