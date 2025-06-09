package request

import (
	"mime/multipart"

	"gorm.io/gorm"
)

// PageInfo Paging common input parameter structure
type PageInfo struct {
	Page     int    `json:"page" form:"page"`         // 页码
	PageSize int    `json:"pageSize" form:"pageSize"` // 每页大小
	Keyword  string `json:"keyword" form:"keyword"`   // 关键字
}

func (r *PageInfo) Paginate() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if r.Page <= 0 {
			r.Page = 1
		}
		switch {
		case r.PageSize > 100:
			r.PageSize = 100
		case r.PageSize <= 0:
			r.PageSize = 10
		}
		offset := (r.Page - 1) * r.PageSize
		return db.Offset(offset).Limit(r.PageSize)
	}
}

// GetById Find by id structure
type GetById struct {
	ID int `json:"id" form:"id"` // 主键ID
}

func (r *GetById) Uint() uint {
	return uint(r.ID)
}

type IdsReq struct {
	Ids []int `json:"ids" form:"ids"`
}

// GetAuthorityId Get role by id structure
type GetAuthorityId struct {
	AuthorityId uint `json:"authorityId" form:"authorityId"` // 角色ID
}

type Empty struct{}

// 添加账号时的请求
type ImportSessionRequest struct {
	MasterUsername string                  `form:"master_username" binding:"required"`
	SessionFiles   []*multipart.FileHeader `form:"session_files[]" binding:"required"`
}

// 查询所有账号时的请求
type GetSessionRequest struct {
	MasterUsername string `json:"master_username"`
	Page           int    `json:"page"`
	PageSize       int    `json:"page_size"`
	Name           string `json:"name"`
	Tag            int    `json:"tag"`
}
type DeleteSessionRequest struct {
	Phone []string `json:"phone"`
}
type LoginSessionRequest struct {
	Phone  []string `json:"phone"`
	Proxy  string   `json:"proxy"`
	Server int      `json:"server"`
}
type OutLoginSessionRequest struct {
	Phone []string `json:"phone"`
}

// 分组请求
type ChangeTagTelegramRequest struct {
	Phone []string `json:"phone"`
	TagId int      `json:"tag_id"`
}
