package controller

import (
	"context"
	"fmt"
	"strconv"

	"cbsr.io/golang-grpc-template/common/controller"
	"cbsr.io/golang-grpc-template/modules/users/application"
	"cbsr.io/golang-grpc-template/modules/users/models"
	"cbsr.io/golang-grpc-template/modules/users/views/requests"
	"cbsr.io/golang-grpc-template/proto/ping"
	"cbsr.io/golang-grpc-template/proto/users"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

var _ users.UserServiceServer = (*userServer)(nil)
var _ controller.IController = (*userServer)(nil)

type userServer struct {
	users.UnimplementedUserServiceServer
	logger      *logrus.Entry
	pingClient  ping.PingServiceClient
	userService application.IUserService
}

func New(l *logrus.Logger, client ping.PingServiceClient, svc application.IUserService) controller.IController {
	return &userServer{
		logger:      l.WithField("service", "users"),
		pingClient:  client,
		userService: svc,
	}
}

func (s *userServer) RegisterController(srv *grpc.Server) {
	users.RegisterUserServiceServer(srv, s)
}

// Only for testing purposes
func (s *userServer) Ping(ctx context.Context, req *users.PingRequest) (*users.PongResponse, error) {
	s.logger.Debug("Received Ping request")
	if req.Message == "" {
		s.logger.Error("message is required")
		return nil, status.Errorf(codes.InvalidArgument, "message is required")
	}

	r, err := s.pingClient.Ping(ctx, &ping.PingRequest{
		Message: req.Message,
	})
	if err != nil {
		return nil, err
	}

	return &users.PongResponse{Reply: r.GetReply()}, nil
}

func (s *userServer) CreateUser(ctx context.Context, req *users.CreateUserRequest) (*users.CreateUserResponse, error) {
	s.logger.Debug("Received CreateUser request")

	if err := requests.ValidateCreateUserRequest(req); err != nil {
		return nil, err
	}

	user, err := s.userService.CreateUser(ctx, req.GetUsername(), req.GetPassword(), req.GetEmail(), req.GetName())
	if err != nil {
		return nil, err
	}

	return &users.CreateUserResponse{
		Id:        fmt.Sprint(user.ID),
		Username:  user.Username,
		Email:     user.Email,
		Name:      user.Name,
		CreatedAt: user.CreatedAt.String(),
		UpdatedAt: user.UpdatedAt.String(),
	}, nil
}

func (s *userServer) GetUserByUsername(ctx context.Context, req *users.GetUserByUsernameRequest) (*users.GetUserResponse, error) {
	s.logger.Debug("Received GetUserByUsername request")

	if err := requests.ValidateGetUserByUsernameRequest(req); err != nil {
		return nil, err
	}

	user, err := s.userService.GetUserByUsername(ctx, req.GetUsername())
	if err != nil {
		return nil, err
	}

	return &users.GetUserResponse{
		Id:        fmt.Sprint(user.ID),
		Username:  user.Username,
		Email:     user.Email,
		Name:      user.Name,
		UpdatedAt: user.UpdatedAt.String(),
		CreatedAt: user.CreatedAt.String(),
	}, nil
}

func (s *userServer) GetUserByEmail(ctx context.Context, req *users.GetUserByEmailRequest) (*users.GetUserResponse, error) {
	s.logger.Debug("Received GetUserByEmail request")

	if err := requests.ValidateGetUserByEmailRequest(req); err != nil {
		return nil, err
	}

	user, err := s.userService.GetUserByEmail(ctx, req.GetEmail())
	if err != nil {
		return nil, err
	}

	return &users.GetUserResponse{
		Id:        fmt.Sprint(user.ID),
		Username:  user.Username,
		Email:     user.Email,
		Name:      user.Name,
		UpdatedAt: user.UpdatedAt.String(),
		CreatedAt: user.CreatedAt.String(),
	}, nil
}

func (s *userServer) GetUser(ctx context.Context, req *users.GetUserRequest) (*users.GetUserResponse, error) {
	s.logger.Debug("Received GetUser request")

	if err := requests.ValidateGetUserByIDRequest(req); err != nil {
		return nil, err
	}

	user, err := s.userService.GetUserByID(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &users.GetUserResponse{
		Id:        fmt.Sprint(user.ID),
		Username:  user.Username,
		Email:     user.Email,
		Name:      user.Name,
		UpdatedAt: user.UpdatedAt.String(),
		CreatedAt: user.CreatedAt.String(),
	}, nil
}

func (s *userServer) UpdateUser(ctx context.Context, req *users.UpdateUserRequest) (*users.UpdateUserResponse, error) {
	s.logger.Debug("Received UpdateUser request")

	if err := requests.ValidateUpdateUserRequest(req); err != nil {
		return nil, err
	}

	id, err := strconv.ParseUint(req.GetId(), 10, 64)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid id")
	}

	user := &models.User{
		Model: gorm.Model{
			ID: uint(id),
		},
		Username: req.GetUsername(),
		Email:    req.GetEmail(),
		Name:     req.GetName(),
	}

	if err := s.userService.UpdateUser(ctx, user); err != nil {
		return nil, err
	}

	return &users.UpdateUserResponse{
		Id:        fmt.Sprint(user.ID),
		Username:  user.Username,
		Email:     user.Email,
		Name:      user.Name,
		UpdatedAt: user.UpdatedAt.String(),
	}, nil
}

func (s *userServer) DeleteUser(ctx context.Context, req *users.DeleteUserRequest) (*users.DeleteUserResponse, error) {
	s.logger.Debug("Received DeleteUser request")

	if err := requests.ValidateDeleteUserRequest(req); err != nil {
		return nil, err
	}

	id, err := strconv.ParseUint(req.GetId(), 10, 64)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid id")
	}

	user := &models.User{
		Model: gorm.Model{
			ID: uint(id),
		},
	}

	if err := s.userService.DeleteUser(ctx, user); err != nil {
		return nil, err
	}

	return &users.DeleteUserResponse{
		Id:        fmt.Sprint(user.ID),
		DeletedAt: user.DeletedAt.Time.String(),
	}, nil
}
