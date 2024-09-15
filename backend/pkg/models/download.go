package models

import "time"

// Download struct, indirme bilgilerini tutar
type Download struct {
	ID        uint   `gorm:"primaryKey"`
	OS        string `gorm:"index"` // windows, mac, linux gibi
	IPAddress string `gorm:"index"`
	Country   string `gorm:"index"`
	CreatedAt time.Time
}

// MonthlyDownloadStats aylık indirme istatistiklerini tutar
type MonthlyDownloadStats struct {
	Month time.Time `json:"month"`
	Count int       `json:"count"`
}
