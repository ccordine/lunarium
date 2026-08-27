package calendar

import (
	"fmt"
	"strings"
	"time"
)

const (
	lateFastAppendixSourceURL = "https://www.sefaria.org/Megillat_Ta%27anit%2C_Adar.20?lang=bi"
	lateFastHalakhotURL       = "https://www.sefaria.org/Halakhot_Gedolot.18.12?lang=bi"
	lateFastTurURL            = "https://www.sefaria.org/Tur%2C_Orach_Chayim.580?lang=bi"
	lateFastShulchanArukhURL  = "https://www.sefaria.org/Shulchan_Arukh%2C_Orach_Chayim.580.2?lang=bi"
	yomaGerizimSourceURL      = "https://www.sefaria.org/Yoma.69a.9-13?lang=bi"
	rainFastSourceURL         = "https://www.sefaria.org/Mishnah_Ta%27anit.1.4-7?lang=bi"
	maamadotSourceURL         = "https://www.sefaria.org/Mishnah_Ta%27anit.4.2-3?lang=bi"
	birkatHachamaSourceURL    = "https://www.sefaria.org/Berakhot.59b.1?lang=bi"
	birkatLevanahSourceURL    = "https://www.sefaria.org/Sanhedrin.41b.12-42a.5?lang=bi"
	fourFastsSourceURL        = "https://www.sefaria.org/Rosh_Hashanah.18b.1-3?lang=bi"
	torahCyclesSourceURL      = "https://www.sefaria.org/Leviticus.25.1-24?lang=bi"
)

const jewishCorpusInclusionRule = "The bounded Torah/Talmud corpus includes explicit recurring Torah times, counts, and sabbatical institutions; calendrical dates, conditional schedules, and Temple institutions in Mishnah Moed; and the named annual, monthly, weekly, or multi-year observances and timing disputes enumerated by this versioned manifest from the Babylonian and Jerusalem Talmuds. It separately preserves complete finite witnesses such as the 35-entry Aramaic Megillat Ta'anit core, the 26-date printed Ma'amar Aharon appendix, and the nine wood-offering dates. Personal vows, undated narrative anniversaries, daily liturgy without a calendar trigger, and later local customs are outside this completeness claim unless separately named; completeness is asserted against these explicit manifests, not every statement about time in rabbinic literature."

const lateFastAppendixCaveat = "The Ma'amar Aharon is a later Hebrew appendix, not part of the original 35-entry Aramaic Megillat Ta'anit. 'Megillat Ta'anit Batra' is a modern scholarly label for this separate, textually unstable fast-list tradition, with late-antique antecedents and an influential Geonic recension. This catalog follows the Hebrew date readings in the Warsaw 1874 printed witness presented by Sefaria; Halakhot Gedolot, Tur, Shulchan Arukh, and other witnesses contain date and content variants."

type jewishCorpusRule struct {
	CatalogID      string
	Name           string
	Month          int
	Day            int
	Adar           bool
	Category       string
	Summary        string
	Note           string
	Status         string
	CalendarCorpus string
	Attestation    string
	Era            string
	SourceName     string
	SourceURL      string
	StartsAtSunset bool
}

