# Lunarium

Lunarium is a Go + React calendar for exploring sacred time, historical calendars, and the observable sky. A day brings together civil, Hebrew, and tabular Hijri dates; living and historical observances; Catholic readings; location-aware prayer windows; and clearly separated astronomical and symbolic lenses.

## What is included

- A responsive Hebrew lunisolar calendar by default, with a clearly labeled Gregorian alternate, tradition filters, moon phases, numerology scores, and selected-day details.
- An annual observance atlas with community/denomination labels, meaning, common practices, scripture references, date caveats, and source links.
- Western and Orthodox Paschal cycles; principal Catholic, Orthodox, Anglican, Lutheran, Reformed, and Protestant annual observances.
- Shabbat, Rosh Chodesh, the Torah-appointed festival cycle and Temple rites, all 49 Omer days, explicit Shemitah and Jubilee records, Purim Katan, traditional fasts, selected modern observances, and separate diaspora/Israel occurrences.
- Fixed, recurring, conditional, and institutional Mishnah/Talmud calendar records, including drought-fast stages, the ma'amadot cycle, Birkat HaLevanah, Birkat HaChama, and the four destruction fasts' Rosh Hashanah 18b status discussion.
- The complete 35-entry/period core of Megillat Ta'anit, all 26 dates in a separately labeled printed Ma'amar Aharon fast-appendix witness, and all nine Mishnaic family wood-offering dates, visibly separated and badged as historical or discontinued where appropriate.
- Jumu'ah, Ramadan, Hajj, both Eids, and major Sunni and Shia observances, with a visible crescent-sighting warning.
- A Catholic daily Mass reading plan with season, Sunday cycle A/B/C, weekday cycle I/II, and an official USCCB link for every date.
- Location-aware Christian prayer rhythm, Jewish proportional-hour windows, and Muslim salah times. Methods and local-authority caveats accompany every schedule.
- A day-level astronomical moon model plus separate, explicitly interpretive Kabbalah, tropical astrology, and Pythagorean numerology cards.
- A dedicated Sky alternate view powered offline by Astronomy Engine: exact lunar quarters, eclipses and contacts, seasons, nodes, apsides, planetary positions and ingresses, apparent retrograde stations, ecliptic aspects, true angular close approaches, and a selected-day almanac for twilight, rise, set, and culmination.
- A deterministic 12-sign horoscope lens that is visibly labeled symbolic entertainment and kept separate from computed astronomy.
- All 45 named rows in Fowler's reconstructed Roman *fasti antiquissimi* table, selected later Roman cycles, an app-defined Athens-anchored Attic lunar study proxy, and sourced Byzantine/Holy Roman local or confessional records.
- A documented living Neo-Pagan eightfold convention, kept distinct from historical Norse and Old English records and from the exact seasonal instants calculated in Sky.
- A closed source manifest for all 58 named or datable rows in UCL Digital Egypt's festival-date inventory, including five separate epagomenal birthdays and explicit period/site layers rather than one synthetic Egyptian calendar.
- Source-bounded Mesopotamian additions from CDLI's Ur III festival inventory—including its probable, unnamed month-IX row—and ORACC BTTo Q004806's scholarly Nisannu list, with uncertain identity, frequency, cult ownership, and tablet line numbers kept visible.
- Ugaritic KTU 1.46 and 1.109 calendar-bearing records plus five separately labeled ritual documents that are not promoted into proven annual holidays.
- A representative nine-record regional and Panhellenic Greek native-calendar catalog—Delphic, Olympian, Spartan, Argive, and Boeotian—alongside Babylonian, Ur III, late Uruk, Norse, Anglo-Saxon, and disputed Attic records that cannot be honestly pinned to one recurring Gregorian anniversary.

No finite universal list exists for “all” sacred or ancient holidays: local churches, saints' calendars, Jewish minhagim, Islamic legal schools, ancient cities, dynasties, regnal calendars, and modern communities all differ. Lunarium exposes corpus, site, era, attestation, living status, and projection confidence instead of flattening those differences.

