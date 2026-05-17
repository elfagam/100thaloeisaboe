package usecase

import (
	"context"
	"backend-rs100/internal/domain"
)

type milestoneUsecase struct {
	milestoneRepo domain.MilestoneRepository
}

// NewMilestoneUsecase creates a new milestone usecase.
func NewMilestoneUsecase(repo domain.MilestoneRepository) domain.MilestoneUsecase {
	return &milestoneUsecase{
		milestoneRepo: repo,
	}
}

func (u *milestoneUsecase) Fetch(ctx context.Context) ([]domain.Milestone, error) {
	return u.milestoneRepo.Fetch(ctx)
}

func (u *milestoneUsecase) Store(ctx context.Context, item *domain.Milestone) error {
	return u.milestoneRepo.Store(ctx, item)
}

func (u *milestoneUsecase) Update(ctx context.Context, item *domain.Milestone) error {
	return u.milestoneRepo.Update(ctx, item)
}

func (u *milestoneUsecase) Delete(ctx context.Context, id string) error {
	return u.milestoneRepo.Delete(ctx, id)
}