// This is the complete 26-date Ma'amar Aharon ("final discourse") in the
// cited Warsaw 1874 Hebrew witness. It is intentionally separate from the
// original Aramaic Megillat Ta'anit and from the living four destruction
// fasts. The two י׳ readings are encoded as 10, despite an English rendering
// on the source page that displays "second" for those Hebrew numerals.
var lateAppendedFastRules = []jewishCorpusRule{
	{"mtb-nisan-aaron-sons", "Late Fast Appendix · Death of Aaron's Sons", 1, 1, false, "Late appended commemorative fast", "The appendix assigns Nisan 1 to the deaths of Aaron's sons.", "This is a later memorial-date tradition, not a separate Torah-appointed annual fast.", "Late appendix fast; not generally observed", "Late appended fast list · Ma'amar Aharon", "Printed Hebrew final discourse appended to Megillat Ta'anit", "Post-Talmudic textual layer; underlying traditions vary", "Megillat Ta'anit · Ma'amar Aharon", lateFastAppendixSourceURL, false},
	{"mtb-nisan-miriam", "Late Fast Appendix · Death of Miriam", 1, 10, false, "Late appended commemorative fast", "The appendix assigns Nisan 10 to Miriam's death and the disappearance of the well associated with her.", "The Hebrew witness reads י׳, ten; the displayed English community translation on the cited page renders the numeral incorrectly as second.", "Late appendix fast; not generally observed", "Late appended fast list · Ma'amar Aharon", "Printed Hebrew final discourse appended to Megillat Ta'anit", "Post-Talmudic textual layer; underlying traditions vary", "Megillat Ta'anit · Ma'amar Aharon", lateFastAppendixSourceURL, false},
	{"mtb-nisan-joshua", "Late Fast Appendix · Death of Joshua", 1, 26, false, "Late appended commemorative fast", "The appendix assigns Nisan 26 to the death of Joshua son of Nun.", "Other later memorial lists preserve variants; this entry follows the named printed witness.", "Late appendix fast; not generally observed", "Late appended fast list · Ma'amar Aharon", "Printed Hebrew final discourse appended to Megillat Ta'anit", "Post-Talmudic textual layer; underlying traditions vary", "Megillat Ta'anit · Ma'amar Aharon", lateFastAppendixSourceURL, false},
	{"mtb-iyar-eli", "Late Fast Appendix · Eli, His Sons, and the Ark", 2, 10, false, "Late appended commemorative fast", "The appendix assigns Iyar 10 to the deaths of Eli and his sons and the capture of the Ark.", "The Hebrew witness reads י׳, ten; the displayed English community translation on the cited page renders the numeral incorrectly as second.", "Late appendix fast; not generally observed", "Late appended fast list · Ma'amar Aharon", "Printed Hebrew final discourse appended to Megillat Ta'anit", "Post-Talmudic textual layer; underlying traditions vary", "Megillat Ta'anit · Ma'amar Aharon", lateFastAppendixSourceURL, false},
	{"mtb-iyar-samuel", "Late Fast Appendix · Death of Samuel", 2, 29, false, "Late appended commemorative fast", "The appendix assigns Iyar 29 to Samuel's death and Israel's mourning for him.", "Later sources disagree over the memorial date; this entry remains witness-specific.", "Late appendix fast; not generally observed", "Late appended fast list · Ma'amar Aharon", "Printed Hebrew final discourse appended to Megillat Ta'anit", "Post-Talmudic textual layer; underlying traditions vary", "Megillat Ta'anit · Ma'amar Aharon", lateFastAppendixSourceURL, false},
	{"mtb-sivan-first-fruits", "Late Fast Appendix · First-Fruits Procession Halted", 3, 23, false, "Late appended commemorative fast", "The appendix says that first fruits ceased to be brought to Jerusalem under Jeroboam on Sivan 23.", "This is the appendix's historical attribution, preserved without treating its chronology as independently established.", "Late appendix fast; not generally observed", "Late appended fast list · Ma'amar Aharon", "Printed Hebrew final discourse appended to Megillat Ta'anit", "Post-Talmudic textual layer; underlying traditions vary", "Megillat Ta'anit · Ma'amar Aharon", lateFastAppendixSourceURL, false},
	{"mtb-sivan-three-sages", "Late Fast Appendix · Deaths of Three Sages", 3, 25, false, "Late appended commemorative fast", "The appendix associates Sivan 25 with the deaths of Rabban Shimon ben Gamliel, Rabbi Ishmael ben Elisha, and Rabbi Hanina the deputy high priest.", "The list combines martyr traditions whose historical synchronization should not be assumed.", "Late appendix fast; not generally observed", "Late appended fast list · Ma'amar Aharon", "Printed Hebrew final discourse appended to Megillat Ta'anit", "Post-Talmudic textual layer; underlying traditions vary", "Megillat Ta'anit · Ma'amar Aharon", lateFastAppendixSourceURL, false},
	{"mtb-sivan-hanina", "Late Fast Appendix · Death of Hanina ben Teradion", 3, 27, false, "Late appended commemorative fast", "The appendix associates Sivan 27 with Rabbi Hanina ben Teradion's execution together with a Torah scroll.", "The date belongs to the later appendix's martyr calendar.", "Late appendix fast; not generally observed", "Late appended fast list · Ma'amar Aharon", "Printed Hebrew final discourse appended to Megillat Ta'anit", "Post-Talmudic textual layer; underlying traditions vary", "Megillat Ta'anit · Ma'amar Aharon", lateFastAppendixSourceURL, false},
	{"mtb-tammuz-seventeen", "Late Fast Appendix · Seventeenth of Tammuz", 4, 17, false, "Late appended commemorative fast", "The appendix collects the broken tablets, cessation of the daily offering, burning of Torah, and placement of an image at the Temple under Tammuz 17.", "This appendix layer overlaps the living Fast of the Seventeenth of Tammuz and does not replace its separate observance card.", "Date remains an active fast; appendix is a separate textual layer", "Late appended fast list · Ma'amar Aharon", "Printed Hebrew final discourse appended to Megillat Ta'anit", "Post-Talmudic textual layer; underlying traditions vary", "Megillat Ta'anit · Ma'amar Aharon", lateFastAppendixSourceURL, false},
	{"mtb-av-aaron", "Late Fast Appendix · Death of Aaron", 5, 1, false, "Late appended commemorative fast", "The appendix assigns Av 1 to the death of Aaron the high priest.", "The underlying date is also stated in Numbers 33:38, but the annual fast belongs to the later memorial list.", "Late appendix fast; not generally observed", "Late appended fast list · Ma'amar Aharon", "Printed Hebrew final discourse appended to Megillat Ta'anit", "Post-Talmudic textual layer; underlying traditions vary", "Megillat Ta'anit · Ma'amar Aharon", lateFastAppendixSourceURL, false},
	{"mtb-av-nine", "Late Fast Appendix · Ninth of Av", 5, 9, false, "Late appended commemorative fast", "The appendix gathers the wilderness decree, both Temple destructions, the fall of Beitar, and the plowing of the city under Av 9.", "This appendix layer overlaps the living Tisha B'Av observance and is retained separately as textual history.", "Date remains an active fast; appendix is a separate textual layer", "Late appended fast list · Ma'amar Aharon", "Printed Hebrew final discourse appended to Megillat Ta'anit", "Post-Talmudic textual layer; underlying traditions vary", "Megillat Ta'anit · Ma'amar Aharon", lateFastAppendixSourceURL, false},
	{"mtb-av-western-light", "Late Fast Appendix · Western Light Extinguished", 5, 18, false, "Late appended commemorative fast", "The appendix says the western sanctuary light was extinguished in the days of Ahaz on Av 18.", "The terse claim is retained as the witness presents it; its historical referent is uncertain.", "Late appendix fast; not generally observed", "Late appended fast list · Ma'amar Aharon", "Printed Hebrew final discourse appended to Megillat Ta'anit", "Post-Talmudic textual layer; underlying traditions vary", "Megillat Ta'anit · Ma'amar Aharon", lateFastAppendixSourceURL, false},
	{"mtb-elul-spies", "Late Fast Appendix · Death of the Spies", 6, 7, false, "Late appended commemorative fast", "The appendix assigns Elul 7 to the deaths of those who brought an evil report about the land.", "This is a late memorial-date tradition rather than an explicit biblical anniversary.", "Late appendix fast; not generally observed", "Late appended fast list · Ma'amar Aharon", "Printed Hebrew final discourse appended to Megillat Ta'anit", "Post-Talmudic textual layer; underlying traditions vary", "Megillat Ta'anit · Ma'amar Aharon", lateFastAppendixSourceURL, false},
	{"mtb-tishrei-gedaliah", "Late Fast Appendix · Death of Gedaliah", 7, 3, false, "Late appended commemorative fast", "The appendix places the killing of Gedaliah and the Jews with him at Mizpah on Tishrei 3.", "This appendix layer overlaps the living Fast of Gedaliah, whose observed day can be postponed from Shabbat.", "Date remains an active fast; appendix is a separate textual layer", "Late appended fast list · Ma'amar Aharon", "Printed Hebrew final discourse appended to Megillat Ta'anit", "Post-Talmudic textual layer; underlying traditions vary", "Megillat Ta'anit · Ma'amar Aharon", lateFastAppendixSourceURL, false},
	{"mtb-tishrei-five", "Late Fast Appendix · Tishrei Five Memorial", 7, 5, false, "Late appended commemorative fast", "The appendix links Tishrei 5 to the deaths of twenty Israelites and to Rabbi Akiva's imprisonment and death.", "The compressed notice conjoins traditions and is textually or historically difficult; no reconstruction is asserted here.", "Late appendix fast; not generally observed", "Late appended fast list · Ma'amar Aharon", "Printed Hebrew final discourse appended to Megillat Ta'anit", "Post-Talmudic textual layer; underlying traditions vary", "Megillat Ta'anit · Ma'amar Aharon", lateFastAppendixSourceURL, false},
	{"mtb-tishrei-seven", "Late Fast Appendix · Decree of Sword and Famine", 7, 7, false, "Late appended commemorative fast", "The appendix tersely associates Tishrei 7 with a decree of sword and famine against the ancestors.", "The witness gives no secure historical referent for this compressed notice.", "Late appendix fast; not generally observed", "Late appended fast list · Ma'amar Aharon", "Printed Hebrew final discourse appended to Megillat Ta'anit", "Post-Talmudic textual layer; underlying traditions vary", "Megillat Ta'anit · Ma'amar Aharon", lateFastAppendixSourceURL, false},
	{"mtb-tishrei-atonement", "Late Fast Appendix · Atonement for the Golden Calf", 7, 10, false, "Late appended commemorative fast", "The appendix assigns Tishrei 10 to atonement for the episode of the Golden Calf.", "Yom Kippur is a Torah-appointed holy day; this card preserves only the appendix's additional rationale.", "Living holy day; appendix rationale is a separate textual layer", "Late appended fast list · Ma'amar Aharon", "Printed Hebrew final discourse appended to Megillat Ta'anit", "Post-Talmudic textual layer; underlying traditions vary", "Megillat Ta'anit · Ma'amar Aharon", lateFastAppendixSourceURL, false},
	{"mtb-cheshvan-zedekiah", "Late Fast Appendix · Zedekiah Blinded", 8, 6, false, "Late appended commemorative fast", "The appendix associates Cheshvan 6 with Zedekiah's sons being killed before him and his blinding.", "The biblical episode is not assigned this annual date in its narrative; the date belongs to the appendix.", "Late appendix fast; not generally observed", "Late appended fast list · Ma'amar Aharon", "Printed Hebrew final discourse appended to Megillat Ta'anit", "Post-Talmudic textual layer; underlying traditions vary", "Megillat Ta'anit · Ma'amar Aharon", lateFastAppendixSourceURL, false},
	{"mtb-kislev-jehoiakim", "Late Fast Appendix · Jehoiakim Burns the Scroll", 9, 7, false, "Late appended commemorative fast", "The appendix assigns Kislev 7 to Jehoiakim's burning of the scroll written by Baruch from Jeremiah's dictation.", "This card records the appendix's annual date, not an independently dated anniversary in Jeremiah.", "Late appendix fast; not generally observed", "Late appended fast list · Ma'amar Aharon", "Printed Hebrew final discourse appended to Megillat Ta'anit", "Post-Talmudic textual layer; underlying traditions vary", "Megillat Ta'anit · Ma'amar Aharon", lateFastAppendixSourceURL, false},
	{"mtb-tevet-greek-torah", "Late Fast Appendix · Torah Rendered into Greek", 10, 8, false, "Late appended commemorative fast", "The appendix associates Tevet 8 with the Torah's Greek translation under King Ptolemy and says darkness followed for three days.", "The notice reflects a later rabbinic memory of translation, not a judgment about contemporary translations or Greek-speaking Jews.", "Late appendix fast; not generally observed", "Late appended fast list · Ma'amar Aharon", "Printed Hebrew final discourse appended to Megillat Ta'anit", "Post-Talmudic textual layer; underlying traditions vary", "Megillat Ta'anit · Ma'amar Aharon", lateFastAppendixSourceURL, false},
	{"mtb-tevet-nine", "Late Fast Appendix · Unexplained Ninth of Tevet", 10, 9, false, "Late appended commemorative fast", "The appendix states that its rabbis did not record why Tevet 9 was a fast.", "No later conjecture is promoted to the status of the witness's missing explanation.", "Late appendix fast; not generally observed", "Late appended fast list · Ma'amar Aharon", "Printed Hebrew final discourse appended to Megillat Ta'anit", "Post-Talmudic textual layer; underlying traditions vary", "Megillat Ta'anit · Ma'amar Aharon", lateFastAppendixSourceURL, false},
	{"mtb-tevet-siege", "Late Fast Appendix · Siege of Jerusalem", 10, 10, false, "Late appended commemorative fast", "The appendix assigns Tevet 10 to the Babylonian king's siege against Jerusalem.", "This appendix layer overlaps the living Tenth of Tevet fast and remains separately labeled.", "Date remains an active fast; appendix is a separate textual layer", "Late appended fast list · Ma'amar Aharon", "Printed Hebrew final discourse appended to Megillat Ta'anit", "Post-Talmudic textual layer; underlying traditions vary", "Megillat Ta'anit · Ma'amar Aharon", lateFastAppendixSourceURL, false},
	{"mtb-shevat-righteous", "Late Fast Appendix · Righteous of Joshua's Generation", 11, 8, false, "Late appended commemorative fast", "The appendix assigns Shevat 8 to the deaths of righteous people from Joshua's generation.", "The people and episode are not further identified in the terse list.", "Late appendix fast; not generally observed", "Late appended fast list · Ma'amar Aharon", "Printed Hebrew final discourse appended to Megillat Ta'anit", "Post-Talmudic textual layer; underlying traditions vary", "Megillat Ta'anit · Ma'amar Aharon", lateFastAppendixSourceURL, false},
	{"mtb-shevat-benjamin", "Late Fast Appendix · Assembly against Benjamin", 11, 23, false, "Late appended commemorative fast", "The appendix associates Shevat 23 with Israel's assembly against Benjamin over the concubine at Gibeah and also mentions Micah's image.", "The notice combines biblical memories without establishing a single historical anniversary.", "Late appendix fast; not generally observed", "Late appended fast list · Ma'amar Aharon", "Printed Hebrew final discourse appended to Megillat Ta'anit", "Post-Talmudic textual layer; underlying traditions vary", "Megillat Ta'anit · Ma'amar Aharon", lateFastAppendixSourceURL, false},
	{"mtb-adar-moses", "Late Fast Appendix · Death of Moses", 12, 7, true, "Late appended commemorative fast", "The appendix assigns Adar 7 to the death of Moses.", "Adar placement in leap years varies in later practice; the app uses Adar II as a clearly labeled study projection.", "Rare or local memorial practices survive; no universal fast", "Late appended fast list · Ma'amar Aharon", "Printed Hebrew final discourse appended to Megillat Ta'anit", "Post-Talmudic textual layer; underlying traditions vary", "Megillat Ta'anit · Ma'amar Aharon", lateFastAppendixSourceURL, false},
	{"mtb-adar-schools", "Late Fast Appendix · Dispute of Beit Shammai and Beit Hillel", 12, 9, true, "Late appended commemorative fast", "The appendix says a fast was established on Adar 9 over the dispute between Beit Shammai and Beit Hillel.", "The terse list does not define which dispute or event it intends; Adar II is used as a study convention in leap years.", "Late appendix fast; not generally observed", "Late appended fast list · Ma'amar Aharon", "Printed Hebrew final discourse appended to Megillat Ta'anit", "Post-Talmudic textual layer; underlying traditions vary", "Megillat Ta'anit · Ma'amar Aharon", lateFastAppendixSourceURL, false},
}

