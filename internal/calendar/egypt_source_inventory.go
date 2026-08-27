package calendar

// uclEgyptSpec keeps the UCL/Schott source inventory compact while requiring
// the period, place, and native-calendar resolution to remain explicit.
type uclEgyptSpec struct {
	ID               string
	Name             string
	Category         string
	Summary          string
	Meaning          string
	NativeDate       string
	AttestationLayer string
	Era              string
	Site             string
	Elements         []string
	Texts            []string
	Status           string
	ProjectionKind   string
	Confidence       string
	Caveat           string
}

func uclEgyptRecord(spec uclEgyptSpec) ancientCatalogRecord {
	category := spec.Category
	if category == "" {
		category = "Egyptian source-inventory festival record"
	}
	meaning := spec.Meaning
	if meaning == "" {
		meaning = "The entry preserves a source-specific piece of Egyptian sacred time without turning it into a timeless national calendar."
	}
	elements := spec.Elements
	if elements == nil {
		elements = []string{"Named or dated rite in the cited festival inventory"}
	}
	texts := spec.Texts
	if texts == nil {
		texts = []string{"Siegfried Schott festival-date synthesis as presented by UCL Digital Egypt"}
	}
	projectionKind := spec.ProjectionKind
	if projectionKind == "" {
		projectionKind = "source-layer-native-date"
	}
	confidence := spec.Confidence
	if confidence == "" {
		confidence = "High for the cited inventory row; wider recurrence and geographic scope vary"
	}
	caveat := spec.Caveat
	if caveat == "" {
		caveat = "UCL warns that Schott's linear list overlooks major variation across time and place; this record retains its named source layer and does not assert uninterrupted national observance."
	}
	status := spec.Status
	if status == "" {
		status = "Historical source-inventory record; no continuous modern observance asserted"
	}
	dayBoundary := "Dawn/civil day; the inventory does not define finer local boundary practice"
	if projectionKind == "source-list-name-only" {
		dayBoundary = "Not preserved in the named festival list"
	}
	return ancientCatalogRecord{
		ID:               spec.ID,
		Name:             spec.Name,
		Communities:      []string{"Ancient Egyptian temple, court, or local communities in the cited source layer"},
		Category:         category,
		Summary:          spec.Summary,
		Meaning:          meaning,
		AttestedElements: elements,
		Texts:            texts,
		Origin:           "Egyptian festival lists and documents inventoried by UCL Digital Egypt after Schott",
		Status:           status,
		HistoricalNote:   caveat,
		CalendarCorpus:   "Egyptian civil and temple calendars",
		NativeDateLabel:  spec.NativeDate,
		AttestationLayer: spec.AttestationLayer,
		Era:              spec.Era,
		Site:             spec.Site,
		ProjectionKind:   projectionKind,
		ProjectionStatus: "Catalog-only: the unreformed 365-day civil calendar drifted against the seasons, and no conversion is made without a specified historical year and source layer.",
		DateConfidence:   confidence,
		AnchorLocation:   spec.Site,
		DayBoundary:      dayBoundary,
		SourceName:       "UCL Digital Egypt · Festivals in the ancient Egyptian calendar",
		SourceURL:        egyptianFestivalSource,
	}
}

