package video

type Video struct {
	Id         int    `json:"id"         db:"id"`
	Url        string `json:"url"        db:"url"`
	File       string `json:"file"       db:"file"`
	CreateDate string `json:"createDate" db:"create_date"`
	InfoId     int    `json:"infoId"     db:"info_id"`
	UserId     int    `json:"userId"     db:"user_id"`
}
