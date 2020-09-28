package main

import (
	"bytes"
	"testing"
)

// Print Tests
func Test_printMenu(t *testing.T) {
	printMenu()
	t.Log("Simple Shopping Card CLI Application\n" +
		"---------------------\n" +
		"1) Create a category\n" +
		"2) Create a product\n" +
		"3) Create a campaign\n" +
		"4) Add coupon to card\"\n" +
		"5) Add product to card\n" +
		"6) Calculate card cost\n" +
		"7) Get all categories\n" +
		"8) Get all products\n" +
		"9) Print menu\n" +
		"10) Exit\n---------------------\n")
}

func Test_printProducts(t *testing.T) {
	c := Category{
		ID:       1,
		ParentID: 0,
		Title:    "shoes",
	}
	p := Product{
		Title:    "nike",
		Price:    100,
		Category: c,
	}
	var products []Product
	products = append(products, p)
	printProducts(products)
	t.Log("Products : \n\t" +
		"-- > Title:nike Price:100.000000\n")
}

func Test_printCategories(t *testing.T) {
	c := Category{
		ID:       1,
		ParentID: 0,
		Title:    "shoes",
	}
	var categories []Category
	categories = append(categories, c)
	printCategories(categories)
	t.Log("Categories : \n\t" +
		"-- > Category = ID:1 ParentID:0 Title:shoes")
}

// ReadCategoryInputs Tests
func Test_readCategoryInputs_Success(t *testing.T) {
	var std bytes.Buffer
	_, err := std.WriteString("1 0 title")
	if err != nil {
		t.Fatal(err)
	}

	want := &categoryInputs{
		id:       1,
		parentId: 0,
		title:    "title",
	}
	got, err := readCategoryInputs(&std)
	if err != nil {
		t.Error(err)
	}

	if *got != *want {
		t.Errorf("got %v and want %v is not same", got, want)
	}
	if err == nil {
		t.Log("Success")
	}
}

func Test_readCategoryInputs_Bad_ID(t *testing.T) {
	var std bytes.Buffer
	_, err := std.WriteString("% 0 title")
	if err != nil {
		t.Fatal(err)
	}

	got, err := readCategoryInputs(&std)
	if err == nil {
		t.Error("function should return error")
	}

	if got != nil {
		t.Errorf("got %v and want %v is not same", got, nil)
	}
	if got == nil || err != nil {
		t.Log("success")
	}
}

func Test_readCategoryInputs_Bad_ParentID(t *testing.T) {
	var std bytes.Buffer
	_, err := std.WriteString("1 % title")
	if err != nil {
		t.Fatal(err)
	}

	got, err := readCategoryInputs(&std)
	if err == nil {
		t.Error("function should return error")
	}

	if got != nil {
		t.Errorf("got %v and want %v is not same", got, nil)
	}
	if got == nil || err != nil {
		t.Log("success")
	}
}

func Test_readCategoryInputs_Bad_Title(t *testing.T) {
	var std bytes.Buffer
	_, err := std.WriteString("1 0")
	if err != nil {
		t.Fatal(err)
	}

	got, err := readCategoryInputs(&std)
	if err == nil {
		t.Error("function should return error")
	}

	if got != nil {
		t.Errorf("got %v and want %v is not same", got, nil)
	}
	if got == nil || err != nil {
		t.Log("success")
	}
}

// ReadProductInputs Tests
func Test_readProductInputs_Success(t *testing.T) {
	var std bytes.Buffer
	_, err := std.WriteString("title 100.0 1")
	if err != nil {
		t.Fatal(err)
	}

	category := Category{
		ID:       1,
		ParentID: 0,
		Title:    "shoes",
	}
	var categories []Category
	categories = append(categories, category)
	want := &productInputs{
		title:      "title",
		price:      100.0,
		categoryID: 1,
	}
	got, err := readProductInputs(&std, []Category{})
	if err != nil {
		t.Error(err)
	}

	if *got != *want {
		t.Errorf("got %v and want %v is not same", got, want)
	}
	if err == nil {
		t.Log("Success")
	}
}

func Test_readProductInputs_BadTitle(t *testing.T) {
	var std bytes.Buffer
	_, err := std.WriteString("")
	if err != nil {
		t.Fatal(err)
	}

	category := Category{
		ID:       1,
		ParentID: 0,
		Title:    "shoes",
	}
	var categories []Category
	categories = append(categories, category)

	got, err := readProductInputs(&std, categories)
	if err == nil {
		t.Error("function should return error")
	}

	if got != nil {
		t.Errorf("got %v and want %v is not same", got, nil)
	}
	if got == nil || err != nil {
		t.Log("success")
	}
}

