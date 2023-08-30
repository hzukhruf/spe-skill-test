package main

import (
	"fmt"
	"math"
)

func isNarsistic(number int) bool {
	n := number
	arr := []int{}
	count := 0
	for n > 0 {
		arr = append(arr, n%10)
		n = n / 10
		count++
	}

	var total int = 0
	for _, a := range arr {
		total += int(math.Pow(float64(a), float64(count)))
	}

	return total == number
}

func parityOutlier(arrInt []int) string {
	if len(arrInt) < 3 {
		return "Array Input not valid. Minimum length 3"
	}

	arrOdd := []int{}
	arrEven := []int{}
	for _, a := range arrInt {
		if a%2 == 0 {
			arrEven = append(arrEven, a)
		} else {
			arrOdd = append(arrOdd, a)
		}
	}

	if len(arrOdd) == 1 {
		return fmt.Sprintf("%d (the only odd number)", arrOdd[0])
	}
	if len(arrEven) == 1 {
		return fmt.Sprintf("%d (the only even number)", arrEven[0])
	}
	if len(arrOdd) == 0 {
		return "All even integer, no outlier"
	}
	if len(arrEven) == 0 {
		return "All odd integer, no outlier"
	}

	return ""
}

func findNeedle(arr []string, needle string) int {
	for i, a := range arr {
		if a == needle {
			return i
		}
	}
	return 0
}

func blueOcean(arr1 []int, arr2 []int) []int {
	arr3 := []int{}
	for _, a1 := range arr1 {
		for _, a2 := range arr2 {
			if a2 != a1 {
				arr1 = append(arr3, a1)
			}
		}
	}

	return arr1
}

func main() {
	fmt.Println(isNarsistic(153))
	fmt.Println(isNarsistic(111))

	fmt.Println(parityOutlier([]int{2, 4, 0, 100, 4, 11, 2602, 36}))
	fmt.Println(parityOutlier([]int{160, 3, 1719, 19, 11, 13, -21}))
	fmt.Println(parityOutlier([]int{11, 13, 15, 19, 9, 13, -21}))
	fmt.Println(parityOutlier([]int{11, 13}))

	fmt.Println(findNeedle([]string{"red", "blue", "yellow", "black", "grey"}, "blue"))

	fmt.Println(blueOcean([]int{1, 2, 3, 4, 6, 10}, []int{1}))
	fmt.Println(blueOcean([]int{1, 5, 5, 5, 5, 3}, []int{5}))

}
