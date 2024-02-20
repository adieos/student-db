package model

import (
	"github.com/google/uuid"
)

type Major struct {
	ID       uuid.UUID `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name     string    `json:"name"`
	Students []Student `json:"students"`
}
