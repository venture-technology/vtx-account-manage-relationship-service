package service

import "github.com/venture-technology/vtx-account-manager/internal/repository"

type DriverService struct {
	driverrepository repository.IDriverRepository
}

func NewDriverService(repo repository.IDriverRepository) *DriverService {
	return &DriverService{
		driverrepository: repo,
	}
}
