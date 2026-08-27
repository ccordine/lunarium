package calendar

import "time"

func islamicObservances(date time.Time) []Observance {
	i := jdToIslamic(gregorianToJD(date))
	var events []Observance

	if date.Weekday() == time.Friday {
		event := islamicEvent("Jumu'ah", []string{"Muslim communities"}, "Weekly congregational prayer", "Friday's congregational midday prayer gathers the community for worship and a sermon.", "Jumu'ah gives the week a public rhythm of remembrance, learning, and mutual presence.", []string{"Ghusl and clean dress", "Khutbah", "Congregational Dhuhr prayer", "Surah al-Kahf in widespread devotional practice"}, []string{"Qur'an 62:9–10"})
		event.StartsAtSunset = false
		event.DateNote = "The congregational prayer occurs around local Dhuhr; consult a local mosque for its iqamah time."
		events = append(events, singleOccurrence(event, date))
	}

	add := func(event Observance) {
		event.StartsAtSunset = true
		if event.DateNote == "" {
			event.DateNote = islamicDateCaveat()
		} else {
			event.DateNote += " " + islamicDateCaveat()
		}
		events = append(events, singleOccurrence(event, date))
	}

	switch {
	case i.Month == 1 && i.Day == 1:
		event := islamicEvent("Islamic New Year · 1 Muharram", []string{"Muslim communities"}, "Calendar observance", "The first day of Muharram begins a new Hijri year, counted from the Prophet Muhammad's migration to Medina.", "The calendar locates communal time around migration, covenant, and the formation of a worshiping society.", []string{"Reflection", "Prayer", "Community teaching; customs vary"}, []string{"Qur'an 9:20", "Qur'an 2:218"})
		event.Rank = "Sacred month"
		add(event)
	case i.Month == 1 && i.Day == 9:
		event := islamicEvent("Tasu'a", []string{"Twelver Shia and other Shia communities"}, "Commemoration", "The ninth of Muharram is a day of mourning and preparation before Ashura.", "Devotion centers loyalty, moral courage, and the suffering of Imam Husayn and his companions at Karbala.", []string{"Mourning gatherings", "Recitation", "Fasting in some communities"}, nil)
		event.DateNote = "Practices differ among Shia traditions."
		add(event)
	case i.Month == 1 && i.Day == 10:
		event := islamicEvent("Ashura", []string{"Sunni Muslims", "Shia Muslims"}, "Fast and commemoration", "The tenth of Muharram carries distinct but overlapping meanings across Islam.", "Sunni traditions emphasize the prophetic fast and deliverance associated with Moses; Shia traditions mourn the martyrdom of Imam Husayn at Karbala.", []string{"Voluntary fasting, especially in Sunni practice", "Mourning assemblies and processions in Shia practice", "Charity and remembrance"}, []string{"Prophetic hadith on fasting Ashura", "Accounts of Karbala"})
		event.Rank = "Major Shia mourning day / recommended Sunni fast"
		event.DateNote = "Theologies and devotional practices differ substantially across Muslim communities."
		add(event)
	case i.Month == 2 && i.Day == 20:
		event := islamicEvent("Arba'een", []string{"Twelver Shia and other Shia communities"}, "Commemoration", "Forty days after Ashura, pilgrims and communities mourn Imam Husayn and the martyrs of Karbala.", "The immense pilgrimage expresses fidelity, hospitality, resistance to injustice, and communal memory.", []string{"Pilgrimage to Karbala", "Mourning gatherings", "Hospitality and service", "Ziyarat Arba'een"}, nil)
		event.DateNote = "Most prominent in Twelver Shia practice."
		add(event)
	case i.Month == 3 && i.Day == 12:
		event := islamicEvent("Mawlid al-Nabi · commonly observed date", []string{"Many Sunni and Sufi communities", "Some Shia communities on a different date"}, "Commemoration", "Many Muslims commemorate the birth of the Prophet Muhammad with praise, teaching, and generosity.", "The gathering expresses gratitude for prophetic guidance and love of the Messenger.", []string{"Recitation of Qur'an", "Salawat and poetry", "Teaching about the Prophet's life", "Charity and shared food"}, []string{"Qur'an 21:107", "Qur'an 33:56"})
		event.DateNote = "Many Sunni communities use 12 Rabi al-Awwal; many Shia communities use the 17th. Some Muslims do not observe Mawlid."
		add(event)
	case i.Month == 3 && i.Day == 17:
		event := islamicEvent("Mawlid al-Nabi · Shia observed date", []string{"Many Twelver Shia communities"}, "Commemoration", "Many Shia Muslims commemorate the births of Prophet Muhammad and Imam Ja'far al-Sadiq.", "Praise, learning, and communal joy honor prophetic and imamic guidance.", []string{"Celebratory gathering", "Qur'an and salawat", "Teaching", "Shared food"}, []string{"Qur'an 21:107", "Qur'an 33:56"})
		event.DateNote = "Many Twelver Shia communities use 17 Rabi al-Awwal; practice is not universal."
		add(event)
	case i.Month == 7 && i.Day == 27:
		event := islamicEvent("Isra and Mi'raj · traditional observance", []string{"Many Muslim communities"}, "Commemoration", "The Night Journey and Ascension remember the Prophet's journey to Jerusalem and through the heavens.", "The narratives link prophecy, sacred geography, divine encounter, and the gift of daily prayer.", []string{"Night prayer", "Qur'an recitation", "Storytelling and teaching"}, []string{"Qur'an 17:1", "Qur'an 53:1–18"})
		event.DateNote = "The 27 Rajab date is traditional rather than historically certain, and observance varies."
		add(event)
	case i.Month == 8 && i.Day == 15:
		event := islamicEvent("Mid-Sha'ban · Laylat al-Bara'ah", []string{"Many Sunni, Shia, and South Asian Muslim communities"}, "Night observance", "Many communities devote the middle night of Sha'ban to prayer, mercy, and remembrance of the dead.", "The night invites repentance and preparation for Ramadan, though its status is debated.", []string{"Voluntary night prayer", "Qur'an recitation", "Visiting graves in some cultures", "Fasting in some communities"}, nil)
		event.DateNote = "Names, practices, and assessments of the supporting hadith vary widely."
		add(event)
	case i.Month == 9:
		start := date.AddDate(0, 0, -(i.Day - 1))
		duration := islamicMonthDays(i.Year, 9)
		event := islamicEvent("Ramadan", []string{"Muslim communities"}, "Sacred month", "Ramadan is the month of fasting, Qur'an, prayer, generosity, and intensified awareness of God.", "Fasting from dawn to sunset trains intention and solidarity while the month celebrates the Qur'an's revelation.", []string{"Fasting from Fajr to sunset for those obligated and able", "Five daily prayers and Tarawih in many Sunni communities", "Qur'an recitation", "Charity", "Iftar and suhur"}, []string{"Qur'an 2:183–187", "Qur'an 97"})
		event.Rank = "Month of obligatory fasting"
		event.DateNote = islamicDateCaveat() + " Exemptions and accommodations protect health, travel, pregnancy, age, and other circumstances under Islamic law."
		event.StartsAtSunset = true
		events = append(events, spanOccurrence(event, date, start, duration))
		if i.Day == 17 {
			nuzul := islamicEvent("Nuzul al-Qur'an · regional observance", []string{"Many Southeast Asian Muslim communities"}, "Commemoration", "A regional observance marks the beginning of the Qur'an's revelation.", "Revelation is honored through recitation, study, and renewed ethical attention.", []string{"Qur'an recitation", "Teaching", "Night prayer"}, []string{"Qur'an 96:1–5"})
			nuzul.DateNote = "The 17 Ramadan civic/religious observance is especially prominent in Southeast Asia; traditions about the first revelation and Laylat al-Qadr are distinguished in different ways."
			add(nuzul)
		}
		if i.Day == 19 || i.Day == 21 || i.Day == 23 {
			qadr := islamicEvent("Laylat al-Qadr · Shia devotional night", []string{"Many Twelver Shia communities"}, "Night observance", "Shia communities especially seek the Night of Power on the 19th, 21st, and 23rd nights of Ramadan.", "Night prayer joins revelation, destiny, repentance, and remembrance of Imam Ali's martyrdom in this period.", []string{"Night vigil", "Qur'an recitation", "Dua", "Acts of repentance"}, []string{"Qur'an 97"})
			qadr.DateNote = "The exact Night of Power is unknown; multiple Ramadan nights are observed."
			add(qadr)
		}
		if i.Day == 27 {
			qadr := islamicEvent("Laylat al-Qadr · widely observed night", []string{"Many Sunni and other Muslim communities"}, "Night observance", "Many Muslims especially observe the twenty-seventh night while seeking the Night of Power throughout Ramadan's last ten nights.", "A night described as better than a thousand months intensifies prayer, repentance, and hope.", []string{"Night prayer", "Qur'an recitation", "Dua", "Retreat in the mosque in some communities"}, []string{"Qur'an 97"})
			qadr.DateNote = "The exact night is intentionally uncertain; prophetic tradition encourages seeking it in the odd nights of Ramadan's final ten."
			add(qadr)
		}
	case i.Month == 10 && i.Day >= 1 && i.Day <= 3:
		start := date.AddDate(0, 0, -(i.Day - 1))
		event := islamicEvent("Eid al-Fitr", []string{"Muslim communities"}, "Festival", "The Festival of Breaking the Fast begins Shawwal with communal prayer, gratitude, and generosity.", "Joy follows disciplined worship, and zakat al-fitr ensures the wider community can share the feast.", []string{"Zakat al-fitr before prayer", "Eid prayer", "Festive dress and food", "Family and community visits"}, []string{"Qur'an 2:185"})
		event.Rank = "One of Islam's two principal festivals"
		event.StartsAtSunset = true
		event.DateNote = islamicDateCaveat() + " Religious rites center on the first day; cultural celebration often continues for up to three days."
		events = append(events, spanOccurrence(event, date, start, 3))
	case i.Month == 12 && i.Day >= 1 && i.Day <= 10:
		start := date.AddDate(0, 0, -(i.Day - 1))
		event := islamicEvent("First ten days of Dhu al-Hijjah", []string{"Muslim communities"}, "Sacred season", "The opening days of Dhu al-Hijjah are a highly honored season culminating in Hajj and Eid al-Adha.", "Worship, remembrance, charity, pilgrimage, and sacrifice converge in a shared orientation toward God.", []string{"Dhikr", "Charity", "Voluntary fasting outside Hajj", "Hajj rites for pilgrims"}, []string{"Qur'an 22:27–37", "Qur'an 89:1–5"})
		event.StartsAtSunset = true
		event.DateNote = islamicDateCaveat()
		events = append(events, spanOccurrence(event, date, start, 10))
		if i.Day == 8 {
			hajj := islamicEvent("Hajj begins · Day of Tarwiyah", []string{"Muslim pilgrims; Muslim communities worldwide accompany in prayer"}, "Pilgrimage", "Pilgrims enter the concentrated rites of Hajj and travel to Mina.", "The pilgrimage gathers Muslims across language, nationality, and status in shared rites.", []string{"Ihram", "Talbiyah", "Travel to Mina", "Prayer"}, []string{"Qur'an 2:196–203", "Qur'an 22:27–37"})
			hajj.DateNote = "Hajj logistics and ritual timing follow Saudi authorities and qualified guides."
			add(hajj)
		}
		if i.Day == 9 {
			arafah := islamicEvent("Day of Arafah", []string{"Muslim pilgrims", "Muslim communities worldwide"}, "Pilgrimage and fast", "Standing at Arafat is the heart of Hajj; non-pilgrims are widely encouraged to fast.", "Supplication at Arafat gathers repentance, mercy, equality, and remembrance of final accountability.", []string{"Wuquf at Arafat for pilgrims", "Supplication", "Recommended fast for non-pilgrims"}, nil)
			arafah.Rank = "Central day of Hajj"
			add(arafah)
		}
		if i.Day == 10 {
			events = append(events, eidAlAdhaOccurrence(date, i))
		}
	case i.Month == 12 && i.Day >= 10 && i.Day <= 13:
		events = append(events, eidAlAdhaOccurrence(date, i))
	case i.Month == 12 && i.Day == 18:
		event := islamicEvent("Eid al-Ghadir", []string{"Twelver Shia and other Shia communities"}, "Festival", "Shia Muslims celebrate the Prophet's declaration concerning Ali at Ghadir Khumm.", "The day centers devotion to the Prophet's family, spiritual authority, covenant, and communal care.", []string{"Festive prayer and gathering", "Renewal of bonds", "Charity", "Teaching about Ghadir"}, nil)
		event.DateNote = "Interpretation of the Ghadir event differs between Sunni and Shia traditions."
		add(event)
	case i.Month == 12 && i.Day == 24:
		event := islamicEvent("Eid al-Mubahala", []string{"Many Shia communities"}, "Commemoration", "The day recalls the Prophet's encounter with Christians of Najran and the Qur'anic event of mutual invocation.", "Shia interpretation emphasizes the spiritual station of the Prophet's household; the story also invites reflection on interreligious encounter.", []string{"Prayer", "Teaching", "Charity"}, []string{"Qur'an 3:61"})
		event.DateNote = "Observed especially in Shia communities."
		add(event)
	}
	return events
}

