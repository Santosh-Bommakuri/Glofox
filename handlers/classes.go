package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"glofox/models"
	"glofox/services"
)


type ClassHandler struct {
	classService services.ClassService
}


func NewClassHandler(classService services.ClassService) *ClassHandler {
	return &ClassHandler{classService: classService}
}


func (h *ClassHandler) CreateClass(c *gin.Context) {
	var newClass models.Class
	if err := c.ShouldBindJSON(&newClass); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body", "details": err.Error()})
		return
	}

	if newClass.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required field: name"})
		return
	}
	if newClass.StartDate.IsZero() || newClass.EndDate.IsZero() {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required fields: start_date or end_date"})
		return
	}
	if newClass.EndDate.Before(newClass.StartDate) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "End date cannot be before start date"})
		return
	}
	if newClass.Capacity <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Capacity must be greater than zero"})
		return
	}

	err := h.classService.CreateClasses(newClass)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create classes", "details": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Classes created successfully"})
}