package controller

import (
	"app/models"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
)

/*
func (c *Controller) Task1() []*models.Order {
    orderData, _ := c.OrderGetList(&models.OrderGetListRequest{})
    orders := orderData.Orders

    sort.Slice(orders, func(i, j int) bool {
        return orders[i].DateTime > orders[j].DateTime
    })

    var result []*models.Order
    for i := range orders {
        result = append(result, *orders[i])
    }

    return result
}
*/
// 1. Order boyicha default holati time sort bolishi kerak. DESC
func (c *Controller) Task1() []models.UserDataS {
	orderData, _ := c.OrderGetList(&models.OrderGetListRequest{})
	orders := orderData.Orders
	productData, _ := c.Strg.Product().GetList(&models.ProductGetListRequest{})
	products := productData.Products
	prodName := make(map[string]string)
	prodPrice := make(map[string]int)
	for _, p := range products {
		prodName[p.Id] = p.Name
		prodPrice[p.Id] = p.Price
	}

	userData := []models.UserDataS{}

	for _, order := range orders {
		for _, item := range order.OrderItems {

			userData = append(userData, models.UserDataS{
				Name:  prodName[item.ProductId],
				Price: prodPrice[item.ProductId],
				Count: item.Count,
				Total: prodPrice[item.ProductId] * item.Count,
				Time:  order.DateTime,
			})

		}
	}

	sort.Slice(userData, func(i, j int) bool {
		return userData[i].Time > userData[j].Time
	})

	return userData
}

// 2. Order Date boyicha filter qoyish kerak

func (c *Controller) Task2(fromDate, toDate string) []models.UserDataS {
	orderData, _ := c.OrderGetList(&models.OrderGetListRequest{})
	orders := orderData.Orders
	productData, _ := c.Strg.Product().GetList(&models.ProductGetListRequest{})
	products := productData.Products
	prodName := make(map[string]string)
	prodPrice := make(map[string]int)
	for _, p := range products {
		prodName[p.Id] = p.Name
		prodPrice[p.Id] = p.Price
	}

	userData := []models.UserDataS{}

	for _, order := range orders {
		if order.DateTime >= fromDate && order.DateTime <= toDate {
			for _, item := range order.OrderItems {
				userData = append(userData, models.UserDataS{
					Name:  prodName[item.ProductId],
					Price: prodPrice[item.ProductId],
					Count: item.Count,
					Total: prodPrice[item.ProductId] * item.Count,
					Time:  order.DateTime[:11],
				})
			}
		}
	}
	for i, data := range userData {
		fmt.Printf("{%s %d %d %s} ", data.Name, data.Price, data.Count, data.Time)
		if i < len(userData)-1 {
			fmt.Print(", ")
		}
	}
	fmt.Println()
	return userData
}

// 3. User history chiqish kerak. Ya'ni sotib olgan mahsulotlari korsatish kerak

func (c *Controller) Task3(id string) []models.UserDataS {
	productData, _ := c.Strg.Product().GetList(&models.ProductGetListRequest{})
	products := productData.Products
	shopcarts, _ := readShopCart("data/shop_cart.json")

	prodName := make(map[string]string)
	prodPrice := make(map[string]int)

	for _, p := range products {
		prodName[p.Id] = p.Name
		prodPrice[p.Id] = p.Price
	}

	userData := []models.UserDataS{}

	for _, s := range shopcarts {
		if s.UserId == id {
			userData = append(userData, models.UserDataS{
				Name:  prodName[s.ProductId],
				Price: prodPrice[s.ProductId],
				Count: s.Count,
				Total: prodPrice[s.ProductId] * s.Count,
				Time:  s.Time,
			})
		}
	}
	return userData

}

// 4. User qancha pul mahsulot sotib olganligi haqida hisobot.

func (c *Controller) Task4(id string) map[string]int {
	productData, _ := c.ProductGetList(&models.ProductGetListRequest{})
	userData, _ := c.UserGetList(&models.UserGetListRequest{})
	users := userData.Users
	products := productData.Products
	shopcarts, _ := readShopCart("data/shop_cart.json")

	prodPrice := make(map[string]int)
	userName := make(map[string]string)
	shopHistory := make(map[string]int)

	for _, p := range products {
		prodPrice[p.Id] = p.Price
	}
	for _, u := range users {
		userName[u.Id] = u.Name
	}

	for _, s := range shopcarts {
		if s.UserId == id {
			shopHistory[userName[s.UserId]] += s.Count * prodPrice[s.ProductId]
		}
	}

	return shopHistory
}

//5. Productlarni Qancha sotilgan boyicha hisobot

func (c *Controller) Task5() map[string]int {
	productData, _ := c.ProductGetList(&models.ProductGetListRequest{})
	products := productData.Products
	shopcarts, _ := readShopCart("data/shop_cart.json")
	prodName := make(map[string]string)
	productCounHistory := make(map[string]int)

	for _, p := range products {
		prodName[p.Id] = p.Name
	}
	for _, v := range shopcarts {
		productCounHistory[prodName[v.ProductId]] += v.Count
	}

	return productCounHistory

}

type ProductCountData struct {
	ProductName string
	Count       int
}

// 6. Top 10 ta sotilayotgan mahsulotlarni royxati.

func (c *Controller) Task6() []ProductCountData {
	productData, _ := c.ProductGetList(&models.ProductGetListRequest{})
	products := productData.Products
	shopcarts, _ := readShopCart("data/shop_cart.json")
	prodName := make(map[string]string)
	productCounHistory := make(map[string]int)

	for _, p := range products {
		prodName[p.Id] = p.Name
	}
	for _, v := range shopcarts {
		productCounHistory[prodName[v.ProductId]] += v.Count
	}

	var countDataSlice []ProductCountData

	for productName, count := range productCounHistory {
		countDataSlice = append(countDataSlice, ProductCountData{
			ProductName: productName,
			Count:       count,
		})
	}
	sort.Slice(countDataSlice, func(i, j int) bool {
		return countDataSlice[i].Count > countDataSlice[j].Count
	})
	return countDataSlice[:10]
}

