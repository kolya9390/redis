// Package classification Geting Adresses.
//
// Documentation of my project API.
//
//	Schemes:
//	- http
//	- https
//	BasePath: /
//	Version: 1.0.0
//
//	Consumes:
//	- application/json
//	- multipart/form-data
//
//	Produces:
//	- application/json
//
// swagger:meta
package docs

import docsresp "studentgit.kata.academy/Nikolai/historysearch/internal/docs/resp"

//go:generate swagger generate spec -o ./swagger.json --scan-models

// swagger:route POST /api/address/search AddressSearch
//
// Search for city information based on the provided address query.

// swagger:parameters AddressSearchRequest
type AddressSearchRequest struct {
	// in: body
	Body docsresp.RequestAddressSearch
	// example: "москва сухонская 11"
}


// swagger:response AddressSearchResponse
type AddressSearchResponse struct {
	// in: body
	Body docsresp.ResponseAddress
	// exeple {"addresses":[{"lat":"59.937887","lon":"30.24818","result":"г Санкт-Петербург"}]}
}

// swagger:route POST /api/address/geocode Geocode
// Geocode the provided address to retrieve city information.
//
// responses:
//   200: GeocodeResponse

// swagger:parameters GeocodeRequest
type GeocodeRequest struct {
	// lat_lon координаты (широрта, долгота)
	// in: body
	Body docsresp.RequestAddressGeocode
	// example: {"lat": "59.93784722564821","lng": "30.24881601333618"}
}

// swagger:response GeocodeResponse
type GeocodeResponse struct {
	// Geocode response structure
	// in: body
	Body docsresp.ResponseAddress
	// exeple {"addresses":[{"lat":"59.937887","lon":"30.24818","result":"г Санкт-Петербург"}]}
}

// swagger:route POST /api/auth/register Register
// Register a new user.
//
// Responses:
//   200: RegisterResponse

// swagger:parameters RegisterRequest
type RegisterRequest struct {
	// in: body
	Body docsresp.RegisterRequest
	// example: {"email": "user@example.com", "password": "password123"}
}

// swagger:response RegisterResponse
type RegisterResponse struct {
	// in: body
	Body docsresp.RegisterResponse
	// example: {"message": "User registered successfully"}
}

// swagger:route POST /api/auth/login Login
// Login to the system.
//
// Responses:
//   200: LoginResponse

// swagger:parameters LoginRequest
type LoginRequest struct {
	// in: body
	Body docsresp.LoginRequest
	// example: {"email": "user@example.com", "password": "password123"}
}


// swagger:response LoginResponse
type LoginResponse struct {
	// in: body
	Body docsresp.AuthResponse
	// example: {"access_token": "your_access_token", "message": "Logged in successfully"}
}

//TODO USER