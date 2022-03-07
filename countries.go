package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"sort"
)

func GetCountries() (Countries, error) {
	data, err := os.Open("data/data.json")

	if err != nil {
		return Countries{}, errors.New(err.Error())
	}

	defer data.Close()

	byteValue, _ := ioutil.ReadAll(data)
	var countries Countries

	json.Unmarshal(byteValue, &countries)

	return countries, nil
}

func (count Countries) SortByHdi() []Country {
	arr := []float64{}
	cnt := []Country{}

	for _, i := range count.Country {
		arr = append(arr, i.Hdi)
	}

	sort.Float64s(arr)

	for _, i := range arr {
		for j, k := range count.Country {
			if i == k.Hdi {
				cnt = append(cnt, count.Country[j])
			}
		}
	}

	return cnt
}

func (count Countries) SortByCountry() []Country {
	arr := []string{}
	cnt := []Country{}

	for _, i := range count.Country {
		arr = append(arr, i.Country)
	}

	sort.Strings(arr)

	for _, i := range arr {
		for j, k := range count.Country {
			if i == k.Country {
				cnt = append(cnt, count.Country[j])
			}
		}
	}

	return cnt
}

func Search(lstByAlp []Country, country string) (Country, int, error) {
	arr := []string{}
	pos := []float64{}

	var index int = -1

	for _, i := range lstByAlp {
		arr = append(arr, i.Country)
		pos = append(pos, i.Hdi)
	}

	defer func() {
		if err := recover(); err != nil {
			index = -1
		}
	}()

	i := sort.SearchStrings(arr, country)

	if i == -1 {
		return Country{}, index, errors.New("Cannot find the country")
	}

	index = sort.SearchFloat64s(pos, lstByAlp[i].Hdi)

	return lstByAlp[i], len(lstByAlp) - index + 1, nil
}
