package domain

import "context"

// Book represents the book entity in the system.
type Book struct {
	ID        int64  `json:"id"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	ISBN      string `json:"isbn"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// BookRepository defines the data access contract for books.
type BookRepository interface {
	Fetch(ctx context.Context) ([]Book, error)
	GetByID(ctx context.Context, id int64) (Book, error)
	Store(ctx context.Context, b *Book) error
	Update(ctx context.Context, b *Book) error
	Delete(ctx context.Context, id int64) error
}

// BookUsecase defines the business logic contract for books.
type BookUsecase interface {
	Fetch(ctx context.Context) ([]Book, error)
	GetByID(ctx context.Context, id int64) (Book, error)
	Store(ctx context.Context, b *Book) error
	Update(ctx context.Context, b *Book) error
	Delete(ctx context.Context, id int64) error
}
