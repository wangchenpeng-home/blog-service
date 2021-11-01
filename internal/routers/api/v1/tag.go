package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/wangchenpeng-home/blog-service/global"
	"github.com/wangchenpeng-home/blog-service/pkg/app"
	"github.com/wangchenpeng-home/blog-service/pkg/errcode"
)

type Tag struct {
}

func NewTag() Tag {
	return Tag{}
}

func (a Tag) Get(c *gin.Context) {}

//@Summary获取多个标签// @Produce json
// Param name query string false "标签名称" maxlength(100)
// @Param state query int false "状态" Enums(0,1) default(1)
// @Param page query int false "页码"
// @Param page size query int false "每页数量"
// @success 200 {object} model.Tag "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags [get]
func (a Tag) List(c *gin.Context) {
	param := struct {
		Name  string
		State uint8
	}{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
		response.ToResponse(errRsp)
		return
	}

	response.ToResponse(gin.H{})
	return
}

// @Summary新增标签
// @Producejson
// @Param name body string true "标签名称" minlength(3) maxlength(100)
// @Param state body int false "状态"Enums(0,1) default(1)
// @Param created by body string false "创建者" minlength(3) maxlength(100)
// @success 200 {object} model.Tag "成功”
// @Failure 400 {object} errcode.Error "请求错误"
// CFailure 500 {obiectl errcode.Error "内部错误”
// @Router /api/v1/tags [post]
func (a Tag) Create(c *gin.Context) {}

// @Summary更新标签
// @Produce json
// @Param id path int true "标签ID"
// @Param name body string false "标签名称" minlength(3) maxlength (100)
// @Param state body int false "状态"Enums(0,1) default(1)
// @Param modified by body string true "修改者" ninlength(3) maxlength(100)
// @success 200 {array} model.Tag "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags/{id} [put]
func (a Tag) Update(c *gin.Context) {}

// @Summary删除标签
// @Produce json
// @Param id path int true "标签ID"
// @success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// GFailure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags/{id} [delete]
func (a Tag) Delete(c *gin.Context) {}
