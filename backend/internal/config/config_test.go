package config

import (
	"strings"
	"testing"
	"time"
)

func setRequired(t *testing.T) {
	t.Setenv("DB_USER", "beehub")
	t.Setenv("JWT_SECRET", strings.Repeat("s", minJWTSecretLength))
	t.Setenv("APP_VERSION", "1.2.1")
}

func TestLoadDefaults(t *testing.T) {
	setRequired(t)

	cfg, err := Load()
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if cfg.HTTPAddr != "127.0.0.1:9000" {
		t.Errorf("HTTPAddr = %q", cfg.HTTPAddr)
	}
	if cfg.Download.DedupWindow != 10*time.Minute {
		t.Errorf("DedupWindow = %v", cfg.Download.DedupWindow)
	}
	if got := strings.Join(cfg.TrustedProxies, ","); got != "127.0.0.1,::1" {
		t.Errorf("TrustedProxies = %q", got)
	}
}

func TestLoadRejectsInvalidConfig(t *testing.T) {
	tests := map[string][2]string{
		"short secret":   {"JWT_SECRET", "123456"},
		"no version":     {"APP_VERSION", ""},
		"no db user":     {"DB_USER", ""},
		"bad duration":   {"DOWNLOAD_DEDUP_WINDOW", "ten minutes"},
		"bad bool":       {"SWAGGER_ENABLED", "maybe"},
		"bad time zone":  {"STATS_TIMEZONE", "Mars/Olympus"},
		"zero token ttl": {"JWT_TTL", "0s"},
	}
	for name, env := range tests {
		t.Run(name, func(t *testing.T) {
			setRequired(t)
			t.Setenv(env[0], env[1])
			if _, err := Load(); err == nil {
				t.Error("Load accepted invalid configuration")
			}
		})
	}
}
