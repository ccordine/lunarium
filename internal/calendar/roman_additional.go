package calendar

import "time"

const (
	romanFeriaeTaxonomySource = "https://penelope.uchicago.edu/Thayer/E/Roman/Texts/secondary/SMIGRA*/Feriae.html"
	romanFornacaliaSource     = "https://penelope.uchicago.edu/Thayer/E/Roman/Texts/secondary/SMIGRA*/Fornacalia.html"
	romanMundusSource         = "https://penelope.uchicago.edu/Thayer/E/Journals/JRS/2/Mundus*.html"
	romanSeptimontiumSource   = "https://penelope.uchicago.edu/Thayer/E/Journals/AJA/12/2/Roma_Quadrata*.html"
)

func init() {
	romanLaterFixedFestivals = append(romanLaterFixedFestivals,
		romanLaterFixedFestival{"roman-mundus-patet-august", time.August, 24, "Mundus Patet · August Opening", "Roman religious-day calendar", "Republican and Imperial testimony", "The mundus was said to stand open on the day after the Volcanalia, a day the juristic tradition classified as religious.", "The opening joined underworld interpretation with older questions about the city's sacred store and seasonal grain cycle.", []string{"Avoidance of public business in later testimony", "Opening of the mundus in the received antiquarian tradition"}, "W. Warde Fowler · Mundus Patet", romanMundusSource, "Varro, Festus/Ateius Capito, Macrobius, and antiquarian analysis", "high for August 24 as one of the three received days; interpretation disputed"},
		romanLaterFixedFestival{"roman-mundus-patet-october", time.October, 5, "Mundus Patet · October Opening", "Roman religious-day calendar", "Republican and Imperial testimony", "Roman antiquarian sources list October 5 as the second of three annual days on which the mundus stood open.", "The day was treated as religiosus; explanations involving the dead and the agricultural storehouse belong to layered interpretation.", []string{"Avoidance of public business in later testimony", "Opening of the mundus in the received antiquarian tradition"}, "W. Warde Fowler · Mundus Patet", romanMundusSource, "Varro, Festus/Ateius Capito, Macrobius, and antiquarian analysis", "high for October 5 as one of the three received days; interpretation disputed"},
		romanLaterFixedFestival{"roman-mundus-patet-november", time.November, 8, "Mundus Patet · November Opening", "Roman religious-day calendar", "Republican and Imperial testimony", "Roman antiquarian sources list November 8 as the third annual opening of the mundus.", "The day was marked as religious rather than treated here as a modernized festival of ghosts.", []string{"Avoidance of public business in later testimony", "Opening of the mundus in the received antiquarian tradition"}, "W. Warde Fowler · Mundus Patet", romanMundusSource, "Varro, Festus/Ateius Capito, Macrobius, and antiquarian analysis", "high for November 8 as one of the three received days; interpretation disputed"},
		romanLaterFixedFestival{"roman-septimontium", time.December, 11, "Septimontium · Festival of the Hills", "Roman neighborhood and topographic cult calendar", "Republican and Imperial Rome; origins disputed", "Sacrifices at named hill districts marked a festival remembered as the Septimontium.", "The rite preserved an old topography of community and cult, although ancient and modern accounts disagree about which 'seven' hills or settlements the name originally described.", []string{"Sacrifices at multiple hill districts", "Public distributions in Imperial testimony"}, "J. B. Carter · Roma Quadrata and the Septimontium", romanSeptimontiumSource, "Festus, Varro, Imperial testimony, and calendar/topographic scholarship", "high for December 11 in later calendars; origin and original district list disputed"},
	)

	ancientCatalogRecords = append(ancientCatalogRecords,
		romanMovableRecord(
			"roman-fornacalia-conceptiva", "Fornacalia · Proclaimed Curial Festival", "Roman curial and civic communities", "Roman proclaimed grain festival",
			"The Curio Maximus proclaimed a schedule for the curiae to honor Fornax so that grain would be properly parched; people who did not know their curia observed the concluding Quirinalia.",
			"The festival tied food preparation to Rome's archaic curial organization and shows why not every annual Roman feast had a fixed date.",
			[]string{"Curia-specific rites for Fornax", "Publicly posted annual schedule", "Concluding observance by the Quirinalia on February 17"},
			"Annually proclaimed before the Quirinalia; each curia received its own date", "Curio Maximus proclamation; final day fixed by the Quirinalia",
			"A fixed February 17 occurrence represents the Quirinalia, not every curia's Fornacalia date.", "Smith's Dictionary · Fornacalia", romanFornacaliaSource,
		),
		romanMovableRecord(
			"roman-sementivae-paganalia", "Sementivae and Paganalia · Proclaimed Sowing Rites", "Roman rural pagi and civic communities", "Roman proclaimed agricultural feriae",
			"Magistrates or priests annually proclaimed the winter sowing observances known as the Sementivae and Paganalia rather than assigning them permanent fasti dates.",
			"The rites located sowing, seed, neighborhood, and civic authority within a flexible agricultural calendar.",
			[]string{"Proclamation of the annual dates", "Rural community gathering", "Offerings connected with sowing and the fields"},
			"Feriae conceptivae in the winter agricultural season; annual dates proclaimed", "Macrobian and Varronian classification as conceptivae",
			"The related rites and their timing changed; no single January day is projected.", "Smith's Dictionary · Feriae", romanFeriaeTaxonomySource,
		),
		romanMovableRecord(
			"roman-feriae-latinae", "Feriae Latinae · Latin Festival", "Roman and Latin civic communities", "Roman-Latin proclaimed federal festival",
			"Rome's magistrates proclaimed the annual Latin Festival on the Alban Mount, where participating communities honored Jupiter Latiaris.",
			"The shared sacrifice renewed a political and sacred relationship among Rome and Latin communities before the consuls could depart on campaign.",
			[]string{"Sacrifice to Jupiter Latiaris", "Distribution of sacrificial portions among participating communities", "Magisterial and federal gathering"},
			"Annual date proclaimed by the consuls; multi-day duration in later evidence", "Literary and antiquarian evidence for feriae conceptivae",
			"Political circumstances could delay the proclamation; a permanent modern anniversary would erase that constitutional role.", "Smith's Dictionary · Feriae Latinae", romanFeriaeTaxonomySource,
		),
		romanMovableRecord(
			"roman-compitalia-conceptiva", "Compitalia · Proclaimed Crossroads Festival", "Roman neighborhoods, households, and collegia", "Roman proclaimed neighborhood festival",
			"Neighborhoods honored the Lares Compitales at crossroads on dates that were historically proclaimed rather than permanently fixed.",
			"The festival joined household protection, local identity, and the organization of urban and rural neighborhoods.",
			[]string{"Garlanding crossroads shrines", "Neighborhood offerings and meals", "Games in later phases"},
			"Annual date proclaimed; late calendars supply conventional January placements", "Macrobian classification as conceptivae and later calendar witnesses",
			"The dated January 3–5 display elsewhere in the app is explicitly a late conventional layer, not the universal Republican rule.", "Smith's Dictionary · Feriae", romanFeriaeTaxonomySource,
		),
	)
}

