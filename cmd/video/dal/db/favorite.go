package db

import (
	"context"

	"gorm.io/gorm"
)

type Favorite struct {
	gorm.Model
	ID      int64 `gorm:"primarykey" json:"id"`
	UserId  int64 `gorm:"not null" json:"user_id"`
	VideoId int64 `gorm:"not null" json:"video_id"`
}

//获得喜欢列表
func MGetFavoriteList(ctx context.Context, userId int64) ([]*Video, error) {
	videos := make([]*Video, 0)
	ids := make([]int64, 0)

	//这里的deleted_at is null很需要 不然会查询到软删除的内容
	if err := DB.WithContext(ctx).Where("id in (?)",
		DB.WithContext(ctx).Table("favorites").Select("video_id").Where("user_id = ? AND deleted_at is null", userId).Find(&ids)).Find(&videos).Error; err != nil {
		return nil, err
	}
	return videos, nil
}

//增加喜欢 增加了事务处理
func MPostFavoriteAction(ctx context.Context, userId int64, videoId int64) error {
	favorite := Favorite{
		UserId:  userId,
		VideoId: videoId,
	}

	//判断喜欢表里是不是已经有喜欢了
	var cnt int64 = 0
	if err := DB.WithContext(ctx).Model(&Favorite{}).Where("user_id = ? and video_id = ?", userId, videoId).Count(&cnt).Error; err != nil {
		return err
	}

	if cnt != 0 {
		return nil
	}

	return DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.WithContext(ctx).Create(&favorite).Error; err != nil {
			return err
		}
		if err := DB.WithContext(ctx).Model(&Video{}).Where("id = ?", videoId).Update("favorite_count", gorm.Expr("favorite_count + ?", 1)).Error; err != nil {
			return err
		}
		return nil
	})

}

//撤销喜欢 增加了事务处理
func MCancelFavoriteAction(ctx context.Context, userId int64, videoId int64) error {
	//判断喜欢表里是否存在
	var cnt int64 = 0
	if err := DB.WithContext(ctx).Model(&Favorite{}).Where("user_id = ? and video_id = ?", userId, videoId).Count(&cnt).Error; err != nil {
		return err
	}
	// fmt.Println("cnt的值是", cnt)
	if cnt == 0 {
		return nil
	}

	return DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.WithContext(ctx).Where("user_id = ? and video_id = ?", userId, videoId).Delete(&Favorite{}).Error; err != nil {
			return err
		}

		if err := tx.WithContext(ctx).Model(&Video{}).Where("id = ?", videoId).Update("favorite_count", gorm.Expr("favorite_count - ?", 1)).Error; err != nil {
			return err
		}
		// fmt.Println("操作完毕了的")
		return nil
	})

	// if err := DB.WithContext(ctx).Model(&Favorite{}).Where("user_id = ? and video_id = ?", userId, videoId).Count(&cnt).Error; err != nil {
	// 	return err
	// }
	// fmt.Println("删除后的cnt为", cnt)

	// return nil

}

// MGetFavoriteIds 返回userId点赞过的videoIds内的视频
func MGetFavoriteIds(ctx context.Context, videoIds []*int64, userId int64) ([]*int64, error) {
	res := make([]*int64, 0)
	if err := DB.WithContext(ctx).Model(&Favorite{}).Select("video_id").Where("user_id = ? and video_id in (?)", userId, videoIds).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
