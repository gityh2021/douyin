package constants

const (
	// 数据库相关常量
	UserTableName     = "user"
	FollowerTableName = "follower"
	MySQLDefaultDSN   = "gorm:gorm@tcp(localhost:9910)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	// JWT相关常量
	SecretKey   = "secret key"
	IdentityKey = "id"
	// Total       = "total"
	// Notes       = "notes"
	// NoteID      = "note_id"
	// 服务相关常量
	ApiServiceName            = "demoapi"
	VideoServiceName          = "demovideo"
	UserServiceName           = "demouser"
	EtcdAddress               = "127.0.0.1:2379"
	CPURateLimit      float64 = 80.0
	DefaultLimit              = 10
	QueryUserInfo     int32   = 1
	QueryFollowList   int32   = 2
	QueryFollowerList int32   = 3
	RelationAdd       int32   = 1
	RelationDel       int32   = 2
)

// const (
// 	NoteTableName           = "note"
// 	UserTableName           = "user"
// 	SecretKey               = "secret key"
// 	IdentityKey             = "id"
// 	Total                   = "total"
// 	Notes                   = "notes"
// 	NoteID                  = "note_id"
// 	ApiServiceName          = "demoapi"
// 	NoteServiceName         = "demonote"
// 	UserServiceName         = "demouser"
// 	MySQLDefaultDSN         = "gorm:gorm@tcp(localhost:9910)/gorm?charset=utf8&parseTime=True&loc=Local"
// 	EtcdAddress             = "127.0.0.1:2379"
// 	CPURateLimit    float64 = 80.0
// 	DefaultLimit            = 10
// )
