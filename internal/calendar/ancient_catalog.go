package calendar

import "time"

const (
	egyptianFestivalSource = "https://www.ucl.ac.uk/museums-static/digitalegypt/ideology/festivaldates.html"
	egyptianOpetSource     = "https://escholarship.org/uc/item/4739r3fr"
	egyptianSedSource      = "https://www.ucl.ac.uk/museums-static/digitalegypt/ideology/sed/index.html"
	babylonAkituSource     = "https://www.metmuseum.org/essays/babylon"
	mesopotamianBTTO       = "https://oracc.museum.upenn.edu/btto/Q004806"
	urIIIAkitiSource       = "https://cdli.earth/articles/cdlj/2017-1"
	urIIICultSource        = "https://cdli.earth/articles/cdlp/9.0.pdf"
	urIIIBabaSource        = "https://isac.uchicago.edu/sites/default/files/uploads/shared/docs/ois8.pdf"
	urukRitualSource       = "https://oracc.museum.upenn.edu/cams/gkab/P363711"
	esheshuSource          = "https://oracc.museum.upenn.edu/epsd2/cbd/sux/o0027345.html"
	ugaritRitualSource     = "https://isac.uchicago.edu/sites/default/files/uploads/shared/docs/Publications/SAOC/saoc73.pdf"
	ugaritPardeeSource     = "https://catalog.folger.edu/record/269276"
	heimskringlaSource     = "https://vsnr.org/wp-content/uploads/2021/11/Heimskringla-I.pdf"
	orkneyingaSource       = "https://www.gutenberg.org/ebooks/57723"
	alfablotSource         = "https://skaldic.org/m.php?i=223726&p=wordtextlp"
	bedeMonthsSource       = "https://www.tha.de/~harsch/Chronologia/Lspost08/Bede/bed_ra15.html"
	adfCalendarSource      = "https://host.adf.org/about/org/constitution.html"
)

type ancientCatalogRecord struct {
	ID               string
	Name             string
	Communities      []string
	Category         string
	Summary          string
	Meaning          string
	AttestedElements []string
	Texts            []string
	Origin           string
	Status           string
	HistoricalNote   string
	CalendarCorpus   string
	NativeDateLabel  string
	AttestationLayer string
	Era              string
	Site             string
	ProjectionKind   string
	ProjectionStatus string
	DateConfidence   string
	AnchorLocation   string
	DayBoundary      string
	SourceName       string
	SourceURL        string
}

// catalogOnlyAncientObservances returns attested historical calendar records
// whose native date cannot responsibly be collapsed into a recurring Gregorian
// day. They are discoverable in the catalog, but deliberately never injected
// into a modern day cell.
func catalogOnlyAncientObservances() []Observance {
	events := make([]Observance, 0, len(ancientCatalogRecords))
	for _, record := range ancientCatalogRecords {
		practices := record.AttestedElements
		if practices == nil {
			practices = []string{}
		}
		status := record.Status
		if status == "" {
			status = "Historical record; no continuous modern observance asserted"
		}
		events = append(events, Observance{
			ID:               record.ID,
			CatalogID:        record.ID,
			Name:             record.Name,
			Tradition:        AncientWorld,
			Communities:      record.Communities,
			Category:         record.Category,
			Summary:          record.Summary,
			Meaning:          record.Meaning,
			Practices:        practices,
			Scripture:        record.Texts,
			DateNote:         "Catalog-only native-calendar record. " + record.ProjectionStatus,
			Origin:           record.Origin,
			ObservanceStatus: status,
			Historical:       true,
			HistoricalNote:   record.HistoricalNote,
			DateCertainty:    record.DateConfidence,
			CalendarCorpus:   record.CalendarCorpus,
			NativeDateLabel:  record.NativeDateLabel,
			AttestationLayer: record.AttestationLayer,
			Era:              record.Era,
			Site:             record.Site,
			ProjectionKind:   record.ProjectionKind,
			ProjectionStatus: record.ProjectionStatus,
			DateConfidence:   record.DateConfidence,
			AnchorLocation:   record.AnchorLocation,
			DayBoundary:      record.DayBoundary,
			CatalogOnly:      true,
			SourceName:       record.SourceName,
			SourceURL:        record.SourceURL,
		})
	}
	return events
}

