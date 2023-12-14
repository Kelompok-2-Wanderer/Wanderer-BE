package service

import (
	"errors"
	"wanderer/features/airlines"
)

func NewAirlineService(repo airlines.Repository) airlines.Service {
	return &airlineService{
		repo: repo,
	}
}

type airlineService struct {
	repo airlines.Repository
}

func (srv *airlineService) Create(newAirline airlines.Airline) error {
	if newAirline.Name == "" {
		return errors.New("validate: name can't be empty")
	}

	if err := srv.repo.Create(newAirline); err != nil {
		return err
	}

	return nil
}

func (srv *airlineService) GetAll() ([]airlines.Airline, error) {
	result, err := srv.repo.GetAll()
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (srv *airlineService) Update(id uint, updateAirline airlines.Airline) error {
	if id == 0 {
		return errors.New("validate: invalid airline id")
	}

	if err := srv.repo.Update(id, updateAirline); err != nil {
		return err
	}

	return nil
}

func (srv *airlineService) Delete(id uint) error {
	if id == 0 {
		return errors.New("validate: invalid airline id")
	}

	if err := srv.repo.Delete(id); err != nil {
		return err
	}

	return nil
}