package client

import "time"

type UserAlertBinding struct {
	UserID        string
	AlertConfigID string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
