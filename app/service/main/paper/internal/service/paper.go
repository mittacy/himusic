package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/mittacy/himusic/service/paper/internal/biz"

	pb "github.com/mittacy/himusic/service/paper/api/paper/v1"
)

type PaperService struct {
	pb.UnimplementedPaperServer

	log   *log.Helper
	paper *biz.PaperUseCase
}

func NewPaperService(useCase *biz.PaperUseCase, logger log.Logger) *PaperService {
	return &PaperService{
		log:   log.NewHelper("paper", logger),
		paper: useCase,
	}
}

func (s *PaperService) CreatePaper(ctx context.Context, req *pb.CreatePaperRequest) (*pb.CreatePaperReply, error) {
	paper := biz.Paper{
		OwnerId:     req.OperatorId,
		CommunityId: req.CommunityId,
		Singer:      req.Singer,
		Song:        req.Song,
		Url:         req.Url,
		Code:        req.Code,
		Coin:        req.Coin,
	}
	id, err := s.paper.CreatePaper(ctx, &paper)
	if err != nil && errors.IsUnknown(err) {
		s.log.Errorf("%+v", err)
	}
	return &pb.CreatePaperReply{Id: id}, err
}

func (s *PaperService) UpdatePaper(ctx context.Context, req *pb.UpdatePaperRequest) (*pb.UpdatePaperReply, error) {
	paper := biz.Paper{
		Id:          req.Id,
		CommunityId: req.CommunityId,
		State:       req.State,
		Singer:      req.Singer,
		Song:        req.Song,
		Url:         req.Url,
		Code:        req.Code,
		Coin:        req.Coin,
		View:        req.View,
		Like:        req.Like,
		ReplyTime:   req.ReplyTime,
		ReplyCount:  req.ReplyCount,
	}
	err := s.paper.UpdatePaper(ctx, &paper)
	if err != nil && errors.IsUnknown(err) {
		s.log.Errorf("%+v", err)
	}
	return &pb.UpdatePaperReply{}, err
}

func (s *PaperService) DeletePaper(ctx context.Context, req *pb.DeletePaperRequest) (*pb.DeletePaperReply, error) {
	err := s.paper.DeletePaper(ctx, req.Id)
	if err != nil && errors.IsUnknown(err) {
		s.log.Errorf("%+v", err)
	}
	return &pb.DeletePaperReply{}, err
}

func (s *PaperService) GetPaper(ctx context.Context, req *pb.GetPaperRequest) (*pb.GetPaperReply, error) {
	paper, err := s.paper.GetPaper(ctx, req.GetId())
	if err != nil && errors.IsUnknown(err) {
		s.log.Errorf("%+v", err)
	}
	return &pb.GetPaperReply{
		Paper: &pb.PaperStruct{
			Id:          paper.Id,
			OwnerId:     paper.OwnerId,
			CommunityId: paper.CommunityId,
			State:       paper.State,
			ReplyCount:  paper.ReplyCount,
			CreateTime:  paper.CreateTime,
			UpdateTime:  paper.UpdateTime,
			ReplyTime:   paper.ReplyTime,
			Singer:      paper.Singer,
			Song:        paper.Song,
			Url:         paper.Url,
			View:        paper.View,
			Like:        paper.Like,
			Coin:        paper.Coin,
		},
	}, err
}

func (s *PaperService) GetPaperCode(ctx context.Context, req *pb.GetPaperCodeRequest) (*pb.GetPaperCodeReply, error) {
	code, err := s.paper.GetPaperCode(ctx, req.GetId())
	if err != nil && errors.IsUnknown(err) {
		s.log.Errorf("%+v", err)
	}
	return &pb.GetPaperCodeReply{Code: code}, err
}

func (s *PaperService) ListPaper(ctx context.Context, req *pb.ListPaperRequest) (*pb.ListPaperReply, error) {
	papers, err := s.paper.ListPaper(ctx, req.CommunityId, req.PageNum, req.PageSize)
	if err != nil {
		if errors.IsUnknown(err) {
			s.log.Errorf("%+v", err)
		}
		return &pb.ListPaperReply{}, err
	}
	reply := &pb.ListPaperReply{}
	for _, paper := range papers {
		reply.Papers = append(reply.Papers, &pb.PaperStruct{
			Id: paper.Id,
			OwnerId: paper.OwnerId,
			CommunityId: paper.CommunityId,
			State: paper.State,
			ReplyCount: paper.ReplyCount,
			CreateTime: paper.CreateTime,
			ReplyTime: paper.ReplyTime,
			Singer: paper.Singer,
			Song: paper.Song,
			View: paper.View,
			Like: paper.Like,
		})
	}
	return reply, nil
}
