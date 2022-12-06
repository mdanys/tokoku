package repository

import (
	"tokoku/features/cart"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type repoQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) cart.Repository {
	return &repoQuery{db: db}
}

func (rq *repoQuery) Insert(data cart.Core) (cart.Core, error) {
	var dbCek Cart
	err := rq.db.Preload("Product").First(&dbCek, "user_id = ? AND product_id = ?", data.UserID, data.ProductID).Error
	if err != nil {
		return cart.Core{}, nil
	}

	var datas Product
	er := rq.db.First(&datas, "id = ?", data.ProductID).Error
	if er != nil {
		return cart.Core{}, nil
	}

	var cnv Cart = FromCore(data)
	tx := rq.db.Create(&cnv).Error
	if tx != nil {
		log.Error("error on insert cart: ", err.Error())
		return cart.Core{}, nil
	}

	res := ToCore(cnv)
	return res, nil
}

func (rq *repoQuery) Edit(data cart.Core) (cart.Core, error) {
	var cnv Cart = FromCore(data)
	err := rq.db.Where("id = ?", cnv.ID).Updates(&cnv).Error
	if err != nil {
		log.Error("error on edit cart: ", err.Error())
		return cart.Core{}, nil
	}

	res := ToCore(cnv)
	return res, nil
}

func (rq *repoQuery) Remove(id uint) error {
	var data Cart
	err := rq.db.Where("id = ?", id).Delete(&data).Error
	if err != nil {
		log.Error("error on delete cart: ", err.Error())
		return err
	}

	return nil
}

func (rq *repoQuery) Get(id uint) ([]cart.Core, error) {
	var data []Cart
	err := rq.db.Preload("Product").Where("id_user = ?", id).Scan(&data).Error
	if err != nil {
		log.Error("error on get cart: ", err.Error())
		return nil, err
	}

	res := ToCoreArray(data)
	return res, nil
}