func islamicEvent(name string, communities []string, category, summary, meaning string, practices, scripture []string) Observance {
	return baseObservance(name, Islam, communities, category, summary, meaning, practices, scripture, "Tabular Hijri calendar methodology", islamicCalendarSource)
}

func islamicDateCaveat() string {
	return "Shown by the tabular Hijri calendar; local crescent sighting or religious authority may move the observance by one day."
}

func eidAlAdhaOccurrence(date time.Time, i islamicDate) Observance {
	start := date.AddDate(0, 0, -(i.Day - 10))
	event := islamicEvent("Eid al-Adha", []string{"Muslim communities"}, "Festival", "The Festival of Sacrifice coincides with the climax of Hajj and remembers Abraham's willingness to surrender to God.", "Sacrifice is expressed through worship, generosity, and distribution of food rather than divine need for blood.", []string{"Eid prayer", "Qurbani or udhiyah for those able", "Sharing meat with family and people in need", "Hajj rites for pilgrims"}, []string{"Qur'an 22:27–37", "Qur'an 37:99–111"})
	event.Rank = "One of Islam's two principal festivals"
	event.StartsAtSunset = true
	event.DateNote = islamicDateCaveat() + " Sacrifice is performed during the festival days according to legal school and local guidance."
	return spanOccurrence(event, date, start, 4)
}
