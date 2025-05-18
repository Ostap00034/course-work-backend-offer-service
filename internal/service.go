package offer

import (
	"context"

	"github.com/Ostap00034/course-work-backend-offer-service/ent"
	"github.com/google/uuid"
)

type Service interface {
	GetMyOrderOffers(ctx context.Context, order_id uuid.UUID) ([]*ent.Offer, error)
	CreateOffer(ctx context.Context, order_id, master_id uuid.UUID, price float32) (*ent.Offer, error)
	UpdateOffer(ctx context.Context, id uuid.UUID, status string) (*ent.Offer, error)
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{repo: r}
}

func (s *service) GetMyOrderOffers(ctx context.Context, order_id uuid.UUID) ([]*ent.Offer, error) {
	return s.repo.GetMyOrderOffers(ctx, order_id)
}

func (s *service) CreateOffer(ctx context.Context, order_id, master_id uuid.UUID, price float32) (*ent.Offer, error) {
	return s.repo.CreateOffer(ctx, order_id, master_id, price)
}

func (s *service) UpdateOffer(ctx context.Context, id uuid.UUID, status string) (*ent.Offer, error) {
	return s.repo.UpdateOffer(ctx, id, status)
}
