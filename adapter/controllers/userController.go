package controllers

import (
	"github.com/MegaBlackLabel/go-clean-architecture-practice/domain/model"
	"github.com/MegaBlackLabel/go-clean-architecture-practice/usecase/interactor"
	"github.com/MegaBlackLabel/go-clean-architecture-practice/utils/logs"
)

// ここでは外部から受け取ったデータをusecase/interactorの実装にわたす。コントローラーでは処理の実装をしない
// 受け取ったデータを構造体でusecaseに渡したい場合はusecase側でinputを作って受け渡し用の構造体を用意する。今回はそこまで無いのでmodel.Userのみ

// UserController userController構造体のインターフェイス
type UserController interface {
	CreateUser(data *model.User) (bool, error)
	GetAllUsers() ([]*model.User, error)
	GetUser(id int) (*model.User, error)
}

type userController struct {
	UserInteractor interactor.UserInteractor
	logs           logs.Logger
}

// NewUserController userのコントローラー。処理の実態であるusecase/interactorを呼び出す [この関数をwireでDI]
func NewUserController(us interactor.UserInteractor, log logs.Logger) UserController {
	return &userController{UserInteractor: us, logs: log}
}

// CreateUser ユーザ登録
func (uc *userController) CreateUser(data *model.User) (bool, error) {
	return uc.UserInteractor.Save(data)
}

// GetAllUsers ユーザ一覧取得
func (uc *userController) GetAllUsers() ([]*model.User, error) {
	us, err := uc.UserInteractor.GetAll()
	if err != nil {
		return nil, err
	}
	return us, nil
}

// GetUser　yu-zashutoku

func (uc *userController) GetUser(id int) (*model.User, error) {
	u, err := uc.UserInteractor.Get(id)
	return u, err
}
