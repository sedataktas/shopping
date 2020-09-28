package main

import "testing"

func TestGetCouponDiscountPercentage(t *testing.T) {
	store := CouponStore{Coupons: nil}
	store.CreateCoupons()

	type args struct {
		code string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{"success", args{code: CouponCode1}, 10, false},
		{"not_exist", args{code: "test"}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := store.GetCouponDiscountPercentage(tt.args.code)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCouponDiscountPercentage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetCouponDiscountPercentage() got = %v, want %v", got, tt.want)
			}
		})
	}
}
