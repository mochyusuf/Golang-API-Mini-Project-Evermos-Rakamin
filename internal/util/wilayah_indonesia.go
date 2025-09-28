package util

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sync"
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


func GetAllProvinces() ([]Province, error) {
	resp, err := http.Get("https://www.emsifa.com/api-wilayah-indonesia/api/provinces.json")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var provinces []Province
	if err := json.NewDecoder(resp.Body).Decode(&provinces); err != nil {
		return nil, err
	}
	return provinces, nil
}

func GetAllCitiesByProvinceID(provinceID string) ([]City, error) {
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
	return cities, nil
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

func GetCityByIDOnly(cityID string) (*City, error) {
	provinceID := "11"
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

var (
	citiesCache map[string]City
	cacheMutex  sync.Mutex
)

func loadCities() error {
	file, err := os.Open("D:/projek/project fix/Evermos-Virtual-Intern/data/regencies.csv")
	if err != nil {
		return fmt.Errorf("failed to open regencies.csv: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("failed to read regencies.csv: %v", err)
	}

	tempCache := make(map[string]City)
	for i, record := range records {
		if i == 0 {
			continue // Skip header if any
		}
		if len(record) < 3 {
			continue // skip malformed rows
		}
		city := City{
			ID:         record[0],
			ProvinceID: record[1],
			Name:       record[2],
		}
		tempCache[city.ID] = city
	}

	cacheMutex.Lock()
	citiesCache = tempCache
	cacheMutex.Unlock()

	return nil
}

func GetCityByIDcityID(cityID string) (*City, error) {
	cacheMutex.Lock()
	defer cacheMutex.Unlock()

	// Load cache if empty
	if len(citiesCache) == 0 {
		if err := loadCities(); err != nil {
			return nil, err
		}
	}

	city, exists := citiesCache[cityID]
	if !exists {
		return nil, fmt.Errorf("City not found")
	}
	return &city, nil
}