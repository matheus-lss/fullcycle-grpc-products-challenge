package repository

import (
	"database/sql"

	"github.com/jinzhu/gorm"
	"github.com/matheuslssilva/fullcycle-grpc-products-challenge/domain/model"
)

type ProductRepositoryDb struct {
	Db *gorm.DB
}

func (r ProductRepositoryDb) Create(product *model.Product) (*model.Product, error) {
	err := r.Db.Create(product).Error
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (r ProductRepositoryDb) FindAll() ([]*model.Product, error) {
	products := []*model.Product{}
	err := r.Db.Find(&products).Error
	if err != nil {
		return nil, err
	}

	if len(products) < 1 {
		return nil, sql.ErrNoRows
	}

	return products, nil
}