var lateFastVariantNotes = map[string]string{
	"mtb-iyar-samuel":       "Halakhot Gedolot, Tur, and Shulchan Arukh place Samuel's memorial on Iyar 28 rather than this witness's Iyar 29.",
	"mtb-av-western-light":  "Halakhot Gedolot reads Av 17; Tur and Shulchan Arukh agree with this witness's Av 18.",
	"mtb-elul-spies":        "Tur and Shulchan Arukh read Elul 17 rather than this witness's Elul 7.",
	"mtb-tishrei-five":      "The printed Halakhot Gedolot numeral appears corrupt, while Tur and Shulchan Arukh agree with this witness's Tishrei 5.",
	"mtb-tishrei-atonement": "This Tishrei 10 item is explicit in the appendix, ambiguous or parenthetical in Halakhot Gedolot, and is not a separate supplementary date in Tur or Shulchan Arukh.",
	"mtb-cheshvan-zedekiah": "Tur and Shulchan Arukh read Cheshvan 7 rather than this witness's Cheshvan 6.",
	"mtb-kislev-jehoiakim":  "Halakhot Gedolot reads Kislev 8, while Tur and Shulchan Arukh read Kislev 28 rather than this witness's Kislev 7.",
	"mtb-tevet-nine":        "Halakhot Gedolot associates Tevet 9 with Ezra and Nehemiah; this appendix, Tur, and Shulchan Arukh leave the reason unrecorded.",
	"mtb-shevat-righteous":  "Other recensions and the later legal lists commonly read Shevat 5 rather than this witness's Shevat 8.",
	"mtb-shevat-benjamin":   "The printed Halakhot Gedolot has Shevat 3, likely reflecting loss of the Hebrew tens letter; Tur and Shulchan Arukh agree with this witness's Shevat 23.",
	"mtb-tammuz-seventeen":  "Halakhot Gedolot and Tur also mention the breached city; the base appendix wording encoded here does not.",
}

