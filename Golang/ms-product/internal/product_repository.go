package internal

type RepositoryProduct interface {
	GetAll() (p map[int]Product, err error)
}
