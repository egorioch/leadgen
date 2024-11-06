package service

import "leadgen/pkg/model"

func (s *S) CreateBuilding(building *model.Building) error {
	return s.r.CreateBuilding(building)
}

func (s *S) ListBuildings(city string, year, floors int) ([]model.Building, error) {
	return s.r.ListBuildings(city, year, floors)
}
