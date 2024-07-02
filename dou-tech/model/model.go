package model

type Item struct {
	Id    int
	Name  string
	Price float32
}

type Order struct {
	Total float32
	Items []Item
	Id    int
}
