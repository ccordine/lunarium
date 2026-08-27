package calendar

import (
	"strings"
	"testing"
)

func ancientInventoryByID(t *testing.T) map[string]Observance {
	t.Helper()
	byID := make(map[string]Observance)
	for _, event := range catalogOnlyAncientObservances() {
		if _, exists := byID[event.CatalogID]; exists {
			t.Fatalf("duplicate ancient inventory ID %q", event.CatalogID)
		}
		byID[event.CatalogID] = event
	}
	return byID
}

func TestUCLFestivalCoverageManifestIsClosedAndResolved(t *testing.T) {
	if got := len(uclFestivalCoverageManifest); got != 58 {
		t.Fatalf("UCL coverage manifest has %d rows, want the page-bounded inventory of 58", got)
	}
	if got := len(uclEgyptExpansionRecords); got != 38 {
		t.Fatalf("UCL expansion has %d records, want 38 source-layer records", got)
	}

	byID := ancientInventoryByID(t)
	seenRows := make(map[string]bool)
	mappedIDs := make(map[string]bool)
	for _, row := range uclFestivalCoverageManifest {
		if row.RowID == "" || row.SourceLabel == "" || len(row.CatalogIDs) == 0 {
			t.Errorf("incomplete UCL manifest row: %#v", row)
			continue
		}
		if seenRows[row.RowID] {
			t.Errorf("duplicate UCL source row %q", row.RowID)
		}
		seenRows[row.RowID] = true
		for _, id := range row.CatalogIDs {
			event, exists := byID[id]
			if !exists {
				t.Errorf("UCL row %q maps to missing catalog record %q", row.RowID, id)
				continue
			}
			mappedIDs[id] = true
			if event.CalendarCorpus != "Egyptian civil and temple calendars" {
				t.Errorf("UCL row %q maps outside the Egyptian corpus: %q", row.RowID, event.CalendarCorpus)
			}
			if event.NativeDateLabel == "" || event.AttestationLayer == "" || event.Era == "" || event.Site == "" {
				t.Errorf("UCL row %q maps to a record without explicit native/period/site layers: %#v", row.RowID, event)
			}
		}
	}
	for _, record := range uclEgyptExpansionRecords {
		if !mappedIDs[record.ID] {
			t.Errorf("UCL expansion record %q is not resolved from a manifest row", record.ID)
		}
	}
	if _, exists := byID["egypt-epagomenal-birthdays"]; exists {
		t.Error("obsolete grouped epagomenal record remains beside the five source rows")
	}

	epagomenal := map[string]string{
		"egypt-epagomenal-osiris":   "First day over the year",
		"egypt-epagomenal-horus":    "Second day over the year",
		"egypt-epagomenal-seth":     "Third day over the year",
		"egypt-epagomenal-isis":     "Fourth day over the year",
		"egypt-epagomenal-nephthys": "Fifth day over the year",
	}
	for id, nativeDate := range epagomenal {
		event, exists := byID[id]
		if !exists {
			t.Errorf("missing separate epagomenal record %q", id)
			continue
		}
		if !strings.Contains(event.NativeDateLabel, nativeDate) {
			t.Errorf("%s native date = %q, want %q", id, event.NativeDateLabel, nativeDate)
		}
	}
}

