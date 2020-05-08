package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/MegaBlackLabel/go-clean-architecture-practice/infrastructure/web"
	"github.com/MegaBlackLabel/go-clean-architecture-practice/registry"
)

func main() {
	// ルーター呼び出し
	r := httprouter.New()

	// DI呼び出し
	app := registry.InitialiseApps()
	if app == nil {
		panic("ERR!!!")
	}

	// infrastrcture/routerでルーティーング処理
	web.NewRouter(r, app)

	// HTTPサーバ起動
	http.ListenAndServe(":8080", r)
}
