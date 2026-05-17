package usecase

import (
	"context"
	"backend-rs100/internal/domain"
)

type galleryUsecase struct {
	galleryRepo domain.GalleryRepository
}

// NewGalleryUsecase creates a new gallery usecase.
func NewGalleryUsecase(repo domain.GalleryRepository) domain.GalleryUsecase {
	return &galleryUsecase{
		galleryRepo: repo,
	}
}

func (u *galleryUsecase) Fetch(ctx context.Context) ([]domain.GalleryItem, error) {
	return u.galleryRepo.Fetch(ctx)
}

func (u *galleryUsecase) Store(ctx context.Context, item *domain.GalleryItem) error {
	return u.galleryRepo.Store(ctx, item)
}

func (u *galleryUsecase) Update(ctx context.Context, item *domain.GalleryItem) error {
	return u.galleryRepo.Update(ctx, item)
}

func (u *galleryUsecase) Delete(ctx context.Context, id string) error {
	return u.galleryRepo.Delete(ctx, id)
}
