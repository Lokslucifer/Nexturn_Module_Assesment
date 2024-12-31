package main

import (
	"fmt"
	"strconv"
	"strings"
)

type City struct {
	Name        string
	Temperature float64
	Rainfall    float64
}

func highestTemperature(cities []City) (City, error) {
	if len(cities) == 0 {
		return City{}, fmt.Errorf("no cities in the list")
	}

	highest := cities[0]
	for _, city := range cities[1:] {
		if city.Temperature > highest.Temperature {
			highest = city
		}
	}
	return highest, nil
}

func lowestTemperature(cities []City) (City, error) {
	if len(cities) == 0 {
		return City{}, fmt.Errorf("no cities in the list")
	}

	lowest := cities[0]
	for _, city := range cities[1:] {
		if city.Temperature < lowest.Temperature {
			lowest = city
		}
	}
	return lowest, nil
}

func averageRainfall(cities []City) float64 {
	var totalRainfall float64
	for _, city := range cities {
		totalRainfall += city.Rainfall
	}
	return totalRainfall / float64(len(cities))
}

func filterCitiesByRainfall(cities []City, threshold float64) []City {
	var filtered []City
	for _, city := range cities {
		if city.Rainfall > threshold {
			filtered = append(filtered, city)
		}
	}
	return filtered
}

func searchByCityName(cities []City, name string) (City, error) {
	for _, city := range cities {
		if strings.EqualFold(city.Name, name) {
			return city, nil
		}
	}
	return City{}, fmt.Errorf("city %s not found", name)
}

func main() {
	cities := []City{
		{"Delhi", 25.0, 150.5},
		{"Mumbai", 28.0, 250.3},
		{"Kolkata", 30.0, 200.8},
		{"Chennai", 32.5, 120.2},
		{"Bangalore", 23.5, 120.1},
		{"Hyderabad", 27.0, 90.4},
		{"Pune", 22.0, 85.3},
		{"Jaipur", 34.0, 150.7},
		{"Ahmedabad", 30.5, 100.2},
		{"Kochi", 29.0, 300.4},
	}

	var choice int
	for {
		fmt.Println("\nMenu:")
		fmt.Println("1. Find highest temperature")
		fmt.Println("2. Find lowest temperature")
		fmt.Println("3. Calculate average rainfall")
		fmt.Println("4. Filter cities by rainfall")
		fmt.Println("5. Search city by name")
		fmt.Println("6. Exit")

		fmt.Print("Enter your choice: ")
		fmt.Scanf("%d", &choice)
		fmt.Scanln() // Clearing the input buffer after scanning integer

		switch choice {
		case 1:
			highest, err := highestTemperature(cities)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Printf("City with highest temperature: %s (%f°C)\n", highest.Name, highest.Temperature)
			}
		case 2:
			lowest, err := lowestTemperature(cities)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Printf("City with lowest temperature: %s (%f°C)\n", lowest.Name, lowest.Temperature)
			}
		case 3:
			avgRainfall := averageRainfall(cities)
			fmt.Printf("Average rainfall across all cities: %.2f mm\n", avgRainfall)
		case 4:
			var thresholdStr string
			fmt.Print("Enter rainfall threshold: ")
			fmt.Scanf("%s", &thresholdStr)
			fmt.Scanln() // Clearing the input buffer after scanning string

			threshold, err := strconv.ParseFloat(thresholdStr, 64)
			if err != nil {
				fmt.Println("Invalid input. Please enter a valid number for the threshold.")
				continue
			}
			filteredCities := filterCitiesByRainfall(cities, threshold)
			if len(filteredCities) == 0 {
				fmt.Println("No cities with rainfall above the given threshold.")
			} else {
				fmt.Println("Cities with rainfall above the threshold:")
				for _, city := range filteredCities {
					fmt.Printf("%s: %.2f mm\n", city.Name, city.Rainfall)
				}
			}
		case 5:
			var cityName string
			fmt.Print("Enter city name to search: ")
			fmt.Scanf("%s", &cityName)
			fmt.Scanln() // Clearing the input buffer after scanning string

			city, err := searchByCityName(cities, cityName)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Printf("City: %s\nTemperature: %.2f°C\nRainfall: %.2f mm\n", city.Name, city.Temperature, city.Rainfall)
			}
		case 6:
			fmt.Println("Exiting program...")
			return
		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}
