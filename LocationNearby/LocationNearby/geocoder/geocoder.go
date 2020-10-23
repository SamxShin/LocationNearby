package geocoder

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"

	"LocationNearby/geocoder/structs"
)

// the user is required to have a geocode API key from google
var APIKey string

//we define the Geocode API URL as a const.
const (
	geocodeAPIUrl = "https://maps.googleapis.com/maps/api/geocode/json?"
)

type Location struct {
	Longitude float64
	Latitude  float64
}

// this is the address structure used in the Geocoding and Geocoding reverse functions
// the FormattedAddress is needed for Geocodingrevrse function
type Address struct {
	Number           int
	Street           string
	City             string
	State            string
	PostalCode       string
	Country          string
	FormattedAddress string
}

//formats the address based on the Address struct and returns it as a string
func (address *Address) AddressFormatter() string {

	//create a slice with allthe ocntent from Address struct
	var data []string

	if address.Number > 0 {
		data = append(data, strconv.Itoa(address.Number)) //convert number to a string
	}
	data = append(data, address.Street)
	data = append(data, address.City)
	data = append(data, address.State)
	data = append(data, address.PostalCode)
	data = append(data, address.Country)

	var formattedAddress string

	//adds a ',' in needed places for the formattedAddress string slices
	for _, value := range data {
		if value != "" {
			if formattedAddress != "" {
				formattedAddress += ", "
			}
			formattedAddress += value
		}
	}
	return formattedAddress
}

//httpRequest function sends the request to decode the JSON and stores in Results struct
func httpRequest(url string) (structs.Results, error) {
	var results structs.Results

	//this code builds the request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return results, err
	}
	// For control over HTTP client headers, redirect policy, and other settings, create a Client
	// A Client is an HTTP client
	client := &http.Client{}

	// Send the request via a client
	// Do sends an HTTP request and returns an HTTP response
	resp, err := client.Do(req)
	if err != nil {
		return results, err
	}
	// Callers should close resp.Body when done reading from it
	// Defer the closing of the body
	defer resp.Body.Close()

	// Use json.Decode for reading streams of JSON data
	err = json.NewDecoder(resp.Body).Decode(&results)
	if err != nil {
		return results, err
	}

	// The "OK" status indicates that no error has occurred, it means
	// the address was analyzed and at least one geographic code was returned
	if strings.ToUpper(results.Status) != "OK" {
		// If the status is not "OK" check what status was returned
		switch strings.ToUpper(results.Status) {
		case "ZERO_RESULTS":
			err = errors.New("No results found.")
			break
		case "OVER_QUERY_LIMIT":
			err = errors.New("You are over your quota.")
			break
		case "REQUEST_DENIED":
			err = errors.New("Your request was denied.")
			break
		case "INVALID_REQUEST":
			err = errors.New("Probably the query is missing.")
			break
		case "UNKNOWN_ERROR":
			err = errors.New("Server error. Please, try again.")
			break
		default:
			break
		}
	}

	return results, err
}

// Geocoding function is used to convert an Address structure
// to a Location structure (latitude and longitude)
func Geocoding(address Address) (Location, error) {

	var location Location

	// Convert whitespaces to +
	formattedAddress := address.AddressFormatter()
	formattedAddress = strings.Replace(formattedAddress, " ", "+", -1)

	// Create the URL based on the formated address
	url := geocodeAPIUrl + "address=" + formattedAddress

	// Use the API Key if it was set
	if APIKey != "" {
		url += "&key=" + APIKey
	}

	// Send the HTTP request and get the results
	results, err := httpRequest(url)
	if err != nil {
		log.Println(err)
		return location, err
	}

	// Get the results (latitude and longitude)
	location.Latitude = results.Results[0].Geometry.Location.Lat
	location.Longitude = results.Results[0].Geometry.Location.Lng

	return location, nil
}

// Convert a structs.Results to a slice of Address structures
func convertResultsToAddress(results structs.Results) (addresses []Address) {

	for index := 0; index < len(results.Results); index++ {
		var address Address

		// Put each component from the AddressComponents slice in the correct field in the Address structure
		for _, component := range results.Results[index].AddressComponents {
			// Check all types of each component
			for _, types := range component.Types {
				switch types {
				case "route":
					address.Street = component.LongName
					break
				case "street_number":
					address.Number, _ = strconv.Atoi(component.LongName)
					break
				case "locality":
					address.City = component.LongName
					break
				case "administrative_area_level_3":
					address.City = component.LongName
					break
				case "administrative_area_level_1":
					address.State = component.LongName
					break
				case "country":
					address.Country = component.LongName
					break
				case "postal_code":
					address.PostalCode = component.LongName
					break
				default:
					break
				}
			}
		}

		address.FormattedAddress = results.Results[index].FormattedAddress

		addresses = append(addresses, address)
	}
	return
}

// GeocodingReverse function is used to convert a Location structure
// to an Address structure
func GeocodingReverse(location Location) ([]Address, error) {

	var addresses []Address

	url := getURLGeocodingReverse(location, "")

	// Send the HTTP request and get the results
	results, err := httpRequest(url)
	if err != nil {
		log.Println(err)
		return addresses, err
	}

	// Convert the results to an Address slice called addresses
	addresses = convertResultsToAddress(results)

	return addresses, nil
}

func GeocodingReverseIntl(location Location, language string) ([]Address, error) {

	var addresses []Address

	url := getURLGeocodingReverse(location, language)

	// Send the HTTP request and get the results
	results, err := httpRequest(url)
	if err != nil {
		log.Println(err)
		return addresses, err
	}

	// Convert the results to an Address slice called addresses
	addresses = convertResultsToAddress(results)

	return addresses, nil
}

func getURLGeocodingReverse(location Location, language string) string {
	// Convert the latitude and longitude from double to string
	latitude := strconv.FormatFloat(location.Latitude, 'f', 8, 64)
	longitude := strconv.FormatFloat(location.Longitude, 'f', 8, 64)

	// Create the URL based on latitude and longitude
	url := geocodeAPIUrl + "latlng=" + latitude + "," + longitude

	// Use the API key if it was set
	if APIKey != "" {
		url += "&key=" + APIKey
	}

	if language != "" {
		url += "&language=" + language
	}

	return url
}
