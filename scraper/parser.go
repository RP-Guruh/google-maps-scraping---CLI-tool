package scraper

import (
	"fmt"
	"gomaps/models"
)

func ParserInit(data models.ScrapeInput) {
	fmt.Println("\nData yang diinput :")
	fmt.Println("=====================")
	fmt.Printf("Place: %s\n", data.Place)
	fmt.Printf("Location: %s\n", data.Location)
	fmt.Printf("Limit: %d\n", data.Limit)
	fmt.Printf("Import File: %s\n", data.ImportFile)
	fmt.Println("=====================")
	ScrapperRun(data.Place, data.Location, data.Limit)
}
