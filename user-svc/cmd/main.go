package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/joho/godotenv"
	"github.com/storyofhis/auth-svc/internal/config"
	"github.com/storyofhis/auth-svc/internal/handler"
	"github.com/storyofhis/auth-svc/internal/repository/gorm"
	"github.com/storyofhis/auth-svc/internal/service"
	pb "github.com/storyofhis/auth-svc/protos"
	"google.golang.org/grpc"
)

func runGRPCServer(grpcPort string, userService service.UserService) {
	listener, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &handler.Server{
		Service: userService,
	})

	log.Printf("gRPC server started on %s", grpcPort)
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func runGatewayServer(grpcPort, httpPort string) {
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := pb.RegisterUserServiceHandlerFromEndpoint(context.Background(), mux, grpcPort, opts)
	if err != nil {
		log.Fatalf("failed to start HTTP server: %v", err)
	}

	log.Printf("HTTP server started on %s", httpPort)
	if err := http.ListenAndServe(httpPort, mux); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	db, err := config.ConnectPostgresGORM()
	if err != nil {
		log.Fatalf("Failed to connect to database : %v", err)
	}

	log.Println("Database connection successfully")

	userRepo := gorm.NewUserRepo(db)
	userSvc := service.NewUserService(userRepo)

	grpcPort := ":50052"
	go runGRPCServer(grpcPort, userSvc)

	httpPort := ":8081"
	runGatewayServer(grpcPort, httpPort)
}
