package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/mittacy/himusic/service/coins/internal/biz"

	pb "github.com/mittacy/himusic/service/coins/api/coins/v1"
)

type CoinsService struct {
	pb.UnimplementedCoinsServer

	log  *log.Helper
	user *biz.UserUseCase
}

func NewCoinsService(useCase *biz.UserUseCase, logger log.Logger) *CoinsService {
	return &CoinsService{
		log:  log.NewHelper("coin", logger),
		user: useCase,
	}
}

func (s *CoinsService) UpdateCoins(ctx context.Context, req *pb.UpdateCoinsRequest) (*pb.UpdateCoinsReply, error) {
	err := s.user.Update(ctx, req.Id, req.Coin)
	if err != nil && errors.IsUnknown(err) {
		s.log.Errorf("%+v", err)
	}
	return &pb.UpdateCoinsReply{}, err
}

func (s *CoinsService) GetCoins(ctx context.Context, req *pb.GetCoinsRequest) (*pb.GetCoinsReply, error) {
	coin, err := s.user.Get(ctx, req.Id)
	if err != nil && errors.IsUnknown(err) {
		s.log.Errorf("%+v", err)
	}
	return &pb.GetCoinsReply{Coin: coin}, err
}
