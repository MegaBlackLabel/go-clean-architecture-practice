package model

// ドメイン ここではuser構造体を定義

// User ドメインモデル
type User struct {
	ID   int
	Name string
}

// Users リスト
type Users []User
