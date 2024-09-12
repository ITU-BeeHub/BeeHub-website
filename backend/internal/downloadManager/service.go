package downloadManager

import (
	"errors"

	"github.com/sirupsen/logrus"
)

type Service struct {
	logger *logrus.Logger
}

func NewService(logger *logrus.Logger) *Service {
	return &Service{logger: logger}
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
