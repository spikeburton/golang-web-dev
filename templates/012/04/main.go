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

type restaurant struct {
	Name, Address string
	Menu          menu
}

type restaurants []restaurant

func main() {}
