package v1

import "github.com/gin-gonic/gin"

type Article struct{}

func NewArticle() Article {
	return Article{}
}

func (a Article) Get(c *gin.Context)    {}
func (a Article) List(c *gin.Context)   {}
func (a Article) Create(c *gin.Context) {}
func (a Article) Update(c *gin.Context) {}

// @Summary删除文章
// @Produce json
// @Param id path int true "标签ID"
// @success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// GFailure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/article/{id} [delete]
func (a Article) Delete(c *gin.Context) {}
