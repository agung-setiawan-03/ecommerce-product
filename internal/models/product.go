package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Product struct {
	ID          int
	Name        string    `json:"name" gorm:"column:name;type:varchar(255)" validate:"required"`
	Description string    `json:"description" gorm:"column:description;type:text" validate:"required"`
	Price       float64   `json:"price" gorm:"column:price;type:decimal(10,2)" validate:"required"`
	UpdatedAt   time.Time `json:"-"`
	CreatedAt   time.Time `json:"-"`
}

func (*Product) TableName() string {
	return "products"
}

func (l Product) Validate() error {
	v := validator.New()
	return v.Struct(l)
}

type ProductCategory struct {
	ID        int
	Name      string    `json:"name" gorm:"column:name;type:varchar(255)" validate:"required"`
	UpdatedAt time.Time `json:"-"`
	CreatedAt time.Time `json:"-"`
}

func (*ProductCategory) TableName() string {
	return "product_categories"
}

func (l ProductCategory) Validate() error {
	v := validator.New()
	return v.Struct(l)
}

type ProductVariants struct {
	ID        int
	ProductID int       `json:"product_id" gorm:"column:product_id;type:int"`
	Color     string    `json:"color" gorm:"column:color;type:varchar(50)" validate:"required"`
	Size      string    `json:"size" gorm:"column:size;type:varchar(10)" validate:"required"`
	Quantity  int       `json:"quantity" gorm:"column:quantity;type:int"`
	UpdatedAt time.Time `json:"-"`
	CreatedAt time.Time `json:"-"`
}

func (*ProductVariants) TableName() string {
	return "product_variants"
}

func (l ProductVariants) Validate() error {
	v := validator.New()
	return v.Struct(l)
}
