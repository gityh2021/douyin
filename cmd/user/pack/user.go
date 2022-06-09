package pack

import (
	"douyin/v1/cmd/user/dal/db"
	"douyin/v1/kitex_gen/user"
)

// User pack user info
func User(u *db.User, IsFollow bool) *user.User {
	if u == nil {
		return nil
	}

	return &user.User{Id: u.ID, Name: u.UserName, FollowCount: u.FollowCount, FollowerCount: u.FollowerCount, IsFollow: IsFollow}
}

func Users(us []*db.User, IsFollowList []bool) []*user.User {
	users := make([]*user.User, 0)
	for i, u := range us {
		if IsFollowList[i] { // == 1
			if user2 := User(u, true); user2 != nil {
				users = append(users, user2)
			}
		} else {
			if user2 := User(u, false); user2 != nil {
				users = append(users, user2)
			}
		}
	}
	return users
}
