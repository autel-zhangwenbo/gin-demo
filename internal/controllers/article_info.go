package controllers

import (
	"gin-demo/internal/utils/errorcode"
	"net/http"

	"github.com/gin-gonic/gin"

	"gin-demo/internal/models"
	"gin-demo/internal/services"

)


//	@Summary	获取单个文章
//	@Produce	json
//	@Param		id	path		int		true	"文章ID"
//	@Success	200	{object}	Article	"成功"
//	@Failure	400	{object}	string	"请求错误"
//	@Failure	500	{object}	string	"内部错误"
//	@Router		/api/v1/articles/{id} [get]
func GetArticle(c *gin.Context) {

	return
}

//	@Summary	获取多个文章
//	@Produce	json
//	@Param		name		query		string	false	"文章名称"
//	@Param		tag_id		query		int		false	"标签ID"
//	@Param		state		query		int		false	"状态"
//	@Param		page		query		int		false	"页码"
//	@Param		page_size	query		int		false	"每页数量"
//	@Success	200			{object}	JsonResult{data=services.PageResponse{data=[]Article}}	"成功"
//	@Failure	400			{object}	string	"请求错误"
//	@Failure	500			{object}	string	"内部错误"
//	@Router		/api/v1/articles [get]
func ListArticles(c *gin.Context) {

	var req services.ListArticleReq
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.JsonRsp(errorcode.CodeErr{
			Code: errorcode.ErrGinShouldBind.Code,
			Msg:  err.Error(),
		}))
		return
	}
	c.JSON(http.StatusOK, models.JsonRsp(services.ListArticles(c, &req)))
}

//	@Summary	创建文章
//	@Produce	json
//	@Param		tag_id			body		string	true	"标签ID"
//	@Param		title			body		string	true	"文章标题"
//	@Param		desc			body		string	false	"文章简述"
//	@Param		cover_image_url	body		string	true	"封面图片地址"
//	@Param		content			body		string	true	"文章内容"
//	@Param		created_by		body		int		true	"创建者"
//	@Param		state			body		int		false	"状态"
//	@Success	200				{object}	Article	"成功"
//	@Failure	400				{object}	string	"请求错误"
//	@Failure	500				{object}	string	"内部错误"
//	@Router		/api/v1/articles [post]
func CreateArticle(c *gin.Context) {

}

//	@Summary	更新文章
//	@Produce	json
//	@Param		tag_id			body		string	false	"标签ID"
//	@Param		title			body		string	false	"文章标题"
//	@Param		desc			body		string	false	"文章简述"
//	@Param		cover_image_url	body		string	false	"封面图片地址"
//	@Param		content			body		string	false	"文章内容"
//	@Param		modified_by		body		string	true	"修改者"
//	@Success	200				{object}	Article	"成功"
//	@Failure	400				{object}	string	"请求错误"
//	@Failure	500				{object}	string	"内部错误"
//	@Router		/api/v1/articles/{id} [put]
func UpdateArticle(c *gin.Context) {
	return
}

//	@Summary	删除文章
//	@Produce	json
//	@Param		id	path		int		true	"文章ID"
//	@Success	200	{string}	string	"成功"
//	@Failure	400	{object}	string	"请求错误"
//	@Failure	500	{object}	string	"内部错误"
//	@Router		/api/v1/articles/{id} [delete]
func DeleteArticle(c *gin.Context) {

	return
}