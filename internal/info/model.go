package info

type Info struct {
	Id          int    `json:"id"           db:"id"`
	CreateDate  string `json:"createDate"   db:"create_date"`
	Stream      string `json:"stream"       db:"stream"`
	StartPeriod string `json:"startPeriod"  db:"start_period"`
	StopPeriod  string `json:"stopPeriod"   db:"stop_period"`
	TimeLife    string `json:"timeLife"     db:"time_life"`
	UserId      int    `json:"userId"       db:"user_id"`
}
