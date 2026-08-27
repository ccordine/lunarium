package calendar

import (
	"strings"
	"testing"
	"time"
)

func TestCalendarConversionsKnownDate(t *testing.T) {
	date := dateAt(2026, time.August, 27)
	h := jdToHebrew(gregorianToJD(date))
	if h.Year != 5786 || h.Month != 6 || h.Day != 14 {
		t.Fatalf("Hebrew conversion = %#v, want 14 Elul 5786", h)
	}
	i := jdToIslamic(gregorianToJD(date))
	if i.Year != 1448 || i.Month != 3 || i.Day != 14 {
		t.Fatalf("Islamic conversion = %#v, want 14 Rabi al-Awwal 1448", i)
	}
}

func TestEasterCalculations2026(t *testing.T) {
	if got := westernEaster(2026).Format("2006-01-02"); got != "2026-04-05" {
		t.Fatalf("western Easter = %s", got)
	}
	if got := orthodoxEaster(2026).Format("2006-01-02"); got != "2026-04-12" {
		t.Fatalf("Orthodox Pascha = %s", got)
	}
}

func TestDailyLenses(t *testing.T) {
	date := dateAt(2026, time.August, 27)
	moon := moonForDate(date)
	if moon.Phase != "Full Moon" || moon.Illumination < 98 {
		t.Fatalf("moon = %#v, expected near-full moon", moon)
	}
	numerology := numerologyForDate(date)
	if numerology.Score != 9 {
		t.Fatalf("numerology score = %d, want 9", numerology.Score)
	}
	if astrologyForDate(date).SunSign != "Virgo" {
		t.Fatal("August 27 should be represented as tropical Virgo")
	}
}

func TestSacredObservancesAppear(t *testing.T) {
	tests := []struct {
		date time.Time
		name string
	}{
		{dateAt(2026, time.April, 5), "Easter Sunday"},
		{dateAt(2026, time.April, 12), "Pascha"},
		{dateAt(2026, time.September, 12), "Rosh Hashanah"},
		{dateAt(2026, time.December, 25), "Christmas"},
	}
	for _, test := range tests {
		events := observancesForDate(test.date)
		if !eventNamed(events, test.name) {
			t.Errorf("%s missing event containing %q: %#v", test.date.Format("2006-01-02"), test.name, events)
		}
	}
}

func TestMonthIncludesSchedulesAndSources(t *testing.T) {
	month := BuildMonth(2026, time.August, DefaultLocation)
	if len(month.Days) != 31 {
		t.Fatalf("August days = %d", len(month.Days))
	}
	day := month.Days[26]
	if day.Observances == nil {
		t.Fatal("quiet-day observances must encode as an empty JSON array, not null")
	}
	if len(day.Prayers) != 3 {
		t.Fatalf("prayer schedules = %d", len(day.Prayers))
	}
	if !strings.HasSuffix(day.Reading.SourceURL, "/082726.cfm") {
		t.Fatalf("reading URL = %q", day.Reading.SourceURL)
	}
	if day.Reading.SundayCycle != "A" || day.Reading.WeekdayCycle != "II" {
		t.Fatalf("reading cycles = %#v", day.Reading)
	}
	for _, schedule := range day.Prayers {
		if len(schedule.Times) < 6 {
			t.Errorf("%s schedule has only %d times", schedule.Tradition, len(schedule.Times))
		}
	}
}

func TestEidAlAdhaAppearsOnFirstFestivalDay(t *testing.T) {
	index := BuildObservanceIndex(2026)
	for _, event := range index.Observances {
		if event.Name != "Eid al-Adha" {
			continue
		}
		date, err := time.Parse("2006-01-02", event.Date)
		if err != nil {
			t.Fatal(err)
		}
		if !eventNamed(observancesForDate(date), "Eid al-Adha") {
			t.Fatalf("Eid al-Adha missing from its first day %s", event.Date)
		}
		return
	}
	t.Fatal("Eid al-Adha missing from annual index")
}

func TestObservanceIndexHasAllFiveTraditions(t *testing.T) {
	index := BuildObservanceIndex(2026)
	minimum := map[Tradition]int{
		Christianity: 25,
		Judaism:      25,
		Islam:        18,
		Polytheist:   8,
		AncientWorld: 100,
	}
	for tradition, want := range minimum {
		if index.Counts[tradition] < want {
			t.Errorf("%s count = %d, want at least %d", tradition, index.Counts[tradition], want)
		}
		if len(index.Coverage[tradition]) == 0 {
			t.Errorf("%s has no index coverage description", tradition)
		}
	}
}

func TestObservanceIndexKeepsNativeOnlyRecordsUndated(t *testing.T) {
	index := BuildObservanceIndex(2026)
	seenCatalogOnly := false
	datedAfterCatalogOnly := false
	for _, event := range index.Observances {
		if event.CatalogOnly {
			seenCatalogOnly = true
			if event.Date != "" || event.NativeDateLabel == "" || event.ProjectionStatus == "" {
				t.Errorf("catalog-only event invents a date or lacks native projection metadata: %+v", event)
			}
			continue
		}
		if seenCatalogOnly {
			datedAfterCatalogOnly = true
		}
	}
	if !seenCatalogOnly {
		t.Fatal("annual index has no native-date-only ancient records")
	}
	if datedAfterCatalogOnly {
		t.Fatal("dated events must sort before catalog-only native records")
	}
}

func TestObservanceIndexCatalogIDsHaveOneGlobalIdentity(t *testing.T) {
	identities := make(map[string]string)
	for _, event := range BuildObservanceIndex(2026).Observances {
		if event.CatalogID == "" {
			t.Errorf("observance %q has an empty CatalogID", event.Name)
			continue
		}
		identity := event.Name + "|" + event.CalendarCorpus
		if previous, exists := identities[event.CatalogID]; exists && previous != identity {
			t.Errorf("catalog ID %q collides across %q and %q", event.CatalogID, previous, identity)
		}
		identities[event.CatalogID] = identity
	}
}

func eventNamed(events []Observance, fragment string) bool {
	for _, event := range events {
		if strings.Contains(event.Name, fragment) {
			return true
		}
	}
	return false
}
