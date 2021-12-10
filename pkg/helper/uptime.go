package helper

import (
	"fmt"
	"time"
)

var Uptime time.Time

func InitUptime() {
	Uptime = time.Now()
}

func GetUptime() string {
	uptimeSeconds := time.Since(Uptime).Seconds()

	days := uptimeSeconds / (60 * 60 * 24)
	hours := (uptimeSeconds - (days * 60 * 60 * 24)) / (60 * 60)
	minutes := ((uptimeSeconds - (days * 60 * 60 * 24)) - (hours * 60 * 60)) / 60

	return fmt.Sprintf(`%d days, %d hours, %d minutes, %d seconds`, int(hours), int(minutes), int(minutes), int(uptimeSeconds))
}
