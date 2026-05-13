package core_config

import (
	"os"
	"time"
)

type TimeZone struct {
	TimeZone *time.Location
}

func GetConfig() *TimeZone {
	tz := os.Getenv("TIMEZONE")
	if tz == "" {
		tz = "UTC"
	}

	t, err := time.LoadLocation(tz)
	if err != nil {
		panic("Не удалось загрузить таймзону")
	}

	return &TimeZone{
		TimeZone: t,
	}
}
