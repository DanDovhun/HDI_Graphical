package main

import (
	"math"

	statistics "github.com/DanDovhun/Statistics"
)

func GetClosest(arr []float64, num float64) int {
	for i, j := range arr {
		if arr[i+1] > num {
			var a, b float64 = j - num, -1 * (j - num)

			if a > b {
				return i + 1
			}

			return i
		}
	}

	return 0
}

func GetQuartiles(lst []Country) Quartiles {
	arr := []float64{}

	for _, i := range lst {
		arr = append(arr, i.Hdi)
	}

	first, err := statistics.Quartile(arr, 1)

	if err != nil {
		panic(err)
	}

	second, err2 := statistics.Quartile(arr, 2)

	if err2 != nil {
		panic(err)
	}

	third, err3 := statistics.Quartile(arr, 3)

	if err3 != nil {
		panic(err)
	}

	return Quartiles{first, second, third}
}

func GetRealQuartiles(lst []Country) (int, int, int) {
	arr := []float64{}

	for _, i := range lst {
		arr = append(arr, i.Hdi)
	}

	quartiles := GetQuartiles(lst)

	firstIndex := GetClosest(arr, quartiles.first)
	secondIndex := GetClosest(arr, quartiles.second)
	thirdIndex := GetClosest(arr, quartiles.third)

	return firstIndex, secondIndex, thirdIndex
}

func round(num float64, to int) float64 {
	return math.Round(num*math.Pow10(to)) / math.Pow10(to)
}