var ancientCatalogRecords = []ancientCatalogRecord{
	// Egyptian evidence. The civil calendar had twelve 30-day months plus five
	// added days and drifted through the seasons without a leap day. Temple and
	// lunar festivals also varied by place and reign.
	{
		ID: "egypt-wepet-renpet", Name: "Wepet Renpet · Opening of the Year", Communities: []string{"Ancient Egyptian temple and court communities"}, Category: "Egyptian new-year festival",
		Summary: "Opening of the Year marked day one of the Egyptian civil year; festival lists also associate it with the birth of Ra-Horakhty.", Meaning: "The boundary renewed ordered time, kingship, and the cult calendar.",
		AttestedElements: []string{"Temple offerings", "New-year rites", "Royal and solar renewal"}, Texts: []string{"Old Kingdom tomb-chapel festival lists", "Medinet Habu festival list"},
		Origin: "Egyptian festival lists", HistoricalNote: "The 365-day civil calendar had no regular leap day, so this date moved through the solar seasons.",
		CalendarCorpus: "Egyptian civil and temple calendars", NativeDateLabel: "I Akhet 1; after the five epagomenal days", AttestationLayer: "Old Kingdom festival lists with later temple-calendar witnesses", Era: "Old Kingdom onward; details vary by period", Site: "Multiple Egyptian sites",
		ProjectionKind: "native-calendar-only", ProjectionStatus: "Not assigned a Gregorian anniversary: the civil year drifted and an absolute conversion requires a specified regnal year.", DateConfidence: "High native date; no year-independent Gregorian date", AnchorLocation: "Egypt; local cult forms varied", DayBoundary: "Dawn/civil day", SourceName: "UCL Digital Egypt · Festivals in the ancient Egyptian calendar", SourceURL: egyptianFestivalSource,
	},
	{
		ID: "egypt-wag-festival", Name: "Wag Festival", Communities: []string{"Ancient Egyptian mortuary and temple communities"}, Category: "Egyptian mortuary festival",
		Summary: "Festival lists place the Wag observance near the start of the inundation-season calendar, in connection with mortuary commemoration and offerings.", Meaning: "The festival joined the renewed year to care for the dead and the maintenance of cult memory.",
		AttestedElements: []string{"Festival eve", "Offerings for the dead", "Processional or mortuary rites varying by period"}, Texts: []string{"Old Kingdom festival lists", "Medinet Habu festival list"},
		Origin: "Egyptian tomb and temple festival lists", HistoricalNote: " Surviving lists are uneven and do not prove identical national practice.",
		CalendarCorpus: "Egyptian civil and temple calendars", NativeDateLabel: "I Akhet 17 eve; Wag on I Akhet 18; combined Wag/Thoth listing on day 19", AttestationLayer: "Old Kingdom tomb lists and New Kingdom Medinet Habu list", Era: "Old Kingdom through New Kingdom witnesses", Site: "Memphite and Theban evidence among other sites",
		ProjectionKind: "native-calendar-only", ProjectionStatus: "The native month-days are retained; no floating civil date is converted without a target regnal year.", DateConfidence: "Moderate to high native sequence; practice varied", AnchorLocation: "Egypt", DayBoundary: "Dawn/civil day; an eve is explicitly listed", SourceName: "UCL Digital Egypt · Festivals in the ancient Egyptian calendar", SourceURL: egyptianFestivalSource,
	},
	{
		ID: "egypt-great-procession-osiris", Name: "Great Procession of Osiris", Communities: []string{"Ancient Egyptian Osirian temple communities"}, Category: "Egyptian divine procession",
		Summary: "An early-year festival list records a Great Procession associated with Osiris; later Abydos evidence preserves especially rich Osirian processional traditions.", Meaning: "Procession made divine presence public and connected temple, landscape, kingship, and the cult of the dead.",
		AttestedElements: []string{"Sacred-bark procession", "Offerings", "Movement through a cult landscape"}, Texts: []string{"Egyptian festival lists", "Abydos Osirian processional evidence"},
		Origin: "Egyptian temple and processional records", HistoricalNote: " Different Osirian processions and later Khoiak rites must not be silently merged into one timeless national ceremony.",
		CalendarCorpus: "Egyptian civil and temple calendars", NativeDateLabel: "I Akhet 22 in the cited linear festival list", AttestationLayer: "Festival-list notice interpreted alongside site-specific processional evidence", Era: "Pharaonic Egypt; local forms changed over time", Site: "Abydos and other Osirian cult centers",
		ProjectionKind: "native-calendar-only", ProjectionStatus: "Cataloged by its received Egyptian month-day, not assigned a universal Gregorian equivalent.", DateConfidence: "Moderate; festival identity and local sequence vary", AnchorLocation: "Abydos, Egypt", DayBoundary: "Dawn/civil day", SourceName: "UCL Digital Egypt · Festivals in the ancient Egyptian calendar", SourceURL: egyptianFestivalSource,
	},
	{
		ID: "egypt-opet", Name: "Opet Festival", Communities: []string{"Ancient Theban temple, court, and festival communities"}, Category: "Egyptian state and temple festival",
		Summary: "Amun's bark, later joined by those of Mut and Khonsu, traveled from Karnak to Luxor in a central Theban festival tied to the renewal of royal ka and kingship.", Meaning: "Opet renewed the relationship among the gods, the king, and Thebes through a major public and temple procession.",
		AttestedElements: []string{"Bark procession from Karnak to Luxor", "Offerings, music, and festival booths", "Royal participation and renewal"}, Texts: []string{"Thutmose III festival list", "Karnak and Luxor relief cycles", "Papyrus Harris I", "Medinet Habu festival list"},
		Origin: "New Kingdom Theban inscriptions and reliefs", HistoricalNote: " Earliest conclusive attestation is Dynasty 18; its start and length changed by reign, from 11 days under Thutmose III to longer Ramesside forms.",
		CalendarCorpus: "Egyptian civil and temple calendars", NativeDateLabel: "II Akhet 15 for 11 days under Thutmose III; II Akhet 19 for 24–27 days under Ramesses III", AttestationLayer: "Dynasty 18–20 monumental and papyrus witnesses", Era: "New Kingdom onward", Site: "Karnak and Luxor, Thebes",
		ProjectionKind: "regnal-variant-native-date", ProjectionStatus: "No single date is projected because the documented schedule changed by reign and the civil calendar drifted.", DateConfidence: "High for cited reign-specific schedules", AnchorLocation: "Thebes (Luxor), Egypt", DayBoundary: "Dawn/civil day; festival eves also attested", SourceName: "UCLA Encyclopedia of Egyptology · Opet Festival", SourceURL: egyptianOpetSource,
	},
	{
		ID: "egypt-khoiak-osiris", Name: "Khoiak Ceremonies · Osiris, Sokar, and Raising the Djed", Communities: []string{"Ancient Egyptian Osirian temple communities"}, Category: "Egyptian multi-day mystery and renewal cycle",
		Summary: "A multi-day sequence centered on the death and regeneration of Osiris included earth-working, the Sokar festival, and the raising of the djed pillar.", Meaning: "The rites expressed renewal through death, germination, restored stability, and the reconstitution of Osiris.",
		AttestedElements: []string{"Day 18 opening rites", "Day 22 Ploughing the Earth", "Day 26 Sokar festival", "Day 30 Raising the Djed"}, Texts: []string{"Abydos Middle Kingdom evidence", "Later temple calendars and Khoiak ritual texts"},
		Origin: "Osirian processional and temple corpora", HistoricalNote: " Middle Kingdom Abydos supports an Osirian procession, while the most detailed surviving Khoiak programs are substantially later; the layers are identified rather than flattened.",
		CalendarCorpus: "Egyptian civil and temple calendars", NativeDateLabel: "IV Akhet 18–30 in the synthesized festival sequence", AttestationLayer: "Middle Kingdom Abydos processional evidence plus later detailed temple programs", Era: "Middle Kingdom through Ptolemaic and Roman periods", Site: "Abydos, Dendera, and other Osirian cult centers",
		ProjectionKind: "layered-native-calendar", ProjectionStatus: "Retained as a native-date range; the exact program depends on period and temple, and no year-independent Gregorian conversion exists.", DateConfidence: "High for later sequence; moderate when extended backward", AnchorLocation: "Abydos, Egypt", DayBoundary: "Dawn/civil day", SourceName: "UCL Digital Egypt · Festivals in the ancient Egyptian calendar", SourceURL: egyptianFestivalSource,
	},
	{
		ID: "egypt-nehebkau", Name: "Festival of Nehebkau · Beginning of Eternity", Communities: []string{"Ancient Egyptian temple communities"}, Category: "Egyptian temple festival",
		Summary: "The first day of the sowing season's first civil month was marked for Nehebkau and called the Beginning of Eternity in a Sety I inscription.", Meaning: "The day linked a calendrical turning point with protection, continuity, and renewed divine order.",
		AttestedElements: []string{"Temple offerings", "Calendrical renewal"}, Texts: []string{"Sety I inscription at Nauri"},
		Origin: "Nauri decree and festival-calendar synthesis", HistoricalNote: " A named inscriptional witness does not establish identical observance throughout Egypt.",
		CalendarCorpus: "Egyptian civil and temple calendars", NativeDateLabel: "I Peret 1", AttestationLayer: "Dynasty 19 inscriptional witness", Era: "New Kingdom, Dynasty 19 attestation", Site: "Nauri; broader cult context uncertain",
		ProjectionKind: "native-calendar-only", ProjectionStatus: "The civil month-day is retained without a Gregorian projection.", DateConfidence: "High for the cited inscription's native date", AnchorLocation: "Nauri, Nubia/Egyptian imperial context", DayBoundary: "Dawn/civil day", SourceName: "UCL Digital Egypt · Festivals in the ancient Egyptian calendar", SourceURL: egyptianFestivalSource,
	},
	{
		ID: "egypt-renenutet-nepri", Name: "Renenutet Festival · Birthday of Nepri", Communities: []string{"Ancient Egyptian agricultural and temple communities"}, Category: "Egyptian harvest festival",
		Summary: "A first-day harvest-season festival honored Renenutet and was also identified with the birthday of Nepri, the personification of grain.", Meaning: "The festival joined the grain harvest to divine nourishment, storage, and protection.",
		AttestedElements: []string{"Harvest offerings", "Temple celebration of Renenutet and Nepri"}, Texts: []string{"Egyptian festival calendars"},
		Origin: "Egyptian temple calendars", HistoricalNote: " Agricultural timing and month naming are local and historical; the civil calendar's drift prevents a permanent modern harvest date.",
		CalendarCorpus: "Egyptian civil and temple calendars", NativeDateLabel: "I Shemu 1", AttestationLayer: "Festival-calendar witness", Era: "Pharaonic Egypt; precise horizon varies by source", Site: "Egyptian agricultural cult centers",
		ProjectionKind: "native-calendar-only", ProjectionStatus: "Native harvest-season label retained; no universal Gregorian harvest date is asserted.", DateConfidence: "Moderate; attested date within a synthesized calendar", AnchorLocation: "Egypt", DayBoundary: "Dawn/civil day", SourceName: "UCL Digital Egypt · Festivals in the ancient Egyptian calendar", SourceURL: egyptianFestivalSource,
	},
	{
		ID: "egypt-min-festival", Name: "Festival and Procession of Min", Communities: []string{"Ancient Theban and Coptite temple communities"}, Category: "Egyptian lunar and agricultural festival",
		Summary: "The Medinet Habu list gives Min a four-day festival beginning at new moon, while older festival lists also name a procession of Min.", Meaning: "Min's procession connected divine fertility, kingship, and the agricultural season.",
		AttestedElements: []string{"Four-day festival", "New-moon timing", "Divine procession", "Royal participation in New Kingdom imagery"}, Texts: []string{"Old Kingdom festival lists", "Ramesses III festival list at Medinet Habu"},
		Origin: "Egyptian festival lists and reliefs", HistoricalNote: "The new-moon rule and civil day number belong to a particular New Kingdom synthesis and should not be universalized.",
		CalendarCorpus: "Egyptian civil and temple calendars", NativeDateLabel: "I Shemu day 11; four days at the new moon in the Medinet Habu list", AttestationLayer: "Early festival-name attestations and Ramesside dated program", Era: "Old Kingdom name attestations; detailed New Kingdom program", Site: "Coptos and Thebes, especially Medinet Habu",
		ProjectionKind: "native-lunar-rule", ProjectionStatus: "Requires both a specified ancient civil year and a reconstruction of the local lunar calendar; not projected automatically.", DateConfidence: "High for the Ramesside list; wider chronology uncertain", AnchorLocation: "Thebes, Egypt", DayBoundary: "Dawn/civil day with lunar trigger", SourceName: "UCL Digital Egypt · Festivals in the ancient Egyptian calendar", SourceURL: egyptianFestivalSource,
	},
	{
		ID: "egypt-beautiful-valley", Name: "Beautiful Festival of the Valley", Communities: []string{"Ancient Theban temple and family communities"}, Category: "Egyptian procession and festival of the dead",
		Summary: "At the new moon Amun traveled from Karnak to west-bank royal cult temples while families feasted with and remembered their dead.", Meaning: "The festival joined public divine procession with family remembrance across Thebes's living and mortuary landscapes.",
		AttestedElements: []string{"Amun bark procession across the Nile", "Flowers and festive banquets", "Visits to tomb chapels and royal cult temples"}, Texts: []string{"Theban tomb and temple scenes", "Festival calendars"},
		Origin: "Theban monumental and funerary evidence", HistoricalNote: " This is a specifically Theban festival whose date and scale changed over its long history.",
		CalendarCorpus: "Egyptian civil and temple calendars", NativeDateLabel: "New moon in II Shemu; exact civil day varies by witness", AttestationLayer: "Middle and New Kingdom Theban monumental evidence with later continuities", Era: "Middle Kingdom roots; especially well attested in the New Kingdom", Site: "Karnak, Deir el-Bahri, and the Theban west bank",
		ProjectionKind: "local-lunar-rule", ProjectionStatus: "A projection needs a selected historical year, Theban lunar model, and source layer; none is silently chosen.", DateConfidence: "High festival identity; variable annual day", AnchorLocation: "Thebes (Luxor), Egypt", DayBoundary: "Dawn/civil day with new-moon trigger", SourceName: "UCL Digital Egypt · Festivals in the ancient Egyptian calendar", SourceURL: egyptianFestivalSource,
	},
	{
		ID: "egypt-menkhet-middle-kingdom", Name: "Menkhet Festival · Middle Kingdom Account", Communities: []string{"Middle Kingdom Egyptian court or temple personnel"}, Category: "Egyptian locally attested festival",
		Summary: "A late Middle Kingdom account names a Menkhet or Cloth festival while listing singers and dancers, but does not preserve a secure annual day.", Meaning: "The record shows a performed festival economy without supplying enough evidence for a universal calendar placement.",
		AttestedElements: []string{"Singers and dancers recorded in an administrative account"}, Texts: []string{"UCL Petrie Museum UC 32191"},
		Origin: "Late Middle Kingdom administrative papyrus", HistoricalNote: " This is exactly the kind of period evidence that should remain visible without being promoted into a fixed national holiday.",
		CalendarCorpus: "Egyptian civil and temple calendars", NativeDateLabel: "Month II Akhet context; day not specified", AttestationLayer: "Single late Middle Kingdom administrative account", Era: "Late Middle Kingdom, approximately 2025–1700 BCE", Site: "Provenance/context of UC 32191; Egyptian court administration",
		ProjectionKind: "attested-without-date", ProjectionStatus: "Catalog-only because the source does not establish an annual day or national recurrence.", DateConfidence: "Festival name attested; calendar placement low", AnchorLocation: "Egypt; exact local anchor unresolved", DayBoundary: "Not preserved", SourceName: "UCL Digital Egypt · Festivals in the ancient Egyptian calendar", SourceURL: egyptianFestivalSource,
	},
	{
		ID: "egypt-mont-middle-kingdom", Name: "Festival of Mont · Boulaq 18 Account", Communities: []string{"Middle Kingdom Theban court and temple personnel"}, Category: "Egyptian possibly nonannual ceremony",
		Summary: "Papyrus Boulaq 18 records a two-day event called the festival of Mont, perhaps a local festival or a one-time shrine consecration.", Meaning: "The entry preserves the source's ambiguity instead of manufacturing a recurring holiday.",
		AttestedElements: []string{"Two-day court or temple provisioning", "Possible shrine ceremony"}, Texts: []string{"Papyrus Boulaq 18"},
		Origin: "Late Middle Kingdom accounts", Status: "Historical event; annual recurrence uncertain", HistoricalNote: "The source may describe one ceremony rather than an annual feast.",
		CalendarCorpus: "Egyptian civil and temple calendars", NativeDateLabel: "II Akhet 27–28 in the account", AttestationLayer: "Single late Middle Kingdom administrative source", Era: "Late Middle Kingdom, approximately 2025–1700 BCE", Site: "Theban administrative context",
		ProjectionKind: "possibly-one-off-native-date", ProjectionStatus: "Not projected because recurrence, regnal year, and the event's exact character are unresolved.", DateConfidence: "High source date; low confidence in annual recurrence", AnchorLocation: "Thebes, Egypt", DayBoundary: "Dawn/civil day", SourceName: "UCL Digital Egypt · Festivals in the ancient Egyptian calendar", SourceURL: egyptianFestivalSource,
	},
	{
		ID: "egypt-heb-sed-jubilee", Name: "Heb Sed · Royal Jubilee", Communities: []string{"Ancient Egyptian royal court and temple communities"}, Category: "Egyptian regnal jubilee",
		Summary: "The Sed festival renewed royal power and kingship through a major jubilee cycle attested from the First Dynasty; many well-documented kings held a first Sed around regnal year 30, but the rule had exceptions.", Meaning: "The rites represented the renewal of the king's physical and supernatural capacity to sustain ordered rule.",
		AttestedElements: []string{"Royal jubilee ceremonies", "Special Sed-festival dress and enthronement imagery", "Ritual movement and offerings reconstructed from relief cycles"}, Texts: []string{"Early Dynastic labels", "Old Kingdom Sed-festival relief cycles", "Later royal regnal records"},
		Origin: "Royal inscriptions and pictorial festival cycles", HistoricalNote: " A convenient 'thirty-year festival' label describes an important pattern, not an invariant rule: rulers celebrated at different intervals and the surviving sequence is reconstructed from partial scenes.",
		CalendarCorpus: "Egyptian civil and temple calendars", NativeDateLabel: "Regnal jubilee, often first attested around regnal year 30–31; later repetitions and exceptions vary by king", AttestationLayer: "First Dynasty onward; especially rich Old and New Kingdom royal evidence", Era: "Early Dynastic through later pharaonic periods", Site: "Royal cult centers across Egypt; programs differed by reign",
		ProjectionKind: "regnal-event-only", ProjectionStatus: "Not assigned an annual or Gregorian date: a projection requires an identified king, regnal chronology, and a specific attested celebration.", DateConfidence: "High festival identity; reign-specific timing", AnchorLocation: "Reign and royal cult site specific", DayBoundary: "Festival sequence varies; no universal civil-day boundary supplied", SourceName: "UCL Digital Egypt · The Sed Festival", SourceURL: egyptianSedSource,
	},

	// Mesopotamian calendars were city-based lunisolar systems. Month names,
	// intercalations, cult owners, and even New Year differed by place and era.
	{
		ID: "babylon-akitu-marduk", Name: "Akītu of Marduk · Babylonian New Year", Communities: []string{"First-millennium BCE Babylonian temple, court, and urban communities"}, Category: "Babylonian new-year festival",
		Summary: "Babylon's spring Akītu brought Marduk and visiting divine statues through the Processional Way to the Akītu house and renewed the bond among god, king, city, and cosmos.", Meaning: "The cycle reaffirmed divine order and legitimate kingship at the opening of the Babylonian year.",
		AttestedElements: []string{"Multi-day temple rites", "Procession of divine statues", "Journey to the Akītu house", "Royal participation"}, Texts: []string{"First-millennium Babylonian ritual tablets", "Royal inscriptions and topographic evidence"},
		Origin: "Babylonian ritual and royal corpora", HistoricalNote: "Akītu changed across cities and centuries; the familiar Marduk-centered Babylonian program is not a timeless pan-Mesopotamian festival.",
		CalendarCorpus: "Babylonian cult calendar", NativeDateLabel: "Nisannu, opening days of the lunisolar year; the mature program spans roughly 11–12 days", AttestationLayer: "First-millennium BCE ritual tablets, inscriptions, and Babylonian topography", Era: "Especially Neo-Assyrian and Neo-Babylonian evidence", Site: "Babylon",
		ProjectionKind: "city-lunisolar-native-date", ProjectionStatus: "No Gregorian date is assigned without a historical year and Babylonian intercalation/new-crescent reconstruction.", DateConfidence: "High native month and festival; year conversion requires choices", AnchorLocation: "Babylon, Iraq", DayBoundary: "Evening; observational lunar month", SourceName: "The Metropolitan Museum of Art · Babylon", SourceURL: babylonAkituSource,
	},
	{
		ID: "babylon-ishtar-nisannu", Name: "Akītu of Ishtar of Nippur—or Ninurta · BTTo Q004806", Communities: []string{"Babylonian scholarly and temple communities"}, Category: "Babylonian theological festival notice",
		Summary: "A bilingual explanatory tablet places an Akītu in its Nisannu list and glosses it as belonging to Ishtar, Queen of Nippur, or alternatively to Ninurta.", Meaning: "The alternative wording preserves the tablet's learned comparison without forcing one cult owner where the source offers two readings.",
		AttestedElements: []string{"Akītu name in a bilingual scholarly list", "Alternative Ishtar-of-Nippur / Ninurta wording"}, Texts: []string{"Bilingual explanatory tablet BTTo Q004806 iii 15′"},
		Origin: "Babylonian temple-theological tablet", HistoricalNote: "This is a scholarly or theological festival list, not a complete liturgy or proof that one public program was enacted unchanged throughout Babylonia; the Ishtar/Ninurta alternative is retained.",
		CalendarCorpus: "Babylonian cult calendar", NativeDateLabel: "Nisannu; list entry iii 15′, without a numbered day of celebration", AttestationLayer: "First-millennium bilingual explanatory tablet", Era: "First millennium BCE", Site: "Babylonia; tablet context",
		ProjectionKind: "scholarly-list-month-only", ProjectionStatus: "Cataloged at month resolution only; neither an exact ancient day nor a Gregorian recurrence is inferred.", DateConfidence: "High tablet wording and month; cult-owner wording explicitly ambiguous; exact day unavailable", AnchorLocation: "Babylonia", DayBoundary: "Evening; lunisolar month", SourceName: "ORACC · Babylonian Temple Texts Online Q004806", SourceURL: mesopotamianBTTO,
	},
	{
		ID: "ur3-akiti-harvest", Name: "Akiti-šekinku · New Year of Harvest at Ur", Communities: []string{"Ur III court, temple, and administrative communities at Ur"}, Category: "Ur III new-year festival",
		Summary: "Administrative records at Ur identify a harvest Akiti at the start of the local year, with gifts and provisions recorded for participants and visitors.", Meaning: "The feast joined agricultural transition, the city's cult, and the redistributive work of the Ur III state.",
		AttestedElements: []string{"Festival provisions", "Livestock gifts", "Court and diplomatic participation"}, Texts: []string{"Ur III administrative tablets"},
		Origin: "Ur III administrative archives", HistoricalNote: "This Ur calendar must not be replaced with the later standard Babylonian month system; local calendars and month names are essential evidence.",
		CalendarCorpus: "Ur III administrative cult calendars", NativeDateLabel: "Month of Harvest (iti še-sag₁₁-ku₅), beginning of Ur's year", AttestationLayer: "Dated Ur III administrative tablets", Era: "Ur III, late third millennium BCE", Site: "Ur",
		ProjectionKind: "city-lunisolar-native-date", ProjectionStatus: "No Gregorian date is assigned without a selected Ur III regnal year and reconstruction of Ur's local lunisolar calendar.", DateConfidence: "High within dated administrative records", AnchorLocation: "Ur, Iraq", DayBoundary: "Evening; lunisolar month", SourceName: "Cuneiform Digital Library Journal · Ur III diplomatic and festival records", SourceURL: urIIIAkitiSource,
	},
	{
		ID: "ur3-akiti-sowing", Name: "Akiti-šununum · New Year of Sowing at Ur", Communities: []string{"Ur III court, temple, and administrative communities at Ur"}, Category: "Ur III midyear/new-year festival",
		Summary: "Ur's second Akiti opened the seventh month and was called the New Year of Sowing or Middle Year in administrative documentation.", Meaning: "The observance marked the agricultural and administrative hinge between harvest and sowing cycles.",
		AttestedElements: []string{"Festival provisions", "Livestock gifts", "Court and diplomatic participation"}, Texts: []string{"Ur III administrative tablets"},
		Origin: "Ur III administrative archives", HistoricalNote: "Calling both Ur festivals 'New Year' reflects the source terminology and a dual agricultural structure, not two Gregorian January-style holidays.",
		CalendarCorpus: "Ur III administrative cult calendars", NativeDateLabel: "First day of Ur month VII, the Akiti month; a₂-ki-ti šu-numun", AttestationLayer: "Dated Ur III administrative tablets", Era: "Ur III, late third millennium BCE", Site: "Ur",
		ProjectionKind: "city-lunisolar-native-date", ProjectionStatus: "Requires a specified Ur III year and local intercalation model; retained only in native form.", DateConfidence: "High within dated administrative records", AnchorLocation: "Ur, Iraq", DayBoundary: "Evening; lunisolar month", SourceName: "Cuneiform Digital Library Journal · Ur III diplomatic and festival records", SourceURL: urIIIAkitiSource,
	},
	{
		ID: "ur3-esheshu-lunar", Name: "Ešeš / Eššešu · Recurring Lunar Festival Offerings", Communities: []string{"Ur III Mesopotamian temple and palace communities"}, Category: "Mesopotamian recurring lunar sacred days",
		Summary: "Hundreds of Ur III records attest ešeš festival offerings, supplied two to four times in a month at lunar phases in local temple economies.", Meaning: "The recurring cycle tied cult provisioning to the visible lunar month rather than to one annual solar anniversary.",
		AttestedElements: []string{"Regular animal and food offerings", "New-, quarter-, or full-moon scheduling depending on local record"}, Texts: []string{"Ur III administrative tablets", "ePSD2 ešeš corpus"},
		Origin: "Sumerian administrative and lexical corpus", HistoricalNote: "The term's 556 attestations span periods, but exact phase schedules and recipients varied; this entry describes a recurring class, not one universal feast day.",
		CalendarCorpus: "Mesopotamian lunar cult calendars", NativeDateLabel: "Two to four ešeš days per lunar month; local phase schedule", AttestationLayer: "Primarily Ur III administrative records, with earlier and later attestations", Era: "Old Akkadian through Neo-Babylonian; strongest Ur III corpus", Site: "Ur, Girsu, Nippur, and other cities",
		ProjectionKind: "recurring-local-lunar-rule", ProjectionStatus: "Not rendered as fixed modern moon phases because each archive's schedule and observational convention must be modeled separately.", DateConfidence: "High existence; variable local phase mapping", AnchorLocation: "City-specific Mesopotamian observation", DayBoundary: "Evening; observational lunar cycle", SourceName: "ORACC ePSD2 · ešeš festival", SourceURL: esheshuSource,
	},
	{
		ID: "ur3-baba-ushim", Name: "U₂-šim Festival of Baba · Girsu", Communities: []string{"Ur III Girsu temple and administrative communities"}, Category: "Ur III local deity festival",
		Summary: "Ur III offering records place a u₂-šim festival of the goddess Baba in Girsu's eleventh local month.", Meaning: "The dossier preserves a local cult calendar and its offerings, distinct from calendars centered on Ur or later Babylon.",
		AttestedElements: []string{"Temple offerings", "Commemoration of deceased rulers within the later Ur III offering system"}, Texts: []string{"Ur III Girsu administrative tablets"},
		Origin: "Girsu administrative cult records", HistoricalNote: "The event is locally attested; neither its month nor its ritual detail should be universalized across Sumer.",
		CalendarCorpus: "Ur III administrative cult calendars", NativeDateLabel: "Girsu month XI; u₂-šim festival of Baba", AttestationLayer: "Later Ur III offering accounts", Era: "Ur III, late third millennium BCE", Site: "Girsu",
		ProjectionKind: "city-lunisolar-month-only", ProjectionStatus: "Cataloged in the Girsu calendar only; no exact day or Gregorian equivalent is supplied.", DateConfidence: "Moderate to high local attestation; day unresolved", AnchorLocation: "Girsu, Iraq", DayBoundary: "Evening; local lunisolar calendar", SourceName: "Institute for the Study of Ancient Cultures · Gudea and Ur III cult calendar", SourceURL: urIIIBabaSource,
	},
	{
		ID: "uruk-anu-antu-akitu", Name: "Akītu of Anu and Antu · Late Uruk", Communities: []string{"Late Babylonian and Seleucid Uruk temple communities"}, Category: "Uruk temple festival",
		Summary: "A detailed Uruk ritual tablet schedules processional boats, movement through streets, service at the Akītu house, meals, music, and offerings for Anu, Antu, Ishtar, and the divine assembly.", Meaning: "The program renewed Uruk's local divine order through a city-specific learned temple liturgy.",
		AttestedElements: []string{"Street and boat processions", "Akītu-house rites", "Morning and evening meals", "Music and offerings"}, Texts: []string{"Uruk ritual tablet P363711"},
		Origin: "Late Uruk ritual tablet", HistoricalNote: "This Anu-centered Uruk program is not interchangeable with Marduk's Akītu at Babylon, even where vocabulary and month names overlap.",
		CalendarCorpus: "Seleucid Uruk ritual calendar", NativeDateLabel: "Tašrītu sequence opening on day 1; a day-7 passage compares the corresponding Nisannu rite", AttestationLayer: "Late Babylonian/Seleucid learned temple ritual tablet", Era: "Late first millennium BCE", Site: "Uruk",
		ProjectionKind: "city-lunisolar-ritual-sequence", ProjectionStatus: "No Gregorian placement without a selected Babylonian year; incomplete tablet context is not filled by analogy.", DateConfidence: "High tablet sequence; absolute year not inherent", AnchorLocation: "Uruk, Iraq", DayBoundary: "Evening; lunisolar month", SourceName: "ORACC · Greek and Late Babylonian Knowledge, P363711", SourceURL: urukRitualSource,
	},
	{
		ID: "uruk-clothing-procession", Name: "Clothing and Processional Rites of Anu, Antu, and Ishtar", Communities: []string{"Late Uruk temple personnel"}, Category: "Uruk multi-day temple rite",
		Summary: "The Uruk schedule separately records clothing divine statues, cleansing the temple, setting an ox in place, music, street movement, boats, and Akītu-house service.", Meaning: "Retaining the sub-cycle makes the tablet's day-by-day ritual labor visible without pretending it was a separate modern public holiday.",
		AttestedElements: []string{"Clothing divine statues", "Temple cleansing", "Musicians and lamentation priests", "Processional boats"}, Texts: []string{"Uruk ritual tablet P363711"},
		Origin: "Late Uruk ritual tablet", Status: "Historical temple sub-cycle; not asserted as an independent public holiday", HistoricalNote: "This record is included as a calendrical ritual sequence, not inflated into a separately attested annual festival name.",
		CalendarCorpus: "Seleucid Uruk ritual calendar", NativeDateLabel: "Within the tablet's Tašrītu Akītu schedule; exact preserved line sequence", AttestationLayer: "Late Babylonian/Seleucid temple ritual tablet", Era: "Late first millennium BCE", Site: "Uruk",
		ProjectionKind: "source-sequence-only", ProjectionStatus: "Discoverable with its parent Akītu program; not given an invented standalone date.", DateConfidence: "High ritual attestation; independent-holiday status not claimed", AnchorLocation: "Uruk, Iraq", DayBoundary: "Evening; lunisolar month", SourceName: "ORACC · Greek and Late Babylonian Knowledge, P363711", SourceURL: urukRitualSource,
	},

	// Ugaritic ritual tablets preserve royal and temple schedules but not a
	// complete, securely ordered twelve-month calendar. Day sequences are kept
	// native and gaps remain gaps.
	{
		ID: "ugarit-first-wine-cycle", Name: "First-Wine Vintage Festival Cycle", Communities: []string{"Late Bronze Age Ugaritic royal and temple communities"}, Category: "Ugaritic vintage and royal cult festival",
		Summary: "KTU 1.41 and its duplicate KTU 1.87 prescribe a long royal ritual sequence in the month called First of the Wine, beginning with a grape cluster offered to El.", Meaning: "The cycle connected the vintage, royal purity, household and dynastic gods, and Ugarit's major deities.",
		AttestedElements: []string{"Grape-cluster offering to El", "Royal washing and purity", "Sacrifices and likely libations", "Temple and rooftop rites"}, Texts: []string{"KTU 1.41", "KTU 1.87"},
		Origin: "Ugaritic alphabetic ritual tablets", HistoricalNote: "The First-Wine month is often placed near the autumn vintage and sometimes treated as a liturgical new year, but the full Ugaritic twelve-month sequence is not preserved securely.",
		CalendarCorpus: "Ugaritic ritual tablets", NativeDateLabel: "Month of First Wine (rʾš yn); rites from new moon through later numbered days", AttestationLayer: "Late Bronze Age prescriptive royal ritual and duplicate", Era: "Late Bronze Age, approximately 13th century BCE", Site: "Ugarit (Ras Shamra)",
		ProjectionKind: "incomplete-native-lunisolar-calendar", ProjectionStatus: "Not assigned to a Gregorian vintage date; the Ugaritic month order, intercalation, and local crescent anchor remain partly reconstructed.", DateConfidence: "High tablet sequence; moderate seasonal placement", AnchorLocation: "Ugarit (Ras Shamra), Syria", DayBoundary: "Evening; lunisolar calendar", SourceName: "Theodore J. Lewis · God [ʾIlu] and King in KTU 1.23 (SAOC 73)", SourceURL: ugaritRitualSource,
	},
	{
		ID: "ugarit-first-wine-new-moon", Name: "First-Wine New-Moon Opening", Communities: []string{"Late Bronze Age Ugaritic royal and temple communities"}, Category: "Ugaritic lunar festival opening",
		Summary: "The First-Wine ritual opens at the new moon before proceeding through explicitly numbered days of the month.", Meaning: "The opening shows that Ugaritic cultic time was structured by a locally observed lunar month as well as the vintage season.",
		AttestedElements: []string{"New-moon opening", "Royal and temple offerings"}, Texts: []string{"KTU 1.41", "KTU 1.87"},
		Origin: "Ugaritic alphabetic ritual tablets", Status: "Historical sub-cycle within the First-Wine festival", HistoricalNote: "This is a separately discoverable calendar phase within KTU 1.41/1.87, not evidence for a generic Canaanite New Moon holiday practiced identically everywhere.",
		CalendarCorpus: "Ugaritic ritual tablets", NativeDateLabel: "New moon of the First-Wine month", AttestationLayer: "Late Bronze Age prescriptive royal ritual", Era: "Late Bronze Age, approximately 13th century BCE", Site: "Ugarit (Ras Shamra)",
		ProjectionKind: "local-observational-lunar-rule", ProjectionStatus: "Requires a selected ancient year and Ugarit-based crescent model; no modern annual recurrence is generated.", DateConfidence: "High relative placement in the tablet; absolute date unavailable", AnchorLocation: "Ugarit (Ras Shamra), Syria", DayBoundary: "Evening; observational lunar month", SourceName: "Theodore J. Lewis · God [ʾIlu] and King in KTU 1.23 (SAOC 73)", SourceURL: ugaritRitualSource,
	},
	{
		ID: "ugarit-baal-royal-ritual-119", Name: "Royal Rite for Baal of Ugarit · KTU 1.119", Communities: []string{"Late Bronze Age Ugaritic royal and temple communities"}, Category: "Ugaritic royal temple rite",
		Summary: "KTU 1.119 opens in the month ʾIbʿaltu/Ibaalat on day 7 and records royal washing and sacrifices centered on Baal of Ugarit, followed by a mid-month ritual sequence.", Meaning: "The tablet witnesses the king's active role in maintaining the city's cult while preserving relative native dates that still cannot be assigned one annual Gregorian recurrence.",
		AttestedElements: []string{"Royal ritual washing", "Sacrifice for Baal of Ugarit", "Offering at the temple of El"}, Texts: []string{"KTU 1.119"},
		Origin: "Ugaritic alphabetic ritual tablet", HistoricalNote: "The opening month/day is explicit, while the interpretation of a possible second-month layer and the full annual month order remain debated.",
		CalendarCorpus: "Ugaritic ritual tablets", NativeDateLabel: "ʾIbʿaltu/Ibaalat day 7 opening; further rites around mid-month (possible second-month layer debated)", AttestationLayer: "Late Bronze Age royal ritual tablet", Era: "Late Bronze Age, approximately 13th century BCE", Site: "Ugarit (Ras Shamra)",
		ProjectionKind: "partial-native-lunisolar-sequence", ProjectionStatus: "Catalog-only: native month and numbered days are retained, but no Gregorian projection is made without an ancient year and defensible Ugaritic month anchor.", DateConfidence: "High opening month/day; moderate reconstruction of later sequence", AnchorLocation: "Ugarit (Ras Shamra), Syria", DayBoundary: "Evening; lunisolar month", SourceName: "Dennis Pardee · Ritual and Cult at Ugarit (SBL, 2002)", SourceURL: ugaritPardeeSource,
	},
	{
		ID: "ugarit-divine-image-entry-112", Name: "Divine-Image Entry Procession · KTU 1.112", Communities: []string{"Late Bronze Age Ugaritic royal and temple communities"}, Category: "Ugaritic royal processional rite",
		Summary: "KTU 1.112 preserves a reconstructed Ḫiyāru/Hyr monthly sequence, including explicitly numbered days, within Ugarit's royal divine-image entry rites.", Meaning: "The tablet makes royal and divine movement through sacred space visible while preserving a partial monthly structure whose interpretation remains debated.",
		AttestedElements: []string{"Procession of divine images", "Participation of king and royal family"}, Texts: []string{"KTU 1.112"},
		Origin: "Ugaritic alphabetic ritual corpus", Status: "Historical processional rite; partial native schedule retained", HistoricalNote: "Scholarship treats the tablet as a Ḫiyāru/Hyr sequence with named days through at least day 17 (day 14 explicit), but the month reconstruction and ritual interpretation remain debated.",
		CalendarCorpus: "Ugaritic ritual tablets", NativeDateLabel: "Reconstructed month Ḫiyāru/Hyr; numbered sequence through at least day 17, with day 14 explicit", AttestationLayer: "Late Bronze Age royal entry-ritual tablet", Era: "Late Bronze Age, approximately 13th century BCE", Site: "Ugarit (Ras Shamra)",
		ProjectionKind: "partially-reconstructed-native-month", ProjectionStatus: "Catalog-only: the partial native sequence is preserved, but no Gregorian date is projected from a debated month reconstruction.", DateConfidence: "High numbered-day sequence; moderate month reconstruction", AnchorLocation: "Ugarit (Ras Shamra), Syria", DayBoundary: "Evening; lunisolar month", SourceName: "Dennis Pardee · Ritual and Cult at Ugarit (SBL, 2002)", SourceURL: ugaritPardeeSource,
	},

	// Old Norse evidence combines near-contemporary skaldic verse with much
	// later Christian-era saga prose. These records expose that distinction.
	{
		ID: "norse-vetrnaetr-blot", Name: "Vetrnætr · Sacrifice at Winter's Beginning", Communities: []string{"Historical Old Norse communities"}, Category: "Old Norse seasonal blót",
		Summary: "Ynglinga saga describes a sacrifice at winter's beginning for a good year, corresponding to the broader Winter Nights seasonal boundary in saga tradition.", Meaning: "The rite marked entry into winter and sought communal well-being for the coming year.",
		AttestedElements: []string{"Blót or sacrificial feast", "Petitions for a good year"}, Texts: []string{"Ynglinga saga, chapter 8", "Later saga references to Winter Nights"},
		Origin: "Medieval Icelandic saga prose", HistoricalNote: "Snorri wrote centuries after conversion. The seasonal triad may preserve older practice, but it is not a contemporary Viking Age festival manual.",
		CalendarCorpus: "Old Norse textual calendar", NativeDateLabel: "At winter's beginning; Old Icelandic calendar boundary", AttestationLayer: "Thirteenth-century saga account drawing on earlier traditions", Era: "Medieval witness to claimed pre-Christian practice", Site: "Scandinavia; saga geography varies",
		ProjectionKind: "seasonal-native-calendar", ProjectionStatus: "Not fixed to a Gregorian date; Old Norse local calendars and later saga schemes require explicit reconstruction.", DateConfidence: "Moderate seasonal placement; low exact-day confidence", AnchorLocation: "Old Norse regional calendar", DayBoundary: "Evening/night; local reckoning", SourceName: "Viking Society · Heimskringla I", SourceURL: heimskringlaSource,
	},
	{
		ID: "norse-jol-midwinter", Name: "Jól · Midwinter Feast", Communities: []string{"Historical Old Norse communities"}, Category: "Old Norse midwinter festival",
		Summary: "Hákonar saga góða describes an older three-night Jól at midwinter and King Hákon's law aligning its start with Christian Christmas, including obligatory ale brewing.", Meaning: "The account preserves both a midwinter feast and the historical process by which its public timing was Christianized.",
		AttestedElements: []string{"Multi-night feast", "Ale brewing and ritual drinking", "Sacrificial banquet in saga narratives"}, Texts: []string{"Hákonar saga góða", "Ynglinga saga"},
		Origin: "Medieval Icelandic kings' sagas", HistoricalNote: "The source is later and Christian; its 'older' timing is not a license to equate Jól automatically with the astronomical solstice or December 25 in every earlier region.",
		CalendarCorpus: "Old Norse textual calendar", NativeDateLabel: "Midwinter night (hǫkunótt) for three nights in the saga's retrospective account", AttestationLayer: "Thirteenth-century saga prose describing tenth-century reform", Era: "Viking Age setting preserved in High Medieval narrative", Site: "Norway, especially Trøndelag in the saga",
		ProjectionKind: "historical-calendar-transition", ProjectionStatus: "Catalog-only because both the earlier native timing and Hákon's reform require year- and source-specific conversion.", DateConfidence: "High that saga reports a timing reform; older exact date debated", AnchorLocation: "Norway", DayBoundary: "Night beginning the feast", SourceName: "Viking Society · Heimskringla I", SourceURL: heimskringlaSource,
	},
	{
		ID: "norse-sigrblot", Name: "Sigrblót · Summer-Beginning Victory Sacrifice", Communities: []string{"Historical Old Norse communities in saga tradition"}, Category: "Old Norse seasonal blót",
		Summary: "Ynglinga saga's three-part annual scheme places a sacrifice for victory at summer's beginning.", Meaning: "The rite marked the opening of the summer half-year and framed hopes for success in the active season.",
		AttestedElements: []string{"Blót", "Petitions for victory"}, Texts: []string{"Ynglinga saga, chapter 8"},
		Origin: "Medieval Icelandic saga prose", HistoricalNote: "The name and annual scheme are preserved in a late literary source, not a contemporary ritual calendar.",
		CalendarCorpus: "Old Norse textual calendar", NativeDateLabel: "At summer's beginning in the Old Norse seasonal half-year", AttestationLayer: "Thirteenth-century saga account", Era: "Medieval witness to claimed pre-Christian practice", Site: "Scandinavian setting in Ynglinga saga",
		ProjectionKind: "seasonal-native-calendar", ProjectionStatus: "No Gregorian day is chosen for the regional beginning of summer.", DateConfidence: "Moderate seasonal placement; low exact-day confidence", AnchorLocation: "Old Norse regional calendar", DayBoundary: "Local day/night boundary", SourceName: "Viking Society · Heimskringla I", SourceURL: heimskringlaSource,
	},
	{
		ID: "norse-alfablot", Name: "Álfablót · Sacrificial Feast for the Elves", Communities: []string{"Historical Swedish households encountered by Sigvatr"}, Category: "Old Norse household sacrifice",
		Summary: "Sigvatr's near-contemporary Austrfararvísur recounts being refused hospitality because households were holding an álfablót.", Meaning: "The verses attest a private household rite and, unusually, the social boundary around it.",
		AttestedElements: []string{"Private household sacrificial feast", "Exclusion of an outside traveler"}, Texts: []string{"Sigvatr Þórðarson, Austrfararvísur"},
		Origin: "Early eleventh-century skaldic verse", HistoricalNote: "The poem is unusually close to the described event, but it gives no full ritual manual and should not be supplemented with later folklore as fact.",
		CalendarCorpus: "Old Norse textual calendar", NativeDateLabel: "Autumn journey in Austrfararvísur; no numbered native day", AttestationLayer: "Near-contemporary eleventh-century skaldic witness", Era: "Early eleventh century CE", Site: "Sweden in the poem's travel setting",
		ProjectionKind: "season-only-attestation", ProjectionStatus: "Retained in its autumn narrative setting; no annual Gregorian night is fabricated.", DateConfidence: "High rite name and episode; low exact-date confidence", AnchorLocation: "Sweden; exact households unresolved", DayBoundary: "Night/household feast", SourceName: "Skaldic Poetry of the Scandinavian Middle Ages · Austrfararvísur", SourceURL: alfablotSource,
	},
	{
		ID: "norse-disablot", Name: "Dísablót and Dísaþing · Late-Winter Tradition", Communities: []string{"Historical and medieval Swedish communities"}, Category: "Old Norse/Swedish cult and assembly tradition",
		Summary: "Later Norse sources associate sacrifice to the dísir with a Swedish assembly and market at Uppsala in late winter.", Meaning: "The tradition connects cult for female powers or ancestors with regional gathering and political assembly.",
		AttestedElements: []string{"Sacrifice to the dísir", "Assembly and market in later accounts"}, Texts: []string{"Ynglinga saga and later Swedish/Icelandic notices"},
		Origin: "Medieval saga and antiquarian tradition", HistoricalNote: "Sources are late and combine cult, assembly, and market evidence; modern fixed-date Dísablót reconstructions should remain separate.",
		CalendarCorpus: "Old Norse textual calendar", NativeDateLabel: "Late winter; source traditions do not yield one secure universal day", AttestationLayer: "Medieval literary witnesses", Era: "Medieval witnesses to earlier Swedish tradition", Site: "Uppsala, Sweden",
		ProjectionKind: "season-only-late-attestation", ProjectionStatus: "No Gregorian projection because the date and relationship among Dísablót, Dísaþing, and regional calendars are disputed.", DateConfidence: "Moderate festival complex; low exact-date confidence", AnchorLocation: "Uppsala, Sweden", DayBoundary: "Not securely specified", SourceName: "Viking Society · Heimskringla I", SourceURL: heimskringlaSource,
	},
	{
		ID: "norse-thorrablot-medieval", Name: "Þorrablót · Medieval Etiological Tradition", Communities: []string{"Medieval Icelandic and Norwegian textual communities"}, Category: "Disputed Old Norse midwinter feast tradition",
		Summary: "Medieval tradition explains the month Þorri through an eponymous midwinter sacrificial feast; the modern Icelandic Þorrablót revival is a separate phenomenon.", Meaning: "The record belongs in the source history of the calendar, with its evidentiary limits visible.",
		AttestedElements: []string{"Midwinter feast in an etiological narrative"}, Texts: []string{"Orkneyinga saga prologue / Hversu Noregr byggðist tradition"},
		Origin: "Medieval Icelandic etiological prose", Status: "Late historical attestation; antiquity and continuity uncertain", HistoricalNote: "A medieval origin story is not proof of an unchanged Viking Age holiday, and the well-known modern Þorrablót dates from a later revival.",
		CalendarCorpus: "Old Norse textual calendar", NativeDateLabel: "Month Þorri / midwinter in medieval calendar tradition", AttestationLayer: "Medieval etiological narrative", Era: "High Medieval textual witness; claimed older setting", Site: "Icelandic/Norwegian literary tradition",
		ProjectionKind: "late-literary-month-only", ProjectionStatus: "Not placed on a modern Þorri date because that would merge medieval evidence with a distinct revival calendar.", DateConfidence: "Moderate medieval month tradition; low ancient-festival confidence", AnchorLocation: "Old Icelandic calendar", DayBoundary: "Not securely specified", SourceName: "Orkneyinga Saga · Hjaltalín/Goudie translation", SourceURL: orkneyingaSource,
	},

	// Bede is the sole detailed witness for several Old English month names and
	// explicitly marks some explanations as his own inference.
	{
		ID: "anglosaxon-modraniht", Name: "Mōdraniht · Mothers' Night", Communities: []string{"Pre-Christian English communities as reported by Bede"}, Category: "Old English year-opening night",
		Summary: "Bede says the English called the night opening their year Mothers' Night and kept ceremonies through the night.", Meaning: "The notice preserves a rare name and nocturnal year-boundary observance, while Bede himself says its rationale is his supposition.",
		AttestedElements: []string{"All-night ceremonies according to Bede"}, Texts: []string{"Bede, De temporum ratione 15"},
		Origin: "Bede's eighth-century computistical account", HistoricalNote: "Bede is a Christian author and effectively the sole source; his explanation 'because, we suspect' must remain marked as inference.",
		CalendarCorpus: "Bede's Old English month record", NativeDateLabel: "Night at the old year opening, aligned by Bede with viii Kalends January in the Julian calendar", AttestationLayer: "Bede, writing in 725 CE about earlier English custom", Era: "Early medieval witness to pre-Christian English practice", Site: "Anglo-Saxon England",
		ProjectionKind: "source-specific-julian-and-lunar", ProjectionStatus: "No permanent Gregorian date: Bede's Julian alignment, lunar months, and historical year must be specified first.", DateConfidence: "High wording in Bede; independent corroboration absent", AnchorLocation: "Anglo-Saxon England", DayBoundary: "Night", SourceName: "Bede · De ratione temporum, chapter 15 (Latin)", SourceURL: bedeMonthsSource,
	},
	{
		ID: "anglosaxon-solmonath-cakes", Name: "Solmōnaþ · Month of Cakes Offered to the Gods", Communities: []string{"Pre-Christian English communities as reported by Bede"}, Category: "Old English ritual month",
		Summary: "Bede explains Solmōnaþ as the month of cakes that the English offered to their gods.", Meaning: "The month notice preserves a seasonal offering practice but not one named universal feast day.",
		AttestedElements: []string{"Cakes offered to deities"}, Texts: []string{"Bede, De temporum ratione 15"},
		Origin: "Bede's eighth-century computistical account", Status: "Historical ritual-month record; not asserted as one feast day", HistoricalNote: "Bede supplies no recipe, deity list, or exact lunar day; later reconstructions must not fill those gaps silently.",
		CalendarCorpus: "Bede's Old English month record", NativeDateLabel: "Solmōnaþ, the lunar month Bede aligns approximately with February", AttestationLayer: "Bede, 725 CE", Era: "Early medieval witness to pre-Christian English practice", Site: "Anglo-Saxon England",
		ProjectionKind: "lunar-month-only", ProjectionStatus: "Cataloged as a ritual month, not placed on February 1 or another invented day.", DateConfidence: "High month notice in Bede; no exact-day evidence", AnchorLocation: "Anglo-Saxon England", DayBoundary: "Lunar month; phase boundary unspecified by Bede", SourceName: "Bede · De ratione temporum, chapter 15 (Latin)", SourceURL: bedeMonthsSource,
	},
	{
		ID: "anglosaxon-hrethmonath", Name: "Hrēþmōnaþ · Sacrifices for Hrēða", Communities: []string{"Pre-Christian English communities as reported by Bede"}, Category: "Old English deity festival month",
		Summary: "Bede says the month Hrēþmōnaþ took its name from the goddess Hrēða, to whom sacrifices were made then.", Meaning: "The short notice is the primary evidence for both the goddess and the month's rites.",
		AttestedElements: []string{"Sacrifices to Hrēða"}, Texts: []string{"Bede, De temporum ratione 15"},
		Origin: "Bede's eighth-century computistical account", HistoricalNote: "Bede is the only substantive witness; attributes, myths, and exact rites assigned to Hrēða by modern writers are not in his text.",
		CalendarCorpus: "Bede's Old English month record", NativeDateLabel: "Hrēþmōnaþ, the lunar month Bede aligns approximately with March", AttestationLayer: "Bede, 725 CE", Era: "Early medieval witness to pre-Christian English practice", Site: "Anglo-Saxon England",
		ProjectionKind: "lunar-month-only", ProjectionStatus: "Not reduced to a modern March holiday because Bede preserves a lunar month, not a numbered feast day.", DateConfidence: "High report in Bede; no independent corroboration", AnchorLocation: "Anglo-Saxon England", DayBoundary: "Lunar month; phase boundary unspecified by Bede", SourceName: "Bede · De ratione temporum, chapter 15 (Latin)", SourceURL: bedeMonthsSource,
	},
	{
		ID: "anglosaxon-eosturmonath", Name: "Ēosturmōnaþ · Feasts for Ēostre", Communities: []string{"Pre-Christian English communities as reported by Bede"}, Category: "Old English deity festival month",
		Summary: "Bede says Ēosturmōnaþ was named for a goddess Ēostre and that feasts were celebrated for her in that month before the name passed to the Paschal season.", Meaning: "The notice documents a remembered English month-name transition, not a detailed universal spring ritual.",
		AttestedElements: []string{"Feasts for Ēostre; details not supplied"}, Texts: []string{"Bede, De temporum ratione 15"},
		Origin: "Bede's eighth-century computistical account", HistoricalNote: "Bede is the sole early source. This entry must not be conflated with the modern Wiccan name Ostara or used to claim that Christian Easter derives wholesale from one pagan feast.",
		CalendarCorpus: "Bede's Old English month record", NativeDateLabel: "Ēosturmōnaþ, the lunar month Bede aligns approximately with April", AttestationLayer: "Bede, 725 CE", Era: "Early medieval witness to pre-Christian English practice", Site: "Anglo-Saxon England",
		ProjectionKind: "lunar-month-only", ProjectionStatus: "Retained at month resolution; no equinox or Gregorian date is inferred.", DateConfidence: "High wording in Bede; independent corroboration absent", AnchorLocation: "Anglo-Saxon England", DayBoundary: "Lunar month; phase boundary unspecified by Bede", SourceName: "Bede · De ratione temporum, chapter 15 (Latin)", SourceURL: bedeMonthsSource,
	},
	{
		ID: "anglosaxon-blotmonath", Name: "Blōtmōnaþ · Month of Immolations", Communities: []string{"Pre-Christian English communities as reported by Bede"}, Category: "Old English sacrifice month",
		Summary: "Bede calls Blōtmōnaþ the month of immolations because animals to be slaughtered were dedicated to the gods then.", Meaning: "The record preserves a seasonal sacrificial context rather than a single named date.",
		AttestedElements: []string{"Dedication of livestock to deities", "Seasonal slaughter context"}, Texts: []string{"Bede, De temporum ratione 15"},
		Origin: "Bede's eighth-century computistical account", HistoricalNote: "The source does not identify one feast day, participating deities, or a uniform rite across all English kingdoms.",
		CalendarCorpus: "Bede's Old English month record", NativeDateLabel: "Blōtmōnaþ, the lunar month Bede aligns approximately with November", AttestationLayer: "Bede, 725 CE", Era: "Early medieval witness to pre-Christian English practice", Site: "Anglo-Saxon England",
		ProjectionKind: "lunar-month-only", ProjectionStatus: "Cataloged as a ritual season; no November day is fabricated.", DateConfidence: "High month notice in Bede; no exact-day evidence", AnchorLocation: "Anglo-Saxon England", DayBoundary: "Lunar month; phase boundary unspecified by Bede", SourceName: "Bede · De ratione temporum, chapter 15 (Latin)", SourceURL: bedeMonthsSource,
	},
	{
		ID: "anglosaxon-giuli-double-month", Name: "Giuli · Double Yule Month", Communities: []string{"Early English communities as reported by Bede"}, Category: "Old English midwinter calendar season",
		Summary: "Bede gives both the month aligned with December and the following month aligned with January the name Giuli, with an intercalary third Litha in leap years elsewhere in the cycle.", Meaning: "The double month is valuable evidence for a lunisolar English year even though Bede does not describe a separate Giuli ritual program here.",
		AttestedElements: []string{"Double month-name around the year boundary"}, Texts: []string{"Bede, De temporum ratione 15"},
		Origin: "Bede's eighth-century computistical account", Status: "Historical calendar season; no independent feast program asserted", HistoricalNote: "The shared Germanic word family invites comparison with Old Norse Jól, but the corpora and their dating rules remain separate in this catalog.",
		CalendarCorpus: "Bede's Old English month record", NativeDateLabel: "Two successive lunar months called Giuli around the English year boundary", AttestationLayer: "Bede, 725 CE", Era: "Early medieval witness to earlier English calendar practice", Site: "Anglo-Saxon England",
		ProjectionKind: "lunisolar-season-only", ProjectionStatus: "No fixed December–January span is generated; Bede's lunar months and Julian frame require a specified year.", DateConfidence: "High month-name report; ritual content not supplied", AnchorLocation: "Anglo-Saxon England", DayBoundary: "Lunar month; phase boundary unspecified by Bede", SourceName: "Bede · De ratione temporum, chapter 15 (Latin)", SourceURL: bedeMonthsSource,
	},
}

