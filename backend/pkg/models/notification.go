// models/notification.go
package models

import (
	"time"
)

type Notification struct {
	Title     string    `json:"title"`
	Message   string    `json:"message"`
	Type      string    `json:"type"` // e.g. "info", "warning", "announcement"
	Active    bool      `json:"active"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}
