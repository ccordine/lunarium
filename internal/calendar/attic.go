package calendar

import (
	"fmt"
	"math"
	"time"
)

// This file projects a research-oriented Attic festival calendar onto modern
// civil dates. It is deliberately not a conversion claim. Classical Athens did
// not use a mechanically fixed calendar: month openings depended on the moon,
// intercalation was civic, and officials could manipulate individual days.
//
// Projection model:
//   - estimate the northern summer solstice with Meeus's mean-solstice
//     polynomial;
//   - find the first mean conjunction whose conjunction+1-day crescent proxy
//     falls after that solstice;
//   - open Hekatombaion on that proxy date and each later month at the next
//     mean lunation;
//   - where thirteen lunations occur before the following projected new year,
//     insert Poseideon II after Poseideon.
//
// The +1-day rule is only a reproducible stand-in for first visibility. It does
// not model weather, horizon extinction, testimony by observers, archon action,
// or the historically variable full/hollow ordering of months.

const (
	atticCalendarCorpus = "Athenian Attic lunisolar festival calendar"
	atticAnchorLocation = "Athens, Attica (23.7° E mean-solar date proxy)"
	atticFestivalSource = "Robert Parker, Attic Festivals: A Check List"
	atticFestivalURL    = "https://academic.oup.com/book/11496/chapter-abstract/160243868"
	atticAIOURL         = "https://www.atticinscriptions.com/browse/byinscriptiontype/sacrificial-calendar/"
	atticMethodURL      = "https://www.cambridge.org/core/journals/annual-of-the-british-school-at-athens/article/illuminating-the-parthenon/D0F077C96D199A00C5294CA6B41B42D1"
	atticLenaiaURL      = "https://academic.oup.com/edited-volume/61673/chapter-abstract/549671129"
	atticDionysiaURL    = "https://ancienttheater.culture.gr/games/game_2/index.htm?lang=en"

	// 2000-01-06 18:14 UT, expressed as a Julian day, is a commonly used
	// modern mean-new-moon epoch. The period is the mean synodic month.
	atticMeanNewMoonEpochJD = 2451550.25972
	atticMeanSynodicDays    = 29.530588853

	// Athens longitude is about 23.7 degrees east, or 94.8 minutes ahead of
	// UTC in local mean solar time. This offset labels civil proxy dates only;
	// it is not an assertion about an ancient clock or time zone.
	atticAthensMeanSolarOffset = 95 * time.Minute
)

type atticFestivalRule struct {
	CatalogID  string
	Name       string
	Month      string
	StartDay   int
	EndDay     int
	Category   string
	Summary    string
	Meaning    string
	Practices  []string
	Confidence string
}

type atticFestivalAnnotation struct {
	NativeDateLabel string
	DateNote        string
	SourceName      string
	SourceURL       string
	CatalogOnly     bool
}

// These annotations keep alternatives and reconstructed windows visible
// instead of converting every scholarly possibility into a continuous span.
var atticFestivalAnnotations = map[string]atticFestivalAnnotation{
	"attic-boedromia": {
		NativeDateLabel: "Boedromion 7 (conventional placement; exact-day evidence is not uniform)",
		DateNote:        "The app uses day 7 as a conventional study placement and labels it low confidence; month-level attestation is stronger than the exact day.",
	},
	"attic-greater-mysteries": {
		NativeDateLabel: "Boedromion 15–23 public sequence (preliminaries from 15; final day apparently 23)",
		DateNote:        "This is a public festival-sequence window, not a claim that one identical rite ran continuously on every included day; restricted initiatory content is not reconstructed.",
	},
	"attic-proerosia": {
		NativeDateLabel: "Pyanepsion 5 proclamation; celebration likely Pyanepsion 6 (Thorikos used Boedromion)",
		DateNote:        "The projection chooses the likely Athenian celebration on day 6. Parker distinguishes the day-5 proclamation and notes local variation.",
	},
	"attic-oschophoria": {
		NativeDateLabel: "Early Pyanepsion, often reconstructed as day 6 or 7",
		DateNote:        "The exact day is disputed, so this record remains native-date-only instead of selecting one alternative as a modern anniversary.",
		CatalogOnly:     true,
	},
	"attic-lenaia": {
		NativeDateLabel: "Gamelion 12 (opening/main day); later dramatic schedule reconstructed across following days",
		DateNote:        "Oxford's reference entry supports day 12. The app does not turn debated reconstructions of the full dramatic schedule into a four-day continuous occurrence.",
		SourceName:      "Oxford Classical Dictionary · Lenaea",
		SourceURL:       atticLenaiaURL,
	},
	"attic-theogamia": {
		NativeDateLabel: "Gamelion 27; classically attested festival name Hieros Gamos",
		DateNote:        "Later and modern sources often use Theogamia; Parker identifies Hieros Gamos as the classically attested name.",
	},
	"attic-city-dionysia": {
		NativeDateLabel: "Elaphebolion 10 opening; the multi-day program varied by period",
		DateNote:        "The Greek Ministry of Culture reconstruction supports the opening on day 10. This entry anchors the festival there without asserting a fixed end date for every period.",
		SourceName:      "Greek Ministry of Culture · Great Dionysia",
		SourceURL:       atticDionysiaURL,
	},
	"attic-bendideia": {
		NativeDateLabel: "Thargelion 19 or 20",
		DateNote:        "The sources preserve alternative dates, not a two-day span; the record therefore remains native-date-only.",
		CatalogOnly:     true,
	},
}