func TestCDLP9UrFestivalRowsStaySourceBounded(t *testing.T) {
	if got := len(urIIISourceInventoryRecords); got != 9 {
		t.Fatalf("Ur III expansion has %d records, want eight CDLP 9 rows plus one Tummal dossier", got)
	}
	byID := ancientInventoryByID(t)
	wantNative := map[string]string{
		"ur3-great-festival":                "month X",
		"ur3-gazelle-feast":                 "month II",
		"ur3-piglet-feast":                  "month III",
		"ur3-ubi-bird-feast":                "month IV",
		"ur3-ninazu-festival":               "month VI",
		"ur3-shulgi-festival":               "month VIII",
		"ur3-probable-month-nine-festival":  "month IX",
		"ur3-annunitum-recurring-festivals": "Probably each Ur lunar month",
	}
	for id, fragment := range wantNative {
		event, exists := byID[id]
		if !exists {
			t.Errorf("missing CDLP 9 inventory record %q", id)
			continue
		}
		if event.SourceURL != urIIICultSource || !strings.Contains(event.NativeDateLabel, fragment) {
			t.Errorf("%s escapes the cited CDLP 9 source resolution: %#v", id, event)
		}
		if !event.CatalogOnly || event.Date != "" || event.CalendarCorpus != "Ur III administrative cult calendars" {
			t.Errorf("%s was falsely projected or moved to another calendar: %#v", id, event)
		}
	}

	ninazu := byID["ur3-ninazu-festival"]
	if !strings.Contains(ninazu.HistoricalNote, "no ezem") || strings.Contains(ninazu.Name, "month V") {
		t.Errorf("month-V Ninazu cult activity was promoted into a named festival: %#v", ninazu)
	}
	annunitum := byID["ur3-annunitum-recurring-festivals"]
	if !strings.Contains(strings.ToLower(annunitum.DateConfidence), "probable") || !strings.Contains(strings.ToLower(annunitum.HistoricalNote), "probably") {
		t.Errorf("Annunitum recurrence lost CDLI's probability qualifier: %#v", annunitum)
	}
	monthNine := byID["ur3-probable-month-nine-festival"]
	if len(monthNine.Practices) != 0 || !strings.Contains(strings.ToLower(monthNine.Name), "probable unnamed") || !strings.Contains(strings.ToLower(monthNine.HistoricalNote), "old babylonian") {
		t.Errorf("probable month-IX evidence was given an invented identity or rite: %#v", monthNine)
	}
}

func TestTummalDossierRetainsInferentialConfidence(t *testing.T) {
	event, exists := ancientInventoryByID(t)["ur3-tummal-festival-dossier"]
	if !exists {
		t.Fatal("Tummal dossier is missing")
	}
	if event.SourceURL != urIIITummalSource || !strings.Contains(strings.ToLower(event.DateConfidence), "moderate") {
		t.Errorf("Tummal dossier overstates its source confidence: %#v", event)
	}
	if !strings.Contains(strings.ToLower(event.HistoricalNote), "usually do not explicitly") || event.Date != "" || !event.CatalogOnly {
		t.Errorf("Tummal delivery clustering was presented as an explicit fixed festival date: %#v", event)
	}
}

func TestBTTOFestivalListDoesNotTurnLineNumbersIntoDays(t *testing.T) {
	if got := len(bttoFestivalListRecords); got != 3 {
		t.Fatalf("BTTo expansion has %d records, want Ninurta, Sin, and Ishtar's Ululu", got)
	}
	byID := ancientInventoryByID(t)
	ids := []string{
		"babylon-ishtar-nisannu",
		"babylon-ninurta-akitu-nisannu",
		"babylon-sin-akitu-nisannu",
		"babylon-ishtar-ululu-nisannu",
	}
	for _, id := range ids {
		event, exists := byID[id]
		if !exists {
			t.Errorf("missing BTTo list record %q", id)
			continue
		}
		label := strings.ToLower(event.NativeDateLabel)
		if event.SourceURL != mesopotamianBTTO || !strings.Contains(label, "nisannu") || !(strings.Contains(label, "line") || strings.Contains(label, "list entry")) {
			t.Errorf("%s loses the Nisannu/list-line context: %#v", id, event)
		}
		if !(strings.Contains(label, "not a numbered") || strings.Contains(label, "no numbered") || strings.Contains(label, "without a numbered")) {
			t.Errorf("%s may misread a tablet line as a festival day: %q", id, event.NativeDateLabel)
		}
		if !strings.Contains(strings.ToLower(event.HistoricalNote), "scholarly") && !strings.Contains(strings.ToLower(event.HistoricalNote), "theological") {
			t.Errorf("%s loses the source's learned-list genre: %q", id, event.HistoricalNote)
		}
	}
	if event := byID["babylon-ishtar-nisannu"]; !strings.Contains(strings.ToLower(event.Name), "or ninurta") && !strings.Contains(event.Name, "—or Ninurta") {
		t.Errorf("BTTo iii 15 ambiguity was erased: %#v", event)
	}
	if event := byID["babylon-ishtar-ululu-nisannu"]; !strings.Contains(strings.ToLower(event.HistoricalNote), "festival name") || !strings.Contains(strings.ToLower(event.HistoricalNote), "not confused") {
		t.Errorf("Ulūlu was confused with a month rather than retained as the Nisannu-list festival: %#v", event)
	}
}

