package main

import (
	"fmt"
	"os"

	"github.com/ITU-BeeHub/BeeHub-website/backend/internal/downloadManager"
	"github.com/ITU-BeeHub/BeeHub-website/backend/internal/versionControl"
	"github.com/ITU-BeeHub/BeeHub-website/backend/pkg/config"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/ITU-BeeHub/BeeHub-website/backend/docs"
)

// @title BeeHub Web Site API
// @version 1.0
// @description Bu, BeeHub web sitesi için API belgeleridir.

func main() {
	// Config ve logger'ı yükle
	cfg := config.LoadConfig()
	logger := config.SetupLogger()

	// Logger'a başlangıç mesajı yaz
	logger.Info("BeeHub backend başlatılıyor...")

	// DownloadService ve handler'ı oluştur (dependency injection)
	downloadService := downloadManager.NewService(logger)
	downloadHandler := downloadManager.NewHandler(downloadService, logger)

	// Gin router'ını başlat
	r := gin.Default()

	if os.Getenv("SWAGGER_ENABLED") == "true" {
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	} else {
		fmt.Println("Swagger is disabled")
	}

	// Basit bir route ekleyelim
	r.GET("/ping", func(c *gin.Context) {
		logger.Info("Ping route'una istek geldi")
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// VersionService ve handler'ı oluştur (dependency injection)
	versionService := versionControl.NewService(logger)
	versionHandler := versionControl.NewHandler(versionService, logger)

	// Download route'u
	r.GET("/download", downloadHandler.Download)

	// Version route'u
	r.GET("/version", versionHandler.GetVersion)

	// Sunucuyu başlat
	port := fmt.Sprintf(":%s", cfg.Port)
	logger.Infof("Sunucu %s portunda başlatılıyor...", cfg.Port)
	r.Run(port)
}
