package repository

import "github.com/MegaBlackLabel/go-clean-architecture-practice/domain/model"

// ここではRepositoryのインターフェイスを定義。adapter/gatewaysで実装する

// UserRepository インターフェイス wireの関係でusecaseとadapterから参照されるのでdomainにあるほうが最適
type UserRepository interface {
	Save(data *model.User) (bool, error)
	FindAll() ([]*model.User, error)
	FindByID(id int) (*model.User, error)
}
