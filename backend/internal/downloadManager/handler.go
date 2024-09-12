package downloadManager

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
	return &Handler{
		service: service,
		logger:  logger,
	}
}

// @Tags Download Manager
// @Summary Download installer for the selected OS
// @Accept json
// @Produce json
// @Param os query string true "Operating System (windows, mac, linux)"
// @Success 200 {file} file "Returns the installer file"
// @Router /download [get]
func (h *Handler) Download(c *gin.Context) {
	// İstekten OS parametresi alınır
	os := c.Query("os")
	h.logger.Infof("Download request for OS: %s", os)
	if os == "" {
		h.logger.Warn("OS parameter is missing")
		c.JSON(http.StatusBadRequest, gin.H{"error": "OS parameter is required"})
		return
	}

	// Service katmanına yönlendirilir
	filePath, err := h.service.GetInstallerFilePath(os)
	if err != nil {
		h.logger.Error("Error while getting file path: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Dosya indirilmeye sunulur
	h.logger.Infof("File found: %s, serving to client", filePath)
	c.File(filePath)
}
