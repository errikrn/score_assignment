package handler

import (
	"assignment-project-new/database"
	"assignment-project-new/helpers"
	"assignment-project-new/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	appJSON = "application/json"
)

func CreateStudent(ctx *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(ctx)
	student := models.Student{}

	if contentType == appJSON {
		if err := ctx.ShouldBindJSON(&student); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error":   "Invalid input",
				"message": err.Error(),
			})
			return
		}
	} else {
		if err := ctx.ShouldBind(&student); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error":   "Invalid input",
				"message": err.Error(),
			})
			return
		}
	}

	if err := db.Debug().Create(&student).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Database error",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    student,
	})
}

func GetAllStudents(ctx *gin.Context) {
	db := database.GetDB()

	var students []models.Student
	if err := db.Debug().Preload("Scores").Find(&students).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Database error",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    students,
	})
}

func UpdateStudent(ctx *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(ctx)
	student := models.Student{}
	studentID := ctx.Param("id")

	if contentType == appJSON {
		if err := ctx.ShouldBindJSON(&student); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error":   "Invalid input",
				"message": err.Error(),
			})
			return
		}
	} else {
		if err := ctx.ShouldBind(&student); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error":   "Invalid input",
				"message": err.Error(),
			})
			return
		}
	}

	var existingStudent models.Student
	if err := db.Debug().Where("id = ?", studentID).Preload("Scores").First(&existingStudent).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Student not found",
		})
		return
	}

	existingStudent.Name = student.Name
	existingStudent.Age = student.Age

	if err := db.Debug().Save(&existingStudent).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Database error",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    existingStudent,
	})
}

func DeleteStudent(ctx *gin.Context) {
	db := database.GetDB()
	studentID := ctx.Param("id")

	var existingStudent models.Student
	if err := db.Debug().Where("id = ?", studentID).First(&existingStudent).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Student not found",
		})
		return
	}

	if err := db.Debug().Delete(&existingStudent).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Database error",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Student deleted successfully",
	})
}
