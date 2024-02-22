package dto

type Student struct {
	Name     string `json:"name" form:"name" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
	Major    string `json:"major" form:"major" binding:"required"`
}

type StudentResponse struct {
	StatusCode int    `json:"statuscode"`
	Message    string `json:"message"`
	Data       any    `json:"data,omitempty"`
}
