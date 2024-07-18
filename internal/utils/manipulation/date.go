package manipulation

import (
	"fmt"
	"time"
)

var (
	serverTimezone = ""
	serverLocation = time.Now().Location()
)

func GetServerTimezone() string {
	if serverTimezone != "" {
		return serverTimezone
	}

	now := time.Now()
	_, offset := now.Zone()
	serverTimezone = fmt.Sprintf("%+03d:%02d", offset/3600, offset%3600/60)
	return serverTimezone
}

func StartOfToday() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
}

func StartOfYesterday() time.Time {
	return StartOfToday().AddDate(0, 0, -1)
}

func IsToday(t time.Time) bool {
	today := StartOfToday()
	return t.Year() == today.Year() && t.Month() == today.Month() && t.Day() == today.Day()
}

func IsYesterday(t time.Time) bool {
	yesterday := StartOfYesterday()
	return t.Year() == yesterday.Year() && t.Month() == yesterday.Month() && t.Day() == yesterday.Day()
}

func ToSQLTimestamp(t time.Time) string {
	return t.UTC().Format("2006-01-02 15:04:05.999999-00:00")
}

func ToSQLDate(t time.Time) string {
	if t.Hour() > 0 || t.Minute() > 0 || t.Second() > 0 {
		t = t.Add(24 * time.Hour).Truncate(24 * time.Hour)
	}
	return t.UTC().Format("2006-01-02")
}

func ToSQLDateFrom(t time.Time) string {
	if t.Hour() > 0 || t.Minute() > 0 || t.Second() > 0 || t.Nanosecond() > 0 {
		t = t.Add(24 * time.Hour).Truncate(24 * time.Hour)
	} else {
		t = t.Truncate(24 * time.Hour)
	}
	return t.UTC().Format("2006-01-02")
}

func ToSQLDateTo(t time.Time) string {
	t = t.Truncate(24 * time.Hour)
	return t.UTC().Format("2006-01-02")
}
