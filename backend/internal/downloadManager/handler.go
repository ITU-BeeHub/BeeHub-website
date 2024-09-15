package downloadManager

import (
	"fmt"
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

	// Kullanıcının IP adresini alıyoruz
	clientIP := c.ClientIP()

	// İndirme bilgilerini kaydediyoruz
	if err := h.service.SaveDownload(c, os, clientIP); err != nil {
		h.logger.Error("Error while saving download info: ", err)
	}

	// Doğru dosya adını belirlemek için dosya adı ekliyoruz
	var fileName string
	switch os {
	case "windows":
		fileName = "installer.exe"
	case "mac":
		fileName = "installer.dmg"
	}

	// Dosya indirilmeye sunulur
	h.logger.Infof("File found: %s, serving to client", filePath)
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))
	c.File(filePath)
}
