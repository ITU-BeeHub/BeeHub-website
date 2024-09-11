package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type VersionResponse struct {
	Version     string `json:"version"`
	DownloadURL string `json:"download_url"`
}

func main() {
	// Gin router'ı oluştur
	r := gin.Default()

	// Version endpoint'i
	r.GET("/version", func(c *gin.Context) {
		response := VersionResponse{
			Version:     "1.0.0",                                         //Sample version
			DownloadURL: "https://yourdomain.com/download/installer.exe", //Sample download URL
		}
		c.JSON(http.StatusOK, response)
	})

	// Download endpoint'i
	r.GET("/download/:os", func(c *gin.Context) {
		os := c.Param("os")

		// İşletim sistemine göre dosya yolu belirle
		var filePath string
		switch os {
		case "windows":
			filePath = "./path/to/your/installer.exe"
		case "mac":
			filePath = "./path/to/your/installer.dmg"
		case "linux":
			filePath = "./path/to/your/installer.deb"
		default:
			c.JSON(http.StatusBadRequest, gin.H{"error": "Desteklenmeyen işletim sistemi"})
			return
		}

		c.FileAttachment(filePath, "installer")
	})

	// Sunucuyu başlat
	r.Run(":8080")
}
