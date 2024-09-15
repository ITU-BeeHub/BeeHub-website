package adminpanel

import (
	"net/http"

	download "github.com/ITU-BeeHub/BeeHub-website/backend/internal/downloadManager"
	"github.com/ITU-BeeHub/BeeHub-website/backend/pkg/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// AdminDownloadStatsHandler indirme istatistiklerini JWT ile korur.
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

// AdminMonthlyDownloadStatsHandler aylık indirme istatistiklerini döner
// @Summary Get monthly download statistics
// @Description Returns the monthly download statistics, accessible only by admin users.
// @Tags Admin
// @Security ApiKeyAuth
// @Produce json
// @Success 200 {array} models.MonthlyDownloadStats "Returns monthly download statistics"
// @Failure 401 {object} models.ErrorResponse "Unauthorized"
// @Router /admin/monthly-download-stats [get]
func AdminMonthlyDownloadStatsHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		stats, err := download.GetMonthlyDownloadStats(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Could not fetch monthly download stats"})
			return
		}

		c.JSON(http.StatusOK, stats)
	}
}

// AdminIPLogsHandler IP loglarını döner
// @Summary Get IP logs
// @Description Returns recent IP logs, accessible only by admin users.
// @Tags Admin
// @Security ApiKeyAuth
// @Produce json
// @Success 200 {array} models.Download "Returns list of downloads with IP and OS"
// @Failure 401 {object} models.ErrorResponse "Unauthorized"
// @Router /admin/ip-logs [get]
func AdminIPLogsHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		logs, err := download.GetIPLogs(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Could not fetch IP logs"})
			return
		}

		c.JSON(http.StatusOK, logs)
	}
}