// atticFestivalRules is an Athens-centered manifest. Lower-confidence or
// alternative dates remain explicitly labeled, and disputed alternatives can
// be catalog-only instead of being materialized as invented continuous spans.
var atticFestivalRules = []atticFestivalRule{
	{"attic-kronia", "Kronia", "Hekatombaion", 12, 12, "Civic festival", "A festival of Kronos marked the twelfth of Hekatombaion.", "The Athenian celebration remembered Kronos and a mythic age of abundance.", []string{"Communal feasting", "Historically remembered relaxation of ordinary social roles"}, "high"},
	{"attic-synoikia", "Synoikia", "Hekatombaion", 16, 16, "Civic festival", "Athens marked the political unification of Attica on Hekatombaion 16.", "The festival connected civic identity with the synoecism traditionally attributed to Theseus.", []string{"Civic sacrifice", "Commemoration of Attic unity"}, "high"},
	{"attic-panathenaia", "Panathenaia", "Hekatombaion", 28, 28, "Major civic festival", "The central procession and sacrifice for Athena are placed on Hekatombaion 28.", "Athens honored its patron through a civic festival whose scale and program changed across periods; the Great Panathenaia followed a four-year cycle.", []string{"Procession to the Acropolis", "Sacrifice to Athena", "Presentation of the peplos and contests in appropriate historical cycles"}, "high"},

	{"attic-genesia", "Genesia", "Boedromion", 5, 5, "Festival of the dead", "A public commemoration of the dead fell on Boedromion 5.", "The day situated ancestral remembrance within the civic sacred calendar.", []string{"Offerings and rites for the dead"}, "high"},
	{"attic-artemis-agrotera", "Artemis Agrotera", "Boedromion", 6, 6, "Civic sacrifice", "Athenians sacrificed to Artemis Agrotera on Boedromion 6.", "The rite was associated with Artemis as huntress and, in later civic memory, the victory at Marathon.", []string{"Public sacrifice", "Commemoration associated with Marathon"}, "high"},
	{"attic-boedromia", "Boedromia", "Boedromion", 7, 7, "Festival of Apollo", "The Boedromia honored Apollo Boedromios in the month bearing the festival's name; day 7 is retained only as a low-confidence conventional placement.", "Apollo was invoked as a helper in battle and civic danger.", []string{"Sacrifice to Apollo Boedromios"}, "low"},
	{"attic-greater-mysteries", "Greater Eleusinian Mysteries", "Boedromion", 15, 23, "Mystery festival", "The publicly recoverable sequence surrounding the Greater Mysteries is represented from preliminaries on Boedromion 15 through the apparent final day on 23.", "The state-supported rites joined Athens and Eleusis in the worship of Demeter and Kore; the initiatory content was restricted in antiquity and is not reconstructed here.", []string{"Assemblies and purification", "Procession between Athens and Eleusis", "Initiation historically restricted to participants"}, "medium"},

	{"attic-proerosia", "Proerosia", "Pyanepsion", 6, 6, "Agricultural festival", "The Proerosia is projected on likely celebration day Pyanepsion 6; day 5 belongs to its proclamation in the Athenian evidence.", "Offerings sought divine favor for sowing and the coming crop.", []string{"Agricultural offering", "Prayer before ploughing and sowing"}, "low"},
	{"attic-pyanopsia", "Pyanopsia", "Pyanepsion", 7, 7, "Festival of Apollo", "The Pyanopsia honored Apollo on Pyanepsion 7.", "The observance joined a pulse stew and the decorated eiresione branch with harvest gratitude and protection.", []string{"Offering of cooked pulses", "Carrying or dedicating the eiresione"}, "high"},
	{"attic-oschophoria", "Oschophoria", "Pyanepsion", 6, 6, "Processional festival", "The Oschophoria belonged early in Pyanepsion, but modern scholarship does not yield one uncontested numbered day.", "A procession bearing vine shoots linked Athena, Dionysos, Theseus traditions, and the grape harvest.", []string{"Procession with vine branches", "Choral and athletic observances in historical accounts"}, "low"},
	{"attic-theseia", "Theseia", "Pyanepsion", 8, 8, "Hero festival", "The Theseia commemorated Theseus on Pyanepsion 8.", "The hero's cult expressed Athenian civic memory and identity.", []string{"Sacrifice and offerings to Theseus", "Historical contests and public distribution"}, "medium"},
	{"attic-stenia", "Stenia", "Pyanepsion", 9, 9, "Women's festival", "The Stenia is projected on Pyanepsion 9 before the Thesmophoria sequence.", "Women performed rites connected with Demeter and Kore; surviving testimony is partial and later chronologies differ.", []string{"Women-only ritual in antiquity", "Ritual joking reported by ancient sources"}, "medium"},
	{"attic-thesmophoria", "Thesmophoria", "Pyanepsion", 11, 13, "Women's festival", "The Athenian Thesmophoria is projected on Pyanepsion 11–13.", "Citizen women honored Demeter and Kore in rites associated with fertility, civic continuity, and agriculture.", []string{"Women-only encampment and ritual in antiquity", "Fasting and offerings within the reconstructed sequence"}, "high"},
	{"attic-chalkeia", "Chalkeia", "Pyanepsion", 30, 30, "Craft festival", "The Chalkeia fell on the thirtieth or last day of Pyanepsion.", "Athena Ergane and Hephaistos were honored in connection with artisanship and civic craft production.", []string{"Offerings by craftspeople", "Beginning work associated with Athena's peplos in later evidence"}, "medium"},

	{"attic-haloa", "Haloa", "Poseideon", 26, 26, "Agricultural festival", "The Haloa at Eleusis is placed on Poseideon 26.", "The winter festival honored Demeter and Dionysos and centered agricultural fertility; important parts were conducted by women.", []string{"Women's feast in antiquity", "Offerings associated with fertility and cultivated foods"}, "high"},

	{"attic-lenaia", "Lenaia", "Gamelion", 12, 12, "Dionysian festival", "The Lenaia is anchored on its attested Gamelion 12 opening or main day; the complete dramatic schedule is not treated as fixed.", "Dionysos was honored with procession and dramatic competition during the winter month.", []string{"Procession and sacrifice", "Historical dramatic competitions"}, "high"},
	{"attic-theogamia", "Hieros Gamos", "Gamelion", 27, 27, "Sacred-marriage festival", "The classically named Hieros Gamos is placed on Gamelion 27.", "The divine marriage of Zeus and Hera framed reflection on marriage within the civic cult calendar.", []string{"Offerings to Zeus and Hera", "Marriage-related observance"}, "medium"},

	{"attic-anthesteria", "Anthesteria", "Anthesterion", 11, 13, "Dionysian festival", "The Anthesteria occupied Anthesterion 11–13: Pithoigia, Choes, and Chytroi.", "New wine, household and civic ritual, and the presence and departure of the dead formed a complex three-day festival.", []string{"Opening the wine jars", "Drinking contests and civic rites", "Offerings associated with the dead on the final day"}, "high"},
	{"attic-diasia", "Diasia", "Anthesterion", 23, 23, "Festival of Zeus", "The Diasia honored Zeus Meilichios on Anthesterion 23.", "A major extra-urban gathering sought the favor of Zeus in a chthonic and protective aspect.", []string{"Family and civic offerings", "Outdoor gathering"}, "high"},

	{"attic-asklepieia", "Asklepieia", "Elaphebolion", 8, 8, "Healing festival", "The Athenian Asklepieia is placed on Elaphebolion 8.", "The rite honored Asklepios immediately before the City Dionysia sequence.", []string{"Sacrifice to Asklepios", "Rites at the Athenian Asklepieion"}, "high"},
	{"attic-city-dionysia", "City Dionysia", "Elaphebolion", 10, 10, "Major Dionysian festival", "The City Dionysia is anchored on its Elaphebolion 10 festival opening; its changing multi-day program is described without assigning a universal end day.", "Athens honored Dionysos Eleuthereus through procession, sacrifice, and dramatic competitions whose detailed schedule varied over time.", []string{"Procession and sacrifice", "Historical dithyrambic, tragic, and comic competitions"}, "high"},

	{"attic-delphinia", "Delphinia", "Mounichion", 6, 6, "Festival of Apollo and Artemis", "The Delphinia was observed on Mounichion 6.", "The festival honored Apollo Delphinios and Artemis and was connected in Athenian story with Theseus's departure for Crete.", []string{"Procession and supplication", "Offerings at the Delphinion"}, "high"},
	{"attic-mounichia", "Mounichia", "Mounichion", 16, 16, "Festival of Artemis", "Artemis Mounichia was honored on Mounichion 16.", "A harbor-side cult at Mounichia connected Artemis, moonlike lights, and Athenian civic memory.", []string{"Procession and sacrifice", "Round cakes bearing lights in later descriptions"}, "high"},
	{"attic-olympieia", "Olympieia", "Mounichion", 19, 19, "Festival of Zeus", "The Olympieia is projected on Mounichion 19.", "The day honored Olympian Zeus in Athens; the scale and institutional form changed substantially across eras.", []string{"Sacrifice to Zeus Olympios", "Historical equestrian observances in some periods"}, "medium"},

	{"attic-thargelia", "Thargelia", "Thargelion", 6, 7, "Festival of Apollo and Artemis", "The Thargelia occupied Thargelion 6–7.", "Purification and first-fruit offerings preceded celebration of Apollo on his traditional seventh day.", []string{"Civic purification in antiquity", "First-fruit offerings", "Choral competition"}, "high"},
	{"attic-bendideia", "Bendideia", "Thargelion", 19, 19, "Adopted civic cult festival", "The Bendideia is attested on either Thargelion 19 or 20; those alternatives are not represented as a two-day duration.", "Athens' Thracian cult of Bendis featured distinct processions and a famous torch race, illustrating cultural exchange within the city.", []string{"Athenian and Thracian processions", "Historical horseback torch race"}, "low"},
	{"attic-plynteria", "Plynteria", "Thargelion", 25, 25, "Festival of Athena", "The Plynteria fell on Thargelion 25.", "Athena's ancient image and garments underwent ritual care while the city regarded the day as inauspicious.", []string{"Ritual washing associated with Athena's image and garments", "Temporary veiling or withdrawal of the cult image"}, "high"},

	{"attic-skira", "Skira", "Skirophorion", 12, 12, "Processional festival", "The Skira was observed on Skirophorion 12.", "A procession involving major Athenian priesthoods and women marked the approach of the year's end; interpretation of its rites remains debated.", []string{"Procession under a ritual canopy", "Women's observances in antiquity"}, "high"},
	{"attic-dipolieia", "Dipolieia", "Skirophorion", 14, 14, "Festival of Zeus", "The Dipolieia honored Zeus Polieus on Skirophorion 14.", "The festival included the unusual ox-sacrifice remembered as the Bouphonia and a ritualized inquiry into responsibility.", []string{"Sacrifice to Zeus Polieus", "Bouphonia and ritual trial in historical accounts"}, "high"},
}

