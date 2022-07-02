package model

type FileInfo struct {
	UserID int64
	FileID int64
	Hash   string // MD5, unique
	Size   int64
	PWD    string
}

type FileGroup struct {
	UserID        int64
	GroupID       int64
	GroupName     string
	GroupDescribe string
}

type FileGroupRelationship struct {
	GroupID int64
	FileID  int64
}

type FileSharing struct {
	FileID      int64
	UserGroupID int64
}
