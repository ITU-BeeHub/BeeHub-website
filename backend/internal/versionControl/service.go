package versionControl

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

type Service struct {
	logger *logrus.Logger
}

func NewService(logger *logrus.Logger) *Service {
	return &Service{logger: logger}
}

// GetVersionInfo, versiyon bilgisini dinamik bir kaynaktan (örneğin .env dosyası) çeker
func (s *Service) GetVersionInfo() (map[string]string, error) {
	version := os.Getenv("APP_VERSION")
	if version == "" {
		version = "1.0.0" // Varsayılan versiyon
		s.logger.Warn("APP_VERSION environment variable is not set, using default version")
	}

	downloadURL := fmt.Sprintf("https://yourdomain.com/download/installer_%s.exe", version)

	versionInfo := map[string]string{
		"version":      version,
		"download_url": downloadURL,
	}

	return versionInfo, nil
}
