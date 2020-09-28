package main

import (
	"reflect"
	"testing"
)

func TestNewProduct(t *testing.T) {
	categoryStore := CategoryStore{nil}
	category, err := categoryStore.NewCategory(1, 0, "shoes", nil)
	if err != nil {
		t.Error(err)
	}

	prodAdidas := Product{
		Title:    "adidas",
		Price:    100.0,
		Category: *category,
	}

	prodNike := Product{
		Title:    "nike",
		Price:    100.0,
		Category: *category,
	}

	var mockProducts []Product
	mockProducts = append(mockProducts, prodAdidas)
	store := ProductStore{Products: mockProducts}

	tests := []struct {
		name    string
		args    Product
		want    *Product
		wantErr bool
	}{
		{"success", prodNike, &prodNike, false},
		{"exist_by_title", prodAdidas, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := store.NewProduct(tt.args.Title, tt.args.Price, &tt.args.Category)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewProduct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewProduct() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_IsProductExistByTitle(t *testing.T) {
	categoryStore := CategoryStore{nil}
	cat, err := categoryStore.NewCategory(1, 0, "shoes", nil)
	if err != nil {
		t.Error(err)
	}

	prodNike := Product{
		Title:    "nike",
		Price:    100.0,
		Category: *cat,
	}

	var mockProducts []Product
	mockProducts = append(mockProducts, prodNike)
	store := ProductStore{Products: mockProducts}

	type args struct {
		title string
		store ProductStore
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"not_exist", args{title: "adidas", store: store}, false},
		{"exist", args{title: "nike", store: store}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsProductExistByTitle(tt.args.title, tt.args.store); got != tt.want {
				t.Errorf("isProductExistByTitle() = %v, want %v", got, tt.want)
			}
		})
	}
}
