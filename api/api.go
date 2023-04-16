package api

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/sinulingga23/learn-prometheus/monitoring"
)

type (
	API struct {
		products []Product
	}

	Product struct {
		Id    int64
		Name  string `json:"name"`
		Stock int    `json:"stock"`
	}

	AddProduct struct {
		Name  string `json:"name"`
		Stock int    `json:"stock"`
	}
)

func NewAPI() API {
	return API{products: make([]Product, 0)}
}

func (api *API) AddProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	serviceName := "add_product"

	bytesBody, errReadAll := io.ReadAll(r.Body)
	if errReadAll != nil {

		go monitoring.TotalRequestApi.WithLabelValues(serviceName, strconv.Itoa(http.StatusBadRequest)).Inc()

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	addProduct := AddProduct{}
	if errUnmarshal := json.Unmarshal(bytesBody, &addProduct); errUnmarshal != nil {

		go monitoring.TotalRequestApi.WithLabelValues(serviceName, strconv.Itoa(http.StatusBadRequest)).Inc()

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if addProduct.Name == "" || addProduct.Stock <= 0 {

		go monitoring.TotalRequestApi.WithLabelValues(serviceName, strconv.Itoa(http.StatusBadRequest)).Inc()

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	api.products = append(api.products, Product{
		Id:    time.Now().Unix(),
		Name:  addProduct.Name,
		Stock: addProduct.Stock,
	})

	go monitoring.TotalRequestApi.WithLabelValues(serviceName, strconv.Itoa(http.StatusOK)).Inc()

	w.WriteHeader(http.StatusOK)
	return
}

func (api *API) GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if len(api.products) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	bytesProducts, errMarshal := json.Marshal(api.products)
	if errMarshal != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(bytesProducts)
	return
}

func (api *API) GetProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idStr := chi.URLParam(r, "id")
	id, errAtoi := strconv.Atoi(idStr)
	if errAtoi != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	for _, product := range api.products {
		if id == int(product.Id) {
			bytesProduct, errMarshal := json.Marshal(product)
			if errMarshal != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			w.WriteHeader(http.StatusOK)
			w.Write(bytesProduct)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	return
}
