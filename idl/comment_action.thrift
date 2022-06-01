namespace go comment_action

struct CAResp {
    1:i64 status_code
    2:string status_message
    3:i64 service_time
}

struct Comment_Action {
    1:i64 user_id
    2:string token
    3:i64 video_id
    4:i64 action_type //1.发布评论 2.删除评论
    5:string comment_text
    6:i64 comment_id
    7:string create_date
}

struct Video_Comments{
    1:i64 comment_id
    2:i64 comment_user_id
    3:string content
    4:string create_date
}

struct CommentActionResponse {
    1:CAResp caresp
    2:Video_Comments video_comments
}

service CommentActionService {
    CommentActionResponse PostCommentActionResponse(1:Comment_Action req)
}