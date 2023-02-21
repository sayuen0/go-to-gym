package db

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID             uuid.UUID  `genorm:"id"`
	Name           string     `genorm:"name"`
	Email          string     `genorm:"email"`
	HashedPassword string     `genorm:"hashed_password"`
	CreatedAt      time.Time  `genorm:"created_at"`
	UpdatedAt      time.Time  `genorm:"updated_at"`
	DeleteAt       *time.Time `genorm:"delete_at"`
}

func (*User) TableName() string {
	return "users"
}
