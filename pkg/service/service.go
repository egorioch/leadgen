package service

import (
	"github.com/sirupsen/logrus"
	"leadgen/pkg/model"
)

type Repository interface {
	CreateBuilding(building *model.Building) error
	ListBuildings(city string, year, floors int) ([]model.Building, error)
}

type S struct {
	log *logrus.Logger
	r   Repository
}

func New(log *logrus.Logger, r Repository) *S {
	return &S{
		log: log,
		r:   r,
	}
}
