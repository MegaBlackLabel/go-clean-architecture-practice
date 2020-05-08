package web

// ここではWEBの処理を書く。実際の処理はadapter/controllersを呼び出す

import (
	"github.com/julienschmidt/httprouter"

	"github.com/MegaBlackLabel/go-clean-architecture-practice/infrastructure/handler"
)

// NewRouter HTTPルーティング処理。infrastructure/handlerで処理を呼び出し
func NewRouter(router *httprouter.Router, handler handler.UserHandler) {
	// USER API
	router.POST("/api/user/register", handler.CreateUser)
	router.GET("/api/user/get", handler.GetAllUsers)
	router.GET("/api/user/get/:id", handler.GetUser)
}
