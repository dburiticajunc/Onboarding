package application

import "project-go-web/internal/product/repository"

type ServerChi struct {
	address string
}

func NewServerChi(address string) *ServerChi {
	defaultAddress := ":8080"

	if address != "" {
		defaultAddress = address
	}
	return &ServerChi{
		address: defaultAddress,
	}
}

func (s *ServerChi) Run() error {

	//initialize
	//repository
	defaultId := 0
	rp := repository.NewRepositoryProductMap(make(map[int]repository.ProductAttributeMap, 0), defaultId)

}
