// Package config loads and validates the application configuration from
// environment variables.
package config

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const minJWTSecretLength = 32

// Config holds every setting the application needs at runtime.
type Config struct {
	HTTPAddr       string
	PublicURL      string
	TrustedProxies []string
	SwaggerEnabled bool

	Database Database
	Auth     Auth
	Download Download

	// StatsLocation is the time zone used to bucket statistics into days and months.
	StatsLocation *time.Location
}

// Database holds the PostgreSQL connection settings.
type Database struct {
	Host           string
	Port           string
	User           string
	Password       string
	Name           string
	SSLMode        string
	ConnectTimeout time.Duration
}

// Auth holds the admin authentication settings.
type Auth struct {
	JWTSecret []byte
	TokenTTL  time.Duration
}

// Download holds the installer distribution settings.
type Download struct {
	AppVersion    string
	InstallersDir string
	// DedupWindow is the period during which repeated downloads from the
	// same IP address for the same platform are counted only once.
	DedupWindow time.Duration
	// CountryHeader is the request header carrying the client's ISO country
	// code, set by the CDN in front of the server.
	CountryHeader string
}

// Load reads the configuration from the environment and validates it.
func Load() (*Config, error) {
	var errs []error
	collect := func(err error) {
		if err != nil {
			errs = append(errs, err)
		}
	}

	cfg := &Config{
		HTTPAddr:       getEnv("HTTP_ADDR", "127.0.0.1:9000"),
		PublicURL:      getEnv("PUBLIC_URL", "https://beehubapp.com"),
		TrustedProxies: splitList(getEnv("TRUSTED_PROXIES", "127.0.0.1,::1")),
		Auth: Auth{
			JWTSecret: []byte(os.Getenv("JWT_SECRET")),
		},
		Download: Download{
			AppVersion:    os.Getenv("APP_VERSION"),
			InstallersDir: getEnv("INSTALLERS_DIR", "/opt/beehub/installers"),
			CountryHeader: getEnv("COUNTRY_HEADER", "CF-IPCountry"),
		},
	}

	var err error
	cfg.Database, err = LoadDatabase()
	collect(err)
	cfg.SwaggerEnabled, err = getBool("SWAGGER_ENABLED", false)
	collect(err)
	cfg.Auth.TokenTTL, err = getDuration("JWT_TTL", 6*time.Hour)
	collect(err)
	cfg.Download.DedupWindow, err = getDuration("DOWNLOAD_DEDUP_WINDOW", 10*time.Minute)
	collect(err)
	cfg.StatsLocation, err = time.LoadLocation(getEnv("STATS_TIMEZONE", "Europe/Istanbul"))
	collect(err)

	collect(cfg.validate())

	if err := errors.Join(errs...); err != nil {
		return nil, fmt.Errorf("invalid configuration: %w", err)
	}
	return cfg, nil
}

// LoadDatabase reads only the database settings, for tools that need
// nothing else.
func LoadDatabase() (Database, error) {
	db := Database{
		Host:     getEnv("DB_HOST", "localhost"),
		Port:     getEnv("DB_PORT", "5432"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Name:     getEnv("DB_NAME", "beehub"),
		SSLMode:  getEnv("DB_SSLMODE", "disable"),
	}

	var errs []error
	if db.User == "" {
		errs = append(errs, errors.New("DB_USER is required"))
	}
	var err error
	if db.ConnectTimeout, err = getDuration("DB_CONNECT_TIMEOUT", time.Minute); err != nil {
		errs = append(errs, err)
	}
	return db, errors.Join(errs...)
}

func (c *Config) validate() error {
	var errs []error
	if len(c.Auth.JWTSecret) < minJWTSecretLength {
		errs = append(errs, fmt.Errorf("JWT_SECRET must be at least %d characters", minJWTSecretLength))
	}
	if c.Download.AppVersion == "" {
		errs = append(errs, errors.New("APP_VERSION is required"))
	}
	if c.Auth.TokenTTL <= 0 {
		errs = append(errs, errors.New("JWT_TTL must be positive"))
	}
	if c.Download.DedupWindow < 0 {
		errs = append(errs, errors.New("DOWNLOAD_DEDUP_WINDOW must not be negative"))
	}
	return errors.Join(errs...)
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok && value != "" {
		return value
	}
	return fallback
}

func getBool(key string, fallback bool) (bool, error) {
	raw, ok := os.LookupEnv(key)
	if !ok || raw == "" {
		return fallback, nil
	}
	value, err := strconv.ParseBool(raw)
	if err != nil {
		return false, fmt.Errorf("%s: %w", key, err)
	}
	return value, nil
}

func getDuration(key string, fallback time.Duration) (time.Duration, error) {
	raw, ok := os.LookupEnv(key)
	if !ok || raw == "" {
		return fallback, nil
	}
	value, err := time.ParseDuration(raw)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", key, err)
	}
	return value, nil
}

func splitList(raw string) []string {
	var items []string
	for _, item := range strings.Split(raw, ",") {
		if item = strings.TrimSpace(item); item != "" {
			items = append(items, item)
		}
	}
	return items
}
