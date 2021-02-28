package biz

import (
	"context"
)

type Community struct {
	Id   int32  `gorm:"primaryKey"`
	Name string `validate:"omitempty,min=1,max=10"`
	Icon string `validate:"omitempty,url,max=255"`
}

// TableName 覆盖gorm的映射表名
func (*Community) TableName() string {
	return "cg_community"
}

type CommunityRepo interface {
	// db
	CreateCommunity(ctx context.Context, community *Community) error
	DeleteCommunity(ctx context.Context, id int32) error
	UpdateCommunity(ctx context.Context, community *Community) error
	GetCommunity(ctx context.Context, id int32) (*Community, error)
	ListCommunity(ctx context.Context) ([]*Community, error)
}

type CommunityUseCase struct {
	repo CommunityRepo
}

func NewCommunityUseCase(repo CommunityRepo) *CommunityUseCase {
	return &CommunityUseCase{repo: repo}
}

func (uc *CommunityUseCase) Create(ctx context.Context, community *Community) (int32, error) {
	return community.Id, uc.repo.CreateCommunity(ctx, community)
}

func (uc *CommunityUseCase) Delete(ctx context.Context, id int32) error {
	return uc.repo.DeleteCommunity(ctx, id)
}

func (uc *CommunityUseCase) Update(ctx context.Context, community *Community) error {
	return uc.repo.UpdateCommunity(ctx, community)
}

func (uc *CommunityUseCase) List(ctx context.Context) (ps []*Community, err error) {
	return uc.repo.ListCommunity(ctx)
}

func (uc *CommunityUseCase) Get(ctx context.Context, id int32) (ps *Community, err error) {
	return uc.repo.GetCommunity(ctx, id)
}