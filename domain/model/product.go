package model

import (
	"math/rand"

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
	ID          int32   `json:"id" gorm:"column:id;primaryKey;type:int" valid:"notnull"`
	Name        string  `json:"name" gorm:"column:name;type:varchar(255)" valid:"notnull"`
	Description string  `json:"description" gorm:"column:description;type:varchar(255)" valid:"notnull"`
	Price       float32 `json:"price" gorm:"column:price;type:float" valid:"notnull"`
}

func NewProduct(name, description string, price float32) (*Product, error) {
	product := Product{
		Name:        name,
		Description: description,
		Price:       price,
	}

	product.ID = int32(rand.Intn(100000))

	if err := product.isValid(); err != nil {
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
