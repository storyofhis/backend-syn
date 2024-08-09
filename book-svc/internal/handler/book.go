package handler

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/storyofhis/book-svc/internal/handler/category"
	"github.com/storyofhis/book-svc/internal/repository/model"
	"github.com/storyofhis/book-svc/internal/service"
	pb "github.com/storyofhis/book-svc/protos"
)

type Server struct {
	pb.UnimplementedBookServiceServer
	Service service.BookService
}

func (s *Server) CreateBook(ctx context.Context, req *pb.CreateBookRequest) (*pb.BookResponse, error) {
	// No need to convert CategoryId to int if you store it as a string
	categoryID := req.CategoryId

	// Fetch the category from category service (assuming category service expects an int ID)
	categoryIDInt, err := strconv.Atoi(categoryID)
	if err != nil {
		log.Println("[ERROR] : Invalid category ID")
		return nil, err
	}

	category, err := category.GetCategoryById(categoryIDInt)
	if err != nil {
		log.Println("[ERROR] : Cannot find the category", err)
		return nil, err
	}

	// Check if category is nil
	if category == nil {
		log.Println("[ERROR] : Category is nil")
		return nil, fmt.Errorf("category not found")
	}

	fmt.Println("Category Name:", category.Name)

	// Create a new book record with the received details
	newBook := model.Book{
		Title:       req.Title,
		Author:      req.Author,
		Description: req.Description,
		CategoryId:  categoryID, // Keep CategoryId as string
		Stock:       req.Stock,
	}

	createdBook, err := s.Service.CreateBook(ctx, &newBook)
	if err != nil {
		log.Println("[ERROR] : ", err.Error())
		return nil, err
	}

	return &pb.BookResponse{
		Book: &pb.Book{
			Id:          string(createdBook.ID), // Assuming ID is a string
			Title:       createdBook.Title,
			Author:      createdBook.Author,
			Description: createdBook.Description,
			CategoryId:  createdBook.CategoryId, // No conversion needed
			Stock:       createdBook.Stock,
		},
	}, nil
}

func (s *Server) GetBooks(ctx context.Context, req *pb.GetBooksRequest) (*pb.GetBooksResponse, error) {
	books, err := s.Service.ListBooks(ctx)
	if err != nil {
		return nil, err
	}
	var pbBooks []*pb.Book
	for _, book := range books {
		pbBooks = append(pbBooks, &pb.Book{
			Id:          string(book.ID),
			Title:       book.Title,
			Author:      book.Author,
			Description: book.Description,
			CategoryId:  book.CategoryId,
			Stock:       book.Stock,
		})
	}
	return &pb.GetBooksResponse{
		Books: pbBooks,
	}, nil
}

func (s *Server) GetBookById(ctx context.Context, req *pb.GetBookByIDRequest) (*pb.BookResponse, error) {
	book, err := s.Service.GetBookById(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.BookResponse{
		Book: &pb.Book{
			Id:          string(book.ID),
			Title:       book.Title,
			Author:      book.Author,
			Description: book.Description,
		},
	}, nil
}