Ancient completeness claims are source-bounded. “58 Egyptian rows” means the named and datable rows on the cited UCL Digital Egypt page, not every rite practiced across Egyptian history; the nine regional/Panhellenic Greek records are explicitly representative, not a complete calendar for the Greek-speaking world.

For the bounded “Torah and Talmud calendar” claim, Lunarium includes explicit recurring Torah times, counts, and sabbatical institutions; calendrical dates, conditional schedules, and Temple institutions in Mishnah Moed; and the named annual, monthly, weekly, or multi-year observances and timing disputes enumerated by its versioned manifest from the Babylonian and Jerusalem Talmuds. It excludes personal vows, undated narrative anniversaries, daily liturgy without a calendar trigger, and later local customs unless a separate catalog explicitly names them. Completeness is asserted against those explicit manifests—not every statement about time in rabbinic literature. Fixed dates are projected into the received Hebrew calendar, recurring rites are labeled as spans, weather-dependent stages remain conditional, and institutionally inactive or disputed schedules remain unprojected Atlas records.

The Ma'amar Aharon appendix is not silently treated as the original Megillat Ta'anit. The app follows all 26 Hebrew dates in the cited Warsaw 1874 printed witness and labels the corpus as a textually unstable late-antique-to-Geonic fast tradition. Other recensions, Halakhot Gedolot, Tur, and Shulchan Arukh preserve date and content variants, and the 26 records are not presented as universal modern fasting obligations.

## Run it

Requirements: Go 1.23+, Node 20.19+ (or Node 22.12+), and npm.

For a production-style local build:

```bash
make install
make build
./bin/lunarium
```

Then open `http://localhost:8080`.

## Docker and WorkNet hosting

Lunarium ships as a multi-stage production image. The final container runs as an unprivileged user with a read-only filesystem, no Linux capabilities, an internal health check, and no published host port.

The application joins the external `dev-net` bridge on the system Docker daemon used by `~/Networking`. WorkNet itself keeps using its separate `worknet_net` infrastructure network; its Published Apps flow connects the edge proxy to `dev-net` and owns the `*.worknet` DNS, TLS, and HTTPS route. The Compose file stays in this repository—hosting through `~/Networking` does not require copying it there.

The shared `dev-net` network must already exist on system Docker. Start the supported WorkNet runtime so its DNS, API, and edge proxy are available:

```bash
cd ~/Networking
./worknet-up.sh --skip-deps --no-stack
```

Then build and start Lunarium:

```bash
cd ~/Projects/tools/lunar-calendar
make docker-up
```

The Docker targets deliberately use `docker --context default`, matching `~/Networking`'s system-Docker runtime even when your interactive shell currently selects the separate rootless Docker context. Override `SYSTEM_DOCKER` only if WorkNet is deliberately running on another Docker endpoint.

Publish it through WorkNet's canonical **Published Apps** API:

```bash
make worknet-publish
```

This persists `lunarium.worknet` in WorkNet, creates its DNS record and local-CA certificate, attaches `worknet-edge-nginx` to `dev-net`, and routes HTTPS to `lunarium-app:8080`. Open:

```text
https://lunarium.worknet
```

The client must use WorkNet DNS and trust the WorkNet root CA. No manual edits to WorkNet's generated DNS or nginx files are required.

For frontend hot reload, use two terminals:

```bash
make dev-api
```

```bash
make dev-web
```

Vite runs at `http://localhost:5173` and proxies `/api` to the Go service.

## Verify

```bash
make test
```

The Go tests validate known Hebrew and Hijri conversions, Gregorian and Orthodox Easter, the August 2026 full moon, daily numerology, lectionary cycles and links, all three prayer schedules, historical manifests, bounded Hebrew and Attic navigation, native-date alternatives, global catalog identities, core observances, Eid duration behavior, and annual multi-corpus coverage. The frontend checks lunar/Gregorian boundary navigation, local and global eclipse circumstances, double-eclipse months, DST and midnight-transition boundaries, J2000 constellation lookup, known stations and approaches, then builds the React production bundle.

