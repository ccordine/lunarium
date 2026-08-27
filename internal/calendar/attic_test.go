package calendar

import (
	"strings"
	"testing"
	"time"
)

func TestAtticFestivalManifest(t *testing.T) {
	if got, want := len(atticFestivalRules), 37; got != want {
		t.Fatalf("Attic festival manifest has %d rules, want %d", got, want)
	}

	seen := make(map[string]bool, len(atticFestivalRules))
	for _, rule := range atticFestivalRules {
		if rule.CatalogID == "" || rule.Name == "" || rule.Month == "" || rule.StartDay < 1 {
			t.Errorf("incomplete Attic rule: %#v", rule)
		}
		if rule.Confidence != "high" && rule.Confidence != "medium" && rule.Confidence != "low" {
			t.Errorf("%s confidence = %q, want high, medium, or low", rule.CatalogID, rule.Confidence)
		}
		if seen[rule.CatalogID] {
			t.Errorf("duplicate Attic catalog ID %q", rule.CatalogID)
		}
		seen[rule.CatalogID] = true
	}

	for _, catalogID := range []string{
		"attic-kronia",
		"attic-greater-mysteries",
		"attic-thesmophoria",
		"attic-anthesteria",
		"attic-city-dionysia",
		"attic-panathenaia",
	} {
		if !seen[catalogID] {
			t.Errorf("Attic manifest missing %s", catalogID)
		}
	}
}

func TestAtticProjectedYearMaterializesEveryFestivalRule(t *testing.T) {
	year := atticYearProjection(2026)
	if len(year.Months) != 12 && len(year.Months) != 13 {
		t.Fatalf("projected Attic year has %d months", len(year.Months))
	}
	if year.Months[0].Name != "Hekatombaion" || year.Months[len(year.Months)-1].Name != "Skirophorion" {
		t.Fatalf("unexpected projected month bounds: %q through %q", year.Months[0].Name, year.Months[len(year.Months)-1].Name)
	}

	materialized := make(map[string]Observance)
	for day := year.Start; day.Before(year.End); day = day.AddDate(0, 0, 1) {
		for _, event := range atticObservances(day) {
			materialized[event.CatalogID] = event
		}
	}
	projectedCount := 0
	for _, rule := range atticFestivalRules {
		if !atticFestivalAnnotations[rule.CatalogID].CatalogOnly {
			projectedCount++
		}
	}
	if got, want := len(materialized), projectedCount; got != want {
		missing := make([]string, 0)
		for _, rule := range atticFestivalRules {
			if !atticFestivalAnnotations[rule.CatalogID].CatalogOnly {
				if _, ok := materialized[rule.CatalogID]; !ok {
					missing = append(missing, rule.CatalogID)
				}
			}
		}
		t.Fatalf("materialized %d/%d Attic rules; missing %v", got, want, missing)
	}

	for catalogID, event := range materialized {
		if event.Tradition != PolytheistAncient || !event.Historical || !event.StartsAtSunset {
			t.Errorf("%s lacks historical/tradition/day-boundary metadata: %#v", catalogID, event)
		}
		if event.CalendarCorpus != atticCalendarCorpus || event.NativeDateLabel == "" || event.ProjectionKind == "" || event.ProjectionStatus == "" {
			t.Errorf("%s lacks Attic projection metadata: %#v", catalogID, event)
		}
		if event.DateConfidence != "high" && event.DateConfidence != "medium" && event.DateConfidence != "low" {
			t.Errorf("%s confidence = %q", catalogID, event.DateConfidence)
		}
		if !strings.Contains(event.DateCertainty, "reconstruction") {
			t.Errorf("%s DateCertainty does not disclose reconstruction: %q", catalogID, event.DateCertainty)
		}
	}
}

func TestAtticContestedDatesAreNotFlattenedIntoFalseSpans(t *testing.T) {
	rules := make(map[string]atticFestivalRule)
	for _, rule := range atticFestivalRules {
		rules[rule.CatalogID] = rule
	}
	if got := rules["attic-proerosia"].StartDay; got != 6 {
		t.Errorf("Proerosia projection day = %d, want likely celebration day 6", got)
	}
	if got := rules["attic-greater-mysteries"]; got.StartDay != 15 || got.EndDay != 23 {
		t.Errorf("Greater Mysteries sequence = %d–%d, want 15–23", got.StartDay, got.EndDay)
	}
	if got := rules["attic-lenaia"]; got.StartDay != 12 || got.EndDay != 12 {
		t.Errorf("Lenaia materialized span = %d–%d, want attested day-12 anchor only", got.StartDay, got.EndDay)
	}
	if got := rules["attic-city-dionysia"]; got.StartDay != 10 || got.EndDay != 10 {
		t.Errorf("City Dionysia materialized span = %d–%d, want attested day-10 opening only", got.StartDay, got.EndDay)
	}
	if got := rules["attic-theogamia"].Name; got != "Hieros Gamos" {
		t.Errorf("classically attested name = %q, want Hieros Gamos", got)
	}

	catalog := catalogOnlyAtticObservances()
	if len(catalog) != 9 {
		t.Fatalf("catalog-only Attic alternatives = %d, want 9", len(catalog))
	}
	seen := make(map[string]Observance)
	for _, event := range catalog {
		seen[event.CatalogID] = event
	}
	for _, id := range []string{"attic-oschophoria", "attic-bendideia"} {
		event, ok := seen[id]
		if !ok || !event.CatalogOnly || event.Date != "" || !strings.Contains(event.NativeDateLabel, "or") {
			t.Errorf("%s alternative date was not preserved honestly: %#v", id, event)
		}
	}
}

