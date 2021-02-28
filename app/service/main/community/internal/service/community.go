package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	pb "github.com/mittacy/himusic/service/community/api/community/v1"
	"github.com/mittacy/himusic/service/community/internal/biz"
)

type CommunityService struct {
	pb.UnimplementedCommunityServiceServer

	log       *log.Helper
	community *biz.CommunityUseCase
}

func NewCommunityService(useCase *biz.CommunityUseCase, logger log.Logger) *CommunityService {
	return &CommunityService{
		log:       log.NewHelper("community", logger),
		community: useCase,
	}
}

func (s *CommunityService) CreateCommunity(ctx context.Context,
	req *pb.CreateCommunityRequest) (*pb.CreateCommunityReply, error) {
	id, err := s.community.Create(ctx, &biz.Community{
		Name: req.Name,
		Icon: req.Icon,
	})
	if err != nil && errors.IsUnknown(err) {
		s.log.Errorf("%+v", err)
	}
	return &pb.CreateCommunityReply{Id: id}, err
}

func (s *CommunityService) DeleteCommunity(ctx context.Context,
	req *pb.DeleteCommunityRequest) (*pb.DeleteCommunityReply, error) {
	err := s.community.Delete(ctx, req.Id)
	if err != nil && errors.IsUnknown(err) {
		s.log.Errorf("%+v", err)
	}
	return &pb.DeleteCommunityReply{}, err
}

func (s *CommunityService) UpdateCommunity(ctx context.Context,
	req *pb.UpdateCommunityRequest) (*pb.UpdateCommunityReply, error) {
	err := s.community.Update(ctx, &biz.Community{
		Id:   req.Id,
		Name: req.Name,
		Icon: req.Icon,
	})
	if err != nil && errors.IsUnknown(err) {
		s.log.Errorf("%+v", err)
	}
	return &pb.UpdateCommunityReply{}, err
}

func (s *CommunityService) GetCommunity(ctx context.Context,
	req *pb.GetCommunityRequest) (*pb.GetCommunityReply, error) {
	community, err := s.community.Get(ctx, req.Id)
	if err != nil && errors.IsUnknown(err) {
		s.log.Errorf("%+v", err)
	}
	return &pb.GetCommunityReply{Community: &pb.Community{
		Id:   community.Id,
		Name: community.Name,
		Icon: community.Icon,
	}}, err
}

func (s *CommunityService) ListCommunity(ctx context.Context,
	req *pb.ListCommunityRequest) (*pb.ListCommunityReply, error) {
	ps, err := s.community.List(ctx)
	reply := &pb.ListCommunityReply{}
	if err != nil && errors.IsUnknown(err) {
		s.log.Errorf("%+v", err)
		return reply, err
	}
	for _, p := range ps {
		reply.Results = append(reply.Results, &pb.Community{
			Id:   p.Id,
			Name: p.Name,
		})
	}
	return reply, err
}
