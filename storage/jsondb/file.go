package jsondb

import (
	"os"

	"app/config"
	"app/storage"
)

type StoreJSON struct {
	branch   *BranchRepo
	user     *UserRepo
	category *CategoryRepo
	product  *ProductRepo
	order    *OrderRepo
}

func NewConnectionJSON(cfg *config.Config) (storage.StorageI, error) {

	branchFile, err := os.Open(cfg.Path + cfg.BranchFileName)
	if err != nil {
		return nil, err
	}
	userFile, err := os.Open(cfg.Path + cfg.UserFileName)
	if err != nil {
		return nil, err
	}

	categoryFile, err := os.Open(cfg.Path + cfg.CategoryFileName)
	if err != nil {
		return nil, err
	}

	productFile, err := os.Open(cfg.Path + cfg.ProductFileName)
	if err != nil {
		return nil, err
	}

	orderFile, err := os.Open(cfg.Path + cfg.OrderFileName)
	if err != nil {
		return nil, err
	}

	return &StoreJSON{
		branch:   NewBranchRepo(cfg.Path+cfg.BranchFileName, branchFile),
		user:     NewUserRepo(cfg.Path+cfg.UserFileName, userFile),
		category: NewCategoryRepo(cfg.Path+cfg.CategoryFileName, categoryFile),
		product:  NewProductRepo(cfg.Path+cfg.ProductFileName, productFile),
		order:    NewOrderRepo(cfg.Path+cfg.OrderFileName, orderFile),
	}, nil
}

func (o *StoreJSON) Branch() storage.BranchRepoI {
	return o.branch
}
func (u *StoreJSON) User() storage.UserRepoI {
	return u.user
}

func (u *StoreJSON) Category() storage.CategoryRepoI {
	return u.category
}

func (p *StoreJSON) Product() storage.ProductRepoI {
	return p.product
}

func (o *StoreJSON) Order() storage.OrderRepoI {
	return o.order
}