// 7. TOP 10 ta Eng past sotilayotgan mahsulotlar royxati
func (c *Controller) Task7() []ProductCountData {
	productData, _ := c.ProductGetList(&models.ProductGetListRequest{})
	products := productData.Products
	shopcarts, _ := readShopCart("data/shop_cart.json")
	prodName := make(map[string]string)
	productCounHistory := make(map[string]int)

	for _, p := range products {
		prodName[p.Id] = p.Name
	}
	for _, v := range shopcarts {
		productCounHistory[prodName[v.ProductId]] += v.Count
	}

	var countDataSlice []ProductCountData

	for productName, count := range productCounHistory {
		countDataSlice = append(countDataSlice, ProductCountData{
			ProductName: productName,
			Count:       count,
		})
	}
	sort.Slice(countDataSlice, func(i, j int) bool {
		return countDataSlice[i].Count < countDataSlice[j].Count
	})
	return countDataSlice[:10]
}

// 8. Qaysi Sanada eng kop mahsulot sotilganligi boyicha jadval
// Create a slice to hold the product sales data
type ProductSales struct {
	Day       string
	ProductID string
	Count     int
}

func (c *Controller) Task8() []ProductSales {
	productData, _ := c.ProductGetList(&models.ProductGetListRequest{})
	products := productData.Products
	shopcarts, _ := readShopCart("data/shop_cart.json")

	prodCount := make(map[string]int)
	prodDay := make(map[string]string)
	prodName := make(map[string]string)
	for _, p := range products {
		prodName[p.Id] = p.Name
	}
	for _, v := range shopcarts {
		prodCount[v.ProductId] += v.Count
		prodDay[v.ProductId] = v.Time
	}

	var productSales []ProductSales

	for productID, count := range prodCount {
		productSales = append(productSales, ProductSales{
			Day:       prodDay[productID],
			ProductID: prodName[productID],
			Count:     count,
		})
	}

	// Sort the productSales slice based on the count in descending order
	sort.Slice(productSales, func(i, j int) bool {
		return productSales[i].Count > productSales[j].Count
	})

	return productSales
}

// 9. Qaysi category larda qancha mahsulot sotilgan boyicha jadval
func (c *Controller) Task9() map[string]int {
	productData, _ := c.ProductGetList(&models.ProductGetListRequest{})
	products := productData.Products
	categoryData, _ := c.CategoryGetList(&models.CategoryGetListRequest{})
	categories := categoryData.Categorys
	shopcarts, _ := readShopCart("data/shop_cart.json")
	categoryName := make(map[string]string)
	productCategory := make(map[string]string)
	categoryCounHistory := make(map[string]int)

	for _, c := range categories {
		categoryName[c.Id] = c.Name
	}

	for _, v := range products {
		productCategory[v.Id] = v.CategoryID
	}
	for _, v := range shopcarts {
		categoryCounHistory[categoryName[productCategory[v.ProductId]]] += v.Count
	}

	return categoryCounHistory

}

// 10. Qaysi User eng Active xaridor. Bitta ma'lumot chiqsa yetarli.
type UserCountData struct {
	UserName string
	Count    int
}

func (c *Controller) Task10() []UserCountData {
	userData, _ := c.UserGetList(&models.UserGetListRequest{})
	users := userData.Users
	shopcarts, _ := readShopCart("data/shop_cart.json")
	userName := make(map[string]string)
	UserCounHistory := make(map[string]int)

	for _, u := range users {
		userName[u.Id] = u.Name
	}
	for _, v := range shopcarts {
		UserCounHistory[userName[v.UserId]] += v.Count
	}

	var countDataSlice []UserCountData

	for userName, count := range UserCounHistory {
		countDataSlice = append(countDataSlice, UserCountData{
			UserName: userName,
			Count:    count,
		})
	}
	sort.Slice(countDataSlice, func(i, j int) bool {
		return countDataSlice[i].Count > countDataSlice[j].Count
	})
	return countDataSlice[:2]
}

// 11. Agar User 9 dan kop mahuslot sotib olgan bolsa,
func (c *Controller) Task11() []int {
	orderData, _ := c.OrderGetList(&models.OrderGetListRequest{})
	orders := orderData.Orders
	TotalSlice := make([]int, 0)
	minPrice := math.MaxInt32

	for _, order := range orders {
		for _, item := range order.OrderItems {
			if item.TotalPrice < minPrice {
				minPrice = item.TotalPrice
			}
		}
		Total := order.Sum - minPrice
		TotalSlice = append(TotalSlice, Total)
	}

	return TotalSlice
}

func readShopCart(data string) ([]models.ShopCartS, error) {
	var shopcarts []models.ShopCartS

	d, err := os.ReadFile(data)
	if err != nil {
		log.Printf("Error while Read data: %+v", err)
		return nil, err
	}
	err = json.Unmarshal(d, &shopcarts)
	if err != nil {
		log.Printf("Error while Unmarshal data: %+v", err)
		return nil, err
	}
	return shopcarts, nil
}

func readOrders(data string) ([]models.Order, error) {
	var orders []models.Order

	d, err := os.ReadFile(data)
	if err != nil {
		log.Printf("Error while Read data: %+v", err)
		return nil, err
	}
	err = json.Unmarshal(d, &orders)
	if err != nil {
		log.Printf("Error while Unmarshal data: %+v", err)
		return nil, err
	}
	return orders, nil
}
