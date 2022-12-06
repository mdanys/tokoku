package service

import (
	"tokoku/features/user"
	"tokoku/utils/middlewares"

	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	qry user.Repository
}

func New(repo user.Repository) user.Service {
	return &userService{qry: repo}
}

func (us *userService) Create(data user.Core) (user.Core, error) {
	generate, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return user.Core{}, err
	}
	data.Password = string(generate)

	res, err := us.qry.Insert(data)
	if err != nil {
		return user.Core{}, err
	}

	return res, nil
}

func (us *userService) Update(data user.Core, id uint) (user.Core, error) {
	if data.Password != "" {
		generate, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
		if err != nil {
			return user.Core{}, err
		}
		data.Password = string(generate)
	}

	res, err := us.qry.Edit(data, id)
	if err != nil {
		return user.Core{}, err
	}

	return res, nil
}

func (us *userService) Delete(id uint) error {
	err := us.qry.Remove(id)
	if err != nil {
		return err
	}

	return nil
}

func (us *userService) Login(input user.Core) (user.Core, error) {
	res, err := us.qry.Login(input)
	if err != nil {
		return user.Core{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(res.Password), []byte(input.Password))
	if err != nil {
		return user.Core{}, err
	}

	res.Token = middlewares.GenerateToken(res.ID, res.Role)

	return res, nil
}
