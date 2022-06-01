namespace go video

struct BaseResp {
    1:i32 status_code
    2:string status_message
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

service VideoService{
    PublishListResponse GetPublishListByUser(1:i64 user_id)
}
