package repositories

import (
	"database/sql"
	"net/http"

	"code.id.northwind/models"
	"code.id.northwind/repositories/dbContext"
	"github.com/gin-gonic/gin"
)

type CategoryRepository struct {
	dbHandler   *sql.DB
	transaction *sql.Tx
}

func NewCategoryRepository(dbHandler *sql.DB) *CategoryRepository {
	return &CategoryRepository{
		dbHandler: dbHandler,
	}
}

func (cr CategoryRepository) GetListCategory(ctx *gin.Context) ([]*models.Category, *models.ResponseError) {

	store := dbContext.New(cr.dbHandler)
	categories, err := store.ListCategories(ctx)

	listCategories := make([]*models.Category, 0)

	for _, v := range categories {
		category := &models.Category{
			CategoryID:   v.CategoryID,
			CategoryName: v.CategoryName,
			Description:  v.Description,
			Picture:      nil,
		}
		listCategories = append(listCategories, category)
	}

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return listCategories, nil
}

func (cr CategoryRepository) GetCategory(ctx *gin.Context, id int64) (*models.Category, *models.ResponseError) {

	store := dbContext.New(cr.dbHandler)
	category, err := store.GetCategory(ctx, int16(id))

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return &category, nil
}

func (cr CategoryRepository) CreateCategory(ctx *gin.Context, categoryParams *dbContext.CreateCategoryParams) (*models.Category, *models.ResponseError) {

	store := dbContext.New(cr.dbHandler)
	category, err := store.CreateCategory(ctx, *categoryParams)

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Message,
			Status:  http.StatusInternalServerError,
		}
	}
	return category, nil
}

func (cr CategoryRepository) UpdateCategory(ctx *gin.Context, categoryParams *dbContext.CreateCategoryParams) *models.ResponseError {

	store := dbContext.New(cr.dbHandler)
	err := store.UpdateCategory(ctx, *categoryParams)

	if err != nil {
		return &models.ResponseError{
			Message: "error when update",
			Status:  http.StatusInternalServerError,
		}
	}
	return &models.ResponseError{
		Message: "data has been update",
		Status:  http.StatusOK,
	}
}

func (cr CategoryRepository) DeleteCategory(ctx *gin.Context, id int64) *models.ResponseError {

	store := dbContext.New(cr.dbHandler)
	err := store.DeleteCategory(ctx, int16(id))

	if err != nil {
		return &models.ResponseError{
			Message: "error when update",
			Status:  http.StatusInternalServerError,
		}
	}
	return &models.ResponseError{
		Message: "data has been deleted",
		Status:  http.StatusOK,
	}
}
