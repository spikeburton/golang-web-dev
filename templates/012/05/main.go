package main

import (
	"encoding/csv"
	"log"
	"net/http"
	"os"
	"text/template"
	"time"
)

type csvRow struct {
	Date                   time.Time
	Open, High, Low, Close string
}

type csvTable []csvRow

var tpl *template.Template

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func parse() csvTable {
	var table csvTable

	file, err := os.Open("table.csv")
	if err != nil {
		log.Fatalln("cannot open .csv: ", err)
	}
	defer file.Close()

	lines, err := csv.NewReader(file).ReadAll()
	if err != nil {
		log.Fatalln("cannot read data from .csv: ", err)
	}

	for _, line := range lines[1:] {
		date, _ := time.Parse("2006-01-02", line[0])
		table = append(
			table,
			csvRow{
				Date:  date,
				Open:  line[1],
				High:  line[2],
				Low:   line[3],
				Close: line[4],
			},
		)
	}

	return table
}

func index(res http.ResponseWriter, req *http.Request) {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
	err := tpl.Execute(res, parse())
	if err != nil {
		log.Fatalln("could not execute template: ", err)
	}
}
