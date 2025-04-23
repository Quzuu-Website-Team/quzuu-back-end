package services

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"godp.abdanhafidz.com/models"
	"godp.abdanhafidz.com/repositories"
)

type ProvinceService struct {
	Service[models.RegionProvince, []models.RegionProvince]
}

type CityService struct {
	Service[models.RegionCity, []models.RegionCity]
}

func seedCity() ([]models.RegionCity, error) {
	log.Println("Seed City")
	_, b, _, _ := runtime.Caller(0)
	basePath := filepath.Dir(b)
	file, err := os.Open(filepath.Join(basePath, "..", "utils", "seeds", "city.json"))
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var cities []models.RegionCity
	if err := json.NewDecoder(file).Decode(&cities); err != nil {
		return nil, err
	}
	return cities, nil
}

func seedProvince() ([]models.RegionProvince, error) {
	log.Println("Seed City")
	_, b, _, _ := runtime.Caller(0)
	basePath := filepath.Dir(b)
	file, err := os.Open(filepath.Join(basePath, "..", "utils", "seeds", "province.json"))
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var provinces []models.RegionProvince
	if err := json.NewDecoder(file).Decode(&provinces); err != nil {
		return nil, err
	}
	return provinces, nil
}

func (s *ProvinceService) Create() {
	provinces, errSeed := seedProvince()
	if errSeed != nil {
		s.Error = errSeed
		s.Exception.InternalServerError = true
		s.Exception.Message = "Failed to seed province"
		return
	}
	createProvince := repositories.BulkCreateProvince(provinces)
	s.Error = createProvince.RowsError
	s.Result = createProvince.Result
}

func (s *CityService) Create() {
	cities, errSeed := seedCity()
	if errSeed != nil {
		s.Error = errSeed
		s.Exception.InternalServerError = true
		s.Exception.Message = "Failed to seed province"
		return
	}
	createCity := repositories.BulkCreateCity(cities)
	s.Error = createCity.RowsError
	s.Result = createCity.Result
}

func (s *ProvinceService) Retrieve() {
	Province := repositories.GetListProvinces()
	s.Error = Province.RowsError
	s.Result = Province.Result
}

func (s *CityService) Retrieve() {
	cities := repositories.GetListCitiesByProvinceId(s.Constructor.ProvinceId)
	s.Error = cities.RowsError
	s.Result = cities.Result
}
