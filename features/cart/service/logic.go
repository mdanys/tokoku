package service

import "tokoku/features/cart"

type cartService struct {
	qry cart.Repository
}

func New(repo cart.Repository) cart.Service {
	return &cartService{qry: repo}
}

func (cs *cartService) Create(data cart.Core) (cart.Core, error) {
	res, err := cs.qry.Insert(data)
	if err != nil {
		return cart.Core{}, err
	}

	return res, nil
}

func (cs *cartService) Update(data cart.Core) (cart.Core, error) {
	res, err := cs.qry.Edit(data)
	if err != nil {
		return cart.Core{}, err
	}

	return res, nil
}

func (cs *cartService) Delete(id uint) error {
	err := cs.qry.Remove(id)
	if err != nil {
		return err
	}

	return nil
}

func (cs *cartService) Show(id uint) ([]cart.Core, error) {
	res, err := cs.qry.Get(id)
	if err != nil {
		return nil, err
	}

	return res, nil
}
