package calendar

import (
	"fmt"
	"sort"
	"time"
)

func BuildMonth(year int, month time.Month, requested ...Location) MonthResponse {
	location := requestedLocation(requested)
	first := dateAt(year, month, 1)
	last := first.AddDate(0, 1, -1)
	previousDate := first.AddDate(0, -1, 0)
	nextDate := first.AddDate(0, 1, 0)
	var previous, next *MonthReference
	if previousDate.Year() >= 1900 {
		ref := gregorianMonthReference(previousDate.Year(), previousDate.Month())
		previous = &ref
	}
	if nextDate.Year() <= 2100 {
		ref := gregorianMonthReference(nextDate.Year(), nextDate.Month())
		next = &ref
	}
	return buildMonthResponse(
		GregorianCalendar,
		year,
		int(month),
		first.Format("January 2006"),
		first,
		last.Day(),
		previous,
		next,
		location,
	)
}

func BuildHebrewMonth(year, month int, requested ...Location) (MonthResponse, error) {
	if year < 1 {
		return MonthResponse{}, fmt.Errorf("Hebrew year must be positive")
	}
	if month < 1 || month > hebrewMonthsInYear(year) {
		return MonthResponse{}, fmt.Errorf("month must be between 1 and %d for Hebrew year %d", hebrewMonthsInYear(year), year)
	}

	if !hebrewMonthWithinSupportedCivilRange(year, month) {
		return MonthResponse{}, fmt.Errorf("Hebrew month must fall within Gregorian years 1900 through 2100")
	}
	first := jdToGregorian(hebrewToJD(year, month, 1))
	dayCount := hebrewMonthDays(year, month)
	previous := supportedHebrewMonthReference(previousHebrewMonthReference(year, month))
	next := supportedHebrewMonthReference(nextHebrewMonthReference(year, month))

	location := requestedLocation(requested)
	return buildMonthResponse(
		HebrewCalendar,
		year,
		month,
		fmt.Sprintf("%s %d", hebrewMonthName(year, month), year),
		first,
		dayCount,
		previous,
		next,
		location,
	), nil
}

func hebrewMonthWithinSupportedCivilRange(year, month int) bool {
	if year < 1 || month < 1 || month > hebrewMonthsInYear(year) {
		return false
	}
	first := jdToGregorian(hebrewToJD(year, month, 1))
	last := first.AddDate(0, 0, hebrewMonthDays(year, month)-1)
	return first.Year() >= 1900 && last.Year() <= 2100
}

func supportedHebrewMonthReference(ref MonthReference) *MonthReference {
	if !hebrewMonthWithinSupportedCivilRange(ref.Year, ref.Month) {
		return nil
	}
	return &ref
}

func requestedLocation(requested []Location) Location {
	if len(requested) > 0 {
		return normalizeLocation(requested[0])
	}
	return DefaultLocation
}

func buildMonthResponse(system CalendarSystem, year, month int, label string, first time.Time, dayCount int, previous, next *MonthReference, location Location) MonthResponse {
	days, observanceCount := buildDays(first, dayCount, location)
	last := first.AddDate(0, 0, dayCount-1)
	return MonthResponse{
		CalendarSystem:  system,
		Year:            year,
		Month:           month,
		Label:           label,
		StartDate:       first.Format("2006-01-02"),
		EndDate:         last.Format("2006-01-02"),
		Previous:        previous,
		Next:            next,
		FirstWeekday:    int(first.Weekday()),
		Days:            days,
		ObservanceCount: observanceCount,
		Coverage:        monthCoverage(),
		Location:        location,
		GeneratedAt:     time.Now().UTC(),
	}
}

func buildDays(first time.Time, dayCount int, location Location) ([]Day, int) {
	loc, _ := time.LoadLocation(location.Timezone)
	today := time.Now().In(loc)
	days := make([]Day, 0, dayCount)
	unique := map[string]bool{}
	for offset := 0; offset < dayCount; offset++ {
		date := first.AddDate(0, 0, offset)
		sacred := sacredDates(date)
		events := observancesForDate(date)
		for _, event := range events {
			unique[event.ID] = true
		}
		days = append(days, Day{
			Date:        date.Format("2006-01-02"),
			Day:         date.Day(),
			Weekday:     date.Weekday().String(),
			IsToday:     date.Year() == today.Year() && date.YearDay() == today.YearDay(),
			SacredDates: sacred,
			Moon:        moonForDate(date),
			Kabbalah:    kabbalahForMonth(sacred.HebrewMonth),
			Astrology:   astrologyForDate(date),
			Numerology:  numerologyForDate(date),
			Prayers:     prayerSchedules(date, location),
			Reading:     catholicReadingForDate(date, events),
			Observances: events,
		})
	}
	return days, len(unique)
}

