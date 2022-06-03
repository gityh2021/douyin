package handlers

import (
	"douyin/v1/cmd/api/rpc"
	"douyin/v1/cmd/api/vo"
	"douyin/v1/kitex_gen/video"
	"douyin/v1/pkg/constants"
	"douyin/v1/pkg/errno"
	"github.com/gin-gonic/gin"
	"strconv"
)

func PostComment(c *gin.Context) {
	token := c.Query("token")
	videoIdStr := c.Query("video_id")
	actionType := c.Query("action_type")
	commentText := c.Query("comment_text")
	commentIdStr := c.Query("comment_id")
	userId := vo.GetUserIdFromToken(c)
	if userId == constants.NotLogin {
		SendPostCommentResponse(c, nil, errno.LoginErr)
		return
	}
	if token == "" || videoIdStr == "" || actionType == "" {
		SendPostCommentResponse(c, nil, errno.ParamErr)
		return
	}
	videoId, err := strconv.ParseInt(videoIdStr, 10, 64)
	if err != nil {
		SendPostCommentResponse(c, nil, err)
		return
	}
	action, err := strconv.ParseInt(actionType, 10, 64)
	if err != nil {
		SendPostCommentResponse(c, nil, err)
		return
	}
	commentActionRequest := video.CommentActionRequest{
		Comment:    nil,
		ActionType: action,
	}
	switch action {
	case 1:
		if commentText == "" || commentIdStr != "" {
			SendPostCommentResponse(c, nil, errno.ParamErr)
			return
		}
		commentActionRequest.SetComment(&video.Comment{
			UserId:  userId,
			VideoId: videoId,
			Content: commentText,
		})
	case 2:
		if commentText != "" || commentIdStr == "" {
			SendPostCommentResponse(c, nil, errno.ParamErr)
			return
		}
		commentId, err := strconv.ParseInt(commentIdStr, 10, 64)
		if err != nil {
			SendPostCommentResponse(c, nil, err)
			return
		}
		commentActionRequest.SetComment(&video.Comment{
			Id: commentId,
		})
	default:
		SendPostCommentResponse(c, nil, errno.ParamErr)
		return
	}

	resp, err := rpc.PostComment(c, &commentActionRequest)
	if err != nil {
		SendPostCommentResponse(c, nil, err)
		return
	}
	SendPostCommentResponse(c, resp.Comment, errno.Success)
}

func QueryComments(c *gin.Context) {
	tokenId := vo.GetUserIdFromToken(c)
	tokenStr := c.Query("token")
	videoIdStr := c.Query("video_id")
	if tokenStr == "" || videoIdStr == "" {
		SendQueryCommentResponse(c, nil, errno.ParamErr)
		return
	}
	if tokenId == -1 {
		SendQueryCommentResponse(c, nil, errno.LoginErr)
		return
	}
	videoId, err := strconv.ParseInt(videoIdStr, 10, 64)
	if err != nil {
		SendQueryCommentResponse(c, nil, err)
		return
	}
	comments, err := rpc.GetCommentsByVideoId(c, videoId)
	if err != nil {
		SendQueryCommentResponse(c, nil, err)
		return
	}
	if len(comments) == 0 {
		SendQueryCommentResponse(c, comments, errno.Success)
		return
	} else {
		ids := make([]int64, len(comments))
		for i := 0; i < len(comments); i++ {
			ids[i] = comments[i].UserId
		}
		users, err := rpc.GetUsersByIds(c, ids, tokenId)
		if err != nil {
			SendQueryCommentResponse(c, nil, err)
			return
		}
		commentVos := vo.PackCommentVos(users, comments)
		SendQueryCommentResponse(c, commentVos, errno.Success)
	}
}
