package manipulation

import (
	"fmt"
	"time"
)

func getLocation(timezone string) *time.Location {
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		fmt.Printf("invalid timezone '%s', defaulting to UTC: %v\n", timezone, err)
		return time.UTC
	}
	return loc
}

func Now(timezone string) time.Time {
	loc := getLocation(timezone)
	return time.Now().In(loc)
}

func NowUTC() time.Time {
	return time.Now().UTC()
}

func StartOfToday(timezone string) time.Time {
	loc := getLocation(timezone)
	now := time.Now().In(loc)
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, loc)
}

func StartOfYesterday(timezone string) time.Time {
	today := StartOfToday(timezone)
	return today.AddDate(0, 0, -1)
}

func IsToday(t time.Time, timezone string) bool {
	today := StartOfToday(timezone)
	t = t.In(today.Location())
	return t.Year() == today.Year() && t.Month() == today.Month() && t.Day() == today.Day()
}

func IsYesterday(t time.Time, timezone string) bool {
	yesterday := StartOfYesterday(timezone)
	t = t.In(yesterday.Location())
	return t.Year() == yesterday.Year() && t.Month() == yesterday.Month() && t.Day() == yesterday.Day()
}

func ToSQLTimestamp(t time.Time, timezone string) string {
	loc := getLocation(timezone)
	t = t.In(loc)
	return t.Format("2006-01-02 15:04:05.999999-07:00")
}

func ToSQLDate(t time.Time, timezone string) string {
	loc := getLocation(timezone)
	t = t.In(loc)
	if t.Hour() > 0 || t.Minute() > 0 || t.Second() > 0 {
		t = t.Add(24 * time.Hour).Truncate(24 * time.Hour)
	}
	return t.Format("2006-01-02")
}

func ToSQLDateFrom(t time.Time, timezone string) string {
	loc := getLocation(timezone)
	t = t.In(loc)
	if t.Hour() > 0 || t.Minute() > 0 || t.Second() > 0 || t.Nanosecond() > 0 {
		t = t.Add(24 * time.Hour).Truncate(24 * time.Hour)
	} else {
		t = t.Truncate(24 * time.Hour)
	}
	return t.Format("2006-01-02")
}

func ToSQLDateTo(t time.Time, timezone string) string {
	loc := getLocation(timezone)
	t = t.In(loc)
	t = t.Truncate(24 * time.Hour)
	return t.Format("2006-01-02")
}
