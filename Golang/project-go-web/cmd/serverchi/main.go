package main

import (
	"errors"
	"fmt"
	"project-go-web/internal/product/storage"
)

func main() {

	a := storage.NewStorageProductJSON("products.json")

	data, err := storage.ReadAll(a)
	if err != nil {
		ErroConsumingReadAll := errors.New("Error consuming ")
		fmt.Println(ErroConsumingReadAll.Error())
	}

	fmt.Println("Name: ", data[19].Name)

}
