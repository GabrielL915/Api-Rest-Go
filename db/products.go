package db

import (
	"database/sql"

	"github.com/GabrielL915/Api-Rest-Go/models"
)

func (db Database) CreateProduct(product *models.Product) error {
	var id string
	var createdAt string
	query := `INSERT INTO products (name, description, price) VALUES ($1, $2, $3) RETURNING id, created_at;`
	err := db.Conn.QueryRow(query, product.Name, product.Description, product.Price).Scan(&id, &createdAt)
	if err != nil {
		return err
	}
	product.Id = id
	product.CreatedAt = createdAt
	return nil
}

func (db Database) GetAllProducts() (*models.ProductsList, error) {
	list := &models.ProductsList{}
	rows, err := db.Conn.Query("SELECT * FROM products")
	if err != nil {
		return list, err
	}
	for rows.Next() {
		var product models.Product
		err := rows.Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.CreatedAt)
		if err != nil {
			return list, err
		}
		list.Products = append(list.Products, product)
	}
	return list, nil
}

func (db Database) GetProductById(id string) (*models.Product, error) {
	product := &models.Product{}
	query := `SELECT * FROM products WHERE id=$1;`
	row := db.Conn.QueryRow(query, id)
	switch err := row.Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.CreatedAt); err {
	case sql.ErrNoRows:
		return product, ErrNoMatch
	default:
		return product, err
	}
}

func (db Database) UpdateProduct(id string, input models.Product) (models.Product, error) {
	product := models.Product{}
	query := `UPDATE products SET name=$1, description=$2, price=$3 WHERE id=$4 RETURNING *;`
	err := db.Conn.QueryRow(query, input.Name, input.Description, input.Price, id).Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return product, ErrNoMatch
		}
		return product, err
	}
	return product, nil
}

func (db Database) DeleteProduct(id string) error {
	query := `DELETE FROM products WHERE id=$1;`
	_, err := db.Conn.Exec(query, id)
	switch err {
	case sql.ErrNoRows:
		return ErrNoMatch
	default:
		return err
	}
}
