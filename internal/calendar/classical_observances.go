package calendar

import "time"

const (
	romanFastiSource = "https://www.gutenberg.org/ebooks/59007"
	holyLanceSource  = "https://doi.org/10.1515/zkg-2023-2002"
)

// romanFixedFestival is one named festival row in the reconstructed
// fasti antiquissimi. Repeated ritual days (for example the three Lemuria)
// remain separate rows because that is how the ancient calendar presents them.
type romanFixedFestival struct {
	CatalogID string
	Month     time.Month
	Day       int
	Name      string
	Summary   string
	Meaning   string
	Practices []string
}

// romanFixedCanon contains all 45 named fixed-date rows in W. Warde Fowler's
// reconstruction of the oldest Republican fasti. LIBERALIA AGONIA is one
// calendar row; later temple anniversaries, games, and movable feriae are not
// silently promoted into this epigraphic core.
var romanFixedCanon = []romanFixedFestival{
	{"roman-agonalia-jan", time.January, 9, "Agonalia of Janus", "The rex sacrorum offered a ram in the Regia on the first of four days marked Agonia.", "The rite joined royal-priestly office to the opening of the civic year; even Roman antiquarians were unsure of the name's origin and divine recipient.", []string{"Ram sacrifice by the rex sacrorum historically", "Public suspension appropriate to a ferial day"}},
	{"roman-carmentalia-first", time.January, 11, "Carmentalia · First Day", "The first of two fixed feasts honored Carmentis, a goddess associated with prophecy, birth, and Rome's remembered past.", "The paired dates preserve an archaic women's and civic cult whose original relationship was already obscure in antiquity.", []string{"Offerings at the Porta Carmentalis historically", "Matronal observance"}},
	{"roman-carmentalia-second", time.January, 15, "Carmentalia · Second Day", "The second Carmentalia returned to the cult of Carmentis four days after its opening feast.", "Its separation from the first day warns against forcing a later continuous festival onto two distinct fasti entries.", []string{"Offerings to Carmentis historically", "Matronal observance"}},

	{"roman-lupercalia", time.February, 15, "Lupercalia", "The Luperci sacrificed and ran a ritual circuit at the foot of the Palatine in one of Rome's best-attested archaic rites.", "Ancient and modern explanations variously stress purification, fertility, protection, and civic boundary; no single explanation exhausts the evidence.", []string{"Sacrifice at the Lupercal historically", "Ritual run by the Luperci", "Contact with ritual thongs in later descriptions"}},
	{"roman-quirinalia", time.February, 17, "Quirinalia", "A public feast honored Quirinus; the same date served as the collective close of the movable Fornacalia.", "The day linked an archaic state god with the curial grain-roasting cycle, while the fasti name only the Quirinalia.", []string{"Sacrifice to Quirinus historically", "Collective Fornacalia observance for those who missed a curial date"}},
	{"roman-feralia", time.February, 21, "Feralia", "The public Feralia closed the principal rites of the Parentalia for the ancestral dead.", "It made care for the dead a civic as well as household obligation without turning later folklore into a complete ritual script.", []string{"Offerings at tombs historically", "Closure of temples and suspension of marriages during the ancestral season"}},
	{"roman-terminalia", time.February, 23, "Terminalia", "Households and the state honored Terminus at boundary stones and at Rome's territorial boundary.", "Shared offerings at a fixed marker treated negotiated limits as sacred protections of neighborly order.", []string{"Garlanding boundary stones historically", "Paired offerings by neighboring landholders", "State rite at the sixth milestone"}},
	{"roman-regifugium", time.February, 24, "Regifugium", "The rex sacrorum performed a rite in the Comitium and then departed in ritual haste.", "The ceremony's meaning is uncertain; the familiar story of the kings' expulsion is an ancient interpretation, not a secure origin.", []string{"Rite of the rex sacrorum historically", "Ritual departure from the Comitium"}},
	{"roman-equirria-feb", time.February, 27, "Equirria · First Race", "Horse races dedicated to Mars formed the first of two Equirria dates around the old year's turning.", "The races prepared the civic and military community for the season associated with Mars.", []string{"Horse racing in the Campus Martius or alternate course historically", "Rites for Mars"}},

	{"roman-equirria-mar", time.March, 14, "Equirria · Second Race", "The second Equirria repeated Mars's horse races shortly before the March Ides.", "Together the two separated dates frame the old Roman new-year transition rather than a modern continuous festival.", []string{"Horse racing historically", "Rites for Mars"}},
	{"roman-liberalia-agonia", time.March, 17, "Liberalia and Agonalia", "The fasti combine LIBERALIA and AGONIA on one day: Liber's public feast coincided with another archaic Agonalia.", "Liberalia became associated with young men's assumption of the toga virilis, while the Agonalia layer remains obscure.", []string{"Offerings to Liber and Libera historically", "Honey cakes associated with priestesses", "Assumption of the toga virilis in later practice"}},
	{"roman-quinquatrus", time.March, 19, "Quinquatrus", "An archaic one-day observance associated with the purification of arms later opened five days of Minervan celebration.", "Its layered history joins the March military cycle with artisans', teachers', and Minerva's later festival culture.", []string{"Purification of arms historically", "Offerings to Minerva in later practice", "Games during the later expansion"}},
	{"roman-tubilustrium-mar", time.March, 23, "Tubilustrium · March", "Sacred war trumpets were purified at the first of two annual Tubilustria.", "The rite prepared ritual and military instruments within the festival sequence of Mars's month.", []string{"Purification of the tubae historically", "Sacrifice and Salian activity"}},

	{"roman-fordicidia", time.April, 15, "Fordicidia", "Pregnant cattle were offered to Tellus in a state and curial rite for the fertility of the earth.", "The historical rite expressed the costly reciprocity imagined between animal, field, harvest, and civic survival; it is described, not prescribed.", []string{"Sacrifice of pregnant cattle historically", "Vestal preservation of material used at the Parilia"}},
	{"roman-cerealia", time.April, 19, "Cerealia", "The fixed feast of Ceres stood within a longer period of games honoring Ceres, Liber, and Libera.", "Grain, plebeian civic identity, and the vulnerability of food supply met in the cult of Ceres.", []string{"Public sacrifice historically", "Circus games in the later Ludi Cereales", "White clothing in some literary accounts"}},
	{"roman-parilia", time.April, 21, "Parilia", "A pastoral purification feast for Pales later shared its date with Rome's traditional birthday.", "The rite renewed flocks, shepherds, and eventually the civic story of Rome without making the 753 BCE foundation date historical certainty.", []string{"Purification of sheepfolds historically", "Bonfires and ritual smoke", "Offerings to Pales", "Celebration of Rome's birthday in later practice"}},
	{"roman-vinalia-priora", time.April, 23, "Vinalia Priora", "The first Vinalia opened the previous vintage and concerned vows and wine offered to Jupiter, with Venus prominent in later calendars.", "The feast distinguished sacred first offering from ordinary consumption and sale.", []string{"Libation of the prior year's wine historically", "Wine tasting and market opening", "Rites for Jupiter and later Venus"}},
	{"roman-robigalia", time.April, 25, "Robigalia", "A procession beyond the city sought protection of grain from destructive rust disease.", "The rite confronted agricultural risk through propitiation of Robigo or Robigus; its animal sacrifice is historical description.", []string{"Procession along the Via Claudia historically", "Prayer for healthy grain", "Dog and sheep sacrifice historically"}},

	{"roman-lemuria-first", time.May, 9, "Lemuria · First Night", "Household rites addressed restless or harmful dead on the first of three alternate Lemuria dates.", "The separated nights acknowledge obligations to unsettled dead distinct from February's ancestral Parentalia.", []string{"Midnight household rite historically", "Casting black beans", "Bronze sounding and dismissal formula"}},
	{"roman-lemuria-second", time.May, 11, "Lemuria · Second Night", "The household Lemuria resumed after an intervening day.", "Repeating the rite on alternate nights formed a ritual cadence rather than a continuous five-day holiday.", []string{"Midnight household rite historically", "Casting black beans", "Bronze sounding and dismissal formula"}},
	{"roman-lemuria-third", time.May, 13, "Lemuria · Third Night", "The third Lemuria completed the household sequence for restless dead.", "Its close restored ordinary relations between household, ancestors, and potentially dangerous spirits.", []string{"Final household rite historically", "Casting black beans", "Ritual dismissal of the lemures"}},
	{"roman-agonalia-may", time.May, 21, "Agonalia of May", "A third fixed Agonalia was associated in some fasti with Vediovis.", "The divine attribution is better preserved than the rite's purpose, so the entry retains ancient uncertainty.", []string{"Public sacrifice historically", "Possible rite for Vediovis"}},
	{"roman-tubilustrium-may", time.May, 23, "Tubilustrium · May", "The second annual Tubilustrium purified sacred trumpets and was linked in a later fasti note with Vulcan.", "It closed another cycle of preparing ritual instruments while preserving uncertainty about its precise divine focus.", []string{"Purification of trumpets historically", "Sacrifice historically"}},

	{"roman-vestalia", time.June, 9, "Vestalia", "The fasti's fixed Vestalia honored Vesta at the center of a longer June sequence involving her sanctuary and priesthood.", "The hearth's guarded fire represented continuity of household and commonwealth; later reconstructions should distinguish the one named day from the broader period.", []string{"Offerings at Vesta's sanctuary historically", "Ritual access for matrons", "Milling and bakery observances in the broader season"}},
	{"roman-matralia", time.June, 11, "Matralia", "Married women honored Mater Matuta, while Fortuna received rites at the neighboring Forum Boarium temples.", "The ritual placed maternal kinship and the care of sisters' children within civic worship, though surviving explanations are layered.", []string{"Matronal offerings historically", "Rites for Mater Matuta and Fortuna"}},

	{"roman-poplifugia", time.July, 5, "Poplifugia", "The 'Flight of the People' was a fixed archaic festival whose historical explanation was already contested.", "Its terse calendar name should be preserved without treating later stories of Romulus or military panic as certain origin.", []string{"Public rite historically; details are poorly preserved"}},
	{"roman-lucaria-first", time.July, 19, "Lucaria · First Day", "The first Lucaria honored or ritually engaged a sacred grove; ancient explanations linked the dates with refuge after the Gallic invasion.", "The name securely signals a grove, while the familiar historical narrative remains an interpretation.", []string{"Rites in a grove historically", "Public ferial observance"}},
	{"roman-lucaria-second", time.July, 21, "Lucaria · Second Day", "The second Lucaria returned to the grove after an intervening day.", "Its paired structure is epigraphically clearer than the lost details of what worshippers did.", []string{"Rites in a grove historically", "Public ferial observance"}},
	{"roman-neptunalia", time.July, 23, "Neptunalia", "Neptune's midsummer feast was celebrated during the dry season, with temporary leafy shelters attested in later sources.", "The observance sought an ordered relation with water and drought rather than simply mirroring the Greek cult of Poseidon.", []string{"Rites for Neptune historically", "Leaf-covered shelters in literary accounts", "Communal outdoor gathering"}},
	{"roman-furrinalia", time.July, 25, "Furrinalia", "A public feast honored the archaic goddess Furrina under the care of her flamen.", "The cult's fixed place in the oldest calendar outlived Roman knowledge of the goddess's original character.", []string{"Sacrifice by the Flamen Furrinalis historically", "Observance at Furrina's grove"}},

	{"roman-portunalia", time.August, 17, "Portunalia", "The feast of Portunus concerned gates, harbors, and passages; keys were treated ritually in a later antiquarian notice.", "The day marked control of thresholds and safe passage in the season's cluster of harvest rites.", []string{"Sacrifice to Portunus historically", "Ritual treatment of keys in later testimony"}},
	{"roman-vinalia-rustica", time.August, 19, "Vinalia Rustica", "The second Vinalia preceded the grape harvest and concerned Jupiter's protection of the vintage, with Venus prominent in later cult.", "It oriented an uncertain harvest toward divine permission and protection before human use.", []string{"First-fruit or vintage prayers historically", "Rites for Jupiter and Venus", "Inspection of grapes before harvest"}},
	{"roman-consualia-aug", time.August, 21, "Consualia · Summer", "An underground altar of Consus was uncovered for sacrifice and games at the summer Consualia.", "The festival joined stored produce, counsel, horses, and communal games in ways no single modern gloss fully captures.", []string{"Uncovering Consus's altar historically", "Horse and mule races", "Rest and garlanding for working animals"}},
	{"roman-vulcanalia", time.August, 23, "Vulcanalia", "Vulcan's feast addressed the danger and utility of fire at the height of summer heat.", "Offerings sought to turn destructive fire away from grain, buildings, and community.", []string{"Offerings in fire historically", "Rites to Vulcan and associated deities", "Beginning work by lamplight in one literary account"}},
	{"roman-opiconsivia", time.August, 25, "Opiconsivia", "Ops Consiva received a restricted rite in the Regia attended by the Vestals and chief priest.", "The hidden ceremonial emphasized gathered abundance and the protected store of the community.", []string{"Restricted Regia rite historically", "Offerings to Ops Consiva"}},
	{"roman-volturnalia", time.August, 27, "Volturnalia", "The Flamen Volturnalis sacrificed to Volturnus on the last named festival of the August sequence.", "The rite is securely calendrical but its deity and purpose are too poorly preserved for confident elaboration.", []string{"Sacrifice by the Flamen Volturnalis historically"}},

	{"roman-meditrinalia", time.October, 11, "Meditrinalia", "New wine was ritually tasted and mixed with old in a feast whose divine recipient was already uncertain.", "The rite framed transformation and first tasting with a protective formula for bodily health.", []string{"Mixing new and old wine historically", "First tasting with a traditional formula"}},
	{"roman-fontinalia", time.October, 13, "Fontinalia", "Springs and wells were garlanded for Fons or Fontus.", "The feast treated reliable fresh water as a sacred public good worthy of gratitude and care.", []string{"Garlanding wells and springs historically", "Offerings to Fons"}},
	{"roman-armilustrium", time.October, 19, "Armilustrium", "Weapons were purified on the Aventine as the campaigning season closed.", "The rite balanced March's preparation for war with ritual closure and cleansing in autumn.", []string{"Purification of arms historically", "Salian music and procession", "Rites for Mars"}},

	{"roman-agonalia-dec", time.December, 11, "Agonalia of December", "The final Agonalia, associated by later sources with Sol Indiges, shared its date with the locally organized Septimontium.", "The fasti preserve the Agonalia securely while the deity and relationship to the hill festival remain debated.", []string{"Public sacrifice historically", "Septimontium observances by hill communities in later testimony"}},
	{"roman-consualia-dec", time.December, 15, "Consualia · Winter", "The second Consualia again uncovered Consus's altar during the midwinter festival sequence.", "Stored abundance and working animals were honored after the agricultural year rather than before harvest.", []string{"Uncovering Consus's altar historically", "Horse and mule games", "Rest and garlanding for working animals"}},
	{"roman-saturnalia", time.December, 17, "Saturnalia · Public Rite", "The oldest fasti mark Saturnalia on one fixed day, when Saturn's public sacrifice and banquet opened a much-loved season.", "Later gift-giving, gaming, relaxed dress, and role reversals expanded around a one-day Republican civic core.", []string{"Sacrifice at Saturn's temple historically", "Public banquet", "Cry of Io Saturnalia", "Gift-giving and licensed play in later practice"}},
	{"roman-opalia", time.December, 19, "Opalia", "Ops received her own fixed feast two days after Saturnalia.", "The paired but distinct dates joined stored plenty with Saturn's remembered age without collapsing Ops into an appendix to Saturnalia.", []string{"Offerings to Ops historically", "Public ferial observance"}},
	{"roman-divalia", time.December, 21, "Divalia · Angeronalia", "A secretive rite honored Diva Angerona, whose image was described with a bound or sealed mouth.", "Ancient explanations range from sacred silence to relief from anguish; the calendar attestation is firmer than any interpretation.", []string{"Sacrifice to Angerona historically", "Priestly rite in the sanctuary of Volupia"}},
	{"roman-larentalia", time.December, 23, "Larentalia", "The flamen Quirinalis performed rites at the tomb or shrine of Acca Larentia.", "Competing stories made Larentia nurse, benefactor, courtesan, or ancestral figure; the rite's civic memory should not be reduced to one tale.", []string{"Funerary offering historically", "Rite by the Flamen Quirinalis"}},
}

