package notification

import (
	"net/http"
	"time"

	"github.com/ITU-BeeHub/BeeHub-website/backend/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// GetActiveNotifications returns notifications that are currently active.
func (h *Handler) GetActiveNotifications() gin.HandlerFunc {
	return func(c *gin.Context) {
		dbInterface, exists := c.Get("db")
		if !exists {
			h.logger.Error("Database connection not found in context")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "database connection not found"})
			return
		}

		// Type assert the database connection
		db, ok := dbInterface.(*gorm.DB)
		if !ok {
			h.logger.Error("Failed to cast database connection")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid database connection"})
			return
		}

		// Log the incoming request.
		h.logger.Infof("Request: %s %s", c.Request.Method, c.Request.URL.Path)

		var notifications []models.Notification
		now := time.Now()

		// Query notifications that are active and within the valid date range.
		if err := db.Where("active = ? AND start_date <= ? AND end_date >= ?", true, now, now).Find(&notifications).Error; err != nil {
			h.logger.Errorf("Error fetching notifications: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch notifications"})
			return
		}

		// Log the successful response.
		h.logger.Infof("Response: %d active notifications returned", len(notifications))
		c.JSON(http.StatusOK, notifications)
	}
}

type Handler struct {
	db     *gorm.DB
	logger *logrus.Logger
}

// NewHandler creates a new instance of NotificationHandler.
func NewHandler(db *gorm.DB, logger *logrus.Logger) *Handler {
	return &Handler{db: db, logger: logger}
}
