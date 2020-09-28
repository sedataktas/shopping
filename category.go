package main

import (
	"errors"
	"fmt"
)

// Category represents product's category
type Category struct {
	ID       int
	ParentID int
	Title    string
}

// CategoryStore stores categories
type CategoryStore struct {
	Categories []Category
}

// NewCategory creates a category
func (store *CategoryStore) NewCategory(id, parentID int,
	title string, campaignStore *CampaignStore) (*Category, error) {
	if title == "" {
		return nil, errors.New("category title can not be empty")
	}

	err := CheckIDsValid(id, parentID)
	if err != nil {
		return nil, err
	}

	err = CheckIDsExist(title, id, parentID, store.Categories)
	if err != nil {
		return nil, err
	}

	newCategory := Category{
		ID:       id,
		ParentID: parentID,
		Title:    title,
	}

	store.Categories = append(store.Categories, newCategory)

	if campaignStore != nil {
		AddCategoryToCampaign(id, parentID, campaignStore)
	}
	return &newCategory, nil
}

// GetSubCategoriesRecursively returns parent category's sub categories
func (store *CategoryStore) GetSubCategoriesRecursively(categoryID int) *[]Category {
	var categories []Category
	for _, cat := range store.Categories {
		if cat.ParentID == categoryID {
			categories = append(categories, cat)
			cats := store.GetSubCategoriesRecursively(cat.ID)
			if cats != nil {
				categories = append(categories, *cats...)
			}
		}
	}
	return &categories
}

// CheckIDsValid return error if ids negative or
// equals each other, otherwise true
func CheckIDsValid(id, parentID int) error {
	if id < 0 || parentID < 0 {
		return errors.New("category id or parent id can not be negative number")
	}

	if id == parentID {
		return errors.New("category id and parent id can not be equal")
	}
	return nil
}

// CheckIDsExist checks id or parent id exist
func CheckIDsExist(title string, id, parentID int, categories []Category) error {
	parentIDExisted := false
	for _, cat := range categories {
		if cat.Title == title {
			return fmt.Errorf("category already exists in this title : %s", title)
		}

		if cat.ID == id {
			return fmt.Errorf("category id exist : %d", id)
		}

		if parentID != 0 && cat.ID == parentID {
			parentIDExisted = true
		}
	}

	if parentID != 0 && !parentIDExisted {
		return fmt.Errorf("entered parent id is not exist. Parent id : %d", parentID)
	}
	return nil
}

// AddCategoryToCampaign add new campaign for category
func AddCategoryToCampaign(id, parentID int, store *CampaignStore) {
	isParentIDExist := false
	var percentage float64
	for _, c := range store.Campaigns {
		if c.CategoryID == id {
			return
		} else {
			if c.CategoryID == parentID {
				percentage = c.DiscountPercentage
				isParentIDExist = true
			}
		}
	}

	if isParentIDExist {
		c := Campaign{
			CategoryID:         id,
			DiscountPercentage: percentage,
		}

		store.Campaigns = append(store.Campaigns, c)
	}
}

// GetCategoryByID returns category by category id
func GetCategoryByID(id int, categories []Category) *Category {
	for _, cat := range categories {
		if cat.ID == id {
			return &cat
		}
	}
	return nil
}
