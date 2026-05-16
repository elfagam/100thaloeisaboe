package domain

import "context"

// Activation represents the account activation entity.
type Activation struct {
	ID        int64  `json:"id"`
	UserID    int64  `json:"user_id"`
	Code      string `json:"code"`
	IsUsed    bool   `json:"is_used"`
	ExpiredAt string `json:"expired_at"`
	CreatedAt string `json:"created_at"`
}

// ActivationRepository defines the data access contract for account activation.
type ActivationRepository interface {
	GetByCode(ctx context.Context, code string) (Activation, error)
	Store(ctx context.Context, act *Activation) error
	Update(ctx context.Context, act *Activation) error
}

// AuthUsecase defines the business logic contract for authentication and activation.
type AuthUsecase interface {
	Register(ctx context.Context, email, password string) error
	Activate(ctx context.Context, code string) error
	Login(ctx context.Context, email, password string) (string, error) // Returns JWT Token
}
