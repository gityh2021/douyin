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

package pack

import (
	"douyin/v1/kitex_gen/comment_action"
	"douyin/v1/pkg/errno"
	"errors"
	"time"
)

// BuildCommentActionResp build caResp from error
func BuildCommentActionResp(err error) *comment_action.CAResp {
	if err == nil {
		return caResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return caResp(e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return caResp(s)
}

func caResp(err errno.ErrNo) *comment_action.CAResp {
	return &comment_action.CAResp{StatusCode: err.ErrCode, StatusMessage: err.ErrMsg, ServiceTime: time.Now().Unix()}
}