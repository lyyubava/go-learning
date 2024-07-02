package repo

import (
	"dou-tech/model"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type FileRepo struct {
}

func (r FileRepo) Load(l string) ([]model.Item, error) {
	filename := fmt.Sprintf("data/%v.json", l)
	jsonFile, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()
	byteValues, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	var items []model.Item
	err = json.Unmarshal(byteValues, &items)
	if err != nil {
		return nil, err
	}

	return items, nil
}

func (r FileRepo) Save(order model.Order) error {
	file, err := os.Create(fmt.Sprintf("data/order.%v.json", order.Id))
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	err = encoder.Encode(order)
	if err != nil {
		return err
	}
	return nil
}
