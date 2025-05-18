package offer

import (
	"context"

	commonpbv1 "github.com/Ostap00034/course-work-backend-api-specs/gen/go/common/v1"
	offerpbv1 "github.com/Ostap00034/course-work-backend-api-specs/gen/go/offer/v1"
	orderpbv1 "github.com/Ostap00034/course-work-backend-api-specs/gen/go/order/v1"
	userpbv1 "github.com/Ostap00034/course-work-backend-api-specs/gen/go/user/v1"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	offerpbv1.UnimplementedOfferServiceServer
	svc      Service
	userSvc  userpbv1.UserServiceClient
	orderSvc orderpbv1.OrderServiceClient
}

func NewServer(svc Service, userSvc userpbv1.UserServiceClient, orderSvc orderpbv1.OrderServiceClient) *Server {
	return &Server{svc: svc, userSvc: userSvc, orderSvc: orderSvc}
}

func (s *Server) CreateOffer(ctx context.Context, req *offerpbv1.CreateOfferRequest) (*offerpbv1.CreateOfferResponse, error) {
	order_id, err := uuid.Parse(req.OrderId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "неправильный формат UUID заказа")
	}

	order, err := s.orderSvc.GetOrderById(ctx, &orderpbv1.GetOrderByIdRequest{Id: req.OrderId})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	if order.Order.Status != "active" {
		return nil, status.Error(codes.InvalidArgument, "заказ не активен")
	}

	master_id, err := uuid.Parse(req.MasterId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "неправильный формат UUID исполнителя")
	}

	master, err := s.userSvc.GetUserById(ctx, &userpbv1.GetUserByIdRequest{UserId: req.MasterId})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	offer, err := s.svc.CreateOffer(ctx, order_id, master_id, req.Price)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &offerpbv1.CreateOfferResponse{Offer: &commonpbv1.OfferData{
		Id:        offer.ID.String(),
		Order:     order.Order,
		Master:    master.User,
		Status:    offer.Status.String(),
		Price:     offer.Price,
		CreatedAt: offer.CreatedAt.String(),
		UpdatedAt: offer.UpdatedAt.String(),
	}}, nil
}

func (s *Server) GetMyOrderOffers(ctx context.Context, req *offerpbv1.GetMyOrderOffersRequest) (*offerpbv1.GetMyOrderOffersResponse, error) {
	order_id, err := uuid.Parse(req.OrderId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "неправильный формат UUID заказа")
	}

	_, err = s.orderSvc.GetOrderById(ctx, &orderpbv1.GetOrderByIdRequest{Id: req.OrderId})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	offers, err := s.svc.GetMyOrderOffers(ctx, order_id)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	var offersData []*commonpbv1.OfferData
	for _, offer := range offers {
		offersData = append(offersData, &commonpbv1.OfferData{
			Id:        offer.ID.String(),
			Status:    offer.Status.String(),
			Price:     offer.Price,
			CreatedAt: offer.CreatedAt.String(),
			UpdatedAt: offer.UpdatedAt.String(),
		})
	}

	return &offerpbv1.GetMyOrderOffersResponse{Offers: offersData}, nil
}

func (s *Server) UpdateOffer(ctx context.Context, req *offerpbv1.UpdateOfferRequest) (*offerpbv1.UpdateOfferResponse, error) {
	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "неправильный формат UUID предложения")
	}

	offer, err := s.svc.UpdateOffer(ctx, id, req.Status)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if req.Status == "accepted" {
		order, err := s.orderSvc.GetOrderById(ctx, &orderpbv1.GetOrderByIdRequest{
			Id: offer.OrderID.String(),
		})
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}

		_, err = s.orderSvc.UpdateOrder(ctx, &orderpbv1.UpdateOrderRequest{
			Id:       offer.OrderID.String(),
			Status:   "in_progress",
			ClientId: order.Order.Client.Id,
			MasterId: offer.MasterID.String(),
		})
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	return &offerpbv1.UpdateOfferResponse{Offer: &commonpbv1.OfferData{
		Id:        offer.ID.String(),
		Status:    offer.Status.String(),
		Price:     offer.Price,
		CreatedAt: offer.CreatedAt.String(),
		UpdatedAt: offer.UpdatedAt.String(),
	}}, nil
}
