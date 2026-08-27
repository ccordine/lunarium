package calendar

import (
	"math"
	"time"
)

const synodicMonth = 29.530588853

// The reference is the well-established new moon near J2000. This mean-lunation
// model is intended for a day-level calendar, not navigation or observatory use.
var referenceNewMoon = time.Date(2000, time.January, 6, 18, 14, 0, 0, time.UTC)

func moonForDate(date time.Time) Moon {
	noon := time.Date(date.Year(), date.Month(), date.Day(), 12, 0, 0, 0, time.UTC)
	days := noon.Sub(referenceNewMoon).Hours() / 24
	age := math.Mod(days, synodicMonth)
	if age < 0 {
		age += synodicMonth
	}
	fraction := age / synodicMonth
	illumination := (1 - math.Cos(2*math.Pi*fraction)) / 2
	name, emoji := phaseLabel(fraction)
	return Moon{
		Phase:        name,
		Emoji:        emoji,
		AgeDays:      round(age, 1),
		Illumination: round(illumination*100, 1),
		Waxing:       fraction < 0.5,
		Angle:        round(fraction*360, 1),
		AccuracyNote: "Day-level estimate from the mean synodic month; exact phase instants can differ by several hours.",
	}
}

func phaseLabel(fraction float64) (string, string) {
	index := int(math.Floor(fraction*8+0.5)) % 8
	labels := []string{"New Moon", "Waxing Crescent", "First Quarter", "Waxing Gibbous", "Full Moon", "Waning Gibbous", "Last Quarter", "Waning Crescent"}
	emoji := []string{"🌑", "🌒", "🌓", "🌔", "🌕", "🌖", "🌗", "🌘"}
	return labels[index], emoji[index]
}

func round(value float64, places int) float64 {
	p := math.Pow10(places)
	return math.Round(value*p) / p
}
