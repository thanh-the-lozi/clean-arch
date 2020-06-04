package main

import (
	"source/my-clean-arch/handler"
	"source/my-clean-arch/repository/student"

	"github.com/gin-gonic/gin"
)

func NewStudentHandler(e *gin.Engine) *gin.Engine {
	repo := student.NewStudentRepository(db)
	h := handler.NewStudentHandler(*repo)

	e.GET("/students", h.GetStudents)
	e.POST("/students", h.CreateStudent)
	// e.PUT("/students", h.Update)
	// e.DELETE("/students", h.Delete)

	return e
}
