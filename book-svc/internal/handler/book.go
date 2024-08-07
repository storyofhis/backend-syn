package handler

import (
	"context"
	"log"

	"github.com/storyofhis/book-svc/internal/repository/model"
	"github.com/storyofhis/book-svc/internal/service"
	pb "github.com/storyofhis/book-svc/protos"
)

type Server struct {
	pb.UnimplementedBookServiceServer
	Service service.BookService
}

func (s *Server) CreateBook(ctx context.Context, req *pb.CreateBookRequest) (*pb.BookResponse, error) {
	newBook := model.Book{
		Title:       req.Title,
		Author:      req.Author,
		Description: req.Description,
	}
	createdBook, err := s.Service.CreateBook(ctx, &newBook)
	if err != nil {
		log.Println("[ERROR] : ", err.Error())
		return nil, err
	}
	return &pb.BookResponse{
		Book: &pb.Book{
			Id:          string(createdBook.ID),
			Title:       createdBook.Title,
			Author:      createdBook.Author,
			Description: createdBook.Description,
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
