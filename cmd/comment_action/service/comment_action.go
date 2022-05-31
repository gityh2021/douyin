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

package service

import (
	"context"
	"douyin/v1/cmd/comment_action/dal/db"
	"douyin/v1/kitex_gen/comment_action"
)

type CommentService struct {
	ctx context.Context
}

// NewCommentService new CommentService
func NewCommentService(ctx context.Context) *CommentService {
	return &CommentService{ctx: ctx}
}

// Comment comment info
func (s *CommentService) CommentService(req *comment_action.Comment_Action) error {
	commentModel := &db.Comment_Action{
		User_id: req.UserId,
		Token: req.Token,
		Video_id: req.VideoId,
		Action_type: req.ActionType,
		Comment_text: req.CommentText,
		Comment_id: req.CommentId,
		Create_date: req.CreateDate,
	}
	return db.VideoCommentAction(s.ctx, commentModel)
}