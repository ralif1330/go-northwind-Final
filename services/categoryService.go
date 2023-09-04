package services

import (
	"net/http"

	"code.id.northwind/models"
	"code.id.northwind/repositories"
	"code.id.northwind/repositories/dbContext"
	"github.com/gin-gonic/gin"
)

type CategoryService struct {
	categoryRepository *repositories.CategoryRepository
}

func NewCategoryService(categoryRepository *repositories.CategoryRepository) *CategoryService {
	return &CategoryService{
		categoryRepository: categoryRepository,
	}
}

func (cs CategoryService) GetListCategory(ctx *gin.Context) ([]*models.Category, *models.ResponseError) {
	return cs.categoryRepository.GetListCategory(ctx)
}

func (cs CategoryService) GetCategory(ctx *gin.Context, id int64) (*models.Category, *models.ResponseError) {
	return cs.categoryRepository.GetCategory(ctx, id)
}

func (cs CategoryService) CreateCategory(ctx *gin.Context, categoryParams *dbContext.CreateCategoryParams) (*models.Category, *models.ResponseError) {
	responseErr := validateCategory(categoryParams)
	if responseErr != nil {
		return nil, responseErr
	}

	return cs.categoryRepository.CreateCategory(ctx, categoryParams)
}

func (cs CategoryService) UpdateCategory(ctx *gin.Context, categoryParams *dbContext.CreateCategoryParams, id int64) *models.ResponseError {
	responseErr := validateCategory(categoryParams)
	if responseErr != nil {
		return responseErr
	}

	return cs.categoryRepository.UpdateCategory(ctx, categoryParams)
}

func (cs CategoryService) DeleteCategory(ctx *gin.Context, id int64) *models.ResponseError {
	return cs.categoryRepository.DeleteCategory(ctx, id)
}

func validateCategory(categoryParams *dbContext.CreateCategoryParams) *models.ResponseError {
	if categoryParams.CategoryID == 0 {
		return &models.ResponseError{
			Message: "Invalid category id",
			Status:  http.StatusBadRequest,
		}
	}

	if categoryParams.CategoryName == "" {
		return &models.ResponseError{
			Message: "Invalid category name",
			Status:  http.StatusBadRequest,
		}
	}

	return nil

}
