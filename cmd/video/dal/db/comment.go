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

//发表评论 增加了事务处理
func MPostCommentAction(ctx context.Context, comment *Comment) error {
	return DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.WithContext(ctx).Create(&comment).Error; err != nil {
			return err
		}
		if err := tx.WithContext(ctx).Model(&Video{}).Where("id = ?", comment.VideoId).Update("comment_count", gorm.Expr("comment_count + ?", 1)).Error; err != nil {
			return err
		}
		return nil
	})
}

//删除评论 增加了事务处理
func MDeleteComment(ctx context.Context, commentId int64) error {
	return DB.Transaction(func(tx *gorm.DB) error {
		videoId := 0
		if err := tx.WithContext(ctx).Model(&Comment{}).Select("video_id").Where("id = ?", commentId).Find(&videoId).Error; err != nil {
			return err
		}
		if err := tx.WithContext(ctx).Model(&Video{}).Where("id = ?", videoId).Update("comment_count", gorm.Expr("comment_count - ?", 1)).Error; err != nil {
			return err
		}
		if err := tx.WithContext(ctx).Where("id = ?", commentId).Delete(&Comment{}).Error; err != nil {
			return err
		}
		return nil
	})
}

//获取评论列表
func MGetCommentsByVideoId(ctx context.Context, videoId int64) ([]*Comment, error) {
	res := make([]*Comment, 0)
	if err := DB.WithContext(ctx).Where("video_id = ?", videoId).Order("created_at desc").Limit(30).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// DB.Transaction(func(tx *gorm.DB) error {
// 	//先在Follow表中更改关注的关系
// 	//再在User表中更改follow_count与follower_count
// 	if req.ActionType == constants.RelationAdd {
// 		if err := tx.WithContext(ctx).Create(&Follower{UserID: req.ToUserId, FollowerID: req.UserId}).Error; err != nil {
// 			return err
// 		}
// 		user1.FollowCount += 1
// 		user2.FollowerCount += 1
// 	} else if req.ActionType == constants.RelationDel {
// 		if err := tx.WithContext(ctx).Where("user_id = ? and follower_id = ?", req.ToUserId, req.UserId).Delete(&Follower{}).Error; err != nil {
// 			return err
// 		}
// 		user1.FollowCount -= 1
// 		user2.FollowerCount -= 1
// 	}
// 	if err := tx.WithContext(ctx).Model(&user1).Select("follow_count").Updates(user1).Error; err != nil {
// 		return err
// 	}

// 	if err := tx.WithContext(ctx).Model(&user2).Select("follower_count").Updates(user2).Error; err != nil {
// 		return err
// 	}
// 	return nil //没有错误则提交事务
// })
