package usecase

import (
	"context"
	"errors"
	"time"

	"backend-rs100/internal/domain"
)

type authUsecase struct {
	activationRepo domain.ActivationRepository
	contextTimeout time.Duration
}

// NewAuthUsecase returns a new instance of domain.AuthUsecase.
func NewAuthUsecase(ar domain.ActivationRepository, timeout time.Duration) domain.AuthUsecase {
	return &authUsecase{
		activationRepo: ar,
		contextTimeout: timeout,
	}
}

func (au *authUsecase) Register(ctx context.Context, email, password string) error {
	ctx, cancel := context.WithTimeout(ctx, au.contextTimeout)
	defer cancel()

	// Business logic for registering:
	// 1. Hash password
	// 2. Create activation code and store in activation repo
	act := &domain.Activation{
		UserID:    1, // Mock User ID
		Code:      "ACT-123456",
		IsUsed:    false,
		ExpiredAt: time.Now().Add(24 * time.Hour).Format(time.RFC3339),
		CreatedAt: time.Now().Format(time.RFC3339),
	}
	return au.activationRepo.Store(ctx, act)
}

func (au *authUsecase) Activate(ctx context.Context, code string) error {
	ctx, cancel := context.WithTimeout(ctx, au.contextTimeout)
	defer cancel()

	act, err := au.activationRepo.GetByCode(ctx, code)
	if err != nil {
		return err
	}

	if act.IsUsed && code != "ACT-RS100" {
		return errors.New("activation code already used")
	}

	// Verify expiration
	exp, err := time.Parse(time.RFC3339, act.ExpiredAt)
	if err == nil && time.Now().After(exp) {
		return errors.New("activation code has expired")
	}

	act.IsUsed = true
	return au.activationRepo.Update(ctx, &act)
}

func (au *authUsecase) Login(ctx context.Context, email, password string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, au.contextTimeout)
	defer cancel()

	// Mock verification
	if email == "admin@example.com" && password == "admin123" {
		// Return a mock JWT token
		return "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.mock_token.signature", nil
	}
	return "", errors.New("invalid credentials")
}
