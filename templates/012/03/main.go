package main

type item struct {
	Name, Description string
	price             float64
}

type meal struct {
	Meal  string
	Items []item
}

type menu []meal

func main() {}