var uclEgyptExpansionRecords = []ancientCatalogRecord{
	// The page first gives an undated Old Kingdom tomb-chapel name list. These
	// six names cannot safely be folded into similarly named later programs.
	uclEgyptRecord(uclEgyptSpec{
		ID: "egypt-thoth-old-kingdom", Name: "Thoth Festival · Old Kingdom List Name",
		Summary:    "Old Kingdom tomb-chapel festival lists name a Thoth festival but preserve no procedure or numbered day on the UCL inventory page.",
		NativeDate: "Named in Old Kingdom tomb-chapel lists; numbered day not supplied", AttestationLayer: "Old Kingdom tomb-chapel festival-name list", Era: "Old Kingdom, approximately 2686–2181 BCE", Site: "Old Kingdom tomb-chapel evidence; locality varies", ProjectionKind: "source-list-name-only", Confidence: "High name attestation in the cited inventory; date and ritual program unavailable",
	}),
	uclEgyptRecord(uclEgyptSpec{
		ID: "egypt-first-of-year-old-kingdom", Name: "First of the Year · Old Kingdom List Name",
		Summary:    "The Old Kingdom list distinguishes a festival called First of the Year from Opening of the Year, but UCL supplies no date or procedure for the distinction.",
		NativeDate: "Named in Old Kingdom tomb-chapel lists; relationship to Opening of the Year unresolved", AttestationLayer: "Old Kingdom tomb-chapel festival-name list", Era: "Old Kingdom, approximately 2686–2181 BCE", Site: "Old Kingdom tomb-chapel evidence; locality varies", ProjectionKind: "source-list-name-only", Confidence: "High listed name; low confidence in identity, schedule, or relationship to Wepet Renpet",
	}),
	uclEgyptRecord(uclEgyptSpec{
		ID: "egypt-sokar-old-kingdom", Name: "Sokar Festival · Old Kingdom List Layer",
		Summary:    "Old Kingdom tomb-chapel lists name a Sokar festival independently of the much later detailed Khoiak schedules.",
		NativeDate: "Named in Old Kingdom tomb-chapel lists; no numbered day supplied for this layer", AttestationLayer: "Old Kingdom festival-name list, kept separate from later Khoiak programs", Era: "Old Kingdom, approximately 2686–2181 BCE", Site: "Memphite and other Old Kingdom mortuary contexts", ProjectionKind: "source-list-name-only", Confidence: "High listed name; date and continuity into later Sokar rites unresolved",
	}),
	uclEgyptRecord(uclEgyptSpec{
		ID: "egypt-great-festival-old-kingdom", Name: "Great Festival · Old Kingdom List Name",
		Summary:    "A festival called simply the Great Festival occurs in Old Kingdom tomb-chapel lists; the abbreviated name does not securely identify one later rite.",
		NativeDate: "Named in Old Kingdom tomb-chapel lists; date and cult owner not supplied", AttestationLayer: "Old Kingdom tomb-chapel festival-name list", Era: "Old Kingdom, approximately 2686–2181 BCE", Site: "Old Kingdom tomb-chapel evidence; locality varies", ProjectionKind: "source-list-name-only", Confidence: "High listed title; low confidence in identification beyond that title",
	}),
	uclEgyptRecord(uclEgyptSpec{
		ID: "egypt-flame-festival-old-kingdom", Name: "Flame Festival · Old Kingdom List Name",
		Summary:    "Old Kingdom tomb-chapel lists preserve the name Flame Festival without a numbered day or surviving procedural account on the UCL page.",
		NativeDate: "Named in Old Kingdom tomb-chapel lists; numbered day not supplied", AttestationLayer: "Old Kingdom tomb-chapel festival-name list", Era: "Old Kingdom, approximately 2686–2181 BCE", Site: "Old Kingdom tomb-chapel evidence; locality varies", ProjectionKind: "source-list-name-only", Confidence: "High listed name; calendar placement and rites unresolved",
	}),
	uclEgyptRecord(uclEgyptSpec{
		ID: "egypt-sadj-old-kingdom", Name: "Sadj Festival · Old Kingdom List Name",
		Summary:    "Old Kingdom tomb-chapel lists name a Sadj festival, but the UCL inventory supplies neither a day nor a ritual reconstruction.",
		NativeDate: "Named in Old Kingdom tomb-chapel lists; numbered day not supplied", AttestationLayer: "Old Kingdom tomb-chapel festival-name list", Era: "Old Kingdom, approximately 2686–2181 BCE", Site: "Old Kingdom tomb-chapel evidence; locality varies", ProjectionKind: "source-list-name-only", Confidence: "High listed name; interpretation and calendar placement unresolved",
	}),

	// Missing rows in the page's explicit month-by-month Schott inventory.
	uclEgyptRecord(uclEgyptSpec{
		ID: "egypt-hapy-amun-flood-offerings", Name: "Offerings to Hapy and Amun for the Flood · Gebel el-Silsila",
		Category: "Egyptian inundation offering", Summary: "Dynasty 19 rock inscriptions at Gebel el-Silsila place offerings to Hapy and Amun on the fifteenth day in two different civil months to seek a good flood.",
		NativeDate: "I Akhet 15 and III Shemu 15 in the cited Dynasty 19 inscriptions", AttestationLayer: "Dynasty 19 rock inscriptions", Era: "New Kingdom, Dynasty 19", Site: "Gebel el-Silsila", Elements: []string{"Offerings to Hapy", "Offerings to Amun", "Petition for a good inundation"}, Confidence: "High for both cited civil dates and site; the relationship between the two entries is not reconstructed",
	}),
	uclEgyptRecord(uclEgyptSpec{
		ID: "egypt-tekh-drunkenness", Name: "Tekh · Festival of Drunkenness",
		Category: "Egyptian festival notice", Summary: "The linear UCL/Schott inventory places Tekh, glossed as drunkenness, on the twentieth day of the year's first civil month.",
		NativeDate: "I Akhet 20", AttestationLayer: "Schott festival-date synthesis; underlying source layer not detailed on the UCL page", Era: "Pharaonic Egypt; precise source horizon not specified on the inventory page", Site: "Egypt; local scope unresolved", Elements: []string{"Named Tekh observance", "Drunkenness gloss in the source inventory"}, Confidence: "Moderate: native row is explicit, but its period, locality, and ritual detail are not supplied",
	}),
	uclEgyptRecord(uclEgyptSpec{
		ID: "egypt-elephantine-khnum-anuqet", Name: "Festival of Khnum and Anuqet · Elephantine",
		Category: "Elephantine local temple festival", Summary: "Thutmose III's Elephantine festival list records a local rite for Khnum and Anuqet.",
		NativeDate: "II Akhet 18", AttestationLayer: "Festival list of Thutmose III at Elephantine", Era: "New Kingdom, reign of Thutmose III", Site: "Elephantine", Elements: []string{"Local festival of Khnum and Anuqet"}, Confidence: "High for the cited reign, site, and civil day; no national scope implied",
	}),
	uclEgyptRecord(uclEgyptSpec{
		ID: "egypt-elephantine-satet-anuqet", Name: "Festival of Satet and Anuqet · Elephantine",
		Category: "Elephantine local temple festival", Summary: "Thutmose III's Elephantine list records a local festival for Satet and Anuqet late in the second month.",
		NativeDate: "II Akhet 28", AttestationLayer: "Festival list of Thutmose III at Elephantine", Era: "New Kingdom, reign of Thutmose III", Site: "Elephantine", Elements: []string{"Local festival of Satet and Anuqet"}, Confidence: "High for the cited reign, site, and civil day; no national scope implied",
	}),
	uclEgyptRecord(uclEgyptSpec{
		ID: "egypt-elephantine-amun", Name: "Festival for Amun · Elephantine",
		Category: "Elephantine local temple festival", Summary: "Thutmose III's Elephantine festival list places a festival for Amun on day nine of the third civil month.",
		NativeDate: "III Akhet 9", AttestationLayer: "Festival list of Thutmose III at Elephantine", Era: "New Kingdom, reign of Thutmose III", Site: "Elephantine", Elements: []string{"Local festival for Amun"}, Confidence: "High for the cited reign, site, and civil day",
	}),
	uclEgyptRecord(uclEgyptSpec{
		ID: "egypt-elephantine-anuqet", Name: "Festival of Anuqet · Elephantine",
		Category: "Elephantine local temple festival", Summary: "Thutmose III's Elephantine festival list records a local observance of Anuqet on the last day of the third civil month.",
		NativeDate: "III Akhet 30", AttestationLayer: "Festival list of Thutmose III at Elephantine", Era: "New Kingdom, reign of Thutmose III", Site: "Elephantine", Elements: []string{"Local festival of Anuqet"}, Confidence: "High for the cited reign, site, and civil day",
	}),
	uclEgyptRecord(uclEgyptSpec{
		ID: "egypt-hathor-medinet-habu", Name: "Festival for Hathor · Medinet Habu List",
		Category: "Theban temple festival", Summary: "Ramesses III's great festival list at Medinet Habu places a festival for Hathor on the first day of the fourth civil month.",
		NativeDate: "IV Akhet 1", AttestationLayer: "Great festival list of Ramesses III at Medinet Habu", Era: "New Kingdom, Dynasty 20", Site: "Medinet Habu, Thebes", Elements: []string{"Festival for Hathor"}, Confidence: "High for the Ramesside list's day and site",
	}),
	uclEgyptRecord(uclEgyptSpec{
		ID: "egypt-sailing-wadjyt", Name: "Sailing of Wadjyt · Karnak",
		Category: "Egyptian sacred-bark rite", Summary: "An inscription of Thutmose III at the Mut temple records the sailing of Wadjyt.",
		NativeDate: "I Peret 20", AttestationLayer: "Thutmose III inscription at the temple of Mut", Era: "New Kingdom, reign of Thutmose III", Site: "Temple of Mut, Karnak, Thebes", Elements: []string{"Sailing rite for Wadjyt"}, Confidence: "High for the cited inscription's civil date and site",
	}),
	uclEgyptRecord(uclEgyptSpec{
		ID: "egypt-sailing-bast", Name: "Sailing of Bast · Karnak",
		Category: "Egyptian sacred-bark rite", Summary: "An inscription of Thutmose III at the Mut temple records a sailing of Bast.",
		NativeDate: "I Peret 29", AttestationLayer: "Thutmose III inscription at the temple of Mut", Era: "New Kingdom, reign of Thutmose III", Site: "Temple of Mut, Karnak, Thebes", Elements: []string{"Sailing rite for Bast"}, Confidence: "High for the cited inscription's civil date and site",
	}),
	uclEgyptRecord(uclEgyptSpec{
		ID: "egypt-raising-willow", Name: "Festival of Raising the Willow · Medinet Habu",
		Category: "Egyptian temple festival", Summary: "Ramesses III's Medinet Habu list places the Raising the Willow festival on the same civil day as a separately attested sailing of Bast.",
		NativeDate: "I Peret 29", AttestationLayer: "Great festival list of Ramesses III at Medinet Habu", Era: "New Kingdom, Dynasty 20", Site: "Medinet Habu, Thebes", Elements: []string{"Raising the willow"}, Confidence: "High for the Ramesside list's name and civil date; its relationship to other same-day rites is not inferred",
	}),
	uclEgyptRecord(uclEgyptSpec{
		ID: "egypt-sailing-shesmet", Name: "Sailing of Shesmet · Thutmose III Layer",
		Category: "Egyptian sacred-bark rite", Summary: "A Thutmose III inscription at the Mut temple records the sailing of Shesmet on the last day of the first sowing-season month.",
		NativeDate: "I Peret 30", AttestationLayer: "Thutmose III inscription at the temple of Mut", Era: "New Kingdom, reign of Thutmose III", Site: "Temple of Mut, Karnak, Thebes", Elements: []string{"Sailing rite for Shesmet"}, Confidence: "High for this reign-specific inscription and civil day",
	}),
	uclEgyptRecord(uclEgyptSpec{
		ID: "egypt-sailing-mut-isheru", Name: "Sailing of Mut, Lady of Isheru · Late New Kingdom Layer",
		Category: "Egyptian sacred-bark rite", Summary: "A late New Kingdom Turin papyrus identifies the same inventory position as the sailing of Mut, Lady of Isheru, rather than silently harmonizing it with the earlier Shesmet notice.",
		NativeDate: "I Peret 30 in the UCL inventory's late New Kingdom variant", AttestationLayer: "Late New Kingdom Turin papyrus", Era: "Late New Kingdom", Site: "Theban Mut cult; Turin papyrus witness", Elements: []string{"Sailing rite for Mut, Lady of Isheru"}, Confidence: "High for the listed textual variant; relationship to the Thutmose III Shesmet entry remains source-layered",
	}),
	uclEgyptRecord(uclEgyptSpec{
		ID: "egypt-sailing-anubis", Name: "Sailing of Anubis",
		Category: "Egyptian sacred-bark rite", Summary: "The UCL/Schott inventory places a sailing rite of Anubis at the opening of the second sowing-season month.",
		NativeDate: "II Peret 1", AttestationLayer: "Schott festival-date synthesis; underlying witness not identified on the UCL page", Era: "Pharaonic Egypt; precise source horizon not specified on the inventory page", Site: "Egypt; local anchor unresolved", Elements: []string{"Sailing rite for Anubis"}, Confidence: "Moderate: name and native date are explicit in the inventory, but source period and locality are not",
	}),
	uclEgyptRecord(uclEgyptSpec{
		ID: "egypt-raising-heaven-cycle", Name: "Amun in the Festival of Raising Heaven · Midyear Cycle",
		Category: "Egyptian multi-day temple cycle", Summary: "A several-day cycle reaches a key point on II Peret 30 with ished branches and culminates the next day with the image's return and the filling of the sacred eye at Iunu.",
		NativeDate: "Key date II Peret 30; culmination and image return on III Peret 1", AttestationLayer: "Multiple sources synthesized by Schott; variants include ished-tree branches and the filling of the sacred eye", Era: "Pharaonic Egypt; source layers vary", Site: "Iunu/Heliopolis traditions with a Theban work-journal notice", Elements: []string{"Raising-heaven festival", "Bringing ished-tree branches in some sources", "Filling the sacred eye at Iunu", "Return of the divine image"}, ProjectionKind: "layered-native-calendar", Confidence: "Moderate to high for the two-day hinge; name and component rites vary by source",
	}),
	uclEgyptRecord(uclEgyptSpec{
		ID: "egypt-ptah-thebes-peret", Name: "Festival of Ptah · Theban Work Journal",
		Category: "Possibly local Theban festival", Summary: "The journal for work on the king's tomb records a festival of Ptah on the first day of the third sowing-season month, perhaps local to Thebes.",
		NativeDate: "III Peret 1", AttestationLayer: "Journal for work on the king's tomb", Era: "New Kingdom Deir el-Medina administrative horizon", Site: "Theban royal-necropolis workforce context", Elements: []string{"Festival of Ptah"}, Confidence: "High source day; geographic scope explicitly uncertain",
	}),
	uclEgyptRecord(uclEgyptSpec{
		ID: "egypt-amenhotep-i-valley", Name: "Festival of King Amenhotep I in the Valley",
		Category: "Theban royal cult festival", Summary: "A festival of the deified Amenhotep I in the valley is listed for day 21, with UCL noting an originally local Theban setting and possible later widening.",
		NativeDate: "III Peret 21", AttestationLayer: "Theban royal-cult calendar evidence summarized by UCL/Schott", Era: "New Kingdom and later Theban royal cult", Site: "Theban valley; later reach uncertain", Elements: []string{"Festival of the deified king Amenhotep I"}, Confidence: "High listed native day; historical expansion beyond Thebes remains a question",
	}),
	uclEgyptRecord(uclEgyptSpec{
		ID: "egypt-amenhotep-i-workforce", Name: "Four-Day Festival of Amenhotep I · Deir el-Medina",
		Category: "Deir el-Medina workforce festival", Summary: "The Deir el-Medina workforce observed a four-day festival of the deified Amenhotep I beginning around day 29.",
		NativeDate: "III Peret 29 start, possibly four days", AttestationLayer: "Deir el-Medina workforce records", Era: "New Kingdom royal-necropolis community", Site: "Deir el-Medina, Thebes", Elements: []string{"Four-day workforce festival", "Cult of Amenhotep I"}, Confidence: "Moderate to high: local festival and duration are attested; UCL marks the start day with uncertainty",
	}),
	uclEgyptRecord(uclEgyptSpec{
		ID: "egypt-bast-onion", Name: "Festival of Bast · Chewing Onions",
		Category: "Egyptian Bast festival", Summary: "The inventory records a festival of Bast also described as the day of chewing onions for Bast.",
		NativeDate: "IV Peret 4", AttestationLayer: "Festival-calendar notice synthesized by Schott", Era: "Pharaonic Egypt; exact source horizon not specified on the UCL page", Site: "Egypt; local scope unresolved", Elements: []string{"Festival of Bast", "Chewing onions for Bast"}, Confidence: "Moderate: name, action, and native day are explicit; site and period are not",
	}),
	uclEgyptRecord(uclEgyptSpec{
		ID: "egypt-bast-boat-appearance", Name: "Appearance of Bast in Her Boat",
		Category: "Egyptian divine appearance and bark rite", Summary: "A Dynasty 26 statue records the appearance of Bast in her boat on the day after the onion rite listed by UCL.",
		NativeDate: "IV Peret 5", AttestationLayer: "Dynasty 26 statue Louvre A88", Era: "Late Period, Dynasty 26", Site: "Provenance of Louvre A88; Bast cult setting", Elements: []string{"Appearance of Bast", "Divine boat"}, Confidence: "High for the statue's listed day; local performance context is incomplete",
	}),
	uclEgyptRecord(uclEgyptSpec{
		ID: "egypt-renenutet-harvest-offering", Name: "Harvest Offering to Renenutet · Theban Tomb 38",
		Category: "Theban agricultural offering", Summary: "A depiction in Theban Tomb 38 places a harvest offering to Renenutet on day 25 of her namesake month.",
		NativeDate: "IV Peret 25", AttestationLayer: "Depiction in Theban Tomb-chapel 38", Era: "New Kingdom Theban funerary evidence", Site: "Theban Tomb 38", Elements: []string{"Harvest offering to Renenutet"}, Confidence: "High for the depicted offering and native day in this tomb source",
	}),
	uclEgyptRecord(uclEgyptSpec{
		ID: "egypt-renenutet-granary-offering", Name: "Granary Offering to Renenutet · Theban Tomb 48",
		Category: "Theban agricultural offering", Summary: "A depiction in Theban Tomb 48 places a granary offering to Renenutet two days after the separate harvest-offering notice.",
		NativeDate: "IV Peret 27", AttestationLayer: "Depiction in Theban Tomb-chapel 48", Era: "New Kingdom Theban funerary evidence", Site: "Theban Tomb 48", Elements: []string{"Granary offering to Renenutet"}, Confidence: "High for the depicted offering and native day in this tomb source",
	}),
	uclEgyptRecord(uclEgyptSpec{
		ID: "egypt-anubis-adoration", Name: "Adoration of Anubis",
		Category: "Egyptian temple rite", Summary: "The UCL/Schott inventory lists an adoration of Anubis on day ten of the first summer-season month.",
		NativeDate: "I Shemu 10", AttestationLayer: "Schott festival-date synthesis; underlying witness not detailed on the UCL page", Era: "Pharaonic Egypt; precise source horizon not specified", Site: "Egypt; local scope unresolved", Elements: []string{"Adoration of Anubis"}, Confidence: "Moderate: rite and native day are explicit in the inventory, while period and site are not",
	}),
	uclEgyptRecord(uclEgyptSpec{
		ID: "egypt-hathor-eve-thebes", Name: "Eve of the Hathor Festival · Thebes",
		Category: "Theban festival eve", Summary: "A stela of Thutmose III records the eve of a Hathor festival at Thebes on the last day of the third summer-season month.",
		NativeDate: "III Shemu 30", AttestationLayer: "Stela of Thutmose III, Cairo CG 34013", Era: "New Kingdom, reign of Thutmose III", Site: "Thebes", Elements: []string{"Festival eve for Hathor"}, Confidence: "High for the cited stela's place and civil day; the following festival program is not supplied",
	}),
	uclEgyptRecord(uclEgyptSpec{
		ID: "egypt-unnamed-iv-shemu-festival", Name: "Unnamed Two-Day Festival · Deir el-Medina Ostracon 209",
		Category: "Egyptian locally attested unnamed festival", Summary: "A late New Kingdom ostracon records a two-day festival but does not preserve its occasion.",
		NativeDate: "IV Shemu 1–2", AttestationLayer: "Late New Kingdom ostracon Deir el-Medina 209 verso, line 4", Era: "Late New Kingdom", Site: "Deir el-Medina, Thebes", Elements: []string{"Two-day festival; occasion unspecified"}, Status: "Historical dated event; festival identity and recurrence unresolved", ProjectionKind: "unnamed-source-event", Confidence: "High source dates; identity, cult owner, and annual recurrence unavailable",
	}),
	uclEgyptRecord(uclEgyptSpec{
		ID: "egypt-ipip-festival", Name: "Ipip Festival · Royal Tomb Work Journal",
		Category: "Egyptian festival notice", Summary: "The journal for work on the king's tomb records an Ipip festival on day two of the twelfth civil month.",
		NativeDate: "IV Shemu 2", AttestationLayer: "Necropolis Journal plate 59, line 19", Era: "New Kingdom royal-necropolis work-journal horizon", Site: "Theban royal necropolis", Elements: []string{"Ipip festival"}, Confidence: "High named notice and civil day; ritual program not supplied",
	}),
	uclEgyptRecord(uclEgyptSpec{
		ID: "egypt-ptah-middle-kingdom-pyramid", Name: "Festival of Ptah · Middle Kingdom Pyramid Inscription",
		Category: "Possibly local Ptah festival", Summary: "A rough inscription on a Middle Kingdom pyramid records a festival of Ptah late in the twelfth civil month; UCL marks its locality as uncertain.",
		NativeDate: "IV Shemu 24", AttestationLayer: "Rough inscription on a Middle Kingdom pyramid", Era: "Middle Kingdom", Site: "Middle Kingdom pyramid context; locality not specified on the UCL page", Elements: []string{"Festival of Ptah"}, Confidence: "High source date; locality and wider recurrence uncertain",
	}),
	uclEgyptRecord(uclEgyptSpec{
		ID: "egypt-eve-opening-year", Name: "Eve of the Opening of the Year",
		Category: "Egyptian year-boundary eve", Summary: "The final ordinary civil day is listed as the eve of the start of the year, before the five days over the year and the next Wepet Renpet.",
		NativeDate: "IV Shemu 30, before the five epagomenal days", AttestationLayer: "Schott festival-date synthesis", Era: "Pharaonic civil-calendar tradition; local rites vary", Site: "Egypt; source list is synthetic", Elements: []string{"Eve of the year-boundary cycle"}, Confidence: "High relative calendar position; specific ritual program not supplied",
	}),

	// Each day over the year is separately discoverable; the previous grouped
	// record was removed so the five UCL rows are not duplicated.
	uclEgyptRecord(uclEgyptSpec{
		ID: "egypt-epagomenal-osiris", Name: "Birthday of Osiris · First Day over the Year",
		Category: "Egyptian epagomenal divine birthday", Summary: "The first of the five days added after the twelve civil months was celebrated as the birthday of Osiris.",
		NativeDate: "First day over the year (civil day 361), after IV Shemu 30", AttestationLayer: "Civil-calendar structure and layered divine-birthday lists", Era: "Pharaonic Egypt; attestations are layered", Site: "Multiple Egyptian temple-calendar traditions", Elements: []string{"Birthday of Osiris"}, Confidence: "High position in the received five-day sequence; local and period detail varies",
	}),
	uclEgyptRecord(uclEgyptSpec{
		ID: "egypt-epagomenal-horus", Name: "Birthday of Horus · Second Day over the Year",
		Category: "Egyptian epagomenal divine birthday", Summary: "The second of the five days added after the twelve civil months was celebrated as the birthday of Horus.",
		NativeDate: "Second day over the year (civil day 362)", AttestationLayer: "Civil-calendar structure and layered divine-birthday lists", Era: "Pharaonic Egypt; attestations are layered", Site: "Multiple Egyptian temple-calendar traditions", Elements: []string{"Birthday of Horus"}, Confidence: "High position in the received five-day sequence; local and period detail varies",
	}),
	uclEgyptRecord(uclEgyptSpec{
		ID: "egypt-epagomenal-seth", Name: "Birthday of Seth · Third Day over the Year",
		Category: "Egyptian epagomenal divine birthday", Summary: "The third of the five days added after the twelve civil months was celebrated as the birthday of Seth.",
		NativeDate: "Third day over the year (civil day 363)", AttestationLayer: "Civil-calendar structure and layered divine-birthday lists", Era: "Pharaonic Egypt; attestations are layered", Site: "Multiple Egyptian temple-calendar traditions", Elements: []string{"Birthday of Seth"}, Confidence: "High position in the received five-day sequence; local and period detail varies",
	}),
	uclEgyptRecord(uclEgyptSpec{
		ID: "egypt-epagomenal-isis", Name: "Birthday of Isis · Fourth Day over the Year",
		Category: "Egyptian epagomenal divine birthday", Summary: "The fourth of the five days added after the twelve civil months was celebrated as the birthday of Isis.",
		NativeDate: "Fourth day over the year (civil day 364)", AttestationLayer: "Civil-calendar structure and layered divine-birthday lists", Era: "Pharaonic Egypt; attestations are layered", Site: "Multiple Egyptian temple-calendar traditions", Elements: []string{"Birthday of Isis"}, Confidence: "High position in the received five-day sequence; local and period detail varies",
	}),
	uclEgyptRecord(uclEgyptSpec{
		ID: "egypt-epagomenal-nephthys", Name: "Birthday of Nephthys · Fifth Day over the Year",
		Category: "Egyptian epagomenal divine birthday", Summary: "The fifth day added after the twelve civil months was celebrated as the birthday of Nephthys and immediately preceded the next civil new year.",
		NativeDate: "Fifth day over the year (civil day 365), before I Akhet 1", AttestationLayer: "Civil-calendar structure and layered divine-birthday lists", Era: "Pharaonic Egypt; attestations are layered", Site: "Multiple Egyptian temple-calendar traditions", Elements: []string{"Birthday of Nephthys"}, Confidence: "High position in the received five-day sequence; local and period detail varies",
	}),
}

