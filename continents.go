package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"sort"
)

func GetContinents() (Continents, error) {
	data, err := os.Open("data/data.json")

	if err != nil {
		return Continents{}, errors.New(err.Error())
	}

	defer data.Close()

	byteValue, _ := ioutil.ReadAll(data)
	var continents Continents

	json.Unmarshal(byteValue, &continents)

	return continents, nil
}

func (cont Continents) Sort() []Continent {
	arr := []float64{}
	cnt := []Continent{}

	for _, i := range cont.Continent {
		arr = append(arr, i.HdiAverage)
	}

	sort.Float64s(arr)

	for _, i := range arr {
		for j, k := range cont.Continent {
			if i == k.HdiAverage {
				cnt = append(cnt, cont.Continent[j])
			}
		}
	}

	return cnt
}
