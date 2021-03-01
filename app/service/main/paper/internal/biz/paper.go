package biz

import (
	"context"
)

type Paper struct {
	Id          int32 `gorm:"primaryKey"`
	OwnerId     int32 `validate:"required"`
	CommunityId int32 `validate:"required"`
	State       int32
	Singer      string `validate:"required,max=64"`
	Song        string `validate:"required,max=64"`
	CreateTime  int64  `gorm:"autoCreateTime"`
	UpdateTime  int64  `gorm:"autoUpdateTime"`
	ReplyTime   int64
	Url         string `validate:"required,url,max=255"`
	Code        string `validate:"omitempty,max=16"`
	View        int32
	Like        int32
	Coin        int32
	ReplyCount  int32
}

func (*Paper) TableName() string {
	return "ms_paper"
}

type PaperRepo interface {
	CreatePaper(ctx context.Context, paper *Paper) (int32, error)
	DeletePaper(ctx context.Context, id int32) error
	UpdatePaper(ctx context.Context, paper *Paper) error
	GetPaper(ctx context.Context, id int32, columns []string) (*Paper, error)
	GetPaperCode(ctx context.Context, id int32) (string, error)
	ListPaper(ctx context.Context, communityId int32, pageNum, pageSize int, columns []string) ([]*Paper, error)
}

type PaperUseCase struct {
	repo PaperRepo
}

func NewPaperUseCase(repo PaperRepo) *PaperUseCase {
	return &PaperUseCase{repo: repo}
}

func (uc *PaperUseCase) CreatePaper(ctx context.Context, paper *Paper) (int32, error) {
	return uc.repo.CreatePaper(ctx, paper)
}

func (uc *PaperUseCase) DeletePaper(ctx context.Context, id int32) error {
	return uc.repo.DeletePaper(ctx, id)
}

func (uc *PaperUseCase) UpdatePaper(ctx context.Context, paper *Paper) error {
	return uc.repo.UpdatePaper(ctx, paper)
}

func (uc *PaperUseCase) GetPaper(ctx context.Context, id int32) (*Paper, error) {
	columns := []string{"id", "owner_id", "community_id", "state", "reply_count", "create_time", "update_time",
		"reply_time", "singer", "song", "url", "view", "like", "coin"}
	return uc.repo.GetPaper(ctx, id, columns)
}

func (uc *PaperUseCase) GetPaperCode(ctx context.Context, id int32) (string, error) {
	return uc.repo.GetPaperCode(ctx, id)
}

func (uc *PaperUseCase) ListPaper(ctx context.Context, communityId, pageNum, pageSize int32) ([]*Paper, error) {
	columns := []string{"id", "owner_id", "community_id", "state", "reply_count", "create_time", "update_time",
		"reply_time", "singer", "song", "url", "view", "like", "coin"}
	startIndex := (pageNum - 1) * pageSize
	return uc.repo.ListPaper(ctx, communityId, int(startIndex), int(pageSize), columns)
}
