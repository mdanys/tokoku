package service

import "tokoku/features/product"

type productService struct {
	qry product.Repository
}

func New(repo product.Repository) product.Service {
	return &productService{qry: repo}
}

func (ps *productService) ShowAll() ([]product.Core, error) {
	res, err := ps.qry.GetAll()
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (ps *productService) ShowByID(id uint) (product.Core, error) {
	res, err := ps.qry.GetByID(id)
	if err != nil {
		return product.Core{}, err
	}

	return res, nil
}
