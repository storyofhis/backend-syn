package handler

import (
	"context"
	"log"

	"github.com/storyofhis/author-service/internal/repository/model"
	"github.com/storyofhis/author-service/internal/service"
	"github.com/storyofhis/author-service/protos"
)

type Server struct {
	protos.UnimplementedAuthorServiceServer
	Service service.AuthorService
}

func (s *Server) CreateAuthor(ctx context.Context, req *protos.CreateAuthorRequest) (*protos.CreateAuthorResponse, error) {
	newAuthor := model.Author{
		Name:      req.Name,
		Biography: req.Biography,
	}
	createdAuthor, err := s.Service.CreateAuthor(ctx, &newAuthor)
	if err != nil {
		log.Println("[ERROR] : ", err.Error())
		return nil, err
	}

	return &protos.CreateAuthorResponse{
		Author: &protos.Author{
			Id:        string(createdAuthor.ID),
			Name:      req.Name,
			Biography: req.Biography,
		},
	}, nil
}

func (s *Server) GetAuthors(ctx context.Context, req *protos.GetAuthorsRequest) (*protos.GetAuthorsResponse, error) {
	authors, err := s.Service.GetAuthors(ctx)
	if err != nil {
		return nil, err
	}
	var pbAuthors []*protos.Author
	for _, author := range authors {
		pbAuthors = append(pbAuthors, &protos.Author{
			Id:        string(author.ID),
			Name:      author.Name,
			Biography: author.Biography,
		})
	}
	return &protos.GetAuthorsResponse{
		Authors: pbAuthors,
	}, nil
}

func (s *Server) GetAuthorById(ctx context.Context, req *protos.GetAuthorByIdRequest) (*protos.GetAuthorByIdResponse, error) {
	author, err := s.Service.GetAuthorById(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &protos.GetAuthorByIdResponse{
		Author: &protos.Author{
			Id:        string(author.ID),
			Name:      author.Name,
			Biography: author.Biography,
		},
	}, nil
}
