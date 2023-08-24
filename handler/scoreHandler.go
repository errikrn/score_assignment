package handler

import (
	"assignment-project-new/database"
	"assignment-project-new/helpers"
	"assignment-project-new/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateScore(ctx *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(ctx)
	score := models.Scores{}

	if contentType == appJSON {
		if err := ctx.ShouldBindJSON(&score); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "message": err.Error()})
			return
		}
	} else {
		if err := ctx.ShouldBind(&score); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "message": err.Error()})
			return
		}
	}

	if err := db.Debug().Create(&score).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Database error", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"success": true, "data": score})
}

func GetAllScores(ctx *gin.Context) {
	db := database.GetDB()

	var scores []models.Scores
	if err := db.Debug().Find(&scores).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"success": true, "data": scores})
}

func UpdateScore(ctx *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(ctx)
	scoreID := ctx.Param("id")
	updatedScore := models.Scores{}

	if contentType == appJSON {
		if err := ctx.ShouldBindJSON(&updatedScore); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "message": err.Error()})
			return
		}
	} else {
		if err := ctx.ShouldBind(&updatedScore); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "message": err.Error()})
			return
		}
	}

	var existingScore models.Scores
	if err := db.Debug().Where("id = ?", scoreID).First(&existingScore).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Score not found"})
		return
	}

	existingScore.AssignmentTitle = updatedScore.AssignmentTitle
	existingScore.Description = updatedScore.Description
	existingScore.Score = updatedScore.Score

	if err := db.Debug().Save(&existingScore).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Database error", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"success": true, "data": existingScore})
}

func DeleteScore(ctx *gin.Context) {
	db := database.GetDB()
	scoreID := ctx.Param("id")
	score := models.Scores{}

	if err := db.Debug().Where("id = ?", scoreID).First(&score).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Score not found", "message": err.Error()})
		return
	}

	if err := db.Debug().Delete(&score).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"success": true, "message": "Score deleted successfully"})
}
