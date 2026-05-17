package repository

import (
	"context"
	"strings"
	"time"

	"backend-rs100/internal/domain"
	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

type milestoneFirestoreRepo struct {
	client *firestore.Client
}

// NewMilestoneFirestoreRepository creates a Firestore implementation of MilestoneRepository.
func NewMilestoneFirestoreRepository(client *firestore.Client) domain.MilestoneRepository {
	return &milestoneFirestoreRepo{
		client: client,
	}
}

func (r *milestoneFirestoreRepo) Fetch(ctx context.Context) ([]domain.Milestone, error) {
	var items []domain.Milestone
	
	// Order by year ascending for timeline
	iter := r.client.Collection("milestones").OrderBy("year", firestore.Asc).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		
		var item domain.Milestone
		doc.DataTo(&item)
		item.ID = doc.Ref.ID
		items = append(items, item)
	}
	
	return items, nil
}

func (r *milestoneFirestoreRepo) Store(ctx context.Context, item *domain.Milestone) error {
	item.CreatedAt = time.Now().Format(time.RFC3339)
	
	// Sanitize UTF-8 strings to prevent gRPC marshaling errors
	item.Year = strings.ToValidUTF8(item.Year, "")
	item.Title = strings.ToValidUTF8(item.Title, "")
	item.Description = strings.ToValidUTF8(item.Description, "")
	
	docRef, _, err := r.client.Collection("milestones").Add(ctx, map[string]interface{}{
		"year":        item.Year,
		"title":       item.Title,
		"description": item.Description,
		"created_at":  item.CreatedAt,
	})
	if err != nil {
		return err
	}
	
	item.ID = docRef.ID
	return nil
}

func (r *milestoneFirestoreRepo) Update(ctx context.Context, item *domain.Milestone) error {
	// Sanitize UTF-8 strings
	item.Year = strings.ToValidUTF8(item.Year, "")
	item.Title = strings.ToValidUTF8(item.Title, "")
	item.Description = strings.ToValidUTF8(item.Description, "")

	_, err := r.client.Collection("milestones").Doc(item.ID).Set(ctx, map[string]interface{}{
		"year":        item.Year,
		"title":       item.Title,
		"description": item.Description,
	}, firestore.MergeAll)
	return err
}

func (r *milestoneFirestoreRepo) Delete(ctx context.Context, id string) error {
	_, err := r.client.Collection("milestones").Doc(id).Delete(ctx)
	return err
}
