// Package server wires the HTTP routes and middleware.
package server

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/ITU-BeeHub/BeeHub-website/backend/internal/admin"
	"github.com/ITU-BeeHub/BeeHub-website/backend/internal/auth"
	"github.com/ITU-BeeHub/BeeHub-website/backend/internal/download"
	"github.com/ITU-BeeHub/BeeHub-website/backend/internal/httpx"
	"github.com/ITU-BeeHub/BeeHub-website/backend/internal/version"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Pinger reports whether a dependency is reachable.
type Pinger interface {
	PingContext(ctx context.Context) error
}

// Dependencies are the components the routes are served by.
type Dependencies struct {
	Log            *slog.Logger
	TrustedProxies []string
	SwaggerEnabled bool
	Database       Pinger
	Tokens         *auth.TokenManager
	Auth           *auth.Handler
	Download       *download.Handler
	Version        *version.Handler
	Admin          *admin.Handler
}

// NewRouter returns the HTTP handler for the whole API.
func NewRouter(deps Dependencies) (http.Handler, error) {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	if err := r.SetTrustedProxies(deps.TrustedProxies); err != nil {
		return nil, err
	}
	r.Use(gin.Recovery(), requestLogger(deps.Log))

	api := r.Group("/api")
	api.GET("/health", health(deps.Database))
	api.GET("/version", deps.Version.Get)
	api.GET("/download", deps.Download.Download)
	api.HEAD("/download", deps.Download.Download)

	api.POST("/admin/login", auth.RateLimit(12*time.Second, 5), deps.Auth.Login)
	adminAPI := api.Group("/admin", auth.RequireAdmin(deps.Tokens))
	adminAPI.GET("/stats/summary", deps.Admin.Summary)
	adminAPI.GET("/stats/daily", deps.Admin.Daily)
	adminAPI.GET("/stats/monthly", deps.Admin.Monthly)
	adminAPI.GET("/downloads", deps.Admin.Downloads)

	if deps.SwaggerEnabled {
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
	return r, nil
}

// health reports whether the API can serve requests.
//
// @Summary  Health check
// @Tags     Health
// @Produce  json
// @Success  200 {object} map[string]string
// @Failure  503 {object} httpx.ErrorResponse
// @Router   /health [get]
func health(db Pinger) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), 2*time.Second)
		defer cancel()
		if err := db.PingContext(ctx); err != nil {
			httpx.Error(c, http.StatusServiceUnavailable, "database unavailable")
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	}
}

func requestLogger(log *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		log.Info("request",
			"method", c.Request.Method,
			"path", c.Request.URL.Path,
			"status", c.Writer.Status(),
			"duration", time.Since(start).Round(time.Millisecond),
			"ip", c.ClientIP(),
		)
	}
}