var talmudFixedHistoricalRules = []jewishCorpusRule{
	{
		CatalogID: "yoma-tevet-gerizim", Name: "Yoma 69a · Mount Gerizim Day", Month: 10, Day: 25,
		Category: "Fixed Talmudic historical festival", Summary: "Yoma 69a transmits a baraita naming Tevet 25 a no-eulogy festival associated with a polemical victory account about Mount Gerizim.",
		Note:   "This is a variant calendrical transmission of the Mount Gerizim story placed on 21 Kislev in the original Megillat Ta'anit, not confidently a second historical event. The narrative reflects ancient Jewish-Samaritan polemic and is not presented as a modern judgment or practice.",
		Status: "Discontinued", CalendarCorpus: "Bavli named calendar observances", Attestation: "Bavli Yoma 69a baraita attributed to a Megillat Ta'anit tradition", Era: "Talmudic reception of Second Temple memory",
		SourceName: "Bavli Yoma 69a", SourceURL: yomaGerizimSourceURL, StartsAtSunset: true,
	},
}

var rainFastThresholdRules = []jewishCorpusRule{
	{
		CatalogID: "rain-fast-individual-threshold", Name: "Rain-Fast Protocol · Individual Stage Threshold", Month: 8, Day: 17,
		Category: "Conditional Mishnaic fast protocol", Summary: "If Marheshvan 17 arrived without rain, designated individuals began a sequence of three daytime fasts.",
		Note:   "This is an annual threshold marker for a weather-dependent court protocol, not a claim that a fast occurs every Marheshvan 17.",
		Status: "Conditional Temple-era protocol; not automatically active", CalendarCorpus: "Mishnah Ta'anit drought protocol", Attestation: "Mishnah Ta'anit 1:4 with Bavli discussion", Era: "Mishnaic and Talmudic",
		SourceName: "Mishnah Ta'anit 1:4–7", SourceURL: rainFastSourceURL,
	},
	{
		CatalogID: "rain-fast-communal-threshold", Name: "Rain-Fast Protocol · First Communal Stage Threshold", Month: 9, Day: 1,
		Category: "Conditional Mishnaic fast protocol", Summary: "If Rosh Chodesh Kislev arrived without rain, the court decreed the first three communal daytime fasts.",
		Note:   "This is the threshold for a conditional court declaration. Later stages depended on the first fasts passing without an answer.",
		Status: "Conditional Temple-era protocol; not automatically active", CalendarCorpus: "Mishnah Ta'anit drought protocol", Attestation: "Mishnah Ta'anit 1:5 with Bavli discussion", Era: "Mishnaic and Talmudic",
		SourceName: "Mishnah Ta'anit 1:4–7", SourceURL: rainFastSourceURL,
	},
}

