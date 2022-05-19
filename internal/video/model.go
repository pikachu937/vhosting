package video

type Video struct {
	Id           int    `json:"id"            db:"id"`
	Url          string `json:"url"           db:"url"`
	Filename     string `json:"file-name"     db:"file_name"`
	UserId       int    `json:"user-id"       db:"user_id"`
	InfoId       int    `json:"info-id"       db:"info_id"`
	CreationDate string `json:"creation-date" db:"creation_date"`
}
