package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func FormatDate(dateStr string) string {
	parts := strings.Split(dateStr, "-")
	if len(parts) != 3 {
		return dateStr
	}

	year, _ := strconv.Atoi(parts[0])
	month, _ := strconv.Atoi(parts[1])
	day, _ := strconv.Atoi(parts[2])

	return fmt.Sprintf("%d %s %d", day, months[month], year)
}

func GetCountry(country string) string {
	return countries[country]
}

func GetUserInput() (string, string) {
	var country, city string

	fmt.Println(WelcomeMsg)
	fmt.Println(DashLine)

	fmt.Println(SelectCountryMsg)
	for key := range countries {
		fmt.Printf("%s. %s\n", key, countries[key])
	}
	fmt.Print(CountryChoiceMsg)
	_, err := fmt.Scan(&country)
	if err != nil {
		return "", ""
	}
	country = countries[country]

	if country == "" {
		log.Fatal(InvalidCountryMsg)
	}

	fmt.Println(SelectCityMsg)
	for key, value := range citiesTR {
		fmt.Printf("%s. %s\n", key, value)
	}
	fmt.Print(CityChoiceMsg)
	_, err = fmt.Scan(&city)
	if err != nil {
		return "", ""
	}
	city = citiesTR[city]

	if city == "" {
		log.Fatal(InvalidCityMsg)
	}

	return country, city
}
