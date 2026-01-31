package repositories

import (
	"database/sql"
	"errors"
	"kasir-api/models"
)

type ProductRepository struct{
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository{
	return &ProductRepository{db: db}
}

func (repo *ProductRepository) GetAll() ([]models.Product, error){
	query := "SELECT p.id, p.name, p.price, p.stock, p.category_id, c.name as category_name FROM product p JOIN category c ON p.category_id = c.id"
	rows, err := repo.db.Query(query)
	if err != nil{
		return nil, err
	}

	defer rows.Close()
	products := make([]models.Product, 0)
	for rows.Next(){
		var product models.Product
		err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.Stock, &product.CategoryID, &product.CategoryName)
		if err != nil{
			return nil, err
		}

		products = append(products, product)
	}

	return products, nil
}

func (repo *ProductRepository) Create(product *models.Product) error{
	query := "INSERT INTO product (name, price, stock, category_id) VALUES ($1, $2, $3, $4) RETURNING id"
	err := repo.db.QueryRow(query, product.Name, product.Price, product.Stock, product.CategoryID).Scan(&product.ID)
	return err
}

func (repo *ProductRepository) GetByID(id int) (*models.Product, error){
	query := "SELECT p.id, p.name, p.price, p.stock, p.category_id, c.name as category_name FROM product p JOIN category c ON p.category_id = c.id WHERE p.id = $1"

	var product models.Product
	err := repo.db.QueryRow(query, id).Scan(&product.ID, &product.Name, &product.Price, &product.Stock, &product.CategoryID, &product.CategoryName)
	if err == sql.ErrNoRows {
		return nil, errors.New("Product Tidak ditemukan")
	}

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (repo *ProductRepository) Update(product *models.Product) error{
	query := "UPDATE product  SET name =  $1, price = $2, stock = $3, category_id = $4 WHERE id = $5"
	result, err := repo.db.Exec(query, product.Name, product.Price, product.Stock, product.CategoryID, product.ID)

	if err != nil{
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil{
		return err
	}

	if rows == 0 {
		return errors.New("Produk Tidak ditemukan")
	}
	return nil
}

func (repo *ProductRepository) Delete(id int) error{
	query := "DELETE FROM product WHERE id = $1"
	result, err := repo.db.Exec(query, id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("Produk Tidak ditemukan")
	}
	return err
}