type atticProjectedMonth struct {
	Name        string
	Start       time.Time
	End         time.Time // exclusive
	Intercalary bool
}

type atticProjectedYear struct {
	AnchorYear int
	Start      time.Time
	End        time.Time // exclusive
	Months     []atticProjectedMonth
}

var atticCommonMonthNames = []string{
	"Hekatombaion",
	"Metageitnion",
	"Boedromion",
	"Pyanepsion",
	"Maimakterion",
	"Poseideon",
	"Gamelion",
	"Anthesterion",
	"Elaphebolion",
	"Mounichion",
	"Thargelion",
	"Skirophorion",
}

var atticIntercalaryMonthNames = []string{
	"Hekatombaion",
	"Metageitnion",
	"Boedromion",
	"Pyanepsion",
	"Maimakterion",
	"Poseideon",
	"Poseideon II",
	"Gamelion",
	"Anthesterion",
	"Elaphebolion",
	"Mounichion",
	"Thargelion",
	"Skirophorion",
}

func atticObservances(date time.Time) []Observance {
	if date.Year() < 1900 || date.Year() > 2100 {
		return nil
	}
	day := atticCivilDay(date)
	year := atticYearForDate(day)
	month, ok := atticMonthForDate(year, day)
	if !ok || month.Intercalary {
		return nil
	}

	monthDays := daysBetween(month.Start, month.End)
	if monthDays < 1 {
		return nil
	}

	events := make([]Observance, 0, 2)
	for _, rule := range atticFestivalRules {
		annotation := atticFestivalAnnotations[rule.CatalogID]
		if rule.Month != month.Name || annotation.CatalogOnly {
			continue
		}

		// A nominal day 30 is historically meaningful even when this mean-moon
		// projection creates a 29-day civil month. In that case, place it on the
		// projected final day and retain the unaltered native label in metadata.
		startOffset := rule.StartDay - 1
		endDay := rule.EndDay
		if endDay == 0 {
			endDay = rule.StartDay
		}
		endOffset := endDay - 1
		clamped := false
		if startOffset >= monthDays {
			startOffset = monthDays - 1
			clamped = true
		}
		if endOffset >= monthDays {
			endOffset = monthDays - 1
			clamped = true
		}
		start := month.Start.AddDate(0, 0, startOffset)
		end := month.Start.AddDate(0, 0, endOffset)
		if day.Before(start) || day.After(end) {
			continue
		}

		event := atticFestivalObservance(rule)
		if clamped {
			event.DateNote += " This native day falls beyond the proxy month's civil length, so it is shown on that projected month's final day rather than inventing an overlapping date."
		}
		duration := daysBetween(start, end) + 1
		if duration > 1 {
			event = spanOccurrence(event, day, start, duration)
		} else {
			event = singleOccurrence(event, day)
		}
		event.ID = rule.CatalogID + "-" + start.Format("2006-01-02")
		events = append(events, event)
	}
	return events
}

