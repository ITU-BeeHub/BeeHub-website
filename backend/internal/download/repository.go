package download

import (
	"context"
	"time"

	"gorm.io/gorm"
)

// Repository persists download events in PostgreSQL.
type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

// RecordUnique stores d unless a download from the same IP address for the
// same platform was already stored within window before d.CreatedAt.
// It reports whether d was stored.
//
// Concurrent requests for the same IP and platform are serialized with a
// transaction-scoped advisory lock, so the existence check and the insert
// cannot interleave and a burst of parallel requests is counted once.
func (r *Repository) RecordUnique(ctx context.Context, d *Download, window time.Duration) (bool, error) {
	stored := false
	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		lockKey := d.IPAddress + "|" + d.OS
		if err := tx.Exec("SELECT pg_advisory_xact_lock(hashtextextended(?, 0))", lockKey).Error; err != nil {
			return err
		}

		var duplicate bool
		err := tx.Raw(
			"SELECT EXISTS (SELECT 1 FROM downloads WHERE ip_address = ? AND os = ? AND created_at > ?)",
			d.IPAddress, d.OS, d.CreatedAt.Add(-window),
		).Scan(&duplicate).Error
		if err != nil || duplicate {
			return err
		}

		if err := tx.Create(d).Error; err != nil {
			return err
		}
		stored = true
		return nil
	})
	return stored, err
}
