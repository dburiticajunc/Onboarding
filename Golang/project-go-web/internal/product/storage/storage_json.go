package storage

import (
	"encoding/json"
	"os"
	"project-go-web/internal/product"
)

type StorageProductJSON struct {
	filePath string
}

type ProductAttributeJSON struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

func NewStorageProductJSON(filePath string) *StorageProductJSON {
	return &StorageProductJSON{
		filePath: filePath,
	}
}

func ReadAll(s *StorageProductJSON) (p map[int]*product.Product, err error) {
	dataFile, err := os.Open(s.filePath)
	if err != nil {
		return
	}
	defer func(dataFile *os.File) {
		err := dataFile.Close()
		if err != nil {
			return
		}
	}(dataFile)

	var products []product.Product
	p = make(map[int]*product.Product)
	err = json.NewDecoder(dataFile).Decode(&products)
	if err != nil {
		err = ErrStorageProductTimeLayout
		return
	}

	for _, pro := range products {
		p[int(pro.Id)] = product.NewProduct(pro.Id, pro.Name, pro.Quantity, pro.CodeValue, pro.IsPublished,
			pro.Expiration, pro.Price)
	}
	return
}

func WriteAll(p map[int]product.Product) (err error) {
	err = ErrStorageProductTimeLayout
	return
}
