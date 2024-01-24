package handler

import (
	"fmt"
	"ms-product/internal"
	"ms-product/platform/web/response"
	"net/http"
)

type HandlerProductDefault struct {
	sv internal.ServiceProduct
}

func NewHandlerProductDefault(sv internal.ServiceProduct) *HandlerProductDefault {
	return &HandlerProductDefault{sv: sv}
}

func (s *HandlerProductDefault) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, re *http.Request) {

		data, err := s.sv.GetAllProduct()
		if err != nil {
			err = fmt.Errorf("error consuming service: %v", err)
			return
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message":  "seccess to get all products",
			"products": data,
		})
	}
}