type modernPolytheistRule struct {
	Month       time.Month
	Day         int
	ID          string
	Name        string
	Communities []string
	Summary     string
	Meaning     string
	Practices   []string
	DateNote    string
}

// modernPolytheistObservances is intentionally a living-practice layer, not an
// assertion that one eight-feast wheel existed unchanged in antiquity. The
// fixed dates follow the documented legal calendar of ADF; communities may use
// astronomical instants, other civil dates, local ecology, or nearby weekends.
func modernPolytheistObservances(date time.Time) []Observance {
	events := make([]Observance, 0, 1)
	for _, rule := range modernPolytheistRules {
		if date.Month() != rule.Month || date.Day() != rule.Day {
			continue
		}
		event := baseObservance(rule.Name, Polytheist, rule.Communities, "Living Neo-Pagan high day", rule.Summary, rule.Meaning, rule.Practices, nil, "ADF Constitution · Article 4 calendar", adfCalendarSource)
		event.CatalogID = rule.ID
		event.Origin = "Modern Neo-Pagan eightfold calendar"
		event.ObservanceStatus = "Living observance; names, dates, and practices vary by tradition and hemisphere"
		event.DateNote = rule.DateNote
		event.DateCertainty = "High for the cited fixed-date convention; not a universal Pagan date"
		event.CalendarCorpus = "Modern Neo-Pagan Wheel of the Year"
		event.NativeDateLabel = date.Format("January 2 fixed civil convention")
		event.AttestationLayer = "ADF Constitution Article 4 and contemporary community practice"
		event.Era = "Modern religious movement"
		event.Site = "Living communities; northern-hemisphere convention shown"
		event.ProjectionKind = "documented-modern-fixed-date"
		event.ProjectionStatus = "Direct placement under ADF's legal calendar; exact astronomical and local ecological observances are separate choices"
		event.DateConfidence = "High for this named modern convention"
		event.AnchorLocation = "Northern Hemisphere convention; seasons reverse in the Southern Hemisphere"
		event.DayBoundary = "Local civil day; some communities begin on the preceding evening"
		event = singleOccurrence(event, date)
		event.ID = rule.ID + "-" + date.Format("2006-01-02")
		events = append(events, event)
	}
	return events
}

