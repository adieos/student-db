package model

import (
	"github.com/google/uuid"
)

type Student struct {
	ID       uuid.UUID `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"-"`
	MajorID  string    `json:"major_id"`
	Major    *Major    `json:"major,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
