package convert

import "time"

func ToDateString(t time.Time) string {
	if (t == time.Time{}) {
		return "--"
	}
	return t.UTC().Format("02.01.2006")
}
