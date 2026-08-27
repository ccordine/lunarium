package calendar

import (
	"fmt"
	"strings"
	"time"
)

func catholicReadingForDate(date time.Time, events []Observance) CatholicReading {
	season := catholicSeason(date)
	title := catholicDayTitle(date, season, events)
	liturgicalYear := date.Year()
	if !date.Before(adventStart(date.Year())) {
		liturgicalYear++
	}
	cycleNames := []string{"A", "B", "C"}
	cycleIndex := positiveMod(liturgicalYear-2026, 3)
	weekdayCycle := "I"
	if date.Year()%2 == 0 {
		weekdayCycle = "II"
	}
	return CatholicReading{
		Title:         "Daily Mass readings",
		LiturgicalDay: title,
		Season:        season,
		SundayCycle:   cycleNames[cycleIndex],
		WeekdayCycle:  weekdayCycle,
		SourceName:    "United States Conference of Catholic Bishops",
		SourceURL:     "https://bible.usccb.org/bible/readings/" + date.Format("010206") + ".cfm",
		ScheduleNote:  "Solemnities and feasts use their proper readings. The official USCCB page resolves optional memorials and local calendar choices.",
	}
}

func catholicDayTitle(date time.Time, season string, events []Observance) string {
	priority := ""
	for _, event := range events {
		if event.Tradition != Christianity || !containsCommunity(event.Communities, "Catholic") {
			continue
		}
		if event.Rank == "Solemnity" || event.Rank == "Triduum" || priority == "" {
			priority = event.Name
		}
	}
	if priority != "" {
		return priority
	}
	weekday := date.Weekday().String()
	switch season {
	case "Advent":
		week := daysBetween(adventStart(date.Year()), date)/7 + 1
		if date.Weekday() == time.Sunday {
			return fmt.Sprintf("%s Sunday of Advent", ordinalWord(week))
		}
		return fmt.Sprintf("%s of the %s week of Advent", weekday, ordinalWord(week))
	case "Lent":
		ash := westernEaster(date.Year()).AddDate(0, 0, -46)
		week := (daysBetween(ash, date)+3)/7 + 1
		if date.Weekday() == time.Sunday {
			return fmt.Sprintf("%s Sunday of Lent", ordinalWord(week))
		}
		return fmt.Sprintf("%s of the %s week of Lent", weekday, ordinalWord(week))
	case "Easter":
		easter := westernEaster(date.Year())
		week := daysBetween(easter, date)/7 + 1
		if date.Weekday() == time.Sunday {
			return fmt.Sprintf("%s Sunday of Easter", ordinalWord(week))
		}
		return fmt.Sprintf("%s of the %s week of Easter", weekday, ordinalWord(week))
	case "Christmas":
		return weekday + " in the Christmas season"
	case "Triduum":
		return weekday + " of the Sacred Paschal Triduum"
	default:
		return weekday + " in Ordinary Time"
	}
}

func catholicSeason(date time.Time) string {
	year := date.Year()
	easter := westernEaster(year)
	ashWednesday := easter.AddDate(0, 0, -46)
	triduum := easter.AddDate(0, 0, -3)
	pentecost := easter.AddDate(0, 0, 49)
	advent := adventStart(year)
	if !date.Before(advent) && date.Before(dateAt(year, time.December, 25)) {
		return "Advent"
	}
	if !date.Before(dateAt(year, time.December, 25)) || date.Before(baptismOfLord(year)) {
		return "Christmas"
	}
	if !date.Before(ashWednesday) && date.Before(triduum) {
		return "Lent"
	}
	if !date.Before(triduum) && date.Before(easter) {
		return "Triduum"
	}
	if !date.Before(easter) && !date.After(pentecost) {
		return "Easter"
	}
	return "Ordinary Time"
}

func baptismOfLord(year int) time.Time {
	// In the U.S. calendar Epiphany is transferred to the Sunday from Jan 2–8.
	epiphany := dateAt(year, time.January, 2)
	for epiphany.Weekday() != time.Sunday {
		epiphany = epiphany.AddDate(0, 0, 1)
	}
	if epiphany.Day() >= 7 {
		return epiphany.AddDate(0, 0, 1)
	}
	return epiphany.AddDate(0, 0, 7)
}

func containsCommunity(communities []string, needle string) bool {
	for _, community := range communities {
		if strings.Contains(strings.ToLower(community), strings.ToLower(needle)) {
			return true
		}
	}
	return false
}

func ordinalWord(number int) string {
	words := map[int]string{1: "First", 2: "Second", 3: "Third", 4: "Fourth", 5: "Fifth", 6: "Sixth", 7: "Seventh"}
	if word, ok := words[number]; ok {
		return word
	}
	return fmt.Sprintf("Week %d", number)
}
