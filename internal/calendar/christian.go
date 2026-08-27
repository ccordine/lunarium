package calendar

import "time"

type christianFixed struct {
	month       time.Month
	day         int
	name        string
	communities []string
	category    string
	rank        string
	color       string
	summary     string
	meaning     string
	practices   []string
	scripture   []string
	source      string
}

var christianFixedDates = []christianFixed{
	{time.January, 1, "Mary, the Holy Mother of God", []string{"Catholic"}, "Solemnity", "Solemnity", "white", "Catholics begin the civil year honoring Mary under her ancient title Theotokos, Mother of God.", "The feast holds together the full divinity and humanity of Christ and Mary's faithful consent.", []string{"Mass", "Prayers for peace", "Marian devotion"}, []string{"Numbers 6:22–27", "Luke 2:16–21"}, usccbSource},
	{time.January, 1, "Feast of the Holy Name / Circumcision of Christ", []string{"Anglican", "Lutheran", "Orthodox"}, "Feast", "Feast", "white", "Several churches remember Jesus' naming and circumcision on the eighth day after the Nativity.", "The observance locates Jesus within the covenant of Israel and celebrates the name of Jesus.", []string{"Divine liturgy or Eucharist", "Prayer in the Holy Name"}, []string{"Genesis 17:9–14", "Luke 2:21"}, usccbSource},
	{time.January, 6, "Epiphany", []string{"Catholic", "Anglican", "Lutheran", "Reformed", "Orthodox on the Gregorian calendar"}, "Feast", "Solemnity", "white", "Epiphany celebrates Christ made manifest to the nations, especially in the visit of the Magi.", "The light revealed in Christ is understood as offered to every people.", []string{"Eucharist", "House blessing", "Chalking the door"}, []string{"Isaiah 60:1–6", "Matthew 2:1–12"}, usccbSource},
	{time.January, 7, "Nativity of Christ · Julian calendar", []string{"Eastern Orthodox using the Julian calendar", "Oriental Orthodox communities"}, "Feast", "Great Feast", "white", "Churches retaining the Julian calendar celebrate the Nativity on January 7 Gregorian through 2099.", "The incarnation proclaims God entering human life in Jesus Christ.", []string{"Divine Liturgy", "Fasting ends", "Family feast"}, []string{"Isaiah 9:2–7", "Matthew 2:1–12"}, ocaSource},
	{time.January, 19, "Theophany · Julian calendar", []string{"Eastern Orthodox using the Julian calendar"}, "Feast", "Great Feast", "white", "Theophany commemorates Christ's baptism and the revelation of the Trinity.", "Water is blessed as a sign of creation renewed through Christ.", []string{"Great Blessing of Waters", "Divine Liturgy"}, []string{"Isaiah 55:1–13", "Matthew 3:13–17"}, ocaSource},
	{time.January, 25, "Conversion of Saint Paul", []string{"Catholic", "Anglican", "Lutheran"}, "Feast", "Feast", "white", "The feast remembers Saul's encounter with the risen Christ and his call as apostle.", "Radical reorientation becomes a witness to grace and vocation.", []string{"Mass or Eucharist", "Prayer for Christian unity"}, []string{"Acts 22:3–16", "Mark 16:15–18"}, usccbSource},
	{time.February, 2, "Presentation of the Lord · Candlemas", []string{"Catholic", "Orthodox", "Anglican", "Lutheran"}, "Feast", "Feast", "white", "Forty days after Christmas, the church remembers Jesus presented in the Temple.", "Simeon recognizes Christ as light for the nations; candles embody that image.", []string{"Blessing of candles", "Procession", "Eucharist"}, []string{"Malachi 3:1–4", "Luke 2:22–40"}, usccbSource},
	{time.March, 19, "Saint Joseph, Spouse of the Blessed Virgin Mary", []string{"Catholic"}, "Solemnity", "Solemnity", "white", "The Catholic Church honors Joseph as Mary's spouse and guardian of Jesus.", "Joseph embodies faithful, protective, and often hidden service.", []string{"Mass", "Works of care for families and workers"}, []string{"2 Samuel 7:4–16", "Matthew 1:16–24"}, usccbSource},
	{time.March, 25, "Annunciation of the Lord", []string{"Catholic", "Orthodox", "Anglican", "Lutheran"}, "Feast", "Solemnity / Great Feast", "white", "Gabriel announces Jesus' conception to Mary nine months before Christmas.", "Mary's consent and the incarnation join divine initiative with human freedom.", []string{"Mass or Divine Liturgy", "Angelus"}, []string{"Isaiah 7:10–14", "Luke 1:26–38"}, usccbSource},
	{time.June, 24, "Nativity of Saint John the Baptist", []string{"Catholic", "Orthodox", "Anglican", "Lutheran"}, "Feast", "Solemnity", "white", "The church celebrates the birth of John, the prophetic forerunner of Jesus.", "John's vocation is to prepare a way and direct attention beyond himself.", []string{"Eucharist", "Evening bonfires in some cultures"}, []string{"Isaiah 49:1–6", "Luke 1:57–66, 80"}, usccbSource},
	{time.June, 29, "Saints Peter and Paul", []string{"Catholic", "Orthodox", "Anglican"}, "Feast", "Solemnity", "red", "Peter and Paul are honored together as foundational apostolic witnesses.", "Different vocations and temperaments are gathered into one mission.", []string{"Mass or Divine Liturgy", "Prayer for the church"}, []string{"Acts 12:1–11", "Matthew 16:13–19"}, usccbSource},
	{time.July, 22, "Saint Mary Magdalene", []string{"Catholic", "Orthodox", "Anglican", "Lutheran"}, "Feast", "Feast", "white", "Mary Magdalene is remembered as a disciple and first witness sent to announce the resurrection.", "Her witness centers faithful presence, liberation, and proclamation.", []string{"Eucharist", "Prayer for faithful witness"}, []string{"Song of Songs 3:1–4", "John 20:1–18"}, usccbSource},
	{time.August, 6, "Transfiguration of the Lord", []string{"Catholic", "Orthodox", "Anglican", "Lutheran"}, "Feast", "Feast / Great Feast", "white", "Jesus is revealed in glory before Peter, James, and John.", "The feast joins contemplation of divine light with the path toward the cross.", []string{"Mass or Divine Liturgy", "Blessing of fruit in some Orthodox churches"}, []string{"Daniel 7:9–14", "Luke 9:28–36"}, usccbSource},
	{time.August, 15, "Assumption / Dormition of Mary", []string{"Catholic", "Orthodox", "Anglican communities"}, "Feast", "Solemnity / Great Feast", "white", "Catholics celebrate Mary's Assumption; Orthodox Christians celebrate her Dormition.", "Distinct formulations share hope that embodied human life is destined for communion with God.", []string{"Mass or Divine Liturgy", "Processions", "Blessing of herbs in some places"}, []string{"Revelation 11:19–12:10", "Luke 1:39–56"}, usccbSource},
	{time.September, 8, "Nativity of the Blessed Virgin Mary", []string{"Catholic", "Orthodox", "Anglican"}, "Feast", "Feast / Great Feast", "white", "The church celebrates the traditional birthday of Mary.", "Her birth is contemplated as part of the long preparation for the incarnation.", []string{"Mass or Divine Liturgy", "Marian hymns"}, []string{"Micah 5:1–4", "Matthew 1:1–16"}, usccbSource},
	{time.September, 14, "Exaltation of the Holy Cross", []string{"Catholic", "Orthodox", "Anglican", "Lutheran"}, "Feast", "Feast", "red", "The cross is honored as the instrument of Christ's self-giving and victory.", "Suffering is not romanticized; divine love is proclaimed as present within and beyond it.", []string{"Veneration of the cross", "Eucharist", "Fasting in Orthodox practice"}, []string{"Numbers 21:4–9", "John 3:13–17"}, usccbSource},
	{time.September, 29, "Holy Archangels Michael, Gabriel, and Raphael", []string{"Catholic", "Anglican", "Lutheran"}, "Feast", "Feast", "white", "The feast honors the named archangels as messengers and servants of God.", "Angelic imagery directs attention to protection, healing, proclamation, and worship.", []string{"Mass or Eucharist", "Prayer for protection and healing"}, []string{"Daniel 7:9–14", "John 1:47–51"}, usccbSource},
	{time.October, 4, "Saint Francis of Assisi", []string{"Catholic", "Anglican", "Lutheran communities"}, "Memorial", "Memorial", "white", "Francis is remembered for evangelical poverty, peacemaking, and kinship with creation.", "His life invites joyful simplicity and solidarity with people at the margins.", []string{"Blessing of animals", "Care for creation", "Service to people in poverty"}, []string{"Galatians 6:14–18", "Matthew 11:25–30"}, usccbSource},
	{time.October, 31, "Reformation Day", []string{"Lutheran", "Reformed", "Many Protestant communities"}, "Commemoration", "Festival", "red", "Many Protestants remember the sixteenth-century Reformation and its call to renew the church.", "The day emphasizes grace, faith, scripture, and continuing reform, while ecumenical observance also names historical wounds.", []string{"Worship", "Study of Reformation history", "Ecumenical prayer"}, []string{"Romans 3:19–28", "John 8:31–36"}, usccbSource},
	{time.November, 1, "All Saints", []string{"Catholic", "Anglican", "Lutheran", "Many Protestant communities"}, "Feast", "Solemnity", "white", "The church celebrates all holy people, known and unknown.", "Holiness is understood as a shared vocation sustained by communion across generations.", []string{"Eucharist", "Naming saints and ancestors in faith"}, []string{"Revelation 7:2–14", "Matthew 5:1–12"}, usccbSource},
	{time.November, 2, "Commemoration of All the Faithful Departed · All Souls", []string{"Catholic", "Anglican communities"}, "Commemoration", "Commemoration", "violet", "The faithful departed are remembered in prayer.", "The day makes room for grief, hope, memory, and care for the dead.", []string{"Mass", "Visiting graves", "Prayer for the dead"}, []string{"Wisdom 3:1–9", "John 6:37–40"}, usccbSource},
	{time.December, 8, "Immaculate Conception of the Blessed Virgin Mary", []string{"Catholic"}, "Solemnity", "Solemnity", "white", "Catholics celebrate Mary as preserved from original sin from the first moment of her conception.", "The doctrine emphasizes grace preceding and enabling human response.", []string{"Mass", "Marian prayer"}, []string{"Genesis 3:9–20", "Luke 1:26–38"}, usccbSource},
	{time.December, 12, "Our Lady of Guadalupe", []string{"Catholic", "Especially Mexican and other American communities"}, "Feast", "Feast", "white", "The feast remembers Mary's appearances to Saint Juan Diego and the image associated with Tepeyac.", "Guadalupe is a powerful sign of divine closeness, inculturation, and dignity in the Americas.", []string{"Mass", "Mañanitas", "Processions", "Dance and flowers"}, []string{"Revelation 11:19–12:6", "Luke 1:39–47"}, usccbSource},
	{time.December, 24, "Christmas Eve", []string{"Catholic", "Protestant", "Anglican", "Orthodox on the revised calendar"}, "Vigil", "Solemnity vigil", "white", "Christian communities enter the celebration of Jesus' Nativity with evening worship.", "The vigil dwells in expectation as divine life is welcomed in vulnerability.", []string{"Vigil Mass or service", "Carols", "Candlelight", "Family gathering"}, []string{"Isaiah 9:1–6", "Luke 2:1–14"}, usccbSource},
	{time.December, 25, "Christmas · Nativity of the Lord", []string{"Catholic", "Protestant", "Anglican", "Orthodox on the revised calendar"}, "Feast", "Solemnity / Great Feast", "white", "Christians celebrate the birth of Jesus Christ.", "The incarnation proclaims divine love entering embodied, vulnerable human life.", []string{"Mass, Divine Liturgy, or worship", "Carols", "Feasting", "Giving"}, []string{"Isaiah 52:7–10", "John 1:1–18"}, usccbSource},
	{time.December, 26, "Saint Stephen, First Martyr", []string{"Catholic", "Orthodox", "Anglican", "Lutheran"}, "Feast", "Feast", "red", "Stephen is remembered as the first Christian martyr.", "Placed immediately after Christmas, his witness joins incarnation with costly discipleship and forgiveness.", []string{"Eucharist", "Works of charity"}, []string{"Acts 6:8–10; 7:54–59", "Matthew 10:17–22"}, usccbSource},
}