type romanLaterFixedFestival struct {
	CatalogID      string
	Month          time.Month
	Day            int
	Name           string
	Corpus         string
	Era            string
	Summary        string
	Meaning        string
	Practices      []string
	SourceName     string
	SourceURL      string
	Attestation    string
	DateConfidence string
}

// romanLaterFixedFestivals is intentionally not folded into romanFixedCanon.
// These records come from literary testimony, temple calendars, or later
// Republican/Imperial fasti rather than the 45 named rows of the oldest layer.
var romanLaterFixedFestivals = []romanLaterFixedFestival{
	{"roman-kalends-january", time.January, 1, "Kalends of January · Roman New Year", "Later Republican and Imperial civic calendar", "From 153 BCE through the Empire", "Consuls entered office, public vows were renewed, auspices were taken, and good-omen gifts marked the January civil new year.", "The date's public importance grew after the consular year moved to January; it should not be back-projected as Rome's original religious new year.", []string{"Public vows and auspices historically", "Consular inauguration", "Exchange of strenae or good-omen gifts"}, "W. Warde Fowler · Roman Festivals", romanFastiSource, "Literary testimony and later civic calendar", "high for January 1 after 153 BCE"},
	{"roman-caristia", time.February, 22, "Caristia · Cara Cognatio", "Roman household calendar", "Republican and Imperial Rome", "Families shared a reconciliatory meal after the ancestral rites of the Parentalia and Feralia.", "The day turned from obligations to the dead toward repaired relationships among living kin.", []string{"Family meal historically", "Reconciliation among relatives", "Household offerings to the Lares"}, "W. Warde Fowler · Roman Festivals", romanFastiSource, "Literary household-calendar testimony", "high for February 22"},
	{"roman-matronalia", time.March, 1, "Matronalia", "Roman literary and temple calendar", "Republican and Imperial Rome", "Married women honored Juno Lucina on the old new-year Kalends, while households exchanged gifts and hospitality.", "The observance joined childbirth, marriage, household reciprocity, and renewal, although Matronalia was not a named row in the archaic fasti.", []string{"Offerings to Juno Lucina historically", "Gifts between spouses", "Household feast and respite for enslaved workers in later accounts"}, "W. Warde Fowler · Roman Festivals", romanFastiSource, "Literary testimony and Juno Lucina temple anniversary", "high for March 1; details vary by source"},
	{"roman-navigium-isidis", time.March, 5, "Navigium Isidis · Ship of Isis", "Roman Egyptian-cult calendar", "Imperial Rome and Mediterranean Isiac communities", "A spring procession launched a sacred model ship and celebrated Isis as protector of seafarers.", "The festival demonstrates the regional and cosmopolitan religious calendar of the Empire rather than an archaic Roman state feast.", []string{"Costumed procession historically", "Carrying and launching a sacred ship", "Offerings and hymns to Isis"}, "Apuleius · Metamorphoses, Book 11", "https://penelope.uchicago.edu/Thayer/E/Roman/Texts/Apuleius/Metamorphoses/11*.html", "Imperial literary testimony and late Roman calendars", "moderate-to-high for the conventional March 5 date; local practice varied"},
	{"roman-anna-perenna", time.March, 15, "Feast of Anna Perenna", "Roman literary festival calendar", "Republican and Imperial Rome", "Crowds picnicked and celebrated the year goddess Anna Perenna near the March Ides.", "Popular prayers for as many years as cups of wine joined spring, longevity, and the old year's first full moon season.", []string{"Outdoor picnic historically", "Drinking and songs", "Prayers for long life"}, "W. Warde Fowler · Roman Festivals", romanFastiSource, "Literary testimony; later fasti annotation", "high for March 15; ritual interpretation is layered"},
	{"roman-veneralia", time.April, 1, "Veneralia", "Roman literary and temple calendar", "Late Republic and Empire", "Women and worshippers honored Venus Verticordia and Fortuna Virilis on the April Kalends.", "The rites concerned beauty, erotic life, moral reputation, and bodily vulnerability in forms filtered through Augustan literary testimony.", []string{"Washing and adorning cult images historically", "Myrtle and bathing rites in literary testimony", "Prayer to Venus and Fortuna"}, "W. Warde Fowler · Roman Festivals", romanFastiSource, "Ovid and later calendar annotations", "high for April 1; reconstruction of practices is moderate"},
	{"roman-mercuralia", time.May, 15, "Mercuralia", "Roman merchants' and temple calendar", "Republican and Imperial Rome", "Merchants honored Mercury and Maia on the May Ides, with a rite at Mercury's spring near the Porta Capena.", "Commercial exchange was placed under divine protection while acknowledging the moral ambiguity Romans associated with profit and persuasion.", []string{"Drawing and sprinkling sacred water historically", "Offerings by merchants", "Prayers for commercial success and pardon"}, "W. Warde Fowler · Roman Festivals", romanFastiSource, "Temple anniversary and literary testimony", "high for May 15"},
	{"roman-fors-fortuna", time.June, 24, "Fors Fortuna", "Roman popular and temple calendar", "Republican and Imperial Rome", "Crowds traveled by boat or on foot to Fortuna's shrines downstream on the Tiber for a midsummer popular festival.", "The outing joined chance, social reversal, river travel, and hope among ordinary worshippers.", []string{"River outing historically", "Offerings at Fortuna's shrines", "Feasting and public celebration"}, "W. Warde Fowler · Roman Festivals", romanFastiSource, "Temple anniversaries and literary testimony", "high for June 24"},
	{"roman-nonae-caprotinae", time.July, 7, "Nonae Caprotinae · Serving Women's Festival", "Roman women's and civic calendar", "Republican and Imperial Rome", "Women, especially enslaved and serving women in later accounts, celebrated beneath a wild fig on the July Nones.", "The day preserved women's ritual agency alongside competing patriotic legends that should not be mistaken for certain origin.", []string{"Rites beneath the caprificus historically", "Women's communal celebration", "Offerings to Juno Caprotina"}, "W. Warde Fowler · Roman Festivals", romanFastiSource, "Varro, Macrobius, and later fasti", "high for July 7; origin narratives uncertain"},
	{"roman-nemoralia", time.August, 13, "Nemoralia · Festival of Diana", "Roman and Latin regional cult calendar", "Republican and Imperial Rome and Aricia", "Diana's Ides festival centered on her Aventine temple and the ancient sanctuary at Lake Nemi.", "The feast connected enslaved people's respite, women's petitions, hunting, woodland, and a cult shared between Rome and Latin communities.", []string{"Offerings to Diana historically", "Processions or visits to the grove at Nemi", "Day of respite for enslaved worshippers", "Lamps and petitions in later testimony"}, "Roman festival fasti summarized by Scullard", "https://penelope.uchicago.edu/encyclopaedia_romana/calendar/antiates.html", "Temple anniversaries, Latin sanctuary, and literary testimony", "high for the August Ides; exact pre-Julian civil equivalence is not asserted"},
	{"roman-epulum-iovis-sep", time.September, 13, "Epulum Iovis · September", "Roman Republican and Imperial state calendar", "Republican and Imperial Rome", "A public banquet for Jupiter, Juno, and Minerva marked the dedication anniversary of the Capitoline temple during the Great Games.", "The gods were represented as civic banquet participants at the symbolic center of the Roman state.", []string{"Lectisternium and divine banquet historically", "Senatorial and priestly participation", "Rites at the Capitoline temple"}, "W. Warde Fowler · Roman Festivals", romanFastiSource, "Fasti annotations and Capitoline temple calendar", "high for September 13"},
	{"roman-augustalia", time.October, 12, "Augustalia", "Roman Imperial cult calendar", "From 19 BCE; games added after 14 CE", "An altar rite for Fortuna Redux commemorated Augustus's return, later developing into honors and games for the deified emperor.", "The feast shows an imperial political anniversary becoming a recurring cult institution rather than belonging to the archaic canon.", []string{"Sacrifice to Fortuna Redux historically", "Imperial-cult honors", "Games in the later Ludi Augustales"}, "Roman festival fasti summarized by Scullard", "https://penelope.uchicago.edu/encyclopaedia_romana/calendar/antiates.html", "Augustan and Imperial fasti", "high for October 12 after institution"},
	{"roman-october-horse", time.October, 15, "October Horse", "Roman military and agricultural ritual calendar", "Republican Rome", "A winning right-hand chariot horse was sacrificed to Mars after a race in the Campus Martius.", "The violent historical rite joined the close of campaigning, competitive civic space, and fertility symbolism; its remains were contested by neighborhoods.", []string{"Chariot race historically", "Horse sacrifice to Mars historically", "Contest over the head and preservation of the tail's blood"}, "W. Warde Fowler · Roman Festivals", romanFastiSource, "Festus, Polybius, and later literary testimony", "high for October 15; interpretations remain disputed"},
	{"roman-epulum-iovis-nov", time.November, 13, "Epulum Iovis · November", "Roman Republican and Imperial games calendar", "Republican and Imperial Rome", "A second annual banquet for the Capitoline triad fell within the Plebeian Games on the November Ides.", "The repeated divine banquet connected plebeian public games to Jupiter's central state cult.", []string{"Lectisternium and divine banquet historically", "Public games", "Priestly and civic participation"}, "W. Warde Fowler · Roman Festivals", romanFastiSource, "Fasti annotations and Plebeian Games calendar", "high for November 13"},
	{"roman-natalis-invicti", time.December, 25, "Natalis Invicti · Birth of the Unconquered", "Late Roman Imperial calendar", "Attested in the fourth-century Roman calendar", "The Chronography of 354 records games for the birthday of the Unconquered on December 25, commonly associated with Sol Invictus.", "The late imperial observance belongs to a precisely attested calendar layer; claims about a simple one-way derivation of Christmas from it go beyond the evidence.", []string{"Thirty circus races in the fourth-century calendar", "Imperial solar-cult observance"}, "Chronography of 354 · Calendar of Filocalus", "https://www.tertullian.org/fathers/chronography_of_354_06_calendar.htm", "Fourth-century Roman calendar", "high for the December 25 calendar notice; interpretation of 'Invictus' is strong but not wholly uncontested"},
}