type jewishCatalogRecord struct {
	CatalogID      string
	Name           string
	Category       string
	NativeDate     string
	Summary        string
	Meaning        string
	Practices      []string
	Scripture      []string
	Status         string
	HistoricalNote string
	CalendarCorpus string
	Attestation    string
	Era            string
	SourceName     string
	SourceURL      string
}

var unprojectedJewishCorpusRecords = []jewishCatalogRecord{
	{
		CatalogID: "rain-fast-second-communal-stage", Name: "Rain-Fast Protocol · Second Communal Stage", Category: "Conditional Mishnaic fast protocol",
		NativeDate: "After the first three communal fasts pass without rain; a further Monday–Thursday–Monday sequence",
		Summary:    "The court could decree three additional, more severe communal fasts if the first communal stage passed without rain.",
		Meaning:    "The relative stage is preserved without inventing a fixed annual date or asserting that a court declared it in a modern year.",
		Practices:  []string{"Fast beginning the prior evening", "Abstention from work, bathing, anointing, shoes, and marital relations in the Mishnaic protocol"},
		Status:     "Conditional Temple-era protocol; no fixed annual projection", HistoricalNote: "Its dates depend on weather, prior unanswered fasts, weekdays, and a court declaration.",
		CalendarCorpus: "Mishnah Ta'anit drought protocol", Attestation: "Mishnah Ta'anit 1:6 with Bavli discussion", Era: "Mishnaic and Talmudic", SourceName: "Mishnah Ta'anit 1:4–7", SourceURL: rainFastSourceURL,
	},
	{
		CatalogID: "rain-fast-final-seven-stage", Name: "Rain-Fast Protocol · Final Seven Communal Fasts", Category: "Conditional Mishnaic fast protocol",
		NativeDate: "After the second communal stage passes without rain; seven alternating Monday/Thursday fasts, ending Monday",
		Summary:    "A final seven severe fasts brought the Mishnaic communal total to thirteen, apart from the first three individual fasts.",
		Meaning:    "The catalog preserves the escalation sequence while refusing to assign its contingent days to every modern winter.",
		Practices:  []string{"Alarm or shofar sounding", "Public prayer and additional blessings", "Shop closures with limited Monday/Thursday provisions"},
		Status:     "Conditional Temple-era protocol; no fixed annual projection", HistoricalNote: "The sequence is relative to two unanswered three-fast stages and therefore has no single Hebrew-date range.",
		CalendarCorpus: "Mishnah Ta'anit drought protocol", Attestation: "Mishnah Ta'anit 1:6 with Bavli discussion", Era: "Mishnaic and Talmudic", SourceName: "Mishnah Ta'anit 1:4–7", SourceURL: rainFastSourceURL,
	},
	{
		CatalogID: "rain-fast-post-thirteen-mourning", Name: "Rain-Fast Protocol · Mourning after Thirteen Fasts", Category: "Conditional Mishnaic drought response",
		NativeDate: "After all thirteen communal fasts pass without rain; communal restrictions and continuing individual fasts until Nisan ends",
		Summary:    "If the full communal sequence passed unanswered, the Mishnah prescribed public curtailment and continuing individual mourning rather than another fixed communal series.",
		Meaning:    "This closing stage completes the named Mishnaic protocol without converting a drought contingency into an annual holiday.",
		Practices:  []string{"Reduced business, building, betrothal, and greetings", "Continuing individual fasts in the source's seasonal framework"},
		Status:     "Conditional Temple-era protocol; no fixed annual projection", HistoricalNote: "The Mishnah says the individuals continue until Nisan ends; local climate and later law require expert interpretation.",
		CalendarCorpus: "Mishnah Ta'anit drought protocol", Attestation: "Mishnah Ta'anit 1:7 with Bavli discussion", Era: "Mishnaic and Talmudic", SourceName: "Mishnah Ta'anit 1:4–7", SourceURL: rainFastSourceURL,
	},
	{
		CatalogID: "maamadot-weekly-cycle", Name: "Ma'amadot · Weekly Temple Representation Cycle", Category: "Temple calendar institution",
		NativeDate: "Each priestly-watch week; creation readings Sunday–Friday and representative fasts Monday–Thursday",
		Summary:    "Israelite representatives accompanied the rotating priestly watches through creation readings, prayer, and a four-day Monday-through-Thursday fast pattern.",
		Meaning:    "A weekly Temple institution is cataloged as a pattern, not projected onto modern weekdays as though a historical watch roster were still operating.",
		Practices:  []string{"Monday–Thursday daytime fasts by members of the ma'amad", "Daily readings from the creation account", "Representation of the wider community alongside the Temple service"},
		Scripture:  []string{"Numbers 28:2", "Genesis 1:1–2:3"}, Status: "Temple institution discontinued; no modern roster projection",
		HistoricalNote: "Mishnah Ta'anit 4:2–3 and Bavli Ta'anit 26a–27b describe the institution; festival and service-day interactions prevent treating it as a simple universal four-day weekly holiday.",
		CalendarCorpus: "Mishnah and Bavli Temple calendar institutions", Attestation: "Mishnah Ta'anit 4:2–3; Bavli Ta'anit 26a–27b", Era: "Second Temple institution in Mishnaic/Talmudic record", SourceName: "Mishnah Ta'anit 4:2–3", SourceURL: maamadotSourceURL,
	},
	{
		CatalogID: "torah-jubilee-cycle", Name: "Yovel · Jubilee Cycle", Category: "Unprojected Torah cycle institution",
		NativeDate: "After seven sabbatical cycles; liberty proclaimed by shofar on 10 Tishrei of the Jubilee year",
		Summary:    "Leviticus appoints a Jubilee after seven cycles of seven years, with liberty proclaimed on Yom Kippur and hereditary land returning.",
		Meaning:    "The Torah institution remains visible without selecting one disputed post-Temple Jubilee count and fabricating modern occurrence years.",
		Practices:  []string{"Yom Kippur shofar proclamation historically", "Release and return provisions in the Torah's land system"},
		Scripture:  []string{"Leviticus 25:8–24"}, Status: "Not operative as a complete national institution; historical year count disputed",
		HistoricalNote: "Rabbinic sources dispute how the Jubilee year relates to the forty-nine-year sequence, and no universally accepted operative modern count supports Gregorian projection.",
		CalendarCorpus: "Torah and rabbinic sabbatical institutions", Attestation: "Torah appointment with later rabbinic counting disputes", Era: "Biblical institution; later legal interpretation", SourceName: "Torah · Leviticus 25", SourceURL: torahCyclesSourceURL,
	},
}

