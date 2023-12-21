package servis

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"studentgit.kata.academy/Nikolai/historysearch/internal/config"
)

type DadataService interface {
	SearchAddress(query string) (respDadataAdres, error)
	GeocodeAddress(lat, lng string) (responseDadataGeo, error)
}



type DadataServiceImpl struct {
	client http.Client
	AuthorizationDADATA	config.AuthorizationDADATA
}

type respDadataAdres []struct {
	Region string `json:"region"`
	GeoLat string `json:"geo_lat"`
	GeoLon string `json:"geo_lon"`
}

type responseDadataGeo struct {
	Suggestions []struct {
		Value string `json:"value"`
		Data  struct {
			GeoLat string `json:"geo_lat"`
			GeoLon string `json:"geo_lon"`
			Result string `json:"region_with_type"`
		} `json:"data"`
	} `json:"suggestions"`
}



func NewDadataService( /* Env? */ ) (DadataService,error) {

	return &DadataServiceImpl{
		client: http.Client{},
		},nil
}


// Реализация методов Запроса к ДаДата
func (d *DadataServiceImpl) makeRequest(ctx context.Context, url, method, contentType string, body io.Reader) (*http.Response, error) {

	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, fmt.Errorf("ошибка создания запроса: %w", err)
	}

	req.Header.Set("Content-Type", contentType)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Token %s", d.AuthorizationDADATA.ApiKeyValue))
	req.Header.Set("X-Secret", d.AuthorizationDADATA.SecretKeyValue)


	return d.client.Do(req)

}

// Реализация методов DadataService
// SearchAddress и GeocodeAddress
func(d *DadataServiceImpl) SearchAddress(query string) (respDadataAdres, error){

	ctx := context.Background()
	// москва сухонская 11

	data := strings.NewReader(fmt.Sprintf(`[ "%s" ]`, query))


	url := "https://cleaner.dadata.ru/api/v1/clean/address"

	resp, err := d.makeRequest(ctx, url, "POST", "application/json", data)
	if err != nil {
		return nil, fmt.Errorf("ошибка запроса к Dadata API: %w", err)
	}

	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var respData respDadataAdres

	//log.Fatalf("%s",bodyText)
	err = json.Unmarshal(bodyText, &respData)
	if err != nil {
		return nil, err
	}

	

	return respData,nil

}

func(d *DadataServiceImpl)	GeocodeAddress(lat, lng string) (responseDadataGeo, error){

	ctx := context.Background()
	// москва сухонская 11

	data := strings.NewReader(fmt.Sprintf(`{ "lat": %s, "lon": %s }`, lat, lng))

	url := "https://suggestions.dadata.ru/suggestions/api/4_1/rs/geolocate/address"

	resp, err := d.makeRequest(ctx, url, "POST", "application/json", data)
	if err != nil {
		return responseDadataGeo{}, fmt.Errorf("ошибка запроса к Dadata API: %w", err)
	}

	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		return responseDadataGeo{}, err
	}

	var respData responseDadataGeo

	err = json.Unmarshal(bodyText, &respData)
	if err != nil {
		return responseDadataGeo{},err
	}

	return respData,nil
}