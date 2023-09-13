package services

import (
	"gin-demo/internal/dao"
	"gin-demo/internal/utils/errorcode"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"go.uber.org/zap"
)

//	@Param		name		query		string	false	"文章名称"
//	@Param		tag_id		query		int		false	"标签ID"
//	@Param		state		query		int		false	"状态"
//	@Param		page		query		int		false	"页码"
//	@Param		page_size	query		int		false	"每页数量"

type ListArticleReq struct {
	Name       string    `json:"name"`           // 名称
	TagId      int      `json:"tag_id"`         // tag id
	State      int      `json:"state"`          // 状态
	Page       int      `json:"sortWeight"`     // 页数
	PageSize   int      `json:"page_size"` // 每页的数量
}


// ListArticles 分页显示文章。
func ListArticles(ctx *gin.Context, req *ListArticleReq) any {

	page := req.Page
	pageSize := req.PageSize
	articleList, err := dao.ArticleList(page, pageSize)
	if err != nil {
		log.Error("dao.GetSPUs failed", zap.Error(err))
		return errorcode.CodeErr{
			Code: errorcode.ErrReadMySQL.Code,
			Msg:  err.Error(),
		}
	}

	c := len(articleList)

	return PageResponse{
		Total: int64(c),
		Data:  articleList,
	}
}