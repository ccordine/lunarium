package calendar

import "time"

const (
	sefariaMegillatTaanit = "https://www.sefaria.org/Megillat_Ta%27anit?tab=contents"
	sefariaMishnahTaanit  = "https://www.sefaria.org/Mishnah_Ta%27anit.4.5"
	sefariaMishnahRH      = "https://www.sefaria.org/Mishnah_Rosh_Hashanah.1.1"
)

// historicalHebrewRule represents a fixed date in the received Hebrew calendar.
// Ancient dates are projected into modern fixed-calendar years for study; they
// are not claims about exact Gregorian anniversaries in an observational era.
type historicalHebrewRule struct {
	CatalogID  string
	Name       string
	Month      int
	Day        int
	EndDay     int
	Adar       bool
	Category   string
	Origin     string
	Status     string
	Summary    string
	Note       string
	SourceName string
	SourceURL  string
}

// The complete 35-entry/period Aramaic core of Megillat Ta'anit. The scroll
// records days on which fasting was forbidden; entries marked in the text as
// stricter also barred eulogies. Rosh Hashanah 18b-19b treats this calendar as
// abrogated after the Second Temple, apart from Hanukkah and Purim.
var megillatTaanitRules = []historicalHebrewRule{
	{"mt-nisan-tamid", "Megillat Ta'anit · Public Tamid Restored", 1, 1, 8, false, "Historical no-fast period", "Megillat Ta'anit", "Discontinued", "An eight-day period marked the restoration of communal funding for the daily Temple offering.", "The later scholion connects this with a Pharisaic-Sadducean dispute; the identification is traditional rather than stated by the terse Aramaic entry.", "Megillat Ta'anit · Nisan", "https://www.sefaria.org/Megillat_Ta%27anit%2C_Nisan"},
	{"mt-nisan-festival-calendar", "Megillat Ta'anit · Festival Calendar Vindicated", 1, 8, 21, false, "Historical no-fast period", "Megillat Ta'anit", "Discontinued", "A period through Passover marked the accepted calculation of Shavuot and the Omer.", "The no-fast period overlaps biblical Passover. Its scholion frames a dispute with Boethusians over whether the Omer count always began on Sunday.", "Megillat Ta'anit · Nisan", "https://www.sefaria.org/Megillat_Ta%27anit%2C_Nisan"},
	{"mt-iyar-wall-dedication", "Megillat Ta'anit · Jerusalem Wall Dedication", 2, 7, 0, false, "Historical no-fast day", "Megillat Ta'anit", "Discontinued", "A dedication of Jerusalem's wall was kept as a no-fast and no-eulogy day.", "The terse source does not securely identify which rebuilding or dedication is meant.", "Megillat Ta'anit · Iyar", "https://www.sefaria.org/Megillat_Ta%27anit%2C_Iyar"},
	{"mt-iyar-pesach-sheni", "Megillat Ta'anit · Minor Passover", 2, 14, 0, false, "Historical no-fast day", "Megillat Ta'anit", "Temple rite discontinued; date still marked", "The Second Passover sacrifice was listed as a no-fast and no-eulogy day.", "The Torah's sacrifice is inactive without the Temple, while Pesach Sheni remains a minor modern observance.", "Megillat Ta'anit · Iyar", "https://www.sefaria.org/Megillat_Ta%27anit%2C_Iyar"},
	{"mt-iyar-akra", "Megillat Ta'anit · Akra Garrison Departed", 2, 23, 0, false, "Historical no-fast day", "Megillat Ta'anit", "Discontinued", "The departure of the foreign garrison from Jerusalem's Akra was commemorated.", "This Second Temple political anniversary is projected onto the fixed Hebrew calendar.", "Megillat Ta'anit · Iyar", "https://www.sefaria.org/Megillat_Ta%27anit%2C_Iyar"},
	{"mt-iyar-tribute", "Megillat Ta'anit · Tribute Removed", 2, 27, 0, false, "Historical no-fast day", "Megillat Ta'anit", "Discontinued", "The end of a tribute or crown-tax imposed on Judah and Jerusalem was commemorated.", "The exact authority and episode behind the terse notice are uncertain.", "Megillat Ta'anit · Iyar", "https://www.sefaria.org/Megillat_Ta%27anit%2C_Iyar"},
	{"mt-sivan-beth-shean", "Megillat Ta'anit · Beth Shean and Valley Victory", 3, 15, 16, false, "Historical no-fast period", "Megillat Ta'anit", "Discontinued", "Two days recalled the removal of hostile inhabitants from Beth Shean and the valley.", "The wording and historical reconstruction should be read in its ancient conflict setting, not as a present practice.", "Megillat Ta'anit · Sivan", "https://www.sefaria.org/Megillat_Ta%27anit%2C_Sivan"},
	{"mt-sivan-tower", "Megillat Ta'anit · Tower Captured", 3, 17, 0, false, "Historical no-fast day", "Megillat Ta'anit", "Discontinued", "The capture of a fortress or tower was commemorated.", "The Hebrew witness says 17 Sivan; an older English translation printed 14. The site's base date follows the Hebrew text.", "Megillat Ta'anit · Sivan", "https://www.sefaria.org/Megillat_Ta%27anit%2C_Sivan"},
	{"mt-sivan-tax-collectors", "Megillat Ta'anit · Tax Collectors Removed", 3, 25, 0, false, "Historical no-fast day", "Megillat Ta'anit", "Discontinued", "The removal of tax collectors from Judah and Jerusalem was commemorated.", "The exact political setting of this terse notice remains uncertain.", "Megillat Ta'anit · Sivan", "https://www.sefaria.org/Megillat_Ta%27anit%2C_Sivan"},
	{"mt-tammuz-decrees", "Megillat Ta'anit · Book of Decrees Abolished", 4, 14, 0, false, "Historical no-fast day", "Megillat Ta'anit", "Discontinued", "A legal victory called the removal of the Book of Decrees was commemorated.", "The Hebrew text reads 14 Tammuz; translation traditions contain conflicting numerals. The scholion relates it to a sectarian legal dispute.", "Megillat Ta'anit · Tammuz", "https://www.sefaria.org/Megillat_Ta%27anit%2C_Tammuz"},
	{"mt-av-wood", "Megillat Ta'anit · Xylophoria / Wood Festival", 5, 15, 0, false, "Historical no-fast day", "Megillat Ta'anit", "Temple rite discontinued; Tu B'Av remains", "The priests' wood-offering festival was a no-fast and no-eulogy day.", "This Temple rite overlaps Tu B'Av, which survives with other meanings.", "Megillat Ta'anit · Av", "https://www.sefaria.org/Megillat_Ta%27anit%2C_Av"},
	{"mt-av-law", "Megillat Ta'anit · Return to the Law", 5, 24, 0, false, "Historical no-fast day", "Megillat Ta'anit", "Discontinued", "The scroll tersely marks a day on which 'we returned to our law' or judgment.", "Its exact historical referent is uncertain; later explanations should not be treated as the wording of the original scroll.", "Megillat Ta'anit · Av", "https://www.sefaria.org/Megillat_Ta%27anit%2C_Av"},
	{"mt-elul-wall", "Megillat Ta'anit · Jerusalem Wall Dedicated", 6, 7, 0, false, "Historical no-fast day", "Megillat Ta'anit", "Discontinued", "A dedication of Jerusalem's wall was commemorated as a stricter no-fast day.", "The source does not securely identify which wall dedication is intended.", "Megillat Ta'anit · Elul", "https://www.sefaria.org/Megillat_Ta%27anit%2C_Elul"},
	{"mt-elul-garrison", "Megillat Ta'anit · Foreign Garrison Evacuated", 6, 17, 0, false, "Historical no-fast day", "Megillat Ta'anit", "Discontinued", "The evacuation of a Roman or foreign garrison from Judah and Jerusalem was commemorated.", "The scroll and later commentary preserve different historical labels.", "Megillat Ta'anit · Elul", "https://www.sefaria.org/Megillat_Ta%27anit%2C_Elul"},
	{"mt-elul-persecutors", "Megillat Ta'anit · Persecutors Defeated", 6, 22, 0, false, "Historical no-fast day", "Megillat Ta'anit", "Discontinued", "An ancient victory over people described by the source as persecutors or the wicked was commemorated.", "This is a historical notice from a conflict text, not a modern ritual prescription.", "Megillat Ta'anit · Elul", "https://www.sefaria.org/Megillat_Ta%27anit%2C_Elul"},
	{"mt-tishrei-divine-name", "Megillat Ta'anit · Divine Name Removed from Contracts", 7, 3, 0, false, "Historical no-fast day", "Megillat Ta'anit", "Discontinued; overlaps an active fast", "The removal of casual divine-name formulas from legal documents was commemorated.", "After Megillat Ta'anit was abrogated, the Fast of Gedaliah again defined this date.", "Megillat Ta'anit · Tishrei", "https://www.sefaria.org/Megillat_Ta%27anit%2C_Tishrei"},
	{"mt-cheshvan-soreg", "Megillat Ta'anit · Soreg Removed", 8, 23, 0, false, "Historical no-fast day", "Megillat Ta'anit", "Discontinued", "The removal or repair of the Temple court barrier called the soreg was commemorated.", "The terse entry's precise referent is obscure.", "Megillat Ta'anit · Cheshvan", "https://www.sefaria.org/Megillat_Ta%27anit%2C_Cheshvan"},
	{"mt-cheshvan-samaria", "Megillat Ta'anit · Samaria Wall Captured", 8, 25, 0, false, "Historical no-fast day", "Megillat Ta'anit", "Discontinued", "The capture of Samaria's wall was commemorated.", "This ancient military anniversary is retained for historical study only.", "Megillat Ta'anit · Cheshvan", "https://www.sefaria.org/Megillat_Ta%27anit%2C_Cheshvan"},
	{"mt-cheshvan-flour", "Megillat Ta'anit · Fine-Flour Offering Restored", 8, 27, 0, false, "Historical no-fast day", "Megillat Ta'anit", "Discontinued", "The restoration of the fine-flour offering on the altar was commemorated.", "The scholion frames the day as a sectarian sacrificial-law victory.", "Megillat Ta'anit · Cheshvan", "https://www.sefaria.org/Megillat_Ta%27anit%2C_Cheshvan"},
	{"mt-kislev-ensigns", "Megillat Ta'anit · Ensigns Removed", 9, 3, 0, false, "Historical no-fast day", "Megillat Ta'anit", "Discontinued", "The removal of ensigns or idolatrous installations from the Temple court was commemorated.", "The underlying episode is known only through terse and later layered accounts.", "Megillat Ta'anit · Kislev", "https://www.sefaria.org/Megillat_Ta%27anit%2C_Kislev"},
	{"mt-kislev-seven", "Megillat Ta'anit · Kislev Festival", 9, 7, 0, false, "Historical no-fast day", "Megillat Ta'anit", "Discontinued", "The core scroll calls this simply a festival.", "A later scholion associates it with Herod's death; that explanation is not present in the Aramaic core.", "Megillat Ta'anit · Kislev", "https://www.sefaria.org/Megillat_Ta%27anit%2C_Kislev"},
	{"mt-kislev-gerizim", "Megillat Ta'anit · Mount Gerizim Day", 9, 21, 0, false, "Historical no-fast day", "Megillat Ta'anit", "Discontinued", "A victory associated with Mount Gerizim was a no-fast and no-eulogy day.", "Yoma 69a transmits 25 Tevet for a related Mount Gerizim story; the traditions should not be silently merged.", "Megillat Ta'anit · Kislev", "https://www.sefaria.org/Megillat_Ta%27anit%2C_Kislev"},
	{"mt-tevet-sanhedrin", "Megillat Ta'anit · Sanhedrin Restored", 10, 28, 0, false, "Historical no-fast day", "Megillat Ta'anit", "Discontinued", "The assembly or Sanhedrin sitting again in judgment was commemorated.", "The precise institutional event remains disputed.", "Megillat Ta'anit · Tevet", "https://www.sefaria.org/Megillat_Ta%27anit%2C_Tevet"},
	{"mt-shevat-two", "Megillat Ta'anit · Shevat Festival", 11, 2, 0, false, "Historical no-fast day", "Megillat Ta'anit", "Discontinued", "The core text designates a festival without naming its cause.", "The Hebrew witness reads 2 Shevat; translations also preserve 7. A later scholion associates it with Alexander Jannaeus's death.", "Megillat Ta'anit · Shevat", "https://www.sefaria.org/Megillat_Ta%27anit%2C_Shevat"},
	{"mt-shevat-temple-decree", "Megillat Ta'anit · Temple Decree Halted", 11, 22, 0, false, "Historical no-fast day", "Megillat Ta'anit", "Discontinued", "The halting of an enemy order concerning the Temple was commemorated.", "Later interpretation commonly links the entry to Caligula's statue decree.", "Megillat Ta'anit · Shevat", "https://www.sefaria.org/Megillat_Ta%27anit%2C_Shevat"},
	{"mt-shevat-antiochus", "Megillat Ta'anit · Antiochus Departed Jerusalem", 11, 28, 0, false, "Historical no-fast day", "Megillat Ta'anit", "Discontinued", "The departure of a ruler named Antiochus from Jerusalem was commemorated.", "Which Antiochus and which episode are uncertain.", "Megillat Ta'anit · Shevat", "https://www.sefaria.org/Megillat_Ta%27anit%2C_Shevat"},
	{"mt-adar-rain-eight", "Megillat Ta'anit · Rain Prayer Answered", 12, 8, 9, true, "Historical no-fast period", "Megillat Ta'anit", "Discontinued", "Two days recalled a communal supplication for rain that was answered.", "This anniversary is distinct from contingent rain fasts declared according to drought conditions.", "Megillat Ta'anit · Adar", "https://www.sefaria.org/Megillat_Ta%27anit%2C_Adar"},
	{"mt-adar-turyanus", "Megillat Ta'anit · Day of Turyanus", 12, 12, 0, true, "Historical no-fast day", "Megillat Ta'anit", "Discontinued", "A deliverance called the Day of Turyanus or Trajan was commemorated.", "The identity and historical event are disputed.", "Megillat Ta'anit · Adar", "https://www.sefaria.org/Megillat_Ta%27anit%2C_Adar"},
	{"mt-adar-nicanor", "Megillat Ta'anit · Nicanor Day", 12, 13, 0, true, "Historical no-fast day", "Megillat Ta'anit", "Discontinued; overlaps Fast of Esther", "The defeat of the Seleucid commander Nicanor was commemorated.", "The later Fast of Esther now occupies the same nominal date.", "Megillat Ta'anit · Adar", "https://www.sefaria.org/Megillat_Ta%27anit%2C_Adar"},
	{"mt-adar-purim", "Megillat Ta'anit · Purim Days", 12, 14, 15, true, "No-fast festival", "Megillat Ta'anit", "Still observed", "The two Purim dates were included among days without fasting or eulogies.", "Purim remains active: the 14th is standard in unwalled places and the 15th in qualifying ancient walled cities.", "Megillat Ta'anit · Adar", "https://www.sefaria.org/Megillat_Ta%27anit%2C_Adar"},
	{"mt-adar-wall", "Megillat Ta'anit · Jerusalem Wall Building Began", 12, 16, 0, true, "Historical no-fast day", "Megillat Ta'anit", "Discontinued", "The beginning of work on Jerusalem's wall was commemorated.", "The source preserves the anniversary without enough detail for a certain reconstruction.", "Megillat Ta'anit · Adar", "https://www.sefaria.org/Megillat_Ta%27anit%2C_Adar"},
	{"mt-adar-chalcis", "Megillat Ta'anit · Deliverance at Chalcis", 12, 17, 0, true, "Historical no-fast day", "Megillat Ta'anit", "Discontinued", "A deliverance of Jewish refugees or sages at Chalcis or Beth Zabdin was commemorated.", "The place names and historical reconstruction are textually uncertain.", "Megillat Ta'anit · Adar", "https://www.sefaria.org/Megillat_Ta%27anit%2C_Adar"},
	{"mt-adar-rain-twenty", "Megillat Ta'anit · Communal Rain Fast Answered", 12, 20, 0, true, "Historical no-fast day", "Megillat Ta'anit", "Discontinued", "A communal fast and prayer for rain that was answered became a no-fast anniversary.", "It records a specific deliverance, not an annually scheduled drought fast.", "Megillat Ta'anit · Adar", "https://www.sefaria.org/Megillat_Ta%27anit%2C_Adar"},
	{"mt-adar-torah-study", "Megillat Ta'anit · Torah Study Restored", 12, 28, 0, true, "Historical no-fast day", "Megillat Ta'anit", "Discontinued", "Good news that Jews would no longer be barred from Torah study was commemorated.", "Bavli Ta'anit 18a preserves a later narrative for this terse anniversary.", "Megillat Ta'anit · Adar", "https://www.sefaria.org/Megillat_Ta%27anit%2C_Adar"},
}

