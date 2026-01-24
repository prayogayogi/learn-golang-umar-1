package service

import (
	"encoding/json"
	"kasir-api/helpers"
	"net/http"
	"strconv"
)

type Produk struct{
	ID int `json:"id"`
	Nama string `json:"nama"`
	Harga int `json:"harga"`
	Stok int `json:"stok"`
}

var produks = []Produk{
		{ID: 1, Nama: "Indome Goreng", Harga: 1000, Stok: 10},
		{ID: 1, Nama: "Kecap", Harga: 1022, Stok: 10},
	}

func GetProdukHandler(w http.ResponseWriter, r *http.Request){
	helpers.JsonResponse(w, http.StatusOK, produks)
}

func CreateProdukHandler(w http.ResponseWriter, r *http.Request){
	var produkBaru Produk

	if err := json.NewDecoder(r.Body).Decode(&produkBaru); err != nil{
		http.Error(w, "Format JSON salah", http.StatusBadRequest)
		return
	}

	produkBaru.ID = len(produks) + 1
	produks = append(produks, produkBaru)

	helpers.JsonResponse(w, http.StatusCreated, produkBaru)
}

func GetProdukByIdHandler(w http.ResponseWriter, r *http.Request){
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)

	if err != nil{
		http.Error(w, "ID Harus Berupa Angka", http.StatusBadRequest)
		return
	}

	for _, produk := range produks{
		if produk.ID == id{
			helpers.JsonResponse(w, http.StatusOK, produk)
			return
		}
	}

	http.Error(w, "Produk tidak ditemukan", http.StatusNotFound)
}

func UpdateProdukHandler(w http.ResponseWriter, r *http.Request){
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)

	if err != nil{
		http.Error(w, "ID Harus Berupa Angkat", http.StatusBadRequest)
		return
	}

	var updatedData Produk
	if err := json.NewDecoder(r.Body).Decode(&updatedData); err != nil{
		http.Error(w, "Format Json salah", http.StatusBadRequest)
		return
	}

	for i, produk := range produks{
		if produk.ID == id {
			updatedData.ID = id
			produks[i] = updatedData
			helpers.JsonResponse(w, http.StatusOK, updatedData)
			return
		}
	}

	http.Error(w, "Produk tidak ditemukan", http.StatusNotFound)
}

func DeleteProdukHandler(w http.ResponseWriter, r *http.Request){
	idStr := r.PathValue("id")
	id , err := strconv.Atoi(idStr)
	if err != nil{
		http.Error(w, "ID Harus Berupa Angka", http.StatusBadRequest)
		return
	}

	for i, produk := range produks{
		if produk.ID == id{
			produks = append(produks[:i], produks[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.Error(w, "Produk tidak ditemukan", http.StatusNotFound)
}
