package main

import (
	"dou-tech/model"
	"dou-tech/repo"
	"dou-tech/service"
	"fmt"
)

var i int = 20

func main() {
	r := repo.FileRepo{}

	svc := service.NewMenuService("ua", r)
	m, err := svc.GetItems()
	if err != nil {
		fmt.Println(err)
	}
	svc.PrintMenu(m)

	orderSvc := service.NewOrderService(r)
	orderSvc.Create(m[1])
}

func getMenu() []model.Item {
	m := []model.Item{
		{Id: 1, Name: "Americano", Price: 30.00},
		{Id: 2, Name: "Capuccino", Price: 40.00},
		{Id: 3, Name: "Latte", Price: 35.00},
	}
	item := model.Item{Id: 4, Name: "Tea", Price: 25.00}
	m = append(m, item)
	return m
}
