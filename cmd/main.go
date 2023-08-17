package main

import (
	"app/config"
	"app/controller"
	"app/storage/jsondb"
	"fmt"
)

func main() {
	cfg := config.Load()
	strg, err := jsondb.NewConnectionJSON(&cfg)
	if err != nil {
		panic("Failed connect to json:" + err.Error())
	}
	con := controller.NewController(&cfg, strg)
	fmt.Println(con.Task11())
	// fromDate := "2023-06-16"
	// toDate := "2023-06-16"

	// fmt.Println(con.Task2(fromDate, toDate))
}
