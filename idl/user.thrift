namespace go user

struct BaseResp {
    1:i32 status_code
    2:string status_msg
    3:i64 service_time
}

struct User {
    1:i64 id
    2:string name
    3:i64 follow_count
    4:i64 follower_count
    5:bool is_follow
}

struct CreateUserRequest {
    1:string username
    2:string password
}

struct CreateUserResponse {
    1:BaseResp base_resp
}

struct CheckUserRequest {
    1:string username
    2:string password
}

struct CheckUserResponse {
    1:BaseResp base_resp
    2:i64 user_id
}

struct InfoGetUserRequest {
    1:i64 user_id
    2:i64 to_user_id
}

struct InfoGetUserResponse {
    1:BaseResp base_resp
    2:User user
}

struct MGetUserRequest {
    1:i64 user_id
    2:i64 to_user_id
    3:i32 action_type
}

struct MGetUserResponse {
    1:BaseResp base_resp
    2:list<User> users
}

struct UpdateUserRequest {
    1:i64 user_id
    2:i64 to_user_id
    3:i32 action_type
}

struct UpdateUserResponse {
    1:BaseResp base_resp
}

// 根据userid list获取user info list
// 调用方法：userClient.GetUserInfoList(ctx, req)
struct GetUserInfoListRequest {
    1:list<i64> user_ids
    2:i64 user_id// 登录用户的userid（用于判断是否已关注视频作者）
}

struct GetUserInfoListResponse {
    1:BaseResp base_resp
    2:list<User> users
}

service UserService {
    CreateUserResponse CreateUser(1:CreateUserRequest req)
    CheckUserResponse CheckUser(1:CheckUserRequest req)
    InfoGetUserResponse InfoGetUser(1:InfoGetUserRequest req)
    MGetUserResponse MGetUser(1:MGetUserRequest req)
    UpdateUserResponse UpdateUser(1:UpdateUserRequest req)
    GetUserInfoListResponse GetUserInfoList(1:GetUserInfoListRequest req)
}