type romanSpanFestival struct {
	CatalogID      string
	Name           string
	StartMonth     time.Month
	StartDay       int
	Duration       int
	Corpus         string
	Era            string
	Summary        string
	Meaning        string
	Practices      []string
	SourceName     string
	SourceURL      string
	Attestation    string
	DateConfidence string
}

// These later, longer, or popular seasons are deliberately separate from the
// 45-row archaic core. This prevents an imperial game span from rewriting the
// single fixed ferial day preserved in the oldest calendars.
var romanSpanFestivals = []romanSpanFestival{
	{"roman-compitalia-late-placement", "Compitalia · Late Conventional Placement", time.January, 3, 3, "Roman neighborhood and household calendar", "Republican movable feast; January 3–5 placement in late antique calendars", "Neighborhoods honored the Lares Compitales at crossroads after the agricultural year; the date was traditionally proclaimed rather than fixed.", "The three displayed days preserve a common late placement without erasing Compitalia's original status as a movable feriae conceptivae.", []string{"Crossroads offerings historically", "Neighborhood games and meals", "Garlanding household and crossroads shrines"}, "W. Warde Fowler · Roman Festivals", romanFastiSource, "Literary witnesses; late calendar placement", "moderate: January 3–5 is a late conventional placement, not a universal Republican fixed date"},
	{"roman-parentalia-season", "Parentalia · Ancestral Season", time.February, 13, 9, "Roman Republican household and civic calendar", "Republican and Imperial Rome", "Nine days of household remembrance for family dead culminated in the public Feralia.", "The season made kinship with the dead a recurring household duty while ordinary public and marital rites were curtailed.", []string{"Visits and offerings at tombs historically", "Family remembrance", "Feralia on the final public day"}, "W. Warde Fowler · Roman Festivals", romanFastiSource, "Literary witnesses and later fasti", "high for the February 13–21 conventional span"},
	{"roman-quinquatria-expanded", "Quinquatria · Expanded Games", time.March, 19, 5, "Later Republican Roman festival calendar", "Late Republic and Empire", "The archaic Quinquatrus expanded into five days associated especially with Minerva, artisans, teachers, and games.", "This layer shows a one-day archaic purification acquiring a wider urban festival life.", []string{"Offerings to Minerva historically", "Artisan and teacher observances", "Games after the opening day"}, "W. Warde Fowler · Roman Festivals", romanFastiSource, "Later fasti and literary witnesses", "high for the conventional March 19–23 span"},
	{"roman-megalesia", "Megalesia · Ludi Megalenses", time.April, 4, 7, "Roman Republican and Imperial games calendar", "From the second century BCE", "Games and performances honored Magna Mater, ending on her temple's dedication day.", "The imported cult became a highly visible part of Roman civic spectacle while retaining distinctive priesthood and music.", []string{"Theatrical performances historically", "Circus games", "Processions and music for Magna Mater"}, "W. Warde Fowler · Roman Festivals", romanFastiSource, "Later fasti; games calendar", "high for April 4–10"},
	{"roman-ludi-cereales", "Ludi Cereales", time.April, 12, 8, "Roman Republican and Imperial games calendar", "Attested by the late third century BCE", "Public games for Ceres led to the fixed Cerealia on April 19.", "The longer games layer should be read beside, not substituted for, the archaic fixed feast.", []string{"Theatre and circus games historically", "Public cult of Ceres, Liber, and Libera"}, "W. Warde Fowler · Roman Festivals", romanFastiSource, "Later fasti; games calendar", "high for April 12–19"},
	{"roman-floralia", "Floralia · Ludi Florae", time.April, 28, 6, "Roman Republican and Imperial games calendar", "Republican restoration and Imperial Rome", "Games for Flora crossed from late April into May and celebrated flowering, fertility, and theatrical license.", "Their exuberant urban character belongs to a later games calendar rather than the 45-row archaic ferial core.", []string{"Theatrical games historically", "Circus events", "Colorful clothing and public festivity in literary testimony"}, "W. Warde Fowler · Roman Festivals", romanFastiSource, "Later fasti; literary witnesses", "high for the Julian-era April 28–May 3 span"},
	{"roman-vestalia-season", "Vestalia · Sanctuary Season", time.June, 7, 9, "Roman Republican and Imperial ritual calendar", "Republican and Imperial Rome", "A broader sequence around the fixed June 9 Vestalia opened Vesta's inner sanctuary and ended with ritual cleansing.", "The span distinguishes the sanctuary's attested access and closure from the single named VESTALIA row.", []string{"Matronal access and offerings historically", "Milling and bakery observances", "Closing purification on June 15"}, "W. Warde Fowler · Roman Festivals", romanFastiSource, "Fasti and literary reconstruction", "moderate-to-high; component rites fall on securely attested dates"},
	{"roman-ludi-apollinares", "Ludi Apollinares", time.July, 6, 8, "Roman Republican and Imperial games calendar", "Annual from 208 BCE", "Apollo's public games grew from a one-day wartime institution into an eight-day annual span.", "The games joined vows for public safety, Greek-style performance, and Roman civic spectacle.", []string{"Theatre and circus games historically", "Public rites for Apollo"}, "W. Warde Fowler · Roman Festivals", romanFastiSource, "Historical games calendar", "high for July 6–13"},
	{"roman-ludi-victoriae-caesaris", "Ludi Victoriae Caesaris", time.July, 20, 11, "Late Republican and Imperial games calendar", "Annual from 45 BCE", "Games founded around Caesar's victory and Venus Genetrix ran through late July.", "This explicitly political festival belongs to the late Republican ruler-cult layer, not to archaic Roman religion.", []string{"Public games historically", "Cult honors for Venus Genetrix and Caesar"}, "Roman festival fasti summarized by Scullard", "https://penelope.uchicago.edu/encyclopaedia_romana/calendar/antiates.html", "Late Republican/Imperial fasti", "high for July 20–30 after institution"},
	{"roman-ludi-romani", "Ludi Romani · Great Games", time.September, 5, 15, "Roman Republican and Imperial games calendar", "Republican and Imperial Rome", "Rome's principal games honored Jupiter Optimus Maximus and culminated around the Ides before later circus days.", "Their changing length demonstrates why a dated historical layer must state which calendar phase it projects.", []string{"Theatrical and circus games historically", "Procession to the Circus Maximus", "Epulum Iovis"}, "Roman festival fasti summarized by Scullard", "https://penelope.uchicago.edu/encyclopaedia_romana/calendar/antiates.html", "Republican and Imperial games calendars", "moderate-to-high; length changed, September 5–19 is the later conventional span"},
	{"roman-ludi-augustales", "Ludi Augustales", time.October, 3, 10, "Roman Imperial cult calendar", "From 14 CE", "Games associated with the Augustalia honored the deified Augustus after his death.", "The span marks the transformation of political memory into recurring imperial cult and spectacle.", []string{"Public games historically", "Imperial-cult observance"}, "Roman festival fasti summarized by Scullard", "https://penelope.uchicago.edu/encyclopaedia_romana/calendar/antiates.html", "Imperial fasti", "high after establishment in 14 CE"},
	{"roman-ludi-victoriae-sullanae", "Ludi Victoriae Sullanae", time.October, 26, 7, "Late Republican games calendar", "Annual from 81 BCE", "Sulla's victory games crossed into November and converted civil-war victory into an annual public spectacle.", "The political origin is retained explicitly rather than treating the games as timeless Roman religion.", []string{"Public games historically", "Victory commemoration"}, "Roman festival fasti summarized by Scullard", "https://penelope.uchicago.edu/encyclopaedia_romana/calendar/antiates.html", "Late Republican and Imperial fasti", "high for October 26–November 1"},
	{"roman-ludi-plebeii", "Ludi Plebeii", time.November, 4, 14, "Roman Republican and Imperial games calendar", "Republican and Imperial Rome", "The Plebeian Games formed a long November cycle including the Epulum Iovis on the Ides.", "The festival linked plebeian political memory, Jupiter's cult, and public spectacle.", []string{"Theatre and circus games historically", "Epulum Iovis", "Public procession"}, "W. Warde Fowler · Roman Festivals", romanFastiSource, "Historical games calendar", "high for the later November 4–17 span"},
	{"byzantine-brumalia", "Brumalia · Constantinopolitan Season", time.November, 24, 24, "Late Roman / Byzantine popular and court calendar", "Late Antiquity through the Middle Byzantine period", "A late Roman winter custom became an alphabetically ordered season of banquets and acclamations at Constantinople.", "Its survival in a Christian empire reflects continuity and adaptation, not a simple unchanged pagan holiday or an official Christian feast.", []string{"Court banquets and acclamations historically", "Gift exchange", "Name-day gatherings by initial letter"}, "Fritz Graf · Roman Festivals in the Greek East", "https://doi.org/10.1017/CBO9781316274988", "Late-antique and Byzantine literary witnesses", "high for the conventional November 24–December 17 court sequence; practice varied by period"},
	{"roman-saturnalia-extended", "Saturnalia · Extended Popular Season", time.December, 17, 7, "Roman Imperial popular calendar", "Late Republic and Empire", "Popular and imperial practice expanded Saturnalia around its December 17 public rite, eventually through December 23.", "Gift exchange, gaming, relaxed dress, and temporary status play belong to the extended social season, not uniformly to every Roman period.", []string{"Feasting and gift exchange historically", "Dice and licensed play", "Relaxed dress", "Temporary ritual role reversal"}, "Macrobius, Saturnalia; Roman fasti", "https://penelope.uchicago.edu/Thayer/E/Roman/Texts/Macrobius/Saturnalia/1*.html", "Literary and Imperial calendar layer", "high for the late Imperial seven-day convention; duration differed earlier"},
}

