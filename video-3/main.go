package main

import (
	"fmt"
	"math"
)

func main() {
	res := MAE()
	fmt.Println(res)
}

func MAE() float64 {
	observedValues := []int{10, 12, 15, 19, 22, 25}
	pValues := []int{9, 11, 19, 18, 26, 19}
	// errorMargin :=  []int {1,2,3,4,5,6}
	count := 6
	var _mae float64
	for i := 0; i < count; i++ {
		_mae += math.Abs(float64(observedValues[i]) - float64(pValues[i]))
	}

	return _mae / float64(count)

}
