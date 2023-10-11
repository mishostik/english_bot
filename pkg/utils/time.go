package utils

import "time"

func GetMoscowTime() time.Time {
	loc, _ := time.LoadLocation("Europe/Moscow")
	return time.Now().In(loc)
}