// Later and regional Roman records need their own places. Keeping this mapping
// separate from the positional source manifests prevents a generic Rome label
// from erasing the local sanctuary or Constantinopolitan setting named by the
// evidence.
func romanLaterFixedLocation(catalogID string) (site, anchor string) {
	switch catalogID {
	case "roman-kalends-january":
		return "Rome and Roman civic communities", "Rome"
	case "roman-caristia":
		return "Roman households in Rome and the wider Roman world", "Rome"
	case "roman-matronalia":
		return "Rome, including Juno Lucina's sanctuary and Roman households", "Rome"
	case "roman-navigium-isidis":
		return "Rome and Mediterranean Isiac communities; Apuleius's procession is set at Cenchreae", "Rome for the conventional calendar date; Cenchreae for Apuleius's narrative"
	case "roman-anna-perenna":
		return "Anna Perenna's suburban gathering place beside the Via Flaminia, Rome", "Rome"
	case "roman-veneralia":
		return "Rome, at sanctuaries and baths named in the literary tradition", "Rome"
	case "roman-mercuralia":
		return "Mercury's spring near the Porta Capena, Rome", "Rome"
	case "roman-fors-fortuna":
		return "Fortuna's shrines downstream on the Tiber outside Rome", "Rome and the lower Tiber"
	case "roman-nonae-caprotinae":
		return "Rome and Latium; the received rite centers on a wild fig", "Rome"
	case "roman-nemoralia":
		return "Diana's Aventine temple at Rome and sanctuary at Lake Nemi near Aricia", "Lake Nemi / Aricia and Rome"
	case "roman-epulum-iovis-sep", "roman-epulum-iovis-nov":
		return "Capitoline sanctuary of Jupiter, Juno, and Minerva, Rome", "Rome"
	case "roman-augustalia":
		return "Rome, centered on the altar of Fortuna Redux", "Rome"
	case "roman-october-horse":
		return "Campus Martius and participating neighborhoods, Rome", "Rome"
	case "roman-natalis-invicti":
		return "Rome, as recorded in the fourth-century calendar", "Rome"
	case "roman-mundus-patet-august", "roman-mundus-patet-october", "roman-mundus-patet-november":
		return "Rome; the mundus's precise identity and location remain disputed", "Rome"
	case "roman-septimontium":
		return "The named hill districts of Rome", "Rome"
	default:
		return "", ""
	}
}

