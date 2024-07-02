package service

import (
	"dou-tech/model"
	"dou-tech/repo"
	"fmt"
)

type OrederService struct {
	repo        repo.Repository
	orderNumber int
}

func NewOrderService(r repo.Repository) OrederService {
	return OrederService{
		repo:        r,
		orderNumber: 1,
	}
}

func (s *OrederService) Create(items []model.Item) {
	var total float32
	for _, v := range items {
		total += v.Price
	}
	var order model.Order = model.Order{
		Id:    s.orderNumber,
		Total: total,
		Items: items,
	}
	fmt.Println(order)
	s.orderNumber++
	err := s.repo.Save(order)
	if err != nil {
		fmt.Println(err)
	}
}
