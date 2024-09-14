package models

import "time"

// Download struct, indirme bilgilerini tutar
type Download struct {
	ID        uint   `gorm:"primaryKey"`
	OS        string `gorm:"index"` // windows, mac, linux gibi
	IPAddress string `gorm:"index"`
	CreatedAt time.Time
}
