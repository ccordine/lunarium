package calendar

import (
	"strings"
	"testing"
	"time"
)

func TestAncientCatalogCorpusCoverageAndMetadata(t *testing.T) {
	events := catalogOnlyAncientObservances()
	wantCounts := map[string]int{
		"Egyptian civil and temple calendars":      50,
		"Babylonian cult calendar":                 5,
		"Ur III administrative cult calendars":     12,
		"Mesopotamian lunar cult calendars":        1,
		"Seleucid Uruk ritual calendar":            2,
		"Ugaritic ritual tablets":                  11,
		"Old Norse textual calendar":               6,
		"Bede's Old English month record":          6,
		"Roman proclaimed and movable feriae":      4,
		"Regional and Panhellenic Greek calendars": 9,
	}
	if len(events) != 106 {
		t.Fatalf("ancient catalog has %d records, want 106", len(events))
	}

	counts := make(map[string]int)
	seen := make(map[string]bool)
	for _, event := range events {
		counts[event.CalendarCorpus]++
		if event.ID == "" || event.CatalogID == "" || event.ID != event.CatalogID {
			t.Errorf("%q has unstable catalog identity: id=%q catalogId=%q", event.Name, event.ID, event.CatalogID)
		}
		if seen[event.CatalogID] {
			t.Errorf("duplicate ancient catalog ID %q", event.CatalogID)
		}
		seen[event.CatalogID] = true
		if event.Tradition != AncientWorld {
			t.Errorf("%q tradition = %q, want %q", event.Name, event.Tradition, AncientWorld)
		}
		if !event.Historical || !event.CatalogOnly {
			t.Errorf("%q must be historical and catalog-only", event.Name)
		}
		if event.Date != "" || event.EndDate != "" {
			t.Errorf("%q invents a projected date: %q–%q", event.Name, event.Date, event.EndDate)
		}
		if event.NativeDateLabel == "" || event.AttestationLayer == "" || event.Era == "" || event.Site == "" {
			t.Errorf("%q missing native/attestation metadata: %#v", event.Name, event)
		}
		if event.ProjectionKind == "" || event.ProjectionStatus == "" || event.DateConfidence == "" {
			t.Errorf("%q missing projection metadata: %#v", event.Name, event)
		}
		if event.AnchorLocation == "" || event.DayBoundary == "" {
			t.Errorf("%q missing calendar anchor metadata", event.Name)
		}
		if event.HistoricalNote == "" || event.DateNote == "" || event.ObservanceStatus == "" {
			t.Errorf("%q missing source-critical caveat", event.Name)
		}
		if event.SourceName == "" || !strings.HasPrefix(event.SourceURL, "https://") {
			t.Errorf("%q has invalid source metadata: %q %q", event.Name, event.SourceName, event.SourceURL)
		}
		if event.Practices == nil {
			t.Errorf("%q must encode attested elements as an empty array rather than null", event.Name)
		}
	}

	for corpus, want := range wantCounts {
		if got := counts[corpus]; got != want {
			t.Errorf("%s count = %d, want %d", corpus, got, want)
		}
	}
	if len(counts) != len(wantCounts) {
		t.Errorf("unexpected corpus partition: %#v", counts)
	}
}

func TestRomanConceptivaeRemainUnprojected(t *testing.T) {
	want := map[string]bool{
		"roman-fornacalia-conceptiva": false,
		"roman-sementivae-paganalia":  false,
		"roman-feriae-latinae":        false,
		"roman-compitalia-conceptiva": false,
	}
	for _, event := range catalogOnlyAncientObservances() {
		if _, ok := want[event.CatalogID]; !ok {
			continue
		}
		want[event.CatalogID] = event.CatalogOnly && event.Date == "" && event.ProjectionKind == "annually-proclaimed-native-rule"
	}
	for id, valid := range want {
		if !valid {
			t.Errorf("movable Roman feriae %s was missing or falsely projected", id)
		}
	}
}

func TestEgyptianCatalogKeepsJosephFramingAtCollectionLevel(t *testing.T) {
	for _, event := range catalogOnlyAncientObservances() {
		if event.CalendarCorpus != "Egyptian civil and temple calendars" {
			continue
		}
		if strings.Contains(event.HistoricalNote, "Joseph") {
			t.Errorf("Egyptian record %q repeats collection-level biblical framing", event.Name)
		}
		if strings.Contains(strings.ToLower(event.ProjectionStatus), "gregorian anniversary") && !strings.Contains(strings.ToLower(event.ProjectionStatus), "not") {
			t.Errorf("Egyptian record %q overstates Gregorian projection: %q", event.Name, event.ProjectionStatus)
		}
	}
	foundCollectionCaveat := false
	for _, item := range About().Methodology {
		if strings.Contains(item, "Joseph cannot be securely") {
			foundCollectionCaveat = true
		}
	}
	if !foundCollectionCaveat {
		t.Error("About methodology omits the collection-level Joseph chronology caveat")
	}
}

func TestAncientNativeDatesRemainCatalogOnly(t *testing.T) {
	want := map[string]bool{
		"egypt-mont-middle-kingdom": false,
		"ur3-akiti-harvest":         false,
		"ugarit-first-wine-cycle":   false,
		"norse-jol-midwinter":       false,
		"anglosaxon-eosturmonath":   false,
	}
	for _, event := range catalogOnlyAncientObservances() {
		if _, tracked := want[event.CatalogID]; tracked {
			want[event.CatalogID] = event.CatalogOnly && event.Date == "" && event.NativeDateLabel != ""
		}
	}
	for id, valid := range want {
		if !valid {
			t.Errorf("representative catalog record %q missing or projected", id)
		}
	}
}

