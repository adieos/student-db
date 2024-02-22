package handler

// TODO: figure out how to use dependency injection here

import (
	"github.com/gin-gonic/gin"
)

type StudentHandler interface {
	Register(ctx *gin.Context)
	GetAll(ctx *gin.Context)
	GetById(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

func Register(ctx *gin)
