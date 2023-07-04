package service

import (
	"learn-gqlgen/domain/repository"
)

type depressionService struct {
	depressionRepository repository.DepressionRepository
}

type DepressionService interface {
	GetForDepressionIds()
}

func (s *depressionService) GetForDepressionIds() {
}

func NewDepressionService(
	depressionRepository repository.DepressionRepository,
) DepressionService {
	return &depressionService{
		depressionRepository,
	}
}
