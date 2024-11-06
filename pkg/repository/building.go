package repository

import "leadgen/pkg/model"

func (s *S) CreateBuilding(building *model.Building) error {
	_, err := s.db.Model(building).Insert()
	return err
}

func (s *S) ListBuildings(city string, year, floors int) ([]model.Building, error) {
	var buildings []model.Building
	query := s.db.Model(&buildings)

	if city != "" {
		query.Where("city = ?", city)
	}
	if year != 0 {
		query.Where("year_built = ?", year)
	}
	if floors != 0 {
		query.Where("floors = ?", floors)
	}

	err := query.Select()
	return buildings, err
}