func Test_readProductInputs_BadPrice(t *testing.T) {
	var std bytes.Buffer
	_, err := std.WriteString("1 % title")
	if err != nil {
		t.Fatal(err)
	}

	category := Category{
		ID:       1,
		ParentID: 0,
		Title:    "shoes",
	}
	var categories []Category
	categories = append(categories, category)

	got, err := readProductInputs(&std, categories)
	if err == nil {
		t.Error("function should return error")
	}

	if got != nil {
		t.Errorf("got %v and want %v is not same", got, nil)
	}
	if got == nil || err != nil {
		t.Log("success")
	}
}

func Test_readProductInputs_BadCategory(t *testing.T) {
	var std bytes.Buffer
	_, err := std.WriteString("1 0 %")
	if err != nil {
		t.Fatal(err)
	}

	category := Category{
		ID:       1,
		ParentID: 0,
		Title:    "shoes",
	}
	var categories []Category
	categories = append(categories, category)

	got, err := readProductInputs(&std, categories)
	if err == nil {
		t.Error("function should return error")
	}

	if got != nil {
		t.Errorf("got %v and want %v is not same", got, nil)
	}
	if got == nil || err != nil {
		t.Log("success")
	}
}

// ReadCampaignInputs Tests
func Test_readCampaignInputs_Success(t *testing.T) {
	var std bytes.Buffer
	_, err := std.WriteString("1 10")
	if err != nil {
		t.Fatal(err)
	}

	want := &campaignInputs{
		categoryID: 1,
		discount:   10,
	}

	got, err := readCampaignInputs(&std)
	if err != nil {
		t.Error(err)
	}

	if *got != *want {
		t.Errorf("got %v and want %v is not same", got, want)
	}
	if err == nil {
		t.Log("Success")
	}
}

func Test_readCampaignInputs_Bad_Discount(t *testing.T) {
	var std bytes.Buffer
	_, err := std.WriteString("1")
	if err != nil {
		t.Fatal(err)
	}

	got, err := readCampaignInputs(&std)
	if err == nil {
		t.Error("function should return error")
	}

	if got != nil {
		t.Errorf("got %v and want %v is not same", got, nil)
	}
	if got == nil || err != nil {
		t.Log("success")
	}
}

func Test_readCampaignInputs_Bad_CategoryID(t *testing.T) {
	var std bytes.Buffer
	_, err := std.WriteString("")
	if err != nil {
		t.Fatal(err)
	}

	got, err := readCampaignInputs(&std)
	if err == nil {
		t.Error("function should return error")
	}

	if got != nil {
		t.Errorf("got %v and want %v is not same", got, nil)
	}
	if got == nil || err != nil {
		t.Log("success")
	}
}

// ReadCouponInputs Tests
func Test_readCouponInputs_Success(t *testing.T) {
	var std bytes.Buffer
	_, err := std.WriteString(CouponCode1)
	if err != nil {
		t.Fatal(err)
	}

	got, err := readCouponInput(&std)
	if err != nil {
		t.Error(err)
	}

	want := CouponCode1
	if got != want {
		t.Errorf("got %v and want %v is not same", got, want)
	}
	if err == nil {
		t.Log("Success")
	}
}

func Test_readCouponInputs_Bad_Code(t *testing.T) {
	var std bytes.Buffer
	_, err := std.WriteString("")
	if err != nil {
		t.Fatal(err)
	}

	got, err := readCouponInput(&std)
	if err == nil {
		t.Error("function should return error")
	}

	if got != "" {
		t.Errorf("got %v and want %s is not same", got, "")
	}
	if got == "" || err != nil {
		t.Log("success")
	}
}

// ReadAddToCardInputs Tests
func Test_readAddToCardInputs_Success(t *testing.T) {
	var std bytes.Buffer
	_, err := std.WriteString("nike 2")
	if err != nil {
		t.Fatal(err)
	}

	want := &addToCardInputs{
		productTitle: "nike",
		quantity:     2,
	}
	got, err := readAddToCardInputs(&std)
	if err != nil {
		t.Error(err)
	}

	if *got != *want {
		t.Errorf("got %v and want %v is not same", got, want)
	}
	if err == nil {
		t.Log("Success")
	}
}

func Test_readAddToCardInputs_BadTitle(t *testing.T) {
	var std bytes.Buffer
	_, err := std.WriteString("")
	if err != nil {
		t.Fatal(err)
	}

	got, err := readAddToCardInputs(&std)
	if err == nil {
		t.Error("function should return error")
	}

	if got != nil {
		t.Errorf("got %v and want %v is not same", got, nil)
	}
	if got == nil || err != nil {
		t.Log("success")
	}
}

func Test_readAddToCardInputs_BadQuantity(t *testing.T) {
	var std bytes.Buffer
	_, err := std.WriteString("nike")
	if err != nil {
		t.Fatal(err)
	}

	got, err := readAddToCardInputs(&std)
	if err == nil {
		t.Error("function should return error")
	}

	if got != nil {
		t.Errorf("got %v and want %v is not same", got, nil)
	}
	if got == nil || err != nil {
		t.Log("success")
	}
}

