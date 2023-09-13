package models

import (
	"time"

	"gorm.io/gorm"

	"gin-demo/internal/utils/errorcode"
)

type Model struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index" `
	CreatorID string         `json:"creatorID" gorm:"size:32"`
	UpdaterID string         `json:"updaterID" gorm:"size:32"`
}

type JsonResult struct {
	Code int    `json:"code"`
	Msg  string `json:"msg" example:"ok"`
	Data any    `json:"data,omitempty"`
}

func JsonRsp(data any) JsonResult {
	var rsp JsonResult

	switch v := data.(type) {
	case error:
		rsp.Msg = v.Error()
		rsp.Code = 9999
		if v, ok := data.(errorcode.CodeErr); ok {
			rsp.Code = v.Code
		}
	default:
		rsp.Data = v
		rsp.Msg = "ok"
	}
	return rsp
}