// uclFestivalCoverageRow is a closed manifest of the named rows on the UCL
// page. Several rows correctly map to one layered record, and a source row
// containing explicit variants may map to more than one record.
type uclFestivalCoverageRow struct {
	RowID       string
	SourceLabel string
	CatalogIDs  []string
}

var uclFestivalCoverageManifest = []uclFestivalCoverageRow{
	{"old-kingdom-opening-year", "Old Kingdom list: Opening of the Year", []string{"egypt-wepet-renpet"}},
	{"old-kingdom-thoth", "Old Kingdom list: Thoth festival", []string{"egypt-thoth-old-kingdom"}},
	{"old-kingdom-first-year", "Old Kingdom list: First of the Year", []string{"egypt-first-of-year-old-kingdom"}},
	{"old-kingdom-wag", "Old Kingdom list: Wag festival", []string{"egypt-wag-festival"}},
	{"old-kingdom-sokar", "Old Kingdom list: Sokar festival", []string{"egypt-sokar-old-kingdom"}},
	{"old-kingdom-great", "Old Kingdom list: Great Festival", []string{"egypt-great-festival-old-kingdom"}},
	{"old-kingdom-flame", "Old Kingdom list: Flame festival", []string{"egypt-flame-festival-old-kingdom"}},
	{"old-kingdom-min", "Old Kingdom list: Procession of Min", []string{"egypt-min-festival"}},
	{"old-kingdom-sadj", "Old Kingdom list: Sadj festival", []string{"egypt-sadj-old-kingdom"}},
	{"i-akhet-01", "I Akhet 1: New Year / Opening of Year / Ra-Horakhty", []string{"egypt-wepet-renpet"}},
	{"i-akhet-15", "I Akhet 15: offerings to Hapy and Amun", []string{"egypt-hapy-amun-flood-offerings"}},
	{"i-akhet-17", "I Akhet 17: eve of Wag", []string{"egypt-wag-festival"}},
	{"i-akhet-18", "I Akhet 18: Wag", []string{"egypt-wag-festival"}},
	{"i-akhet-19", "I Akhet 19: Wag and Thoth", []string{"egypt-wag-festival"}},
	{"i-akhet-20", "I Akhet 20: Tekh", []string{"egypt-tekh-drunkenness"}},
	{"i-akhet-22", "I Akhet 22: Great Procession of Osiris", []string{"egypt-great-procession-osiris"}},
	{"ii-akhet-15", "II Akhet 15: 11-day Ipet under Thutmose III", []string{"egypt-opet"}},
	{"ii-akhet-19", "II Akhet 19: 27-day Ipet under Ramesses III", []string{"egypt-opet"}},
	{"ii-akhet-18", "II Akhet 18: Khnum and Anuqet at Elephantine", []string{"egypt-elephantine-khnum-anuqet"}},
	{"ii-akhet-27", "II Akhet 27: two-day Mont event", []string{"egypt-mont-middle-kingdom"}},
	{"ii-akhet-28", "II Akhet 28: Satet and Anuqet at Elephantine", []string{"egypt-elephantine-satet-anuqet"}},
	{"ii-akhet-undated", "II Akhet, day unspecified: Menkhet", []string{"egypt-menkhet-middle-kingdom"}},
	{"iii-akhet-09", "III Akhet 9: Amun at Elephantine", []string{"egypt-elephantine-amun"}},
	{"iii-akhet-30", "III Akhet 30: Anuqet at Elephantine", []string{"egypt-elephantine-anuqet"}},
	{"iv-akhet-01", "IV Akhet 1: Hathor at Medinet Habu", []string{"egypt-hathor-medinet-habu"}},
	{"iv-akhet-18", "IV Akhet 18: start Khoiak ceremonies", []string{"egypt-khoiak-osiris"}},
	{"iv-akhet-22", "IV Akhet 22: Ploughing the Earth", []string{"egypt-khoiak-osiris"}},
	{"iv-akhet-26", "IV Akhet 26: Sokar festival", []string{"egypt-khoiak-osiris"}},
	{"iv-akhet-30", "IV Akhet 30: raising the djed", []string{"egypt-khoiak-osiris"}},
	{"i-peret-01", "I Peret 1: Nehebkau / Beginning of Eternity", []string{"egypt-nehebkau"}},
	{"i-peret-20", "I Peret 20: sailing of Wadjyt", []string{"egypt-sailing-wadjyt"}},
	{"i-peret-29-bast", "I Peret 29: sailing of Bast", []string{"egypt-sailing-bast"}},
	{"i-peret-29-willow", "I Peret 29: Raising the Willow", []string{"egypt-raising-willow"}},
	{"i-peret-30", "I Peret 30: Shesmet / late Mut-of-Isheru variant", []string{"egypt-sailing-shesmet", "egypt-sailing-mut-isheru"}},
	{"ii-peret-01", "II Peret 1: sailing of Anubis", []string{"egypt-sailing-anubis"}},
	{"ii-peret-30", "II Peret 30: Raising Heaven / ished branches", []string{"egypt-raising-heaven-cycle"}},
	{"iii-peret-01", "III Peret 1: Ptah and return in Raising Heaven cycle", []string{"egypt-ptah-thebes-peret", "egypt-raising-heaven-cycle"}},
	{"iii-peret-21", "III Peret 21: Amenhotep I in the valley", []string{"egypt-amenhotep-i-valley"}},
	{"iii-peret-29", "III Peret 29: four-day Amenhotep I workforce festival", []string{"egypt-amenhotep-i-workforce"}},
	{"iv-peret-04", "IV Peret 4: Bast / chewing onions", []string{"egypt-bast-onion"}},
	{"iv-peret-05", "IV Peret 5: appearance of Bast in her boat", []string{"egypt-bast-boat-appearance"}},
	{"iv-peret-25", "IV Peret 25: harvest offering to Renenutet", []string{"egypt-renenutet-harvest-offering"}},
	{"iv-peret-27", "IV Peret 27: granary offering to Renenutet", []string{"egypt-renenutet-granary-offering"}},
	{"i-shemu-01", "I Shemu 1: Renenutet / birthday of Nepri", []string{"egypt-renenutet-nepri"}},
	{"i-shemu-10", "I Shemu 10: adoration of Anubis", []string{"egypt-anubis-adoration"}},
	{"i-shemu-11", "I Shemu 11: four-day Min festival at new moon", []string{"egypt-min-festival"}},
	{"ii-shemu-new-moon", "II Shemu new moon: Festival of the Valley", []string{"egypt-beautiful-valley"}},
	{"iii-shemu-15", "III Shemu 15: offerings to Hapy and Amun", []string{"egypt-hapy-amun-flood-offerings"}},
	{"iii-shemu-30", "III Shemu 30: eve of Hathor at Thebes", []string{"egypt-hathor-eve-thebes"}},
	{"iv-shemu-01-02", "IV Shemu 1–2: unnamed festival", []string{"egypt-unnamed-iv-shemu-festival"}},
	{"iv-shemu-02", "IV Shemu 2: Ipip festival", []string{"egypt-ipip-festival"}},
	{"iv-shemu-24", "IV Shemu 24: Ptah festival in pyramid inscription", []string{"egypt-ptah-middle-kingdom-pyramid"}},
	{"iv-shemu-30", "IV Shemu 30: eve of start of year", []string{"egypt-eve-opening-year"}},
	{"epagomenal-01", "First day over year: birthday of Osiris", []string{"egypt-epagomenal-osiris"}},
	{"epagomenal-02", "Second day over year: birthday of Horus", []string{"egypt-epagomenal-horus"}},
	{"epagomenal-03", "Third day over year: birthday of Seth", []string{"egypt-epagomenal-seth"}},
	{"epagomenal-04", "Fourth day over year: birthday of Isis", []string{"egypt-epagomenal-isis"}},
	{"epagomenal-05", "Fifth day over year: birthday of Nephthys", []string{"egypt-epagomenal-nephthys"}},
}

func init() {
	ancientCatalogRecords = append(ancientCatalogRecords, uclEgyptExpansionRecords...)
}
