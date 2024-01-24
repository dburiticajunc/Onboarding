package repository

import (
	"ms-product/internal"
)

type RepositoryProductMap struct {
	lastId int
	db     map[int]internal.Product
}

func NewRepositoryProductMap(lastId int, p map[int]internal.Product) *RepositoryProductMap {
	return &RepositoryProductMap{
		lastId: lastId,
		db:     p,
	}
}

func (pr *RepositoryProductMap) GetAll() (p map[int]internal.Product, err error) {
	p = make(map[int]internal.Product, len(pr.db))
	for k, v := range pr.db {
		p[k] = v
	}
	return
}
