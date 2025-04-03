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
		userGroup.DELETE("/:id", DeleteUserById)
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
	c.String(200, "user added")
}

func UpdateUser(c *gin.Context) {
	c.String(200, "user added")
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