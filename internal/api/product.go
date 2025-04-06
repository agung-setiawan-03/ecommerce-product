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

type ProductAPI struct {
	ProductService interfaces.IProductService
}

func (api *ProductAPI) CreateProduct(e echo.Context) error {
	var (
		log = helpers.Logger
	)
	req := models.Product{}

	if err := e.Bind(&req); err != nil {
		log.Error("failed to parse request: ", err)
		return helpers.SendResponseHTTP(e, http.StatusBadRequest, constants.ErrFailedBadRequest, err)
	}

	if err := req.Validate(); err != nil {
		log.Error("failed to validate request: ", err)
		return helpers.SendResponseHTTP(e, http.StatusBadRequest, constants.ErrFailedBadRequest, err)
	}

	reps, err := api.ProductService.CreateProduct(e.Request().Context(), &req)
	if err != nil {
		log.Error("failed to create product: ", err)
		return helpers.SendResponseHTTP(e, http.StatusInternalServerError, constants.ErrServerError, err)
	}

	return helpers.SendResponseHTTP(e, http.StatusOK, constants.Success, reps)
}

func (api *ProductAPI) UpdateProduct(e echo.Context) error {
	var (
		log          = helpers.Logger
		productIDstr = e.Param("id")
	)

	productID, err := strconv.Atoi(productIDstr)
	if err != nil || productID == 0 {
		log.Error("failed to get productID: ", err)
		return helpers.SendResponseHTTP(e, http.StatusBadRequest, constants.ErrFailedBadRequest, err)
	}

	req := models.Product{}

	if err := e.Bind(&req); err != nil {
		log.Error("failed to parse request: ", err)
		return helpers.SendResponseHTTP(e, http.StatusBadRequest, constants.ErrFailedBadRequest, err)
	}

	err = api.ProductService.UpdateProduct(e.Request().Context(), productID, req)
	if err != nil {
		log.Error("failed to update product: ", err)
		return helpers.SendResponseHTTP(e, http.StatusInternalServerError, constants.ErrServerError, err)
	}

	return helpers.SendResponseHTTP(e, http.StatusOK, constants.Success, nil)
}

func (api *ProductAPI) UpdateProductVariant(e echo.Context) error {
	var (
		log          = helpers.Logger
		variantIDstr = e.Param("id")
	)

	variantID, err := strconv.Atoi(variantIDstr)
	if err != nil || variantID == 0 {
		log.Error("failed to get variantID: ", err)
		return helpers.SendResponseHTTP(e, http.StatusBadRequest, constants.ErrFailedBadRequest, err)
	}

	req := models.ProductVariants{}

	if err := e.Bind(&req); err != nil {
		log.Error("failed to parse request: ", err)
		return helpers.SendResponseHTTP(e, http.StatusBadRequest, constants.ErrFailedBadRequest, err)
	}

	err = api.ProductService.UpdateProductVariant(e.Request().Context(), variantID, req)
	if err != nil {
		log.Error("failed to update product variants: ", err)
		return helpers.SendResponseHTTP(e, http.StatusInternalServerError, constants.ErrServerError, err)
	}

	return helpers.SendResponseHTTP(e, http.StatusOK, constants.Success, nil)
}

func (api *ProductAPI) DeleteProduct(e echo.Context) error {
	var (
		log          = helpers.Logger
		productIDstr = e.Param("id")
	)

	productID, err := strconv.Atoi(productIDstr)
	if err != nil || productID == 0 {
		log.Error("failed to get productID: ", err)
		return helpers.SendResponseHTTP(e, http.StatusBadRequest, constants.ErrFailedBadRequest, err)
	}

	err = api.ProductService.DeleteProduct(e.Request().Context(), productID)
	if err != nil {
		log.Error("failed to delete product: ", err)
		return helpers.SendResponseHTTP(e, http.StatusInternalServerError, constants.ErrServerError, err)
	}

	return helpers.SendResponseHTTP(e, http.StatusOK, constants.Success, nil)
}

func (api *ProductAPI) GetAllProducts(e echo.Context) error {
	var (
		log      = helpers.Logger
		pageStr  = e.QueryParam("page")
		limitStr = e.QueryParam("limit")
	)

	page, err := strconv.Atoi(pageStr)
	if err != nil && pageStr != "" {
		log.Error("failed to get page: ", err)
		return helpers.SendResponseHTTP(e, http.StatusBadRequest, constants.ErrFailedBadRequest, err)
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil && limitStr != "" {
		log.Error("failed to get limit: ", err)
		return helpers.SendResponseHTTP(e, http.StatusBadRequest, constants.ErrFailedBadRequest, err)
	}

	resp, err := api.ProductService.GetAllProducts(e.Request().Context(), page, limit)
	if err != nil {
		log.Error("failed to delete product: ", err)
		return helpers.SendResponseHTTP(e, http.StatusInternalServerError, constants.ErrServerError, err)
	}

	return helpers.SendResponseHTTP(e, http.StatusOK, constants.Success, resp)
}

func (api *ProductAPI) GetProductDetail(e echo.Context) error {
	var (
		log          = helpers.Logger
		productIDstr = e.Param("id")
	)

	productID, err := strconv.Atoi(productIDstr)
	if err != nil || productID == 0 {
		log.Error("failed to get productID: ", err)
		return helpers.SendResponseHTTP(e, http.StatusBadRequest, constants.ErrFailedBadRequest, err)
	}

	resp, err := api.ProductService.GetProductDetail(e.Request().Context(), productID)
	if err != nil {
		log.Error("failed to get product detail: ", err)
		return helpers.SendResponseHTTP(e, http.StatusInternalServerError, constants.ErrServerError, err)
	}

	return helpers.SendResponseHTTP(e, http.StatusOK, constants.Success, resp)
}
