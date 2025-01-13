package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"net/http"
)

const (
	clear = "\033[0m"
	BrightRed = "\033[31;1m"
	BrightGreen = "\033[32;1m"
	BrightYellow = "\033[33;1m"
	BrightWhite = "\033[37;1m"
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

	bands, err := fetchBandData()
	if err != nil {
		fmt.Println("Failed to retrieve band data. %w", err)
		os.Exit(1)
	}

	dayBands, nightBands := separateBandsPerTime(bands)

	renderBandData(dayBands, nightBands)
}

func renderBandData(day, night []Band) {

	fmt.Println(colorText(BrightWhite, "\tHF Propagation\n"))
	fmt.Println(colorText(BrightWhite, " Band\t\t Day\t Night"))
	fmt.Println("-------------------------------")

	for i := 0; i < len(day); i++ {
		fmt.Printf(" %s\t %s\t %s\n",
			colorText(BrightWhite, day[i].Name),
			colorCondititon(day[i].Condition),
			colorCondititon(night[i].Condition))
	}
}
func colorText(color, text string) string {
	return fmt.Sprintf("%s%s%s", color, text, clear)
}

func colorCondititon(condition string) string {

	const (
		GOOD = "Good"
		FAIR = "Fair"
		POOR = "Poor"
		CLOSED = "Closed"
	)

	switch (condition) {
	case GOOD:
		return fmt.Sprintf("%s%s%s", BrightGreen, condition, clear)
	case FAIR:
		return fmt.Sprintf("%s%s%s", BrightYellow, condition, clear)
	case POOR:
		return fmt.Sprintf("%s%s%s", BrightRed, condition, clear)
	case CLOSED:
		return fmt.Sprintf("%s%s%s", BrightRed, condition, clear)
	default:
		return condition
	}
}

func fetchBandData() ([]Band, error) {
	apiEndpoint := "https://www.hamqsl.com/solarxml.php"
	
	resp, err := http.Get(apiEndpoint)
	if err != nil {
		panic(err)
	}
	
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Failed to retrieve XML data from [%s]: %w", apiEndpoint, err)
	}
	defer resp.Body.Close()

	xmlData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Failed to read xml data from response body: %w", err)
	}

	solar := Solar{}
	err = xml.Unmarshal(xmlData, &solar)
	if err != nil {
		return nil, fmt.Errorf("Failed to unmarshal xml data from response: %w", err)
	}

	return solar.SolarData.CalculatedConditition.Bands, nil
}

func separateBandsPerTime(bands []Band) ([]Band, []Band) {
	const (
		day = "day"
		night = "night"
	)

	dayBands := []Band{}
	nightBands := []Band{}

	for _, condition := range bands {
		if condition.Time == day {
			dayBands = append(dayBands, condition)
		} else {
			nightBands = append(nightBands, condition)
		}
	}

	return dayBands, nightBands
}
