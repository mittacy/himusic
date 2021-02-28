package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	apierr "github.com/mittacy/himusic/service/coins/api/coins/v1/apierr"
	"github.com/mittacy/himusic/service/coins/internal/biz"
	errors2 "github.com/pkg/errors"
	"gorm.io/gorm"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper("coin_repo", logger),
	}
}

func (ur *userRepo) UpdateCoin(ctx context.Context, id int32, count int32) error {
	if count == 0 {
		return nil
	}
	var user = biz.User{}
	user.Id = id
	if err := ur.data.db.Select("coin").First(&user).Error; err != nil {
		if errors2.Is(err, gorm.ErrRecordNotFound) {
			return errors.NotFound(apierr.Errors_UserNoFound, "用户不存在")
		}
	}

	user.Coin = user.Coin + count
	if user.Coin < 0 { // 查询硬币余额是否足够
		return errors.Unavailable(apierr.Errors_CoinsNoEnough, "硬币余额不足")
	}

	if err := ur.data.db.Model(&user).Update("coin", user.Coin).Error; err != nil {
		if errors2.Is(err, gorm.ErrRecordNotFound) {
			err = errors.NotFound(apierr.Errors_UserNoFound, "用户不存在")
		} else {
			err = errors2.WithStack(errors.Unknown(apierr.Errors_MysqlErr, ""))
		}
		return err
	}
	return nil
}

func (ur *userRepo) GetCoin(ctx context.Context, id int32) (int32, error) {
	user := biz.User{}
	user.Id = id
	err := ur.data.db.Select("coin").First(&user).Error
	if err != nil {
		if errors2.Is(err, gorm.ErrRecordNotFound) {
			err = errors.NotFound(apierr.Errors_UserNoFound, "用户不存在")
		} else {
			err = errors2.WithStack(errors.Unknown(apierr.Errors_MysqlErr, ""))
		}
	}
	return user.Coin, err
}
