package handler

import (
	"context"
	"log"

	"github.com/storyofhis/category-svc/internal/repository/model"
	"github.com/storyofhis/category-svc/internal/service"
	pb "github.com/storyofhis/category-svc/protos"
)

type Server struct {
	pb.UnimplementedCategoryServiceServer
	Service service.CategoryService
}

func (s *Server) CreateCategory(ctx context.Context, req *pb.CreateCategoryRequest) (*pb.CreateCategoryResponse, error) {
	newCategoryy := model.Category{
		Name:        req.Name,
		Description: req.Description,
	}

	createdCategory, err := s.Service.CreateCategory(ctx, &newCategoryy)
	if err != nil {
		log.Println("[ERROR] : ", err.Error())
		return nil, err
	}

	return &pb.CreateCategoryResponse{
		Category: &pb.Category{
			Id:          string(createdCategory.ID),
			Name:        createdCategory.Name,
			Description: createdCategory.Description,
		},
	}, nil
}

func (s *Server) GetCategories(ctx context.Context, req *pb.GetCategoryRequest) (*pb.GetCategoryResponse, error) {
	categories, err := s.Service.GetCategories(ctx)
	if err != nil {
		return nil, err
	}
	var pbCategories []*pb.Category
	for _, category := range categories {
		pbCategories = append(pbCategories, &pb.Category{
			Id:          string(category.ID),
			Name:        category.Name,
			Description: category.Description,
		})
	}
	return &pb.GetCategoryResponse{
		Categories: pbCategories,
	}, nil
}

func (s *Server) GetCategoryById(ctx context.Context, req *pb.GetCategoryByIdRequest) (*pb.GetCategoryByIdResponse, error) {
	category, err := s.Service.GetCategoryById(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.GetCategoryByIdResponse{
		Category: &pb.Category{
			Id:          string(category.ID),
			Name:        category.Name,
			Description: category.Description,
		},
	}, nil
}
