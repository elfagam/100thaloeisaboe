package domain

import "context"

// Milestone represents a historical milestone of the hospital.
type Milestone struct {
	ID          string `json:"id"`
	Year        string `json:"year"`
	Title       string `json:"title"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
}

// MilestoneRepository defines the data access contract for milestones.
type MilestoneRepository interface {
	Fetch(ctx context.Context) ([]Milestone, error)
	Store(ctx context.Context, item *Milestone) error
	Update(ctx context.Context, item *Milestone) error
	Delete(ctx context.Context, id string) error
}

// MilestoneUsecase defines the business logic contract for milestones.
type MilestoneUsecase interface {
	Fetch(ctx context.Context) ([]Milestone, error)
	Store(ctx context.Context, item *Milestone) error
	Update(ctx context.Context, item *Milestone) error
	Delete(ctx context.Context, id string) error
}
