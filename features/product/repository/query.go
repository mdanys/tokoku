package repository

import (
	"tokoku/features/product"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type repoQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) product.Repository {
	return &repoQuery{db: db}
}

func (rq *repoQuery) GetAll() ([]product.Core, error) {
	var data []Product
	err := rq.db.Find(&data).Error
	if err != nil {
		log.Error("error on get all product: ", err.Error())
		return nil, err
	}

	res := ToCoreArray(data)
	return res, nil
}

func (rq *repoQuery) GetByID(id uint) (product.Core, error) {
	var data Product
	err := rq.db.First(&data, "id = ?", id).Error
	if err != nil {
		log.Error("error on get by id product: ", err.Error())
		return product.Core{}, err
	}

	res := ToCore(data)
	return res, nil
}
