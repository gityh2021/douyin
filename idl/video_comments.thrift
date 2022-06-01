namespace go video_comments

struct VCResp {
    1:i64 status_code
    2:string status_message
    3:i64 service_time
}

struct Video_Comments{
    1:i64 comment_id
    2:i64 comment_user_id
    3:string content
    4:string create_date
}

struct PublishVideoCommentResponse{
    1:VCResp vcresp
    2:list<Video_Comments> video_comments
}

service VideoCommentsService{
    PublishVideoCommentResponse GetPublishVideoCommentByVideo(1:i64 video_id)
}
