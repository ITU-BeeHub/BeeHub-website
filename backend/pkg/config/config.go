package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

// Config yapısı, uygulamanın genel yapılandırmalarını tutar.
type Config struct {
	Port string
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
		Port: getEnv("PORT", "8080"),
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
func SetupLogger() *logrus.Logger {
	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	log.SetLevel(logrus.InfoLevel) // Varsayılan olarak DEBUG seviyesi
	log.SetOutput(os.Stdout)
	return log
}
