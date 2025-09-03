package include

import (
	"fmt"
	"time"
)

var startTime time.Time

func GetUptime() int64 {
	return time.Now().Unix() - int64(time.Since(startTime).Seconds())
}

func GetUptimeFormatted() string {
	uptime := time.Since(startTime)

	days := int(uptime.Hours()) / 24
	hours := int(uptime.Hours()) % 24
	minutes := int(uptime.Minutes()) % 60
	seconds := int(uptime.Seconds()) % 60

	return fmt.Sprintf("%d days %d hours %d minutes %d seconds", days, hours, minutes, seconds)
}

func init() {
	startTime = time.Now()
}
