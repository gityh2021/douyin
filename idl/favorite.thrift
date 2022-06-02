namespace go favorite

struct FavoriteActionRequest {
   1:i64 user_id
   2:string token
   3:i64 video_id
   4:i64 action_type
}
struct BaseResponse {
    1:i64 status_code
    2:string status_message
}
struct FavoriteListRequest {
   1:i64 user_id
   2:string token
}
struct FavoriteListResponse {
      1:BaseResponse base_resp
      2:list<Video> video_list
}
service FavoriteService{
    BaseResponse FavoriteByUser(1:FavoriteActionRequest request)
    FavoriteListResponse GetFavoriteListBYUser(1:FavoriteListRequest request)
}
