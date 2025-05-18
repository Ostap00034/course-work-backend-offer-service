package main

import (
	"log"
	"net"
	"os"

	offerpbv1 "github.com/Ostap00034/course-work-backend-api-specs/gen/go/offer/v1"
	orderpbv1 "github.com/Ostap00034/course-work-backend-api-specs/gen/go/order/v1"
	userpbv1 "github.com/Ostap00034/course-work-backend-api-specs/gen/go/user/v1"

	"github.com/Ostap00034/course-work-backend-offer-service/db"
	offer "github.com/Ostap00034/course-work-backend-offer-service/internal"
	"github.com/joho/godotenv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	dbString, ok := os.LookupEnv("DB_CONN_STRING")
	if !ok {
		log.Fatal("DB_CONN_STRING is not set")
	}
	client := db.NewClient(dbString)
	defer client.Close()

	repo := offer.NewRepo(client)
	svc := offer.NewService(repo)

	userAddr, ok := os.LookupEnv("USER_SERVICE_ADDR")
	if !ok {
		log.Fatal("USER_SERVICE_ADDR is not set")
	}
	userConn, err := grpc.NewClient(
		userAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("failed to dial UserService: %v", err)
	}
	defer userConn.Close()
	userSvc := userpbv1.NewUserServiceClient(userConn)

	orderAddr, ok := os.LookupEnv("ORDER_SERVICE_ADDR")
	if !ok {
		log.Fatal("ORDER_SERVICE_ADDR is not set")
	}
	orderConn, err := grpc.NewClient(
		orderAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("failed to dial OrderService: %v", err)
	}
	defer orderConn.Close()
	orderSvc := orderpbv1.NewOrderServiceClient(orderConn)

	lis, err := net.Listen("tcp", ":50055")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcSrv := grpc.NewServer()
	srv := offer.NewServer(svc, userSvc, orderSvc)
	offerpbv1.RegisterOfferServiceServer(grpcSrv, srv)

	log.Println("OfferService is listening on :50055")
	log.Fatal(grpcSrv.Serve(lis))
}
