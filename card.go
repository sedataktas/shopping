package main

import (
	"fmt"
)

const (
	// ExtraCost represents if the number in the basket is greater than
	//the ProductNumberForRaise it indicates the amount to add
	ExtraCost = 10

	// ProductNumberForRaise represents the minimum required to add extra costs
	ProductNumberForRaise = 5
)

// Card represents card object
type Card struct {
	Items       []CardItem
	CampStore   *CampaignStore
	CouponStore *CouponStore
	CouponCode  string
}

// CardItem stores product and quantity infos
type CardItem struct {
	Product  Product
	Quantity int
}

// NewCard creates new card
func NewCard(items []CardItem, couponCode string) *Card {
	return &Card{
		Items:      items,
		CouponCode: couponCode,
	}
}

// CalculateCardCost take card object and return card's cost
func CalculateCardCost(c *Card) (float64, error) {
	var cost float64
	for _, info := range c.Items {
		CalculateCampaign(&info.Product, c.CampStore)
		cost = cost + (float64(info.Quantity) * info.Product.Price)
	}

	if c.CouponCode != "" {
		if cost < MinimumAmount {
			return 0,
				fmt.Errorf("the total card amount is not sufficient for the coupon to apply."+
					"Minimum total card amount should be : %d", MinimumAmount)
		}

		percentage, err := c.CouponStore.GetCouponDiscountPercentage(c.CouponCode)
		if err != nil {
			return 0, err
		}

		cost = cost * (1 - float64(percentage)/100)
	}

	prodNumber := GetProductNumberInCard(c)
	if CheckRaiseAddToCardCost(prodNumber) {
		cost += ExtraCost
	}

	return cost, nil
}

// GetProductNumberInCard returns product number in car
func GetProductNumberInCard(card *Card) int {
	totalProductsNumber := 0
	if card != nil {
		for _, i := range card.Items {
			totalProductsNumber += i.Quantity
		}
	}

	return totalProductsNumber
}

// CheckRaiseAddToCardCost return true if product number
// bigger than ProductNumberForRaise, otherwise return false
func CheckRaiseAddToCardCost(productNumber int) bool {
	return productNumber >= ProductNumberForRaise
}
