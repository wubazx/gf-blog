// ==========================================================================
// This is auto-generated by gf cli tool. DO NOT EDIT THIS FILE MANUALLY.
// ==========================================================================

package internal

import (
	"github.com/gogf/gf/os/gtime"
)

// Tag is the golang structure for table t_tag.
type Tag struct {
    Id           int         `orm:"id,primary"    json:"id"             c:"-"`             // 标签ID
    Name         string      `orm:"name"          json:"name"           c:"name"`          // 标签名称
    ArticleCount int         `orm:"article_count" json:"article_count"  c:"article_count"` // 标签下的文章数量
    CreatedAt    *gtime.Time `orm:"created_at"    json:"created_at"     c:"created_at"`    // 标签创建时间
    UpdatedAt    *gtime.Time `orm:"updated_at"    json:"updated_at"     c:"updated_at"`    // 标签更新时间
}