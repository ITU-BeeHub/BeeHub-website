package downloadManager

import (
	"errors"
	"fmt"
	"net"
	"os"
	"time"

	"github.com/ITU-BeeHub/BeeHub-website/backend/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/oschwald/geoip2-golang"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Service struct {
	logger *logrus.Logger
}

func NewService(logger *logrus.Logger) *Service {
	return &Service{logger: logger}
}

func (s *Service) SaveDownload(c *gin.Context, os, ip string) error {
	// Veritabanı bağlantısını bağlamdan al
	db, exists := c.Get("db")
	if !exists {
		s.logger.Error("Database connection not found in context")
		return errors.New("database connection not found")
	}

	// IP'den ülke bilgisini alın (Eğer bir kütüphane kullanıyorsanız)
	country := GetCountryFromIP(ip)

	download := models.Download{
		OS:        os,
		IPAddress: ip,
		Country:   country,
		CreatedAt: time.Now(),
	}

	// Veritabanına kaydet
	if err := db.(*gorm.DB).Create(&download).Error; err != nil {
		s.logger.Error("Error while saving download info: ", err)
		return err
	}
	s.logger.Infof("Download info saved: OS=%s, IP=%s, Country=%s", os, ip, country)
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

// GetMonthlyDownloadStats, aylık indirme istatistiklerini döner
func GetMonthlyDownloadStats(db *gorm.DB) ([]models.MonthlyDownloadStats, error) {
	var monthlyStats []models.MonthlyDownloadStats

	err := db.Model(&models.Download{}).
		Select("DATE_TRUNC('month', created_at) as month, COUNT(*) as count").
		Group("month").
		Order("month").
		Scan(&monthlyStats).Error

	if err != nil {
		return nil, err
	}

	return monthlyStats, nil
}

// GetIPLogs, son 100 indirme kaydını döner
func GetIPLogs(db *gorm.DB) ([]models.Download, error) {
	var logs []models.Download
	err := db.Order("created_at DESC").Limit(100).Find(&logs).Error
	if err != nil {
		return nil, err
	}

	return logs, nil
}

// GetInstallerFilePath iş mantığını içerir ve doğru dosya yolunu döner.
func (s *Service) GetInstallerFilePath(osystem string) (string, error) {
	var filePath string
	version := os.Getenv("APP_VERSION")

	s.logger.Infof("Getting installer file path for OS: %s", osystem)
	switch osystem {
	case "windows":
		filePath = fmt.Sprintf("./installers/installer%s.exe", version)
	case "mac":
		filePath = fmt.Sprintf("./installers/installer%s.dmg", version)
	default:
		s.logger.Warn("Unsupported OS provided: ", osystem)
		return "", errors.New("unsupported OS")
	}
	s.logger.Infof("File path found for OS %s: %s", osystem, filePath)
	return filePath, nil
}

func GetCountryFromIP(ipStr string) string {
	db, err := geoip2.Open("GeoLite2-Country.mmdb")
	if err != nil {
		// Hata durumunda boş string dönebilir veya varsayılan bir değer kullanabilirsiniz
		return ""
	}
	defer db.Close()

	ip := net.ParseIP(ipStr)
	record, err := db.Country(ip)
	if err != nil {
		return ""
	}

	return record.Country.IsoCode
}
