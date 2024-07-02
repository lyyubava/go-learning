package repo

import (
	"dou-tech/model"
)

type Repository interface {
	Load(string) ([]model.Item, error)
	Save([]model.Item)
}
