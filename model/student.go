package model

import (
	"github.com/adieos/student-db/utils"
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

func GeneratePass(pass string) (string, error) {
	var err error
	pass, err = utils.GeneratePassword(pass)
	if err != nil {
		return "", err
	}

	return pass, nil
}
