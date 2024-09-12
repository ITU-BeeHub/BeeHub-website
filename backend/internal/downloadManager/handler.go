package downloadManager

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Tags Download Manager
// @Summary download manager
// @Accept json
// @Produce json
// @Router /download [get]
func DownloadHandler(c *gin.Context) {
	os := c.Param("os")
	var filePath string
	switch os {
	case "windows":
		filePath = "./path/to/your/installer.exe"
	case "mac":
		filePath = "./path/to/your/installer.dmg"
	case "linux":
		filePath = "./path/to/your/installer.deb"
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unsupported OS"})
		return
	}
	c.File(filePath)
}
