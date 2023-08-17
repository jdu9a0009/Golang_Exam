package controller

import (
	"app/models"
	"encoding/json"
	"sort"

	"log"
	"os"
)

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