var woodOfferingRules = []historicalHebrewRule{
	{"wood-arah", "Wood Offering · Arah ben Judah", 1, 1, 0, false, "Temple family festival", "Mishnah Ta'anit 4:5", "Private rite discontinued", "A designated family brought wood for the Temple altar and kept the day as its own festival.", "This was a family-specific Temple service day, not a national holiday.", "Mishnah Ta'anit 4:5", sefariaMishnahTaanit},
	{"wood-david", "Wood Offering · David ben Judah", 4, 20, 0, false, "Temple family festival", "Mishnah Ta'anit 4:5", "Private rite discontinued", "The family of David ben Judah brought its wood offering.", "This was a family-specific Temple service day, not a national holiday.", "Mishnah Ta'anit 4:5", sefariaMishnahTaanit},
	{"wood-parosh-av", "Wood Offering · Parosh ben Judah", 5, 5, 0, false, "Temple family festival", "Mishnah Ta'anit 4:5", "Private rite discontinued", "The family of Parosh ben Judah brought its wood offering.", "This was a family-specific Temple service day, not a national holiday.", "Mishnah Ta'anit 4:5", sefariaMishnahTaanit},
	{"wood-jonadab", "Wood Offering · Jonadab ben Rechab", 5, 7, 0, false, "Temple family festival", "Mishnah Ta'anit 4:5", "Private rite discontinued", "The family of Jonadab ben Rechab brought its wood offering.", "This was a family-specific Temple service day, not a national holiday.", "Mishnah Ta'anit 4:5", sefariaMishnahTaanit},
	{"wood-senaah", "Wood Offering · Senaah ben Benjamin", 5, 10, 0, false, "Temple family festival", "Mishnah Ta'anit 4:5", "Private rite discontinued", "The family of Senaah ben Benjamin brought its wood offering.", "Ta'anit 12a remembers limited post-Temple family observance of this day.", "Mishnah Ta'anit 4:5", sefariaMishnahTaanit},
	{"wood-zattu", "Wood Offering · Zattu and Uncertain Lineages", 5, 15, 0, false, "Temple family festival", "Mishnah Ta'anit 4:5", "Private rite discontinued", "Zattu, priests, Levites, Israelites of uncertain lineage, and named occupational groups brought wood.", "This major wood-offering date overlaps Tu B'Av and Megillat Ta'anit's Xylophoria.", "Mishnah Ta'anit 4:5", sefariaMishnahTaanit},
	{"wood-pahat-moab", "Wood Offering · Pahat-Moab ben Judah", 5, 20, 0, false, "Temple family festival", "Mishnah Ta'anit 4:5", "Private rite discontinued", "The family of Pahat-Moab ben Judah brought its wood offering.", "This was a family-specific Temple service day, not a national holiday.", "Mishnah Ta'anit 4:5", sefariaMishnahTaanit},
	{"wood-adin", "Wood Offering · Adin ben Judah", 6, 20, 0, false, "Temple family festival", "Mishnah Ta'anit 4:5", "Private rite discontinued", "The family of Adin ben Judah brought its wood offering.", "This was a family-specific Temple service day, not a national holiday.", "Mishnah Ta'anit 4:5", sefariaMishnahTaanit},
	{"wood-parosh-tevet", "Wood Offering · Parosh ben Judah (Tevet)", 10, 1, 0, false, "Temple family festival", "Mishnah Ta'anit 4:5", "Private rite discontinued", "The family of Parosh ben Judah brought a second annual wood offering.", "The date can overlap Hanukkah and Rosh Chodesh.", "Mishnah Ta'anit 4:5", sefariaMishnahTaanit},
}

