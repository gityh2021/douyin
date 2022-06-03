namespace go video

struct BaseResp {
    1:i32 status_code
    2:string status_msg
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

// favorite
struct FavoriteActionRequest {
   1:i64 user_id
   3:i64 video_id
   4:i64 action_type
}

struct FavoriteListResponse {
      1:BaseResp base_resp
      2:list<Video> video_list
}

// comment
struct Comment{
    1:i64    id
    2:i64    user_id
    3:i64    video_id
    4:string content
    5:string create_date
}
struct CommentListResponse {
      1:BaseResp base_resp
      2:list<Comment> comment_list
}
struct CommentActionResponse {
      1:BaseResp base_resp
      2:Comment comment
}
struct CommentActionRequest {
      1:Comment comment
      2:i64 action_type
}
service VideoService{
    PublishListResponse GetPublishListByUser(1:i64 user_id)
    VideoFeedResponse GetVideosByLastTime(1:i64 last_time, 2:i64 user_id)
    BaseResp PublishVideo(1:Video published_video)
    BaseResp FavoriteByUser(1:FavoriteActionRequest request)
    FavoriteListResponse GetFavoriteListBYUser(1:i64 user_id)
    CommentListResponse GetCommentListByVideo(1:i64 video_id)
    CommentActionResponse PostComment(1:CommentActionRequest comment_action_request)

}
