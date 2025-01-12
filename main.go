package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"net/http"
)

type Solar struct {
	SolarData struct {
		CalculatedConditition struct {
			Bands []Band `xml:"band"`
		} `xml:"calculatedconditions"`
	} `xml:"solardata"`
}

type Band struct {
	Name string `xml:"name,attr"`
	Time string `xml:"time,attr"`
	Condition string `xml:",chardata"`
}

func main() {
	apiEndpoint := "https://www.hamqsl.com/solarxml.php"
	
	resp, err := http.Get(apiEndpoint)
	if err != nil {
		panic(err)
	}
	
	if resp.StatusCode != 200 {
		fmt.Printf("Failed to retrieve XML data from [%s].\n", apiEndpoint)
		os.Exit(1)
	}
	defer resp.Body.Close()

	xmlData, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Failed to read HTTP response body.")
		os.Exit(1)
	}

	solar := Solar{}
	err = xml.Unmarshal(xmlData, &solar)
	if err != nil {
		fmt.Println("Failed to unmarshall solar data, exiting.")
	}

	for _, condition := range solar.SolarData.CalculatedConditition.Bands {
		fmt.Printf("%s: %s (%s)\n", condition.Name, condition.Condition, condition.Time)
	}
}
