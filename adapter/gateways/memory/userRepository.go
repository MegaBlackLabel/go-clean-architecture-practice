package memory

import (
	"github.com/MegaBlackLabel/go-clean-architecture-practice/domain/model"
	"github.com/MegaBlackLabel/go-clean-architecture-practice/domain/repository"
	"github.com/MegaBlackLabel/go-clean-architecture-practice/infrastructure/memory"
)

// ここではinfrastrucureのuserMemoryとの接続を行う。その際にdomain/repositoryのインターフェイスの実装を行う

type userRepository struct {
	users memory.UserList
}

// NewUserRepository userRepository構造体のインターフェイス
func NewUserRepository(users memory.UserList) repository.UserRepository {
	return &userRepository{users: users}
}

// Store Userを保存
func (ur *userRepository) Save(data *model.User) bool {
	if data.Name == "" {
		return false
	}
	user := memory.User{
		Name: data.Name,
	}
	return ur.users.Store(user)
}

// FindAll　全ユーザ取得
func (ur *userRepository) FindAll() ([]*model.User, error) {
	users := []*model.User{}
	userList, err := ur.users.GetAll()
	if err != nil {
		return nil, err
	}
	for _, u := range userList {
		user := &model.User{
			ID:   u.ID,
			Name: u.Name,
		}
		users = append(users, user)
	}
	return users, nil
}

// FindById IDでユーザ検索
func (ur *userRepository) FindByID(id int) (*model.User, error) {
	u, err := ur.users.Get(id)
	if err != nil {
		return nil, err
	}
	user := &model.User{
		ID:   u.ID,
		Name: u.Name,
	}
	return user, err
}
