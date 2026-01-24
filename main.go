package main

import (
	"fmt"
	"kasir-api/service"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// Route Api Produk
	mux.HandleFunc("GET /api/produk", service.GetProdukHandler)
	mux.HandleFunc("GET /api/produk/{id}", service.GetProdukByIdHandler)
	mux.HandleFunc("POST /api/produk/", service.CreateProdukHandler)
	mux.HandleFunc("PUT /api/produk/{id}", service.UpdateProdukHandler)
	mux.HandleFunc("DELETE /api/produk/{id}", service.DeleteProdukHandler)

	// Route Api Category

	mux.HandleFunc("GET /api/category/", service.GetCategoryHandler)
	mux.HandleFunc("GET /api/category/{id}", service.GetCategoryByIDHandler)
	mux.HandleFunc("POST /api/category", service.CreatedCategoryHandler)
	mux.HandleFunc("PUT /api/category/{id}", service.UpdateCategoryHandler)
	mux.HandleFunc("DELETE /api/category/{id}", service.DeleteCategoryHandler)


	fmt.Println("Server Berjalan di port 8080")
	if err := http.ListenAndServe(":8080",  mux); err != nil{
		fmt.Println("Gagal menjalankan server:", err)
	}

}


