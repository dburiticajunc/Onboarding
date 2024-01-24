package handler

import (
	"net/http"
	"project-go-web/internal/product/repository"
	"project-go-web/platform/web/response"
)

type HandlerProducts struct {
	st repository.RepositoryProduct
}

func NewHandlerProducts(st repository.RepositoryProduct) *HandlerProducts {
	return &HandlerProducts{
		st: st,
	}
}

type ProductJSON struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

func (h *HandlerProducts) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, re *http.Request) {
		pr, err := h.st.Get()
		if err != nil {
			return
		}
		data := make([]ProductJSON, 0, len(pr))
		for _, v := range pr {
			data = append(data, ProductJSON{Id: v.Id, Name: v.Name, Quantity: v.Quantity, CodeValue: v.CodeValue, Expiration: v.Expiration, Price: v.Price})
		}

		response.JSON(w, http.StatusOK, map[string]any{"message": "product", "data": data})

	}
}
