package calendar

import (
	"strings"
	"testing"
	"time"
)

func TestRomanFixedCanonHasAllFortyFiveNamedRows(t *testing.T) {
	if got := len(romanFixedCanon); got != 45 {
		t.Fatalf("roman fixed canon has %d rows, want 45", got)
	}
	wantByMonth := map[time.Month]int{
		time.January: 3, time.February: 6, time.March: 4,
		time.April: 5, time.May: 5, time.June: 2,
		time.July: 5, time.August: 6, time.October: 3,
		time.December: 6,
	}
	gotByMonth := make(map[time.Month]int)
	seenIDs := make(map[string]bool)
	seenDates := make(map[string]bool)
	for _, rule := range romanFixedCanon {
		if seenIDs[rule.CatalogID] {
			t.Errorf("duplicate Roman catalog id %q", rule.CatalogID)
		}
		seenIDs[rule.CatalogID] = true
		gotByMonth[rule.Month]++
		dateKey := dateAt(2026, rule.Month, rule.Day).Format("01-02")
		if seenDates[dateKey] {
			t.Errorf("duplicate fixed-canon calendar row for %s", dateKey)
		}
		seenDates[dateKey] = true
		if rule.Name == "" || rule.Summary == "" || rule.Meaning == "" || len(rule.Practices) == 0 {
			t.Errorf("%s lacks descriptive content", rule.CatalogID)
		}
	}
	for month, want := range wantByMonth {
		if got := gotByMonth[month]; got != want {
			t.Errorf("%s has %d fixed rows, want %d", month, got, want)
		}
	}
}

func TestEveryRomanFixedRuleMaterializesWithProjectionMetadata(t *testing.T) {
	for _, rule := range romanFixedCanon {
		date := dateAt(2026, rule.Month, rule.Day)
		var found *Observance
		for _, event := range romanObservances(date) {
			if event.CatalogID == rule.CatalogID {
				copy := event
				found = &copy
				break
			}
		}
		if found == nil {
			t.Errorf("%s did not materialize on %s", rule.CatalogID, date.Format("2006-01-02"))
			continue
		}
		if !found.Historical || found.Tradition != PolytheistAncient {
			t.Errorf("%s missing historical/polytheist classification", rule.CatalogID)
		}
		if found.Date != date.Format("2006-01-02") || found.DurationDays != 1 {
			t.Errorf("%s occurrence date/duration = %s/%d", rule.CatalogID, found.Date, found.DurationDays)
		}
		if found.CalendarCorpus == "" || found.NativeDateLabel == "" || found.AttestationLayer == "" || found.ProjectionKind == "" || found.DateConfidence == "" {
			t.Errorf("%s lacks provenance/projection metadata: %+v", rule.CatalogID, *found)
		}
	}
}

func TestLaterRomanTierIsSeparateAndManifested(t *testing.T) {
	if got := len(romanLaterFixedFestivals); got != 19 {
		t.Fatalf("later/regional Roman tier has %d records, want 19", got)
	}
	seen := make(map[string]bool)
	for _, rule := range romanLaterFixedFestivals {
		if seen[rule.CatalogID] {
			t.Errorf("duplicate later-tier catalog id %q", rule.CatalogID)
		}
		seen[rule.CatalogID] = true
		event := findClassicalEvent(romanObservances(dateAt(2026, rule.Month, rule.Day)), rule.CatalogID)
		if event == nil {
			t.Errorf("%s did not materialize", rule.CatalogID)
			continue
		}
		if !event.Historical || event.Category != "Later or regional Roman observance" {
			t.Errorf("%s was folded into the wrong status/tier: %+v", rule.CatalogID, *event)
		}
		if event.CalendarCorpus == "Roman Republican fasti · archaic fixed canon" || event.AttestationLayer == "" || event.DateConfidence == "" {
			t.Errorf("%s lacks distinct later-tier provenance: %+v", rule.CatalogID, *event)
		}
		if event.Site == "" || event.AnchorLocation == "" {
			t.Errorf("%s lacks record-specific location metadata: %+v", rule.CatalogID, *event)
		}
	}
	if !seen["roman-augustalia"] || !seen["roman-natalis-invicti"] || !seen["roman-navigium-isidis"] {
		t.Errorf("key imperial/regional records are missing: %v", seen)
	}
	for _, id := range []string{"roman-mundus-patet-august", "roman-mundus-patet-october", "roman-mundus-patet-november", "roman-septimontium"} {
		if !seen[id] {
			t.Errorf("source-bounded later Roman record %s is missing", id)
		}
	}
	nemoralia := findClassicalEvent(romanObservances(dateAt(2026, time.August, 13)), "roman-nemoralia")
	if nemoralia == nil || !strings.Contains(nemoralia.Site, "Lake Nemi") || !strings.Contains(nemoralia.AnchorLocation, "Aricia") {
		t.Errorf("Nemoralia lost its Nemi/Aricia setting: %+v", nemoralia)
	}
}