func atticFestivalObservance(rule atticFestivalRule) Observance {
	annotation := atticFestivalAnnotations[rule.CatalogID]
	sourceName := atticFestivalSource
	sourceURL := atticFestivalURL
	if annotation.SourceName != "" {
		sourceName = annotation.SourceName
	}
	if annotation.SourceURL != "" {
		sourceURL = annotation.SourceURL
	}
	event := baseObservance(
		rule.Name,
		PolytheistAncient,
		[]string{"Classical Athens and Attica (historical study)"},
		rule.Category,
		rule.Summary,
		rule.Meaning,
		rule.Practices,
		nil,
		sourceName,
		sourceURL,
	)
	event.CatalogID = rule.CatalogID
	event.Origin = "Ancient Athenian civic and sacred calendars"
	event.ObservanceStatus = "Historical festival; modern date is an app-defined study projection"
	event.Historical = true
	event.HistoricalNote = "Festival selection follows Parker's Athens-centered checklist, supported by surviving sacrificial-calendar inscriptions collected by Attic Inscriptions Online (" + atticAIOURL + "). Details and schedules varied by period."
	event.DateCertainty = "Native Attic placement is " + rule.Confidence + " confidence; any civil date is an explicit mean-lunation reconstruction"
	event.CalendarCorpus = atticCalendarCorpus
	event.NativeDateLabel = atticNativeDate(rule)
	if annotation.NativeDateLabel != "" {
		event.NativeDateLabel = annotation.NativeDateLabel
	}
	event.AttestationLayer = "Literary and epigraphic evidence for Athenian civic cult"
	event.Era = "Primarily Classical and Hellenistic Athens; festival forms changed over time"
	event.Site = "Athens / Attica"
	event.ProjectionKind = "App-defined mean-new-moon Attic lunisolar study proxy"
	event.ProjectionStatus = "Projected, not an exact proleptic anniversary"
	event.DateConfidence = rule.Confidence
	event.AnchorLocation = atticAnchorLocation
	event.DayBoundary = "Sunset-to-sunset; civil label denotes the daylight-bearing date"
	event.StartsAtSunset = true
	event.DateNote = "Study proxy: Hekatombaion opens at the first mean conjunction + one-day crescent proxy after the northern summer solstice; later months follow mean synodic lunations. The Athens date label uses a fixed mean-solar offset. Method context: " + atticMethodURL + "."
	if annotation.DateNote != "" {
		event.DateNote += " " + annotation.DateNote
	}
	return event
}

