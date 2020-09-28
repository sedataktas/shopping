package main

import (
	"testing"
)

func TestNewCampaign_CreatedSuccessful(t *testing.T) {
	categoryStore := CategoryStore{Categories: nil}
	_, err := categoryStore.NewCategory(1, 0, "shoes", nil)
	if err != nil {
		t.Error(err)
	}

	campaignStore := CampaignStore{Campaigns: nil}
	campaignStore.NewCampaign(1, 10.0, &categoryStore)

	isCampaignExist := false
	for _, c := range campaignStore.Campaigns {
		if c.CategoryID == 1 {
			isCampaignExist = true
		}
	}

	if !isCampaignExist {
		t.Errorf("campaign not added with this id : %d", 1)
	}
}

func TestNewCampaign_WithSubCategories(t *testing.T) {
	categoryStore := CategoryStore{Categories: nil}
	_, err := categoryStore.NewCategory(1, 0, "shoes", nil)
	if err != nil {
		t.Error(err)
	}

	_, err = categoryStore.NewCategory(2, 1, "leather-shoes", nil)
	if err != nil {
		t.Error(err)
	}

	campaignStore := CampaignStore{Campaigns: nil}
	campaignStore.NewCampaign(1, 10.0, &categoryStore)
	campaignStore.NewCampaign(2, 10.0, &categoryStore)
	isCampaignExist := false
	isSubCatCampaignExist := false
	for _, c := range campaignStore.Campaigns {
		if c.CategoryID == 1 {
			isCampaignExist = true
		}
		if c.CategoryID == 2 {
			isSubCatCampaignExist = true
		}
	}

	if !isCampaignExist {
		t.Errorf("campaign not added with this id : %d", 1)
	}

	if !isSubCatCampaignExist {
		t.Errorf("campaign not added with this sub category id : %d", 1)
	}
}

func Test_CalculateCampaign(t *testing.T) {
	categoryStore := CategoryStore{Categories: nil}
	cat, err := categoryStore.NewCategory(1, 0, "shoes", nil)
	if err != nil {
		t.Error(err)
	}
	p := &Product{
		Title:    "adidas",
		Price:    100.0,
		Category: *cat,
	}

	store := CampaignStore{Campaigns: nil}
	store.NewCampaign(1, 10.0, &categoryStore)

	var discount float64
	for _, c := range store.Campaigns {
		if c.CategoryID == 1 {
			c.DiscountPercentage = discount
		}
	}
	CalculateCampaign(p, &store)

	got := p.Price
	discountedPrice := p.Price * (1 - discount/100)
	want := discountedPrice
	if p.Price != discountedPrice {
		t.Errorf("NewCampaign() got = %v, want %v", got, want)
	}
}

func Test_CheckCampaignExist(t *testing.T) {
	cat := Category{
		ID:       1,
		ParentID: 0,
		Title:    "shoes",
	}
	var cats []Category
	cats = append(cats, cat)

	camp := Campaign{
		CategoryID:         1,
		DiscountPercentage: 10,
	}
	var camps []Campaign
	camps = append(camps, camp)

	type args struct {
		campaigns  []Campaign
		categoryID int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"success", args{
			campaigns:  camps,
			categoryID: 1,
		}, true},
		{"not_exist", args{
			campaigns:  camps,
			categoryID: 2,
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckCampaignExist(tt.args.campaigns, tt.args.categoryID); got != tt.want {
				t.Errorf("checkCampaignExist() = %v, want %v", got, tt.want)
			}
		})
	}
}
