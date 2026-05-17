package domain

import "context"

// GalleryItem represents a photo and narration in the museum history gallery.
type GalleryItem struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Narration string `json:"narration"`
	PhotoPath string `json:"photoPath"` // Stores the GCS URL or local path
	CreatedAt string `json:"created_at"`
}

// GalleryRepository defines the data access contract for gallery items.
type GalleryRepository interface {
	Fetch(ctx context.Context) ([]GalleryItem, error)
	Store(ctx context.Context, item *GalleryItem) error
	Update(ctx context.Context, item *GalleryItem) error
	Delete(ctx context.Context, id string) error
}

// GalleryUsecase defines the business logic contract for gallery items.
type GalleryUsecase interface {
	Fetch(ctx context.Context) ([]GalleryItem, error)
	Store(ctx context.Context, item *GalleryItem) error
	Update(ctx context.Context, item *GalleryItem) error
	Delete(ctx context.Context, id string) error
}
