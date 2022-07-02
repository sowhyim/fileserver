package model

type Userinfo struct {
	UserID      int64
	Account     string
	Username    string
	InviteCode  string
	InviterID   int64
	StoragePath string
	StorageSize int64 // 存储空间大小
}

// UserGroup 用户组群
type UserGroup struct {
	GroupID   int64
	GroupName string
	UserID    int64
}

type UserPassword struct {
	UserID   int64
	Password string
}
