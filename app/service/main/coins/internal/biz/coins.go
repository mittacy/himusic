package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type User struct {
	Id   int32
	Coin int32
}

// TableName 覆盖gorm的映射表名
func (*User) TableName() string {
	return "ums_user"
}

type UserRepo interface {
	UpdateCoin(ctx context.Context, id int32, count int32) error
	GetCoin(ctx context.Context, id int32) (int32, error)
}

type UserUseCase struct {
	repo UserRepo
}

func NewUserUseCase(repo UserRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{repo: repo}
}

func (uc *UserUseCase) Update(ctx context.Context, id int32, count int32) error {
	return uc.repo.UpdateCoin(ctx, id, count)
}

func (uc *UserUseCase) Get(ctx context.Context, id int32) (int32, error) {
	return uc.repo.GetCoin(ctx, id)
}