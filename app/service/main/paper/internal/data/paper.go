package data

import (
	"context"
	errors3 "errors"
	"fmt"
	"github.com/go-kratos/kratos/v2/errors"
	apierr "github.com/mittacy/himusic/service/paper/api/paper/v1/apierr"
	"github.com/mittacy/himusic/service/paper/internal/biz"
	"github.com/mittacy/tools/validator"
	errors2 "github.com/pkg/errors"
	"gorm.io/gorm"
)

type paperRepo struct {
	data *Data
}

func NewPaperRepo(data *Data) biz.PaperRepo {
	return &paperRepo{data: data}
}

func (pr *paperRepo) CreatePaper(ctx context.Context, paper *biz.Paper) (int32, error) {
	if err := validator.ValidatorStruct(paper); err != nil {
		return 0, errors.InvalidArgument(apierr.Errors_FieldInvalid, err.Error())
	}
	err := pr.data.mysqlDB.Create(paper).Error
	return paper.Id, err
}

func (pr *paperRepo) DeletePaper(ctx context.Context, id int32) error {
	err := pr.data.mysqlDB.Delete(biz.Paper{Id: id}).Error
	if err != nil {
		if errors3.Is(err, gorm.ErrRecordNotFound) {
			err = errors.InvalidArgument(apierr.Errors_PaperNoFound, "稿件不存在")
		} else {
			err = errors2.WithStack(errors.Unknown(apierr.Errors_MysqlErr, err.Error()))
		}
	}
	return err
}

func (pr *paperRepo) UpdatePaper(ctx context.Context, paper *biz.Paper) error {
	return nil
}

func (pr *paperRepo) GetPaper(ctx context.Context, id int32, columns []string) (*biz.Paper, error) {
	paper := biz.Paper{Id: id}
	tx := pr.data.mysqlDB
	if len(columns) > 0 {
		tx = tx.Select(columns)
	}
	err := tx.First(&paper).Error
	if err != nil {
		if errors2.Is(err, gorm.ErrRecordNotFound) {
			err = errors.NotFound(apierr.Errors_PaperNoFound, "稿件找不到")
		} else {
			err = errors2.WithStack(errors.Unknown(apierr.Errors_MysqlErr, err.Error()))
		}
	}
	return &paper, err
}

func (pr *paperRepo) GetPaperCode(ctx context.Context, id int32) (string, error) {
	paper := biz.Paper{Id: id}
	err := pr.data.mysqlDB.Select("code").First(&paper).Error
	if err != nil {
		if errors2.Is(err, gorm.ErrRecordNotFound) {
			err = errors.NotFound(apierr.Errors_PaperNoFound, "稿件找不到")
		} else {
			err = errors2.WithStack(errors.Unknown(apierr.Errors_MysqlErr, err.Error()))
		}
	}
	return paper.Code, err
}

func (pr *paperRepo) ListPaper(ctx context.Context, communityId int32, startIndex, pageSize int, columns []string) ([]*biz.Paper, error) {
	if pageSize <= 0 {
		fmt.Println("返回错误")
		return nil, errors.InvalidArgument(apierr.Errors_PageSizeNoZero, "page_num参数必须大于0")
	}
	var papers []*biz.Paper
	tx := pr.data.mysqlDB
	if len(columns) > 0 {
		tx = tx.Select(columns)
	}
	if communityId != 0 {
		tx = tx.Where("community_id = ?", communityId)
	}
	err := tx.Offset(startIndex).Limit(pageSize).Find(&papers).Error
	if err != nil {
		err = errors2.WithStack(errors.Unknown(apierr.Errors_MysqlErr, err.Error()))
	}
	return papers, err
}

