package timestamp

import "time"

func MakeTimestamp() string {
	var time string = time.Now().Round(time.Microsecond).String()
	var lenOfTime int = len(time)
	return time[:lenOfTime-4]
}
