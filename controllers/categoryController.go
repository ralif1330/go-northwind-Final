package controllers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"code.id.northwind/repositories/dbContext"
	"code.id.northwind/services"
	"github.com/gin-gonic/gin"

)

type CategoryController struct {
	categoryService *services.CategoryService
}

// declare constructor
func NewCategoryController(categoryService *services.CategoryService) *CategoryController {
	return &CategoryController{
		categoryService: categoryService,
	}
}

// method
func (categoryController CategoryController) GetListCategory(ctx *gin.Context) {
	response, responseErr := categoryController.categoryService.GetListCategory(ctx)

	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)

	// ctx.JSON(http.StatusOK, "Hello gin framework")
}

func (categoryController CategoryController) GetCategory(ctx *gin.Context) {

	categoryId, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response, responseErr := categoryController.categoryService.GetCategory(ctx, int64(categoryId))
	if responseErr != nil {

		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (categoryController CategoryController) CreateCategory(ctx *gin.Context) {

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading create category request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var category dbContext.CreateCategoryParams
	err = json.Unmarshal(body, &category)
	if err != nil {
		log.Println("Error while unmarshaling create category request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response, responseErr := categoryController.categoryService.CreateCategory(ctx, &category)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)

}

func (categoryController CategoryController) UpdateCategory(ctx *gin.Context) {

	categoryId, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading update category request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var category dbContext.CreateCategoryParams
	err = json.Unmarshal(body, &category)
	if err != nil {
		log.Println("Error while unmarshaling update category request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response := categoryController.categoryService.UpdateCategory(ctx, &category, int64(categoryId))
	if response != nil {
		ctx.AbortWithStatusJSON(response.Status, response)
		return
	}

	ctx.JSON(http.StatusOK, response)

}

func (categoryController CategoryController) DeleteCategory(ctx *gin.Context) {

	categoryId, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	responseErr := categoryController.categoryService.DeleteCategory(ctx, int64(categoryId))
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.Status(http.StatusNoContent)
}