var rabbinicCalendarRules = []historicalHebrewRule{
	{"mishnah-new-year-kings", "New Year for Kings and Festivals", 1, 1, 0, false, "Mishnaic calendar marker", "Mishnah Rosh Hashanah 1:1", "Administrative rule discontinued", "Nisan 1 began regnal-year counting and ordered the pilgrimage festivals.", "The Mishnah calls this a new year for legal calculation, not a separate work-free festival.", "Mishnah Rosh Hashanah 1:1", sefariaMishnahRH},
	{"mishnah-new-year-animal-tithe", "New Year for Animal Tithe", 6, 1, 0, false, "Mishnaic calendar marker", "Mishnah Rosh Hashanah 1:1", "Temple/agricultural rule discontinued", "Elul 1 separated animal-tithe years according to the first opinion.", "Rabbi Elazar and Rabbi Shimon place this cutoff on Tishrei 1 instead.", "Mishnah Rosh Hashanah 1:1", sefariaMishnahRH},
	{"mishnah-new-year-years", "New Year for Years, Release, and Planting", 7, 1, 0, false, "Mishnaic calendar marker", "Mishnah Rosh Hashanah 1:1", "Underlying date still observed", "Tishrei 1 begins years and the legal cycles for sabbatical years, Jubilee, planting, and vegetables.", "This legal-agricultural layer shares the date of Rosh Hashanah.", "Mishnah Rosh Hashanah 1:1", sefariaMishnahRH},
	{"mishnah-new-year-trees-shammai", "New Year for Trees · Beit Shammai", 11, 1, 0, false, "Mishnaic minority date", "Mishnah Rosh Hashanah 1:1", "Not normative", "Beit Shammai placed the tree-tithe new year on Shevat 1.", "The accepted ruling follows Beit Hillel on Shevat 15; this entry preserves the Mishnah's dissenting date.", "Mishnah Rosh Hashanah 1:1", sefariaMishnahRH},
	{"mishnah-new-year-trees-hillel", "New Year for Trees · Beit Hillel", 11, 15, 0, false, "Mishnaic calendar marker", "Mishnah Rosh Hashanah 1:1", "Date still observed as Tu BiShvat", "Beit Hillel placed the tree-tithe new year on Shevat 15.", "The Mishnah describes an agricultural cutoff; fruit seders and ecological themes developed much later.", "Mishnah Rosh Hashanah 1:1", sefariaMishnahRH},
	{"mishnah-early-megillah", "Early Megillah Reading Window", 12, 11, 13, true, "Historical calendar procedure", "Mishnah Megillah 1:1-2", "Discontinued as routine practice", "Villages could advance the Megillah reading to local court-and-market assembly days.", "These were conditional reading dates, not three separate Purim festivals.", "Mishnah Megillah 1:1-2", "https://www.sefaria.org/Mishnah_Megillah.1.1-2"},
}

