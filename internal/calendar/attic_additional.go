package calendar

// These entries complete the major Athens-centered festival families named in
// Parker's checklist while preserving the places where the evidence gives only
// a month, a range of scholarly proposals, or a deme-specific schedule. The
// placeholder StartDay on catalog-only rules is never projected or displayed;
// the annotation carries the honest native-date statement.

const (
	atticRuralDionysiaURL = "https://www.ascsa.edu.gr/uploads/media/hesperia/40981054.pdf"
)

func init() {
	atticFestivalRules = append(atticFestivalRules,
		atticFestivalRule{"attic-metageitnia", "Metageitnia", "Metageitnion", 7, 7, "Festival of Apollo", "The festival that gave Metageitnion its name honored Apollo Metageitnios, but the surviving evidence does not secure one universal numbered day.", "Apollo's epithet and the month name evoke movement between neighboring communities and civic relationships.", []string{"Sacrifice to Apollo Metageitnios"}, "low"},
		atticFestivalRule{"attic-apatouria", "Apatouria", "Pyanepsion", 19, 21, "Phratry festival", "A multi-day festival of the Ionian phratries took place in Pyanepsion, with the detailed schedule managed through kinship and phratry institutions.", "The festival renewed civic kinship by admitting and registering children and honoring the gods of the phratries.", []string{"Phratry meals and sacrifices", "Registration of children", "Hair offerings in historical accounts"}, "medium"},
		atticFestivalRule{"attic-pompaia-maimakteria", "Pompaia / Maimakteria", "Maimakterion", 10, 10, "Festival of Zeus", "Sparse testimony places purificatory rites for Zeus in Maimakterion but does not establish one secure numbered day or prove that Pompaia and Maimakteria were identical in every period.", "The surviving rites sought protection and favorable winter weather from Zeus in his storm and mild-weather aspects.", []string{"Procession with a fleece in later testimony", "Purificatory and protective rites"}, "low"},
		atticFestivalRule{"attic-rural-dionysia", "Rural Dionysia · Dionysia at the Demes", "Poseideon", 1, 1, "Deme Dionysian festivals", "Attic demes held their own Dionysia during Poseideon on locally selected dates rather than one Athenian citywide day.", "Local processions, sacrifices, and performances honored Dionysos while expressing each deme's own civic life.", []string{"Deme procession and sacrifice", "Phallic songs in literary testimony", "Local dramatic performances in participating demes"}, "high"},
		atticFestivalRule{"attic-lesser-mysteries", "Lesser Eleusinian Mysteries", "Anthesterion", 20, 26, "Mystery festival", "The Lesser Mysteries were held at Agrai as preparation for the Greater Mysteries, but their exact Anthesterion schedule remains disputed.", "The rites formed a preparatory stage within the Eleusinian initiatory system; restricted content is not reconstructed here.", []string{"Purification and preliminary initiation at Agrai", "Restricted rites not reproduced"}, "low"},
		atticFestivalRule{"attic-pandia", "Pandia", "Elaphebolion", 17, 17, "Festival of Zeus", "The Pandia followed the City Dionysia and is conventionally placed on Elaphebolion 17.", "A pan-Attic gathering probably honored Zeus; ancient evidence does not securely establish later lunar interpretations of the name.", []string{"Assembly and sacrifice to Zeus in the received reconstruction"}, "medium"},
		atticFestivalRule{"attic-kallynteria", "Kallynteria", "Thargelion", 24, 24, "Festival of Athena", "The Kallynteria involved cleansing Athena's sanctuary before the Plynteria, but proposals for its numbered day conflict.", "Ritual cleaning prepared the sanctuary and cult image for the following garment-washing rites.", []string{"Cleansing Athena's sanctuary"}, "low"},
		atticFestivalRule{"attic-arrephoria", "Arrephoria", "Skirophorion", 3, 3, "Festival service of Athena", "The Arrephoroi completed a nocturnal service for Athena near the end of the Attic year; the exact numbered day is not secure enough for projection.", "The rite marked the culmination of the young attendants' sacred service and connected the Acropolis with a hidden nocturnal route.", []string{"Nocturnal carrying rite by the Arrephoroi", "Restricted objects not reconstructed"}, "low"},
	)

	for catalogID, annotation := range map[string]atticFestivalAnnotation{
		"attic-metageitnia": {
			NativeDateLabel: "Metageitnion; exact numbered day not securely attested",
			DateNote:        "Month-level placement is retained without selecting a conjectural modern day.",
			CatalogOnly:     true,
		},
		"attic-apatouria": {
			NativeDateLabel: "Pyanepsion; three-day phratry sequence, exact civic dates and local schedules varied",
			DateNote:        "The rule's commonly reconstructed three-day character is preserved, but it is not flattened into one universal date range.",
			CatalogOnly:     true,
		},
		"attic-pompaia-maimakteria": {
			NativeDateLabel: "Maimakterion; exact day and relationship between Pompaia and Maimakteria disputed",
			DateNote:        "The related names remain together as a source problem rather than two fabricated fixed anniversaries.",
			CatalogOnly:     true,
		},
		"attic-rural-dionysia": {
			NativeDateLabel: "Poseideon; date selected independently by each participating Attic deme",
			DateNote:        "Deme calendars deliberately used different dates, so no citywide projection is generated.",
			SourceName:      "American School of Classical Studies at Athens · Deme Theaters in Attica",
			SourceURL:       atticRuralDionysiaURL,
			CatalogOnly:     true,
		},
		"attic-lesser-mysteries": {
			NativeDateLabel: "Anthesterion, commonly reconstructed within days 20–26; exact sequence disputed",
			DateNote:        "The month and preparatory role are firmer than any single proposed range.",
			CatalogOnly:     true,
		},
		"attic-pandia": {
			NativeDateLabel: "Elaphebolion 17 (conventional placement after the City Dionysia)",
			DateNote:        "The day-17 projection is medium confidence; the festival's divine interpretation remains debated.",
		},
		"attic-kallynteria": {
			NativeDateLabel: "Late Thargelion before Plynteria; proposed dates include 19, 22, 24, and 28",
			DateNote:        "Conflicting reconstructions are retained rather than presented as one attested date.",
			CatalogOnly:     true,
		},
		"attic-arrephoria": {
			NativeDateLabel: "Skirophorion; exact numbered day uncertain",
			DateNote:        "The nocturnal rite is well known, but the calendar day is not projected from later convention alone.",
			CatalogOnly:     true,
		},
	} {
		atticFestivalAnnotations[catalogID] = annotation
	}
}
