// Copyright 2021 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package handlers

import (
	"context"
	"douyin/v1/cmd/api/rpc"
	"douyin/v1/kitex_gen/comment_action"
	"douyin/v1/pkg/errno"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

// Create or Delete Comment comment info
func CommentAction(c *gin.Context) {
	var commentVar VideoCommentParam
	if err := c.ShouldBind(&commentVar); err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}

	if len(commentVar.Content) == 0 || commentVar.CommentId <= 0 || commentVar.CommentUserId <= 0 {
		SendResponse(c, errno.ParamErr, nil)
		return
	}

	claims := jwt.ExtractClaims(c)
	userID := int64(claims["User_Id"].(float64))
	token := string(claims["Token"].(string))
	videoID := int64(claims["Video_Id"].(float64))
	action_type := int64(claims["Action_Type"].(float64))
	comment_text := string(claims["Comment_Text"].(string))
	comment_id := int64(claims["Comment_ID"].(float64))
	create_date := string(claims["Create_Date"].(string))
	err := rpc.CommentAction(context.Background(), &comment_action.Comment_Action{
		UserId:      userID,
		Token:       token,
		VideoId:     videoID,
		ActionType:  action_type,
		CommentText: comment_text,
		CommentId:   comment_id,
		CreateDate:  create_date,
	})
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	SendResponse(c, errno.Success, nil)
}
