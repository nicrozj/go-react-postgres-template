package handlers

import (
	"backend/internal/db"
	"backend/internal/models"
	"backend/internal/types"
	"database/sql"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	userGroup := r.Group("/users")
	{
		userGroup.GET("", GetUsers)
		userGroup.GET("/:id", GetUserById)
		userGroup.POST("/:name", CreateUser)
		userGroup.DELETE("/:id", DeleteUserById)
		userGroup.PATCH("/:id", UpdateUser)
	}
}

func GetUsers(c *gin.Context) {
	users, err := models.GetUsers(db.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch users",
			"details": err.Error(),
		})
		return
	}
	if users == nil {
		users = []types.User{}
	}

	c.JSON(http.StatusOK, users)
}

func GetUserById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid user ID format",
			"details": "ID must be an integer",
		})
		return
	}

	user, err := models.GetUserById(db.DB, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "User not found",
			})
		} else {
			log.Printf("Failed to fetch user %d: %v", id, err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to fetch user",
			})
		}
		return
	}

	c.JSON(http.StatusOK, user)
}

func CreateUser(c *gin.Context) {
	name := c.Param("name")

	lastId, err := models.CreateUser(db.DB, &types.User{Name: name})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, lastId)
}

func UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
			c.JSON(400, gin.H{"error": "Invalid ID"})
			return
	}

	var input struct {
			Name string `json:"name"`
	}
	
	if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(400, gin.H{"error": "Bad request"})
			return
	}

	err = models.UpdateUser(c.Request.Context(), db.DB, id, input.Name)
	if err != nil {
			c.JSON(500, gin.H{"error": "Update failed: " + err.Error()})
			return
	}

	c.JSON(200, gin.H{
			"status": "updated",
			"id":     id,
	})
}

func DeleteUserById(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid user ID format",
			"details": "ID must be an integer",
		})
		return
	}

	if err := models.DeleteUserById(db.DB, id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User not found",
		}) 
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user deleted!",
	})
}