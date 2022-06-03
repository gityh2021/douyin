package vo

type CommentVo struct {
	ID         int64  `json:"id"`
	Author     Author `json:"user"`
	Content    string `json:"content"`
	CreateDate string `json:"create_date"`
}
