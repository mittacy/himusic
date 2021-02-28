package biz

type User struct {
	Id    int32 `gorm:"primaryKey"`
	Score int32 `validate:"required"`
}

func (*User) TableName() string {
	return "ums_user"
}

type ScoresRepo interface {
	UpdateScores(user *User) error
	GetScores(id int32) (int32, error)
}

type ScoresUseCase struct {
	repo ScoresRepo
}

func NewScoresUseRepo(repo ScoresRepo) *ScoresUseCase {
	return &ScoresUseCase{repo: repo}
}

func (uc *ScoresUseCase) UpdateScores(user *User) error {
	return uc.repo.UpdateScores(user)
}

func (uc *ScoresUseCase) GetScores(id int32) (int32, error) {
	return uc.repo.GetScores(id)
}