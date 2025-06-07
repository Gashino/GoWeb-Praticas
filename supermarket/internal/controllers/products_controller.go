package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"local/go-web/supermarket/internal/domain/entities"
	"local/go-web/supermarket/internal/services"
	"net/http"
	"strconv"
)

var service = services.ProductService{}

func AttachProductsController(router *chi.Mux) {
	fmt.Println("Mounting routes...")
	router.Route("/products", func(r chi.Router) {
		getAllProducts(r, "/")       // GET /products
		getById(r, "/{id}")          // GET /products/{id}
		getGreaterThan(r, "/filter") // GET /products/filter?greaterThan=10
		publishProducts(r, "/")      // POST /products
	})
}

// Handlers
func publishProducts(r chi.Router, url string) {
	r.Post(url, func(writer http.ResponseWriter, request *http.Request) {

		var newProduct entities.Product

		err := json.NewDecoder(request.Body).Decode(&newProduct)

		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
		}

		status, err := service.PostProduct(&services.Data, newProduct)

		fmt.Println(status)

		writer.WriteHeader(http.StatusCreated)
	})
}

func getAllProducts(r chi.Router, url string) {
	r.Get(url, func(writer http.ResponseWriter, request *http.Request) {
		products, err := service.GetAll(services.Data)

		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
		}

		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)

		errEncoding := json.NewEncoder(writer).Encode(products)

		if errEncoding != nil {
			http.Error(writer, errEncoding.Error(), http.StatusBadRequest)
		}

	})
}

func getById(r chi.Router, url string) {
	r.Get(url, func(writer http.ResponseWriter, request *http.Request) {

		idStr := chi.URLParam(request, "id")

		id, err := strconv.Atoi(idStr)

		product, err := service.GetById(services.Data, id)

		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
		}

		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)

		errEncoding := json.NewEncoder(writer).Encode(product)

		if errEncoding != nil {
			http.Error(writer, errEncoding.Error(), http.StatusBadRequest)
		}

	})
}

func getGreaterThan(r chi.Router, url string) {
	r.Get(url, func(writer http.ResponseWriter, request *http.Request) {
		gtStr := request.URL.Query().Get("greaterThan")
		gt, err := strconv.ParseFloat(gtStr, 64)

		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}

		products, err := service.GetGreaterThan(services.Data, gt)

		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)

		errEncoding := json.NewEncoder(writer).Encode(products)

		if errEncoding != nil {
			http.Error(writer, errEncoding.Error(), http.StatusBadRequest)
		}

	})
}
