package models

import "github.com/google/uuid"

type UserToAuth struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	ID       uuid.UUID `json:"id"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	Name     string    `json:"name"`
}