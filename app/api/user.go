package api

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
	"go-gf-blog/app/model"
	"go-gf-blog/app/service"
	"go-gf-blog/library/response"
)

// 用户API管理对象
var User = new(apiUser)

type apiUser struct{}

// @summary 用户注册接口
// @tags    用户服务
// @produce json
// @param   entity body model.ApiUserSignUpReq true "注册请求"
// @router  /go-gf-blog/user/signup [POST]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *apiUser) SignUp(r *ghttp.Request) {
	var (
		apiReq     *model.ApiUserSignUpReq
		serviceReq *model.ServiceUserSignUpReq
	)
	if err := r.ParseForm(&apiReq); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if err := gconv.Struct(apiReq, &serviceReq); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if err := service.User.SignUp(serviceReq); err != nil {
		response.JsonExit(r, 1, err.Error())
	} else {
		response.JsonExit(r, 0, "ok")
	}
}

// @summary 检测用户账号接口(唯一性校验)
// @tags    用户服务
// @produce json
// @param   passport query string true "用户账号"
// @router  /go-gf-blog/user/checkpassport [GET]
// @success 200 {object} response.JsonResponse "执行结果:`true/false`"
func (a *apiUser) CheckPassport(r *ghttp.Request) {
	var (
		data *model.ApiUserCheckPassportReq
	)
	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if data.Passport != "" && !service.User.CheckPassport(data.Passport) {
		response.JsonExit(r, 1, "账号已经存在", false)
	}
	response.JsonExit(r, 0, "", true)
}

// @summary 检测用户昵称接口(唯一性校验)
// @tags    用户服务
// @produce json
// @param   nickname query string true "用户昵称"
// @router  /go-gf-blog/user/checknickname [GET]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *apiUser) CheckNickName(r *ghttp.Request) {
	var (
		data *model.ApiUserCheckNickNameReq
	)
	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if data.Nickname != "" && !service.User.CheckNickName(data.Nickname) {
		response.JsonExit(r, 1, "昵称已经存在", false)
	}
	response.JsonExit(r, 0, "ok", true)
}

// @summary 用户退出登录接口
// @tags    用户服务
// @produce json
// @router  /go-gf-blog/user/logout [POST]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *apiUser) Logout(r *ghttp.Request) {
	// todo 退出前操作
	response.JsonExit(r, 0, "退出成功")
}