func monthCoverage() []string {
	return []string{
		"Christian: Catholic, Orthodox, Anglican, Lutheran, Reformed, and other Protestant observances",
		"Jewish: " + sourceLayerSummary() + ", alongside living fasts, later customs, modern days, Rosh Chodesh, and weekly Shabbat",
		"Islamic: Sunni and Shia sacred days, Ramadan, Hajj season, both Eids, and weekly Jumu'ah",
		"Pagan and polytheist: a documented modern Neo-Pagan eightfold convention, kept distinct from historical Norse and Old English records",
		"Ancient world day cells: nominal Roman dates and an explicitly reconstructed Attic lunar study proxy; the Observance Atlas separately adds native-date-only Egyptian, Mesopotamian, Ugaritic, regional/Panhellenic Greek, Norse, Old English, and disputed Attic records",
	}
}

func gregorianMonthReference(year int, month time.Month) MonthReference {
	return MonthReference{
		Year:  year,
		Month: int(month),
		Label: dateAt(year, month, 1).Format("January 2006"),
	}
}

func previousHebrewMonthReference(year, month int) MonthReference {
	if month == 7 {
		year--
		month = 6
	} else if month == 1 {
		month = hebrewMonthsInYear(year)
	} else {
		month--
	}
	return hebrewMonthReference(year, month)
}

func nextHebrewMonthReference(year, month int) MonthReference {
	if month == 6 {
		year++
		month = 7
	} else if month == hebrewMonthsInYear(year) {
		month = 1
	} else {
		month++
	}
	return hebrewMonthReference(year, month)
}

func hebrewMonthReference(year, month int) MonthReference {
	return MonthReference{
		Year:  year,
		Month: month,
		Label: fmt.Sprintf("%s %d", hebrewMonthName(year, month), year),
	}
}

func BuildObservanceIndex(year int) ObservanceIndex {
	start := dateAt(year, time.January, 1)
	end := dateAt(year, time.December, 31)
	seen := map[string]bool{}
	var observances []Observance
	for date := start; !date.After(end); date = date.AddDate(0, 0, 1) {
		for _, event := range observancesForDate(date) {
			if seen[event.ID] {
				continue
			}
			seen[event.ID] = true
			observances = append(observances, event)
		}
	}
	for _, event := range catalogOnlyJewishObservances() {
		if seen[event.ID] {
			continue
		}
		seen[event.ID] = true
		observances = append(observances, event)
	}
	for _, event := range catalogOnlyAncientObservances() {
		if seen[event.ID] {
			continue
		}
		seen[event.ID] = true
		observances = append(observances, event)
	}
	for _, event := range catalogOnlyAtticObservances() {
		if seen[event.ID] {
			continue
		}
		seen[event.ID] = true
		observances = append(observances, event)
	}
	sort.Slice(observances, func(i, j int) bool {
		if observances[i].CatalogOnly != observances[j].CatalogOnly {
			return !observances[i].CatalogOnly
		}
		if observances[i].Date == observances[j].Date {
			if observances[i].CalendarCorpus == observances[j].CalendarCorpus {
				return observances[i].Name < observances[j].Name
			}
			return observances[i].CalendarCorpus < observances[j].CalendarCorpus
		}
		return observances[i].Date < observances[j].Date
	})
	counts := map[Tradition]int{Christianity: 0, Judaism: 0, Islam: 0, Polytheist: 0, AncientWorld: 0}
	for _, event := range observances {
		counts[event.Tradition]++
	}
	return ObservanceIndex{
		Year:        year,
		Observances: observances,
		Counts:      counts,
		Coverage: map[Tradition][]string{
			Christianity: {"Western and Orthodox Paschal cycles", "Principal Catholic solemnities and feasts", "Selected Anglican, Lutheran, Reformed, and Protestant commemorations", "Julian-calendar Nativity and Theophany"},
			Judaism:      {"Weekly Shabbat, monthly Rosh Chodesh, and the monthly lunar-blessing timing discussion", "Torah-appointed times, the complete 49-day Omer count, and explicit Shemitah/Jubilee records", "Fixed, recurring, conditional, and institutional Mishnah/Talmud calendar records", "All 35 entries or periods in the original Megillat Ta'anit", "All 26 dates in the separately labeled printed Ma'amar Aharon fast appendix witness", "All nine private wood-offering dates", "Separate diaspora and Israel festival occurrences", "Selected later and modern days"},
			Islam:        {"Weekly Jumu'ah and Ramadan", "Hajj, Eid al-Fitr, and Eid al-Adha", "Sunni and Shia observances with difference notes", "Tabular dates with crescent-sighting caveat"},
			Polytheist:   {"Documented ADF modern Neo-Pagan high-day convention", "Eight seasonal and cross-quarter dates with astronomical and local-variation caveats", "Living practice kept distinct from historical Norse and Old English records"},
			AncientWorld: {"All 45 named rows in Fowler's reconstructed fasti-antiquissimi table plus selected later Roman cycles", "Athens-anchored Attic lunar study proxy with disputed alternatives kept native-date-only", "Closed mapping of all 58 named or datable rows in UCL Digital Egypt's festival-date inventory, separated by site and period", "CDLI Ur III festival records and ORACC BTTo Q004806 scholarly-list entries with source uncertainty retained", "Two additional calendar-bearing Ugaritic KTU records plus five ritual-document records not asserted as annual holidays", "Representative nine-record regional and Panhellenic Greek native-calendar catalog", "Old Norse and Bede/Old English textual-calendar records"},
		},
	}
}

