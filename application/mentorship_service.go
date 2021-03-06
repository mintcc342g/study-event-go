package application

import (
	"context"
	"study-event-go/application/dto"
	"study-event-go/domain/entity"
	"study-event-go/domain/interfaces"
	"study-event-go/types"

	"github.com/juju/errors"
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

	_, err := m.mentorshipRepo.MentorshipByName(ctx, mentorshipDTO.Name)
	if err != nil && !errors.IsNotFound(err) {
		return nil, err
	}

	mentorship, err := entity.NewMentorship(mentorshipDTO)
	if err != nil {
		return nil, err
	}

	if mentorship, err = m.mentorshipRepo.Save(ctx, mentorship); err != nil {
		return nil, err
	}

	return mentorship.DTO(), nil
}

// Get ...
func (m *MentorshipService) Get(ctx context.Context, id types.MentorshipID) (*dto.Mentorship, error) {

	mentorship, err := m.mentorshipRepo.Mentorship(ctx, id)
	if err != nil {
		return nil, err
	}

	return mentorship.DTO(), nil
}

// List ...
func (m *MentorshipService) List(ctx context.Context, offset uint32) ([]*dto.Mentorship, error) {

	// TODO: cursor pagination

	mentorships, err := m.mentorshipRepo.Mentorships(ctx, offset)
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

// Update ...
func (m *MentorshipService) Update(ctx context.Context, mentorshipDTO *dto.Mentorship) (*dto.Mentorship, error) {

	comparable, err := m.mentorshipRepo.MentorshipByName(ctx, mentorshipDTO.Name)
	if err != nil && !errors.IsNotFound(err) {
		return nil, err
	}

	mentorship, err := m.mentorshipRepo.Mentorship(ctx, mentorshipDTO.ID)
	if err != nil {
		return nil, err
	}

	if err = mentorship.Update(mentorshipDTO, comparable); err != nil {
		return nil, err
	}

	if mentorship, err = m.mentorshipRepo.Update(ctx, mentorship); err != nil {
		return nil, err
	}

	return mentorship.DTO(), nil
}

// SoftDelete ...
func (m *MentorshipService) SoftDelete(ctx context.Context, id types.MentorshipID) (*dto.Mentorship, error) {

	mentorship, err := m.mentorshipRepo.Mentorship(ctx, id)
	if err != nil {
		return nil, err
	}

	mentorship.Delete()

	if mentorship, err = m.mentorshipRepo.Update(ctx, mentorship); err != nil {
		return nil, err
	}

	return mentorship.DTO(), nil
}
