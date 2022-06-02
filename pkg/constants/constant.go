package constants

const (
	// 数据库相关常量
	UserTableName     = "user"
	FollowerTableName = "follower"
	MySQLDefaultDSN   = "root:Yang75769933@tcp(139.224.195.12:3305)/douyin?parseTime=True&loc=Local"
	// JWT相关常量
	SecretKey   = "secret key"
	IdentityKey = "id"
	// Total       = "total"
	// Notes       = "notes"
	// NoteID      = "note_id"
	// 服务相关常量
	ApiServiceName            = "api"
	VideoServiceName          = "video"
	UserServiceName           = "user"
	EtcdAddress               = "139.224.195.12:2379"
	PlayURL                   = "http://127.0.0.1:8080/videos/"
	CoverURL                  = ""
	CPURateLimit      float64 = 80.0
	DefaultLimit              = 10
	QueryUserInfo     int32   = 1
	QueryFollowList   int32   = 2
	QueryFollowerList int32   = 3
	RelationAdd       int32   = 1
	RelationDel       int32   = 2
	NotLogin          int64   = -1
)

// const (
// 	NoteTableName           = "note"
// 	UserTableName           = "user"
// 	SecretKey               = "secret key"
// 	IdentityKey             = "id"
// 	Total                   = "total"
// 	Notes                   = "notes"
// 	NoteID                  = "note_id"
// 	ApiServiceName          = "api"
// 	NoteServiceName         = "note"
// 	UserServiceName         = "user"
// 	MySQLDefaultDSN         = "gorm:gorm@tcp(localhost:9910)/gorm?charset=utf8&parseTime=True&loc=Local"
// 	EtcdAddress             = "127.0.0.1:2379"
// 	CPURateLimit    float64 = 80.0
// 	DefaultLimit            = 10
// )
