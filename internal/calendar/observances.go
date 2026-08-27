package calendar

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

const (
	usccbSource           = "https://www.usccb.org/committees/divine-worship/liturgical-calendar"
	ocaSource             = "https://www.oca.org/fs/paschal-cycle"
	hebcalSource          = "https://www.hebcal.com/holidays/"
	islamicCalendarSource = "https://www.britannica.com/topic/Islamic-calendar"
)

func observancesForDate(date time.Time) []Observance {
	events := make([]Observance, 0)
	events = append(events, christianObservances(date)...)
	events = append(events, imperialChristianObservances(date)...)
	events = append(events, jewishObservances(date)...)
	events = append(events, islamicObservances(date)...)
	events = append(events, romanObservances(date)...)
	events = append(events, atticObservances(date)...)
	events = append(events, modernPolytheistObservances(date)...)
	sort.SliceStable(events, func(i, j int) bool {
		if events[i].Tradition == events[j].Tradition {
			if events[i].Historical != events[j].Historical {
				return !events[i].Historical
			}
			return events[i].Name < events[j].Name
		}
		return events[i].Tradition < events[j].Tradition
	})
	return events
}

func baseObservance(name string, tradition Tradition, communities []string, category, summary, meaning string, practices, scripture []string, sourceName, sourceURL string) Observance {
	return Observance{
		CatalogID:   string(tradition) + "-" + slug(name),
		Name:        name,
		Tradition:   tradition,
		Communities: communities,
		Category:    category,
		Summary:     summary,
		Meaning:     meaning,
		Practices:   practices,
		Scripture:   scripture,
		SourceName:  sourceName,
		SourceURL:   sourceURL,
	}
}

func singleOccurrence(event Observance, date time.Time) Observance {
	if event.CatalogID == "" {
		event.CatalogID = string(event.Tradition) + "-" + slug(event.Name)
	}
	event.ID = event.CatalogID + "-" + date.Format("2006-01-02")
	event.Date = date.Format("2006-01-02")
	event.DurationDays = 1
	event.DayIndex = 1
	return event
}

func spanOccurrence(event Observance, current, start time.Time, duration int) Observance {
	if event.CatalogID == "" {
		event.CatalogID = string(event.Tradition) + "-" + slug(event.Name)
	}
	event.ID = event.CatalogID + "-" + start.Format("2006-01-02")
	event.Date = start.Format("2006-01-02")
	event.EndDate = start.AddDate(0, 0, duration-1).Format("2006-01-02")
	event.DurationDays = duration
	event.DayIndex = daysBetween(start, current) + 1
	return event
}

func daysBetween(a, b time.Time) int {
	a = time.Date(a.Year(), a.Month(), a.Day(), 12, 0, 0, 0, time.UTC)
	b = time.Date(b.Year(), b.Month(), b.Day(), 12, 0, 0, 0, time.UTC)
	return int(b.Sub(a).Hours() / 24)
}

func slug(value string) string {
	value = strings.ToLower(value)
	var out strings.Builder
	lastDash := false
	for _, r := range value {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') {
			out.WriteRune(r)
			lastDash = false
		} else if !lastDash && out.Len() > 0 {
			out.WriteByte('-')
			lastDash = true
		}
	}
	return strings.Trim(out.String(), "-")
}

func sameDay(a, b time.Time) bool {
	return a.Year() == b.Year() && a.YearDay() == b.YearDay()
}

func dateAt(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 12, 0, 0, 0, time.UTC)
}

func westernEaster(year int) time.Time {
	a := year % 19
	b := year / 100
	c := year % 100
	d := b / 4
	e := b % 4
	f := (b + 8) / 25
	g := (b - f + 1) / 3
	h := (19*a + b - d - g + 15) % 30
	i := c / 4
	k := c % 4
	l := (32 + 2*e + 2*i - h - k) % 7
	m := (a + 11*h + 22*l) / 451
	month := (h + l - 7*m + 114) / 31
	day := (h+l-7*m+114)%31 + 1
	return dateAt(year, time.Month(month), day)
}

func orthodoxEaster(year int) time.Time {
	a := year % 4
	b := year % 7
	c := year % 19
	d := (19*c + 15) % 30
	e := (2*a + 4*b - d + 34) % 7
	month := (d + e + 114) / 31
	day := (d+e+114)%31 + 1
	julianDifference := year/100 - year/400 - 2
	return dateAt(year, time.Month(month), day).AddDate(0, 0, julianDifference)
}

func adventStart(year int) time.Time {
	date := dateAt(year, time.December, 3)
	return date.AddDate(0, 0, -int(date.Weekday()))
}

func fmtRange(start time.Time, days int) string {
	if days <= 1 {
		return start.Format("Jan 2")
	}
	return fmt.Sprintf("%s–%s", start.Format("Jan 2"), start.AddDate(0, 0, days-1).Format("Jan 2"))
}
