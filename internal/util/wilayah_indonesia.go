package util

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Province struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type City struct {
	ID         string `json:"id"`
	ProvinceID string `json:"province_id"`
	Name       string `json:"name"`
}

func GetProvinceByID(id string) (*Province, error) {
	resp, err := http.Get("https://www.emsifa.com/api-wilayah-indonesia/api/provinces.json")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var provinces []Province
	if err := json.NewDecoder(resp.Body).Decode(&provinces); err != nil {
		return nil, err
	}

	for _, province := range provinces {
		if province.ID == id {
			return &province, nil
		}
	}
	return nil, fmt.Errorf("Province not found")
}

func GetCityByID(provinceID, cityID string) (*City, error) {
	url := fmt.Sprintf("https://www.emsifa.com/api-wilayah-indonesia/api/regencies/%s.json", provinceID)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var cities []City
	if err := json.NewDecoder(resp.Body).Decode(&cities); err != nil {
		return nil, err
	}

	for _, city := range cities {
		if city.ID == cityID {
			return &city, nil
		}
	}
	return nil, fmt.Errorf("City not found")
}