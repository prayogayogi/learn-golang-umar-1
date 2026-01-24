package service

import (
	"encoding/json"
	"kasir-api/helpers"
	"net/http"
	"strconv"
)

type Category struct{
	ID int `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
}

var categoryes = []Category{
	{ID: 1, Name: "Minuman", Description: "Minuman Bersoda"},
	{ID: 2, Name: "Makanan", Description: "Makanan Berminyak"},
}

func GetCategoryHandler(w http.ResponseWriter, r *http.Request){
	helpers.JsonResponse(w, http.StatusOK, categoryes)
}

func CreatedCategoryHandler(w http.ResponseWriter, r *http.Request){
	var CategoryBaru Category

	if err := json.NewDecoder(r.Body).Decode(&CategoryBaru); err != nil{
		http.Error(w, "Format Json Salah", http.StatusBadRequest)
		return
	}

	CategoryBaru.ID = len(categoryes) + 1
	categoryes = append(categoryes, CategoryBaru)

	helpers.JsonResponse(w, http.StatusCreated, CategoryBaru)
}

func GetCategoryByIDHandler(w http.ResponseWriter, r *http.Request){
	idStr := r.PathValue("id")
	id,  err := strconv.Atoi(idStr)

	if err != nil{
		http.Error(w, "ID Harus Berupa angka", http.StatusBadRequest)
		return
	}

	for _, category := range categoryes{
		if category.ID == id{
			helpers.JsonResponse(w, http.StatusOK, category)
			return
		}
	}

	http.Error(w, "Category tidak ditemukan", http.StatusNotFound)
}

func UpdateCategoryHandler(w http.ResponseWriter, r *http.Request){
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)

	if err != nil{
		http.Error(w, "ID Harus Berupa Angkat", http.StatusBadRequest)
		return
	}

	var updatedData Category
	if err := json.NewDecoder(r.Body).Decode(&updatedData); err != nil{
		http.Error(w, "Format Json Salah", http.StatusBadRequest)
		return
	}

	for i, category := range categoryes{
		if category.ID == id{
			categoryes[i] = updatedData
			helpers.JsonResponse(w, http.StatusOK, updatedData)
			return
		}
	}

	http.Error(w, "Category tidak ditemukan", http.StatusNotFound)
}

func DeleteCategoryHandler(w http.ResponseWriter, r *http.Request){
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)

	if err != nil{
		http.Error(w, "ID Harus Berupa Angkat", http.StatusBadRequest)
		return
	}

	for i, category := range categoryes{
		if category.ID == id{
			categoryes =  append(categoryes[:i], categoryes[i+1:]...)
			return
		}
	}

	http.Error(w, "Category tidak ditemukan", http.StatusNotFound)
}