func About() AboutResponse {
	return AboutResponse{
		Methodology: []string{
			"Gregorian and Hebrew lunisolar month views share the same canonical civil days. Dates are converted locally with the arithmetic fixed Hebrew calendar and tabular Hijri calendar; the civil-day card represents the daylight portion of a sacred date that began the evening before.",
			"Historical Jewish dates are projected annually onto the modern fixed Hebrew calendar. That makes the traditional month and day explorable without claiming an exact Gregorian anniversary for eras that used an observational calendar.",
			jewishCorpusInclusionRule,
			"Jewish records are typed rather than flattened: fixed annual observances receive date occurrences; recurring counts and lunar rites receive labeled spans; drought schedules are conditional threshold or unprojected relative-stage records; Temple rosters and disputed Jubilee years remain institutional catalog records. A catalog record documents a source and does not assert that the practice occurs today.",
			"Roman festival dates are nominal same-month/day projections, not asserted conversions of pre-Julian years. Attic dates use an app-defined, Athens-anchored mean-lunation study proxy informed by scholarship and carry uncertainty because observation and archon intervention changed the civic calendar.",
			"Egyptian, Mesopotamian, Ugaritic, and regional Greek rites without a defensible unique modern date remain searchable native-calendar records in the Observance Atlas. The Egyptian inventory closes all 58 named or datable rows on the cited UCL page with separate period and site layers; UCL itself warns that Schott's linear synthesis obscures variation across time and place, and the page-bounded manifest is not a claim to every rite in Egyptian history. The app does not fabricate a single Egyptian 'Joseph-era' calendar: Joseph cannot be securely attached to one dynasty or pharaoh.",
			"Ancient ritual documents and festival calendars are not treated as synonyms. BTTo Q004806 remains a scholarly/theological Nisannu list whose column-line numbers are not days; CDLI's unnamed month-IX festival and monthly Annunītum recurrence remain probable, while its Tummal dossier remains inferential; five Ugaritic ritual tablets without annual dates remain document records rather than invented holidays.",
			"The nine regional and Panhellenic Greek records are an explicitly representative native-date/cycle catalog. Delphi, Olympia, Sparta, Argos, and Boeotia used local evidence and calendars; no universal Greek calendar or automatic Gregorian recurrence is inferred.",
			"Western Easter uses the Gregorian computus; Orthodox Pascha uses the Julian computus converted to Gregorian dates for 1900–2100.",
			"The calendar-card Moon remains a day-level mean-synodic estimate. For civil years 1900–2100, the separate Sky view uses Astronomy Engine for offline ephemeris calculations, including exact phases, seasons, eclipse and transit contacts, nodes, apsides, planetary positions, stations, ecliptic aspects, true angular approaches, twilight, rise, set, and culmination.",
			"Prayer windows use solar declination and equation-of-time calculations for the selected coordinates and IANA timezone.",
			"The Catholic schedule identifies the liturgical season and lectionary cycles, then links each date to the official USCCB readings page.",
			"Numerology uses an explicitly shown digit-reduction formula. Kabbalah and astrology are labeled interpretive lenses, separate from astronomical calculations.",
		},
		Sources: []Source{
			{Name: "USCCB Liturgical Calendar", URL: usccbSource, Use: "Catholic seasons, ranks, cycles, and principal celebrations"},
			{Name: "USCCB Daily Readings", URL: "https://bible.usccb.org/readings/calendar", Use: "Authoritative date-specific U.S. Catholic Mass readings"},
			{Name: "Orthodox Church in America · Paschal cycle", URL: ocaSource, Use: "Orthodox Paschal framing and calendar cross-check"},
			{Name: "Hebcal", URL: "https://www.hebcal.com/api-docs/", Use: "Jewish-calendar cross-check and terminology; Hebcal API content is CC BY 4.0"},
			{Name: "Sefaria · Torah appointed times", URL: "https://www.sefaria.org/Leviticus.23", Use: "Primary text for the Torah festival cycle"},
			{Name: "Mishnah Rosh Hashanah 1:1", URL: sefariaMishnahRH, Use: "The four legal and agricultural new years"},
			{Name: "Mishnah Ta'anit 4:5", URL: sefariaMishnahTaanit, Use: "Temple-era family wood-offering festivals"},
			{Name: "Megillat Ta'anit", URL: sefariaMegillatTaanit, Use: "All 35 original no-fast entries or periods, including discontinued Second Temple observances"},
			{Name: "Megillat Ta'anit · printed Ma'amar Aharon appendix", URL: lateFastAppendixSourceURL, Use: "All 26 dates in one explicitly named printed Hebrew fast-list witness, kept separate from the original Aramaic scroll"},
			{Name: "Halakhot Gedolot 18:12", URL: lateFastHalakhotURL, Use: "Geonic recension cross-check for the unstable late fast-list tradition"},
			{Name: "Tur, Orach Chayim 580", URL: lateFastTurURL, Use: "Later legal-list cross-check and date variants for supplementary fasts"},
			{Name: "Shulchan Arukh, Orach Chayim 580:2", URL: lateFastShulchanArukhURL, Use: "The later 21-date supplementary selection, distinguished from the printed 26-date appendix witness"},
			{Name: "Mishnah Ta'anit 1:4–7", URL: rainFastSourceURL, Use: "Conditional individual and communal drought-fast thresholds, stages, and post-thirteen-fast response"},
			{Name: "Mishnah Ta'anit 4:2–3", URL: maamadotSourceURL, Use: "Weekly ma'amadot creation readings and Monday-through-Thursday representative fast pattern"},
			{Name: "Bavli Yoma 69a", URL: yomaGerizimSourceURL, Use: "The variant Tevet 25 Mount Gerizim festival transmission"},
			{Name: "Bavli Berakhot 59b", URL: birkatHachamaSourceURL, Use: "The twenty-eight-year blessing-of-the-sun cycle"},
			{Name: "Bavli Sanhedrin 41b–42a", URL: birkatLevanahSourceURL, Use: "Monthly moon-blessing practice and alternative seven-/sixteen-day latest limits"},
			{Name: "Bavli Rosh Hashanah 18b–19b", URL: "https://www.sefaria.org/Rosh_Hashanah.18b-19b", Use: "Rabbinic discussion of the abrogation and surviving status of Megillat Ta'anit"},
			{Name: "Torah · Leviticus 25", URL: torahCyclesSourceURL, Use: "Shemitah and Jubilee cycle institutions"},
			{Name: "U.S. Naval Observatory · Moon phases", URL: "https://aa.usno.navy.mil/faq/moon_phases", Use: "Astronomical phase definitions and validation"},
			{Name: "NOAA Solar Calculator", URL: "https://gml.noaa.gov/grad/solcalc/calcdetails.html", Use: "Solar-position equations underlying location-aware prayer windows"},
			{Name: "Astronomy Engine", URL: "https://github.com/cosinekitty/astronomy", Use: "Offline planetary ephemerides, eclipses, phases, seasons, nodes, apsides, and local sky positions"},
			{Name: "Roman fasti and Ovid's Fasti", URL: "https://penelope.uchicago.edu/encyclopaedia_romana/calendar/antiates.html", Use: "Nominal ancient Roman festival dates and calendar classifications"},
			{Name: "Attic Inscriptions Online", URL: "https://www.atticinscriptions.com/browse/byinscriptiontype/sacrificial-calendar/", Use: "Primary epigraphic evidence for Athenian sacrificial calendars"},
			{Name: "Oxford Classical Dictionary · Lenaea", URL: atticLenaiaURL, Use: "Attested Gamelion day-12 anchor for the Lenaia"},
			{Name: "Greek Ministry of Culture · Great Dionysia", URL: atticDionysiaURL, Use: "Elaphebolion day-10 opening and reconstructed festival program"},
			{Name: "UCL Digital Egypt · Festival dates", URL: egyptianFestivalSource, Use: "Closed 58-row witness manifest for named and datable Egyptian festival entries, with site and period distinctions"},
			{Name: "University of Zurich · The Joseph Story", URL: "https://www.pentateuch.uzh.ch/en/Subprojects/Subproject-A.html", Use: "Chronological uncertainty behind the decision not to invent a single Joseph-era Egyptian calendar"},
			{Name: "CDLI Preprint 9 · Ur III festival inventory", URL: urIIICultSource, Use: "Great, Gazelle, Piglet, Ubi-bird, Ninazu, and Shulgi festivals, plus a probable unnamed month-IX festival and probable recurring Annunītum festivals at Ur"},
			{Name: "CDLI Journal 2016:1 · Tummal dossier", URL: urIIITummalSource, Use: "Inferential festival identification from clustered Drehem disbursements, retained at moderate confidence"},
			{Name: "ORACC BTTo Q004806", URL: mesopotamianBTTO, Use: "Nisannu scholarly/theological festival list; column-line references, ambiguous Ishtar/Ninurta wording, and Ulūlu festival name"},
			{Name: "Ugaritic ritual-text corpus", URL: ugaritInstitutionalInventory, Use: "KTU tablet identifiers and the distinction between calendar-bearing rites and ritual documents without proven annual dates"},
			{Name: "Dennis Pardee · Ritual and Cult at Ugarit", URL: ugaritPardeeSource, Use: "Partial native schedules in KTU 1.119 and KTU 1.112"},
			{Name: "Mission de Ras Shamra · Pardee, Les textes rituels", URL: ugaritPardeeRitualEdition, Use: "Fragmentary unnamed-month sequences in KTU 1.46 and KTU 1.109, with reconstruction limits"},
			{Name: "Hellenic Ministry of Culture · Delphi", URL: delphiOfficialFestivalSource, Use: "Official institutional account of the reorganized Pythian Games' Delphic month and cycle"},
			{Name: "Hellenic Ministry of Culture · Olympia", URL: olympiaOfficialGamesSource, Use: "Official institutional account of the Olympic Games' four-year cycle"},
			{Name: "Orkneyinga Saga", URL: orkneyingaSource, Use: "Direct witness for the medieval etiological Þorri/Þorrablót tradition"},
			{Name: "ADF Constitution · High Days", URL: adfCalendarSource, Use: "One documented living Neo-Pagan eightfold calendar convention"},
		},
		Disclaimers: []string{
			"No single list can include every local saint day, denomination, minhag, tariqa, jurisdiction, or civil observance. The catalog covers principal annual and weekly sacred observances and names its represented communities.",
			"Local Jewish calendars and rabbinic authorities should determine practical zmanim. Local mosques or Islamic authorities should determine prayer and crescent-sighting dates.",
			"Historical notices are source records, not recommendations for modern observance. Megillat Ta'anit's terse Aramaic core and later explanatory scholia are distinct textual layers, and uncertain identifications are labeled.",
			"The 26-date Ma'amar Aharon fast appendix is a textually unstable late-antique-to-Geonic tradition preserved in later printed Megillat Ta'anit witnesses, not part of the original 35-entry Aramaic no-fast scroll. Its witness-specific projections do not make all 26 dates current fast obligations.",
			"Conditional rain-fast stages, ma'amadot, Jubilee, Birkat HaLevanah, and Birkat HaChama are study records, not practical rulings. Weather, visibility, locality, legal authority, and operative institutions matter.",
			"There was no universal Greek, Roman-era pagan, Holy Roman imperial, Egyptian, Mesopotamian, Canaanite, or modern Pagan calendar. Corpus, site, era, living status, and projection method are kept visible; catalog coverage can never be literally exhaustive of every locality and period.",
			"Astronomical calculations describe geometry, not omens. Ecliptic aspects are projected angular relationships, while close approaches use true apparent sky separation. Horoscopes are explicitly symbolic entertainment prompts, not scientific predictions or advice.",
			"Rise and set calculations assume standard atmospheric refraction and an ideal horizon; local terrain, buildings, elevation above nearby ground, and weather can shift observed times.",
			"Astrology, Kabbalah correspondences, and numerology are offered for respectful study and reflection, not as scientific predictions or as substitutes for a living religious teacher.",
		},
	}
}
