package downloadManager

import (
	"errors"

	"github.com/ITU-BeeHub/BeeHub-website/backend/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Service struct {
	logger *logrus.Logger
}

func NewService(logger *logrus.Logger) *Service {
	return &Service{logger: logger}
}

// SaveDownload, indirme bilgilerini veritabanına kaydeder
func (s *Service) SaveDownload(c *gin.Context, os, ip string) error {
	// Veritabanı bağlantısını bağlamdan al
	db, exists := c.Get("db")
	if !exists {
		s.logger.Error("Database connection not found in context")
		return errors.New("database connection not found")
	}

	download := models.Download{
		OS:        os,
		IPAddress: ip,
	}

	// Veritabanına kaydet
	if err := db.(*gorm.DB).Create(&download).Error; err != nil {
		s.logger.Error("Error while saving download info: ", err)
		return err
	}
	s.logger.Infof("Download info saved: OS=%s, IP=%s", os, ip)
	return nil
}

// GetDownloadStats, indirme istatistiklerini işletim sistemine göre gruplar ve döner
func GetDownloadStats(db *gorm.DB) (map[string]int, error) {
	var countByOS []struct {
		OS    string
		Count int
	}

	// OS'ye göre gruplama ve indirme sayısını alıyoruz
	err := db.Model(&models.Download{}).Select("os, count(*) as count").Group("os").Scan(&countByOS).Error
	if err != nil {
		return nil, err
	}

	// İstatistikleri bir map içinde topluyoruz
	stats := make(map[string]int)
	for _, row := range countByOS {
		stats[row.OS] = row.Count
	}

	return stats, nil
}

// GetInstallerFilePath iş mantığını içerir ve doğru dosya yolunu döner.
func (s *Service) GetInstallerFilePath(os string) (string, error) {
	var filePath string

	s.logger.Infof("Getting installer file path for OS: %s", os)
	switch os {
	case "windows":
		filePath = "./installers/installer.exe"
	case "mac":
		filePath = "./installers/installer.dmg"
	case "linux":
		filePath = "./installers/installer.deb"
	default:
		s.logger.Warn("Unsupported OS provided: ", os)
		return "", errors.New("unsupported OS")
	}
	s.logger.Infof("File path found for OS %s: %s", os, filePath)
	return filePath, nil
}
