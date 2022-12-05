package repository

import (
	"tokoku/features/user"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type repoQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) user.Repository {
	return &repoQuery{db: db}
}

func (rq *repoQuery) Insert(data user.Core) (user.Core, error) {
	var cnv User = FromCore(data)
	err := rq.db.Create(&cnv).Error
	if err != nil {
		log.Error("error on create user: ", err.Error())
		return user.Core{}, err
	}

	res := ToCore(cnv)
	return res, nil
}

func (rq *repoQuery) Edit(data user.Core, id uint) (user.Core, error) {
	var cnv User = FromCore(data)
	err := rq.db.Where("id = ?", id).Updates(&cnv).Error
	if err != nil {
		log.Error("error on edit user: ", err.Error())
		return user.Core{}, err
	}

	er := rq.db.First(&cnv, "id = ?", id).Error
	if er != nil {
		log.Error("error on getting after edit: ", err.Error())
		return user.Core{}, er
	}

	res := ToCore(cnv)
	return res, nil
}

func (rq *repoQuery) Remove(id uint) error {
	var data User
	err := rq.db.Where("id = ?", id).Delete(&data).Error
	if err != nil {
		log.Error("error on delete user: ", err.Error())
		return err
	}

	return nil
}

func (rq *repoQuery) Login(input user.Core) (user.Core, error) {
	var data User = FromCore(input)
	err := rq.db.First(&data, "email = ?", input.Email).Error
	if err != nil {
		log.Error("error get data: ", err.Error())
		return user.Core{}, nil
	}

	res := ToCore(data)
	return res, nil
}
