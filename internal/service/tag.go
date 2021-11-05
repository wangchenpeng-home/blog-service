package service

import (
	"github.com/wangchenpeng-home/blog-service/internal/model"
	"github.com/wangchenpeng-home/blog-service/pkg/app"
)

// 处理标签模块的业务逻辑

type CountTagRequest struct {
	Name string `form:"name" binding:"max=100"`
	State uint8 `form:"state,default=1" binding:"oneof=0 1"`
}


type TagListRequest struct {
	Name string `form:"name" binding:"max=100"`
	State uint8 `form:"state,default=1" binding:"oneof=0 1"`
}

type CreateTagRequest struct {
	Name string `form:"name" binding:"required,min=3,max=100"`
	CreateBy string `form:"created_by" binding:"required,min=3,max=100"`
	State uint8 `form:"state,default=1" binding:"oneof=0 1"`
}

type UpdateTagRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
	Name string `form:"name" binding:"required,min=3,max=100"`
	State uint8 `form:"state,default=1" binding:"required,oneof=0 1"`
	ModifiedBy string `form:"modified_by" binding:"required,min=3,max=100"`
}

type DeleteTagRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}

func (svc *Service) CountTag(param *CountTagRequest) (int, error) {
	return svc.dao.CountTag(param.Name, param.State)
}

func (svc *Service) GetTagList(param *TagListRequest, paper *app.Paper) ([]*model.Tag, error) {
	return svc.dao.GetTagList(param.Name, param.State, paper.Page, paper.PageSize)
}

func (svc *Service) CreateTag(param *CreateTagRequest) error {
	return svc.dao.CreateTag(param.Name, param.State, param.CreateBy)
}

func (svc *Service) UpdateTag(param *UpdateTagRequest) error {
	return svc.dao.UpdateTag(param.ID, param.Name, param.State, param.ModifiedBy)
}

func (svc* Service) DeleteTag(param *DeleteTagRequest) error {
	return svc.dao.DeleteTag(param.ID)
}
