// Package admin provides the read-only reporting API behind the admin panel.
package admin

import (
	"context"
	"time"

	"github.com/ITU-BeeHub/BeeHub-website/backend/internal/download"
	"gorm.io/gorm"
)

// Totals are download counts over fixed periods.
type Totals struct {
	Total      int64 `json:"total"`
	Today      int64 `json:"today"`
	Last7Days  int64 `json:"last7Days"`
	Last30Days int64 `json:"last30Days"`
	UniqueIPs  int64 `json:"uniqueIps"`
}

// PlatformCount is the number of downloads for one platform.
type PlatformCount struct {
	OS    string `json:"os"`
	Count int64  `json:"count"`
}

// bucketCount is the number of downloads for one platform in one period.
type bucketCount struct {
	Bucket string
	OS     string
	Count  int64
}

// DownloadFilter narrows a download listing.
type DownloadFilter struct {
	Search string
	OS     string
	Limit  int
	Offset int
}

// Repository runs the reporting queries.
type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

// Totals counts downloads since each of the given period starts.
func (r *Repository) Totals(ctx context.Context, today, last7Days, last30Days time.Time) (Totals, error) {
	var totals Totals
	err := r.db.WithContext(ctx).Raw(`
		SELECT COUNT(*)                                     AS total,
		       COUNT(*) FILTER (WHERE created_at >= ?)      AS today,
		       COUNT(*) FILTER (WHERE created_at >= ?)      AS last7_days,
		       COUNT(*) FILTER (WHERE created_at >= ?)      AS last30_days,
		       COUNT(DISTINCT ip_address)                   AS unique_ips
		FROM downloads`, today, last7Days, last30Days).
		Scan(&totals).Error
	return totals, err
}

// CountByPlatform counts all downloads per platform, most downloaded first.
func (r *Repository) CountByPlatform(ctx context.Context) ([]PlatformCount, error) {
	counts := []PlatformCount{}
	err := r.db.WithContext(ctx).Model(&download.Download{}).
		Select("os, COUNT(*) AS count").
		Group("os").
		Order("count DESC").
		Scan(&counts).Error
	return counts, err
}

// countByBucket counts downloads since `since` per platform and period. The
// period is the creation time formatted with the PostgreSQL `layout` in the
// given time zone, e.g. "YYYY-MM-DD" for days.
func (r *Repository) countByBucket(ctx context.Context, layout string, loc *time.Location, since time.Time) ([]bucketCount, error) {
	var counts []bucketCount
	err := r.db.WithContext(ctx).Raw(`
		SELECT to_char(created_at AT TIME ZONE ?, ?) AS bucket, os, COUNT(*) AS count
		FROM downloads
		WHERE created_at >= ?
		GROUP BY bucket, os`, loc.String(), layout, since).
		Scan(&counts).Error
	return counts, err
}

// ListDownloads returns a page of downloads, newest first, and the number of
// downloads matching the filter.
func (r *Repository) ListDownloads(ctx context.Context, f DownloadFilter) ([]download.Download, int64, error) {
	query := r.db.WithContext(ctx).Model(&download.Download{})
	if f.OS != "" {
		query = query.Where("os = ?", f.OS)
	}
	if f.Search != "" {
		pattern := "%" + f.Search + "%"
		query = query.Where("ip_address ILIKE ? OR country ILIKE ?", pattern, pattern)
	}

	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	downloads := []download.Download{}
	err := query.Order("created_at DESC").Limit(f.Limit).Offset(f.Offset).Find(&downloads).Error
	return downloads, total, err
}
