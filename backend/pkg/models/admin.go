package models

import "gorm.io/gorm"

// Admin yapısı, veritabanında admin kullanıcıları temsil eder
type Admin struct {
	gorm.Model
	Username string `gorm:"unique"`
	Password string // Hashlenmiş şifre
}
