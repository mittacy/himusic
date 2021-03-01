package data

import (
	"context"
	errors2 "errors"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/gomodule/redigo/redis"
	"github.com/mittacy/himusic/service/account/api/account/v1/apierr"
	"github.com/mittacy/himusic/service/account/internal/biz"
	"github.com/mittacy/tools/encryption"
	"github.com/mittacy/tools/validator"
	errors3 "github.com/pkg/errors"
	"gorm.io/gorm"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper("user_repo", logger),
	}
}

func (ur *userRepo) CreateUser(ctx context.Context, user *biz.User) error {
	if err := validator.ValidatorStruct(user); err != nil {
		return errors.InvalidArgument(apierr.Errors_FieldInvalid, err.Error())
	}
	/*
	 * 1. 检查邮箱是否已存在
	 * 2. 检查名字是否已存在
	 * 3. 加密密码并创建用户
	 */
	// 检查邮箱是否已占用
	exists, err := ur.emailExists(user.Email)
	if err != nil {
		return err
	}
	if exists {
		return errors.AlreadyExists(apierr.Errors_EmailExists, "邮箱已被注册")
	}
	// 检查name是否已占用
	exists, err = ur.nameExists(user.Name)
	if err != nil {
		return errors3.WithStack(errors.Unknown(apierr.Errors_MysqlErr, err.Error()))
	}
	if exists {
		return errors.AlreadyExists(apierr.Errors_NameExists, "名字已被注册")
	}

	user.Password, user.Salt = encryption.Encryption(user.Password)
	err = ur.data.db.Create(user).Error
	if err != nil {
		return errors3.WithStack(errors.Unknown(apierr.Errors_MysqlErr, err.Error()))
	}
	return nil
}

func (ur *userRepo) DeleteUser(ctx context.Context, id int32) error {
	err := ur.data.db.Delete(&biz.User{}, id).Error
	if err == nil {
		return nil
	}
	if errors3.Is(err, gorm.ErrRecordNotFound) {
		return errors.NotFound(apierr.Errors_UserNoFound, "用户不存在")
	}
	return errors3.WithStack(errors.Unknown(apierr.Errors_MysqlErr, err.Error()))
}

func (ur *userRepo) UpdateUser(ctx context.Context, user *biz.User) error {
	// 针对更新场景进行数据校验和处理
	var err error
	if user.Password != "" {
		if err := validator.ValidatorVar(user.Password, "password", "required,min=1,max=20"); err != nil {
			return errors.InvalidArgument(apierr.Errors_FieldInvalid, err.Error())
		}
		user.Password, user.Salt = encryption.Encryption(user.Password)
	}
	if user.Email != "" {
		if err := validator.ValidatorVar(user.Email, "email", "required,email"); err != nil {
			return errors.InvalidArgument(apierr.Errors_FieldInvalid, err.Error())
		}
		exists, err := ur.emailExists(user.Email)
		if err != nil {
			return errors3.WithStack(errors.Unknown(apierr.Errors_MysqlErr, err.Error()))
		}
		if exists {
			return errors.AlreadyExists(apierr.Errors_EmailExists, "邮箱已被注册")
		}
	}
	if user.Name != "" {
		if err := validator.ValidatorVar(user.Name, "name", "required,min=1,max=20"); err != nil {
			return errors.InvalidArgument(apierr.Errors_FieldInvalid, err.Error())
		}
		exists, err := ur.nameExists(user.Name)
		if err != nil {
			return errors3.WithStack(errors.Unknown(apierr.Errors_MysqlErr, err.Error()))
		}
		if exists {
			return errors.AlreadyExists(apierr.Errors_NameExists, "名字已被注册")
		}
	}
	err = ur.data.db.Model(user).Updates(user).Error
	if err != nil {
		if errors3.Is(err, gorm.ErrRecordNotFound) {
			err = errors.NotFound(apierr.Errors_UserNoFound, "用户不存在")
		} else {
			err = errors3.WithStack(errors.Unknown(apierr.Errors_MysqlErr, err.Error()))
		}
	}
	return err
}

func (ur *userRepo) GetUser(ctx context.Context, field, fieldVal string, column []string) (*biz.User, error) {
	user := biz.User{}
	tx := ur.data.db
	if len(column) > 0 {
		tx = tx.Select(column)
	}
	err := tx.Where(field + " = ?", fieldVal).First(&user).Error
	if err != nil {
		if errors2.Is(err, gorm.ErrRecordNotFound) {
			err = errors.NotFound(apierr.Errors_UserNoFound, "用户不存在")
		} else {
			err = errors3.WithStack(errors.Unknown(apierr.Errors_MysqlErr, err.Error()))
		}
	}
	return &user, err
}

func (ur *userRepo) ListUser(ctx context.Context, startIndex, pageSize int, column []string) (users []*biz.User, err error) {
	if startIndex < 0 || pageSize <= 0 {
		return nil, errors.InvalidArgument(apierr.Errors_FieldInvalid, "page_num >= 1, page_size > 0")
	}
	tx := ur.data.db
	if len(column) == 0 {
		tx = tx.Select(column)
	}
	err = tx.Offset(startIndex).Limit(pageSize).Find(&users).Error
	if err != nil {
		err = errors3.WithStack(errors.Unknown(apierr.Errors_MysqlErr, err.Error()))
	}
	return
}

func (ur *userRepo) GetUserByName(ctx context.Context, name string, column []string) (*biz.User, error) {
	msg := "获取用户详细信息"
	user := biz.User{}
	tx := ur.data.db
	if len(column) > 0 {
		tx = tx.Select(column)
	}
	err := tx.Where("name = ?", name).First(&user).Error
	if err != nil {
		if errors2.Is(err, gorm.ErrRecordNotFound) {
			err = errors.NotFound(apierr.Errors_UserNoFound, msg)
		} else {
			err = errors3.WithStack(errors.Unknown(apierr.Errors_MysqlErr, err.Error()))
		}
	}
	return &user, err
}

func (ur *userRepo) GetUserByEmail(ctx context.Context, email string, column []string) (*biz.User, error) {
	msg := "获取用户详细信息"
	user := biz.User{}
	tx := ur.data.db
	if len(column) > 0 {
		tx = tx.Select(column)
	}
	err := tx.Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors2.Is(err, gorm.ErrRecordNotFound) {
			err = errors.NotFound(apierr.Errors_UserNoFound, msg)
		} else {
			err = errors3.WithStack(errors.Unknown(apierr.Errors_MysqlErr, err.Error()))
		}
	}
	return &user, err
}

// 验证验证码是否正确
// @Params key redis字符串键名
// @Params code 待验证的验证码
// Return bool 验证码是否正确
// Return error 错误信息
func (ur *userRepo) verifyCode(key string, code string) (bool, error) {
	val, err := redis.String(ur.data.redisDB.Do("get", key))
	if errors3.Is(err, redis.ErrNil) {
		return false, nil
	}
	return val == code, err
}

func (ur *userRepo) emailExists(email string) (is bool, err error) {
	user := biz.User{}
	r := 0
	err = ur.data.db.Raw("select 1 from "+user.TableName()+" where email = ?", email).Scan(&r).Error
	if err != nil {
		return
	}
	if r == 1 {
		is = true
	}
	return
}

func (ur *userRepo) nameExists(name string) (is bool, err error) {
	user := biz.User{}
	r := 0
	err = ur.data.db.Raw("select 1 from "+user.TableName()+" where name = ?", name).Scan(&r).Error
	if err != nil {
		return
	}
	if r == 1 {
		is = true
	}
	return
}
