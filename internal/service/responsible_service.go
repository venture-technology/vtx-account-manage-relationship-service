package service

import "github.com/venture-technology/vtx-account-manager/internal/repository"

type ResponsibleService struct {
	responsiblerepository repository.IResponsibleRepository
}

func NewResponsibleService(repo repository.IResponsibleRepository) *ResponsibleService {
	return &ResponsibleService{
		responsiblerepository: repo,
	}
}
