package application

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"local/go-web/supermarket/internal/domain/entities"
	"local/go-web/supermarket/internal/services"
	"net/http"
	"strconv"
)

var service = services.StartProdService()

func AttachProductsController(router *chi.Mux) {
	fmt.Println("Mounting routes...")
	router.Route("/products", func(r chi.Router) {
		getAllProducts(r, "/")       // GET /products
		getById(r, "/{id}")          // GET /products/{id}
		getGreaterThan(r, "/filter") // GET /products/filter?greaterThan=10
		publishProducts(r, "/")      // POST /products
		deleteProduct(r, "/{id}")    //DELETE /products/{id}
		//updatePatchProduct(r, "/{id}") //PATCH /products/{id}
		updatePutProduct(r, "/{id}") // PUT /products/{id}
	})
}

// Handlers

func deleteProduct(r chi.Router, url string) {
	r.Delete(url, func(writer http.ResponseWriter, request *http.Request) {

		idStr := chi.URLParam(request, "id")

		id, err := strconv.Atoi(idStr)

		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		errDeleting := service.DeleteProduct(id)

		if errDeleting != nil {
			http.Error(writer, errDeleting.Error(), http.StatusNotFound)
			return
		}

		writer.WriteHeader(http.StatusOK)
	})
}

func updatePutProduct(r chi.Router, url string) {
	r.Put(url, func(writer http.ResponseWriter, request *http.Request) {
		var newProduct entities.Product

		err := json.NewDecoder(request.Body).Decode(&newProduct)

		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		idStr := chi.URLParam(request, "id")

		id, err := strconv.Atoi(idStr)

		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		newProduct.Id = id

		validationErr := service.ValidateProduct(newProduct)

		if validationErr != nil {
			http.Error(writer, validationErr.Error(), http.StatusBadRequest)
			return
		}

		errPutting := service.UpdateProduct(newProduct)

		if errPutting != nil {
			http.Error(writer, errPutting.Error(), http.StatusNotFound)
			return
		}
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)

		errJson := json.NewEncoder(writer).Encode(&newProduct)

		if errJson != nil {
			http.Error(writer, errJson.Error(), http.StatusInternalServerError)
			return
		}

	})
}

//func updatePutProduct(r chi.Router, url string) {
//	r.Put(url, func(writer http.ResponseWriter, request *http.Request) {
//
//	})
//}

func publishProducts(r chi.Router, url string) {
	r.Post(url, func(writer http.ResponseWriter, request *http.Request) {

		var newProduct entities.Product

		err := json.NewDecoder(request.Body).Decode(&newProduct)

		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}

		validationErr := service.ValidateProduct(newProduct)

		if validationErr != nil {
			http.Error(writer, validationErr.Error(), http.StatusBadRequest)
			return
		}

		status, _ := service.PostProduct(newProduct)

		if !status {
			http.Error(writer, "undefined error", http.StatusInternalServerError)
			return
		}

		writer.WriteHeader(http.StatusCreated)
	})
}

func getAllProducts(r chi.Router, url string) {
	r.Get(url, func(writer http.ResponseWriter, request *http.Request) {
		products, err := service.GetAll()

		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
		}

		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)

		errEncoding := json.NewEncoder(writer).Encode(products)

		if errEncoding != nil {
			http.Error(writer, errEncoding.Error(), http.StatusBadRequest)
			return
		}

	})
}

func getById(r chi.Router, url string) {
	r.Get(url, func(writer http.ResponseWriter, request *http.Request) {

		idStr := chi.URLParam(request, "id")

		id, err := strconv.Atoi(idStr)

		product, err := service.GetById(id)

		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}

		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)

		errEncoding := json.NewEncoder(writer).Encode(product)

		if errEncoding != nil {
			http.Error(writer, errEncoding.Error(), http.StatusBadRequest)
			return
		}

	})
}

func getGreaterThan(r chi.Router, url string) {
	r.Get(url, func(writer http.ResponseWriter, request *http.Request) {
		gtStr := request.URL.Query().Get("greaterThan")
		gt, err := strconv.ParseFloat(gtStr, 64)

		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		products, err := service.GetGreaterThan(gt)

		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)

		errEncoding := json.NewEncoder(writer).Encode(products)

		if errEncoding != nil {
			http.Error(writer, errEncoding.Error(), http.StatusBadRequest)
			return
		}

	})
}