func sourceBoundJewishObservances(date time.Time, h hebrewDate) []Observance {
	var events []Observance
	events = append(events, materializeJewishCorpusRules(date, h, lateAppendedFastRules)...)
	events = append(events, materializeJewishCorpusRules(date, h, talmudFixedHistoricalRules)...)
	events = append(events, materializeJewishCorpusRules(date, h, rainFastThresholdRules)...)
	events = append(events, omerCountObservances(date, h)...)
	events = append(events, birkatLevanahObservances(date, h)...)
	events = append(events, birkatHachamaObservances(date)...)
	events = append(events, shemitahCycleObservances(date, h)...)
	return events
}

func materializeJewishCorpusRules(date time.Time, h hebrewDate, rules []jewishCorpusRule) []Observance {
	var events []Observance
	for _, rule := range rules {
		month := rule.Month
		if rule.Adar && hebrewLeap(h.Year) {
			month = 13
		}
		if h.Month != month || h.Day != rule.Day {
			continue
		}
		meaning := "This source record remains visible for historical study without presenting a discontinued or conditional observance as current practice."
		event := baseObservance(rule.Name, Judaism, []string{"Jewish textual-calendar study"}, rule.Category, rule.Summary, meaning, nil, nil, rule.SourceName, rule.SourceURL)
		event.CatalogID = rule.CatalogID
		event.Origin = rule.Attestation
		event.ObservanceStatus = rule.Status
		event.Historical = true
		event.HistoricalNote = rule.Note
		if strings.HasPrefix(rule.CatalogID, "mtb-") {
			event.HistoricalNote += " " + lateFastAppendixCaveat
			if variant := lateFastVariantNotes[rule.CatalogID]; variant != "" {
				event.HistoricalNote += " " + variant
			}
		}
		event.CalendarCorpus = rule.CalendarCorpus
		event.AttestationLayer = rule.Attestation
		event.Era = rule.Era
		if strings.HasPrefix(rule.CatalogID, "mtb-") {
			event.Era = "Late-antique antecedents and Geonic recension; unstable later printed witness"
		}
		event.NativeDateLabel = fmt.Sprintf("%d %s", rule.Day, hebrewMonthName(h.Year, month))
		event.ProjectionKind = "Received Hebrew month/day projected into the fixed calendar"
		event.ProjectionStatus = "Study projection"
		event.DateConfidence = "Witness-specific traditional date"
		event.DateCertainty = "Traditional Hebrew date projected into the modern fixed calendar"
		event.StartsAtSunset = rule.StartsAtSunset
		if rule.StartsAtSunset {
			event.DayBoundary = "Hebrew date begins at sunset"
		} else {
			event.DayBoundary = "Daytime fast or daytime threshold; not represented as a sunset-starting fast"
		}
		event.DateNote = "Historical study projection; not practical fasting guidance."
		event = singleOccurrence(event, date)
		event.ID = rule.CatalogID + "-" + date.Format("2006-01-02")
		events = append(events, event)
	}
	return events
}

