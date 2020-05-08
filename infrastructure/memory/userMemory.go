package memory

import (
	"errors"
	"fmt"
)

// User情報を配列への保存と参照を実装(DB代わり)

// UserMemory インターフェイス
type UserMemory interface {
	Store(data User) (bool, error)
	Get(id int) (*User, error)
	GetAll() ([]*User, error)
}

// NewUserMemory User情報をメモリに保存 [この関数をwireでDI]
func NewUserMemory() UserList {
	return UserList{[]*User{}}
}

// User 戻り値の形式を定義。ドメインモデルのデータと同じだけど層をまたいで参照させないためにここでも定義している
type User struct {
	ID   int
	Name string
}

// UserList ユーザ情報一覧構造体
type UserList struct {
	userList []*User
}

// generateID generate Key
func (u *UserList) generateID() (int, error) {
	const initID int = 1

	if len(u.userList) == 0 {
		return initID, nil
	}

	var lu = u.userList[len(u.userList)-1]
	if lu == nil {
		return initID, nil
	}
	var id = lu.ID + 1
	return id, nil
}

// Store ユーザ情報を保存
func (u *UserList) Store(data User) (bool, error) {
	id, err := u.generateID()
	if err != nil {
		return false, errors.New("No ID")
	}

	user := &User{
		ID:   id,
		Name: data.Name,
	}

	u.userList = append(u.userList, user)
	return true, nil
}

// Get IDでユーザ情報を取得
func (u UserList) Get(id int) (*User, error) {
	for _, ul := range u.userList {
		if ul.ID == id {
			return ul, nil
		}
	}
	return nil, fmt.Errorf("Error: %s", "no user data")
}

// GetAll ユーザ情報一覧を取得
func (u UserList) GetAll() ([]*User, error) {
	return u.userList, nil
}
