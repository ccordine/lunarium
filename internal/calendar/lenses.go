package calendar

import "time"

type kabbalahCorrespondence struct {
	letter, sign, tribe, sense, theme, practice string
}

var kabbalahMonths = map[string]kabbalahCorrespondence{
	"Nisan":    {"ה · Heh", "Aries", "Judah", "Speech", "Liberation and sacred beginning", "Name one constriction you are ready to leave, then take one practical step toward freedom."},
	"Iyar":     {"ו · Vav", "Taurus", "Issachar", "Thought", "Healing through patient refinement", "Choose one small practice you can repeat consistently rather than dramatically."},
	"Sivan":    {"ז · Zayin", "Gemini", "Zebulun", "Walking", "Revelation, relationship, and learning", "Study a short sacred passage with another person and listen for what changes between you."},
	"Tammuz":   {"ח · Chet", "Cancer", "Reuben", "Sight", "Guarding vision and beginning repair", "Notice what you habitually look toward; gently redirect attention toward what nourishes."},
	"Av":       {"ט · Tet", "Leo", "Simeon", "Hearing", "Grief transformed into consolation", "Make room for an honest lament, then listen for the need beneath it."},
	"Elul":     {"י · Yod", "Virgo", "Gad", "Action", "Return, discernment, and preparation", "Review the day without judgment: repair one harm and recognize one act of goodness."},
	"Tishrei":  {"ל · Lamed", "Libra", "Ephraim", "Touch", "Balance, judgment, and renewal", "Set one intention that joins accountability with compassion."},
	"Cheshvan": {"נ · Nun", "Scorpio", "Manasseh", "Smell", "Integration after sacred intensity", "Bring one insight from ritual into an ordinary task."},
	"Kislev":   {"ס · Samekh", "Sagittarius", "Benjamin", "Sleep", "Trust, dreams, and light within darkness", "Record a dream or quiet hope, then identify the support it needs."},
	"Tevet":    {"ע · Ayin", "Capricorn", "Dan", "Anger", "Discipline and the redirection of power", "Pause before reacting and ask which boundary or value needs clear expression."},
	"Shevat":   {"צ · Tzadi", "Aquarius", "Asher", "Taste", "Rooted renewal and the wisdom of trees", "Eat one fruit attentively and consider the unseen systems that sustained it."},
	"Adar":     {"ק · Kuf", "Pisces", "Naphtali", "Laughter", "Joy, reversals, and hidden possibility", "Practice a form of joy that makes room for, rather than denies, complexity."},
	"Adar I":   {"ק · Kuf", "Pisces", "Naphtali", "Laughter", "Spacious preparation and hidden possibility", "Use the leap month as unhurried space for a practice you usually rush."},
	"Adar II":  {"ק · Kuf", "Pisces", "Naphtali", "Laughter", "Joy, reversals, and hidden possibility", "Practice a form of joy that makes room for, rather than denies, complexity."},
}

func kabbalahForMonth(month string) KabbalahLens {
	c := kabbalahMonths[month]
	return KabbalahLens{
		Month:    month,
		Letter:   c.letter,
		Sign:     c.sign,
		Tribe:    c.tribe,
		Sense:    c.sense,
		Theme:    c.theme,
		Practice: c.practice,
		Caveat:   "Month correspondences draw on one widely taught reading of Sefer Yetzirah; assignments and interpretations vary across Jewish mystical lineages.",
	}
}

type zodiacSign struct {
	name, symbol, element, mode, theme string
}