func historicalJewishObservances(date time.Time, h hebrewDate) []Observance {
	events := make([]Observance, 0, 4)
	for _, rule := range append(append([]historicalHebrewRule{}, megillatTaanitRules...), append(woodOfferingRules, rabbinicCalendarRules...)...) {
		month := rule.Month
		if rule.Adar {
			month = 12
			if hebrewLeap(h.Year) {
				month = 13
			}
		}
		end := rule.EndDay
		if end == 0 {
			end = rule.Day
		}
		if h.Month != month || h.Day < rule.Day || h.Day > end {
			continue
		}
		event := baseObservance(rule.Name, Judaism, []string{"Historical Jewish calendar study"}, rule.Category, rule.Summary, "This date is retained so discontinued and minority calendars remain visible beside living practice.", nil, nil, rule.SourceName, rule.SourceURL)
		event.CatalogID = rule.CatalogID
		event.Origin = rule.Origin
		event.ObservanceStatus = rule.Status
		event.HistoricalNote = rule.Note
		event.Historical = true
		event.DateCertainty = "Traditional Hebrew date projected into the modern fixed calendar"
		event.StartsAtSunset = true
		event.DateNote = "Historical projection: ancient observational dates cannot always be assigned an exact proleptic Gregorian anniversary. " + rule.Note
		start := date.AddDate(0, 0, -(h.Day - rule.Day))
		if end > rule.Day {
			event = spanOccurrence(event, date, start, end-rule.Day+1)
		} else {
			event = singleOccurrence(event, date)
		}
		event.ID = rule.CatalogID + "-" + start.Format("2006-01-02")
		events = append(events, event)
	}

	// Hanukkah crosses the Kislev/Tevet boundary, so its Megillat Ta'anit
	// catalog occurrence cannot be expressed as a same-month fixed range.
	if index := hanukkahDay(h); index > 0 {
		start := date.AddDate(0, 0, -(index - 1))
		event := baseObservance("Megillat Ta'anit · Hanukkah", Judaism, []string{"Jewish communities"}, "No-fast festival", "The scroll records Hanukkah as an eight-day dedication festival without fasting or eulogies.", "Unlike most entries in Megillat Ta'anit, Hanukkah remains a living observance.", nil, []string{"Bavli Shabbat 21b"}, "Megillat Ta'anit · Kislev", "https://www.sefaria.org/Megillat_Ta%27anit%2C_Kislev")
		event.CatalogID, event.Origin, event.ObservanceStatus = "mt-kislev-hanukkah", "Megillat Ta'anit and Bavli Shabbat 21b", "Still observed"
		event.HistoricalNote = "One of the two festival families explicitly retained after the Megillat Ta'anit calendar was abrogated."
		event.Historical = true
		event.DateCertainty = "Received fixed-calendar date"
		event.StartsAtSunset = true
		event = spanOccurrence(event, date, start, 8)
		event.ID = event.CatalogID + "-" + start.Format("2006-01-02")
		events = append(events, event)
	}
	return events
}

