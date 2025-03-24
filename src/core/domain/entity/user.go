package entity

import "time"

type User struct {
	ID            string
	IntegrationID string
	Name          string
	Email         string
	Password      string
	PasswordHash  string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
