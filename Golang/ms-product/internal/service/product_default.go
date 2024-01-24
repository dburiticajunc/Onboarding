package service

import (
	"fmt"
	"ms-product/internal"
)

type ServiceProductDefault struct {
	rp internal.RepositoryProduct
}

func NewServiceProductDefault(rp internal.RepositoryProduct) *ServiceProductDefault {
	return &ServiceProductDefault{
		rp: rp,
	}
}

func (s *ServiceProductDefault) GetAllProduct() (p map[int]internal.Product, err error) {

	data, err := s.rp.GetAll()
	if err != nil {
		err = fmt.Errorf("error consuming the repo: %v", err)
		return
	}
	p = make(map[int]internal.Product, len(data))

	for k, v := range data {
		p[k] = v
	}
	return
}
