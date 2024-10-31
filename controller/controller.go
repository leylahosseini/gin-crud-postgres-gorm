package controller

import (
	"gin-crud-postgres-gorm/config"
	"gin-crud-postgres-gorm/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

var db = config.Connection_DB()

// GET All USERS
// @title Users API
// @version 1.0
// @description This is the API for managing users.
// @host localhost:8080
// @BasePath
// @Summary Get all users
// @Description Get a list of all users
// @Produce json
// @Success 200 {array} model.User
// @Router /users [get]
func GetAllUsers(c *gin.Context) {
	var users []model.User
	db.Find(&users)
	c.JSON(http.StatusOK, users)
}

// Get one user
// @Summary Get a user by ID
// @Description Get details of a user by their ID
// @Produce json
// @Param id path string true "User ID" example("1")
// @Success 200 {object} model.User
// @Router /users/{id} [get]
func GetUserByID(c *gin.Context) {
	id := c.Param("id")
	var user model.User

	if err := db.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User NOT FOUND"})
		return
	}
	c.JSON(http.StatusOK, user)

}

// POST Users
// @Summary Create a new user
// @Description Create a new user with the given information
// @Produce json
// @Param user body model.User true "User to create"
// @Success 201 {object} model.User

// @Router /users [post]
func CreateUser(c *gin.Context) {
	var user model.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	db.Create(&user)
	c.JSON(http.StatusOK, user)
}

// Update User
// @Summary Update a user
// @Description Update a user by their ID
// @Produce json
// @Param id path string true "User ID" example("1")
// @Param user body model.User true "User data to update"
// @Success 200 {object} model.User
// @Router /users/{id} [put]
func UpdateUser(c *gin.Context) {

	id := c.Param("id")
	var user model.User
	if err := db.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User Not Found"})
		return
	}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Save(&user)
	c.JSON(http.StatusOK, user)

}

// Delete User
// @Summary Delete a user
// @Description Delete a user by their ID
// @Produce json
// @Param id path string true "User ID"
// @Router /users/{id} [delete]
func DeleteUser(c *gin.Context) {

	id := c.Param("id")
	var user model.User
	if err := db.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "User Not Found"})
		return
	}

	db.Delete(&user)

	c.JSON(http.StatusOK, gin.H{"Message": "User Deleting ....."})
	return

}
