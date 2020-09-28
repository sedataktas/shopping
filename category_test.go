package main

import (
	"reflect"
	"testing"
)

func Test_CheckIDsValid(t *testing.T) {
	type args struct {
		id       int
		parentID int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"success", args{id: 1, parentID: 0}, false},
		{"negative_numbers", args{id: -1, parentID: 0}, true},
		{"same_numbers", args{id: 1, parentID: 1}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CheckIDsValid(tt.args.id, tt.args.parentID); (err != nil) != tt.wantErr {
				t.Errorf("checkIDsValid() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_CheckIDsExist(t *testing.T) {
	campaignStore := CampaignStore{Campaigns: nil}
	categoryStore := CategoryStore{Categories: nil}

	_, err := categoryStore.NewCategory(1, 0, "shoes", &campaignStore)
	if err != nil {
		t.Error(err)
	}

	type args struct {
		title      string
		id         int
		parentID   int
		categories []Category
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"success_parent_cat",
			args{
				title:      "bags",
				id:         2,
				parentID:   0,
				categories: categoryStore.Categories}, false},
		{"success_sub_cat",
			args{
				title:      "silvers",
				id:         3,
				parentID:   1,
				categories: categoryStore.Categories}, false},
		{"title_exist",
			args{
				title:      "shoes",
				id:         2,
				parentID:   0,
				categories: categoryStore.Categories}, true},
		{"cat_id_exist",
			args{
				title:      "sports",
				id:         1,
				parentID:   0,
				categories: categoryStore.Categories}, true},
		{"parent_id_not_exist",
			args{
				title:      "home",
				id:         4,
				parentID:   7,
				categories: categoryStore.Categories}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CheckIDsExist(tt.args.title, tt.args.id,
				tt.args.parentID, tt.args.categories); (err != nil) != tt.wantErr {
				t.Errorf("checkIDsExist() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCategoryStore_NewCategory(t *testing.T) {
	campaignStore := CampaignStore{Campaigns: nil}
	categoryStore := CategoryStore{Categories: nil}

	cat := Category{
		ID:       1,
		ParentID: 0,
		Title:    "shoes",
	}

	catForExistID := Category{
		ID:       7,
		ParentID: 0,
		Title:    "pencil",
	}
	categoryStore.Categories = append(categoryStore.Categories, catForExistID)

	type args struct {
		id            int
		parentID      int
		title         string
		campaignStore *CampaignStore
	}
	tests := []struct {
		name       string
		categories []Category
		args       args
		want       *Category
		wantErr    bool
	}{
		{"success", categoryStore.Categories,
			args{id: 1, parentID: 0, title: "shoes", campaignStore: &campaignStore},
			&cat, false},

		{"empty_title", categoryStore.Categories,
			args{id: 1, parentID: 0, title: ""}, nil, true},

		{"invalid_id", categoryStore.Categories,
			args{id: 1, parentID: 1, title: "bags"}, nil, true},

		{"id_exist", categoryStore.Categories,
			args{id: 7, parentID: 0, title: "home"}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			store := &CategoryStore{
				Categories: tt.categories,
			}
			got, err := store.NewCategory(tt.args.id,
				tt.args.parentID, tt.args.title, tt.args.campaignStore)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewCategory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCategory() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_AddCategoryToCampaign(t *testing.T) {
	var campaigns []Campaign
	campaign := Campaign{
		CategoryID:         1,
		DiscountPercentage: 10,
	}
	campaigns = append(campaigns, campaign)
	campaignStore := CampaignStore{Campaigns: campaigns}

	type args struct {
		id       int
		parentID int
		store    *CampaignStore
	}
	tests := []struct {
		name string
		args args
	}{
		{"success", args{id: 2, parentID: 1, store: &campaignStore}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AddCategoryToCampaign(tt.args.id, tt.args.parentID, tt.args.store)
		})
	}

	isCampaignExist := false
	for _, c := range campaignStore.Campaigns {
		if c.CategoryID == 2 {
			isCampaignExist = true
		}
	}

	if !isCampaignExist {
		t.Errorf("campaign not added with this id : %d", 2)
	}
}

func Test_GetCategoryByID(t *testing.T) {
	cat := Category{
		ID:       1,
		ParentID: 0,
		Title:    "shoes",
	}
	var cats []Category
	cats = append(cats, cat)

	type args struct {
		id         int
		categories []Category
	}
	tests := []struct {
		name string
		args args
		want *Category
	}{
		{"success", args{
			id:         1,
			categories: cats,
		}, &cat},
		{"nil", args{
			id:         2,
			categories: cats,
		}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetCategoryByID(tt.args.id, tt.args.categories); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getCategoryByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCategoryStore_GetSubCategoriesRecursively(t *testing.T) {
	parent := Category{
		ID:       1,
		ParentID: 0,
		Title:    "shoes",
	}
	subCat1 := Category{
		ID:       2,
		ParentID: 1,
		Title:    "leather-shoes",
	}
	subCat2 := Category{
		ID:       3,
		ParentID: 2,
		Title:    "small-leather-shoes",
	}
	var catsWithSubCat []Category
	catsWithSubCat = append(catsWithSubCat, parent)
	catsWithSubCat = append(catsWithSubCat, subCat1)
	catsWithSubCat = append(catsWithSubCat, subCat2)

	var catsWithNoSubCat []Category
	catsWithNoSubCat = append(catsWithNoSubCat, parent)

	var subCats []Category
	subCats = append(subCats, subCat1)
	subCats = append(subCats, subCat2)

	var emptyCategory []Category

	type fields struct {
		Categories []Category
	}
	type args struct {
		categoryID int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *[]Category
	}{
		{"success", fields{Categories: catsWithSubCat},
			args{categoryID: 1}, &subCats},
		{"no_sub_cat", fields{Categories: catsWithNoSubCat},
			args{categoryID: 1}, &emptyCategory},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			store := &CategoryStore{
				Categories: tt.fields.Categories,
			}
			if got := store.GetSubCategoriesRecursively(tt.args.categoryID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSubCategoriesRecursively() = %v, want %v", got, tt.want)
			}
		})
	}
}
