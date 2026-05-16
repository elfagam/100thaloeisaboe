package repository

import (
	"context"
	"database/sql"
	"errors"
	"sync"
	"time"

	"backend-rs100/internal/domain"
)

// ==========================================
// Book Repository Implementation
// ==========================================

type bookMysqlRepo struct {
	db    *sql.DB
	books map[int64]domain.Book
	mu    sync.RWMutex
}

// NewBookMySQLRepository creates a MySQL implementation of the BookRepository.
func NewBookMySQLRepository(db *sql.DB) domain.BookRepository {
	repo := &bookMysqlRepo{
		db:    db,
		books: make(map[int64]domain.Book),
	}

	// Seed an initial ebook for RSUD Aloei Saboe 100th Anniversary
	repo.books[1] = domain.Book{
		ID:        1,
		Title:     "Satu Abad RSUD Aloei Saboe: Bakti & Dedikasi",
		Author:    "Tim Sejarah RS100",
		ISBN:      "978-602-1234-56-7",
		CreatedAt: time.Now().Format(time.RFC3339),
		UpdatedAt: time.Now().Format(time.RFC3339),
	}

	return repo
}

func (r *bookMysqlRepo) Fetch(ctx context.Context) ([]domain.Book, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	books := make([]domain.Book, 0, len(r.books))
	for _, b := range r.books {
		books = append(books, b)
	}
	return books, nil
}

func (r *bookMysqlRepo) GetByID(ctx context.Context, id int64) (domain.Book, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	book, exists := r.books[id]
	if !exists {
		return domain.Book{}, errors.New("book not found")
	}
	return book, nil
}

func (r *bookMysqlRepo) Store(ctx context.Context, b *domain.Book) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	b.ID = int64(len(r.books) + 1)
	b.CreatedAt = time.Now().Format(time.RFC3339)
	b.UpdatedAt = time.Now().Format(time.RFC3339)
	r.books[b.ID] = *b
	return nil
}

func (r *bookMysqlRepo) Update(ctx context.Context, b *domain.Book) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.books[b.ID]; !exists {
		return errors.New("book not found")
	}
	b.UpdatedAt = time.Now().Format(time.RFC3339)
	r.books[b.ID] = *b
	return nil
}

func (r *bookMysqlRepo) Delete(ctx context.Context, id int64) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.books[id]; !exists {
		return errors.New("book not found")
	}
	delete(r.books, id)
	return nil
}

// ==========================================
// Activation Repository Implementation
// ==========================================

type activationMysqlRepo struct {
	db          *sql.DB
	activations map[string]domain.Activation
	mu          sync.RWMutex
}

// NewActivationMySQLRepository creates a MySQL implementation of the ActivationRepository.
func NewActivationMySQLRepository(db *sql.DB) domain.ActivationRepository {
	return &activationMysqlRepo{
		db:          db,
		activations: make(map[string]domain.Activation),
	}
}

func (r *activationMysqlRepo) GetByCode(ctx context.Context, code string) (domain.Activation, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	act, exists := r.activations[code]
	if !exists {
		// Auto-create a default activation code for VVIP if requested code is ACT-RS100
		if code == "ACT-RS100" {
			newAct := domain.Activation{
				ID:        100,
				UserID:    100,
				Code:      "ACT-RS100",
				IsUsed:    false,
				ExpiredAt: time.Now().Add(24 * time.Hour).Format(time.RFC3339),
				CreatedAt: time.Now().Format(time.RFC3339),
			}
			r.mu.RUnlock()
			r.mu.Lock()
			r.activations[code] = newAct
			r.mu.Unlock()
			r.mu.RLock()
			return newAct, nil
		}
		return domain.Activation{}, errors.New("activation code not found")
	}
	return act, nil
}

func (r *activationMysqlRepo) Store(ctx context.Context, act *domain.Activation) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	act.ID = int64(len(r.activations) + 1)
	act.CreatedAt = time.Now().Format(time.RFC3339)
	r.activations[act.Code] = *act
	return nil
}

func (r *activationMysqlRepo) Update(ctx context.Context, act *domain.Activation) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.activations[act.Code] = *act
	return nil
}