// CreateCategoryFromConsole Tests
func Test_createCategoryFromConsole_Success(t *testing.T) {
	var std bytes.Buffer

	_, err := std.WriteString("1 0 shoes")
	if err != nil {
		t.Fatal(err)
	}

	err = createCategoryFromConsole(&CategoryStore{}, &CampaignStore{}, &std)
	if err != nil {
		t.Error(err)
	}
}

func Test_createCategoryFromConsole_BadInputs(t *testing.T) {
	var std bytes.Buffer
	_, err := std.Write([]byte("sedat 0 shoes"))
	if err != nil {
		t.Fatal(err)
	}

	err = createCategoryFromConsole(&CategoryStore{}, &CampaignStore{}, &std)
	if err == nil {
		t.Error("func should return error")
	}

	if err != nil {
		t.Log("Success")
	}
}

func Test_createCategoryFromConsole_BadCategory(t *testing.T) {
	var std bytes.Buffer
	_, err := std.Write([]byte("1 1 shoes"))
	if err != nil {
		t.Fatal(err)
	}

	err = createCategoryFromConsole(&CategoryStore{}, &CampaignStore{}, &std)
	if err == nil {
		t.Error("func should return error")
	}

	if err != nil {
		t.Log(err)
	}
}

// CreateProductFromConsole Tests
func Test_createProductFromConsole_Success(t *testing.T) {
	var std bytes.Buffer

	_, err := std.WriteString("title 100.0 1")
	if err != nil {
		t.Fatal(err)
	}

	c := Category{
		ID:       1,
		ParentID: 0,
		Title:    "shoes",
	}
	var categories []Category
	categories = append(categories, c)
	catStore := CategoryStore{Categories: categories}

	p := Product{
		Title:    "leather-shoes",
		Price:    100.0,
		Category: c,
	}
	var products []Product
	products = append(products, p)
	prodStore := ProductStore{Products: products}

	err = createProductFromConsole(&prodStore, &catStore, &std)
	if err != nil {
		t.Error(err)
	}
}

func Test_createProductFromConsole_EmptyCategoryError(t *testing.T) {
	var std bytes.Buffer

	_, err := std.WriteString("title 100 1")
	if err != nil {
		t.Fatal(err)
	}

	err = createProductFromConsole(&ProductStore{}, &CategoryStore{}, &std)
	if err == nil {
		t.Error("func should return error")
	}

	if err != nil {
		t.Log("Success")
	}
}

func Test_createProductFromConsole_ProductExistError(t *testing.T) {
	var std bytes.Buffer

	_, err := std.WriteString("nike 100 1")
	if err != nil {
		t.Fatal(err)
	}

	c := Category{
		ID:       1,
		ParentID: 0,
		Title:    "shoes",
	}
	var categories []Category
	categories = append(categories, c)
	catStore := CategoryStore{Categories: categories}

	p := Product{
		Title:    "nike",
		Price:    100.0,
		Category: c,
	}
	var products []Product
	products = append(products, p)
	prodStore := ProductStore{Products: products}

	err = createProductFromConsole(&prodStore, &catStore, &std)
	if err == nil {
		t.Error("func should return error")
	}

	if err != nil {
		t.Log("Success")
	}
}

// CreateCampaignFromConsole Tests
func Test_createCampaignFromConsole_Success(t *testing.T) {
	var std bytes.Buffer

	_, err := std.WriteString("1 10")
	if err != nil {
		t.Fatal(err)
	}

	c := Category{
		ID:       1,
		ParentID: 0,
		Title:    "shoes",
	}
	var categories []Category
	categories = append(categories, c)
	catStore := CategoryStore{Categories: categories}

	camp := Campaign{
		CategoryID:         1,
		DiscountPercentage: 10,
	}

	var camps []Campaign
	camps = append(camps, camp)
	campStore := CampaignStore{Campaigns: camps}
	err = createCampaignFromConsole(&campStore, &catStore, &std)
	if err != nil {
		t.Error(err)
	}
}

func Test_createCampaignFromConsole_BadInputs(t *testing.T) {
	var std bytes.Buffer

	_, err := std.WriteString("1")
	if err != nil {
		t.Fatal(err)
	}

	err = createCampaignFromConsole(&CampaignStore{}, &CategoryStore{}, &std)
	if err == nil {
		t.Error("func should return error")
	}

	if err != nil {
		t.Log("Success")
	}
}

