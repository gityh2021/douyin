package vo

import (
	"douyin/v1/kitex_gen/user"
	"douyin/v1/kitex_gen/video"
	"douyin/v1/pkg/constants"
	"douyin/v1/pkg/myjwt"
	"github.com/gin-gonic/gin"
)

func PackVideoVos(users []*user.User, videos []*video.Video) []*VideoVo {
	userDict := make(map[int64]*user.User)
	for i := 0; i < len(users); i++ {
		userDict[users[i].GetId()] = users[i]
	}
	videoVos := make([]*VideoVo, len(videos))
	for i := 0; i < len(videos); i++ {
		videoVo := VideoVo{
			ID: videos[i].Id,
			Author: Author{
				ID:            userDict[videos[i].AuthorId].Id,
				Name:          userDict[videos[i].AuthorId].Name,
				FollowCount:   userDict[videos[i].AuthorId].FollowCount,
				FollowerCount: userDict[videos[i].AuthorId].FollowerCount,
				IsFollow:      userDict[videos[i].AuthorId].IsFollow,
			},
			PlayURL:       videos[i].PlayUrl,
			CoverURL:      videos[i].CoverUrl,
			FavoriteCount: videos[i].FavoriteCount,
			CommentCount:  videos[i].CommentCount,
			IsFavorite:    videos[i].IsFavorite,
			Title:         videos[i].Title,
		}
		videoVos[i] = &videoVo
	}
	return videoVos
}

func PackCommentVos(users []*user.User, comments []*video.Comment) []*CommentVo {
	userDict := make(map[int64]*user.User)
	for i := 0; i < len(users); i++ {
		userDict[users[i].GetId()] = users[i]
	}
	commentVos := make([]*CommentVo, len(comments))
	for i := 0; i < len(comments); i++ {
		commentVo := CommentVo{
			ID: comments[i].Id,
			Author: Author{
				ID:            userDict[comments[i].UserId].Id,
				Name:          userDict[comments[i].UserId].Name,
				FollowCount:   userDict[comments[i].UserId].FollowCount,
				FollowerCount: userDict[comments[i].UserId].FollowerCount,
				IsFollow:      userDict[comments[i].UserId].IsFollow,
			},
			Content:    comments[i].Content,
			CreateDate: comments[i].CreateDate,
		}
		commentVos[i] = &commentVo
	}
	return commentVos
}

func GetUserIdFromToken(c *gin.Context) int64 {
	claims := myjwt.ExtractClaims(c)
	if claims[constants.IdentityKey] == nil {
		return -1
	}
	return int64(claims[constants.IdentityKey].(float64))
}
