package calendar

import "time"

func jewishObservances(date time.Time) []Observance {
	h := jdToHebrew(gregorianToJD(date))
	var events []Observance

	if date.Weekday() == time.Saturday {
		event := baseObservance("Shabbat", Judaism, []string{"Jewish communities"}, "Weekly sacred time", "The seventh day is set apart for rest, delight, prayer, and release from ordinary labor.", "Shabbat remembers both creation and liberation from Egypt, making holy time a recurring practice of freedom.", []string{"Candle lighting before Friday sunset", "Kiddush and shared meals", "Synagogue prayer and Torah reading", "Havdalah after nightfall Saturday"}, []string{"Genesis 2:1–3", "Exodus 20:8–11", "Deuteronomy 5:12–15"}, "Hebcal Jewish calendar", hebcalSource)
		markJewishLayer(&event, "Torah", "Still observed", "Received weekly cycle")
		event.StartsAtSunset = true
		event.DateNote = "Begins Friday at local sunset and ends after nightfall Saturday; exact times depend on location and custom."
		events = append(events, singleOccurrence(event, date))
	}

	if h.Day == 1 || h.Day == 30 {
		month := hebrewMonthName(h.Year, h.Month)
		if h.Day == 30 {
			month = hebrewMonthName(h.Year, nextHebrewMonth(h.Year, h.Month))
		}
		event := baseObservance("Rosh Chodesh "+month, Judaism, []string{"Jewish communities; practice differs by movement and gender custom"}, "Monthly observance", "The new Hebrew month is welcomed at the lunar renewal.", "Rosh Chodesh makes recurring time visible through additional prayer, blessing, and in some communities women's gathering.", []string{"Hallel or additional psalms", "Musaf in traditional liturgy", "Festive food or study"}, []string{"Numbers 28:11–15"}, "Hebcal Jewish calendar", hebcalSource)
		markJewishLayer(&event, "Torah and rabbinic calendar", "Still observed", "Received fixed-calendar date")
		event.StartsAtSunset = true
		event.DateNote = "A Hebrew month may have one or two Rosh Chodesh days depending on the preceding month's length."
		events = append(events, singleOccurrence(event, date))
	}

	add := func(event Observance) {
		event.StartsAtSunset = true
		events = append(events, singleOccurrence(event, date))
	}
	addDaytimeFast := func(event Observance) {
		event.StartsAtSunset = false
		events = append(events, singleOccurrence(event, date))
	}

	tomorrow := date.AddDate(0, 0, 1)
	tomorrowHebrew := jdToHebrew(gregorianToJD(tomorrow))
	if observedYomHaAtzmaut(tomorrow, tomorrowHebrew) {
		event := jewishEvent("Yom HaZikaron · Israeli Memorial Day", "Modern commemoration", "Israel's memorial day honors fallen members of the security services and victims of terrorism immediately before Independence Day.", "The sequence from mourning into celebration deliberately holds public grief beside national life; interpretation and participation vary.", []string{"Sirens and public silence in Israel", "Memorial ceremonies", "Prayer and remembrance"}, nil)
		event.DateNote = "A modern Israeli civil and communal observance; the date follows the shifted Yom HaAtzmaut schedule."
		add(event)
	}

	switch {
	case h.Month == 7 && h.Day == 1:
		event := jewishEvent("Rosh Hashanah", "High Holy Day", "The Jewish new year begins a season of judgment, remembrance, and return.", "The shofar awakens moral attention while divine sovereignty and the creation of the world are remembered.", []string{"Shofar", "Synagogue prayer", "Festive meals", "Tashlich in many communities"}, []string{"Genesis 21", "1 Samuel 1–2"})
		markJewishLayer(&event, "Torah day of teruah and rabbinic Rosh Hashanah", "Still observed", "Torah appoints Tishrei 1; received practice observes two days")
		event.Rank, event.DurationDays = "Yom Tov", 2
		event.DateNote = "Observed for two days in most communities; begins the previous evening."
		event.HistoricalNote = "Mishnah Rosh Hashanah 1:1 also makes Tishrei 1 the cutoff for years, Shemitah, Jubilee, planting, and vegetables."
		events = append(events, spanOccurrence(event, date, date, 2))
	case h.Month == 7 && h.Day == 2:
		start := date.AddDate(0, 0, -1)
		event := jewishEvent("Rosh Hashanah", "High Holy Day", "The second day continues the Jewish new year and season of return.", "Prayer turns toward memory, accountability, mercy, and the renewal of life.", []string{"Shofar unless Shabbat", "Synagogue prayer", "Festive meals"}, []string{"Genesis 22", "Jeremiah 31:1–19"})
		markJewishLayer(&event, "Rabbinic extension of the Torah day of teruah", "Still observed", "Received two-day observance")
		event.Rank = "Yom Tov"
		events = append(events, spanOccurrence(event, date, start, 2))
	case observedHebrewFast(date, h, 7, 3, false):
		event := jewishEvent("Fast of Gedaliah", "Fast", "A minor fast mourns the assassination of Gedaliah and the collapse of the remaining Judean community after the First Temple's destruction.", "The fast reflects on political violence, communal fracture, and the consequences of hatred.", []string{"Daytime fasting in traditional practice", "Selichot", "Torah reading"}, []string{"Jeremiah 40–41"})
		event.DateNote = "Postponed to Sunday when 3 Tishrei falls on Shabbat."
		addDaytimeFast(event)
	case h.Month == 7 && h.Day == 10:
		event := jewishEvent("Yom Kippur", "High Holy Day", "The Day of Atonement is devoted to fasting, confession, prayer, and reconciliation.", "Human beings seek forgiveness from God while harms between people require direct repair.", []string{"Approximately 25-hour fast", "Kol Nidrei", "Five prayer services", "Confession", "Ne'ilah"}, []string{"Leviticus 16", "Isaiah 57:14–58:14", "Jonah"})
		markJewishLayer(&event, "Torah", "Still observed; Temple rite inactive", "Received fixed-calendar date")
		event.Rank = "Most solemn fast"
		event.DateNote = "Begins before sunset and ends after nightfall the following day. Health and safety exceptions are integral to Jewish law."
		event.HistoricalNote = "In an operative Jubilee year, the Torah also appoints a nationwide liberty-proclaiming shofar on Yom Kippur; the Jubilee cycle is no longer operative and its historical year count is disputed."
		add(event)
	case h.Month == 7 && h.Day >= 15 && h.Day <= 21:
		start := date.AddDate(0, 0, -(h.Day - 15))
		event := jewishEvent("Sukkot", "Pilgrimage festival", "The seven-day festival of booths joins harvest gratitude with memory of wilderness vulnerability.", "Dwelling in a fragile sukkah makes impermanence, hospitality, and divine shelter tangible.", []string{"Dwelling and eating in a sukkah", "Lulav and etrog", "Hallel", "Welcoming guests"}, []string{"Leviticus 23:33–43", "Ecclesiastes"})
		markJewishLayer(&event, "Torah", "Still observed; Temple pilgrimage inactive", "Received fixed-calendar date")
		event.Rank = "Yom Tov / Chol HaMoed"
		events = append(events, spanOccurrence(event, date, start, 7))
		if h.Day == 21 {
			hoshana := jewishEvent("Hoshana Rabbah", "Festival day", "The seventh day of Sukkot culminates its prayers for help, rain, and blessing.", "Circling and willow rituals intensify the festival's plea for life and renewal.", []string{"Seven hakafot", "Willow beating custom", "Hallel"}, []string{"Psalm 118"})
			add(hoshana)
		}
	case h.Month == 7 && h.Day == 22:
		event := jewishEvent("Shemini Atzeret", "Festival", "An eighth-day sacred assembly follows Sukkot as a distinct festival.", "The day lingers in divine-human closeness and includes prayer for rain.", []string{"Yom Tov meals", "Geshem prayer for rain", "Yizkor in many communities"}, []string{"Deuteronomy 14:22–16:17"})
		markJewishLayer(&event, "Torah", "Still observed; Temple rite inactive", "Received fixed-calendar date")
		event.Rank = "Yom Tov"
		add(event)
	case h.Month == 7 && h.Day == 23:
		event := jewishEvent("Simchat Torah · Diaspora", "Festival", "Diaspora communities complete and immediately restart the annual Torah-reading cycle.", "Completion becomes beginning; Torah is celebrated through embodied joy and communal participation.", []string{"Dancing with Torah scrolls", "Hakafot", "Reading Deuteronomy's end and Genesis' beginning"}, []string{"Deuteronomy 33:1–34:12", "Genesis 1:1–2:3"})
		event.Rank = "Yom Tov"
		event.DateNote = "In Israel, Simchat Torah is observed together with Shemini Atzeret on 22 Tishrei."
		add(event)
	case hanukkahDay(h) > 0:
		index := hanukkahDay(h)
		start := date.AddDate(0, 0, -(index - 1))
		event := jewishEvent("Hanukkah", "Festival", "The eight-day Festival of Lights recalls the rededication of the Jerusalem Temple after the Maccabean revolt.", "Increasing light becomes a practice of dedication, resilience, and public memory.", []string{"Lighting the hanukkiah", "Blessings and songs", "Oil foods", "Games and gifts in many communities"}, []string{"1 Maccabees 4:36–59", "Zechariah 4:1–7"})
		markJewishLayer(&event, "Megillat Ta'anit and Bavli Shabbat 21b", "Still observed", "Received fixed-calendar date")
		event.SourceName, event.SourceURL = "Bavli Shabbat 21b", "https://www.sefaria.org/Shabbat.21b"
		event.Rank = "Minor festival"
		events = append(events, spanOccurrence(event, date, start, 8))
	case h.Month == 10 && h.Day == 10:
		event := jewishEvent("Tenth of Tevet · Asarah B'Tevet", "Fast", "A minor fast marks the beginning of the Babylonian siege of Jerusalem.", "The day remembers how destruction develops over time and invites attention before crisis becomes irreversible.", []string{"Daytime fasting in traditional practice", "Selichot", "Torah reading"}, []string{"2 Kings 25:1–4", "Zechariah 8:19"})
		addDaytimeFast(event)
	case h.Month == 11 && h.Day == 15:
		event := jewishEvent("Tu BiShvat", "Festival", "The new year for trees has become a celebration of fruit, land, ecological responsibility, and mystical symbolism.", "Kabbalists in Safed developed a fruit-and-wine seder contemplating four worlds and the flow of divine vitality.", []string{"Eating fruits of the Land of Israel", "Tu BiShvat seder", "Ecological learning and planting"}, []string{"Deuteronomy 8:7–10"})
		markJewishLayer(&event, "Mishnah Rosh Hashanah 1:1 and later custom", "Still observed as a minor day", "Received Beit Hillel date")
		event.Rank = "Minor festival"
		add(event)
	case isFastOfEsther(date, h):
		event := jewishEvent("Fast of Esther", "Fast", "A minor fast precedes Purim and recalls communal fasting in Esther's story.", "Preparation places courage, solidarity, and dependence before celebration.", []string{"Daytime fasting in traditional practice", "Charity", "Afternoon Torah reading"}, []string{"Esther 4:15–17"})
		event.DateNote = "Usually 13 Adar; advanced to Thursday when the date would conflict with Shabbat."
		addDaytimeFast(event)
	case purimMonth(h) && h.Day == 14:
		event := jewishEvent("Purim", "Festival", "Purim celebrates the deliverance narrated in the Book of Esther.", "Costume, reversal, satire, generosity, and public story confront fear and hidden power.", []string{"Reading the Megillah", "Gifts of food", "Gifts to people in need", "Festive meal", "Costumes"}, []string{"Book of Esther"})
		markJewishLayer(&event, "Book of Esther, Mishnah, and Talmud", "Still observed", "Adar II in leap years")
		event.Rank = "Minor festival"
		add(event)
	case purimMonth(h) && h.Day == 15:
		event := jewishEvent("Shushan Purim", "Festival", "Walled cities traditionally celebrate Purim one day later, with Jerusalem the principal example today.", "The date preserves the distinct timeline of deliverance in the city of Shushan.", []string{"Megillah and Purim practices in Jerusalem", "Festive meal"}, []string{"Esther 9:18–19"})
		event.DateNote = "Observed as the main Purim date in Jerusalem and certain ancient walled cities."
		add(event)
	case h.Month == 1 && h.Day >= 15 && h.Day <= 22:
		start := date.AddDate(0, 0, -(h.Day - 15))
		event := jewishEvent("Passover · Pesach", "Pilgrimage festival", "Passover remembers liberation from slavery in Egypt through story, ritual food, and removal of leaven.", "Each generation is invited to experience liberation as personal and unfinished, joining memory to responsibility.", []string{"Seder", "Eating matzah", "Retelling the Exodus", "Avoiding chametz", "Hallel"}, []string{"Exodus 12–15", "Song of Songs"})
		markJewishLayer(&event, "Torah; eighth day is a rabbinic diaspora extension", "Still observed; Temple pilgrimage inactive", "Seven Torah days in Israel; eight received days in the diaspora")
		event.Rank = "Yom Tov / Chol HaMoed"
		event.DateNote = "Eight days in the diaspora and seven in Israel; the first night begins at sunset."
		events = append(events, spanOccurrence(event, date, start, 8))
	case observedYomHaShoah(date, h):
		event := jewishEvent("Yom HaShoah", "Modern commemoration", "Holocaust Remembrance Day honors the six million Jews murdered in the Shoah and remembers resistance and survival.", "Memory becomes testimony, mourning, and a demand to resist antisemitism and dehumanization.", []string{"Memorial ceremonies", "Testimony and study", "Candle lighting", "Silence"}, nil)
		event.DateNote = "The Israeli civil observance shifts when 27 Nisan adjoins Shabbat."
		add(event)
	case observedYomHaAtzmaut(date, h):
		event := jewishEvent("Yom HaAtzmaut · Israeli Independence Day", "Modern observance", "The modern State of Israel's declaration of independence is celebrated, with religious meaning varying widely among Jewish communities.", "Celebration, gratitude, peoplehood, politics, and critique coexist differently across communities.", []string{"Public celebration", "Hallel in many religious-Zionist communities", "Study and gathering"}, nil)
		event.DateNote = "The civil date shifts to avoid conflict with Shabbat; observance and interpretation vary."
		add(event)
	case h.Month == 2 && h.Day == 14:
		event := jewishEvent("Pesach Sheni", "Minor observance", "The biblical second Passover offered another opportunity to bring the sacrifice when the first date could not be kept.", "Later interpretation emphasizes that return and a second chance remain possible.", []string{"Eating a piece of matzah in some communities", "Omitting penitential prayers"}, []string{"Numbers 9:1–14"})
		markJewishLayer(&event, "Torah", "Temple rite inactive; date still marked", "Received fixed-calendar date")
		add(event)
	case h.Month == 2 && h.Day == 18:
		event := jewishEvent("Lag BaOmer", "Minor festival", "The thirty-third day of the Omer interrupts semi-mourning customs and carries rabbinic and mystical associations.", "The day is linked to Rabbi Shimon bar Yochai, hidden Torah, resilience, and joy.", []string{"Bonfires", "Outdoor gatherings", "Weddings", "Pilgrimage to Meron in some communities"}, nil)
		event.DateNote = "Customs and historical explanations vary substantially."
		add(event)
	case h.Month == 2 && h.Day == 28:
		event := jewishEvent("Yom Yerushalayim · Jerusalem Day", "Modern observance", "A modern Israeli day marking the 1967 capture of East Jerusalem and access to the Old City.", "Religious and political interpretations differ sharply; the day can carry celebration, grief, or both.", []string{"Public ceremonies", "Hallel in many religious-Zionist communities", "Study and reflection"}, nil)
		event.DateNote = "Observed chiefly in Israel and religious-Zionist communities; perspectives vary."
		add(event)
	case h.Month == 3 && (h.Day == 6 || h.Day == 7):
		start := date.AddDate(0, 0, -(h.Day - 6))
		event := jewishEvent("Shavuot", "Pilgrimage festival", "Shavuot marks the wheat harvest and rabbinic tradition associates it with the giving of Torah at Sinai.", "Revelation is approached as renewed covenant, learning, and responsibility.", []string{"All-night Torah study", "Reading Ruth", "Dairy foods in many communities", "Yizkor"}, []string{"Exodus 19–20", "Book of Ruth"})
		markJewishLayer(&event, "Torah fifty-day count; fixed date and Sinai meaning are rabbinic", "Still observed; Temple first-fruit rites inactive", "Sivan 6 in the fixed calendar; Sivan 7 is the diaspora extension")
		event.Rank = "Yom Tov"
		event.DateNote = "One day in Israel and two in the diaspora."
		event.HistoricalNote = "The Torah fixes Shavuot as the fiftieth day of the Omer rather than naming Sivan 6; before fixed month lengths, rabbinic sources allow Sivan 5, 6, or 7."
		events = append(events, spanOccurrence(event, date, start, 2))
	case observedHebrewFast(date, h, 4, 17, false):
		event := jewishEvent("Fast of the Seventeenth of Tammuz", "Fast", "A minor fast begins the Three Weeks of mourning leading to Tisha B'Av.", "Breached walls become an image for communal vulnerability and the early signs of destruction.", []string{"Daytime fasting in traditional practice", "Selichot", "Beginning Three Weeks customs"}, []string{"Zechariah 8:19"})
		event.DateNote = "Postponed to Sunday when 17 Tammuz falls on Shabbat."
		addDaytimeFast(event)
	case observedHebrewFast(date, h, 5, 9, false):
		event := jewishEvent("Tisha B'Av", "Major fast", "The Ninth of Av mourns the destruction of both Temples and other catastrophes in Jewish memory.", "Collective lament refuses to rush grief and asks how hatred, violence, and exile are remembered.", []string{"Approximately 25-hour fast", "Reading Lamentations", "Sitting low", "Kinot lament poems"}, []string{"Book of Lamentations"})
		event.DateNote = "Postponed to Sunday when 9 Av falls on Shabbat. Health and safety exceptions apply."
		add(event)
	case h.Month == 5 && h.Day == 15:
		event := jewishEvent("Tu B'Av", "Minor festival", "An ancient day of matchmaking and reconciliation has become a celebration of love.", "After Tisha B'Av, the date turns gently toward consolation, relationship, and renewal.", []string{"Celebration of love", "Music and gathering", "Omitting penitential prayers"}, nil)
		markJewishLayer(&event, "Mishnah Ta'anit 4:8 and Bavli Ta'anit 30b–31a", "Still observed as a minor day; Temple customs inactive", "Received fixed-calendar date")
		add(event)
	}
	events = append(events, torahCalendarLayers(date, h)...)
	events = append(events, leapYearMinorPurims(date, h)...)
	events = append(events, purimMeshulash(date, h)...)
	events = append(events, isruChagObservances(date, h)...)
	events = append(events, templeEraRabbinicLayers(date, h)...)
	events = append(events, fastOfFirstborn(date, h)...)
	events = append(events, historicalJewishObservances(date, h)...)
	return events
}

