package interactor

import (
	"github.com/MegaBlackLabel/go-clean-architecture-practice/domain/model"
	"github.com/MegaBlackLabel/go-clean-architecture-practice/domain/repository"
)

// ここではuserInteractorの構造体を定義してusecase/repositoryを実装する

// UserInteractor インターフェイスはNewUserInteractorで使用する
type UserInteractor interface {
	Save(data *model.User) (bool, error)
	GetAll() ([]*model.User, error)
	Get(id int) (*model.User, error)
}

type userInteractor struct {
	UserRepository repository.UserRepository
}

// NewUserInteractor userの実処理。実処理と登録・参照処理としてusecase/repositoryを呼び出す [この関数をwireでDI]
func NewUserInteractor(ur repository.UserRepository) UserInteractor {
	return &userInteractor{UserRepository: ur}
}

// Save ユーザ情報を保存
func (us *userInteractor) Save(data *model.User) (bool, error) {
	return us.UserRepository.Save(data)
}

// GetAll ユーザ一覧取得
func (us *userInteractor) GetAll() ([]*model.User, error) {
	u, err := us.UserRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return u, nil
}

// Get ユーザ情報取得
func (us *userInteractor) Get(id int) (*model.User, error) {
	u, err := us.UserRepository.FindByID(id)
	return u, err
}
