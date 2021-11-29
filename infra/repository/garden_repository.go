package repository

import (
	"context"
	"study_event_go/domain/entity"
	"study_event_go/domain/interfaces"
)

type gardenRepository struct {
}

// NewGardenRepository ...
func NewGardenRepository() interfaces.GardenRepository {
	return &gardenRepository{}
}

func (g *gardenRepository) Garden(ctx context.Context, id uint64) (*entity.Garden, error) {

	// TODO: RDB

	return nil, nil
}
