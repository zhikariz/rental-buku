package handler

import (
	"net/http"
	"rental-buku/auth"
	"rental-buku/book"
	"rental-buku/helper"
	"rental-buku/user"

	"github.com/gin-gonic/gin"
)

type bookHandler struct {
	bookService book.Service
	authService auth.Service
}

func NewBookHandler(bookService book.Service, authService auth.Service) *bookHandler {
	return &bookHandler{bookService, authService}
}

func (h *bookHandler) CreateBook(c *gin.Context) {
	var input book.CreateBookInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Create Book failed !", http.StatusUnprocessableEntity, "error", errorMessage)
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

	newBook, err := h.bookService.CreateBook(input)

	if err != nil {
		response := helper.APIResponse("Create Book failed !", http.StatusUnprocessableEntity, "error", err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := book.FormatBook(newBook)
	response := helper.APIResponse("Book has been created !", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)

}

func (h *bookHandler) UpdateBook(c *gin.Context) {
	var uri book.BookUriInput
	var input book.UpdateBookInput

	err := c.ShouldBindUri(&uri)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Update book failed !", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err = c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Update book failed !", http.StatusUnprocessableEntity, "error", errorMessage)
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

	updatedBook, err := h.bookService.UpdateBook(uri.ID, input)

	if err != nil {
		response := helper.APIResponse("Update book failed !", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := book.FormatBook(updatedBook)
	response := helper.APIResponse("Book has been updated!", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)

}

func (h *bookHandler) DeleteBook(c *gin.Context) {
	var input book.BookUriInput

	err := c.ShouldBindUri(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Delete book failed !", http.StatusUnprocessableEntity, "error", errorMessage)
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

	isDeleted, err := h.bookService.DeleteBook(input)

	if err != nil {
		response := helper.APIResponse("Delete book failed !", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data := gin.H{
		"is_deleted": isDeleted,
	}

	metaMessage := "Book cannot be deleted !"

	if isDeleted {
		metaMessage = "Book has been deleted !"
	}

	response := helper.APIResponse(metaMessage, http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}

func (h *bookHandler) GetBooks(c *gin.Context) {
	books, err := h.bookService.GetBooks()
	if err != nil {
		response := helper.APIResponse("Get books failed !", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIResponse("List of book", http.StatusOK, "success", books)
	c.JSON(http.StatusOK, response)
	return
}

func (h *bookHandler) GetBookById(c *gin.Context) {
	var uri book.BookUriInput

	err := c.ShouldBindUri(&uri)

	if err != nil {
		response := helper.APIResponse("Get book failed !", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	book, err := h.bookService.GetBookById(uri.ID)
	if err != nil {
		response := helper.APIResponse("Get book failed !", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIResponse("Detail book", http.StatusOK, "success", book)
	c.JSON(http.StatusOK, response)
}
