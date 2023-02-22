package utils

import uuid "github.com/satori/go.uuid"

// NewUUID creates uuid.UUID
func NewUUID() uuid.UUID {
	return uuid.NewV4()
}

// NewUUIDStr creates uuid.UUID as its string representation
func NewUUIDStr() string {
	return NewUUID().String()
}
