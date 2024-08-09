package handler

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/golang/protobuf/ptypes"
	"github.com/storyofhis/auth-svc/internal/handler/book"
	"github.com/storyofhis/auth-svc/internal/repository/model"
	"github.com/storyofhis/auth-svc/internal/service"
	"github.com/storyofhis/auth-svc/protos"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type BorrowServer struct {
	protos.UnimplementedBorrowServiceServer
	Service service.BorrowService
}

func (s *BorrowServer) CreateBorrow(ctx context.Context, req *protos.CreateBorrowRequest) (*protos.CreateBorrowResponse, error) {
	// Convert UserId from string to int
	userId, err := strconv.Atoi(req.UserId)
	if err != nil {
		return nil, fmt.Errorf("invalid user ID: %v", err)
	}

	bookID, err := strconv.Atoi(req.BookId) // Convert BookId from string to int
	if err != nil {
		return nil, fmt.Errorf("invalid book ID: %v", err)
	}

	book, err := book.GetBookById(bookID)
	if err != nil {
		log.Println("[ERROR] : Cannot find the book", err)
		return nil, err
	}

	// Check if book is nil
	if book == nil {
		log.Println("[ERROR] : Book is nil")
		return nil, fmt.Errorf("book not found")
	}

	fmt.Println("Book Name:", book.Title)

	newBorrow := model.Borrow{
		UserId:     userId,
		BookId:     bookID, // Use bookID (int) here
		BorrowDate: req.BorrowDate.AsTime(),
		ReturnDate: req.ReturnDate.AsTime(),
		Status:     req.Status,
	}

	createdBorrow, err := s.Service.CreateBorrow(ctx, &newBorrow)
	if err != nil {
		log.Println("[ERROR] : ", err.Error())
		return nil, err
	}

	// Convert time.Time to *timestamppb.Timestamp
	borrowDatePb := timestamppb.New(createdBorrow.BorrowDate)
	returnDatePb := timestamppb.New(createdBorrow.ReturnDate)

	return &protos.CreateBorrowResponse{
		Borrow: &protos.Borrow{
			Id:         fmt.Sprintf("%d", createdBorrow.ID), // Assuming ID is an int, convert to string
			UserId:     fmt.Sprintf("%d", createdBorrow.UserId),
			BookId:     fmt.Sprintf("%d", createdBorrow.BookId),
			BorrowDate: borrowDatePb,
			ReturnDate: returnDatePb,
			Status:     req.Status,
		},
	}, nil
}
func (s *BorrowServer) GetBorrows(ctx context.Context, req *protos.GetBorrowsRequest) (*protos.GetBorrowsResponse, error) {
	borrows, err := s.Service.GetBorrows(ctx)
	if err != nil {
		return nil, err
	}

	var pbBorrows []*protos.Borrow
	for _, borrow := range borrows {
		borrowDatePb, err := ptypes.TimestampProto(borrow.BorrowDate)
		if err != nil {
			return nil, fmt.Errorf("failed to convert BorrowDate: %v", err)
		}
		returrnDatePb, err := ptypes.TimestampProto(borrow.ReturnDate)
		if err != nil {
			return nil, fmt.Errorf("failed to convert ReturnDate: %v", err)
		}
		pbBorrows = append(pbBorrows, &protos.Borrow{
			Id:         string(borrow.ID),
			UserId:     string(borrow.UserId),
			BookId:     string(borrow.BookId),
			BorrowDate: borrowDatePb,
			ReturnDate: returrnDatePb,
		})
	}

	return &protos.GetBorrowsResponse{
		Borrows: pbBorrows,
	}, nil
}

func (s *BorrowServer) GetBorrowById(ctx context.Context, req *protos.GetBorrowByIdRequest) (*protos.GetBorrowByIdResponse, error) {
	borrow, err := s.Service.GetBorrowById(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	borrowDatePb, err := ptypes.TimestampProto(borrow.BorrowDate)
	if err != nil {
		return nil, fmt.Errorf("failed to convert BorrowDate: %v", err)
	}
	returrnDatePb, err := ptypes.TimestampProto(borrow.ReturnDate)
	if err != nil {
		return nil, fmt.Errorf("failed to convert ReturnDate: %v", err)
	}
	return &protos.GetBorrowByIdResponse{
		Borrow: &protos.Borrow{
			Id:         string(req.Id),
			UserId:     string(borrow.UserId),
			BookId:     string(borrow.BookId),
			BorrowDate: borrowDatePb,
			ReturnDate: returrnDatePb,
		},
	}, nil
}
