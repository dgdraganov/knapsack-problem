package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/draganov89/knapsack-problem/core"
	"github.com/draganov89/knapsack-problem/model"
)

type Item struct {
	Weight int
	Value  int
}

func main() {
	file, err := os.Open("parameters.json")
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer file.Close()

	sack := &model.Knapsack{}
	err = json.NewDecoder(file).Decode(sack)
	if err != nil {
		log.Fatalf(err.Error())
	}

	resultItems := core.Solve(sack)

	totalValue := 0
	totalWeight := 0
	for _, val := range resultItems {
		totalValue += val.Value
		totalWeight += val.Weight
	}

	fmt.Printf("Total value: \t%v\n", totalValue)
	fmt.Printf("Total weight: \t%v\n", totalWeight)
	fmt.Println("Items included:")
	for i := 0; i < len(resultItems); i++ {
		item := resultItems[i]
		fmt.Printf("\tid: %v\t value: %v\t weight: %v\n", item.Id, item.Value, item.Weight)
	}
}
