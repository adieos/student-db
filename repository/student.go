package repository

import (
	"net/http"

	"github.com/adieos/student-db/dto"
	"github.com/adieos/student-db/model"
	"gorm.io/gorm"
)

type StudentRepository interface {
	CreateStudent(studentDTO dto.Student)
	GetAllStudents()
	GetStudentById(studentId string)
	UpdateStudent(studentDTO dto.Student)
	DeleteStudent(studentId string)
}

type studentRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) StudentRepository {
	return &studentRepository{
		db: db,
	}
}

func (s *studentRepository) CreateStudent(studentDTO dto.Student) (dto.StudentResponse, error) {

	// generate password
	password, err := model.GeneratePass(studentDTO.Password)
	if err != nil {
		return dto.StudentResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Unable to hash password",
		}, err
	}

	// get major
	var major model.Major
	if err = s.db.Where("name = ?", studentDTO.Major).First(&major).Error; err != nil {
		return dto.StudentResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Unable to get major",
		}, err
	}

	// generate the student and insert it
	student := model.Student{
		Name:     studentDTO.Name,
		Email:    studentDTO.Email,
		Password: password,
		MajorID:  major.ID.String(),
	}

	if err = s.db.Create(&student).Error; err != nil {
		return dto.StudentResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Unable to register student",
		}, err
	}

	return dto.StudentResponse{
		StatusCode: http.StatusCreated,
		Message:    "Student registered successfully",
		Data:       student,
	}, nil
}
