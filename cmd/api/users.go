package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app *application) getAllUsers(c *gin.Context) {
	users, err := app.models.Users.GetAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retreive users"})
	}

	c.JSON(http.StatusOK, users)
}
