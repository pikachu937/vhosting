package video

type Video struct {
	Id           int    `json:"id"            db:"id"`
	Url          string `json:"url"           db:"url"`
	Filename     string `json:"fileName"     db:"file_name"`
	UserId       int    `json:"userId"       db:"user_id"`
	InfoId       int    `json:"infoId"       db:"info_id"`
	CreationDate string `json:"creationDate" db:"creation_date"`
}
