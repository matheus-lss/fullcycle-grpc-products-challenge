package model

import (
	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type ProductRepositoryInterface interface {
	Create(product *Product) (*Product, error)
	FindAll() ([]*Product, error)
}

type Product struct {
	ID    string  `json:"id" gorm:"column:id;primaryKey;type:uuid" valid:"uuid"`
	Name  string  `json:"name" gorm:"column:name;type:varchar(255)" valid:"notnull"`
	Value float64 `json:"value" gorm:"column:value;type:float" valid:"notnull"`
}

func NewProduct(name string, value float64) (*Product, error) {
	product := Product{
		Name:  name,
		Value: value,
	}

	product.ID = uuid.NewV4().String()

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
