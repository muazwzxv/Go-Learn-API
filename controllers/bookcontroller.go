package controllers

import (
	"Go-Learn-API/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateBookDTO := DTO new book
type CreateBookDTO struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

// UpdateBookDTO := DTO update book
type UpdateBookDTO struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

// GetAllBooks := returns all
func GetAllBooks(ctx *gin.Context) {
	books, err := models.Model.GetAllBooks()
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"erorr": err})
	}

	ctx.JSON(http.StatusOK, gin.H{"Data": books})
}

// CreateBooks := create new book
func CreateBooks(ctx *gin.Context) {
	var input CreateBookDTO

	if err := ctx.ShouldBindJSON(&input); err != nil {
		 ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book := models.Book{
		Title:  input.Title,
		Author: input.Author,
	}

	_, err := models.Model.CreateBooks(&book)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return 
	}

	ctx.JSON(http.StatusOK, gin.H{"data": book})
}

// GetByIDBooks := returns by ID
func GetByIDBooks(ctx *gin.Context) {

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		panic(err)
	}

	if book, err := models.Model.GetByIDBooks(id); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err})
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{"data": book})
		return
	}
}

// UpdateByIDBooks := update by ID
func UpdateByIDBooks(ctx *gin.Context) {

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		panic(err)
	}

	// validate input
	var validate UpdateBookDTO
	if err := ctx.ShouldBindJSON(&validate); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book := models.Book{
		Title: validate.Title,
		Author: validate.Author,
	}	

	if updated, err := models.Model.UpdateByIDBooks(id, &book); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err})
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{"data": updated})
		return 
	}
}

// DeleteByIDBooks := delete by id
func DeleteByIDBooks(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		panic(err)
	}

	if err := models.Model.DeleteByIDBooks(id); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err})
		return
	} 
	ctx.JSON(http.StatusOK, gin.H{"deleted": true})
	return
}
