package offer

import (
	"context"
	"errors"

	"github.com/Ostap00034/course-work-backend-offer-service/ent"
	"github.com/Ostap00034/course-work-backend-offer-service/ent/offer"
	"github.com/google/uuid"
)

var (
	ErrOfferAlreadyExists     = errors.New("такая заявка уже существует")
	ErrOfferCreateFailed      = errors.New("ошибка при создании заявки")
	ErrOfferNotFound          = errors.New("заявка не найдена")
	ErrUpdateOfferFailed      = errors.New("ошибка при обновлении заявки")
	ErrGetMyOrderOffersFailed = errors.New("ошибка при получении заявок")
)

type Repository interface {
	GetMyOrderOffers(ctx context.Context, order_id uuid.UUID) ([]*ent.Offer, error)
	CreateOffer(ctx context.Context, order_id, master_id uuid.UUID, price float32) (*ent.Offer, error)
	UpdateOffer(ctx context.Context, id uuid.UUID, status string) (*ent.Offer, error)
}

type repo struct {
	client *ent.Client
}

func NewRepo(client *ent.Client) Repository {
	return &repo{client: client}
}

func (r *repo) CreateOffer(ctx context.Context, order_id, master_id uuid.UUID, price float32) (*ent.Offer, error) {
	offer, err := r.client.Offer.Create().
		SetOrderID(order_id).
		SetMasterID(master_id).
		SetPrice(price).
		Save(ctx)
	if err != nil {
		if ent.IsConstraintError(err) {
			return nil, ErrOfferAlreadyExists
		}
		return nil, ErrOfferCreateFailed
	}
	return offer, nil
}

func (r *repo) UpdateOffer(ctx context.Context, id uuid.UUID, status string) (*ent.Offer, error) {
	builder := r.client.Offer.UpdateOneID(id)

	if status != "" {
		builder = builder.SetStatus(offer.Status(status))
	}
	updated, err := builder.Save(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, ErrOfferNotFound
		}
		return nil, ErrUpdateOfferFailed
	}

	return updated, nil
}

func (r *repo) GetMyOrderOffers(ctx context.Context, order_id uuid.UUID) ([]*ent.Offer, error) {
	offers, err := r.client.Offer.Query().
		Where(offer.OrderIDEQ(order_id)).
		All(ctx)
	if err != nil {
		return nil, ErrGetMyOrderOffersFailed
	}
	return offers, nil
}
