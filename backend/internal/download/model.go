// Package download serves the desktop installers and records download events.
package download

import "time"

// Download is a single counted installer download.
type Download struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	OS        string    `gorm:"index;index:idx_downloads_dedup,priority:2" json:"os"`
	IPAddress string    `gorm:"index;index:idx_downloads_dedup,priority:1" json:"ipAddress"`
	Country   string    `gorm:"index" json:"country"`
	CreatedAt time.Time `gorm:"index;index:idx_downloads_dedup,priority:3" json:"createdAt"`
}
