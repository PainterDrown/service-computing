package services

import (
	"net/http"
	"strconv"

	"github.com/painterdrown/service-computing/cloudgo-data/entities"
	"github.com/unrolled/render"
)

func postUserInfoHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		if len(req.Form["username"]) == 0 {
			panic("缺乏 username 参数")
		}
		if len(req.Form["username"][0]) == 0 {
			panic("参数 username 错误")
		}
		u := entities.NewUser(req.Form["username"][0])
		entities.UserService.Save(u)
		formatter.JSON(w, http.StatusOK, u)
	}
}

func getUserInfoHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		if len(req.Form["user_id"]) == 0 {
			all := entities.UserService.FindAll()
			formatter.JSON(w, http.StatusOK, all)
		} else {
			i, _ := strconv.ParseInt(req.Form["user_id"][0], 10, 32)
			u := entities.UserService.FindByID(int(i))
			formatter.JSON(w, http.StatusBadRequest, u)
		}
	}
}