func romanSpanLocation(catalogID string) (site, anchor string) {
	switch catalogID {
	case "roman-compitalia-late-placement":
		return "Crossroads and neighborhoods of Rome and Roman communities", "Rome"
	case "roman-parentalia-season":
		return "Roman households, family tombs, and the city of Rome", "Rome"
	case "roman-quinquatria-expanded":
		return "Rome, including Minerva's Aventine sanctuary", "Rome"
	case "roman-megalesia":
		return "Palatine sanctuary of Magna Mater and public games venues, Rome", "Rome"
	case "roman-ludi-cereales":
		return "Rome, including Ceres's Aventine sanctuary and games venues", "Rome"
	case "roman-floralia":
		return "Rome, including Flora's sanctuary and public games venues", "Rome"
	case "roman-vestalia-season":
		return "Temple of Vesta and the Forum Romanum, Rome", "Rome"
	case "roman-ludi-apollinares":
		return "Rome, especially the Circus Maximus and Apollo's civic cult", "Rome"
	case "roman-ludi-victoriae-caesaris", "roman-ludi-augustales", "roman-ludi-victoriae-sullanae":
		return "Public games venues and imperial or victory cult sites at Rome", "Rome"
	case "roman-ludi-romani":
		return "Rome, including the Capitoline and Circus Maximus", "Rome"
	case "roman-ludi-plebeii":
		return "Rome, including the Circus Flaminius and Capitoline rites", "Rome"
	case "byzantine-brumalia":
		return "Constantinople's court and urban communities", "Constantinople"
	case "roman-saturnalia-extended":
		return "Rome and communities across the Roman world", "Rome"
	default:
		return "", ""
	}
}