func romanMovableRecord(id, name, communities, category, summary, meaning string, elements []string, nativeDate, attestation, note, sourceName, sourceURL string) ancientCatalogRecord {
	return ancientCatalogRecord{
		ID:               id,
		Name:             name,
		Communities:      []string{communities},
		Category:         category,
		Summary:          summary,
		Meaning:          meaning,
		AttestedElements: elements,
		Texts:            []string{"Varro, De lingua Latina 6", "Macrobius, Saturnalia 1.16", "Roman antiquarian and historical witnesses named by the source"},
		Origin:           "Roman feriae conceptivae",
		HistoricalNote:   note,
		CalendarCorpus:   "Roman proclaimed and movable feriae",
		NativeDateLabel:  nativeDate,
		AttestationLayer: attestation,
		Era:              "Roman Republic with later literary and calendar witnesses",
		Site:             "Rome and the relevant curial, neighborhood, rural, or Latin federal communities",
		ProjectionKind:   "annually-proclaimed-native-rule",
		ProjectionStatus: "Catalog-only: the ancient date depended on an annual priestly or magisterial proclamation, so the app does not manufacture a fixed Gregorian recurrence.",
		DateConfidence:   "High classification as a recurring movable feriae; no year-independent day",
		AnchorLocation:   "Rome and Latium",
		DayBoundary:      "Roman civil day; the annual proclamation supplied the operative date",
		SourceName:       sourceName,
		SourceURL:        sourceURL,
	}
}
