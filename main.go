package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
)

type (
	categoryInputs struct {
		id       int
		parentId int
		title    string
	}
	productInputs struct {
		title      string
		price      float64
		categoryID int
	}
	campaignInputs struct {
		categoryID int
		discount   float64
	}
	addToCardInputs struct {
		productTitle string
		quantity     int
	}
)

func main() {
	printMenu()
	var (
		categoryStore CategoryStore
		productStore  ProductStore
		campaignStore CampaignStore
		couponStore   CouponStore
		items         []CardItem
		card          *Card
	)

	couponStore.CreateCoupons()
	card = NewCard(items, "")

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println(err)
		}

		switch input {
		case 1:
			fmt.Println("Create a category")
			err := createCategoryFromConsole(&categoryStore, &campaignStore, os.Stdin)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("---------------------")
		case 2:
			fmt.Println("Create a product")
			err := createProductFromConsole(&productStore, &categoryStore, os.Stdin)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("---------------------")
		case 3:
			fmt.Println("Create a campaign")
			err := createCampaignFromConsole(&campaignStore, &categoryStore, os.Stdin)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("---------------------")
		case 4:
			fmt.Println("Add coupon to card")
			err := addCouponToCard(&couponStore, card, os.Stdin)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("---------------------")
		case 5:
			err := addProductToCard(&productStore, card, os.Stdin)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("---------------------")
		case 6:
			fmt.Println("Calculate card cost")
			err := createCardCostFromConsole(card, &campaignStore, &couponStore)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("---------------------")
		case 7:
			fmt.Println("Get all categories")
			printCategories(categoryStore.Categories)
			fmt.Println("---------------------")
		case 8:
			fmt.Println("Get all products")
			printProducts(productStore.Products)
			fmt.Println("---------------------")
		case 9:
			printMenu()
		case 10:
			fmt.Println("terminated")
			fmt.Println("---------------------")
			os.Exit(0)
		default:
			fmt.Println("entered wrong number")
		}
	}
	if scanner.Err() != nil {
		panic("an error occurred when read input")
	}
}

// Read
func readCategoryInputs(stdin io.Reader) (*categoryInputs, error) {
	fmt.Print("ID: ")
	var id int
	_, err := fmt.Fscanf(stdin, "%d", &id)
	if err != nil {
		return nil, err
	}

	fmt.Print("Parent ID: ")
	var parentID int
	_, err = fmt.Fscanf(stdin, "%d", &parentID)
	if err != nil {
		return nil, err
	}

	fmt.Print("Title: ")
	var title string
	_, err = fmt.Fscanf(stdin, "%s", &title)
	if err != nil {
		return nil, err
	}

	inputs := &categoryInputs{
		id:       id,
		parentId: parentID,
		title:    title,
	}

	return inputs, nil
}

func readProductInputs(stdin io.Reader, categories []Category) (*productInputs, error) {
	fmt.Print("Title: ")
	var title string
	_, err := fmt.Fscanf(stdin, "%s", &title)
	if err != nil {
		return nil, err
	}

	fmt.Print("Price: ")
	var price float64
	_, err = fmt.Fscanf(stdin, "%f", &price)
	if err != nil {
		return nil, err
	}

	printCategories(categories)
	fmt.Print("Category ID : ")
	var categoryID int
	_, err = fmt.Fscanf(stdin, "%d", &categoryID)
	if err != nil {
		return nil, err
	}

	inputs := productInputs{
		title:      title,
		price:      price,
		categoryID: categoryID,
	}

	return &inputs, err
}

func readCampaignInputs(stdin io.Reader) (*campaignInputs, error) {
	fmt.Print("Category ID : ")
	var categoryID int
	_, err := fmt.Fscanf(stdin, "%d", &categoryID)
	if err != nil {
		return nil, err
	}

	fmt.Print("Discount: ")
	var discount float64
	_, err = fmt.Fscanf(stdin, "%f", &discount)
	if err != nil {
		return nil, err
	}

	inputs := campaignInputs{
		categoryID: categoryID,
		discount:   discount,
	}
	return &inputs, err
}

func readCouponInput(stdin io.Reader) (string, error) {
	fmt.Print("Coupon code : ")
	var code string
	_, err := fmt.Fscanf(stdin, "%s", &code)
	if err != nil {
		return "", err
	}
	return code, err
}

func readAddToCardInputs(stdin io.Reader) (*addToCardInputs, error) {
	fmt.Print("Enter product title that added to card : ")
	var title string
	_, err := fmt.Fscanf(stdin, "%s", &title)
	if err != nil {
		return nil, err
	}

	var quantity int
	fmt.Print("Quantity : ")
	_, err = fmt.Fscanf(stdin, "%d", &quantity)
	if err != nil {
		return nil, err
	}

	input := &addToCardInputs{
		productTitle: title,
		quantity:     quantity,
	}
	return input, err
}

