package main

import (
	"flag"
	"fmt"
	"hello/csv"
	"log"
	"strings"
)

func init() {
	log.SetPrefix("Zipcode Converter:")
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func main() {
	filename := flag.String("filename", "zipcode.csv", "CSV file name of Zipcode.")
	word := flag.String("filterword", "", "Filter CSV data by this word.")
	flag.Parse()

	fmt.Printf("file name: %v\n", *filename)
	fmt.Printf("filter word: %v\n", *word)

	c := csv.NewCsv(*filename)
	fmt.Printf("filename in Csv struct: %v\n", c.GetFilename())

	data := c.Read()

	var results [][]string
	for _, line := range data {
		if strings.Contains(line[8], *word) || strings.Contains(line[9], *word) {
			results = append(results, line)
		}
	}

	for _, r := range results {
		fmt.Println(r)
	}
}