// romanObservances projects nominal Roman month/day labels onto the selected
// modern civil year. It does not claim that a pre-Julian festival occurred on
// the same astronomical or proleptic Gregorian day in every ancient year.
func romanObservances(date time.Time) []Observance {
	var events []Observance
	for _, rule := range romanFixedCanon {
		if date.Month() != rule.Month || date.Day() != rule.Day {
			continue
		}
		event := baseObservance(rule.Name, PolytheistAncient, []string{"Ancient Roman civic and household cult", "Historical calendar study"}, "Archaic fixed Roman festival", rule.Summary, rule.Meaning, rule.Practices, nil, "W. Warde Fowler · The Roman Festivals of the Period of the Republic", romanFastiSource)
		applyRomanMetadata(&event, rule.CatalogID, rule.Month, rule.Day, "Fasti antiquissimi (named festival row)", "Roman Republic; later reception varied", "high for the nominal Roman date")
		event = singleOccurrence(event, date)
		event.ID = rule.CatalogID + "-" + date.Format("2006-01-02")
		events = append(events, event)
	}
	for _, rule := range romanLaterFixedFestivals {
		if date.Month() != rule.Month || date.Day() != rule.Day {
			continue
		}
		event := baseObservance(rule.Name, PolytheistAncient, []string{"Ancient Roman and regional historical calendar study"}, "Later or regional Roman observance", rule.Summary, rule.Meaning, rule.Practices, nil, rule.SourceName, rule.SourceURL)
		event.CatalogID = rule.CatalogID
		event.Origin = rule.Attestation
		event.ObservanceStatus = "Historical ancient observance; no universal Roman state-calendar obligation survives"
		event.Historical = true
		event.HistoricalNote = "This literary, regional, temple, popular, or imperial record is intentionally separate from the 45 named rows of the archaic fixed fasti."
		event.DateCertainty = rule.DateConfidence + "; no exact pre-Julian-to-Gregorian equivalence is asserted"
		event.CalendarCorpus = rule.Corpus
		event.NativeDateLabel = romanDateLabel(rule.Month, rule.Day)
		event.AttestationLayer = rule.Attestation
		event.Era = rule.Era
		event.Site, event.AnchorLocation = romanLaterFixedLocation(rule.CatalogID)
		event.ProjectionKind = "Nominal later Roman month/day projection"
		event.ProjectionStatus = "Documented historical calendar layer; non-converted study projection"
		event.DateConfidence = rule.DateConfidence
		event.DayBoundary = "Roman civil day rendered as a modern civil date"
		event.DateNote = "Later/regional tier: the nominal ancient month/day is displayed on the same-numbered modern civil date. Calendar phase and locality are stated in the record."
		event = singleOccurrence(event, date)
		event.ID = rule.CatalogID + "-" + date.Format("2006-01-02")
		events = append(events, event)
	}

	for _, rule := range romanSpanFestivals {
		start := dateAt(date.Year(), rule.StartMonth, rule.StartDay)
		if date.Before(start) || date.After(start.AddDate(0, 0, rule.Duration-1)) {
			continue
		}
		event := baseObservance(rule.Name, PolytheistAncient, []string{"Ancient Roman and late-antique historical calendar study"}, "Historical Roman festival span", rule.Summary, rule.Meaning, rule.Practices, nil, rule.SourceName, rule.SourceURL)
		event.CatalogID = rule.CatalogID
		event.Origin = rule.Attestation
		event.ObservanceStatus = "Historical; no longer a Roman state-calendar obligation"
		event.Historical = true
		event.HistoricalNote = "This later or extended span is shown separately from the 45 named rows of the archaic fixed fasti."
		event.DateCertainty = rule.DateConfidence
		event.CalendarCorpus = rule.Corpus
		event.NativeDateLabel = romanDateLabel(rule.StartMonth, rule.StartDay) + " (opening)"
		event.AttestationLayer = rule.Attestation
		event.Era = rule.Era
		event.Site, event.AnchorLocation = romanSpanLocation(rule.CatalogID)
		event.ProjectionKind = "Nominal historical calendar projection"
		event.ProjectionStatus = "Study projection; duration and civil-day equivalence can vary by period"
		event.DateConfidence = rule.DateConfidence
		event.DayBoundary = "Historical civil day rendered as a modern civil date"
		event.DateNote = "Historical layer: dates retain their named Roman/Julian calendar position and are not a claim of exact pre-Julian seasonal equivalence."
		event = spanOccurrence(event, date, start, rule.Duration)
		event.ID = rule.CatalogID + "-" + start.Format("2006-01-02")
		events = append(events, event)
	}
	return events
}

func applyRomanMetadata(event *Observance, catalogID string, month time.Month, day int, attestation, era, confidence string) {
	event.CatalogID = catalogID
	event.Origin = "Fasti antiquissimi; reconstructed from surviving Roman fasti"
	event.ObservanceStatus = "Historical Roman state or household rite; ancient civic obligation discontinued"
	event.Historical = true
	event.HistoricalNote = "Ritual details summarize ancient testimony for study; they are not prescriptions for modern practice."
	event.DateCertainty = confidence + "; no exact pre-Julian-to-Gregorian equivalence is asserted"
	event.CalendarCorpus = "Roman Republican fasti · archaic fixed canon"
	event.NativeDateLabel = romanDateLabel(month, day)
	event.AttestationLayer = attestation
	event.Era = era
	event.Site = "Rome"
	event.ProjectionKind = "Nominal Roman month/day projection"
	event.ProjectionStatus = "High-confidence calendar label; non-converted study projection"
	event.DateConfidence = confidence
	event.AnchorLocation = "Rome"
	event.DayBoundary = "Roman civil day rendered as a modern civil date"
	event.DateNote = "The Roman nominal month/day is displayed on the same-numbered modern civil date. Before the Julian reform, intercalation and calendar drift prevent treating this as an exact proleptic Gregorian anniversary."
}

