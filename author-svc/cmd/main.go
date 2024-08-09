package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/joho/godotenv"
	"github.com/storyofhis/author-service/internal/config"
	"github.com/storyofhis/author-service/internal/handler"
	"github.com/storyofhis/author-service/internal/repository/gorm"
	"github.com/storyofhis/author-service/internal/service"
	"github.com/storyofhis/author-service/protos"
	"google.golang.org/grpc"
)

func runGRPCServer(grpcPort string, authorService service.AuthorService) {
	listener, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	protos.RegisterAuthorServiceServer(s, &handler.Server{
		Service: authorService,
	})

	log.Printf("gRPC server started on %s", grpcPort)
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func runGatewayServer(grpcPort, httpPort string) {
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := protos.RegisterAuthorServiceHandlerFromEndpoint(context.Background(), mux, grpcPort, opts)
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
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Database connection successful")
	// Initialize repository and service
	categoryRepo := gorm.NewAuthorRepo(db)
	bookService := service.NewAuthorService(categoryRepo)

	// Start gRPC server
	grpcPort := ":50054"
	go runGRPCServer(grpcPort, bookService)

	// Start gRPC-Gateway server
	httpPort := ":8083"
	runGatewayServer(grpcPort, httpPort)
}
