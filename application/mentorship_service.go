package application

import (
	"context"
	"study-event-go/application/dto"
	"study-event-go/domain/entity"
	"study-event-go/domain/interfaces"
	"study-event-go/types"
)

// MentorshipService ...
type MentorshipService struct {
	mentorshipRepo interfaces.MentorshipRepository
}

// NewMentorshipService ...
func NewMentorshipService(
	mentorshipRepo interfaces.MentorshipRepository,
) *MentorshipService {
	return &MentorshipService{
		mentorshipRepo: mentorshipRepo,
	}
}

// TODO: logger

// New ...
func (m *MentorshipService) New(ctx context.Context, mentorshipDTO *dto.Mentorship) (*dto.Mentorship, error) {

	mentorship, err := m.mentorshipRepo.New(ctx, entity.NewMentorship(mentorshipDTO))
	if err != nil {
		return nil, err
	}

	return mentorship.DTO(), nil
}

// Get ...
func (m *MentorshipService) Get(ctx context.Context, id types.MentorshipSystemID) (*dto.Mentorship, error) {

	mentorship, err := m.mentorshipRepo.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return mentorship.DTO(), nil
}

// List ...
func (m *MentorshipService) List(ctx context.Context, offset uint32) ([]*dto.Mentorship, error) {

	// TODO: cursor pagination

	mentorships, err := m.mentorshipRepo.List(ctx, offset)
	if err != nil {
		return nil, err
	}

	results := make([]*dto.Mentorship, len(mentorships))
	for i, v := range mentorships {
		results[i] = &dto.Mentorship{
			ID:        v.ID,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
			DeletedAt: v.DeletedAt,
			Name:      v.Name,
		}
	}

	return results, nil
}