func omerCountObservances(date time.Time, h hebrewDate) []Observance {
	start := jdToGregorian(hebrewToJD(h.Year, 1, 16))
	day := daysBetween(start, date) + 1
	if day < 1 || day > 49 {
		return nil
	}
	event := baseObservance(
		"Counting of the Omer",
		Judaism,
		[]string{"Jewish communities"},
		"Torah-appointed recurring count",
		"The Torah appoints a count of seven complete weeks, forty-nine days, from the Omer offering toward Shavuot.",
		"Each numbered evening links liberation to responsibility and harvest beginning to harvest culmination.",
		[]string{"Count the day and completed weeks after nightfall", "Recite the customary blessing when applicable"},
		[]string{"Leviticus 23:15–16", "Deuteronomy 16:9–10"},
		"Torah · Leviticus 23:15–16",
		"https://www.sefaria.org/Leviticus.23.15-16?lang=bi",
	)
	event.CatalogID = "torah-omer-count"
	event.Origin = "Torah count; rabbinic fixed-calendar placement"
	event.ObservanceStatus = "Still observed; Temple offering inactive"
	event.CalendarCorpus = "Torah and rabbinic calendrical institutions"
	event.NativeDateLabel = fmt.Sprintf("Omer day %d of 49", day)
	event.DateCertainty = "Nisan 16 through Sivan 5 in the received fixed Hebrew calendar"
	event.DateNote = "The count is recited after nightfall at the opening of each Hebrew date; the day card represents the daylight portion of that same numbered Omer day."
	event.DayBoundary = "Counted after nightfall"
	event.StartsAtSunset = true
	event = spanOccurrence(event, date, start, 49)
	event.ID = event.CatalogID + "-" + start.Format("2006-01-02")
	return []Observance{event}
}

func birkatLevanahObservances(date time.Time, h hebrewDate) []Observance {
	if h.Day < 1 || h.Day > 16 {
		return nil
	}
	start := date.AddDate(0, 0, -(h.Day - 1))
	event := baseObservance(
		"Birkat HaLevanah · Talmudic Monthly Timing Band",
		Judaism,
		[]string{"Jewish communities; timing customs vary"},
		"Recurring Talmudic lunar rite",
		"Sanhedrin 41b–42a preserves two latest-time opinions for blessing the renewing moon: seven days and sixteen days into the month.",
		"The monthly blessing turns visible lunar renewal into praise while the timing dispute keeps observation and calendrical judgment visible.",
		[]string{"Bless the renewing moon when visible", "Stand for the blessing in the Talmudic account"},
		[]string{"Exodus 12:2", "Exodus 15:2"},
		"Bavli Sanhedrin 41b–42a",
		birkatLevanahSourceURL,
	)
	event.CatalogID = "talmud-birkat-levanah-window"
	event.Origin = "Bavli Sanhedrin 41b–42a"
	event.ObservanceStatus = "Still observed; earliest time and final calculation vary by halakhic tradition"
	event.CalendarCorpus = "Bavli named calendar observances"
	event.AttestationLayer = "Bavli timing dispute: seven-day and sixteen-day latest limits"
	event.NativeDateLabel = fmt.Sprintf("Study band day %d; alternative Talmudic endpoints at days 7 and 16", h.Day)
	event.DateCertainty = "Hebrew-day visualization of Talmudic endpoint opinions, not a practical molad calculation"
	event.DateNote = "This 1–16 span visualizes the two latest-time opinions; it does not claim the blessing may be recited from day 1. Later authorities define earliest times and calculate the deadline from the molad, and visibility and weather still matter."
	event.HistoricalNote = "The Bavli supplies latest limits, not one universal modern practice window. Consult a local authority for practical timing."
	event.DayBoundary = "Nighttime lunar observation within a Hebrew date"
	event.StartsAtSunset = true
	event = spanOccurrence(event, date, start, 16)
	event.ID = event.CatalogID + "-" + start.Format("2006-01-02")
	return []Observance{event}
}

