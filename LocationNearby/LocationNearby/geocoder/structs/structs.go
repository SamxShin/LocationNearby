package structs

// All results from the JSON
type Results struct {
	Results []Result `json:"results"`
	Status  string   `json:"status"`
}

// Result store each result from the JSON
type Result struct {
	AddressComponents []Address `json:"address_components"`
	FormattedAddress  string    `json:"formatted_address"`
	Geometry          Geometry  `json:"geometry"`
	PlaceId           string    `json:"place_id"`
	Types             []string  `json:"types"`
}

// Address store each address is identified by the 'types'
type Address struct {
	LongName  string   `json:"long_name"`
	ShortName string   `json:"short_name"`
	Types     []string `json:"types"`
}

// Geometry store each value in the geometry
type Geometry struct {
	Location LatLng `json:"location"`
}

// LatLng store the latitude and longitude
type LatLng struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}
