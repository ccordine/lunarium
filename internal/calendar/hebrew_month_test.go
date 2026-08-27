package calendar

import (
	"testing"
	"time"
)

func TestJulianDayGregorianRoundTripSupportedRange(t *testing.T) {
	start := dateAt(1900, time.January, 1)
	end := dateAt(2100, time.December, 31)
	for date := start; !date.After(end); date = date.AddDate(0, 0, 1) {
		got := jdToGregorian(gregorianToJD(date))
		if !sameDay(got, date) {
			t.Fatalf("JD round trip for %s = %s", date.Format("2006-01-02"), got.Format("2006-01-02"))
		}
	}
}

func TestHebrewGregorianRoundTripSupportedRange(t *testing.T) {
	start := dateAt(1900, time.January, 1)
	end := dateAt(2100, time.December, 31)
	for date := start; !date.After(end); date = date.AddDate(0, 0, 1) {
		h := jdToHebrew(gregorianToJD(date))
		got := jdToGregorian(hebrewToJD(h.Year, h.Month, h.Day))
		if !sameDay(got, date) {
			t.Fatalf("Hebrew round trip for %s via %#v = %s", date.Format("2006-01-02"), h, got.Format("2006-01-02"))
		}
	}
}

func TestBuildHebrewMonthElul5786(t *testing.T) {
	month, err := BuildHebrewMonth(5786, 6, DefaultLocation)
	if err != nil {
		t.Fatal(err)
	}
	if month.CalendarSystem != HebrewCalendar || month.Year != 5786 || month.Month != 6 || month.Label != "Elul 5786" {
		t.Fatalf("month identity = %#v", month)
	}
	if month.StartDate != "2026-08-14" || month.EndDate != "2026-09-11" {
		t.Fatalf("Gregorian range = %s through %s", month.StartDate, month.EndDate)
	}
	if month.FirstWeekday != int(time.Friday) {
		t.Fatalf("first weekday = %d, want Friday", month.FirstWeekday)
	}
	if len(month.Days) != 29 {
		t.Fatalf("Elul days = %d, want 29", len(month.Days))
	}
	if month.Previous != (MonthReference{Year: 5786, Month: 5, Label: "Av 5786"}) {
		t.Fatalf("previous = %#v", month.Previous)
	}
	if month.Next != (MonthReference{Year: 5787, Month: 7, Label: "Tishrei 5787"}) {
		t.Fatalf("next = %#v", month.Next)
	}

	for index, day := range month.Days {
		wantHebrewDay := index + 1
		if day.SacredDates.HebrewYear != 5786 || day.SacredDates.HebrewMonthNumber != 6 || day.SacredDates.HebrewDay != wantHebrewDay {
			t.Fatalf("day %d sacred date = %#v", index, day.SacredDates)
		}
		wantGregorian := dateAt(2026, time.August, 14).AddDate(0, 0, index).Format("2006-01-02")
		if day.Date != wantGregorian {
			t.Fatalf("day %d Gregorian date = %s, want %s", index, day.Date, wantGregorian)
		}
	}
}

func TestBuildHebrewMonthLengths(t *testing.T) {
	tests := []struct {
		year  int
		month int
		days  int
	}{
		{year: 5786, month: 6, days: 29},
		{year: 5787, month: 7, days: 30},
		{year: 5784, month: 12, days: 30},
		{year: 5784, month: 13, days: 29},
		{year: 5785, month: 12, days: 29},
	}
	for _, test := range tests {
		response, err := BuildHebrewMonth(test.year, test.month, DefaultLocation)
		if err != nil {
			t.Errorf("BuildHebrewMonth(%d, %d): %v", test.year, test.month, err)
			continue
		}
		if len(response.Days) != test.days {
			t.Errorf("BuildHebrewMonth(%d, %d) days = %d, want %d", test.year, test.month, len(response.Days), test.days)
		}
		last := response.Days[len(response.Days)-1].SacredDates
		if last.HebrewDay != test.days || last.HebrewMonthNumber != test.month || last.HebrewYear != test.year {
			t.Errorf("BuildHebrewMonth(%d, %d) last sacred date = %#v", test.year, test.month, last)
		}
	}
}

func TestHebrewMonthReferencesAcrossYearAndAdarBoundaries(t *testing.T) {
	tests := []struct {
		name     string
		year     int
		month    int
		previous MonthReference
		next     MonthReference
	}{
		{
			name:     "Elul advances Hebrew year",
			year:     5786,
			month:    6,
			previous: MonthReference{Year: 5786, Month: 5, Label: "Av 5786"},
			next:     MonthReference{Year: 5787, Month: 7, Label: "Tishrei 5787"},
		},
		{
			name:     "Tishrei retreats Hebrew year",
			year:     5787,
			month:    7,
			previous: MonthReference{Year: 5786, Month: 6, Label: "Elul 5786"},
			next:     MonthReference{Year: 5787, Month: 8, Label: "Cheshvan 5787"},
		},
		{
			name:     "common Adar advances to Nisan",
			year:     5785,
			month:    12,
			previous: MonthReference{Year: 5785, Month: 11, Label: "Shevat 5785"},
			next:     MonthReference{Year: 5785, Month: 1, Label: "Nisan 5785"},
		},
		{
			name:     "common Nisan retreats to Adar",
			year:     5785,
			month:    1,
			previous: MonthReference{Year: 5785, Month: 12, Label: "Adar 5785"},
			next:     MonthReference{Year: 5785, Month: 2, Label: "Iyar 5785"},
		},
		{
			name:     "leap Adar I advances to Adar II",
			year:     5784,
			month:    12,
			previous: MonthReference{Year: 5784, Month: 11, Label: "Shevat 5784"},
			next:     MonthReference{Year: 5784, Month: 13, Label: "Adar II 5784"},
		},
		{
			name:     "leap Adar II advances to Nisan",
			year:     5784,
			month:    13,
			previous: MonthReference{Year: 5784, Month: 12, Label: "Adar I 5784"},
			next:     MonthReference{Year: 5784, Month: 1, Label: "Nisan 5784"},
		},
		{
			name:     "leap Nisan retreats to Adar II",
			year:     5784,
			month:    1,
			previous: MonthReference{Year: 5784, Month: 13, Label: "Adar II 5784"},
			next:     MonthReference{Year: 5784, Month: 2, Label: "Iyar 5784"},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := previousHebrewMonthReference(test.year, test.month); got != test.previous {
				t.Errorf("previous = %#v, want %#v", got, test.previous)
			}
			if got := nextHebrewMonthReference(test.year, test.month); got != test.next {
				t.Errorf("next = %#v, want %#v", got, test.next)
			}
		})
	}
}

func TestBuildHebrewMonthRejectsAdarIIInCommonYear(t *testing.T) {
	if _, err := BuildHebrewMonth(5785, 13, DefaultLocation); err == nil {
		t.Fatal("BuildHebrewMonth accepted Adar II in a common year")
	}
}

func TestBuildGregorianMonthIncludesUniformPeriodMetadata(t *testing.T) {
	month := BuildMonth(2026, time.August, DefaultLocation)
	if month.CalendarSystem != GregorianCalendar || month.StartDate != "2026-08-01" || month.EndDate != "2026-08-31" {
		t.Fatalf("Gregorian period metadata = %#v", month)
	}
	if month.Previous != (MonthReference{Year: 2026, Month: 7, Label: "July 2026"}) {
		t.Fatalf("previous = %#v", month.Previous)
	}
	if month.Next != (MonthReference{Year: 2026, Month: 9, Label: "September 2026"}) {
		t.Fatalf("next = %#v", month.Next)
	}
}
