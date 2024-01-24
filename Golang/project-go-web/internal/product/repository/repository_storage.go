package repository

import (
	"project-go-web/internal/product"
	"project-go-web/internal/product/storage"
)

type RepositoryProductStorage struct {
	st     storage.StorageProduct
	lastId int
}

func NewRepositoryProductStorage(st storage.StorageProduct, lastId int) *RepositoryProductStorage {
	return &RepositoryProductStorage{
		st:     st,
		lastId: lastId,
	}
}

func (r *RepositoryProductStorage) Get() (p []product.Product, err error) {
	data, e := r.st.ReadAll()
	if e != nil {
		err = ErrRepositoryProductInternal
		return
	}
	p = make([]product.Product, 0, len(data))
	for _, product := range data {
		p = append(p, product)
	}
	return p, nil
}