func birkatHachamaObservances(date time.Time) []Observance {
	if date.Month() != time.April || date.Day() != 8 || (date.Year()-2009)%28 != 0 {
		return nil
	}
	event := baseObservance(
		"Birkat HaChama · Blessing of the Sun",
		Judaism,
		[]string{"Jewish communities when the received cycle returns"},
		"Rare Talmudic solar-cycle rite",
		"Berakhot 59b records a blessing when the sun returns to the beginning of its twenty-eight-year cycle in the traditional calculation.",
		"The rare blessing frames cosmic regularity as creation remembered rather than as a modern astronomical measurement.",
		[]string{"Recite the blessing upon seeing the sun on the appointed morning", "Gather for psalms and teachings in many later communities"},
		[]string{"Genesis 1:14–19"},
		"Bavli Berakhot 59b",
		birkatHachamaSourceURL,
	)
	event.CatalogID = "talmud-birkat-hachama"
	event.Origin = "Bavli Berakhot 59b; received twenty-eight-year cycle"
	event.ObservanceStatus = "Still observed when the cycle returns"
	event.CalendarCorpus = "Bavli named calendar observances"
	event.AttestationLayer = "Abaye's twenty-eight-year cycle in Bavli Berakhot 59b"
	event.NativeDateLabel = "Traditional Nisan solar-cycle return; Wednesday morning"
	event.DateCertainty = "Received cycle anchored to 8 April 2009 within the app's 1900–2100 civil-date range"
	event.DateNote = "The projected civil dates follow the traditional twenty-eight-year reckoning, not the measured modern astronomical equinox. Practical recitation also depends on seeing the sun and later halakhic rules."
	event.DayBoundary = "Daytime blessing"
	event.StartsAtSunset = false
	event = singleOccurrence(event, date)
	event.ID = event.CatalogID + "-" + date.Format("2006-01-02")
	return []Observance{event}
}

func shemitahCycleObservances(date time.Time, h hebrewDate) []Observance {
	if h.Month != 7 || h.Day != 1 || h.Year%7 != 0 {
		return nil
	}
	event := baseObservance(
		"Shemitah · Sabbatical Year Begins",
		Judaism,
		[]string{"Jewish agricultural and halakhic communities, especially in the Land of Israel"},
		"Septennial Torah institution",
		"Every seventh year the Torah appoints a land sabbath and release; Tishrei 1 opens the received sabbatical year count.",
		"Shemitah joins ecological rest, release, ownership limits, and trust across an entire year rather than creating one festival day.",
		[]string{"Agricultural rest in the Land of Israel under applicable law", "Debt-release procedures and study according to halakhic practice"},
		[]string{"Exodus 23:10–11", "Leviticus 25:1–7", "Deuteronomy 15:1–11"},
		"Torah · Leviticus 25",
		torahCyclesSourceURL,
	)
	event.CatalogID = "torah-shemitah-cycle"
	event.Origin = "Torah institution; received rabbinic sabbatical count"
	event.ObservanceStatus = "Still legally significant; scope and practice vary"
	event.CalendarCorpus = "Torah and rabbinic sabbatical institutions"
	event.NativeDateLabel = fmt.Sprintf("Hebrew year %d, from Tishrei through Elul", h.Year)
	event.DateCertainty = "Received sabbatical sequence: Hebrew years divisible by seven"
	event.DateNote = "This card marks the year's opening rather than appearing on every day of the sabbatical year. Practical agricultural and financial law requires qualified guidance."
	event.DayBoundary = "Sabbatical year begins at Tishrei 1 sunset"
	event.StartsAtSunset = true
	event = singleOccurrence(event, date)
	event.ID = event.CatalogID + "-" + date.Format("2006-01-02")
	return []Observance{event}
}

func catalogOnlyJewishObservances() []Observance {
	events := make([]Observance, 0, len(unprojectedJewishCorpusRecords))
	for _, record := range unprojectedJewishCorpusRecords {
		event := baseObservance(record.Name, Judaism, []string{"Jewish textual-calendar study"}, record.Category, record.Summary, record.Meaning, record.Practices, record.Scripture, record.SourceName, record.SourceURL)
		event.ID = record.CatalogID
		event.CatalogID = record.CatalogID
		event.Origin = record.Attestation
		event.ObservanceStatus = record.Status
		event.Historical = true
		event.HistoricalNote = record.HistoricalNote
		event.CalendarCorpus = record.CalendarCorpus
		event.NativeDateLabel = record.NativeDate
		event.AttestationLayer = record.Attestation
		event.Era = record.Era
		event.CatalogOnly = true
		event.ProjectionKind = "No modern occurrence projected"
		event.ProjectionStatus = "Unprojected because the schedule is conditional, relative, institutionally inactive, or disputed"
		event.DateConfidence = "Native textual schedule only"
		event.DateCertainty = "No unique modern occurrence is asserted"
		event.DateNote = "Catalog record only; shown in the Observance Atlas without a fabricated Gregorian or fixed-Hebrew occurrence."
		events = append(events, event)
	}
	return events
}

func applyFourDestructionFastsLayer(event *Observance, fastName string, tishaBAv bool) {
	event.Origin = "Zechariah 8:19 and Bavli Rosh Hashanah 18b"
	event.SourceName = "Bavli Rosh Hashanah 18b"
	event.SourceURL = fourFastsSourceURL
	event.CalendarCorpus = "Bavli named calendar observances"
	event.AttestationLayer = "Rav Pappa's conditional status discussion of the four destruction fasts"
	event.ObservanceStatus = "Still observed; Bavli preserves conditional historical statuses"
	event.HistoricalNote = "Bavli Rosh Hashanah 18b reads Zechariah's four fasts as joy in a time of peace, fasting in persecution, and optional in an intermediate condition. This records the Talmudic status framework and is not presented as a modern practical exemption."
	if tishaBAv {
		event.HistoricalNote += " The passage distinguishes the Ninth of Av because multiple catastrophes were associated with it and the community accepted its fast."
	}
	event.DateNote = strings.TrimSpace(event.DateNote + " " + fastName + " is one of Zechariah 8:19's four destruction fasts; consult current halakhic guidance for practice.")
}
