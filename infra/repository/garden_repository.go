package repository

import (
	"context"
	"study-event-go/domain/entity"
	"study-event-go/domain/interfaces"
	"study-event-go/ent"
	entGarden "study-event-go/ent/garden"
	"study-event-go/types"

	"github.com/juju/errors"
)

type gardenRepository struct {
	conn *ent.Client
}

// NewGardenRepository ...
func NewGardenRepository(conn *ent.Client) interfaces.GardenRepository {
	return &gardenRepository{
		conn: conn,
	}
}

func (g *gardenRepository) Save(ctx context.Context, garden *entity.Garden) (*entity.Garden, error) {

	id, err := g.conn.Garden.
		Create().
		SetName(garden.Name).
		SetLocation(garden.Location).
		SetMentorshipID(garden.MentorshipID).
		SetLegionSystem(garden.LegionSystem).
		OnConflict().
		UpdateLocation().
		UpdateMentorshipID().
		UpdateUpdatedAt().
		ClearDeletedAt().
		ID(ctx)
	if err != nil {
		// logger
		return nil, errors.New("internal server error")
	}

	entModel, err := g.conn.Garden.Get(ctx, id)
	if err != nil {
		// logger
		return nil, errors.New("internal server error")
	}

	garden.ID = entModel.ID
	garden.CreatedAt = entModel.CreatedAt
	garden.UpdatedAt = entModel.UpdatedAt
	garden.DeletedAt = entModel.DeletedAt
	garden.Name = entModel.Name
	garden.Location = entModel.Location
	garden.MentorshipID = entModel.MentorshipID
	garden.LegionSystem = entModel.LegionSystem

	return garden, nil
}

func (g *gardenRepository) Garden(ctx context.Context, id types.GardenID) (*entity.Garden, error) {

	entModel, err := g.conn.Garden.
		Query().
		Where(
			entGarden.ID(id),
			entGarden.DeletedAtIsNil()).
		Only(ctx)
	if err != nil {
		// logger
		return nil, errors.NotFoundf("id[%d]", id)
	}

	return &entity.Garden{
		ID:           entModel.ID,
		CreatedAt:    entModel.CreatedAt,
		UpdatedAt:    entModel.UpdatedAt,
		DeletedAt:    entModel.DeletedAt,
		Name:         entModel.Name,
		Location:     entModel.Location,
		MentorshipID: entModel.MentorshipID,
		LegionSystem: entModel.LegionSystem,
	}, nil
}

func (g *gardenRepository) GardenByName(ctx context.Context, name string) (*entity.Garden, error) {

	entModel, err := g.conn.Garden.
		Query().
		Where(
			entGarden.Name(name),
			entGarden.DeletedAtIsNil()).
		Only(ctx)
	if err != nil {
		// logger
		if ent.IsNotFound(err) { // TODO: converter (ent error -> error)
			return nil, errors.NotFoundf("The name[%s]", name)
		}
		return nil, errors.New("internal server error")
	}

	return &entity.Garden{
		ID:           entModel.ID,
		CreatedAt:    entModel.CreatedAt,
		UpdatedAt:    entModel.UpdatedAt,
		DeletedAt:    entModel.DeletedAt,
		Name:         entModel.Name,
		Location:     entModel.Location,
		MentorshipID: entModel.MentorshipID,
		LegionSystem: entModel.LegionSystem,
	}, nil
}

func (g *gardenRepository) Gardens(ctx context.Context, offset uint32) ([]*entity.Garden, error) {

	entModels, err := g.conn.Garden.
		Query().
		Where(
			entGarden.DeletedAtIsNil()).
		Limit(10).
		Offset(int(offset)).
		All(ctx)
	if err != nil {
		// logger
		return nil, errors.New("internal server error")
	}

	gardens := make([]*entity.Garden, len(entModels))
	for i, v := range entModels {
		gardens[i] = &entity.Garden{
			ID:           v.ID,
			CreatedAt:    v.CreatedAt,
			UpdatedAt:    v.UpdatedAt,
			DeletedAt:    v.DeletedAt,
			Name:         v.Name,
			Location:     v.Location,
			MentorshipID: v.MentorshipID,
			LegionSystem: v.LegionSystem,
		}
	}

	return gardens, nil
}

func (g *gardenRepository) Update(ctx context.Context, garden *entity.Garden) (*entity.Garden, error) {

	entModel, err := g.conn.Garden.
		UpdateOneID(garden.ID).
		SetName(garden.Name).
		SetLocation(garden.Location).
		SetMentorshipID(garden.MentorshipID).
		SetLegionSystem(garden.LegionSystem).
		SetNillableDeletedAt(garden.DeletedAt).
		Save(ctx)
	if err != nil {
		// logger
		return nil, errors.New("internal server error")
	}

	garden.ID = entModel.ID
	garden.CreatedAt = entModel.CreatedAt
	garden.UpdatedAt = entModel.UpdatedAt
	garden.DeletedAt = entModel.DeletedAt
	garden.Name = entModel.Name
	garden.Location = entModel.Location
	garden.MentorshipID = entModel.MentorshipID
	garden.LegionSystem = entModel.LegionSystem

	return garden, nil
}
