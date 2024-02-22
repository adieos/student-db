package repository

import (
	"errors"
	"net/http"

	"github.com/adieos/student-db/dto"
	"github.com/adieos/student-db/model"
	"gorm.io/gorm"
)

type StudentRepository interface {
	CreateStudent(studentDTO dto.Student) (dto.StudentResponse, error)
	GetAllStudents() (dto.StudentResponse, error)
	GetStudentById(studentId string) (dto.StudentResponse, error)
	UpdateStudent(studentDTO dto.Student) (dto.StudentResponse, error)
	DeleteStudent(studentId string) (dto.StudentResponse, error)
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

func (s *studentRepository) GetAllStudents() (dto.StudentResponse, error) {
	var students []model.Student
	if err := s.db.Find(&students).Error; err != nil {
		return dto.StudentResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "Unable to retrieve all students",
		}, err
	}

	return dto.StudentResponse{
		StatusCode: http.StatusOK,
		Message:    "Retrieved all students successfully",
		Data:       students,
	}, nil
}

func (s *studentRepository) GetStudentById(studentId string) (dto.StudentResponse, error) {
	var student model.Student
	if err := s.db.Where("id = ?", studentId).First(&student).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return dto.StudentResponse{
				StatusCode: http.StatusNotFound,
				Message:    "Student not found",
			}, err
		}
		return dto.StudentResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "Unable to retrieve student",
		}, err
	}

	return dto.StudentResponse{
		StatusCode: http.StatusOK,
		Message:    "Retrieved student successfully",
		Data:       student,
	}, nil
}

func (s *studentRepository) UpdateStudent(studentDTO dto.Student) (dto.StudentResponse, error) {
	// get the old student data (for id and major)
	var student model.Student
	if err := s.db.Model(&student).Where("id = ?", studentDTO.ID).First(&student).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return dto.StudentResponse{
				StatusCode: http.StatusNotFound,
				Message:    "Student not found",
			}, err
		}

		return dto.StudentResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "Unable to retrieve student",
		}, err
	}

	// get major
	var major model.Major
	if err := s.db.Model(&major).Where("name = ?", studentDTO.Major).First(&major).Error; err != nil {
		return dto.StudentResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "Unable to retrieve major",
		}, err
	}

	// get password
	password, err := model.GeneratePass(studentDTO.Password)
	if err != nil {
		return dto.StudentResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Unable to hash password",
		}, err
	}

	new := model.Student{
		ID:       student.ID,
		Name:     studentDTO.Name,
		Email:    studentDTO.Email,
		Password: password,
		MajorID:  major.ID.String(),
	}

	// update data
	if err := s.db.Updates(&new).Error; err != nil {
		return dto.StudentResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "Unable to update student",
		}, err
	}

	return dto.StudentResponse{
		StatusCode: http.StatusOK,
		Message:    "Student updated successfully",
		Data:       new,
	}, nil
}

func (s *studentRepository) DeleteStudent(studentId string) (dto.StudentResponse, error) {
	if err := s.db.Delete(&model.Student{}, studentId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return dto.StudentResponse{
				StatusCode: http.StatusNotFound,
				Message:    "Student not found",
			}, err
		}

		return dto.StudentResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "Unable to delete student",
		}, err
	}

	return dto.StudentResponse{
		StatusCode: http.StatusNoContent,
		Message:    "Student deleted successfully",
	}, nil
}
