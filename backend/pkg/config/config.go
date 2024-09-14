package config

import (
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

// Config yapısı, uygulamanın genel yapılandırmalarını tutar.
type Config struct {
	Port        string
	LogOutput   string
	LogFilePath string
	DBHost      string
	DBUser      string
	DBPassword  string
	DBName      string
	DBPort      string
}

// LoadConfig, yapılandırmaları yükler ve bir Config nesnesi döner.
func LoadConfig() *Config {
	// .env dosyasını yükle
	err := godotenv.Load()
	if err != nil {
		log.Printf(".env dosyası bulunamadı, varsayılan değerler kullanılacak: %v", err)
	}

	// Port bilgisi .env dosyasından okunur, bulunamazsa 8080 olarak ayarlanır.
	return &Config{
		Port:        getEnv("PORT", "8080"),
		LogOutput:   getEnv("LOG_OUTPUT", "file"), //default olarak file seçilmiştir
		LogFilePath: getEnv("LOG_FILE_PATH", getDefaultLogPath()),
		DBHost:      getEnv("DB_HOST", "localhost"),
		DBUser:      getEnv("DB_USER", ""),
		DBPassword:  getEnv("DB_PASSWORD", ""),
		DBName:      getEnv("DB_NAME", "beehub"),
		DBPort:      getEnv("DB_PORT", "5432"),
	}
}

// getEnv, çevre değişkenini okur; yoksa varsayılan değeri döner.
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

// SetupLogger, uygulama için basit bir logger ayarlar.
func SetupLogger(cfg *Config) *logrus.Logger {
	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	if cfg.LogOutput == "file" {
		file, err := os.OpenFile(cfg.LogFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("Log dosyası açılamadı: %v", err)
		}
		log.SetOutput(file)
	} else {
		log.SetOutput(os.Stdout)
	}

	log.SetLevel(logrus.InfoLevel) // Varsayılan olarak DEBUG seviyesi
	log.SetOutput(os.Stdout)
	return log
}

// getDefaultLogPath işletim sistemine göre varsayılan log dosyası yolunu döner
func getDefaultLogPath() string {
	osType := runtime.GOOS
	var logPath string

	switch osType {
	case "linux":
		logPath = "/var/log/beehub.log"
	case "windows":
		programData := os.Getenv("PROGRAMDATA")
		if programData == "" {
			programData = "C:\\ProgramData" // Varsayılan yol
		}
		logPath = filepath.Join(programData, "BeeHub", "beehub.log")
	case "darwin": // MacOS için varsayılan log yeri
		logPath = "/usr/local/var/log/beehub.log"
	default:
		logPath = "./beehub.log" // Bilinmeyen OS durumunda
	}

	return logPath
}
