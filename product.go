package main

import "fmt"

// Product represents product object
type Product struct {
	Title    string
	Price    float64
	Category Category
}

// ProductStore stores products
type ProductStore struct {
	Products []Product
}

// NewProduct creates a new product
func (store *ProductStore) NewProduct(title string, price float64, cat *Category) (*Product, error) {
	if IsProductExistByTitle(title, *store) {
		return nil,
			fmt.Errorf("product exist with this title: %s", title)
	}

	p := Product{
		Title:    title,
		Price:    price,
		Category: *cat,
	}

	store.Products = append(store.Products, p)
	return &p, nil
}

// IsProductExistByTitle checks if product exist by product title
func IsProductExistByTitle(title string, store ProductStore) bool {
	for _, p := range store.Products {
		if p.Title == title {
			return true
		}
	}
	return false
}
