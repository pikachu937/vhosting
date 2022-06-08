package logger

const (
	ErrLevelInfo    = "info"
	ErrLevelWarning = "warning"
	ErrLevelError   = "error"
	ErrLevelFatal   = "fatal"

	TypeOfUser  = "*user.User"
	TypeOfUsers = "map[int]*user.User"
	GotUser     = "Got user"
	GotAllUsers = "Got all users"

	TypeOfGroup  = "*group.Group"
	TypeOfGroups = "map[int]*group.Group"
	GotGroup     = "Got group"
	GotAllGroups = "Got all groups"

	TypeOfPermIds = "*permission.PermIds"
	GotUserPerms  = "Got user permissions"
	TypeOfPerms   = "map[int]*permission.Perm"
	GotAllPerms   = "Got all permissions"

	TypeOfInfo  = "*info.Info"
	GotInfo     = "Got info"
	TypeOfInfos = "map[int]*info.Info"
	GotAllInfos = "Got all infos"

	TypeOfVideo  = "*video.Video"
	GotVideo     = "Got video"
	TypeOfVideos = "map[int]*video.Video"
	GotAllVideos = "Got all videos"

	TypeOfGroupIds = "*group.GroupIds"
	GotUserGroups  = "Got user groups"

	TypeOfDownload = "*download.Download"
	GotDownload    = "Got download link"

	TableName     = "public.logs"
	Id            = "id"
	ClientID      = "client_id"
	ErrLevel      = "error_level"
	SessionOwner  = "session_owner"
	RequestMethod = "request_method"
	RequestPath   = "request_path"
	StatusCode    = "status_code"
	ErrCode       = "error_code"
	Message       = "message"
	CreationDate  = "creation_date"
)
