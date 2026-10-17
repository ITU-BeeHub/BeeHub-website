package download

import (
	"context"
	"log/slog"
	"strings"
	"time"
)

// Store persists download events.
type Store interface {
	RecordUnique(ctx context.Context, d *Download, window time.Duration) (bool, error)
}

// Service resolves installers and records downloads.
type Service struct {
	catalog     Catalog
	store       Store
	dedupWindow time.Duration
	now         func() time.Time
	log         *slog.Logger
}

func NewService(catalog Catalog, store Store, dedupWindow time.Duration, log *slog.Logger) *Service {
	return &Service{
		catalog:     catalog,
		store:       store,
		dedupWindow: dedupWindow,
		now:         time.Now,
		log:         log,
	}
}

// Installer returns the installer for the given platform.
func (s *Service) Installer(platform string) (Installer, error) {
	return s.catalog.Lookup(platform)
}

// Record counts a download, ignoring repeats within the dedup window.
func (s *Service) Record(ctx context.Context, platform, ip, country string) error {
	d := &Download{
		OS:        platform,
		IPAddress: ip,
		Country:   normalizeCountry(country),
		CreatedAt: s.now(),
	}

	stored, err := s.store.RecordUnique(ctx, d, s.dedupWindow)
	if err != nil {
		return err
	}
	s.log.Info("download requested", "os", platform, "ip", ip, "country", d.Country, "counted", stored)
	return nil
}

// normalizeCountry returns an ISO 3166-1 alpha-2 code, or "" when the value
// is missing or the CDN could not determine the country ("XX").
func normalizeCountry(code string) string {
	code = strings.ToUpper(strings.TrimSpace(code))
	if len(code) != 2 || code == "XX" {
		return ""
	}
	return code
}
