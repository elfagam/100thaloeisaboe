package repository

import (
	"context"
	"strings"
	"time"

	"backend-rs100/internal/domain"
	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

type galleryFirestoreRepo struct {
	client *firestore.Client
}

// NewGalleryFirestoreRepository creates a Firestore implementation of GalleryRepository.
func NewGalleryFirestoreRepository(client *firestore.Client) domain.GalleryRepository {
	return &galleryFirestoreRepo{
		client: client,
	}
}

func (r *galleryFirestoreRepo) Fetch(ctx context.Context) ([]domain.GalleryItem, error) {
	var items []domain.GalleryItem
	
	// Get all documents from "gallery" collection ordered by created_at descending
	iter := r.client.Collection("gallery").OrderBy("created_at", firestore.Desc).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		
		var item domain.GalleryItem
		doc.DataTo(&item)
		item.ID = doc.Ref.ID
		items = append(items, item)
	}
	
	return items, nil
}

func (r *galleryFirestoreRepo) Store(ctx context.Context, item *domain.GalleryItem) error {
	item.CreatedAt = time.Now().Format(time.RFC3339)
	
	// Sanitize UTF-8 strings
	item.Title = strings.ToValidUTF8(item.Title, "")
	item.Narration = strings.ToValidUTF8(item.Narration, "")
	item.PhotoPath = strings.ToValidUTF8(item.PhotoPath, "")
	
	// Add a new document with generated ID
	docRef, _, err := r.client.Collection("gallery").Add(ctx, map[string]interface{}{
		"title":      item.Title,
		"narration":  item.Narration,
		"photoPath":  item.PhotoPath,
		"created_at": item.CreatedAt,
	})
	if err != nil {
		return err
	}
	
	item.ID = docRef.ID
	return nil
}

func (r *galleryFirestoreRepo) Update(ctx context.Context, item *domain.GalleryItem) error {
	// Sanitize UTF-8 strings
	item.Title = strings.ToValidUTF8(item.Title, "")
	item.Narration = strings.ToValidUTF8(item.Narration, "")
	item.PhotoPath = strings.ToValidUTF8(item.PhotoPath, "")

	_, err := r.client.Collection("gallery").Doc(item.ID).Set(ctx, map[string]interface{}{
		"title":     item.Title,
		"narration": item.Narration,
		"photoPath": item.PhotoPath,
	}, firestore.MergeAll)
	return err
}

func (r *galleryFirestoreRepo) Delete(ctx context.Context, id string) error {
	_, err := r.client.Collection("gallery").Doc(id).Delete(ctx)
	return err
}
