package main

import (
	"reflect"
	"testing"
)

func TestNewCard(t *testing.T) {
	var categories []Category
	category := Category{
		ID:       1,
		ParentID: 0,
		Title:    "shoes",
	}
	categories = append(categories, category)
	_ = CategoryStore{Categories: categories}

	var products []Product
	product := Product{
		Title:    "adidas",
		Price:    100,
		Category: Category{},
	}
	products = append(products, product)
	_ = ProductStore{Products: products}

	var items []CardItem
	item := CardItem{
		Product:  product,
		Quantity: 2,
	}
	items = append(items, item)

	card := Card{
		Items: items,
	}

	type args struct {
		items      []CardItem
		couponCode string
	}
	tests := []struct {
		name string
		args args
		want *Card
	}{
		{"success", args{items: items, couponCode: ""}, &card},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCard(tt.args.items, tt.args.couponCode); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCard() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_CalculateCardCost_WithCoupon(t *testing.T) {
	couponStore := CouponStore{Coupons: nil}
	couponStore.CreateCoupons()

	card := createTestCard(MinimumAmount, CouponCode1)
	card.CouponStore = &couponStore
	card.CampStore = &CampaignStore{}
	got, err := CalculateCardCost(card)
	if err != nil {
		t.Error(err)

	}
	perc, err := couponStore.GetCouponDiscountPercentage(card.CouponCode)
	if err != nil {
		t.Error(err)
	}

	quantity := float64(card.Items[0].Quantity)
	price := card.Items[0].Product.Price

	want := (quantity * price) * (1 - float64(perc)/100)
	num := GetProductNumberInCard(card)
	if CheckRaiseAddToCardCost(num) {
		want += ExtraCost
	}
	if got != want {
		t.Error(err)
	}

}

func Test_CalculateCardCost_WrongCoupon(t *testing.T) {
	couponStore := CouponStore{Coupons: nil}
	couponStore.CreateCoupons()

	card := createTestCard(MinimumAmount, "test")
	card.CouponStore = &couponStore
	card.CampStore = &CampaignStore{}
	got, err := CalculateCardCost(card)
	if err == nil {
		t.Error("CalculateCardCost func should return err")

	}

	if got != 0 {
		t.Errorf("got and want not equal. got %f, want %d", got, 0)
	}

}

func Test_CalculateCardCost_WithCampaign(t *testing.T) {
	// create campaign
	var campaigns []Campaign
	campaign := Campaign{
		CategoryID:         1,
		DiscountPercentage: 10,
	}
	campaigns = append(campaigns, campaign)
	campStore := CampaignStore{Campaigns: campaigns}

	card := createTestCard(MinimumAmount, "")

	card.CouponStore = &CouponStore{}
	card.CampStore = &campStore
	got, err := CalculateCardCost(card)
	if err != nil {
		t.Error(err)

	}

	quantity := float64(card.Items[0].Quantity)
	price := card.Items[0].Product.Price

	want := (price * (1 - (campaign.DiscountPercentage / 100))) * quantity
	num := GetProductNumberInCard(card)
	if CheckRaiseAddToCardCost(num) {
		want += ExtraCost
	}
	if got != want {
		t.Error(err)
	}

}

func Test_CalculateCardCost_SmallThanMinAmount(t *testing.T) {
	couponStore := CouponStore{Coupons: nil}
	couponStore.CreateCoupons()

	card := createTestCard(1, CouponCode1)

	card.CampStore = &CampaignStore{}
	card.CouponStore = &couponStore
	got, err := CalculateCardCost(card)
	if err == nil {
		t.Error("test should return error")
	}

	if got != 0 {
		t.Error(err)
	}
}

func Test_CheckAddRaiseToCardCost(t *testing.T) {
	type args struct {
		productNumber int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"true", args{productNumber: 5}, true},
		{"false", args{productNumber: 2}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckRaiseAddToCardCost(tt.args.productNumber); got != tt.want {
				t.Errorf("checkAddRaiseToCardCost() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_CalculateProductNumberInCard(t *testing.T) {
	card := createTestCard(100, "")

	totalProductsNumber := 0
	if card != nil {
		for _, i := range card.Items {
			totalProductsNumber += i.Quantity
		}
	}
	type args struct {
		card *Card
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"success", args{card: card}, totalProductsNumber},
		{"nil", args{card: nil}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetProductNumberInCard(tt.args.card); got != tt.want {
				t.Errorf("calculateProductNumberInCard() = %v, want %v", got, tt.want)
			}
		})
	}
}

func createTestCard(prodPrice float64, couponCode string) *Card {
	// create category
	var categories []Category
	category := Category{
		ID:       1,
		ParentID: 0,
		Title:    "shoes",
	}
	categories = append(categories, category)
	_ = CategoryStore{Categories: categories}

	// create product
	var products []Product
	prod := Product{
		Title:    "adidas",
		Price:    prodPrice,
		Category: category,
	}
	products = append(products, prod)
	_ = ProductStore{Products: products}

	// create card items
	var items []CardItem
	item := CardItem{
		Product:  prod,
		Quantity: ProductNumberForRaise,
	}
	items = append(items, item)

	card := Card{
		Items:      items,
		CouponCode: couponCode,
	}
	return &card
}
