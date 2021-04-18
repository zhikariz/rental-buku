package handler

import (
	"net/http"
	"rental-buku/auth"
	"rental-buku/category"
	"rental-buku/helper"
	"rental-buku/user"

	"github.com/gin-gonic/gin"
)

type categoryHandler struct {
	categoryService category.Service
	authService     auth.Service
}

func NewCategoryHandler(categoryService category.Service, authService auth.Service) *categoryHandler {
	return &categoryHandler{categoryService, authService}
}

func (h *categoryHandler) CreateCategory(c *gin.Context) {
	var input category.CreateCategoryInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Create Category failed !", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	role := currentUser.Role

	if role != "Admin" {
		response := helper.APIResponse("U're not authorized to do this!", http.StatusUnauthorized, "error", nil)
		c.JSON(http.StatusUnauthorized, response)
		return
	}

	newCategory, err := h.categoryService.CreateCategory(input)

	if err != nil {
		response := helper.APIResponse("Create Category failed !", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := category.FormatCategory(newCategory)
	response := helper.APIResponse("Category has been created !", http.StatusOK, "succes", formatter)

	c.JSON(http.StatusOK, response)

}

func (h *categoryHandler) UpdateCategory(c *gin.Context) {
	var uri category.CategoryUriInput
	var input category.UpdateCategoryInput

	err := c.ShouldBindUri(&uri)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Update category failed !", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err = c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Update category failed !", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	role := currentUser.Role

	if role != "Admin" {
		response := helper.APIResponse("U're not authorized to do this!", http.StatusUnauthorized, "error", nil)
		c.JSON(http.StatusUnauthorized, response)
		return
	}

	newCategory, err := h.categoryService.UpdateCategory(uri.ID, input)

	if err != nil {
		response := helper.APIResponse("Update category failed !", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := category.FormatCategory(newCategory)
	response := helper.APIResponse("Category has been updated!", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)

}

func (h *categoryHandler) DeleteCategory(c *gin.Context) {
	var input category.CategoryUriInput

	err := c.ShouldBindUri(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Delete category failed !", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	role := currentUser.Role

	if role != "Admin" {
		response := helper.APIResponse("U're not authorized to do this!", http.StatusUnauthorized, "error", nil)
		c.JSON(http.StatusUnauthorized, response)
		return
	}

	isDeleted, err := h.categoryService.DeleteCategory(input)

	if err != nil {
		response := helper.APIResponse("Delete category failed !", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data := gin.H{
		"is_deleted": isDeleted,
	}

	metaMessage := "Category cannot be deleted !"

	if isDeleted {
		metaMessage = "Category has been deleted !"
	}

	response := helper.APIResponse(metaMessage, http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}

func (h *categoryHandler) GetCategories(c *gin.Context) {
	campaigns, err := h.categoryService.GetCategories()
	if err != nil {
		response := helper.APIResponse("Get categories failed !", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIResponse("List of category", http.StatusOK, "success", campaigns)
	c.JSON(http.StatusOK, response)
	return

}

func (h *categoryHandler) GetCategoryById(c *gin.Context) {
	var uri category.CategoryUriInput

	err := c.ShouldBindUri(&uri)

	if err != nil {
		response := helper.APIResponse("Get category failed !", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	campaign, err := h.categoryService.GetCategoryById(uri.ID)
	if err != nil {
		response := helper.APIResponse("Get category failed !", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIResponse("Detail category", http.StatusOK, "success", campaign)
	c.JSON(http.StatusOK, response)
}
