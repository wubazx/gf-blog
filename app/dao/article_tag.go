// ============================================================================
// This is auto-generated by gf cli tool only once. Fill this file as you wish.
// ============================================================================

package dao

import (
	"go-gf-blog/app/dao/internal"
)

// articleTagDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type articleTagDao struct {
	*internal.ArticleTagDao
}

var (
	// ArticleTag is globally public accessible object for table t_article_tag operations.
	ArticleTag = &articleTagDao{
		internal.ArticleTag,
	}
)

// Fill with you ideas below.