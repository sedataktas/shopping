# Shopping Card Application

Shopping card app is a simple cli application. With this app,
you can create a card and calculate card cost by product number,coupon
and campaign.

Here is the app menu : 

1) Create a category
2) Create a product
3) Create a campaign
4) Add a coupon to card
5) Add the product to card
6) Calculate card cost
7) Get all categories
8) Get all products
9) Print menu
10) Exit

### Rules and Infos

- When app starts, card object creates empty
- When app starts, there is no category or product created
- You should create a product to add to card
- To create a product, you should create a category (because product has a category)
- Categories may have subcategories
- **To create a parent category, set parent id 0**
- Campaigns are applicable to a product a category, 
so a category must be created to make a campaign
- **The campaign can be applied to only 1 category 
and also applies to subcategories of this campaign category.**
- **Coupons generated hard coded**
- If you want to use coupon, use this coupon code : **77ec3a60**
- This coupon applies a 10 percent discount 
if the amount of the basket is greater than the minimum
- **The minimum card amount determined as 100 (hard codded)**
- **Delivery cost based on number of products. If number of products greater than
or equal 5, extra cost is added total cost.**
- **Extra cost is determined 10(hard codded)**


### Example Scenario


**Create a category**\
1\
Create a category\
Create a category : Usage : {id,parentID,'title'}\
!Hint : if you want to create parent category, you must set parentID 0\
ID: 1\
Parent ID: 0\
Title: shoes\
New Category = ID:1 ParentID:0 Title:shoes


**Create a product**\
2\
Create a product\
Title: nike\
Price: 100\
Categories : \
         -- > Category = ID:1 ParentID:0 Title:shoes\
Category ID : 1\
New Product = Title:nike Price:100.000000 Category Title:shoes\


**Create a campaign**\
3\
Create a campaign\
Categories : \
         -- > Category = ID:1 ParentID:0 Title:shoes\
Category ID : 1\
Discount: 10\
Campaign created and sub categories\


**Add a coupon to card**\
4\
Add coupon to card\
Coupon code : 77ec3a60\
Coupon added to card\


**Add the product to card**\
5\
Products : \
         -- > Title:nike Price:100.000000\
Enter product title that added to card : nike\
Quantity : 2\


**Calculate cost**\
6\
Calculate card cost\
Cost:162.000000\


### Next Steps

- Campaign end date should be added
- When applies a campaign, discount is given to the real price of the product.
The product must have a discounted price, 
and the prices should be withdrawn from there
- Delete and update options for categories,campaigns,products and coupons
- Delete to product from card
- Decrease product quantity in card
- If deletes all category's product deletes campaign
- Add expire time or used info for coupon


