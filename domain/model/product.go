package model

import (
	"time"
	uuid "github.com/satori/go.uuid"
	"github.com/asaskevich/govalidator"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type ProductRepositoryInterface interface {
	Create(product *Product) (*Product, error)
	FindAll() ([]*Product, error)
}

type Product struct {
	ID string `json:"id" valid:"uuid"`
	Name string `json:"name" valid:"notnull"`
	Value float64 `json:"value" valid:"notnull"`
	CreatedAt time.Time `json:"createdAt" valid:"-"`
	UpdatedAt time.Time `json:"updatedAt" valid:"-"`
}

func NewProduct(name string, value float64) (*Product, error) {
	product  := Product{
		Name: name,
		Value: value
	}

	product.ID = uuid.NewV4().String()
	product.CreatedAt = time.Now()

	if err := product.isValid(); err != nil{
		return nil, err
	}

	return &product, nil
}

func (p *Product) isValid() error {
	_, err := govalidator.ValidateStruct(p)
	if err != nil {
		return err
	}

	return nil
}