func catalogOnlyAtticObservances() []Observance {
	var events []Observance
	for _, rule := range atticFestivalRules {
		annotation := atticFestivalAnnotations[rule.CatalogID]
		if !annotation.CatalogOnly {
			continue
		}
		event := atticFestivalObservance(rule)
		event.ID = rule.CatalogID
		event.CatalogOnly = true
		event.StartsAtSunset = false
		event.ObservanceStatus = "Historical Attic festival; alternative native dates retained without modern projection"
		event.ProjectionKind = "native-calendar-only"
		event.ProjectionStatus = "Not projected because selecting one disputed alternative would create false precision."
		event.DateNote = "Catalog-only native-calendar record. " + event.DateNote
		events = append(events, event)
	}
	return events
}

func atticNativeDate(rule atticFestivalRule) string {
	if rule.EndDay > rule.StartDay {
		return fmt.Sprintf("%s %d–%d", rule.Month, rule.StartDay, rule.EndDay)
	}
	return fmt.Sprintf("%s %d", rule.Month, rule.StartDay)
}

func atticYearForDate(date time.Time) atticProjectedYear {
	day := atticCivilDay(date)
	year := atticYearProjection(day.Year())
	if day.Before(year.Start) {
		return atticYearProjection(day.Year() - 1)
	}
	if !day.Before(year.End) {
		return atticYearProjection(day.Year() + 1)
	}
	return year
}

