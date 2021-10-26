package main

import "fmt"

func CalculateOutputCurrentAtJunction(inputCurrents []int) int {

	if inputCurrents == nil {
		return 0
	} else if len(inputCurrents) == 0 {
		return 0
	} else {
		totalInputCurr := 0
		for _, curr := range inputCurrents {
			totalInputCurr = totalInputCurr + curr
		}
		return totalInputCurr
	}
}
func main() {

	totalInputCurr := []int{5, 3, 9, 2}

	totalOutputCurr := CalculateOutputCurrentAtJunction(totalInputCurr)

	fmt.Println("kirchoffs law total outlput current is ", totalOutputCurr, "A")

}