// Create
func createCategoryFromConsole(catStore *CategoryStore,
	campStore *CampaignStore, stdin io.Reader) error {
	fmt.Println("Create a category : Usage : {id,parentID,'title'}")
	fmt.Println("!Hint : if you want to create parent category, you must set parentID 0")

	inputs, err := readCategoryInputs(stdin)
	if err != nil {
		return err
	}

	cat, err := catStore.NewCategory(inputs.id, inputs.parentId, inputs.title, campStore)
	if err != nil {
		return err
	}

	fmt.Printf("New Category = ID:%d ParentID:%d Title:%s\n",
		cat.ID, cat.ParentID, cat.Title)
	return nil
}

func createProductFromConsole(prodStore *ProductStore,
	catStore *CategoryStore, stdin io.Reader) error {
	inputs, err := readProductInputs(stdin, catStore.Categories)
	if err != nil {
		return err
	}

	var cat *Category
	if catStore.Categories != nil {
		cat = GetCategoryByID(inputs.categoryID, catStore.Categories)
		if cat == nil {
			return fmt.Errorf("entered category id is not exist. ID : %d", inputs.categoryID)
		}
	} else {
		return fmt.Errorf("there is no category created")
	}

	prod, err := prodStore.NewProduct(inputs.title, inputs.price, cat)
	if err != nil {
		return err
	}
	fmt.Printf("New Product = Title:%s Price:%f Category Title:%s\n",
		prod.Title, prod.Price, prod.Category.Title)
	return nil
}

func createCampaignFromConsole(campStore *CampaignStore,
	catStore *CategoryStore, stdin io.Reader) error {
	printCategories(catStore.Categories)

	inputs, err := readCampaignInputs(stdin)
	if err != nil {
		return err
	}

	if len(catStore.Categories) == 0 {
		return errors.New("no category created")
	}

	cat := GetCategoryByID(inputs.categoryID, catStore.Categories)
	if cat == nil {
		return fmt.Errorf("there is no category with this id %d", inputs.categoryID)
	}

	campStore.NewCampaign(inputs.categoryID, inputs.discount, catStore)
	fmt.Println("Campaign created with sub categories")
	return nil
}

func createCardCostFromConsole(c *Card, campStore *CampaignStore,
	couponStore *CouponStore) error {
	c.CouponStore = couponStore
	c.CampStore = campStore
	cost, err := CalculateCardCost(c)
	if err != nil {
		return err
	}

	fmt.Printf("Cost:%f\n", cost)
	return nil
}

func addCouponToCard(couponStore *CouponStore, c *Card, stdin io.Reader) error {
	code, err := readCouponInput(stdin)
	if err != nil {
		return err
	}

	isCouponExist := false
	for _, coup := range couponStore.Coupons {
		if coup.Code == code {
			isCouponExist = true
		}
	}

	if !isCouponExist {
		return errors.New("coupon is not exist")
	}

	c.CouponCode = code
	fmt.Println("Coupon added to card")
	return nil
}

func addProductToCard(prodStore *ProductStore, c *Card, stdin io.Reader) error {
	printProducts(prodStore.Products)

	inputs, err := readAddToCardInputs(stdin)
	if err != nil {
		return err
	}

	// check prod exist by title
	var product Product
	for _, p := range prodStore.Products {
		if p.Title == inputs.productTitle {
			product = p
		}
	}

	isProdExistInCard := false
	for i := range c.Items {
		// is added product exist in car
		if c.Items[i].Product.Title == inputs.productTitle {
			// add quantity
			c.Items[i].Quantity += inputs.quantity
			isProdExistInCard = true
		}
	}

	if !isProdExistInCard {
		item := CardItem{
			Product:  product,
			Quantity: inputs.quantity,
		}
		c.Items = append(c.Items, item)
	}

	return nil
}

// Print
func printCategories(categories []Category) {
	fmt.Println("Categories : ")
	for _, cat := range categories {
		fmt.Printf("\t -- > Category = ID:%d ParentID:%d Title:%s\n", cat.ID, cat.ParentID, cat.Title)
	}
}

func printProducts(products []Product) {
	fmt.Println("Products : ")
	for _, p := range products {
		fmt.Printf("\t -- > Title:%s Price:%f\n", p.Title, p.Price)
	}
}

func printMenu() {
	fmt.Println("Simple Shopping Card CLI Application")
	fmt.Println("---------------------")
	fmt.Println("1) Create a category")
	fmt.Println("2) Create a product")
	fmt.Println("3) Create a campaign")
	fmt.Println("4) Add coupon to card")
	fmt.Println("5) Add product to card")
	fmt.Println("6) Calculate card cost")
	fmt.Println("7) Get all categories")
	fmt.Println("8) Get all products")
	fmt.Println("9) Print menu")
	fmt.Println("10) Exit")
	fmt.Println("---------------------")
}
