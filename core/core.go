package core

import (
	"github.com/draganov89/knapsack-problem/model"
)

func Solve(sack *model.Knapsack) []*model.Item {
	valueMatrix, keepMatrix := GetSolvingMatrices(len(sack.Items), sack.Capacity)
	for row := 1; row < len(valueMatrix); row++ {
		item := sack.Items[row-1]
		for col := 1; col < len(valueMatrix[row]); col++ {
			value := valueMatrix[row-1][col]
			if col >= item.Weight {
				// ITEM CAN FIT IN
				tempVal := item.Value
				tempVal += valueMatrix[row-1][col-item.Weight]
				if tempVal > value {
					// TAKE
					valueMatrix[row][col] = tempVal
					keepMatrix[row][col] = 1
					continue
				}
			}
			// DONT TAKE
			valueMatrix[row][col] = value
			keepMatrix[row][col] = 0
		}
	}

	items := []*model.Item{}

	row := len(sack.Items)
	col := sack.Capacity

	for row > 0 {
		value := keepMatrix[row][col]
		for value == 0 {
			row--
			value = keepMatrix[row][col]
		}
		// TAKE ITEM
		itemToAdd := sack.Items[row-1]
		items = append(items, itemToAdd)
		col -= itemToAdd.Weight
		row--
	}

	return items
}

func GetSolvingMatrices(items int, capacity int) ([][]int, [][]byte) {
	capacity++
	items++

	variationMatrix := make([][]int, items)
	for ind := 0; ind < len(variationMatrix); ind++ {
		variationMatrix[ind] = make([]int, capacity)
	}

	keepMatrix := make([][]byte, items)
	for ind := 0; ind < len(keepMatrix); ind++ {
		keepMatrix[ind] = make([]byte, capacity)
	}
	return variationMatrix, keepMatrix
}
