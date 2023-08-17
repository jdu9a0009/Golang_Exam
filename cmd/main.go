package main

import (
	"app/config"
	"app/controller"
	"app/models"
	"app/storage/jsondb"
)

func main() {
	cfg := config.Load()
	strg, err := jsondb.NewConnectionJSON(&cfg)
	if err != nil {
		panic("Failed connect to json:" + err.Error())
	}
	con := controller.NewController(&cfg, strg)
	con.OrderPayment(&models.OrderPayment{OrderId: "ff9aa3f6-7dd2-4b2e-9376-93bc47391e82"})
	// fromDate := "2023-06-16"
	// toDate := "2023-06-16"

	// fmt.Println(con.Task2(fromDate, toDate))
}
