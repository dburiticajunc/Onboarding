package storage

import (
	"errors"
	"project-go-web/internal/product"
)

var (
	ErrStorageProductTimeLayout = errors.New("storage: time layout invalid")
)

type StorageProduct interface {
	ReadAll() (p map[int]product.Product, err error)
	WriteAll(p map[int]product.Product) (err error)
}
