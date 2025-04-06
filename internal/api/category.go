package api

import (
	"ecommerce-product/constants"
	"ecommerce-product/helpers"
	"ecommerce-product/internal/interfaces"
	"ecommerce-product/internal/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CategoryAPI struct {
	CategoryService interfaces.ICategoryService
}

func (api *CategoryAPI) CreateCategory(e echo.Context) error {
	var (
		log = helpers.Logger
	)
	req := models.ProductCategory{}

	if err := e.Bind(&req); err != nil {
		log.Error("failed to parse request: ", err)
		return helpers.SendResponseHTTP(e, http.StatusBadRequest, constants.ErrFailedBadRequest, err)
	}

	if err := req.Validate(); err != nil {
		log.Error("failed to validate request: ", err)
		return helpers.SendResponseHTTP(e, http.StatusBadRequest, constants.ErrFailedBadRequest, err)
	}

	reps, err := api.CategoryService.CreateCategory(e.Request().Context(), &req)
	if err != nil {
		log.Error("failed to create category: ", err)
		return helpers.SendResponseHTTP(e, http.StatusInternalServerError, constants.ErrServerError, err)
	}

	return helpers.SendResponseHTTP(e, http.StatusOK, constants.Success, reps)
}

func (api *CategoryAPI) UpdateProductCategory(e echo.Context) error {
	var (
		log           = helpers.Logger
		categoryIDstr = e.Param("id")
	)

	categoryID, err := strconv.Atoi(categoryIDstr)
	if err != nil || categoryID == 0 {
		log.Error("failed to get categoryID: ", err)
		return helpers.SendResponseHTTP(e, http.StatusBadRequest, constants.ErrFailedBadRequest, err)
	}

	req := models.ProductCategory{}

	if err := e.Bind(&req); err != nil {
		log.Error("failed to parse request: ", err)
		return helpers.SendResponseHTTP(e, http.StatusBadRequest, constants.ErrFailedBadRequest, err)
	}

	err = api.CategoryService.UpdateProductCategory(e.Request().Context(), categoryID, req)
	if err != nil {
		log.Error("failed to update product category: ", err)
		return helpers.SendResponseHTTP(e, http.StatusInternalServerError, constants.ErrServerError, err)
	}

	return helpers.SendResponseHTTP(e, http.StatusOK, constants.Success, nil)
}

func (api *CategoryAPI) DeleteCategory(e echo.Context) error {
	var (
		log           = helpers.Logger
		categoryIDstr = e.Param("id")
	)

	categoryID, err := strconv.Atoi(categoryIDstr)
	if err != nil || categoryID == 0 {
		log.Error("failed to get categoryID: ", err)
		return helpers.SendResponseHTTP(e, http.StatusBadRequest, constants.ErrFailedBadRequest, err)
	}


	err = api.CategoryService.DeleteCategory(e.Request().Context(), categoryID)
	if err != nil {
		log.Error("failed to delete product category: ", err)
		return helpers.SendResponseHTTP(e, http.StatusInternalServerError, constants.ErrServerError, err)
	}

	return helpers.SendResponseHTTP(e, http.StatusOK, constants.Success, nil)
}

func (api *CategoryAPI) GetAllCategory(e echo.Context) error {
	var (
		log           = helpers.Logger
	)

	resp, err := api.CategoryService.GetAllCategory(e.Request().Context())
	if err != nil {
		log.Error("failed to get all product category: ", err)
		return helpers.SendResponseHTTP(e, http.StatusInternalServerError, constants.ErrServerError, err)
	}

	return helpers.SendResponseHTTP(e, http.StatusOK, constants.Success, resp)
}
