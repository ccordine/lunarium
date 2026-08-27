package calendar

import (
	"fmt"
	"math"
	"time"
)

var DefaultLocation = Location{
	Name:      "New York, NY",
	Latitude:  40.7128,
	Longitude: -74.0060,
	Timezone:  "America/New_York",
}

type solarDay struct {
	date        time.Time
	latitude    float64
	longitude   float64
	offsetMins  float64
	declination float64
	equation    float64
	noonUTC     float64
}

func prayerSchedules(date time.Time, location Location) []PrayerSchedule {
	location = normalizeLocation(location)
	solar := newSolarDay(date, location)
	sunrise, hasSunrise := solar.morningAtAltitude(-0.833)
	sunset, hasSunset := solar.eveningAtAltitude(-0.833)
	noon := solar.local(solar.noonUTC)

	christian := PrayerSchedule{
		Tradition: Christianity,
		Name:      "Christian daily prayer",
		Times: []PrayerTime{
			{Name: "Office of Readings", Time: "06:00", Note: "May be prayed at any hour"},
			{Name: "Lauds · Morning Prayer", Time: preferredSolarTime(sunrise, hasSunrise, "07:00")},
			{Name: "Terce · Midmorning", Time: "09:00", Note: "Optional daytime hour"},
			{Name: "Sext · Midday", Time: formatMinute(noon), Note: "Solar noon"},
			{Name: "None · Midafternoon", Time: "15:00", Note: "Optional daytime hour"},
			{Name: "Vespers · Evening Prayer", Time: preferredSolarTime(sunset, hasSunset, "18:00")},
			{Name: "Compline · Night Prayer", Time: "21:00"},
		},
		Method: "A practical Liturgy of the Hours rhythm: Morning and Evening Prayer follow local sunrise and sunset; daytime hours use customary clock times.",
		Caveat: "Christian practice varies by church and vocation. The Roman rite permits flexibility; this is a devotional plan, not a canonical obligation timetable.",
	}

	fajr, hasFajr := solar.morningAtAltitude(-18)
	isha, hasIsha := solar.eveningAtAltitude(-17)
	asr, hasAsr := solar.asr(1)
	islamic := PrayerSchedule{
		Tradition: Islam,
		Name:      "Five daily salah",
		Times: []PrayerTime{
			{Name: "Fajr", Time: availableTime(fajr, hasFajr), Note: "18° dawn"},
			{Name: "Sunrise", Time: availableTime(sunrise, hasSunrise), Note: "Prayer window boundary"},
			{Name: "Dhuhr", Time: formatMinute(noon + 2), Note: "Just after solar noon"},
			{Name: "Asr", Time: availableTime(asr, hasAsr), Note: "Standard shadow factor 1"},
			{Name: "Maghrib", Time: availableTime(sunset, hasSunset), Note: "At sunset"},
			{Name: "Isha", Time: availableTime(isha, hasIsha), Note: "17° nightfall"},
		},
		Method: "Solar calculation using Muslim World League-style angles (Fajr 18°, Isha 17°) and standard Asr shadow length.",
		Caveat: "Mosques use different conventions, elevation adjustments, and high-latitude rules. Confirm times with a trusted local mosque.",
	}

	dawn, hasDawn := solar.morningAtAltitude(-16.1)
	nightfall, hasNightfall := solar.eveningAtAltitude(-8.5)
	latestShema, latestShacharit, minchaGedola, minchaKetana := math.NaN(), math.NaN(), math.NaN(), math.NaN()
	if hasSunrise && hasSunset {
		seasonalHour := (sunset - sunrise) / 12
		latestShema = sunrise + 3*seasonalHour
		latestShacharit = sunrise + 4*seasonalHour
		minchaGedola = sunrise + 6.5*seasonalHour
		minchaKetana = sunrise + 9.5*seasonalHour
	}
	jewish := PrayerSchedule{
		Tradition: Judaism,
		Name:      "Jewish daily prayer windows",
		Times: []PrayerTime{
			{Name: "Alot hashachar", Time: availableTime(dawn, hasDawn), Note: "16.1° dawn; opinions vary"},
			{Name: "Sunrise · Netz", Time: availableTime(sunrise, hasSunrise)},
			{Name: "Latest Shema", Time: availableTime(latestShema, !math.IsNaN(latestShema)), Note: "3 seasonal hours after sunrise"},
			{Name: "Latest Shacharit", Time: availableTime(latestShacharit, !math.IsNaN(latestShacharit)), Note: "4 seasonal hours after sunrise"},
			{Name: "Chatzot", Time: formatMinute(noon), Note: "Solar midday"},
			{Name: "Mincha gedola", Time: availableTime(minchaGedola, !math.IsNaN(minchaGedola)), Note: "Earliest common Mincha window"},
			{Name: "Mincha ketana", Time: availableTime(minchaKetana, !math.IsNaN(minchaKetana)), Note: "Preferred later Mincha point"},
			{Name: "Sunset · Shkiah", Time: availableTime(sunset, hasSunset)},
			{Name: "Nightfall · Tzeit", Time: availableTime(nightfall, hasNightfall), Note: "8.5° calculation; opinions vary"},
		},
		Method: "Proportional hours (sha'ot zmaniyot) divide local sunrise-to-sunset into twelve; dawn and nightfall use stated solar angles.",
		Caveat: "Halachic authorities use multiple dawn, Shema, sunset, and nightfall opinions. Confirm practical times with a rabbi or trusted local luach.",
	}

	return []PrayerSchedule{christian, jewish, islamic}
}

