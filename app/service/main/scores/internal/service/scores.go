
package service

import(
	"context"
	"github.com/mittacy/himusic/service/scores/internal/biz"

	pb "github.com/mittacy/himusic/service/scores/api/scores/v1"
)

type ScoresService struct {
	pb.UnimplementedScoresServer

	scores *biz.ScoresUseCase
}

func NewScoresService(useCase *biz.ScoresUseCase) *ScoresService {
	return &ScoresService{scores: useCase}
}

func (s *ScoresService) UpdateScores(ctx context.Context, req *pb.UpdateScoresRequest) (*pb.UpdateScoresReply, error) {
	err := s.scores.UpdateScores(&biz.User{
		Id:    req.Id,
		Score: req.Score,
	})
	return &pb.UpdateScoresReply{}, err
}

func (s *ScoresService) GetScores(ctx context.Context, req *pb.GetScoresRequest) (*pb.GetScoresReply, error) {
	score, err := s.scores.GetScores(req.Id)
	return &pb.GetScoresReply{Score: score}, err
}
