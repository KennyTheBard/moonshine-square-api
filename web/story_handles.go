package web

import (
	"net/http"

	"github.com/gin-gonic/gin"

	data "../data"
	model "../model"
)

// CreateStory creates a story
func CreateStory(c *gin.Context) {
	story := model.Story{Title: c.PostForm("title"), Text: c.PostForm("text")}
	data.DB.Save(&story)

	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Story created successfully!", "resourceId": story.ID})
}

// GetAllStories retrieves all stories
func GetAllStories(context *gin.Context) {
	var stories []model.Story
	data.DB.Find(&stories)

	var dtos []model.StoryDto
	for _, item := range stories {
		dtos = append(dtos, model.StoryDto{ID: item.ID, Title: item.Title, Text: item.Text})
	}
	context.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "body": dtos})
}

// GetStory retrieves a story
func GetStory(context *gin.Context) {
	var story model.Story
	id := context.Param("id")
	data.DB.First(&story, id)

	dto := model.StoryDto{ID: story.ID, Title: story.Title, Text: story.Text}

	context.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "body": dto})
}

// UpdateStory updates a story
func UpdateStory(context *gin.Context) {
	var story model.Story
	id := context.Param("id")

	data.DB.First(&story, id)

	if story.ID == 0 {
		context.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
		return
	}

	data.DB.Model(&story).Update("title", context.PostForm("title"))
	data.DB.Model(&story).Update("text", context.PostForm("text"))

	context.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Todo updated successfully!"})
}

// DeleteStory removes a story
func DeleteStory(context *gin.Context) {
	var story model.Story
	id := context.Param("id")

	data.DB.First(&story, id)

	if story.ID == 0 {
		context.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
		return
	}

	data.DB.Delete(&story)
	context.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Story deleted successfully!"})
}