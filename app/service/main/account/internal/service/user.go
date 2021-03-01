package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	pb "github.com/mittacy/himusic/service/account/api/account/v1"
	"github.com/mittacy/himusic/service/account/internal/biz"
)

type UserService struct {
	pb.UnimplementedUserServer

	log  *log.Helper
	user *biz.UserUseCase
}

func NewUserService(useCase *biz.UserUseCase, logger log.Logger) *UserService {
	return &UserService{
		log:  log.NewHelper("user", logger),
		user: useCase,
	}
}

func (s *UserService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserReply, error) {
	user := &biz.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}
	err := s.user.CreateUser(ctx, user)
	if err != nil && errors.IsUnknown(err) {
		s.log.Errorf("%+v", err)
	}
	return &pb.CreateUserReply{
		Id: user.Id,
	}, err
}

func (s *UserService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserReply, error) {
	err := s.user.UpdateUser(ctx, &biz.User{
		Id:       req.Id,
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
		Icon:     req.Icon,
	})
	if err != nil && errors.IsUnknown(err) {
		s.log.Errorf("%+v", err)
	}
	return &pb.UpdateUserReply{}, err
}

func (s *UserService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserReply, error) {
	err := s.user.DeleteUser(ctx, req.Id)
	if err != nil && errors.IsUnknown(err) {
		s.log.Errorf("%+v", err)
	}
	return &pb.DeleteUserReply{}, err
}

func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserReply, error) {
	user, err := s.user.GetUser(ctx, req.Id)
	if err != nil && errors.IsUnknown(err) {
		s.log.Errorf("%+v", err)
	}
	return &pb.GetUserReply{User: &pb.UserStruct{
		Id:         user.Id,
		Name:       user.Name,
		Email:      user.Email,
		Icon:       user.Icon,
		PaperCount: user.PaperCount,
		BuyCount:   user.BuyCount,
		Coin:       user.Coin,
		Score:      user.Score,
		CreateTime: user.CreateTime,
		UpdateTime: user.UpdateTime,
		LoginTime:  user.LoginTime,
	}}, err
}

func (s *UserService) VerifyPasswordByEmail(ctx context.Context, req *pb.VerifyPasswordByEmailRequest) (*pb.VerifyPasswordReply, error) {
	isCorrect, err := s.user.VerifyPasswordByEmail(ctx, req.Email, req.Password)
	if err != nil && errors.IsUnknown(err) {
		s.log.Errorf("%+v", err)
	}
	return &pb.VerifyPasswordReply{Correct: isCorrect}, err
}

func (s *UserService) ListUser(ctx context.Context, req *pb.ListUserRequest) (*pb.ListUserReply, error) {
	reply := &pb.ListUserReply{}
	users, err := s.user.ListUser(ctx, int(req.PageNum), int(req.PageSize))
	if err != nil {
		if errors.IsUnknown(err) {
			s.log.Errorf("%+v", err)
		}
		return nil, err
	}
	for _, p := range users {
		reply.Users = append(reply.Users, &pb.UserStruct{
			Id:         p.Id,
			Name:       p.Name,
			Email:      p.Email,
			PaperCount: p.PaperCount,
			BuyCount:   p.BuyCount,
			Coin:       p.Coin,
			Score:      p.Score,
			CreateTime: p.CreateTime,
			UpdateTime: p.UpdateTime,
			LoginTime:  p.LoginTime,
		})
	}
	return reply, nil
}
