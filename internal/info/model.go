package info

type Info struct {
	Id           int    `json:"id"            db:"id"`
	Stream       string `json:"stream"        db:"stream"`
	StartPeriod  string `json:"start-period"  db:"start_period"`
	StopPeriod   string `json:"stop-period"   db:"stop_period"`
	LifeTime     string `json:"life-time"     db:"life_time"`
	UserId       int    `json:"user-id"       db:"user_id"`
	CreationDate string `json:"creation-date" db:"creation_date"`
}
