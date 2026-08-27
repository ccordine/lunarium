package calendar

import (
	"strconv"
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

func TestBoundedTorahTalmudCorpusManifests(t *testing.T) {
	if got := len(lateAppendedFastRules); got != 26 {
		t.Fatalf("printed Ma'amar Aharon manifest has %d dates, want 26", got)
	}
	if got := len(rainFastThresholdRules); got != 2 {
		t.Fatalf("dated rain-fast threshold manifest has %d records, want 2", got)
	}
	if got := len(unprojectedJewishCorpusRecords); got != 5 {
		t.Fatalf("unprojected Jewish institutional/conditional manifest has %d records, want 5", got)
	}

	seen := map[string]bool{}
	for _, rule := range append(append(append([]jewishCorpusRule{}, lateAppendedFastRules...), talmudFixedHistoricalRules...), rainFastThresholdRules...) {
		if rule.CatalogID == "" || rule.Name == "" || rule.Category == "" || rule.Status == "" || rule.SourceName == "" || rule.SourceURL == "" || rule.CalendarCorpus == "" || rule.Attestation == "" {
			t.Errorf("incomplete bounded Jewish dated rule: %#v", rule)
		}
		if seen[rule.CatalogID] {
			t.Errorf("duplicate bounded Jewish catalog ID %q", rule.CatalogID)
		}
		seen[rule.CatalogID] = true
	}
	for _, record := range unprojectedJewishCorpusRecords {
		if record.CatalogID == "" || record.Name == "" || record.Category == "" || record.NativeDate == "" || record.SourceName == "" || record.SourceURL == "" || record.CalendarCorpus == "" || record.Attestation == "" {
			t.Errorf("incomplete unprojected Jewish record: %#v", record)
		}
		if seen[record.CatalogID] {
			t.Errorf("duplicate bounded Jewish catalog ID %q", record.CatalogID)
		}
		seen[record.CatalogID] = true
	}
}

func TestEveryPrintedFastAppendixDateMaterializesFromNamedWitness(t *testing.T) {
	const year = 5786
	seen := map[string]Observance{}
	for month := 1; month <= hebrewMonthsInYear(year); month++ {
		for day := 1; day <= hebrewMonthDays(year, month); day++ {
			date := jdToGregorian(hebrewToJD(year, month, day))
			for _, event := range sourceBoundJewishObservances(date, hebrewDate{Year: year, Month: month, Day: day}) {
				if strings.HasPrefix(event.CatalogID, "mtb-") {
					seen[event.CatalogID] = event
				}
			}
		}
	}
	if got := len(seen); got != 26 {
		t.Fatalf("materialized %d printed fast-appendix dates, want 26: %#v", got, seen)
	}
	for catalogID, event := range seen {
		if !event.Historical || event.StartsAtSunset || event.CalendarCorpus != "Late appended fast list · Ma'amar Aharon" || !strings.Contains(event.HistoricalNote, "not part of the original 35-entry") || !strings.Contains(event.Era, "Geonic") {
			t.Errorf("%s erases appendix provenance or daytime-fast status: %#v", catalogID, event)
		}
	}
	if samuel := seen["mtb-iyar-samuel"]; !strings.Contains(samuel.HistoricalNote, "Iyar 28") {
		t.Errorf("Samuel date variant is not disclosed: %#v", samuel)
	}

	miriam := findCatalogEvent(sourceBoundJewishObservances(jdToGregorian(hebrewToJD(year, 1, 10)), hebrewDate{Year: year, Month: 1, Day: 10}), "mtb-nisan-miriam")
	if miriam == nil || !strings.Contains(miriam.HistoricalNote, "Hebrew witness reads") {
		t.Fatalf("10 Nisan Hebrew-witness reading not preserved: %#v", miriam)
	}
	if event := findCatalogEvent(sourceBoundJewishObservances(jdToGregorian(hebrewToJD(year, 1, 2)), hebrewDate{Year: year, Month: 1, Day: 2}), "mtb-nisan-miriam"); event != nil {
		t.Fatal("Miriam fast must follow the Hebrew witness's 10 Nisan, not the erroneous displayed English numeral")
	}
	eli := findCatalogEvent(sourceBoundJewishObservances(jdToGregorian(hebrewToJD(year, 2, 10)), hebrewDate{Year: year, Month: 2, Day: 10}), "mtb-iyar-eli")
	if eli == nil {
		t.Fatal("10 Iyar Eli entry did not materialize")
	}
}

func TestYomaMountGerizimTransmissionIsSeparateAndDated(t *testing.T) {
	const year = 5786
	tevet25 := jdToGregorian(hebrewToJD(year, 10, 25))
	event := findCatalogEvent(jewishObservances(tevet25), "yoma-tevet-gerizim")
	if event == nil {
		t.Fatal("Yoma 69a Tevet 25 Mount Gerizim tradition did not materialize")
	}
	if event.SourceURL != yomaGerizimSourceURL || !event.Historical || !strings.Contains(event.HistoricalNote, "21 Kislev") {
		t.Errorf("Yoma Mount Gerizim record lacks distinct transmission metadata: %#v", *event)
	}
	tevet21 := jdToGregorian(hebrewToJD(year, 10, 21))
	if findCatalogEvent(jewishObservances(tevet21), "yoma-tevet-gerizim") != nil {
		t.Fatal("Yoma's Tevet 25 transmission was silently moved to another date")
	}
}

func TestConditionalRainFastThresholdsAndRelativeStages(t *testing.T) {
	const year = 5786
	cheshvan17 := jdToGregorian(hebrewToJD(year, 8, 17))
	individual := findCatalogEvent(jewishObservances(cheshvan17), "rain-fast-individual-threshold")
	if individual == nil || !strings.Contains(individual.Category, "Conditional") || !strings.Contains(individual.DateNote, "not practical") {
		t.Fatalf("Marheshvan threshold is missing or presented as an automatic fast: %#v", individual)
	}
	kislev1 := jdToGregorian(hebrewToJD(year, 9, 1))
	communal := findCatalogEvent(jewishObservances(kislev1), "rain-fast-communal-threshold")
	if communal == nil || !strings.Contains(communal.Summary, "first three") {
		t.Fatalf("Rosh Chodesh Kislev communal threshold missing: %#v", communal)
	}

	catalog := map[string]Observance{}
	for _, event := range catalogOnlyJewishObservances() {
		catalog[event.CatalogID] = event
		if !event.CatalogOnly || event.Date != "" || event.NativeDateLabel == "" || event.ProjectionStatus == "" {
			t.Errorf("unprojected Jewish record fabricates or omits schedule metadata: %#v", event)
		}
	}
	for _, catalogID := range []string{"rain-fast-second-communal-stage", "rain-fast-final-seven-stage", "rain-fast-post-thirteen-mourning", "maamadot-weekly-cycle", "torah-jubilee-cycle"} {
		if _, ok := catalog[catalogID]; !ok {
			t.Errorf("missing unprojected Jewish corpus record %q", catalogID)
		}
	}
}

func TestFullOmerCountCarriesEveryDayIndex(t *testing.T) {
	checks := []struct {
		month int
		day   int
		want  int
	}{
		{1, 16, 1},
		{2, 18, 33},
		{3, 5, 49},
	}
	for _, check := range checks {
		date := jdToGregorian(hebrewToJD(5786, check.month, check.day))
		event := findCatalogEvent(jewishObservances(date), "torah-omer-count")
		if event == nil {
			t.Errorf("Omer day %d did not materialize", check.want)
			continue
		}
		if event.DayIndex != check.want || event.DurationDays != 49 || !strings.Contains(event.NativeDateLabel, "day "+strconv.Itoa(check.want)) {
			t.Errorf("Omer %d metadata = day %d/%d, label %q", check.want, event.DayIndex, event.DurationDays, event.NativeDateLabel)
		}
	}
	day50 := jdToGregorian(hebrewToJD(5786, 3, 6))
	if findCatalogEvent(jewishObservances(day50), "torah-omer-count") != nil {
		t.Fatal("Omer count incorrectly extends onto Shavuot/day 50")
	}
}

func TestMonthlyMoonBlessingBandPreservesBothTalmudicEndpoints(t *testing.T) {
	for _, day := range []int{1, 7, 16} {
		date := jdToGregorian(hebrewToJD(5786, 8, day))
		event := findCatalogEvent(jewishObservances(date), "talmud-birkat-levanah-window")
		if event == nil || event.DayIndex != day || event.DurationDays != 16 || !strings.Contains(event.HistoricalNote, "latest limits") {
			t.Errorf("Birkat HaLevanah study band day %d missing endpoint caveat: %#v", day, event)
		}
	}
	day17 := jdToGregorian(hebrewToJD(5786, 8, 17))
	if findCatalogEvent(jewishObservances(day17), "talmud-birkat-levanah-window") != nil {
		t.Fatal("Birkat HaLevanah Talmudic study band extends after day 16")
	}
}

func TestBirkatHachamaUsesReceivedTwentyEightYearCycle(t *testing.T) {
	for _, year := range []int{1925, 1953, 1981, 2009, 2037, 2065, 2093} {
		event := findCatalogEvent(jewishObservances(dateAt(year, 4, 8)), "talmud-birkat-hachama")
		if event == nil || event.StartsAtSunset || !strings.Contains(event.DateNote, "not the measured") {
			t.Errorf("Birkat HaChama %d occurrence/projection caveat missing: %#v", year, event)
		}
	}
	if event := findCatalogEvent(jewishObservances(dateAt(2010, 4, 8)), "talmud-birkat-hachama"); event != nil {
		t.Fatal("Birkat HaChama materialized outside its received 28-year cycle")
	}
}

func TestShemitahProjectedButJubileeRemainsUnprojected(t *testing.T) {
	shemitahStart := jdToGregorian(hebrewToJD(5782, 7, 1))
	event := findCatalogEvent(jewishObservances(shemitahStart), "torah-shemitah-cycle")
	if event == nil || event.Historical || !strings.Contains(event.NativeDateLabel, "5782") {
		t.Fatalf("received Shemitah cycle did not materialize correctly: %#v", event)
	}
	nonShemitah := jdToGregorian(hebrewToJD(5783, 7, 1))
	if findCatalogEvent(jewishObservances(nonShemitah), "torah-shemitah-cycle") != nil {
		t.Fatal("Shemitah marker materialized in a non-sabbatical year")
	}
	jubilee := findCatalogEvent(catalogOnlyJewishObservances(), "torah-jubilee-cycle")
	if jubilee == nil || !jubilee.CatalogOnly || jubilee.Date != "" || !strings.Contains(jubilee.ProjectionStatus, "disputed") {
		t.Fatalf("Jubilee must remain an explicit unprojected cycle record: %#v", jubilee)
	}
}

func TestFourDestructionFastsExposeRoshHashanahStatusLayer(t *testing.T) {
	targets := []struct {
		month int
		day   int
		name  string
	}{
		{7, 3, "Fast of Gedaliah"},
		{10, 10, "Tenth of Tevet · Asarah B'Tevet"},
		{4, 17, "Fast of the Seventeenth of Tammuz"},
		{5, 9, "Tisha B'Av"},
	}
	for _, target := range targets {
		var found *Observance
		base := jdToGregorian(hebrewToJD(5786, target.month, target.day))
		for offset := -2; offset <= 1 && found == nil; offset++ {
			for _, event := range jewishObservances(base.AddDate(0, 0, offset)) {
				if event.Name == target.name {
					copy := event
					found = &copy
					break
				}
			}
		}
		if found == nil {
			t.Errorf("living destruction fast %q did not materialize", target.name)
			continue
		}
		if found.SourceURL != fourFastsSourceURL || !strings.Contains(found.HistoricalNote, "peace") || !strings.Contains(found.HistoricalNote, "persecution") || !strings.Contains(found.DateNote, "halakhic guidance") {
			t.Errorf("%s lacks Rosh Hashanah 18b's conditional status caveat: %#v", target.name, *found)
		}
	}
}

func TestIsraelAndDiasporaSimchatTorahAreSeparateOccurrences(t *testing.T) {
	tishrei22 := jdToGregorian(hebrewToJD(5786, 7, 22))
	events22 := jewishObservances(tishrei22)
	if !hasNamedEvent(events22, "Shemini Atzeret") || !hasNamedEvent(events22, "Simchat Torah · Israel") || hasNamedEvent(events22, "Simchat Torah · Diaspora") {
		t.Fatalf("Tishrei 22 regional festival layers incorrect: %#v", eventNames(events22))
	}
	tishrei23 := jdToGregorian(hebrewToJD(5786, 7, 23))
	events23 := jewishObservances(tishrei23)
	if !hasNamedEvent(events23, "Simchat Torah · Diaspora") || hasNamedEvent(events23, "Simchat Torah · Israel") {
		t.Fatalf("Tishrei 23 regional festival layers incorrect: %#v", eventNames(events23))
	}
}

func TestAboutPublishesBoundedJewishInclusionMethod(t *testing.T) {
	about := About()
	method := strings.Join(about.Methodology, " ")
	if !strings.Contains(method, "bounded Torah/Talmud corpus") || !strings.Contains(method, "conditional schedules") || !strings.Contains(method, "Personal vows") {
		t.Errorf("About methodology omits bounded Jewish inclusion rule: %q", method)
	}
	if !strings.Contains(sourceLayerSummary(), "all 26 dates") || !strings.Contains(sourceLayerSummary(), "full 49-day Omer") {
		t.Errorf("coverage summary does not state finite Jewish manifests: %q", sourceLayerSummary())
	}
}

func findCatalogEvent(events []Observance, catalogID string) *Observance {
	for _, event := range events {
		if event.CatalogID == catalogID {
			copy := event
			return &copy
		}
	}
	return nil
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
