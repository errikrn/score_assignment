package handler

import (
	"assignment-project-new/database"
	"assignment-project-new/helpers"
	"assignment-project-new/models"
	"net/http"
	"strconv"

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
	studentID := ctx.Param("id")
	updatedStudent := models.Student{}

	if err := ctx.ShouldBindJSON(&updatedStudent); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "message": err.Error()})
		return
	}

	id, err := strconv.Atoi(studentID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student ID"})
		return
	}

	var existingStudent models.Student
	if err := db.Debug().Preload("Scores").Where("id = ?", id).First(&existingStudent).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}

	if err := db.Debug().Where("student_id = ?", id).Delete(models.Scores{}).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Database error", "message": err.Error()})
		return
	}

	existingStudent.Scores = updatedStudent.Scores

	if err := db.Debug().Save(&existingStudent).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Database error", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Student updated", "data": existingStudent})
}

func DeleteStudent(ctx *gin.Context) {
	db := database.GetDB()
	studentID := ctx.Param("id")

	var student models.Student
	if err := db.Debug().Preload("Scores").Where("id = ?", studentID).First(&student).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}

	if err := db.Debug().Where("student_id = ?", studentID).Delete(&models.Scores{}).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Database error", "message": err.Error()})
		return
	}

	if err := db.Debug().Delete(&student).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Database error", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Student and scores deleted successfully"})
}
