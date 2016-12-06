//https://www.hackerrank.com/challenges/30-operators

package main

import (
	"fmt"
	"math"
)

func Round(val float64) float64 {
	var newVal float64

	if math.Abs(val-math.Ceil(val)) <= 0.5 {
		newVal = math.Ceil(val)
	} else {
		newVal = math.Floor(val)
	}
	return newVal
}

func main() {
	var mealCost float64
	var tipPercent float64
	var taxPercent float64
	var totalCost float64

	fmt.Scan(&mealCost)
	fmt.Scan(&tipPercent)
	fmt.Scan(&taxPercent)

	totalCost = mealCost + mealCost*tipPercent*0.01 + mealCost*taxPercent*0.01
	fmt.Printf("%s %.0f %s", "The total meal cost is", Round(totalCost), "dollars.")

}
