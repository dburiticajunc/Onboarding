package storage

import (
	"encoding/json"
	"fmt"
	"ms-product/internal"
	"os"
)

type StorageProductJSON struct {
	filePath string
}

func NewStorageProductJSON(filePath string) *StorageProductJSON {
	return &StorageProductJSON{
		filePath: filePath,
	}
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

func (s *StorageProductJSON) ReadAll() (p map[int]internal.Product, err error) {

	dataFile, err := os.Open("products.json")
	if err != nil {
		err = fmt.Errorf("error open file: %v", err)
		return
	}
	defer dataFile.Close()

	var auxPro []ProductAttributeJSON
	p = make(map[int]internal.Product, len(auxPro))

	err = json.NewDecoder(dataFile).Decode(&auxPro)
	if err != nil {
		err = fmt.Errorf("error decoding: %v", err)
		return
	}

	for _, pr := range auxPro {
		p[pr.Id] = internal.NewProduct(pr.Id, pr.Name, pr.Quantity, pr.CodeValue, pr.IsPublished, pr.Expiration,
			pr.Price)
	}
	return
}
