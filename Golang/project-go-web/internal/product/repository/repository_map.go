package repository

import "project-go-web/internal/product"

type ProductAttributeMap struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

type RepositoryProductMap struct {
	db     map[int]ProductAttributeMap
	lastId int
}

func NewRepositoryProductMap(db map[int]ProductAttributeMap, lastId int) *RepositoryProductMap {
	return &RepositoryProductMap{
		db:     db,
		lastId: lastId,
	}
}

func (s *RepositoryProductMap) Get() (p []product.Product, err error) {
	p = make([]product.Product, 0, len(s.db))
	for k, v := range s.db {
		// serialization
		p = append(p, *product.NewProduct(k, v.Name, v.Quantity, v.CodeValue, v.IsPublished,
			v.Expiration, v.Price))
	}
	return p, nil
}
