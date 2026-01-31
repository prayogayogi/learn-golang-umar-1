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

type CategoryHandler struct {
	service *services.CategoryService
}

func NewCategoryHandler(service *services.CategoryService) *CategoryHandler {
	return &CategoryHandler{service: service}
}

func (handler *CategoryHandler) HandleCategory(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handler.GetAll(w, r)
	case http.MethodPost:
		handler.Create(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (handler *CategoryHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	categories, err := handler.service.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	helpers.JsonResponse(w, http.StatusOK, "Data Berhasil Diambil", categories)
}


func (handler *CategoryHandler) Create(w http.ResponseWriter, r *http.Request) {
	var category models.Category
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = handler.service.Create(&category)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	helpers.JsonResponse(w, http.StatusCreated, "Data Berhasil Ditambahkan", category)
}

func (handler *CategoryHandler) HandleCategoryByID(w http.ResponseWriter, r *http.Request) {
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

func (handler *CategoryHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/category/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		helpers.JsonResponse(w, http.StatusBadRequest, "Invalid category ID", nil)
		return
	}

	category, err := handler.service.GetByID(id)
	if err != nil {
		helpers.JsonResponse(w, http.StatusNotFound, "Data Tidak Ditemukan", nil)
		return
	}
	helpers.JsonResponse(w, http.StatusOK, "Data Berhasil Diambil", category)
}

func (handler *CategoryHandler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/category/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		helpers.JsonResponse(w, http.StatusBadRequest, "Invalid category ID", nil)
		return
	}

	var category models.Category
	err = json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		helpers.JsonResponse(w, http.StatusBadRequest, "Invalid request body", nil)
		return
	}

	category.ID = id
	err = handler.service.Update(&category)
	if err != nil {
		helpers.JsonResponse(w, http.StatusBadRequest, "Invalid request body", nil)
		return
	}

	helpers.JsonResponse(w, http.StatusOK, "Data Berhasil Diupdate", category)
}

// Delete - DELETE /api/category/{id}
func (handler *CategoryHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/category/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		helpers.JsonResponse(w, http.StatusBadRequest, "Invalid category ID", nil)
		return
	}

	err = handler.service.Delete(id)
	if err != nil {
		helpers.JsonResponse(w, http.StatusInternalServerError, "Invalid request body", nil)
		return
	}

	helpers.JsonResponse(w, http.StatusOK, "Data Berhasil Dihapus", nil)
}
