package model

type Item struct {
	Id     int `json:"id"`
	Weight int `json:"weight"`
	Value  int `json:"value"`
}

type Knapsack struct {
	Capacity int     `json:"capacity"`
	Items    []*Item `json:"items"`
}