func torahCalendarLayers(date time.Time, h hebrewDate) []Observance {
	var events []Observance
	add := func(catalogID string, event Observance, status, note string) {
		event.CatalogID = catalogID
		event.Origin = "Torah"
		event.ObservanceStatus = status
		event.HistoricalNote = note
		event.DateCertainty = "Received fixed-calendar date; the Torah uses numbered months"
		event.StartsAtSunset = true
		event = singleOccurrence(event, date)
		event.ID = catalogID + "-" + date.Format("2006-01-02")
		events = append(events, event)
	}
	if h.Month == 1 && h.Day == 14 {
		event := baseObservance("Korban Pesach · Passover Offering", Judaism, []string{"Temple-era Israel; remembered across Jewish communities"}, "Torah-appointed rite", "At twilight on Nisan 14, the Torah appoints the Passover sacrifice, distinct from the seven-day Feast of Unleavened Bread.", "The Seder preserves the offering's memory, while the sacrifice itself depends on the Temple.", []string{"Temple sacrifice historically", "Remembered in the Haggadah and Seder plate"}, []string{"Exodus 12:6–14", "Leviticus 23:5", "Numbers 9:1–5"}, "Torah · Leviticus 23:5", "https://www.sefaria.org/Leviticus.23.5")
		add("torah-korban-pesach", event, "Temple rite inactive; memory observed", "Nisan 10 lamb selection belonged to the one-time Passover in Egypt and is not projected as an annual holiday.")
	}
	if h.Month == 1 && h.Day == 16 {
		event := baseObservance("Omer Offering · Counting Begins", Judaism, []string{"Jewish communities; Temple offering inactive"}, "Torah-appointed count", "The barley sheaf was elevated and a count of seven complete weeks began toward Shavuot.", "Counting remains active each evening; the agricultural and sacrificial rites await the Temple.", []string{"Count the Omer after nightfall", "Temple barley offering historically"}, []string{"Leviticus 23:9–16", "Deuteronomy 16:9"}, "Torah · Leviticus 23:9–16", "https://www.sefaria.org/Leviticus.23.9-16")
		add("torah-omer-opening", event, "Count still observed; Temple rite inactive", "The Torah says 'the morrow after the Sabbath' without a month-day; Nisan 16 is the received rabbinic interpretation.")
	}
	if h.Month == 3 && h.Day == 6 {
		event := baseObservance("Bikkurim Season Opens", Judaism, []string{"Temple-era agricultural Israel"}, "Torah agricultural rite", "Individual first fruits could be brought with their historical declaration from Shavuot through Sukkot in the Mishnaic Temple calendar.", "The declaration joined harvest gratitude to the memory of wandering, oppression, liberation, and land.", []string{"First-fruit procession and presentation historically", "Recitation of the bikkurim declaration"}, []string{"Deuteronomy 26:1–11", "Mishnah Bikkurim 1:6"}, "Torah · Deuteronomy 26", "https://www.sefaria.org/Deuteronomy.26.1-11")
		add("torah-bikkurim-season", event, "Temple rite inactive", "The Torah gives an agricultural season, not a fixed opening date; the Shavuot opening follows Mishnah Bikkurim 1:6.")
	}
	// In the received cycle, Hebrew years divisible by seven are Shemitah.
	// Hakhel followed at Sukkot in the eighth year; Mishnah Sotah places it
	// at the beginning of Chol HaMoed rather than giving a Torah month-day.
	if h.Month == 7 && h.Day == 16 && h.Year%7 == 1 {
		event := baseObservance("Hakhel · Septennial Torah Assembly", Judaism, []string{"Temple-era Israel; modern commemorative gatherings"}, "Torah septennial assembly", "After the release year, the whole people assembled at Sukkot to hear Torah read publicly.", "Hakhel centers covenantal learning across age, gender, and social position.", []string{"Public Torah reading historically", "Commemorative assemblies in some modern communities"}, []string{"Deuteronomy 31:10–13", "Mishnah Sotah 7:8"}, "Torah · Deuteronomy 31", "https://www.sefaria.org/Deuteronomy.31.10-13")
		add("torah-hakhel", event, "Temple institution inactive; sometimes commemorated", "The Torah gives a relative Sukkot schedule; 16 Tishrei follows the Mishnaic placement and is not an explicit Torah date.")
	}
	return events
}

