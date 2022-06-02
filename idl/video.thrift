namespace go video

struct BaseResp {
    1:i32 status_code
    2:string status_msg
    3:i64 service_time
}

struct Video{
    1:i64 id
    2:i64 author_id
    3:string play_url
    4:string cover_url
    5:i64 favorite_count
    6:i64 comment_count
    7:bool is_favorite
    8:string title
}

struct PublishListResponse{
    1:BaseResp base_resp
    2:list<Video> video_list
}

struct VideoFeedResponse{
    1:BaseResp base_resp
    2:i64 next_time
    3:list<Video> video_list
}
// comment
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

service VideoService{
    PublishListResponse GetPublishListByUser(1:i64 user_id)
    VideoFeedResponse GetVideosByLastTime(1:i64 last_time)
    BaseResp PublishVideo(1:Video published_video)
    CommentActionResponse PostCommentActionResponse(1:Comment_Action req)
    PublishVideoCommentResponse GetPublishVideoCommentByVideo(1:i64 video_id)
}
