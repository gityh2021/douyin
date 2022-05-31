package db

import (
	"context"

	"gorm.io/gorm"
)

type Video_Comments struct {
	gorm.Model
	Comment_ID      int64  `gorm:"primarykey" json:"comment_id"`
	Comment_User_ID int64  `gorm:"not null" json:"comment_user_id"`
	Content         string `gorm:"not null" json:"content"`
	Create_Date     string `gorm:"not null" json:"create_date"`
}

func MGetVideosComments(ctx context.Context, videoId int64) ([]*Video_Comments, error) {
	res := make([]*Video_Comments, 0)
	if err := DB.WithContext(ctx).Where("video_id = ?", videoId).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