func christianObservances(date time.Time) []Observance {
	var events []Observance
	for _, fixed := range christianFixedDates {
		if date.Month() != fixed.month || date.Day() != fixed.day {
			continue
		}
		event := baseObservance(fixed.name, Christianity, fixed.communities, fixed.category, fixed.summary, fixed.meaning, fixed.practices, fixed.scripture, sourceNameFor(fixed.source), fixed.source)
		event.Rank = fixed.rank
		event.LiturgicalColor = fixed.color
		events = append(events, singleOccurrence(event, date))
	}

	easter := westernEaster(date.Year())
	relative := []struct {
		offset                                        int
		name, category, rank, color, summary, meaning string
		communities, practices, scripture             []string
	}{
		{-46, "Ash Wednesday", "Seasonal observance", "Day of fast and abstinence", "violet", "Western Christians begin Lent with ashes and a call to repentance.", "Mortality and mercy frame a forty-day return toward baptismal life.", []string{"Catholic", "Anglican", "Lutheran", "Methodist and other Protestant communities"}, []string{"Receiving ashes", "Fasting", "Prayer", "Almsgiving"}, []string{"Joel 2:12–18", "Matthew 6:1–18"}},
		{-7, "Palm Sunday of the Lord's Passion", "Holy Week", "Sunday", "red", "Jesus' entry into Jerusalem and the Passion are proclaimed together.", "Triumph and suffering are held in tension at the threshold of Holy Week.", []string{"Catholic", "Protestant", "Anglican"}, []string{"Palm procession", "Reading of the Passion", "Eucharist"}, []string{"Matthew 21:1–11", "Matthew 26–27"}},
		{-3, "Holy Thursday · Maundy Thursday", "Holy Week", "Triduum", "white", "Christians remember the Last Supper, Jesus washing feet, and the commandment of love.", "Eucharistic gift and servant love interpret one another.", []string{"Catholic", "Protestant", "Anglican"}, []string{"Eucharist", "Foot washing", "Stripping the altar", "Night watch"}, []string{"Exodus 12:1–14", "John 13:1–15"}},
		{-2, "Good Friday", "Holy Week", "Triduum", "red", "Christians commemorate the crucifixion and death of Jesus.", "The cross is contemplated with grief, repentance, and hope in self-giving love.", []string{"Catholic", "Protestant", "Anglican"}, []string{"Fasting", "Passion reading", "Veneration of the cross", "Solemn prayer"}, []string{"Isaiah 52:13–53:12", "John 18–19"}},
		{-1, "Holy Saturday · Easter Vigil", "Holy Week", "Triduum", "white", "The church waits at the tomb and, after nightfall, begins the celebration of resurrection.", "Darkness gives way to light through scripture, baptism, and Eucharist.", []string{"Catholic", "Protestant", "Anglican"}, []string{"Great Vigil", "Lighting the Paschal candle", "Baptisms", "Eucharist"}, []string{"Genesis 1:1–2:2", "Romans 6:3–11", "Mark 16:1–7"}},
		{0, "Easter Sunday · Resurrection of the Lord", "Feast", "Solemnity", "white", "Western Christians celebrate Jesus Christ raised from the dead.", "Resurrection is the center of Christian hope and the beginning of new creation.", []string{"Catholic", "Protestant", "Anglican"}, []string{"Festal Eucharist", "Renewal of baptismal promises", "Feasting"}, []string{"Acts 10:34–43", "John 20:1–9"}},
		{39, "Ascension of the Lord", "Feast", "Solemnity", "white", "Forty days after Easter, the church celebrates the risen Christ's ascension.", "Christ's departure commissions disciples toward witness and opens a new mode of presence.", []string{"Catholic where observed Thursday", "Anglican", "Some Protestant communities"}, []string{"Eucharist", "Prayer for mission"}, []string{"Acts 1:1–11", "Matthew 28:16–20"}},
		{49, "Pentecost", "Feast", "Solemnity", "red", "The Holy Spirit descends upon the disciples fifty days after Easter.", "Diverse languages become a sign of shared mission rather than forced sameness.", []string{"Catholic", "Protestant", "Anglican"}, []string{"Eucharist", "Red vesture", "Prayer for the gifts of the Spirit"}, []string{"Acts 2:1–11", "John 20:19–23"}},
		{56, "Trinity Sunday", "Feast", "Solemnity", "white", "Western churches celebrate the mystery of God as Trinity.", "Divine life is confessed as communion: Father, Son, and Holy Spirit.", []string{"Catholic", "Anglican", "Many Protestant communities"}, []string{"Eucharist", "Creedal prayer"}, []string{"Deuteronomy 4:32–40", "Matthew 28:16–20"}},
		{60, "Corpus Christi · Body and Blood of Christ", "Feast", "Solemnity", "white", "Catholics honor Christ's presence in the Eucharist.", "The sacrament joins worship, self-giving, and the formation of one body.", []string{"Catholic"}, []string{"Mass", "Eucharistic procession", "Adoration"}, []string{"Exodus 24:3–8", "Mark 14:12–26"}},
	}
	for _, item := range relative {
		when := easter.AddDate(0, 0, item.offset)
		if !sameDay(date, when) {
			continue
		}
		event := baseObservance(item.name, Christianity, item.communities, item.category, item.summary, item.meaning, item.practices, item.scripture, "USCCB liturgical calendar", usccbSource)
		event.Rank, event.LiturgicalColor = item.rank, item.color
		if item.offset == 39 {
			event.DateNote = "In many Catholic dioceses this solemnity is transferred to the following Sunday."
		}
		events = append(events, singleOccurrence(event, date))
	}

	pascha := orthodoxEaster(date.Year())
	orthodoxRelative := []struct {
		offset                 int
		name, summary, meaning string
	}{
		{-7, "Orthodox Palm Sunday", "Orthodox Christians commemorate Christ's entry into Jerusalem.", "The feast opens Holy Week with branches, procession, and attention to the coming Passion."},
		{-2, "Orthodox Great and Holy Friday", "Orthodox churches commemorate the saving Passion and burial of Christ.", "Lament and veneration contemplate divine self-emptying love."},
		{0, "Pascha · Orthodox Easter", "Orthodox Christians celebrate the resurrection of Jesus Christ.", "The proclamation “Christ is risen” announces life overcoming death and begins the festal season."},
		{39, "Orthodox Ascension", "The Orthodox Church celebrates Christ's ascension forty days after Pascha.", "Human nature is contemplated as raised into communion with God."},
		{49, "Orthodox Pentecost", "The Orthodox Church celebrates the descent of the Holy Spirit.", "The Spirit gathers and renews the church for witness."},
	}
	for _, item := range orthodoxRelative {
		when := pascha.AddDate(0, 0, item.offset)
		if sameDay(date, when) {
			event := baseObservance(item.name, Christianity, []string{"Eastern Orthodox"}, "Paschal cycle", item.summary, item.meaning, []string{"Divine Liturgy", "Festal hymns", "Fasting or feasting according to the day"}, []string{"Acts 1–2", "Gospel accounts of the Passion and Resurrection"}, "Orthodox Church in America", ocaSource)
			event.DateNote = "Calculated using the Orthodox Paschal cycle; local calendars should be consulted."
			events = append(events, singleOccurrence(event, date))
		}
	}

	advent := adventStart(date.Year())
	if sameDay(date, advent) {
		event := baseObservance("First Sunday of Advent", Christianity, []string{"Catholic", "Anglican", "Lutheran", "Many Protestant communities"}, "Seasonal observance", "The Western church begins a new liturgical year in watchful preparation for Christ's coming.", "Advent holds memory, present hope, and future expectation together.", []string{"Lighting an Advent wreath", "Eucharist", "Prayer and works of mercy"}, []string{"Isaiah 2:1–5", "Matthew 24:37–44"}, "USCCB liturgical calendar", usccbSource)
		event.Rank, event.LiturgicalColor = "Sunday", "violet"
		events = append(events, singleOccurrence(event, date))
	}
	christKing := advent.AddDate(0, 0, -7)
	if sameDay(date, christKing) {
		event := baseObservance("Our Lord Jesus Christ, King of the Universe", Christianity, []string{"Catholic", "Anglican and Protestant communities with related observances"}, "Feast", "The last Sunday of the Western liturgical year celebrates Christ's reign.", "Kingship is interpreted through truth, service, justice, and the cross rather than domination.", []string{"Mass or Eucharist", "Prayer for justice and peace"}, []string{"Daniel 7:13–14", "John 18:33–37"}, "USCCB liturgical calendar", usccbSource)
		event.Rank, event.LiturgicalColor = "Solemnity", "white"
		events = append(events, singleOccurrence(event, date))
	}
	return events
}

func sourceNameFor(source string) string {
	if source == ocaSource {
		return "Orthodox Church in America"
	}
	return "USCCB liturgical calendar"
}
