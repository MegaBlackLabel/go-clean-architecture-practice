package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"

	"github.com/MegaBlackLabel/go-clean-architecture-practice/adapter/controllers"
	"github.com/MegaBlackLabel/go-clean-architecture-practice/domain/model"
	"github.com/MegaBlackLabel/go-clean-architecture-practice/utils/logs"
)

// UserHandler routerのHTTPハンドラー処理
type UserHandler interface {
	CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params)
	GetAllUsers(w http.ResponseWriter, r *http.Request, _ httprouter.Params)
	GetUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
}

type userHandler struct {
	UserController controllers.UserController
	logs           logs.Logger
}

// NewUserHandler routerのHTTPハンドラー処理。adapter/controllersを呼び出す [この関数をwireでDI]
func NewUserHandler(uc controllers.UserController, log logs.Logger) UserHandler {
	return &userHandler{UserController: uc, logs: log}
}

func (uh *userHandler) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	data := &model.User{}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		fmt.Fprintln(w, "Bad request: "+err.Error())
		return
	}

	fmt.Println(data)

	result, err := uh.UserController.CreateUser(data)
	if err == nil {
		if result == true {
			uh.logs.Infof("registerd")
			fmt.Fprintln(w, "registerd")
			return
		}
	}
	uh.logs.Infof("failed")
	fmt.Fprintln(w, "failed")
}

func (uh *userHandler) GetAllUsers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u, _ := uh.UserController.GetAllUsers()
	// JSON変換処理はPresenterに置くべきだけど今回は省略
	json.NewEncoder(w).Encode(u)
}

func (uh *userHandler) GetUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, _ := strconv.Atoi(ps.ByName("id"))
	uh.logs.Infof("target id: %d", id)
	u, _ := uh.UserController.GetUser(id)
	// JSON変換処理はPresenterに置くべきだけど今回は省略
	json.NewEncoder(w).Encode(u)
}
