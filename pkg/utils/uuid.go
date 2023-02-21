package utils

import uuid "github.com/satori/go.uuid"

func NewUUID() uuid.UUID {
	return uuid.NewV4()
}

func NewUUIDStr() string {
	return NewUUID().String()
}
