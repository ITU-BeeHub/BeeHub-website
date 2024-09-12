package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/ITU-BeeHub/BeeHub-website/backend/internal/versionControl"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title BeeHub Web Site API
// @version 1.0
// @description Bu, BeeHub web sitesi için API belgeleridir.

// @host localhost:8080
// @BasePath /
// @schemes http
func main() {
	// .env dosyasını yükle
	LoadEnvVariables()

	r := gin.Default()

	fmt.Println(os.Getenv("SWAGGER_ENABLED"))
	if os.Getenv("SWAGGER_ENABLED") == "true" {
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	} else {
		fmt.Println("Swagger is disabled")
	}

	// Version control route
	r.GET("/version", versionControl.VersionHandler)

	// Other routes can be added here

	r.Run(":8080") // start service at port 8080
}

func LoadEnvVariables() {
	file, err := os.Open(".env")
	if err != nil {
		log.Fatalf("Error opening .env file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 || strings.HasPrefix(line, "#") {
			// Boş veya yorum satırlarını atla
			continue
		}

		pair := strings.SplitN(line, "=", 2)
		if len(pair) != 2 {
			log.Printf("Ignoring malformed line: %s", line)
			continue
		}

		key := strings.TrimSpace(pair[0])
		value := strings.TrimSpace(pair[1])

		// Çevresel değişkeni ayarla
		os.Setenv(key, value)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading .env file: %v", err)
	}
}