func Test_createCampaignFromConsole_NoCategoryError(t *testing.T) {
	var std bytes.Buffer

	_, err := std.WriteString("1 10")
	if err != nil {
		t.Fatal(err)
	}

	err = createCampaignFromConsole(&CampaignStore{}, &CategoryStore{}, &std)
	if err == nil {
		t.Error("func should return error")
	}

	if err != nil {
		t.Log("Success")
	}
}

func Test_createCampaignFromConsole_CatIDNotFound(t *testing.T) {
	var std bytes.Buffer

	_, err := std.WriteString("2 10")
	if err != nil {
		t.Fatal(err)
	}

	c := Category{
		ID:       1,
		ParentID: 0,
		Title:    "shoes",
	}
	var categories []Category
	categories = append(categories, c)
	catStore := CategoryStore{Categories: categories}

	err = createCampaignFromConsole(&CampaignStore{}, &catStore, &std)
	if err == nil {
		t.Error("func should return error")
	}

	if err != nil {
		t.Log("Success")
	}
}

// CreateCardFromConsole Tests
func Test_createCardFromConsole_Success(t *testing.T) {
	card := createTestCard(MinimumAmount, "")
	err := createCardCostFromConsole(card, &CampaignStore{}, &CouponStore{})
	if err != nil {
		t.Error(err)
	}
}

func Test_createCardFromConsole_WrongCouponCode(t *testing.T) {
	card := createTestCard(MinimumAmount, "test")
	err := createCardCostFromConsole(card, &CampaignStore{}, &CouponStore{})
	if err == nil {
		t.Error("func should return error")
	}

	if err != nil {
		t.Log("Success")
	}
}

// CreateCardFromConsole Tests
func Test_addCouponToCard_Success(t *testing.T) {
	var std bytes.Buffer
	_, err := std.WriteString(CouponCode1)
	if err != nil {
		t.Fatal(err)
	}

	card := createTestCard(MinimumAmount, "")
	var couponStore CouponStore
	couponStore.CreateCoupons()
	err = addCouponToCard(&couponStore, card, &std)
	if err != nil {
		t.Error(err)
	}
}

// CreateCardFromConsole Tests
func Test_addCouponToCard_WrongCoupon(t *testing.T) {
	var std bytes.Buffer
	_, err := std.WriteString("")
	if err != nil {
		t.Fatal(err)
	}

	card := createTestCard(MinimumAmount, "")
	var couponStore CouponStore
	couponStore.CreateCoupons()
	err = addCouponToCard(&couponStore, card, &std)
	if err == nil {
		t.Error("func should return error")
	}

	if err != nil {
		t.Log("Success")
	}
}

func Test_addProductToCard_Success(t *testing.T) {
	var std bytes.Buffer
	_, err := std.WriteString("adidas 2")
	if err != nil {
		t.Fatal(err)
	}

	c := Category{
		ID:       1,
		ParentID: 0,
		Title:    "shoes",
	}
	p := Product{
		Title:    "adidas",
		Price:    100,
		Category: c,
	}

	var prods []Product
	prods = append(prods, p)
	prodStore := ProductStore{Products: prods}

	err = addProductToCard(&prodStore, &Card{}, &std)
	if err != nil {
		t.Error(err)
	} else {
		t.Log("success")
	}
}

func Test_addProductToCad_ExistProd(t *testing.T) {
	var std bytes.Buffer
	_, err := std.WriteString("adidas 2")
	if err != nil {
		t.Fatal(err)
	}

	c := Category{
		ID:       1,
		ParentID: 0,
		Title:    "shoes",
	}
	p := Product{
		Title:    "adidas",
		Price:    100,
		Category: c,
	}

	var prods []Product
	prods = append(prods, p)
	prodStore := ProductStore{Products: prods}

	card := createTestCard(100, "")
	err = addProductToCard(&prodStore, card, &std)
	if err != nil {
		t.Error(err)
	} else {
		t.Log("success")
	}
}

func Test_addProductToCad_WrongProdTitle(t *testing.T) {
	var std bytes.Buffer
	_, err := std.WriteString("nike 2")
	if err != nil {
		t.Fatal(err)
	}

	c := Category{
		ID:       1,
		ParentID: 0,
		Title:    "shoes",
	}
	p := Product{
		Title:    "adidas",
		Price:    100,
		Category: c,
	}

	var prods []Product
	prods = append(prods, p)
	prodStore := ProductStore{Products: prods}

	card := createTestCard(100, "")
	err = addProductToCard(&prodStore, card, &std)
	if err != nil {
		t.Error(err)
	} else {
		t.Log("success")
	}
}

func Test_addProductToCad_BadInputs(t *testing.T) {
	var std bytes.Buffer
	_, err := std.WriteString("")
	if err != nil {
		t.Fatal(err)
	}

	err = addProductToCard(&ProductStore{}, &Card{}, &std)
	if err == nil {
		t.Error("func should return error")
	}

	if err != nil {
		t.Log("Success")
	}
}
