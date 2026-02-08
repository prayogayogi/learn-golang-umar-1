package handlers

import (
	"encoding/json"
	"kasir-api/helpers"
	"kasir-api/models"
	"kasir-api/services"
	"net/http"
	"strconv"
	"strings"
)

type ProductHandler struct {
	service *services.ProductService
}

func NewProductHandler(service *services.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (handler *ProductHandler) HandleProducts(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handler.GetAll(w, r)
	case http.MethodPost:
		handler.Create(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (handler *ProductHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	filter := models.ProdukFilter{
		Name: query.Get("name"),
		Active: query.Get("active"),
	}
	products, err := handler.service.GetAll(filter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	helpers.JsonResponse(w, http.StatusOK, "Data Berhasil Diambil", products)
}


func (handler *ProductHandler) Create(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = handler.service.Create(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	helpers.JsonResponse(w, http.StatusCreated, "Data Berhasil Ditambahkan", product)
}

func (handler *ProductHandler) HandleProductByID(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handler.GetByID(w, r)
	case http.MethodPut:
		handler.Update(w, r)
	case http.MethodDelete:
		handler.Delete(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (handler *ProductHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		helpers.JsonResponse(w, http.StatusBadRequest, "Invalid product ID", nil)
		return
	}

	product, err := handler.service.GetByID(id)
	if err != nil {
		helpers.JsonResponse(w, http.StatusNotFound, "Data Tidak Ditemukan", nil)
		return
	}
	helpers.JsonResponse(w, http.StatusOK, "Data Berhasil Diambil", product)
}

func (handler *ProductHandler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		helpers.JsonResponse(w, http.StatusBadRequest, "Invalid product ID", nil)
		return
	}

	var product models.Product
	err = json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		helpers.JsonResponse(w, http.StatusBadRequest, "Invalid request body", nil)
		return
	}

	product.ID = id
	err = handler.service.Update(&product)
	if err != nil {
		helpers.JsonResponse(w, http.StatusBadRequest, "Invalid request body", nil)
		return
	}

	helpers.JsonResponse(w, http.StatusOK, "Data Berhasil Diupdate", product)
}

// Delete - DELETE /api/produk/{id}
func (handler *ProductHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		helpers.JsonResponse(w, http.StatusBadRequest, "Invalid product ID", nil)
		return
	}

	err = handler.service.Delete(id)
	if err != nil {
		helpers.JsonResponse(w, http.StatusInternalServerError, "Invalid request body", nil)
		return
	}

	helpers.JsonResponse(w, http.StatusOK, "Data Berhasil Dihapus", nil)
}
