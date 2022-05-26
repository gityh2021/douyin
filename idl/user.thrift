namespace go userdemo

struct BaseResp {
    1:i32 status_code
    2:string status_msg
    3:i64 service_time
}

struct User {
    1:i64 user_id
    2:string username
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
    2:i32 action_type
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

service UserService {
    CreateUserResponse CreateUser(1:CreateUserRequest req)
    CheckUserResponse CheckUser(1:CheckUserRequest req)
    InfoGetUserResponse InfoGetUser(1:InfoGetUserRequest req)
    MGetUserResponse MGetUser(1:MGetUserRequest req)
    UpdateUserResponse UpdateUser(1:UpdateUserRequest req)
}