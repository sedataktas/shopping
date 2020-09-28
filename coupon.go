package main

import "fmt"

const (
	MinimumAmount = 100
	CouponCode1   = "77ec3a60"
	CouponCode2   = "3aa996eb"
)

// Coupon represents coupon object
type Coupon struct {
	Code               string
	DiscountPercentage int
}

// CouponStore stores coupons
type CouponStore struct {
	Coupons []Coupon
}

// CreateCoupons for hard-codded
func (store *CouponStore) CreateCoupons() {
	var coupons []Coupon

	coupon1 := Coupon{
		Code:               CouponCode1,
		DiscountPercentage: 10,
	}
	coupons = append(coupons, coupon1)

	coupon2 := Coupon{
		Code:               CouponCode2,
		DiscountPercentage: 15,
	}
	coupons = append(coupons, coupon2)
	store.Coupons = coupons
}

// GetCouponDiscountPercentage return percentage for coupon code
func (store *CouponStore) GetCouponDiscountPercentage(code string) (int, error) {
	for _, coupon := range store.Coupons {
		if coupon.Code == code {
			return coupon.DiscountPercentage, nil
		}
	}

	return 0,
		fmt.Errorf("there is no coupon with this code : %s", code)
}
