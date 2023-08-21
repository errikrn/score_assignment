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

	Score := models.Score{}

	if contentType == appJSON {
		ctx.ShouldBindJSON(&Score)
	} else {
		ctx.ShouldBind(&Score)
	}

	err := db.Debug().Create(&Score).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": Score,
	})
}

func UpdateScore(ctx *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(ctx)

	var updatedScore models.Score

	if contentType == appJSON {
		err := ctx.ShouldBindJSON(&updatedScore)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error":   "Bad request",
				"message": err.Error(),
			})
			return
		}
	} else {
		err := ctx.ShouldBind(&updatedScore)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error":   "Bad request",
				"message": err.Error(),
			})
			return
		}
	}

	var existingScore models.Score
	scoreID := ctx.Param("id")

	if err := db.Debug().Where("id = ?", scoreID).First(&existingScore).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Score not found",
		})
		return
	}

	existingScore.AssignmentTitle = updatedScore.AssignmentTitle
	existingScore.Description = updatedScore.Description

	if err := db.Debug().Save(&existingScore).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Database error",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    existingScore,
	})
}

func GetAllScoreByStudent(ctx *gin.Context) {
	db := database.GetDB()

	var scores []models.Score
	if err := db.Debug().Find(&scores).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Internal Server Error",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    scores,
	})
}

func DeleteScore(ctx *gin.Context) {
	db := database.GetDB()

	scoreID := ctx.Param("id")
	var score models.Score

	if err := db.Debug().Where("id = ?", scoreID).First(&score).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error":   "Score not found",
			"message": err.Error(),
		})
		return
	}

	if err := db.Debug().Delete(&score).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Internal Server Error",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Score deleted successfully",
	})
}
