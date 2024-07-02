package service

import (
	"dou-tech/model"
	"dou-tech/repo"
	"fmt"
)

type MenuService struct {
	repo     repo.Repository
	language string
}

func NewMenuService(l string, repo repo.Repository) MenuService {
	return MenuService{
		repo:     repo,
		language: l,
	}
}

func (s *MenuService) GetItems() ([]model.Item, error) {
	r := repo.FileRepo{}
	return r.Load(s.language)
}

func (s MenuService) PrintMenu(m []model.Item) {
	for _, v := range m {
		fmt.Printf("%v - %.2f\n", v.Name, v.Price)
	}
}