func astrologyForDate(date time.Time) AstrologyLens {
	m, d := date.Month(), date.Day()
	var sign zodiacSign
	switch {
	case (m == time.March && d >= 21) || (m == time.April && d <= 19):
		sign = zodiacSign{"Aries", "♈", "Fire", "Cardinal", "Initiative and courageous beginnings"}
	case (m == time.April && d >= 20) || (m == time.May && d <= 20):
		sign = zodiacSign{"Taurus", "♉", "Earth", "Fixed", "Steadiness, embodiment, and values"}
	case (m == time.May && d >= 21) || (m == time.June && d <= 20):
		sign = zodiacSign{"Gemini", "♊", "Air", "Mutable", "Curiosity, language, and exchange"}
	case (m == time.June && d >= 21) || (m == time.July && d <= 22):
		sign = zodiacSign{"Cancer", "♋", "Water", "Cardinal", "Care, memory, and belonging"}
	case (m == time.July && d >= 23) || (m == time.August && d <= 22):
		sign = zodiacSign{"Leo", "♌", "Fire", "Fixed", "Creative presence and generosity"}
	case (m == time.August && d >= 23) || (m == time.September && d <= 22):
		sign = zodiacSign{"Virgo", "♍", "Earth", "Mutable", "Discernment, service, and craft"}
	case (m == time.September && d >= 23) || (m == time.October && d <= 22):
		sign = zodiacSign{"Libra", "♎", "Air", "Cardinal", "Reciprocity, beauty, and balance"}
	case (m == time.October && d >= 23) || (m == time.November && d <= 21):
		sign = zodiacSign{"Scorpio", "♏", "Water", "Fixed", "Depth, honesty, and transformation"}
	case (m == time.November && d >= 22) || (m == time.December && d <= 21):
		sign = zodiacSign{"Sagittarius", "♐", "Fire", "Mutable", "Meaning, exploration, and perspective"}
	case (m == time.December && d >= 22) || (m == time.January && d <= 19):
		sign = zodiacSign{"Capricorn", "♑", "Earth", "Cardinal", "Responsibility, structure, and mastery"}
	case (m == time.January && d >= 20) || (m == time.February && d <= 18):
		sign = zodiacSign{"Aquarius", "♒", "Air", "Fixed", "Community, imagination, and reform"}
	default:
		sign = zodiacSign{"Pisces", "♓", "Water", "Mutable", "Compassion, mystery, and surrender"}
	}
	return AstrologyLens{
		SunSign: sign.name,
		Symbol:  sign.symbol,
		Element: sign.element,
		Mode:    sign.mode,
		Theme:   sign.theme,
		Caveat:  "Tropical Sun-sign astrology is a symbolic interpretive tradition, not an astronomical or scientific prediction.",
	}
}

func numerologyForDate(date time.Time) NumerologyLens {
	y, m, d := date.Date()
	sum := digitSum(y) + digitSum(int(m)) + digitSum(d)
	for sum > 9 {
		sum = digitSum(sum)
	}
	titles := map[int]string{1: "The Initiator", 2: "The Diplomat", 3: "The Communicator", 4: "The Builder", 5: "The Explorer", 6: "The Nurturer", 7: "The Seeker", 8: "The Steward", 9: "The Humanitarian"}
	themes := map[int]string{1: "Beginnings · agency · focus", 2: "Partnership · patience · receptivity", 3: "Expression · delight · creativity", 4: "Structure · effort · reliability", 5: "Change · freedom · adaptability", 6: "Care · duty · harmony", 7: "Study · contemplation · insight", 8: "Power · resources · accountability", 9: "Completion · compassion · release"}
	prompts := map[int]string{1: "What deserves a clear first step?", 2: "Where could listening change the relationship?", 3: "What truth wants a more creative form?", 4: "Which foundation needs patient attention?", 5: "What can change without losing your center?", 6: "What does responsible care look like today?", 7: "Which question is more useful than a quick answer?", 8: "How can influence be used with integrity?", 9: "What is complete enough to bless and release?"}
	return NumerologyLens{
		Score:  sum,
		Title:  titles[sum],
		Theme:  themes[sum],
		Prompt: prompts[sum],
		Method: "Pythagorean universal-day reduction: add every digit in YYYY-MM-DD, then reduce to 1–9.",
		Caveat: "Numerology is presented as a reflective symbolic practice, not a scientific measurement or forecast.",
	}
}

func digitSum(value int) int {
	if value < 0 {
		value = -value
	}
	total := 0
	for value > 0 {
		total += value % 10
		value /= 10
	}
	return total
}
