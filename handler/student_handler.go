package handler

import (
	"source/my-clean-arch/entity"
	"source/my-clean-arch/repository/student"

	"github.com/gin-gonic/gin"
)

type StudentHandler struct {
	service student.Service
}

func NewStudentHandler(repo student.StudentRepository) *StudentHandler {
	return &StudentHandler{
		service: *student.NewService(repo),
	}
}

func (h *StudentHandler) GetStudents(c *gin.Context) {
	ids := []uint64{5921, 8294, 7283}
	data, err := h.service.GetAllById(ids)
	if err != nil {
		c.JSON(500, gin.H{
			"code":    500,
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"code": 200,
		"data": data,
	})
}

func (h *StudentHandler) CreateStudent(c *gin.Context) {
	student := entity.Student{}
	if err := c.ShouldBind(&student); err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": err.Error(),
		})
		return
	}

	err := h.service.Create(&student)
	if err != nil {
		c.JSON(500, gin.H{
			"code":    500,
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"code": 200,
		"data": student,
	})
}
