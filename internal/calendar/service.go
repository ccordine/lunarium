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
	return buildMonthResponse(
		GregorianCalendar,
		year,
		int(month),
		first.Format("January 2006"),
		first,
		last.Day(),
		gregorianMonthReference(previousDate.Year(), previousDate.Month()),
		gregorianMonthReference(nextDate.Year(), nextDate.Month()),
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

	first := jdToGregorian(hebrewToJD(year, month, 1))
	dayCount := hebrewMonthDays(year, month)
	last := first.AddDate(0, 0, dayCount-1)
	if first.Year() < 1900 || last.Year() > 2100 {
		return MonthResponse{}, fmt.Errorf("Hebrew month must fall within Gregorian years 1900 through 2100")
	}

	location := requestedLocation(requested)
	return buildMonthResponse(
		HebrewCalendar,
		year,
		month,
		fmt.Sprintf("%s %d", hebrewMonthName(year, month), year),
		first,
		dayCount,
		previousHebrewMonthReference(year, month),
		nextHebrewMonthReference(year, month),
		location,
	), nil
}

func requestedLocation(requested []Location) Location {
	if len(requested) > 0 {
		return normalizeLocation(requested[0])
	}
	return DefaultLocation
}

func buildMonthResponse(system CalendarSystem, year, month int, label string, first time.Time, dayCount int, previous, next MonthReference, location Location) MonthResponse {
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
	sort.Slice(observances, func(i, j int) bool {
		if observances[i].Date == observances[j].Date {
			return observances[i].Name < observances[j].Name
		}
		return observances[i].Date < observances[j].Date
	})
	counts := map[Tradition]int{Christianity: 0, Judaism: 0, Islam: 0}
	for _, event := range observances {
		counts[event.Tradition]++
	}
	return ObservanceIndex{
		Year:        year,
		Observances: observances,
		Counts:      counts,
		Coverage: map[Tradition][]string{
			Christianity: {"Western and Orthodox Paschal cycles", "Principal Catholic solemnities and feasts", "Selected Anglican, Lutheran, Reformed, and Protestant commemorations", "Julian-calendar Nativity and Theophany"},
			Judaism:      {"Weekly Shabbat and monthly Rosh Chodesh", "Torah-appointed times and Temple rites", "Mishnah and Talmud calendar rules, including private wood-offering festivals", "All 35 entries or periods in the original Megillat Ta'anit", "Diaspora and Israel duration notes", "Selected later and modern days"},
			Islam:        {"Weekly Jumu'ah and Ramadan", "Hajj, Eid al-Fitr, and Eid al-Adha", "Sunni and Shia observances with difference notes", "Tabular dates with crescent-sighting caveat"},
		},
	}
}

func About() AboutResponse {
	return AboutResponse{
		Methodology: []string{
			"Gregorian and Hebrew lunisolar month views share the same canonical civil days. Dates are converted locally with the arithmetic fixed Hebrew calendar and tabular Hijri calendar; the civil-day card represents the daylight portion of a sacred date that began the evening before.",
			"Historical Jewish dates are projected annually onto the modern fixed Hebrew calendar. That makes the traditional month and day explorable without claiming an exact Gregorian anniversary for eras that used an observational calendar.",
			"Western Easter uses the Gregorian computus; Orthodox Pascha uses the Julian computus converted to Gregorian dates for 1900–2100.",
			"Moon phase and illumination are day-level estimates from the mean synodic month. Prayer windows use solar declination and equation-of-time calculations for the selected coordinates and IANA timezone.",
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
			{Name: "Bavli Rosh Hashanah 18b–19b", URL: "https://www.sefaria.org/Rosh_Hashanah.18b-19b", Use: "Rabbinic discussion of the abrogation and surviving status of Megillat Ta'anit"},
			{Name: "U.S. Naval Observatory · Moon phases", URL: "https://aa.usno.navy.mil/faq/moon_phases", Use: "Astronomical phase definitions and validation"},
			{Name: "NOAA Solar Calculator", URL: "https://gml.noaa.gov/grad/solcalc/calcdetails.html", Use: "Solar-position equations underlying location-aware prayer windows"},
		},
		Disclaimers: []string{
			"No single list can include every local saint day, denomination, minhag, tariqa, jurisdiction, or civil observance. The catalog covers principal annual and weekly sacred observances and names its represented communities.",
			"Local Jewish calendars and rabbinic authorities should determine practical zmanim. Local mosques or Islamic authorities should determine prayer and crescent-sighting dates.",
			"Historical notices are source records, not recommendations for modern observance. Megillat Ta'anit's terse Aramaic core and later explanatory scholia are distinct textual layers, and uncertain identifications are labeled.",
			"Astrology, Kabbalah correspondences, and numerology are offered for respectful study and reflection, not as scientific predictions or as substitutes for a living religious teacher.",
		},
	}
}