// imperialChristianObservances adds only recurring Byzantine or Holy Roman
// imperial/local observances with a documented calendar rule. It intentionally
// does not turn coronations, diets, battles, or constitutional acts into annual
// holidays merely because their historical dates are known.
func imperialChristianObservances(date time.Time) []Observance {
	var events []Observance
	addFixed := func(catalogID, name string, month time.Month, day int, communities []string, category, summary, meaning string, practices []string, sourceName, sourceURL, corpus, era, site, status, note, confidence string) {
		if date.Month() != month || date.Day() != day {
			return
		}
		event := baseObservance(name, Christianity, communities, category, summary, meaning, practices, nil, sourceName, sourceURL)
		event.CatalogID = catalogID
		event.Origin = corpus
		event.ObservanceStatus = status
		event.Historical = false
		event.HistoricalNote = note
		event.DateCertainty = confidence
		event.CalendarCorpus = corpus
		event.NativeDateLabel = date.Format("January 2")
		event.AttestationLayer = "Documented recurring imperial, civic, or local liturgical observance"
		event.Era = era
		event.Site = site
		event.ProjectionKind = "Recurring calendar rule"
		event.ProjectionStatus = "Historical layer; consult the named living community for current practice"
		event.DateConfidence = confidence
		event.AnchorLocation = site
		event.DayBoundary = "Liturgical/civil day; vigils and local usages can begin earlier"
		event = singleOccurrence(event, date)
		event.ID = catalogID + "-" + date.Format("2006-01-02")
		events = append(events, event)
	}

	addFixed("byzantine-constantinople-dedication", "Dedication of Constantinople", time.May, 11,
		[]string{"Byzantine civic calendar", "Eastern Orthodox communities retaining the commemoration"}, "Byzantine civic and liturgical commemoration",
		"Constantinople's dedication in 330 became an annual birthday of the imperial capital.",
		"The commemoration joined the city's Roman civic identity, imperial foundation story, and later Christian protection of the Theotokos.",
		[]string{"Civic celebrations historically", "Processions and church commemoration", "Forty-day inaugural celebration in the foundation tradition"},
		"Orthodox Church in America · Founding of Constantinople", "https://www.oca.org/saints/all-lives/2025/05/11", "Byzantine Constantinopolitan calendar", "From 330 CE; Byzantine and later Orthodox reception", "Constantinople", "Living as a limited Orthodox commemoration; imperial civic festival discontinued", "The foundation narrative contains ceremonial and legendary layers; May 11 as the dedication date is strongly received.", "high for May 11 in the Julian/Byzantine calendar; displayed nominally")

	addFixed("byzantine-indiction", "Beginning of the Indiction · Byzantine New Year", time.September, 1,
		[]string{"Byzantine Empire historically", "Eastern Orthodox churches"}, "Byzantine civil and ecclesiastical new year",
		"September 1 opened the Constantinopolitan indiction and eventually the Byzantine civil and church year.",
		"A fiscal dating cycle became a prayerful threshold for the year, seasons, harvest, and public order.",
		[]string{"Imperial year dating historically", "Divine Liturgy and hymns", "Prayers for favorable seasons and the work of creation"},
		"Orthodox Church in America · Church New Year", "https://www.oca.org/saints/lives/2047/09/01/501-church-new-year", "Byzantine indiction calendar", "Late Roman and Byzantine Empire; living Orthodox reception", "Constantinople and Eastern Orthodox churches", "Still observed as the Orthodox church-year opening; Byzantine civil use discontinued", "Different indiction systems began on different dates; this entry is specifically the Constantinopolitan September indiction.", "high for the Constantinopolitan September 1 rule")

	lastJanuaryDay := dateAt(date.Year(), time.January, 31)
	lastSundayInJanuary := lastJanuaryDay.AddDate(0, 0, -int(lastJanuaryDay.Weekday()))
	if sameDay(date, lastSundayInJanuary) {
		event := baseObservance(
			"Karlsfest · Feast of Charlemagne at Aachen",
			Christianity,
			[]string{"Aachen Cathedral and historic local cult"},
			"Holy Roman imperial-local feast",
			"Aachen celebrates Karlsfest on the last Sunday in January in commemoration of Charlemagne's death on January 28, 814.",
			"The feast expresses Aachen's identity as Charlemagne's burial and coronation city while never becoming a universal Roman-calendar saint's day.",
			[]string{"Solemn high Mass on the last Sunday in January", "Medieval and early-medieval chants", "Veneration at Charlemagne's shrine"},
			nil,
			"Aachen Cathedral · Charlemagne",
			"https://www.aachenerdom.de/en/a-place-of-history/charlemagne/",
		)
		event.CatalogID = "hre-karlsfest"
		event.Origin = "Holy Roman Empire · Aachen local calendar"
		event.ObservanceStatus = "Living local cathedral observance on the last Sunday in January; never a universal canonization or empire-wide holy day"
		event.Historical = false
		event.HistoricalNote = "Frederick Barbarossa arranged Charlemagne's 1165 canonization through antipope Paschal III; Rome never recognized it as a universal papal canonization. The local cult endured, and Aachen Cathedral transfers the principal celebration from the January 28 death anniversary to the last Sunday in January."
		event.DateCertainty = "high for the living last-Sunday rule and the January 28 death anniversary"
		event.CalendarCorpus = "Holy Roman Empire · Aachen local calendar"
		event.NativeDateLabel = "Last Sunday in January; commemorates January 28, 814"
		event.AttestationLayer = "Documented recurring local cathedral observance"
		event.Era = "From the later twelfth century; living local reception"
		event.Site = "Aachen Cathedral"
		event.ProjectionKind = "Last Sunday in January"
		event.ProjectionStatus = "Living local rule documented by Aachen Cathedral"
		event.DateConfidence = "high for the living last-Sunday rule and the January 28 death anniversary"
		event.AnchorLocation = "Aachen"
		event.DayBoundary = "Liturgical/civil day; the cathedral publishes current service times"
		event.DateNote = "January 28 is the nominal death and feast anniversary; the living solemn high Mass is placed on the last Sunday in January."
		event = singleOccurrence(event, date)
		event.ID = event.CatalogID + "-" + date.Format("2006-01-02")
		events = append(events, event)
	}

	addFixed("hre-cunigunde-bamberg", "Solemnity of Saint Cunigunde · Bamberg", time.March, 3,
		[]string{"Archdiocese of Bamberg", "Catholic communities honoring Saint Cunigunde"}, "Holy Roman imperial-local saint feast",
		"Bamberg keeps March 3 as the solemnity of Cunigunde, empress, diocesan patron, and cofounder of the see with Henry II.",
		"The feast receives an imperial ruler through a local church's living saint calendar, not as a civic holiday imposed across the Empire.",
		[]string{"Solemn Mass", "Pilgrimage and women's diocesan observances", "Veneration at the imperial couple's tomb in Bamberg Cathedral"},
		"Archdiocese of Bamberg · Saint Cunigunde", "https://heilige.erzbistum-bamberg.de/heilige-kunigunde", "Holy Roman Empire · Bamberg diocesan calendar", "Canonized in 1200; living Bamberg observance", "Bamberg", "Living diocesan solemnity; not an empire-wide medieval feast", "The source calls March 3 a Hochfest in Bamberg. Dates and rank in other Catholic calendars can differ.", "high for the Bamberg March 3 solemnity")

	addFixed("hre-henry-ii-bamberg", "Feast of Saint Henry II · Bamberg", time.July, 13,
		[]string{"Archdiocese of Bamberg", "Catholic communities honoring Saint Henry II"}, "Holy Roman imperial-local saint feast",
		"Bamberg marks Henry II, the canonized emperor and founder of its bishopric, on his July 13 death anniversary.",
		"The observance joins local diocesan identity and imperial memory while remaining a saint's feast rather than an anniversary of coronation or rule.",
		[]string{"Mass and diocesan festival", "Pilgrimage", "Veneration at the imperial couple's tomb in Bamberg Cathedral"},
		"Vatican News · Saint Henry II", "https://www.vaticannews.va/en/saints/07/13/st-henry-ii-emperor.html", "Holy Roman Empire · Bamberg diocesan calendar", "Canonized in 1146; living Catholic reception", "Bamberg and Catholic calendars", "Living saint's feast with particular Bamberg prominence; not an empire-wide public holiday", "July 13 is the received feast and death date. Local rank and transfer rules vary.", "high for July 13")

	addFixed("hre-augsburg-confession", "Augsburg Confession Day", time.June, 25,
		[]string{"Lutheran churches", "Augsburg and Reformation-history communities"}, "Holy Roman confessional commemoration",
		"Lutheran communities commemorate the presentation of the Augsburg Confession before Emperor Charles V on June 25, 1530.",
		"The day remembers a defining public confession within the religious and constitutional life of the Holy Roman Empire.",
		[]string{"Reading or study of the Augsburg Confession", "Lutheran worship", "Ecumenical historical reflection"},
		"Lutheran Church—Missouri Synod · Presentation of the Augsburg Confession", "https://resources.lcms.org/worship-planning/lectionary-summary-for-augsburg-confession-commemoration/", "Holy Roman Empire · Lutheran confessional calendar", "From 1530; later Lutheran commemoration", "Augsburg and Lutheran churches", "Still commemorated in parts of the Lutheran tradition; not an empire-wide public holiday", "The historical presentation date is secure; the form and prominence of annual observance vary by church and region.", "high for June 25")

	addFixed("hre-augsburg-peace-festival", "Augsburg High Peace Festival · Hohes Friedensfest", time.August, 8,
		[]string{"City of Augsburg", "Ecumenical and civic communities"}, "Holy Roman civic-religious peace festival",
		"Augsburg's Protestants first celebrated restored freedom of worship on August 8, 1650, after the settlements ending the Thirty Years' War.",
		"A confessional thanksgiving became a continuing civic festival of peace and religious coexistence.",
		[]string{"Ecumenical worship", "Civic peace program", "Public holiday within Augsburg"},
		"City of Augsburg · High Peace Festival", "https://www.augsburg.de/kultur/friedensstadt-augsburg/hohes-friedensfest", "Holy Roman Empire · Free Imperial City of Augsburg", "From 1650; living civic observance", "Augsburg", "Still observed; a legal public holiday within the city of Augsburg", "The festival is local to Augsburg, not a general holiday of the former Empire.", "high for August 8")

	// Original imperial feast established for Germany and Bohemia: Friday after
	// the Easter octave (Friday after Quasimodogeniti), i.e. Easter + 12 days.
	// Later local Roman-calendar feasts used a different Lent placement and are
	// not used to back-project Charles IV's medieval observance.
	easter := westernEaster(date.Year())
	lanceDay := easter.AddDate(0, 0, 12)
	if sameDay(date, lanceDay) {
		event := baseObservance("Feast of the Holy Lance and Nails", Christianity, []string{"Holy Roman Empire historically", "Medieval Bohemia and German dioceses"}, "Holy Roman imperial relic feast", "At Charles IV's request, the papacy authorized an annual feast and public showing centered on the imperial Holy Lance and a relic identified as a nail of the Cross.", "Relic veneration, pilgrimage, indulgence, and imperial representation converged in a specifically medieval imperial feast.", []string{"Mass and office historically", "Public relic showing", "Pilgrimage and almsgiving"}, []string{"John 19:31–37"}, "Blood in Stone and the Second Coming · Zeitschrift für Kirchengeschichte", holyLanceSource)
		event.CatalogID = "hre-holy-lance-nails"
		event.Origin = "Papal grant of 1354 at Charles IV's request"
		event.ObservanceStatus = "Historical imperial/local feast; later calendars moved or discontinued it"
		event.Historical = true
		event.HistoricalNote = "The original medieval rule was Friday after the Easter octave. Later Passion-feast supplements often placed a feast of the same title on Friday after the first Sunday of Lent; those are distinct calendar phases."
		event.DateCertainty = "high for the original annual Friday-after-Easter-octave rule"
		event.CalendarCorpus = "Holy Roman Empire · Charles IV relic calendar"
		event.NativeDateLabel = "Feria VI post octavam Paschae / Friday after Quasimodogeniti"
		event.AttestationLayer = "1354 papal indulgence and imperial relic-display practice"
		event.Era = "From 1354 in Bohemia and parts of the Empire"
		event.Site = "Prague initially; later Nuremberg and participating dioceses"
		event.ProjectionKind = "Western Easter-relative historical rule"
		event.ProjectionStatus = "Computed from Gregorian Easter for study; medieval communities used their contemporary Julian computus"
		event.DateConfidence = "high for rule; projected modern occurrence"
		event.AnchorLocation = "Prague"
		event.DayBoundary = "Liturgical day"
		event.DateNote = "Modern display uses Easter + 12. Before Gregorian reform this recurring rule was calculated from the Julian Easter date, so this is not a proleptic anniversary conversion."
		event = singleOccurrence(event, date)
		event.ID = event.CatalogID + "-" + date.Format("2006-01-02")
		events = append(events, event)
	}

	// The first Sunday of Orthodox Great Lent is six weeks before Pascha.
	orthodoxySunday := orthodoxEaster(date.Year()).AddDate(0, 0, -42)
	if sameDay(date, orthodoxySunday) {
		event := baseObservance("Sunday of Orthodoxy · Triumph of the Icons", Christianity, []string{"Eastern Orthodox churches", "Byzantine historical calendar"}, "Byzantine historical feast", "The first Sunday of Great Lent commemorates the restoration of icons at Constantinople in 843 under Empress Theodora, Michael III, and Patriarch Methodios.", "The living feast interprets icon veneration through the incarnation while retaining memory of Byzantine iconoclasm and imperial power.", []string{"Divine Liturgy", "Procession with icons", "Synodikon or local profession of faith"}, []string{"John 1:43–51", "Hebrews 11:24–12:2"}, "Orthodox Church in America · End of Iconoclasm", "https://www.oca.org/orthodoxy/the-orthodox-faith/church-history/ninth-century/the-end-of-iconoclasm")
		event.CatalogID = "byzantine-sunday-orthodoxy"
		event.Origin = "Constantinopolitan restoration of icons in 843"
		event.ObservanceStatus = "Still observed throughout Eastern Orthodoxy"
		event.Historical = false
		event.HistoricalNote = "A living Orthodox feast with a specific Byzantine imperial-historical origin; it is not merely a secular victory anniversary."
		event.DateCertainty = "high; first Sunday of Orthodox Great Lent"
		event.CalendarCorpus = "Byzantine / Eastern Orthodox Paschal cycle"
		event.NativeDateLabel = "First Sunday of Great Lent"
		event.AttestationLayer = "Byzantine synodal and liturgical commemoration"
		event.Era = "Annual from 843; living observance"
		event.Site = "Constantinople; Eastern Orthodox churches worldwide"
		event.ProjectionKind = "Orthodox Pascha-relative rule"
		event.ProjectionStatus = "Calculated with the app's Julian Paschal computus conversion"
		event.DateConfidence = "high for 1900–2100 app computus range"
		event.AnchorLocation = "Constantinople"
		event.DayBoundary = "Liturgical day"
		event = singleOccurrence(event, date)
		event.ID = event.CatalogID + "-" + date.Format("2006-01-02")
		events = append(events, event)
	}

	// Corpus Christi already exists in the general Christian layer. This overlay
	// records its distinctive public, civic-processional life inside the Empire.
	corpusChristi := easter.AddDate(0, 0, 60)
	if sameDay(date, corpusChristi) {
		event := baseObservance("Corpus Christi · Imperial-City Processional Layer", Christianity, []string{"Catholic territories and cities of the Holy Roman Empire"}, "Holy Roman civic-liturgical observance", "Corpus Christi processions became major public rites in many imperial cities, with civic authorities, guilds, clergy, and neighborhoods participating.", "The overlay records how a universal feast took a distinct civic form in the Empire; it does not claim uniform practice across its Catholic and Protestant territories.", []string{"Mass historically", "Eucharistic procession", "Civic and guild participation", "Street and market decoration"}, []string{"1 Corinthians 11:23–26"}, "Aachen historical city accounts · church feasts", "https://de.wikisource.org/wiki/Aachener_Stadtrechnungen_aus_dem_XIV._Jahrhundert/Kirchenfeste_und_Betheiligung")
		event.CatalogID = "hre-corpus-christi-civic"
		event.Origin = "Medieval Catholic feast received through local imperial-city calendars"
		event.ObservanceStatus = "Historical imperial-city overlay on a living Catholic feast; no continuous civic form asserted"
		event.Historical = true
		event.HistoricalNote = "This is a cultural overlay on the general Corpus Christi entry, not a second theological feast. Protestant territories rejected or altered the procession after the Reformation."
		event.DateCertainty = "high for the Thursday after Trinity Sunday; local transfers possible"
		event.CalendarCorpus = "Holy Roman Empire · Catholic civic calendars"
		event.NativeDateLabel = "Thursday after Trinity Sunday"
		event.AttestationLayer = "Municipal accounts and local liturgical custom"
		event.Era = "Later Middle Ages through the early modern Empire"
		event.Site = "Catholic imperial cities and territories"
		event.ProjectionKind = "Western Easter-relative rule"
		event.ProjectionStatus = "Calculated modern recurrence; local calendars may transfer the feast"
		event.DateConfidence = "high for rule, variable for local civic form"
		event.AnchorLocation = "Aachen as the cited example"
		event.DayBoundary = "Liturgical/civil day"
		event = singleOccurrence(event, date)
		event.ID = event.CatalogID + "-" + date.Format("2006-01-02")
		events = append(events, event)
	}

	return events
}