func TestLateUrukTabletRetainsItsTashrituSetting(t *testing.T) {
	tracked := map[string]bool{
		"uruk-anu-antu-akitu":      false,
		"uruk-clothing-procession": false,
	}
	for _, event := range catalogOnlyAncientObservances() {
		if _, ok := tracked[event.CatalogID]; !ok {
			continue
		}
		if !strings.Contains(event.NativeDateLabel, "Tašrītu") {
			t.Errorf("%s mislabels the P363711 Tašrītu sequence: %q", event.CatalogID, event.NativeDateLabel)
		}
		tracked[event.CatalogID] = true
	}
	for id, found := range tracked {
		if !found {
			t.Errorf("missing late Uruk record %s", id)
		}
	}
}

func TestUgariticTabletsRetainPartialNativeSchedules(t *testing.T) {
	want := map[string][]string{
		"ugarit-baal-royal-ritual-119":  {"Ibaalat", "day 7"},
		"ugarit-divine-image-entry-112": {"Hyr", "day 14", "day 17"},
	}
	for _, event := range catalogOnlyAncientObservances() {
		fragments, ok := want[event.CatalogID]
		if !ok {
			continue
		}
		for _, fragment := range fragments {
			if !strings.Contains(event.NativeDateLabel, fragment) {
				t.Errorf("%s native date %q omits %q", event.CatalogID, event.NativeDateLabel, fragment)
			}
		}
		if event.ProjectionKind == "attested-without-secure-date" || !strings.Contains(event.SourceName, "Pardee") {
			t.Errorf("%s still erases its partial schedule/source: %#v", event.CatalogID, event)
		}
		delete(want, event.CatalogID)
	}
	if len(want) != 0 {
		t.Fatalf("missing Ugaritic regressions: %#v", want)
	}
}

func TestModernPolytheistFixedConvention(t *testing.T) {
	want := map[string]string{
		"2026-02-01": "Winter Cross-Quarter",
		"2026-03-21": "Spring Equinox",
		"2026-05-01": "Spring Cross-Quarter",
		"2026-06-21": "Summer Solstice",
		"2026-08-01": "Summer Cross-Quarter",
		"2026-09-21": "Autumn Equinox",
		"2026-11-01": "Autumn Cross-Quarter",
		"2026-12-21": "Winter Solstice",
	}
	for rawDate, name := range want {
		date, err := time.Parse("2006-01-02", rawDate)
		if err != nil {
			t.Fatal(err)
		}
		events := modernPolytheistObservances(date)
		if len(events) != 1 {
			t.Fatalf("%s has %d modern Polytheist events, want 1", rawDate, len(events))
		}
		event := events[0]
		if !strings.Contains(event.Name, name) {
			t.Errorf("%s event = %q, want name containing %q", rawDate, event.Name, name)
		}
		if event.Tradition != Polytheist || event.Historical || event.CatalogOnly {
			t.Errorf("%s living event is conflated with ancient catalog: %#v", rawDate, event)
		}
		if event.CalendarCorpus != "Modern Neo-Pagan Wheel of the Year" || event.ProjectionKind != "documented-modern-fixed-date" {
			t.Errorf("%s missing modern-calendar provenance: %#v", rawDate, event)
		}
		if event.Date != rawDate || event.SourceURL != adfCalendarSource {
			t.Errorf("%s occurrence/source mismatch: %#v", rawDate, event)
		}
	}
}

func TestModernPolytheistYearHasEightDistinctHighDays(t *testing.T) {
	seen := make(map[string]bool)
	for date := dateAt(2026, time.January, 1); date.Year() == 2026; date = date.AddDate(0, 0, 1) {
		for _, event := range modernPolytheistObservances(date) {
			if seen[event.CatalogID] {
				t.Errorf("duplicate modern high day %q", event.CatalogID)
			}
			seen[event.CatalogID] = true
		}
	}
	if len(seen) != 8 {
		t.Fatalf("modern Polytheist calendar has %d high days, want 8: %#v", len(seen), seen)
	}
	if events := modernPolytheistObservances(dateAt(2026, time.January, 15)); events == nil || len(events) != 0 {
		t.Errorf("quiet modern Polytheist day must return a non-nil empty slice, got %#v", events)
	}
}

func TestAncientAndModernMidwinterRecordsStaySeparate(t *testing.T) {
	var ancientYule *Observance
	for _, event := range catalogOnlyAncientObservances() {
		if event.CatalogID == "norse-jol-midwinter" {
			copy := event
			ancientYule = &copy
			break
		}
	}
	if ancientYule == nil {
		t.Fatal("ancient Norse Jól record missing")
	}
	modern := modernPolytheistObservances(dateAt(2026, time.December, 21))
	if len(modern) != 1 {
		t.Fatalf("modern winter high day count = %d", len(modern))
	}
	if ancientYule.Tradition == modern[0].Tradition || ancientYule.CatalogID == modern[0].CatalogID {
		t.Fatalf("ancient Jól and living winter-solstice observance were conflated: %#v / %#v", ancientYule, modern[0])
	}
	if !strings.Contains(ancientYule.HistoricalNote, "not a license") || !strings.Contains(modern[0].DateNote, "not treated here as identical") {
		t.Fatal("ancient/modern separation caveats missing")
	}
}
