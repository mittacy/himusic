package biz

import (
	"github.com/go-kratos/kratos/v2/errors"
	v1 "github.com/mittacy/himusic/service/account/api/account/v1"
	"github.com/mittacy/himusic/service/account/api/account/v1/apierr"
	"github.com/mittacy/tools/encryption"
)

type User struct {
	Id         int32  `gorm:"primaryKey"`
	Name       string `validate:"required,min=1,max=20"`
	Email      string `validate:"required,email"`
	Password   string `validate:"required,min=1,max=32"`
	Icon       string `validaste:"omitempty,url"`
	PaperCount int32
	BuyCount   int32
	Coin       int32
	Score      int32
	CreateTime int64 `gorm:"autoCreateTime"`
	UpdateTime int64 `gorm:"autoUpdateTime"`
	LoginTime  int64
	Salt       string
}

func (*User) TableName() string {
	return "ums_user"
}

type UserRepo interface {
	CreateUser(user *User, codeKey, codeVal string) error
	DeleteUser(id int32) error
	UpdateUser(user *User, way v1.UpdateUserRequest_Way) error
	GetUser(id int32, column ...string) (*User, error)
	GetUserByName(name string, column ...string) (*User, error)
	GetUserByEmail(email string, column ...string) (*User, error)
	ListUser(startIndex, num int, column ...string) ([]*User, error)
}

type UserUseCase struct {
	repo UserRepo
}

func NewUserUseRepo(repo UserRepo) *UserUseCase {
	return &UserUseCase{repo: repo}
}

func (uc *UserUseCase) CreateUser(user *User, codeKey, codeVal string) error {
	return uc.repo.CreateUser(user, codeKey, codeVal)
}

func (uc *UserUseCase) DeleteUser(id int32) error {
	return uc.repo.DeleteUser(id)
}


func (uc *UserUseCase) UpdateUser(user *User, way v1.UpdateUserRequest_Way) error {
	return uc.repo.UpdateUser(user, way)
}

//func (uc *UserUseCase) UpdateIcon(id int32, icon string) error {
//	return uc.repo.UpdateUser(&User{Id: id, Icon: icon})
//}
//
//func (uc *UserUseCase) UpdatePassword(id int32, password string) error {
//	return uc.repo.UpdateUser(&User{Id: id, Password: password})
//}
//
//func (uc *UserUseCase) UpdateEmail(id int32, email string) error {
//	return uc.repo.UpdateUser(&User{Id: id, Email: email})
//}
//
//func (uc *UserUseCase) UpdateName(id int32, name string) error {
//	return uc.repo.UpdateUser(&User{Id: id, Name: name})
//}

func (uc *UserUseCase) VerifyPasswordByName(name string, password string) (bool, error) {
	user, err := uc.repo.GetUserByName(name, "password", "salt")
	if err != nil {
		return false, err
	}
	if encryption.EncryptionBySalt(password, user.Salt) == user.Password {
		return true, nil
	}
	return false, nil
}

func (uc *UserUseCase) VerifyPasswordByEmail(email string, password string) (bool, error) {
	user, err := uc.repo.GetUserByEmail(email, "password", "salt")
	if err != nil {
		return false, err
	}
	if encryption.EncryptionBySalt(password, user.Salt) == user.Password {
		return true, nil
	}
	return false, nil
}

func (uc *UserUseCase) GetUser(id int32) (*User, error) {
	return uc.repo.GetUser(id, "id", "name", "email", "icon", "paper_count", "buy_count", "coin",
		"score", "create_time", "update_time", "login_time")
}

func (uc *UserUseCase) ListUser(pageNum, pageSize int) ([]*User, error) {
	if pageNum <= 0 || pageSize <= 0 {
		return nil, errors.InvalidArgument(apierr.Errors_PageInvalid, "每页大小必须>0,页码必须>=1")
	}
	i := (pageNum - 1) * pageSize
	return uc.repo.ListUser(i, pageSize, "id", "name", "email", "paper_count", "buy_count",
		"coin", "score", "create_time", "update_time", "login_time")
}
