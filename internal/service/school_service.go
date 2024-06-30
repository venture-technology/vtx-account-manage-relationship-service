package service

import "github.com/venture-technology/vtx-account-manager/internal/repository"

type SchoolService struct {
	schoolrepository repository.ISchoolRepository
}

func NewSchoolService(repo repository.ISchoolRepository) *SchoolService {
	return &SchoolService{
		schoolrepository: repo,
	}
}