var modernPolytheistRules = []modernPolytheistRule{
	{time.February, 1, "modern-pagan-winter-cross-quarter", "Winter Cross-Quarter · Imbolc in Many Traditions", []string{"ADF Druids", "Many Wiccans", "Other Neo-Pagan communities"}, "A living high day at the seasonal turn toward spring, often called Imbolc in Celtic-oriented and Wiccan practice.", "Communities emphasize returning light, hearth, inspiration, purification, and the first local signs of spring.", []string{"Community or household ritual", "Lighting candles or a hearth flame", "Seasonal reflection and offerings according to one's tradition"}, "ADF fixes this cross-quarter high day on February 1 for legal purposes. Other communities use February 2, an astronomical midpoint, local seasonal signs, or another nearby date; ancient Irish evidence is not being projected here."},
	{time.March, 21, "modern-pagan-spring-equinox", "Spring Equinox · Modern Pagan High Day", []string{"ADF Druids", "Many Wiccans", "Other Neo-Pagan communities"}, "A modern solar high day for the spring equinox, called by several names across living traditions.", "The observance attends to balance, emergence, and renewed growth.", []string{"Seasonal ritual", "Attention to dawn, plants, or local ecology", "Community gathering"}, "ADF's legal calendar uses March 21. The astronomical equinox can fall on a different civil date and is calculated separately; 'Ostara' and 'Alban Eilir' are tradition-specific modern names, not universal ancient titles."},
	{time.May, 1, "modern-pagan-spring-cross-quarter", "Spring Cross-Quarter · Beltane in Many Traditions", []string{"ADF Druids", "Many Wiccans", "Other Neo-Pagan communities"}, "A living spring cross-quarter high day, often called Beltane, centered on summer's approach and flourishing life.", "Communities variously emphasize vitality, relationship, protection, fire, and the greening land.", []string{"Community celebration", "Fire or flower symbolism where safe and appropriate", "Seasonal offerings according to one's tradition"}, "ADF fixes the cross-quarter on May 1. Some groups begin on April 30, use the astronomical midpoint, or follow local ecology; this modern high-day entry is separate from reconstructing medieval Irish Beltaine."},
	{time.June, 21, "modern-pagan-summer-solstice", "Summer Solstice · Modern Pagan High Day", []string{"ADF Druids", "Many Wiccans", "Other Neo-Pagan communities"}, "A modern solar high day near the northern-hemisphere year's longest daylight.", "The day invites celebration of fullness, light, growth, and the turning toward shorter days.", []string{"Outdoor or community ritual", "Solar and fire symbolism", "Gratitude for seasonal abundance"}, "ADF's legal calendar uses June 21. The actual solstice instant is calculated separately and may fall on June 20 or another local civil date; names such as Midsummer, Litha, and Alban Hefin belong to particular communities."},
	{time.August, 1, "modern-pagan-summer-cross-quarter", "Summer Cross-Quarter · Lughnasadh in Many Traditions", []string{"ADF Druids", "Many Wiccans", "Other Neo-Pagan communities"}, "A living harvest-season high day, often called Lughnasadh or Lammas in modern Pagan practice.", "Communities give thanks for first fruits, skilled work, reciprocity, and the costs and gifts of harvest.", []string{"Sharing seasonal food", "Offerings of first fruits or crafted work", "Community games, ritual, or reflection"}, "ADF fixes the cross-quarter on August 1. Historical Irish Lugnasad, Christian Lammas, and modern Neo-Pagan observances have overlapping but nonidentical histories and are not collapsed into one ancient rite."},
	{time.September, 21, "modern-pagan-autumn-equinox", "Autumn Equinox · Modern Pagan High Day", []string{"ADF Druids", "Many Wiccans", "Other Neo-Pagan communities"}, "A living solar high day around the autumn equinox and harvest season.", "Communities reflect on balance, gratitude, ripening, release, and preparation for winter.", []string{"Harvest thanksgiving", "Community or household ritual", "Seasonal reflection and food sharing"}, "ADF's legal calendar uses September 21. The astronomical equinox can occur on a different civil day and is calculated separately; 'Mabon' and 'Alban Elfed' are modern, tradition-specific names rather than securely attested ancient universal titles."},
	{time.November, 1, "modern-pagan-autumn-cross-quarter", "Autumn Cross-Quarter · Samhain in Many Traditions", []string{"ADF Druids", "Many Wiccans", "Other Neo-Pagan communities"}, "A living autumn cross-quarter high day, often called Samhain, with strong themes of ancestors, endings, and winter's approach.", "The day creates space for remembrance, hospitality, grief, continuity, and reflection at a seasonal threshold.", []string{"Ancestor remembrance", "Community or household ritual", "Seasonal offerings according to one's tradition"}, "ADF's legal calendar fixes the cross-quarter on November 1. Many groups begin on October 31, calculate a midpoint, or follow local tradition; this living entry does not claim one unchanged pan-Celtic ancient New Year."},
	{time.December, 21, "modern-pagan-winter-solstice", "Winter Solstice · Modern Pagan High Day", []string{"ADF Druids", "Many Wiccans", "Other Neo-Pagan communities"}, "A living solar high day near the northern-hemisphere year's longest night, called Yule or other names in many communities.", "The observance emphasizes endurance, kinship, darkness, and the returning light.", []string{"Lighting candles or a hearth fire", "Community feast or vigil", "Seasonal gifts, offerings, or reflection"}, "ADF's legal calendar uses December 21. The actual solstice is calculated separately; modern Yule is not treated here as identical to every Old Norse Jól date or practice."},
}
