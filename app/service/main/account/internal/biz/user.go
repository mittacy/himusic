package biz

import (
	"context"
	"github.com/mittacy/tools/encryption"
	"strconv"
)

type User struct {
	Id         int32 `gorm:"primaryKey"`
	CreateTime int64 `gorm:"autoCreateTime"`
	UpdateTime int64 `gorm:"autoUpdateTime"`
	LoginTime  int64
	Name       string `validate:"required,min=1,max=20"`
	Email      string `validate:"required,email"`
	Password   string `validate:"required,min=1,max=32"`
	Salt       string
	Icon       string `validate:"omitempty,url"`
	PaperCount int32
	BuyCount   int32
	Coin       int32
	Score      int32
}

func (*User) TableName() string {
	return "ums_user"
}

type UserRepo interface {
	CreateUser(ctx context.Context, user *User) error
	DeleteUser(ctx context.Context, id int32) error
	UpdateUser(ctx context.Context, user *User) error
	GetUser(ctx context.Context, field, fieldVal string, column []string) (*User, error)
	ListUser(ctx context.Context, startIndex, num int, column []string) ([]*User, error)
}

type UserUseCase struct {
	repo UserRepo
}

func NewUserUseRepo(repo UserRepo) *UserUseCase {
	return &UserUseCase{repo: repo}
}

func (uc *UserUseCase) CreateUser(ctx context.Context, user *User) error {
	return uc.repo.CreateUser(ctx, user)
}

func (uc *UserUseCase) DeleteUser(ctx context.Context,id int32) error {
	return uc.repo.DeleteUser(ctx, id)
}

func (uc *UserUseCase) UpdateUser(ctx context.Context, user *User) error {
	return uc.repo.UpdateUser(ctx, user)
}

func (uc *UserUseCase) VerifyPasswordByEmail(ctx context.Context, email string, password string) (bool, error) {
	columns := []string{"password", "salt"}
	user, err := uc.repo.GetUser(ctx, "email", email, columns)
	if err != nil {
		return false, err
	}
	if encryption.EncryptionBySalt(password, user.Salt) == user.Password {
		return true, nil
	}
	return false, nil
}

func (uc *UserUseCase) GetUser(context context.Context, id int32) (*User, error) {
	columns := []string{"id", "name", "email", "icon", "paper_count", "buy_count", "coin",
		"score", "create_time", "update_time", "login_time"}
	return uc.repo.GetUser(context, "id", strconv.Itoa(int(id)), columns)
}

func (uc *UserUseCase) ListUser(ctx context.Context, pageNum, pageSize int) ([]*User, error) {
	columns := []string{"id", "name", "email", "paper_count", "buy_count",
		"coin", "score", "create_time", "update_time", "login_time"}
	return uc.repo.ListUser(ctx, (pageNum - 1) * pageSize, pageSize, columns)
}
