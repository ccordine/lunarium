package calendar

import "time"

type Tradition string

type CalendarSystem string

const (
	Christianity Tradition = "christianity"
	Judaism      Tradition = "judaism"
	Islam        Tradition = "islam"

	GregorianCalendar CalendarSystem = "gregorian"
	HebrewCalendar    CalendarSystem = "hebrew"
)

type Observance struct {
	ID               string    `json:"id"`
	CatalogID        string    `json:"catalogId,omitempty"`
	Name             string    `json:"name"`
	Tradition        Tradition `json:"tradition"`
	Communities      []string  `json:"communities"`
	Date             string    `json:"date"`
	EndDate          string    `json:"endDate,omitempty"`
	Category         string    `json:"category"`
	Rank             string    `json:"rank,omitempty"`
	Summary          string    `json:"summary"`
	Meaning          string    `json:"meaning"`
	Practices        []string  `json:"practices"`
	Scripture        []string  `json:"scripture,omitempty"`
	StartsAtSunset   bool      `json:"startsAtSunset"`
	DateNote         string    `json:"dateNote,omitempty"`
	Origin           string    `json:"origin,omitempty"`
	ObservanceStatus string    `json:"observanceStatus,omitempty"`
	Historical       bool      `json:"historical,omitempty"`
	HistoricalNote   string    `json:"historicalNote,omitempty"`
	DateCertainty    string    `json:"dateCertainty,omitempty"`
	SourceName       string    `json:"sourceName"`
	SourceURL        string    `json:"sourceUrl"`
	DayIndex         int       `json:"dayIndex,omitempty"`
	DurationDays     int       `json:"durationDays,omitempty"`
	LiturgicalColor  string    `json:"liturgicalColor,omitempty"`
}

type Moon struct {
	Phase        string  `json:"phase"`
	Emoji        string  `json:"emoji"`
	AgeDays      float64 `json:"ageDays"`
	Illumination float64 `json:"illumination"`
	Waxing       bool    `json:"waxing"`
	Angle        float64 `json:"angle"`
	AccuracyNote string  `json:"accuracyNote"`
}

type SacredDates struct {
	Hebrew            string `json:"hebrew"`
	HebrewDay         int    `json:"hebrewDay"`
	HebrewMonth       string `json:"hebrewMonth"`
	HebrewMonthNumber int    `json:"hebrewMonthNumber"`
	HebrewYear        int    `json:"hebrewYear"`
	Islamic           string `json:"islamic"`
	IslamicDay        int    `json:"islamicDay"`
	IslamicMonth      string `json:"islamicMonth"`
	IslamicYear       int    `json:"islamicYear"`
}

type KabbalahLens struct {
	Month    string `json:"month"`
	Letter   string `json:"letter"`
	Sign     string `json:"sign"`
	Tribe    string `json:"tribe"`
	Sense    string `json:"sense"`
	Theme    string `json:"theme"`
	Practice string `json:"practice"`
	Caveat   string `json:"caveat"`
}

type AstrologyLens struct {
	SunSign string `json:"sunSign"`
	Symbol  string `json:"symbol"`
	Element string `json:"element"`
	Mode    string `json:"mode"`
	Theme   string `json:"theme"`
	Caveat  string `json:"caveat"`
}

type NumerologyLens struct {
	Score  int    `json:"score"`
	Title  string `json:"title"`
	Theme  string `json:"theme"`
	Prompt string `json:"prompt"`
	Method string `json:"method"`
	Caveat string `json:"caveat"`
}

type PrayerTime struct {
	Name string `json:"name"`
	Time string `json:"time"`
	Note string `json:"note,omitempty"`
}

type PrayerSchedule struct {
	Tradition Tradition    `json:"tradition"`
	Name      string       `json:"name"`
	Times     []PrayerTime `json:"times"`
	Method    string       `json:"method"`
	Caveat    string       `json:"caveat"`
}

type Location struct {
	Name      string  `json:"name"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Timezone  string  `json:"timezone"`
}

type CatholicReading struct {
	Title         string `json:"title"`
	LiturgicalDay string `json:"liturgicalDay"`
	Season        string `json:"season"`
	SundayCycle   string `json:"sundayCycle"`
	WeekdayCycle  string `json:"weekdayCycle"`
	SourceName    string `json:"sourceName"`
	SourceURL     string `json:"sourceUrl"`
	ScheduleNote  string `json:"scheduleNote"`
}

type Day struct {
	Date        string           `json:"date"`
	Day         int              `json:"day"`
	Weekday     string           `json:"weekday"`
	IsToday     bool             `json:"isToday"`
	SacredDates SacredDates      `json:"sacredDates"`
	Moon        Moon             `json:"moon"`
	Kabbalah    KabbalahLens     `json:"kabbalah"`
	Astrology   AstrologyLens    `json:"astrology"`
	Numerology  NumerologyLens   `json:"numerology"`
	Prayers     []PrayerSchedule `json:"prayers"`
	Reading     CatholicReading  `json:"reading"`
	Observances []Observance     `json:"observances"`
}

type MonthReference struct {
	Year  int    `json:"year"`
	Month int    `json:"month"`
	Label string `json:"label"`
}

type MonthResponse struct {
	CalendarSystem  CalendarSystem `json:"calendarSystem"`
	Year            int            `json:"year"`
	Month           int            `json:"month"`
	Label           string         `json:"label"`
	StartDate       string         `json:"startDate"`
	EndDate         string         `json:"endDate"`
	Previous        MonthReference `json:"previous"`
	Next            MonthReference `json:"next"`
	FirstWeekday    int            `json:"firstWeekday"`
	Days            []Day          `json:"days"`
	ObservanceCount int            `json:"observanceCount"`
	Coverage        []string       `json:"coverage"`
	Location        Location       `json:"location"`
	GeneratedAt     time.Time      `json:"generatedAt"`
}

type ObservanceIndex struct {
	Year        int                    `json:"year"`
	Observances []Observance           `json:"observances"`
	Counts      map[Tradition]int      `json:"counts"`
	Coverage    map[Tradition][]string `json:"coverage"`
}

type Source struct {
	Name string `json:"name"`
	URL  string `json:"url"`
	Use  string `json:"use"`
}

type AboutResponse struct {
	Methodology []string `json:"methodology"`
	Sources     []Source `json:"sources"`
	Disclaimers []string `json:"disclaimers"`
}
