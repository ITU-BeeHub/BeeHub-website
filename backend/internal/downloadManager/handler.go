package downloadManager

import (
	"net/http"

	"github.com/ITU-BeeHub/BeeHub-website/backend/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
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

	// Kullanıcının IP adresini alıyoruz
	clientIP := c.ClientIP()

	// İndirme bilgilerini kaydediyoruz
	if err := h.service.SaveDownload(c, os, clientIP); err != nil {
		h.logger.Error("Error while saving download info: ", err)
	}

	// Dosya indirilmeye sunulur
	h.logger.Infof("File found: %s, serving to client", filePath)
	c.File(filePath)
}

// DownloadStatsHandler, tüm kullanıcılar tarafından erişilebilen indirilen dosya istatistiklerini döner.
// @Summary Get download statistics
// @Description Returns download statistics for the application.
// @Tags Download Manager
// @Produce json
// @Success 200 {object} map[string]int "Download statistics grouped by OS"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /download-stats [get]
func DownloadStatsHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		stats, err := GetDownloadStats(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Could not fetch download stats"})
			return
		}

		c.JSON(http.StatusOK, stats)
	}
}