func normalizeLocation(location Location) Location {
	if location.Name == "" {
		location.Name = DefaultLocation.Name
	}
	if location.Timezone == "" {
		location.Timezone = DefaultLocation.Timezone
	}
	if location.Latitude < -66 || location.Latitude > 66 || location.Longitude < -180 || location.Longitude > 180 {
		return DefaultLocation
	}
	if _, err := time.LoadLocation(location.Timezone); err != nil {
		location.Timezone = "UTC"
	}
	return location
}

func newSolarDay(date time.Time, location Location) solarDay {
	loc, _ := time.LoadLocation(location.Timezone)
	localNoon := time.Date(date.Year(), date.Month(), date.Day(), 12, 0, 0, 0, loc)
	_, offset := localNoon.Zone()
	jd := gregorianToJD(date)
	t := (jd - 2451545.0) / 36525
	l0 := normalizeDegrees(280.46646 + t*(36000.76983+t*0.0003032))
	m := 357.52911 + t*(35999.05029-0.0001537*t)
	e := 0.016708634 - t*(0.000042037+0.0000001267*t)
	mr := degreesToRadians(m)
	c := math.Sin(mr)*(1.914602-t*(0.004817+0.000014*t)) + math.Sin(2*mr)*(0.019993-0.000101*t) + math.Sin(3*mr)*0.000289
	trueLong := l0 + c
	omega := 125.04 - 1934.136*t
	appLong := trueLong - 0.00569 - 0.00478*math.Sin(degreesToRadians(omega))
	meanObliq := 23 + (26+(21.448-t*(46.815+t*(0.00059-t*0.001813)))/60)/60
	obliq := meanObliq + 0.00256*math.Cos(degreesToRadians(omega))
	decl := radiansToDegrees(math.Asin(math.Sin(degreesToRadians(obliq)) * math.Sin(degreesToRadians(appLong))))
	y := math.Pow(math.Tan(degreesToRadians(obliq)/2), 2)
	l0r := degreesToRadians(l0)
	eq := 4 * radiansToDegrees(y*math.Sin(2*l0r)-2*e*math.Sin(mr)+4*e*y*math.Sin(mr)*math.Cos(2*l0r)-0.5*y*y*math.Sin(4*l0r)-1.25*e*e*math.Sin(2*mr))
	noonUTC := 720 - 4*location.Longitude - eq
	return solarDay{date: date, latitude: location.Latitude, longitude: location.Longitude, offsetMins: float64(offset / 60), declination: decl, equation: eq, noonUTC: noonUTC}
}

func (s solarDay) morningAtAltitude(altitude float64) (float64, bool) {
	h, ok := s.hourAngle(altitude)
	if !ok {
		return 0, false
	}
	return s.local(s.noonUTC - 4*h), true
}

func (s solarDay) eveningAtAltitude(altitude float64) (float64, bool) {
	h, ok := s.hourAngle(altitude)
	if !ok {
		return 0, false
	}
	return s.local(s.noonUTC + 4*h), true
}

func (s solarDay) hourAngle(altitude float64) (float64, bool) {
	lat := degreesToRadians(s.latitude)
	decl := degreesToRadians(s.declination)
	cosH := (math.Sin(degreesToRadians(altitude)) - math.Sin(lat)*math.Sin(decl)) / (math.Cos(lat) * math.Cos(decl))
	if cosH < -1 || cosH > 1 {
		return 0, false
	}
	return radiansToDegrees(math.Acos(cosH)), true
}

func (s solarDay) asr(shadowFactor float64) (float64, bool) {
	latDelta := math.Abs(degreesToRadians(s.latitude - s.declination))
	altitude := -radiansToDegrees(math.Atan(1 / (shadowFactor + math.Tan(latDelta))))
	return s.eveningAtAltitude(altitude)
}

func (s solarDay) local(utcMinute float64) float64 {
	return normalizeMinute(utcMinute + s.offsetMins)
}

func availableTime(minute float64, available bool) string {
	if !available || math.IsNaN(minute) {
		return "—"
	}
	return formatMinute(minute)
}

func preferredSolarTime(minute float64, available bool, fallback string) string {
	if !available {
		return fallback
	}
	return formatMinute(minute)
}

func formatMinute(minute float64) string {
	minute = normalizeMinute(minute)
	total := int(math.Round(minute)) % 1440
	return fmt.Sprintf("%02d:%02d", total/60, total%60)
}

func normalizeMinute(value float64) float64 {
	value = math.Mod(value, 1440)
	if value < 0 {
		value += 1440
	}
	return value
}

func normalizeDegrees(value float64) float64 {
	value = math.Mod(value, 360)
	if value < 0 {
		value += 360
	}
	return value
}

func degreesToRadians(value float64) float64 { return value * math.Pi / 180 }
func radiansToDegrees(value float64) float64 { return value * 180 / math.Pi }
