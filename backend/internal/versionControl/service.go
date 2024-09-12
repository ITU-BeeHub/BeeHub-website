package versionControl

import "github.com/gin-gonic/gin"

// @Summary Get version information
// @Description Returns the current version of the BeeHub Web API
// @Tags Version
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]string
// @Router /version [get]
func VersionHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"version":      "1.0.0",
		"download_url": "https://yourdomain.com/download/installer.exe",
	})
}