func TestAdditionalAtticCoreKeepsUncertainDatesNativeOnly(t *testing.T) {
	wantCatalogOnly := map[string]bool{
		"attic-metageitnia":         false,
		"attic-apatouria":           false,
		"attic-pompaia-maimakteria": false,
		"attic-rural-dionysia":      false,
		"attic-lesser-mysteries":    false,
		"attic-kallynteria":         false,
		"attic-arrephoria":          false,
	}
	for _, event := range catalogOnlyAtticObservances() {
		if _, ok := wantCatalogOnly[event.CatalogID]; ok {
			wantCatalogOnly[event.CatalogID] = event.CatalogOnly && event.Date == "" && event.NativeDateLabel != ""
		}
	}
	for id, valid := range wantCatalogOnly {
		if !valid {
			t.Errorf("uncertain Attic core record %s was not retained native-only", id)
		}
	}

	for day := atticYearProjection(2026).Start; day.Before(atticYearProjection(2026).End); day = day.AddDate(0, 0, 1) {
		for _, event := range atticObservances(day) {
			if event.CatalogID == "attic-pandia" {
				if event.NativeDateLabel == "" || event.DateConfidence != "medium" {
					t.Fatalf("Pandia projection lacks its qualified native date: %#v", event)
				}
				return
			}
		}
	}
	t.Fatal("Pandia did not materialize in the bounded Attic projection")
}

func TestAtticProjectionIsBoundedAndReturnsSupportedYears(t *testing.T) {
	for _, year := range []int{1899, 1900, 2000, 2100} {
		months := len(atticYearProjection(year).Months)
		if months != 12 && months != 13 {
			t.Errorf("supported year %d produced %d months", year, months)
		}
	}
	if got := len(atticYearProjection(1800).Months); got != 0 {
		t.Errorf("unsupported projection materialized %d months", got)
	}
}

func TestAtticIntercalationInsertsPoseideonII(t *testing.T) {
	var leap atticProjectedYear
	found := false
	for anchorYear := 2000; anchorYear <= 2100; anchorYear++ {
		candidate := atticYearProjection(anchorYear)
		if len(candidate.Months) == 13 {
			leap = candidate
			found = true
			break
		}
	}
	if !found {
		t.Fatal("mean-lunation projection found no 13-month Attic year in 2000–2100")
	}

	if leap.Months[5].Name != "Poseideon" || leap.Months[6].Name != "Poseideon II" || !leap.Months[6].Intercalary || leap.Months[7].Name != "Gamelion" {
		t.Fatalf("intercalary order = %q, %q, %q", leap.Months[5].Name, leap.Months[6].Name, leap.Months[7].Name)
	}

	// The intercalary month is handled as a calendar month, not invented as a
	// festival: ordinary Poseideon rules must not be duplicated into it.
	poseideonII := leap.Months[6]
	for day := poseideonII.Start; day.Before(poseideonII.End); day = day.AddDate(0, 0, 1) {
		for _, event := range atticObservances(day) {
			if event.CatalogID == "attic-haloa" {
				t.Fatalf("Haloa was incorrectly duplicated in Poseideon II on %s", day.Format("2006-01-02"))
			}
		}
	}
}

func TestAtticProjectionUsesSolsticeAnchoredLunarYear(t *testing.T) {
	year := atticYearProjection(2026)
	if year.Start.Month() != time.June && year.Start.Month() != time.July {
		t.Fatalf("Hekatombaion projection begins %s, want first lunation after June solstice", year.Start.Format("2006-01-02"))
	}
	if year.Start.Before(atticCivilDay(atticTimeFromJulianDay(atticNorthernSummerSolsticeJD(2026)))) {
		t.Fatalf("projected new year %s precedes summer solstice proxy", year.Start.Format("2006-01-02"))
	}
}
