package calendar

import (
	"strings"
	"testing"
)

func TestHistoricalCatalogManifest(t *testing.T) {
	if got := len(megillatTaanitRules) + 1; got != 35 { // Hanukkah is the cross-month rule.
		t.Fatalf("Megillat Ta'anit manifest has %d entries/periods, want 35", got)
	}
	if got := len(woodOfferingRules); got != 9 {
		t.Fatalf("wood-offering manifest has %d dates, want 9", got)
	}

	seen := map[string]bool{}
	for _, group := range [][]historicalHebrewRule{megillatTaanitRules, woodOfferingRules, rabbinicCalendarRules} {
		for _, rule := range group {
			if rule.CatalogID == "" || rule.SourceURL == "" || rule.Status == "" {
				t.Fatalf("incomplete historical rule: %#v", rule)
			}
			if seen[rule.CatalogID] {
				t.Fatalf("duplicate catalog ID %q", rule.CatalogID)
			}
			seen[rule.CatalogID] = true
		}
	}
}

func TestEveryMegillatTaanitEntryMaterializes(t *testing.T) {
	const year = 5786 // common year exercises ordinary Adar handling.
	seen := map[string]bool{}
	for month := 1; month <= hebrewMonthsInYear(year); month++ {
		for day := 1; day <= hebrewMonthDays(year, month); day++ {
			date := jdToGregorian(hebrewToJD(year, month, day))
			for _, event := range historicalJewishObservances(date, hebrewDate{Year: year, Month: month, Day: day}) {
				if strings.HasPrefix(event.CatalogID, "mt-") {
					seen[event.CatalogID] = true
					if !event.Historical || event.Origin == "" || event.ObservanceStatus == "" || event.DateCertainty == "" {
						t.Fatalf("historical metadata missing from %#v", event)
					}
				}
			}
		}
	}
	if got := len(seen); got != 35 {
		t.Fatalf("materialized %d Megillat Ta'anit entries/periods, want 35: %#v", got, seen)
	}
}

func TestHistoricalAndLivingLayersCanOverlap(t *testing.T) {
	// Find an Adar 13 that is not Shabbat, so both the received Fast of Esther
	// and the discontinued Nicanor Day occupy the same calendar cell.
	for year := 5770; year <= 5800; year++ {
		month := 12
		if hebrewLeap(year) {
			month = 13
		}
		date := jdToGregorian(hebrewToJD(year, month, 13))
		events := jewishObservances(date)
		if hasNamedEvent(events, "Fast of Esther") && hasNamedEvent(events, "Megillat Ta'anit · Nicanor Day") {
			return
		}
	}
	t.Fatal("no test year materialized both Fast of Esther and Nicanor Day")
}

func TestTorahTempleLayersAndPurimKatan(t *testing.T) {
	nisan14 := jdToGregorian(hebrewToJD(5786, 1, 14))
	events := jewishObservances(nisan14)
	if !hasNamedEvent(events, "Korban Pesach · Passover Offering") || !hasNamedEvent(events, "Fast of the Firstborn · Ta'anit Bechorot") {
		t.Fatalf("Nisan 14 missing Torah/rabbinic layers: %#v", eventNames(events))
	}

	adarI14 := jdToGregorian(hebrewToJD(5784, 12, 14))
	if events := jewishObservances(adarI14); !hasNamedEvent(events, "Purim Katan") {
		t.Fatalf("Adar I 14 missing Purim Katan: %#v", eventNames(events))
	}
}

func TestMinorFastsDoNotClaimSunsetStart(t *testing.T) {
	for year := 5780; year <= 5790; year++ {
		date := jdToGregorian(hebrewToJD(year, 10, 10))
		for _, event := range jewishObservances(date) {
			if event.Name == "Tenth of Tevet · Asarah B'Tevet" {
				if event.StartsAtSunset {
					t.Fatal("Tenth of Tevet should be represented as a daytime fast")
				}
				return
			}
		}
	}
	t.Fatal("Tenth of Tevet did not materialize")
}

func TestPurimMeshulashGeneratesThreeDayOccurrence(t *testing.T) {
	for year := 5770; year <= 5900; year++ {
		month := 12
		if hebrewLeap(year) {
			month = 13
		}
		fifteenth := jdToGregorian(hebrewToJD(year, month, 15))
		if fifteenth.Weekday().String() != "Saturday" {
			continue
		}
		for offset := -1; offset <= 1; offset++ {
			date := fifteenth.AddDate(0, 0, offset)
			events := jewishObservances(date)
			if !hasNamedEvent(events, "Purim Meshulash · Walled Cities") {
				t.Fatalf("Purim Meshulash missing on %s", date.Format("2006-01-02"))
			}
		}
		return
	}
	t.Fatal("no Purim Meshulash year found in test range")
}

func TestTempleOfferingAndIsruChagLayers(t *testing.T) {
	shavuot := jdToGregorian(hebrewToJD(5786, 3, 6))
	if events := jewishObservances(shavuot); !hasNamedEvent(events, "Shavuot Tashlumin · Offering Make-up Days") {
		t.Fatalf("Shavuot missing tashlumin layer: %#v", eventNames(events))
	}
	israelAfterday := jdToGregorian(hebrewToJD(5786, 3, 7))
	if events := jewishObservances(israelAfterday); !hasNamedEvent(events, "Isru Chag · Israel") {
		t.Fatalf("Sivan 7 missing Israel Isru Chag: %#v", eventNames(events))
	}
}

func hasNamedEvent(events []Observance, name string) bool {
	for _, event := range events {
		if event.Name == name {
			return true
		}
	}
	return false
}

func eventNames(events []Observance) []string {
	names := make([]string, 0, len(events))
	for _, event := range events {
		names = append(names, event.Name)
	}
	return names
}
