package internal

type ServiceProduct interface {
	GetAllProduct() (p map[int]Product, err error)
}
