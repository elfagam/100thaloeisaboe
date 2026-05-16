package usecase

import (
	"context"
	"time"

	"backend-rs100/internal/domain"
)

type bookUsecase struct {
	bookRepo       domain.BookRepository
	contextTimeout time.Duration
}

// NewBookUsecase returns a new instance of domain.BookUsecase.
func NewBookUsecase(br domain.BookRepository, timeout time.Duration) domain.BookUsecase {
	return &bookUsecase{
		bookRepo:       br,
		contextTimeout: timeout,
	}
}

func (bu *bookUsecase) Fetch(ctx context.Context) ([]domain.Book, error) {
	ctx, cancel := context.WithTimeout(ctx, bu.contextTimeout)
	defer cancel()
	return bu.bookRepo.Fetch(ctx)
}

func (bu *bookUsecase) GetByID(ctx context.Context, id int64) (domain.Book, error) {
	ctx, cancel := context.WithTimeout(ctx, bu.contextTimeout)
	defer cancel()
	return bu.bookRepo.GetByID(ctx, id)
}

func (bu *bookUsecase) Store(ctx context.Context, b *domain.Book) error {
	ctx, cancel := context.WithTimeout(ctx, bu.contextTimeout)
	defer cancel()
	return bu.bookRepo.Store(ctx, b)
}

func (bu *bookUsecase) Update(ctx context.Context, b *domain.Book) error {
	ctx, cancel := context.WithTimeout(ctx, bu.contextTimeout)
	defer cancel()
	return bu.bookRepo.Update(ctx, b)
}

func (bu *bookUsecase) Delete(ctx context.Context, id int64) error {
	ctx, cancel := context.WithTimeout(ctx, bu.contextTimeout)
	defer cancel()
	return bu.bookRepo.Delete(ctx, id)
}
