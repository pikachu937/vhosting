package timestamp

import (
	"strings"
	"time"
)

func WriteThisTimestamp() string {
	var time string = time.Now().Round(time.Microsecond).String()
	timeWithoutGMT := time[:len(time)-4]
	gmtIncrement := timeWithoutGMT[len(timeWithoutGMT)-5:]
	timeWithoutInc := timeWithoutGMT[:len(timeWithoutGMT)-6]

	digits := len(timeWithoutInc) - 1 - strings.LastIndex(timeWithoutInc, ".")
	for i := digits; i < 6; i++ {
		timeWithoutInc += "0"
	}

	return timeWithoutInc + " " + gmtIncrement
}
