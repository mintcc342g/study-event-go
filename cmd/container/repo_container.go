package container

import (
	"context"
	"study-event-go/domain/interfaces"
	"study-event-go/ent"
	"study-event-go/ent/migrate"
	"study-event-go/infra/repository"

	"github.com/RichardKnop/machinery/v1"
	"github.com/go-redis/redis/v8"
)

// RepositoryContainer ...
type RepositoryContainer struct {
	RedisRepo      interfaces.RedisRepository
	LilyRepo       interfaces.LilyRepository
	GardenRepo     interfaces.GardenRepository
	EventRepo      interfaces.EventRepository
	MentorshipRepo interfaces.MentorshipRepository
}

func newRepositoryContainer(db *ent.Client, redis redis.Cmdable, machineryServer *machinery.Server) *RepositoryContainer {
	return &RepositoryContainer{
		RedisRepo:      repository.NewRedisRepository(redis),
		LilyRepo:       repository.NewLilyRepository(db),
		GardenRepo:     repository.NewGardenRepository(db),
		EventRepo:      repository.NewMachineryRepository(machineryServer),
		MentorshipRepo: repository.NewMentorshipRepository(db),
	}
}

// InitRepositoryContainer ...
func InitRepositoryContainer(db *ent.Client, redis redis.Cmdable, machineryServer *machinery.Server) (*RepositoryContainer, error) {
	ctx := context.Background()
	if err := db.Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	); err != nil {
		println("InitRepositoryContainer", "Create", "error", err) // TODO: logger
		return nil, err
	}

	return newRepositoryContainer(db, redis, machineryServer), nil
}
