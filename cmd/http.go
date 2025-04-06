package cmd

import (
	"ecommerce-product/external"
	"ecommerce-product/helpers"
	"ecommerce-product/internal/api"
	"ecommerce-product/internal/interfaces"
	"ecommerce-product/internal/repository"
	"ecommerce-product/internal/services"

	"github.com/labstack/echo/v4"
)

func ServeHTTP() {
	d := dependencyInject()
	healthcheckAPI := &api.HealthCheckAPI{}

	e := echo.New()
	e.GET("/healthcheck", healthcheckAPI.HealthCheck)

	productV1 := e.Group("/products/v1")
	productV1.POST("", d.ProductAPI.CreateProduct, d.MiddlewareValidateAuth)
	productV1.PUT("/:id", d.ProductAPI.UpdateProduct, d.MiddlewareValidateAuth)
	productV1.PUT("/variant/:id", d.ProductAPI.UpdateProductVariant, d.MiddlewareValidateAuth)
	productV1.DELETE("/:id", d.ProductAPI.DeleteProduct, d.MiddlewareValidateAuth)
	productV1.GET("/list", d.ProductAPI.GetAllProducts)
	productV1.GET("/:id", d.ProductAPI.GetProductDetail)

	categoryV1 := e.Group("/products/v1/category")
	categoryV1.POST("/create", d.CategoryAPI.CreateCategory, d.MiddlewareValidateAuth)
	categoryV1.PUT("/:id", d.CategoryAPI.UpdateProductCategory, d.MiddlewareValidateAuth)
	categoryV1.DELETE("/:id", d.CategoryAPI.DeleteCategory, d.MiddlewareValidateAuth)
	categoryV1.GET("", d.CategoryAPI.GetAllCategory)

	e.Start(":" + helpers.GetEnv("PORT", "9001"))
}

type Dependency struct {
	External       interfaces.IExternal
	HealthcheckAPI *api.HealthCheckAPI

	ProductAPI  interfaces.IProductAPI
	CategoryAPI interfaces.ICategoryAPI
}

func dependencyInject() Dependency {
	external := &external.External{}

	productRepo := &repository.ProductRepo{
		DB: helpers.DB,
		Redis: helpers.RedisClient,
	}
	categoryRepo := &repository.CategoryRepo{
		DB: helpers.DB,
	}

	productSvc := &services.ProductService{
		ProductRepo: productRepo,
	}
	categorySvc := &services.CategoryService{
		CategoryRepo: categoryRepo,
	}

	productAPI := &api.ProductAPI{
		ProductService: productSvc,
	}
	categoryAPI := &api.CategoryAPI{
		CategoryService: categorySvc,
	}

	return Dependency{
		External: 	 external,
		HealthcheckAPI: &api.HealthCheckAPI{},
		ProductAPI:     productAPI,
		CategoryAPI:    categoryAPI,
	}
}
