package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	apierr "github.com/mittacy/himusic/service/community/api/community/v1/errors"
	"github.com/mittacy/himusic/service/community/internal/biz"
	"github.com/mittacy/tools/validator"
	errors3 "github.com/pkg/errors"
	"gorm.io/gorm"
)

type communityRepo struct {
	data *Data
	log  *log.Helper
}

func NewCommunityRepo(data *Data, logger log.Logger) biz.CommunityRepo {
	return &communityRepo{
		data: data,
		log:  log.NewHelper("community_repo", logger),
	}
}

func (cr *communityRepo) CreateCommunity(ctx context.Context, community *biz.Community) error {
	if err := validator.ValidatorStruct(community); err != nil {
		return errors.InvalidArgument(apierr.Errors_FieldInvalid, err.Error())
	}
	r := 0
	err := cr.data.db.Raw("select 1 from "+community.TableName()+
		" where name = ? limit 1", community.Name).Scan(&r).Error
	if err != nil {
		return errors3.WithStack(errors.Unknown(apierr.Errors_MysqlErr, ""))
	}
	if r == 1 {
		return errors.InvalidArgument(apierr.Errors_NameExists, "名字冲突")
	}
	if err = cr.data.db.Create(&community).Error; err != nil {
		return errors3.WithStack(errors.Unknown(apierr.Errors_MysqlErr, ""))
	}
	return nil
}

func (cr *communityRepo) DeleteCommunity(ctx context.Context, id int32) error {
	if err := cr.data.db.Delete(&biz.Community{}, id).Error; err != nil {
		return errors3.WithStack(errors.Unknown(apierr.Errors_MysqlErr, ""))
	}
	return nil
}

func (cr *communityRepo) UpdateCommunity(ctx context.Context, community *biz.Community) error {
	if err := validator.ValidatorStruct(community); err != nil {
		return errors.InvalidArgument(apierr.Errors_FieldInvalid, err.Error())
	}
	var id int32
	err := cr.data.db.Raw("select id from "+community.TableName()+
		" where name = ?", community.Name).Scan(&id).Error
	if err != nil {
		return errors3.WithStack(errors.Unknown(apierr.Errors_MysqlErr, ""))
	}
	if id != community.Id {
		return errors.InvalidArgument(apierr.Errors_NameExists, "社区已存在")
	}
	if err = cr.data.db.Model(community).Updates(community).Error; err != nil {
		if errors3.Is(err, gorm.ErrRecordNotFound) {
			err = errors.InvalidArgument(apierr.Errors_CommunityNoFound, "社区不存在")
		} else {
			err = errors3.WithStack(errors.Unknown(apierr.Errors_MysqlErr, ""))
		}
		return err
	}
	return nil
}

func (cr *communityRepo) GetCommunity(ctx context.Context, id int32) (*biz.Community, error) {
	community := biz.Community{
		Id: id,
	}
	err := cr.data.db.First(&community).Error
	if errors3.Is(err, gorm.ErrRecordNotFound) {
		err = errors.InvalidArgument(apierr.Errors_CommunityNoFound, "社区不存在")
	} else if err != nil {
		err = errors3.WithStack(errors.Unknown(apierr.Errors_MysqlErr, ""))
	}
	return &community, err
}

func (cr *communityRepo) ListCommunity(ctx context.Context) (cs []*biz.Community, err error) {
	err = cr.data.db.Select("id", "name").Find(&cs).Error
	if err != nil {
		err = errors3.WithStack(errors.Unknown(apierr.Errors_MysqlErr, ""))
	}
	return
}
