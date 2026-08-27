package calendar

import (
	"fmt"
	"math"
	"time"
)

const (
	hebrewEpoch = 347995.5
	// Midnight alignment used by the civil/tabular calendar. The religious day
	// itself begins at sunset on the preceding Gregorian evening.
	islamicEpoch = 1948438.5
)

var hebrewMonthNames = map[int]string{
	1: "Nisan", 2: "Iyar", 3: "Sivan", 4: "Tammuz", 5: "Av", 6: "Elul",
	7: "Tishrei", 8: "Cheshvan", 9: "Kislev", 10: "Tevet", 11: "Shevat",
}

var islamicMonthNames = []string{
	"Muharram", "Safar", "Rabi al-Awwal", "Rabi al-Thani", "Jumada al-Awwal", "Jumada al-Thani",
	"Rajab", "Sha'ban", "Ramadan", "Shawwal", "Dhu al-Qadah", "Dhu al-Hijjah",
}

type hebrewDate struct {
	Year  int
	Month int
	Day   int
}

type islamicDate struct {
	Year  int
	Month int
	Day   int
}

func gregorianToJD(date time.Time) float64 {
	year, month, day := date.Date()
	y := year
	m := int(month)
	if m <= 2 {
		y--
		m += 12
	}
	a := int(math.Floor(float64(y) / 100))
	b := 2 - a + int(math.Floor(float64(a)/4))
	return math.Floor(365.25*float64(y+4716)) + math.Floor(30.6001*float64(m+1)) + float64(day+b) - 1524.5
}

func jdToGregorian(jd float64) time.Time {
	julianDay := int(math.Floor(jd + 0.5))
	a := julianDay + 32044
	b := (4*a + 3) / 146097
	c := a - (146097*b)/4
	d := (4*c + 3) / 1461
	e := c - (1461*d)/4
	m := (5*e + 2) / 153

	day := e - (153*m+2)/5 + 1
	month := m + 3 - 12*(m/10)
	year := 100*b + d - 4800 + m/10
	return dateAt(year, time.Month(month), day)
}

func hebrewLeap(year int) bool {
	return positiveMod(7*year+1, 19) < 7
}

func hebrewMonthsInYear(year int) int {
	if hebrewLeap(year) {
		return 13
	}
	return 12
}

func hebrewDelay1(year int) int {
	months := (235*year - 234) / 19
	parts := 12084 + 13753*months
	day := 29*months + parts/25920
	if positiveMod(3*(day+1), 7) < 3 {
		day++
	}
	return day
}

func hebrewDelay2(year int) int {
	last := hebrewDelay1(year - 1)
	present := hebrewDelay1(year)
	next := hebrewDelay1(year + 1)
	if next-present == 356 {
		return 2
	}
	if present-last == 382 {
		return 1
	}
	return 0
}

func hebrewNewYear(year int) float64 {
	return hebrewEpoch + float64(hebrewDelay1(year)+hebrewDelay2(year))
}

func hebrewYearDays(year int) int {
	return int(math.Round(hebrewNewYear(year+1) - hebrewNewYear(year)))
}

func longCheshvan(year int) bool { return hebrewYearDays(year)%10 == 5 }
func shortKislev(year int) bool  { return hebrewYearDays(year)%10 == 3 }

func hebrewMonthDays(year, month int) int {
	if month == 2 || month == 4 || month == 6 || month == 10 || month == 13 {
		return 29
	}
	if month == 12 && !hebrewLeap(year) {
		return 29
	}
	if month == 8 && !longCheshvan(year) {
		return 29
	}
	if month == 9 && shortKislev(year) {
		return 29
	}
	return 30
}

func hebrewToJD(year, month, day int) float64 {
	jd := hebrewNewYear(year) + float64(day+1)
	if month < 7 {
		for m := 7; m <= hebrewMonthsInYear(year); m++ {
			jd += float64(hebrewMonthDays(year, m))
		}
		for m := 1; m < month; m++ {
			jd += float64(hebrewMonthDays(year, m))
		}
	} else {
		for m := 7; m < month; m++ {
			jd += float64(hebrewMonthDays(year, m))
		}
	}
	return jd
}

func jdToHebrew(jd float64) hebrewDate {
	jd = math.Floor(jd) + 0.5
	year := int(math.Floor((jd-hebrewEpoch)/365.2468)) + 1
	for jd >= hebrewToJD(year+1, 7, 1) {
		year++
	}
	for jd < hebrewToJD(year, 7, 1) {
		year--
	}
	month := 7
	if jd >= hebrewToJD(year, 1, 1) {
		month = 1
	}
	for jd > hebrewToJD(year, month, hebrewMonthDays(year, month)) {
		month++
		if month > hebrewMonthsInYear(year) {
			month = 1
		}
	}
	day := int(jd-hebrewToJD(year, month, 1)) + 1
	return hebrewDate{Year: year, Month: month, Day: day}
}

func hebrewMonthName(year, month int) string {
	if month == 12 {
		if hebrewLeap(year) {
			return "Adar I"
		}
		return "Adar"
	}
	if month == 13 {
		return "Adar II"
	}
	return hebrewMonthNames[month]
}

func islamicLeap(year int) bool {
	return positiveMod(11*year+14, 30) < 11
}

func islamicMonthDays(year, month int) int {
	if month%2 == 1 || (month == 12 && islamicLeap(year)) {
		return 30
	}
	return 29
}

func islamicToJD(year, month, day int) float64 {
	return float64(day) + math.Ceil(29.5*float64(month-1)) + float64((year-1)*354) +
		math.Floor(float64(3+11*year)/30) + islamicEpoch - 1
}

func jdToIslamic(jd float64) islamicDate {
	jd = math.Floor(jd) + 0.5
	year := int(math.Floor((30*(jd-islamicEpoch) + 10646) / 10631))
	month := int(math.Min(12, math.Ceil((jd-(29+islamicToJD(year, 1, 1)))/29.5)+1))
	if month < 1 {
		month = 1
	}
	day := int(jd-islamicToJD(year, month, 1)) + 1
	return islamicDate{Year: year, Month: month, Day: day}
}

func sacredDates(date time.Time) SacredDates {
	jd := gregorianToJD(date)
	h := jdToHebrew(jd)
	i := jdToIslamic(jd)
	hm := hebrewMonthName(h.Year, h.Month)
	im := islamicMonthNames[i.Month-1]
	return SacredDates{
		Hebrew:            fmt.Sprintf("%d %s %d", h.Day, hm, h.Year),
		HebrewDay:         h.Day,
		HebrewMonth:       hm,
		HebrewMonthNumber: h.Month,
		HebrewYear:        h.Year,
		Islamic:           fmt.Sprintf("%d %s %d AH", i.Day, im, i.Year),
		IslamicDay:        i.Day,
		IslamicMonth:      im,
		IslamicYear:       i.Year,
	}
}

func positiveMod(value, divisor int) int {
	result := value % divisor
	if result < 0 {
		return result + divisor
	}
	return result
}
