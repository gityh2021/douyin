package db

import (
	"context"
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	ID      int64  `gorm:"primarykey" json:"id"`
	UserId  int64  `gorm:"not null" json:"user_id"`
	VideoId int64  `gorm:"not null" json:"video_id"`
	Content string `gorm:"not null" json:"content"`
}

func MPostCommentAction(ctx context.Context, comment *Comment) error {
	return DB.WithContext(ctx).Create(&comment).Error
}

func MDeleteComment(ctx context.Context, commentId int64) error {
	return DB.WithContext(ctx).Where("id = ?", commentId).Delete(&Comment{}).Error
}

func MGetCommentsByVideoId(ctx context.Context, videoId int64) ([]*Comment, error) {
	res := make([]*Comment, 0)
	if err := DB.WithContext(ctx).Where("video_id = ?", videoId).Order("created_at desc").Limit(30).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
