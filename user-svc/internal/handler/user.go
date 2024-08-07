package handler

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/storyofhis/auth-svc/internal/common"
	"github.com/storyofhis/auth-svc/internal/config"
	"github.com/storyofhis/auth-svc/internal/repository/model"
	"github.com/storyofhis/auth-svc/internal/service"
	pb "github.com/storyofhis/auth-svc/protos"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type Server struct {
	pb.UnimplementedUserServiceServer
	Service service.UserService
}

// Register
func (s *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	_, err := s.Service.GetUserByEmail(ctx, req.Email)
	if err == nil {
		log.Println("EMAIL_ALREADY_USED")
		return nil, status.Error(codes.AlreadyExists, "email already used")
	} else if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Println("INTERNAL_SERVER_ERROR : ", err)
		return nil, status.Error(codes.Internal, "internal server error")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("INTERNAL_SERVER_ERROR : ", err)
		return nil, err
	}

	// input
	newUser := model.User{
		Username: req.Username,
		FullName: req.Fullname,
		Email:    req.Email,
		Address:  req.Address,
		Password: string(hashedPassword),
		Phone:    req.PhoneNumber,
	}

	user, err := s.Service.CreateUser(ctx, &newUser)
	if err != nil {
		log.Println("INTERNAL_SERVER_ERROR : ", err)
		return nil, err
	}
	return &pb.CreateUserResponse{
		User: &pb.User{
			Id:          string(rune(user.ID)),
			Username:    user.Username,
			Fullname:    user.FullName,
			Email:       user.Email,
			Address:     user.Address,
			PhoneNumber: user.Username,
			Password:    user.Password,
		},
	}, nil
}

// Login
func (s *Server) LoginUser(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	user, err := s.Service.GetUserByEmail(ctx, req.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("STATUS_BAD_REQUEST : user not found")
			return nil, status.Error(codes.NotFound, "user not found")
		}
		log.Println("STATUS_INTERNAL_ERROR : ", err)
		return nil, status.Error(codes.Internal, "internal server error")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		log.Println("STATUS_BAD_REQUEST : incorrect password")
		return nil, status.Error(codes.Unauthenticated, "incorrect password")
	}

	claims := &common.CustomClaims{
		Id: int(user.ID),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * time.Duration(config.GetJwtExpiredTime())).Unix(),
			Subject:   user.Email,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(config.GetJwtSignature())
	if err != nil {
		log.Println("STATUS_INTERNAL_ERROR : could not generate token")
		return nil, status.Error(codes.Internal, "could not generate token")
	}

	return &pb.LoginResponse{
		Token: ss,
	}, nil
}