func jewishEvent(name, category, summary, meaning string, practices, scripture []string) Observance {
	event := baseObservance(name, Judaism, []string{"Jewish communities"}, category, summary, meaning, practices, scripture, "Hebcal Jewish calendar", hebcalSource)
	markJewishLayer(&event, "Rabbinic and later Jewish tradition", "Still observed; practice varies", "Received fixed-calendar date")
	return event
}

func markJewishLayer(event *Observance, origin, status, certainty string) {
	event.Origin = origin
	event.ObservanceStatus = status
	event.DateCertainty = certainty
}

func nextHebrewMonth(year, month int) int {
	if month == hebrewMonthsInYear(year) {
		return 1
	}
	if month == 6 {
		return 7
	}
	return month + 1
}

func hanukkahDay(h hebrewDate) int {
	if h.Month == 9 && h.Day >= 25 {
		return h.Day - 24
	}
	if h.Month == 10 {
		kislevDays := hebrewMonthDays(h.Year, 9)
		index := kislevDays - 24 + h.Day
		if index <= 8 {
			return index
		}
	}
	return 0
}

func purimMonth(h hebrewDate) bool {
	if hebrewLeap(h.Year) {
		return h.Month == 13
	}
	return h.Month == 12
}

func observedHebrewFast(date time.Time, h hebrewDate, month, day int, advance bool) bool {
	if h.Month != month || h.Day < day-2 || h.Day > day+1 {
		return false
	}
	base := date.AddDate(0, 0, day-h.Day)
	observed := base
	if base.Weekday() == time.Saturday {
		if advance {
			observed = base.AddDate(0, 0, -2)
		} else {
			observed = base.AddDate(0, 0, 1)
		}
	}
	return sameDay(date, observed)
}

func isFastOfEsther(date time.Time, h hebrewDate) bool {
	month := 12
	if hebrewLeap(h.Year) {
		month = 13
	}
	return observedHebrewFast(date, h, month, 13, true)
}

func observedYomHaShoah(date time.Time, h hebrewDate) bool {
	if h.Month != 1 || h.Day < 26 || h.Day > 28 {
		return false
	}
	base := date.AddDate(0, 0, 27-h.Day)
	observed := base
	if base.Weekday() == time.Friday {
		observed = base.AddDate(0, 0, -1)
	} else if base.Weekday() == time.Sunday {
		observed = base.AddDate(0, 0, 1)
	}
	return sameDay(date, observed)
}

func observedYomHaAtzmaut(date time.Time, h hebrewDate) bool {
	if h.Month != 2 || h.Day < 2 || h.Day > 6 {
		return false
	}
	base := date.AddDate(0, 0, 5-h.Day)
	observed := base
	switch base.Weekday() {
	case time.Friday:
		observed = base.AddDate(0, 0, -1)
	case time.Saturday:
		observed = base.AddDate(0, 0, -2)
	case time.Monday:
		observed = base.AddDate(0, 0, 1)
	}
	return sameDay(date, observed)
}
