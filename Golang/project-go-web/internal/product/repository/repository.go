package repository

import (
	"errors"
	"project-go-web/internal/product"
)

var (
	ErrRepositoryProductInternal = errors.New("repository: Internal Error")
)

type RepositoryProduct interface {
	Get() (p []product.Product, err error)
}