func atticYearProjection(anchorYear int) atticProjectedYear {
	if anchorYear < 1899 || anchorYear > 2100 {
		return atticProjectedYear{AnchorYear: anchorYear}
	}
	startIndex, start := atticNewYear(anchorYear)
	nextIndex, end := atticNewYear(anchorYear + 1)
	monthCount := nextIndex - startIndex
	names := atticCommonMonthNames
	if monthCount == 13 {
		names = atticIntercalaryMonthNames
	}

	// The bounded proxy should yield exactly 12 or 13 lunations. Do not silently
	// repair an anomalous result into a plausible-looking calendar.
	if monthCount != 12 && monthCount != 13 {
		return atticProjectedYear{AnchorYear: anchorYear, Start: start, End: end}
	}

	months := make([]atticProjectedMonth, 0, monthCount)
	for i := 0; i < monthCount; i++ {
		monthStart := atticCrescentDate(startIndex + i)
		monthEnd := end
		if i+1 < monthCount {
			monthEnd = atticCrescentDate(startIndex + i + 1)
		}
		name := names[i]
		months = append(months, atticProjectedMonth{
			Name:        name,
			Start:       monthStart,
			End:         monthEnd,
			Intercalary: name == "Poseideon II",
		})
	}

	return atticProjectedYear{
		AnchorYear: anchorYear,
		Start:      start,
		End:        end,
		Months:     months,
	}
}

func atticMonthForDate(year atticProjectedYear, date time.Time) (atticProjectedMonth, bool) {
	for _, month := range year.Months {
		if !date.Before(month.Start) && date.Before(month.End) {
			return month, true
		}
	}
	return atticProjectedMonth{}, false
}

func atticNewYear(year int) (int, time.Time) {
	solsticeJD := atticNorthernSummerSolsticeJD(year)
	index := int(math.Ceil((solsticeJD - atticMeanNewMoonEpochJD - 1.0) / atticMeanSynodicDays))
	proxyJD := atticMeanNewMoonEpochJD + float64(index)*atticMeanSynodicDays + 1.0
	if proxyJD <= solsticeJD {
		index++
	}
	return index, atticCrescentDate(index)
}

func atticCrescentDate(lunationIndex int) time.Time {
	proxyJD := atticMeanNewMoonEpochJD + float64(lunationIndex)*atticMeanSynodicDays + 1.0
	instant := atticTimeFromJulianDay(proxyJD).Add(atticAthensMeanSolarOffset)
	return time.Date(instant.Year(), instant.Month(), instant.Day(), 12, 0, 0, 0, time.UTC)
}

// Meeus's polynomial gives the mean June-solstice JDE for years near the
// present era. Delta-T is immaterial at the one-civil-day precision claimed by
// this projection.
func atticNorthernSummerSolsticeJD(year int) float64 {
	t := float64(year-2000) / 1000.0
	return 2451716.56767 +
		365241.62603*t +
		0.00325*t*t +
		0.00888*t*t*t -
		0.00030*t*t*t*t
}

func atticTimeFromJulianDay(jd float64) time.Time {
	seconds := (jd - 2440587.5) * 86400.0
	whole, fraction := math.Modf(seconds)
	return time.Unix(int64(whole), int64(math.Round(fraction*1e9))).UTC()
}

func atticCivilDay(date time.Time) time.Time {
	return time.Date(date.Year(), date.Month(), date.Day(), 12, 0, 0, 0, time.UTC)
}