func romanDateLabel(month time.Month, day int) string {
	abbreviations := [...]string{"", "Ian.", "Feb.", "Mart.", "Apr.", "Mai.", "Iun.", "Quint.", "Sext.", "Sept.", "Oct.", "Nov.", "Dec."}
	monthLengths := [...]int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	nones := 5
	if month == time.March || month == time.May || month == time.July || month == time.October {
		nones = 7
	}
	ides := nones + 8
	abbr := abbreviations[int(month)]
	if day == 1 {
		return "Kal. " + abbr
	}
	if day <= nones {
		return romanAnteDate(nones-day+1, "Non. "+abbr)
	}
	if day <= ides {
		return romanAnteDate(ides-day+1, "Id. "+abbr)
	}
	nextMonth := month + 1
	if nextMonth > time.December {
		nextMonth = time.January
	}
	return romanAnteDate(monthLengths[int(month)]-day+2, "Kal. "+abbreviations[int(nextMonth)])
}

func romanAnteDate(inclusiveCount int, target string) string {
	if inclusiveCount == 1 {
		return target
	}
	if inclusiveCount == 2 {
		return "prid. " + target
	}
	return "a.d. " + smallRomanNumeral(inclusiveCount) + " " + target
}

func smallRomanNumeral(value int) string {
	values := []struct {
		value int
		label string
	}{{10, "X"}, {9, "IX"}, {5, "V"}, {4, "IV"}, {1, "I"}}
	result := ""
	for _, item := range values {
		for value >= item.value {
			result += item.label
			value -= item.value
		}
	}
	return result
}