func leapYearMinorPurims(date time.Time, h hebrewDate) []Observance {
	if !hebrewLeap(h.Year) || h.Month != 12 || (h.Day != 14 && h.Day != 15) {
		return nil
	}
	name := "Purim Katan"
	catalogID := "purim-katan"
	if h.Day == 15 {
		name = "Shushan Purim Katan"
		catalogID = "shushan-purim-katan"
	}
	event := baseObservance(name, Judaism, []string{"Jewish communities in Hebrew leap years"}, "Minor rabbinic observance", "Adar I preserves a minor echo of Purim before the full festival in Adar II.", "The Mishnah distinguishes the two Adars chiefly by the Megillah and gifts to the poor; later law adds the day's customary festive character.", []string{"Omitting penitential prayer in traditional liturgy", "Avoiding fasting and eulogies", "Optional festive meal"}, []string{"Mishnah Megillah 1:4"}, "Mishnah Megillah 1:4", "https://www.sefaria.org/Mishnah_Megillah.1.4")
	event.CatalogID, event.Origin, event.ObservanceStatus = catalogID, "Mishnah and later halakhah", "Still observed as a minor day"
	event.DateCertainty = "Received fixed-calendar date"
	event.StartsAtSunset = true
	event = singleOccurrence(event, date)
	event.ID = catalogID + "-" + date.Format("2006-01-02")
	return []Observance{event}
}

