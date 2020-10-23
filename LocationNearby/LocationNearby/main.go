package main

import (
	"LocationNearby/geocoder"
	"bufio"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	//replace this with personal api code attained from google geocode API
	geocoder.APIKey = "AIzaSyCL6dhFaEbyzxOfSHlcmkthpS48xu0jSYQ"
	var street, city, state, country string
	var number int
	var lng, lat float64

	scanner := bufio.NewScanner(os.Stdin)

loop:
	for {
		fmt.Println("\nType 'one' to find longitude and latitude by using address name \nType 'two' to find address name using" +
			" latitude and longitude\nType 'three' to exit")
		var choice string
		fmt.Scanln(&choice)

		switch choice {
		case "one":
			fmt.Print("Enter street number: ")
			fmt.Scanln(&number)

			fmt.Print("Enter the street name: ")
			scanner.Scan()
			street := scanner.Text()

			fmt.Print("Enter the city name: ")
			scanner.Scan()
			city := scanner.Text()

			fmt.Print("Enter the state name: ")
			scanner.Scan()
			state := scanner.Text()

			fmt.Print("Enter the country name: ")
			scanner.Scan()
			country := scanner.Text()

			//initialize the data into geocoder.Address struct
			address := geocoder.Address{
				Street:  street,
				Number:  number,
				City:    city,
				State:   state,
				Country: country,
			}

			location := geocoder.Location{
				Latitude:  lat,
				Longitude: lng,
			}
			//this converts address to location(latitude and longitude)
			location, err := geocoder.Geocoding(address)
			//error message
			if err != nil {
				fmt.Println("Could not get the location: ", err)
			} else {
				fmt.Println("Latitude: ", location.Latitude) //prints data
				fmt.Println("Longitude: ", location.Longitude)
			}
			if err != nil {
				fmt.Println("Could not get the addresses: ", err)
			} else {
				// Print the address formatted by the geocoder package
				fmt.Println(address.AddressFormatter())
			}
			break
		case "two":
			fmt.Print("\nEnter latitude: ")
			fmt.Scanln(&lat)
			fmt.Print("Enter longitude: ")
			fmt.Scanln(&lng)

			address := geocoder.Address{
				Street:  street,
				Number:  number,
				City:    city,
				State:   state,
				Country: country,
			}

			location := geocoder.Location{
				Latitude:  lat,
				Longitude: lng,
			}
			//converts location to slices of addresses
			addresses, err := geocoder.GeocodingReverse(location)
			if err != nil {
				fmt.Println("Could not get the addresses: ", err)
			} else {

				address = addresses[0]
				// Print the formatted address from the API
				fmt.Println(address.FormattedAddress)
			}
			break
		case "three":
			break loop
		default:
			fmt.Println("invalid input, try again")
		}
	}
}
