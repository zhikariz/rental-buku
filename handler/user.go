package handler

import (
	"fmt"
	"net/http"
	"rental-buku/auth"
	"rental-buku/helper"
	"rental-buku/user"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
	authService auth.Service
}

func NewUserHandler(userService user.Service, authService auth.Service) *userHandler {
	return &userHandler{userService, authService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Register account failed !", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	isEmailAvailable, err := h.userService.IsEmailAvailable(input.Email)
	if err != nil {
		errorMessage := gin.H{"errors": "Server Error"}
		response := helper.APIResponse("Email checking failed !", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	if !isEmailAvailable {
		response := helper.APIResponse("Email has been registered !", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := h.userService.RegisterUser(input)

	if err != nil {
		response := helper.APIResponse("Register account failed !", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	token, err := h.authService.GenerateToken(newUser)

	if err != nil {
		response := helper.APIResponse("Register account failed !", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := user.FormatUser(newUser, token)
	response := helper.APIResponse("Account has been registered", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)

}

func (h *userHandler) Login(c *gin.Context) {
	var input user.LoginUserInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Login failed !", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loggedInUser, err := h.userService.Login(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Login failed !", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	token, err := h.authService.GenerateToken(loggedInUser)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Login failed !", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(loggedInUser, token)
	response := helper.APIResponse("Successfully Logged in!", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) UploadPhoto(c *gin.Context) {
	file, err := c.FormFile("photo")
	if err != nil {
		data := gin.H{
			"is_uploaded": false,
		}
		response := helper.APIResponse("Failed to upload photo image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
	}

	currentUser := c.MustGet("currentUser").(user.User)
	userId := currentUser.ID

	path := fmt.Sprintf("images/%d-%s", userId, file.Filename)

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{
			"is_uploaded": false,
		}
		response := helper.APIResponse("Failed to upload photo image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
	}

	_, err = h.userService.SavePhoto(int(userId), path)
	if err != nil {
		data := gin.H{
			"is_uploaded": false,
		}
		response := helper.APIResponse("Failed to upload avatar image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
	}

	data := gin.H{
		"is_uploaded": true,
	}
	response := helper.APIResponse("Photo successfully uploaded", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)

}

func (h *userHandler) ResetPassword(c *gin.Context) {
	var input user.ResetPasswordInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Reset password account failed !", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	_, err = h.userService.ResetPassword(input)
	if err != nil {
		data := gin.H{
			"is_reset": false,
		}
		response := helper.APIResponse("Failed to reset password", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{
		"is_reset": true,
	}
	response := helper.APIResponse("Password successfully changed", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)

}
