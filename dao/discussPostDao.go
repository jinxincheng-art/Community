package dao

import (
	"community/model"
	"fmt"
	"go.uber.org/zap"
)

// SelectDiscussPosts 分页查询帖子
func SelectDiscussPosts(page, limit int) []*model.DiscussPost{
	discussPosts := make([]*model.DiscussPost,10)
	if err := DB.Table("discuss_post").Offset((page - 1) * limit).Limit(limit).Find(&discussPosts).Error; err != nil{
		fmt.Printf("查询失败:%v\n",zap.Error(err))
	}
	return discussPosts
}
