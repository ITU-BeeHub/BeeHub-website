// Package auth authenticates admin users and protects the admin API.
package auth

import "gorm.io/gorm"

// Admin is a user allowed to access the admin panel.
type Admin struct {
	gorm.Model
	Username     string `gorm:"unique;not null"`
	PasswordHash string `gorm:"column:password;not null"`
}
