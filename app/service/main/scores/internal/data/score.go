package data

import (
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/mittacy/himusic/service/scores/api/scores/v1/apierr"
	"github.com/mittacy/himusic/service/scores/internal/biz"
	"github.com/mittacy/tools/validator"
	errors3 "github.com/pkg/errors"
	"gorm.io/gorm"
)

func (sr *scoresRepo) UpdateScores(user *biz.User) error {
	if err := validator.ValidatorStruct(user); err != nil {
		return errors.InvalidArgument(apierr.Errors_FieldInvalid, "字段不符合要求")
	}

	var err error
	if user.Score < 0 {
		sr.data.db.Model(user).UpdateColumn("score", gorm.Expr("score - ?", -user.Score))
	} else {
		sr.data.db.Model(user).UpdateColumn("score", gorm.Expr("score + ?", user.Score))
	}
	if err != nil {
		if errors3.Is(err, gorm.ErrRecordNotFound) {
			err = errors.NotFound(apierr.Errors_UserNoFound, "用户不存在")
		} else {
			err = errors3.WithStack(errors.Unknown(apierr.Errors_MysqlErr, ""))
		}
	}
	return err
}

func (sr *scoresRepo) GetScores(id int32) (int32, error) {
	user := biz.User{}
	user.Id = id
	err := sr.data.db.Select("score").First(&user).Error
	if err != nil {
		if errors3.Is(err, gorm.ErrRecordNotFound) {
			err = errors.NotFound(apierr.Errors_UserNoFound, "用户不存在")
		} else {
			err = errors3.WithStack(errors.Unknown(apierr.Errors_MysqlErr, ""))
		}
	}
	return user.Score, err
}
