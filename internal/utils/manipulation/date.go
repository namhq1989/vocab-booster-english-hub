package manipulation

import (
	"fmt"
	"time"
)

func getLocation(tz string) *time.Location {
	loc, err := time.LoadLocation(tz)
	if err != nil {
		fmt.Printf("invalid timezone '%s', defaulting to UTC: %v\n", tz, err)
		return time.UTC
	}
	return loc
}

func Now(tz string) time.Time {
	loc := getLocation(tz)
	return time.Now().In(loc)
}

func NowUTC() time.Time {
	return time.Now().UTC()
}

func StartOfToday(tz string) time.Time {
	loc := getLocation(tz)
	now := time.Now().In(loc)
	startOfToday := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, loc)
	return startOfToday.UTC()
}

func StartOfYesterday(tz string) time.Time {
	today := StartOfToday(tz)
	return today.AddDate(0, 0, -1)
}

func IsToday(t time.Time, tz string) bool {
	today := StartOfToday(tz)
	t = t.In(today.Location())
	return t.Year() == today.Year() && t.Month() == today.Month() && t.Day() == today.Day()
}

func IsYesterday(t time.Time, tz string) bool {
	yesterday := StartOfYesterday(tz)
	t = t.In(yesterday.Location())
	return t.Year() == yesterday.Year() && t.Month() == yesterday.Month() && t.Day() == yesterday.Day()
}

func ToSQLTimestamp(t time.Time, tz string) string {
	t = t.In(getLocation(tz))
	return t.Format("2006-01-02 15:04:05.999999-07:00")
}

func ToSQLDate(t time.Time, tz string) string {
	t = t.In(getLocation(tz))
	if t.Hour() > 0 || t.Minute() > 0 || t.Second() > 0 {
		t = t.Add(24 * time.Hour).Truncate(24 * time.Hour)
	}
	return t.Format("2006-01-02")
}

func ToSQLDateFrom(t time.Time, tz string) string {
	t = t.In(getLocation(tz))
	return t.Format("2006-01-02")
}

func ToSQLDateTo(t time.Time, tz string) string {
	t = t.In(getLocation(tz))
	return t.Format("2006-01-02")
}
