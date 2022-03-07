package main

type Country struct {
	Country   string  `json:"country"`
	Continent string  `json:"continent"`
	Hdi       float64 `json:"hdi"`
}

type Countries struct {
	Country []Country `json:"countries"`
}

type Continent struct {
	Continent  string  `json:"coninent"`
	Countries  int     `json:"countries"`
	HdiAverage float64 `json:"average_hdi"`
}

type Continents struct {
	Continent []Continent `json:"continents"`
}

type JSON struct {
	Continent []Continent `json:"continents"`
	Country   []Country   `json:"countries"`
}

type Quartiles struct {
	first  float64
	second float64
	third  float64
}
