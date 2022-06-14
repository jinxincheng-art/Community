package model

import "time"

type DiscussPost struct {
	Id int `gorm:"column:id"`
	//用户id
	UserId int `gorm:"column:user_id"`
	//标题
	Title string `gorm:"column:title"`
	//内容
	Content string `gorm:"column:content"`
	//类型
	Type int `gorm:"column:type"`
	//状态
	Status int `gorm:"column:status"`
	//创建时间
	CreateTime time.Time `gorm:"column:create_time"`
	//评论数量
	CommentCount int `gorm:"column:comment_count"`
	//评分
	Score float64 `gorm:"column:score"`
}

func NewDiscussPost() DiscussPost {
	return DiscussPost{}
}
