package repository

import (
	"learn-gqlgen/graph/model"

	"gorm.io/gorm"
)

type depressionRepository struct{}

type DepressionRepository interface {
	FindByIDs(tx *gorm.DB, ids uint64) (*model.Depression, error)
}

func (r *depressionRepository) FindByIDs(tx *gorm.DB, ids uint64) (*model.Depression, error) {
	var depressionIds *model.Depression
	result := tx.Where("despression.id IN (?)", ids).Find(&depressionIds)
	if result.Error != nil {
		return nil, result.Error
	}

	return depressionIds, nil
}

func NewDepressionRepository() DepressionRepository {
	return &depressionRepository{}
}
