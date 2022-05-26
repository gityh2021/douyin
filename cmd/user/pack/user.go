package pack

import (
	"github.com/Baojiazhong/dousheng-ubuntu/cmd/user/dal/db"
	"github.com/Baojiazhong/dousheng-ubuntu/kitex_gen/userdemo"
)

// User pack user info
func User(u *db.User, IsFollow bool) *userdemo.User {
	if u == nil {
		return nil
	}

	return &userdemo.User{UserId: u.ID, Username: u.UserName, FollowCount: u.FollowCount, FollowerCount: u.FollowerCount, IsFollow: IsFollow}
}

// Users pack list of user info
// func Users(us []*db.User, IsFollow bool) []*userdemo.User {
// 	users := make([]*userdemo.User, 0)
// 	for _, u := range us {
// 		if user2 := User(u, IsFollow); user2 != nil {
// 			users = append(users, user2)
// 		}
// 	}
// 	return users
// }

func Users(us []*db.User, IsFollowList []int64) []*userdemo.User {
	users := make([]*userdemo.User, 0)
	for i, u := range us {
		if IsFollowList[i] > 0 { // == 1
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
