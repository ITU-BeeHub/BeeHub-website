package main

import (
	"fmt"
	"os"

	_ "github.com/ITU-BeeHub/BeeHub-website/backend/docs"
	adminpanel "github.com/ITU-BeeHub/BeeHub-website/backend/internal/adminPanel"
	"github.com/ITU-BeeHub/BeeHub-website/backend/internal/auth"
	"github.com/ITU-BeeHub/BeeHub-website/backend/internal/downloadManager"
	"github.com/ITU-BeeHub/BeeHub-website/backend/internal/versionControl"
	"github.com/ITU-BeeHub/BeeHub-website/backend/pkg/config"
	"github.com/ITU-BeeHub/BeeHub-website/backend/pkg/models"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	_ "modernc.org/sqlite"
)

// @title BeeHub Web Site API
// @version 1.0
// @description Bu, BeeHub web sitesi için API belgeleridir.

func main() {
	// Config ve logger'ı yükle
	cfg := config.LoadConfig()
	logger := config.SetupLogger(cfg)

	// Logger'a başlangıç mesajı yaz
	logger.Info("BeeHub backend başlatılıyor...")

	// Gin router'ını başlat
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "http://localhost", "https://localhost:443", "https://localhost"}, // İzin verilen domain'ler
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// PostgreSQL bağlantısı için DSN (Data Source Name) oluşturuyoruz
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)

	// PostgreSQL için Gorm Open kullanımı
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		//logger.Fatal("Database connection failed: ", err)
	}

	// Veritabanı modellerini migrate ederek tabloları oluşturun
	err = db.AutoMigrate(&models.Download{}, &models.Admin{}) // Admin ve Download tablolarını oluşturuyoruz
	if err != nil {
		//logger.Fatal("Failed to migrate database models: ", err)
	}

	// Veritabanı middleware'ini ekleyin
	r.Use(DBMiddleware(db))
	// Authentication ve admin paneli yapılandırması
	adminService := auth.NewAdminService(db)
	adminHandler := auth.NewAdminHandler(adminService)

	// DownloadService ve handler'ı oluştur (dependency injection)
	downloadService := downloadManager.NewService(logger)
	downloadHandler := downloadManager.NewHandler(downloadService, logger)

	r.POST("/api/admin/login", adminHandler.LoginHandler)

	// Admin paneli için middleware
	// İndirme istatistikleri endpoint'i (JWT ile korunmuş)
	r.GET("/api//admin/download-stats", auth.AuthMiddleware(), adminpanel.AdminDownloadStatsHandler(db))
	r.GET("/api/admin/monthly-download-stats", auth.AuthMiddleware(), adminpanel.AdminMonthlyDownloadStatsHandler(db))
	r.GET("/api/admin/ip-logs", auth.AuthMiddleware(), adminpanel.AdminIPLogsHandler(db))

	if os.Getenv("SWAGGER_ENABLED") == "true" {
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	} else {
		fmt.Println("Swagger is disabled")
	}

	// VersionService ve handler'ı oluştur (dependency injection)
	versionService := versionControl.NewService(logger)
	versionHandler := versionControl.NewHandler(versionService, logger)

	// Download route'u
	r.GET("/api/download", downloadHandler.Download)

	// Version route'u
	r.GET("/api/version", versionHandler.GetVersion)

	// Sunucuyu başlat
	port := fmt.Sprintf(":%s", cfg.Port)
	logger.Infof("Sunucu %s portunda başlatılıyor...", cfg.Port)
	r.Run(port)
}

// DBMiddleware veritabanı bağlantısını bağlama ekleyen middleware
func DBMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Veritabanı bağlantısını bağlama ekle
		c.Set("db", db)
		c.Next()
	}
}