func TestRomanDateLabelsUseInclusiveRomanReckoning(t *testing.T) {
	cases := []struct {
		month time.Month
		day   int
		want  string
	}{
		{time.January, 1, "Kal. Ian."},
		{time.January, 4, "prid. Non. Ian."},
		{time.January, 9, "a.d. V Id. Ian."},
		{time.March, 14, "prid. Id. Mart."},
		{time.April, 25, "a.d. VII Kal. Mai."},
		{time.February, 27, "a.d. III Kal. Mart."},
		{time.December, 23, "a.d. X Kal. Ian."},
	}
	for _, tc := range cases {
		if got := romanDateLabel(tc.month, tc.day); got != tc.want {
			t.Errorf("romanDateLabel(%s, %d) = %q, want %q", tc.month, tc.day, got, tc.want)
		}
	}
}

func TestRomanExtendedSpansCrossMonthBoundaries(t *testing.T) {
	floralia := findClassicalEvent(romanObservances(dateAt(2026, time.May, 1)), "roman-floralia")
	if floralia == nil {
		t.Fatal("Floralia missing on May 1")
	}
	if floralia.Date != "2026-04-28" || floralia.EndDate != "2026-05-03" || floralia.DayIndex != 4 {
		t.Errorf("Floralia span = %s..%s day %d", floralia.Date, floralia.EndDate, floralia.DayIndex)
	}

	sullan := findClassicalEvent(romanObservances(dateAt(2026, time.November, 1)), "roman-ludi-victoriae-sullanae")
	if sullan == nil || sullan.Date != "2026-10-26" || sullan.DayIndex != 7 {
		t.Errorf("Sullan games cross-month occurrence = %+v", sullan)
	}

	brumalia := findClassicalEvent(romanObservances(dateAt(2026, time.December, 1)), "byzantine-brumalia")
	if brumalia == nil || brumalia.Site != "Constantinople's court and urban communities" || brumalia.AnchorLocation != "Constantinople" {
		t.Errorf("Brumalia lost its Constantinopolitan setting: %+v", brumalia)
	}
}

func TestImperialChristianRecurringRules(t *testing.T) {
	cases := []struct {
		date      time.Time
		catalogID string
		corpus    string
		living    bool
	}{
		{dateAt(2026, time.January, 25), "hre-karlsfest", "Holy Roman Empire · Aachen local calendar", true},
		{dateAt(2026, time.March, 3), "hre-cunigunde-bamberg", "Holy Roman Empire · Bamberg diocesan calendar", true},
		{dateAt(2026, time.March, 1), "byzantine-sunday-orthodoxy", "Byzantine / Eastern Orthodox Paschal cycle", true},
		{dateAt(2026, time.April, 17), "hre-holy-lance-nails", "Holy Roman Empire · Charles IV relic calendar", false},
		{dateAt(2026, time.May, 11), "byzantine-constantinople-dedication", "Byzantine Constantinopolitan calendar", true},
		{dateAt(2026, time.June, 4), "hre-corpus-christi-civic", "Holy Roman Empire · Catholic civic calendars", false},
		{dateAt(2026, time.June, 25), "hre-augsburg-confession", "Holy Roman Empire · Lutheran confessional calendar", true},
		{dateAt(2026, time.July, 13), "hre-henry-ii-bamberg", "Holy Roman Empire · Bamberg diocesan calendar", true},
		{dateAt(2026, time.August, 8), "hre-augsburg-peace-festival", "Holy Roman Empire · Free Imperial City of Augsburg", true},
		{dateAt(2026, time.September, 1), "byzantine-indiction", "Byzantine indiction calendar", true},
	}
	for _, tc := range cases {
		event := findClassicalEvent(imperialChristianObservances(tc.date), tc.catalogID)
		if event == nil {
			t.Errorf("%s missing on %s", tc.catalogID, tc.date.Format("2006-01-02"))
			continue
		}
		if event.CalendarCorpus != tc.corpus || event.DateConfidence == "" || event.SourceURL == "" {
			t.Errorf("%s metadata = %+v", tc.catalogID, *event)
		}
		if event.Historical == tc.living {
			t.Errorf("%s historical=%v, want %v for living=%v", tc.catalogID, event.Historical, !tc.living, tc.living)
		}
	}
	karlsfest := findClassicalEvent(imperialChristianObservances(dateAt(2026, time.January, 25)), "hre-karlsfest")
	if karlsfest == nil || !strings.Contains(karlsfest.NativeDateLabel, "Last Sunday") || !strings.Contains(karlsfest.DateNote, "January 28") {
		t.Errorf("Karlsfest lacks its living Sunday rule and nominal anniversary: %+v", karlsfest)
	}
	if event := findClassicalEvent(imperialChristianObservances(dateAt(2026, time.January, 28)), "hre-karlsfest"); event != nil {
		t.Errorf("Karlsfest incorrectly materialized its current solemn Mass on the nominal anniversary: %+v", event)
	}
}

func TestImperialLayerDoesNotInventCoronationHoliday(t *testing.T) {
	for _, event := range imperialChristianObservances(dateAt(2026, time.December, 25)) {
		if event.CatalogID == "hre-charlemagne-coronation" || event.CatalogID == "hre-imperial-coronation" {
			t.Errorf("known coronation anniversary was incorrectly materialized as recurring holiday: %+v", event)
		}
	}
}

func findClassicalEvent(events []Observance, catalogID string) *Observance {
	for _, event := range events {
		if event.CatalogID == catalogID {
			copy := event
			return &copy
		}
	}
	return nil
}
