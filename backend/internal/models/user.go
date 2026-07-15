package models

import "time"

type UserRole string

const (
	RoleAdmin UserRole = "admin"
	RoleUser  UserRole = "user"
)

type User struct {
	ID              string     `json:"id"`
	Name            string     `json:"name"`
	Email           string     `json:"email"`
	PasswordHash    string     `json:"-"`
	Role            UserRole   `json:"role"`
	CreatedAt       time.Time  `json:"createdAt"`
	UpdatedAt       time.Time  `json:"updatedAt"`
	EmailVerifiedAt *time.Time `json:"emailVerifiedAt,omitempty"`
}