The Hebrew-month suite also exhaustively round-trips every supported civil date from 1900 through 2100, verifies common/leap-year Adar navigation, checks Hebrew month lengths, and asserts the complete original Megillat Ta'anit, 26-date printed fast-appendix witness, wood-offering, rain-fast protocol, and Torah/Talmud cycle manifests.

## API

### `GET /api/v1/calendar`

Query parameters:

- `year` — 1900–2100
- `month` — 1–12
- `latitude` — −66 to 66 (the supported solar-calculation range)
- `longitude` — −180 to 180
- `timezone` — an IANA timezone such as `America/New_York`
- `locationName` — display label

Example:

```text
/api/v1/calendar?year=2026&month=8&latitude=40.7128&longitude=-74.006&timezone=America%2FNew_York
```

### `GET /api/v1/calendar/hebrew`

Returns a real Hebrew lunisolar month whose 29 or 30 day cells carry canonical Gregorian ISO dates. The response supplies server-calculated `previous` and `next` month references so clients navigate Tishrei-to-Elul order and Adar I/II correctly. At the first and last complete Hebrew months inside the supported 1900–2100 civil range, the unavailable boundary reference is `null`.

Query parameters:

- `year` — Hebrew year whose resulting civil dates fall within 1900–2100
- `month` — Nisan = 1 through Elul = 6, Tishrei = 7 through Adar = 12, or Adar II = 13 in leap years
- the same location parameters accepted by the Gregorian endpoint

Example:

```text
/api/v1/calendar/hebrew?year=5786&month=6
```

This returns Elul 5786, spanning Gregorian 2026-08-14 through 2026-09-11.

### `GET /api/v1/observances?year=2026`

Returns the deduplicated annual atlas and counts by tradition.

### `GET /api/v1/about`

Returns calculation methodology, sources, and pastoral/data cautions rendered in the About dialog.

## Accuracy and respectful use

- Jewish and Islamic observances begin at sunset even though each card is anchored to its following civil daylight date.
- Hebrew dates use the arithmetic fixed lunisolar calendar. Historical dates from observational-calendar eras are annual projections of their received Hebrew month/day, not claims of exact proleptic Gregorian anniversaries. Hijri dates are tabular estimates and may differ by a day from local crescent sighting or an official national calendar.
- Conditional Jewish drought-fast thresholds do not assert that a fast occurs every year. Relative drought stages, the ma'amadot roster, and the disputed Jubilee count remain unprojected Atlas records; Birkat HaLevanah and Birkat HaChama bands are study visualizations rather than practical halakhic calculations.
- Prayer times are planning aids. A rabbi or trusted luach should determine practical Jewish zmanim, and a local mosque should determine salah conventions and iqamah times.
- Catholic reading links point to the USCCB schedule for the United States. Optional memorials and diocesan calendars can differ.
- The calendar-card Moon is a day-level mean-lunation estimate. The alternate Sky view uses a tested ephemeris engine for civil years 1900–2100, but remains unsuitable for navigation, ritual rulings, eclipse safety decisions, or spacecraft work. Rise/set times assume standard refraction; weather and the real horizon can shift observed times.
- Ancient Roman dates are nominal projections. Attic dates are a labeled app-defined lunar study proxy, bounded to 1900–2100; disputed alternatives remain native-date-only. Egyptian, Mesopotamian, Ugaritic, and regional Greek native dates are left unprojected where an epoch, city, ruler, intercalation decision, or observational anchor is missing.
- There is no securely datable “Joseph era” or identified pharaoh from which to derive one Egyptian holiday calendar; Egyptian records are presented by their actually attested sites and periods.
- BTTo Q004806 is a scholarly/theological list whose column-line references are not festival days; CDLI's unnamed month-IX festival, monthly Annunītum frequency, and Tummal festival identification retain their published probability and inferential limits. Ugaritic ritual documents without calendar rubrics are discoverable as documents, not asserted annual holidays.
- Kabbalistic month correspondences vary by lineage. Astrology, horoscopes, and numerology are symbolic reflection tools, not scientific forecasts or advice.

The About screen links the calculation engines, primary corpora, scholarly references, and living-community sources used for design and validation.
