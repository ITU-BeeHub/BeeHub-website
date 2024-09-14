package adminpanel

import (
	"net/http"

	download "github.com/ITU-BeeHub/BeeHub-website/backend/internal/downloadManager"
	"github.com/ITU-BeeHub/BeeHub-website/backend/pkg/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// AdminDownloadStatsHandler döndürülen indirme istatistiklerini JWT ile korur.
// @Summary Get download statistics
// @Description Returns the download statistics for the application, accessible only by admin users.
// @Tags Admin
// @Security ApiKeyAuth
// @Produce json
// @Success 200 {object} map[string]int "Returns download statistics grouped by OS"
// @Failure 401 {object} models.ErrorResponse "Unauthorized"
// @Router /admin/download-stats [get]
func AdminDownloadStatsHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Sadece doğrulanmış admin'ler bu işlemi yapabilir
		stats, err := download.GetDownloadStats(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Could not fetch download stats"})
			return
		}

		c.JSON(http.StatusOK, stats)
	}
}
