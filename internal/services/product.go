package services

import (
	"context"
	"ecommerce-product/internal/interfaces"
	"ecommerce-product/internal/models"
	"encoding/json"

	"github.com/pkg/errors"
)

type ProductService struct {
	ProductRepo interfaces.IProductRepository
}

func (s *ProductService) CreateProduct(ctx context.Context, req *models.Product) (*models.Product, error) {
	resp := req
	err := s.ProductRepo.InsertNewProduct(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to insert new product")
	}
	return resp, nil
}

func (s *ProductService) UpdateProduct(ctx context.Context, productID int, req models.Product) error {
	jsonReq, err := json.Marshal(req)
	if err != nil {
		return errors.Wrap(err, "failed to marshal from request")	
	}

	newData := map[string]interface{}{}
	err = json.Unmarshal(jsonReq, &newData)
	if err != nil {
		return errors.Wrap(err, "failed to unmarshal to map")
	}

	err = s.ProductRepo.UpdateProduct(ctx, productID, newData)
	if err != nil {
		return errors.Wrap(err, "failed to update product")
	}

	
	return nil 
}

func (s *ProductService) UpdateProductVariant(ctx context.Context, variantID int, req models.ProductVariants) error {
	jsonReq, err := json.Marshal(req)
	if err != nil {
		return errors.Wrap(err, "failed to marshal from request")
	}

	newData := map[string]interface{}{}
	err = json.Unmarshal(jsonReq, &newData)
	if err != nil {
		return errors.Wrap(err, "failed to unmarshal to map")
	}

	err = s.ProductRepo.UpdateProductVariant(ctx, variantID, newData)
	if err != nil {
		return errors.Wrap(err, "failed to update product variant")
	}

	
	return nil
}

func (s *ProductService) DeleteProduct(ctx context.Context, productID int) error {
	return s.ProductRepo.DeleteProduct(ctx, productID)
}


func (s *ProductService) GetAllProducts(ctx context.Context, page, limit int) ([]models.Product, error) {
	return s.ProductRepo.GetAllProducts(ctx, page, limit)
}
func (s *ProductService) GetProductDetail(ctx context.Context, productID int) (models.Product, error) {
	return s.ProductRepo.GetProductDetail(ctx, productID)
}
