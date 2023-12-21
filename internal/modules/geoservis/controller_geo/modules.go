package controllergeo

type Address struct {
	GeoLat string `json:"lat"`
	GeoLon string `json:"lon"`

	Result string `json:"result"`
}

type RequestAddressSearch struct {
	Query string `json:"query"`
}

type ResponseAddress struct {
	Addresses []Address `json:"addresses"`
}

type RequestAddressGeocode struct {
	Lat string `json:"lat"`
	Lng string `json:"lng"`
}