func purimMeshulash(date time.Time, h hebrewDate) []Observance {
	if !purimMonth(h) || h.Day < 14 || h.Day > 16 {
		return nil
	}
	fifteenth := date.AddDate(0, 0, 15-h.Day)
	if fifteenth.Weekday() != time.Saturday {
		return nil
	}
	start := fifteenth.AddDate(0, 0, -1)
	event := baseObservance("Purim Meshulash · Walled Cities", Judaism, []string{"Jerusalem and qualifying walled cities"}, "Conditional rabbinic observance", "When Shushan Purim falls on Shabbat, its obligations are distributed across Friday, Shabbat, and Sunday.", "The three-day pattern keeps the Megillah, liturgy, gifts, and feast within their legal times without reading the scroll on Shabbat.", []string{"Friday: Megillah and gifts to people in need", "Shabbat: Purim liturgy and Torah reading", "Sunday: festive meal and commonly gifts of food"}, []string{"Jerusalem Talmud Megillah 1:4"}, "Jerusalem Talmud Megillah 1:4", "https://www.sefaria.org/Jerusalem_Talmud_Megillah.1.4")
	event.CatalogID, event.Origin, event.ObservanceStatus = "purim-meshulash", "Jerusalem Talmud and later halakhah", "Still observed when the calendar condition occurs"
	event.DateCertainty = "Generated only when 15 Adar falls on Shabbat"
	event.StartsAtSunset = true
	event = spanOccurrence(event, date, start, 3)
	event.ID = event.CatalogID + "-" + start.Format("2006-01-02")
	return []Observance{event}
}

func isruChagObservances(date time.Time, h hebrewDate) []Observance {
	type isruDate struct {
		month, day int
		region     string
		catalogID  string
	}
	dates := []isruDate{
		{1, 22, "Israel", "isru-chag-pesach-israel"}, {1, 23, "Diaspora", "isru-chag-pesach-diaspora"},
		{3, 7, "Israel", "isru-chag-shavuot-israel"}, {3, 8, "Diaspora", "isru-chag-shavuot-diaspora"},
		{7, 23, "Israel", "isru-chag-sukkot-israel"}, {7, 24, "Diaspora", "isru-chag-sukkot-diaspora"},
	}
	var events []Observance
	for _, rule := range dates {
		if h.Month != rule.month || h.Day != rule.day {
			continue
		}
		event := baseObservance("Isru Chag · "+rule.region, Judaism, []string{rule.region + " and communities following its festival length"}, "Post-festival custom", "The day after a pilgrimage festival retains a small measure of festivity.", "A Talmudic interpretation links the festival offering with bonds; later custom gives the following day additional food and omits penitential prayer.", []string{"Additional food or drink", "Omitting Tachanun in traditional liturgy", "Avoiding voluntary fasts in many communities"}, []string{"Bavli Sukkah 45b"}, "Bavli Sukkah 45b", "https://www.sefaria.org/Sukkah.45b")
		event.CatalogID = rule.catalogID
		event.Origin, event.ObservanceStatus = "Bavli Sukkah 45b and later custom", "Still observed as a minor custom"
		event.DateCertainty = "Region-specific day after the received festival length"
		event.StartsAtSunset = true
		event = singleOccurrence(event, date)
		event.ID = event.CatalogID + "-" + date.Format("2006-01-02")
		events = append(events, event)
	}
	return events
}

