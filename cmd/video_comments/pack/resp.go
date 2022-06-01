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
	"douyin/v1/kitex_gen/video_comments"
	"douyin/v1/pkg/errno"
	"errors"
	"time"
)

// BuildVideoCommentResp build vcResp from error
func BuildVideoCommentResp(err error) *video_comments.VCResp {
	if err == nil {
		return vcResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return vcResp(e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return vcResp(s)
}

func vcResp(err errno.ErrNo) *video_comments.VCResp {
	return &video_comments.VCResp{StatusCode: err.ErrCode, StatusMessage: err.ErrMsg, ServiceTime: time.Now().Unix()}
}
