package storage

import (
	"errors"
	"ms-product/internal"
)

var (
	ErrReadFileJSON = errors.New("storage: error read file json")
)

type StorageProduct interface {
	ReadAll() (p map[int]internal.Product, err error)
}