func templeEraRabbinicLayers(date time.Time, h hebrewDate) []Observance {
	var events []Observance
	addSpan := func(catalogID string, event Observance, start time.Time, duration int) {
		event.CatalogID = catalogID
		event.Origin = event.SourceName
		event.ObservanceStatus = "Temple practice discontinued"
		event.Historical = true
		event.DateCertainty = "Traditional Hebrew date projected into the modern fixed calendar"
		event.StartsAtSunset = true
		event = spanOccurrence(event, date, start, duration)
		event.ID = catalogID + "-" + start.Format("2006-01-02")
		events = append(events, event)
	}
	if h.Month == 7 && h.Day >= 16 && h.Day <= 21 {
		start := date.AddDate(0, 0, -(h.Day - 16))
		event := baseObservance("Simchat Beit HaSho'eva · Water-Drawing Rejoicing", Judaism, []string{"Temple-era Jerusalem"}, "Temple festival ceremony", "Extraordinary public rejoicing accompanied the water-drawing ceremony on the intermediate nights of Sukkot.", "The Mishnah remembers torchlight, music, dancing, and worship in the Temple courts as a summit of festival joy.", []string{"Water-drawing procession historically", "Music, torch dancing, and all-night celebration"}, []string{"Mishnah Sukkah 5:1–4"}, "Mishnah Sukkah 5", "https://www.sefaria.org/Mishnah_Sukkah.5")
		event.HistoricalNote = "A ceremony within Sukkot rather than an additional work-restricted holiday."
		addSpan("temple-water-drawing", event, start, 6)
	}
	if h.Month == 3 && h.Day >= 6 && h.Day <= 12 {
		start := date.AddDate(0, 0, -(h.Day - 6))
		event := baseObservance("Shavuot Tashlumin · Offering Make-up Days", Judaism, []string{"Temple-era pilgrims"}, "Temple offering period", "Six days after Shavuot remained available for pilgrims to complete festival offerings.", "The make-up window extended sacrificial responsibility without turning the following days into work-restricted festivals.", []string{"Pilgrimage offerings historically"}, []string{"Bavli Chagigah 17a"}, "Bavli Chagigah 17a", "https://www.sefaria.org/Chagigah.17a")
		event.HistoricalNote = "These were valid offering days, not seven additional festival days."
		addSpan("shavuot-tashlumin", event, start, 7)
	}
	return events
}

func fastOfFirstborn(date time.Time, h hebrewDate) []Observance {
	if h.Month != 1 || h.Day != 14 {
		return nil
	}
	event := baseObservance("Fast of the Firstborn · Ta'anit Bechorot", Judaism, []string{"Firstborn Jews in communities that retain the custom"}, "Limited rabbinic fast", "A dawn-to-Passover fast recalls the sparing of Israelite firstborn in Egypt.", "Participation is limited and a siyum celebration commonly ends or replaces the fast.", []string{"Daytime fasting by firstborn participants", "Siyum and celebratory meal in many communities"}, []string{"Jerusalem Talmud Pesachim 10:1", "Tractate Soferim 21:3"}, "Jerusalem Talmud Pesachim 10:1", "https://www.sefaria.org/Jerusalem_Talmud_Pesachim.10.1")
	event.CatalogID, event.Origin, event.ObservanceStatus = "fast-firstborn", "Rabbinic/custom; explicit formulation is post-Mishnaic", "Still observed in some communities"
	event.DateNote = "A limited dawn fast, not a sunset-starting Torah festival; its Talmudic basis is ambiguous. Calendar shifts can alter practical scheduling."
	event = singleOccurrence(event, date)
	event.ID = event.CatalogID + "-" + date.Format("2006-01-02")
	return []Observance{event}
}

func sourceLayerSummary() string {
	return "Torah-appointed times and the full 49-day Omer count; explicit Shemitah and unprojected Jubilee records; fixed, recurring, conditional, and institutional Mishnah/Talmud calendar rules; all 35 original entries or periods of Megillat Ta'anit; all 26 dates in the separately labeled printed Ma'amar Aharon appendix witness; and all nine Temple-family wood-offering dates"
}
