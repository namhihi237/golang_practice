package utils

import "time"

func ConvertStringToTime(str string) (*time.Time, error) {
	time, err := time.Parse("2006-01-02", str)
	return &time, err
}
