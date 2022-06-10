package info

type Info struct {
	Id           int    `json:"id"            db:"id"`
	Stream       string `json:"stream"        db:"stream"`
	StartPeriod  string `json:"startPeriod"  db:"start_period"`
	StopPeriod   string `json:"stopPeriod"   db:"stop_period"`
	LifeTime     string `json:"lifeTime"     db:"life_time"`
	UserId       int    `json:"userId"       db:"user_id"`
	CreationDate string `json:"creationDate" db:"creation_date"`
}