func TestUgariticRitualDocumentsRemainDocuments(t *testing.T) {
	if got := len(ugariticSourceInventoryRecords); got != 7 {
		t.Fatalf("Ugaritic expansion has %d records, want two calendar-bearing plus five catalog-only ritual documents", got)
	}
	byID := ancientInventoryByID(t)
	calendarBearing := []string{"ugarit-unnamed-month-ritual-146", "ugarit-unnamed-month-ritual-109"}
	for _, id := range calendarBearing {
		event, exists := byID[id]
		if !exists {
			t.Errorf("missing calendar-bearing Ugaritic tablet %q", id)
			continue
		}
		if event.ProjectionKind == "ritual-document-without-annual-date" || !strings.Contains(strings.ToLower(event.NativeDateLabel), "unnamed month") {
			t.Errorf("%s lost its partial calendar-bearing sequence: %#v", id, event)
		}
	}
	ritualDocuments := []string{
		"ugarit-palace-processional-143",
		"ugarit-civic-atonement-140",
		"ugarit-ancestral-kings-113",
		"ugarit-royal-sacrifice-115",
		"ugarit-athtartu-sacrifice-116",
	}
	for _, id := range ritualDocuments {
		event, exists := byID[id]
		if !exists {
			t.Errorf("missing Ugaritic ritual document %q", id)
			continue
		}
		if event.ProjectionKind != "ritual-document-without-annual-date" || !strings.Contains(strings.ToLower(event.ObservanceStatus), "not asserted") {
			t.Errorf("%s was promoted from ritual document to recurring holiday: %#v", id, event)
		}
		if event.Date != "" || !event.CatalogOnly || event.SourceURL != ugaritInstitutionalInventory {
			t.Errorf("%s must remain an institutional, catalog-only document record: %#v", id, event)
		}
	}
}

func TestRegionalGreekCatalogStaysNativeAndCycleOnly(t *testing.T) {
	if got := len(greekRegionalInventoryRecords); got != 9 {
		t.Fatalf("regional Greek inventory has %d records, want representative catalog of 9", got)
	}
	byID := ancientInventoryByID(t)
	ids := []string{
		"greek-pythian-games-delphi",
		"greek-olympic-games-olympia",
		"greek-spartan-karneia",
		"greek-spartan-hyakinthia",
		"greek-spartan-gymnopaidiai",
		"greek-delphic-theoxenia",
		"greek-delphic-septerion",
		"greek-argive-heraia",
		"greek-boeotian-daidala",
	}
	for _, id := range ids {
		event, exists := byID[id]
		if !exists {
			t.Errorf("missing regional Greek catalog record %q", id)
			continue
		}
		if event.CalendarCorpus != "Regional and Panhellenic Greek calendars" || event.Date != "" || !event.CatalogOnly {
			t.Errorf("%s was projected or moved into a universal calendar: %#v", id, event)
		}
		kind := strings.ToLower(event.ProjectionKind)
		if !strings.Contains(kind, "native") && !strings.Contains(kind, "cycle") {
			t.Errorf("%s lacks a native/cycle-only projection kind: %q", id, event.ProjectionKind)
		}
		if !strings.HasPrefix(event.SourceURL, "https://") || event.NativeDateLabel == "" || event.HistoricalNote == "" {
			t.Errorf("%s lacks its source-bounded native metadata: %#v", id, event)
		}
	}
	if byID["greek-pythian-games-delphi"].SourceURL != delphiOfficialFestivalSource {
		t.Error("Pythian Games record is not anchored to the official Delphi source")
	}
	if byID["greek-olympic-games-olympia"].SourceURL != olympiaOfficialGamesSource {
		t.Error("Olympic Games record is not anchored to the official Olympia source")
	}
	if event := byID["greek-boeotian-daidala"]; !strings.Contains(strings.ToLower(event.DateConfidence), "low to moderate") || !strings.Contains(strings.ToLower(event.HistoricalNote), "could not calculate") {
		t.Errorf("Daidala cycle lost Pausanias' stated uncertainty: %#v", event)
	}
}
