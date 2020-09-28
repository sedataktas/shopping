package main

// Campaign stores campaign infos
type Campaign struct {
	CategoryID         int
	DiscountPercentage float64
}

// CampaignStore stores campaigns
type CampaignStore struct {
	Campaigns []Campaign
}

// NewCampaign creates campaign for category
// If category has sub categories, campaign apply them too
func (store *CampaignStore) NewCampaign(categoryID int, discount float64,
	categoryStore *CategoryStore) {
	// add itself to campaign
	if !CheckCampaignExist(store.Campaigns, categoryID) {
		c := Campaign{
			CategoryID:         categoryID,
			DiscountPercentage: discount,
		}
		store.Campaigns = append(store.Campaigns, c)
	}

	// add sub-categories to campaign
	categories := categoryStore.GetSubCategoriesRecursively(categoryID)
	if categories != nil {
		for _, cat := range *categories {
			c := Campaign{
				CategoryID:         cat.ID,
				DiscountPercentage: discount,
			}
			store.Campaigns = append(store.Campaigns, c)
		}
	}
}

// CheckCampaignExist checks the camping exist by category id
func CheckCampaignExist(campaigns []Campaign, categoryID int) bool {
	for _, camp := range campaigns {
		if camp.CategoryID == categoryID {
			return true
		}
	}
	return false
}

// CalculateCampaign updates product price
// if product has campaign
func CalculateCampaign(p *Product, store *CampaignStore) {
	for _, c := range store.Campaigns {
		if p.Category.ID == c.CategoryID {
			p.Price = p.Price * (1 - c.DiscountPercentage/100)
		}
	}
}
