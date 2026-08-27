# Lunarium

Lunarium is a Go + React calendar for exploring sacred time across Christianity, Judaism, and Islam. A day brings together civil, Hebrew, and tabular Hijri dates; lunar phase; observance details; Catholic readings; location-aware prayer windows; and clearly labeled Kabbalah, astrology, and numerology lenses.

## What is included

- Responsive Gregorian and Hebrew lunisolar month views with tradition filters, moon phases, numerology scores, and selected-day details.
- An annual observance atlas with community/denomination labels, meaning, common practices, scripture references, date caveats, and source links.
- Western and Orthodox Paschal cycles; principal Catholic, Orthodox, Anglican, Lutheran, Reformed, and Protestant annual observances.
- Shabbat, Rosh Chodesh, the Torah-appointed festival cycle and Temple rites, Mishnah/Talmud calendar rules, Purim Katan, traditional fasts, selected modern observances, and diaspora/Israel notes.
- The complete 35-entry/period core of Megillat Ta'anit and all nine Mishnaic family wood-offering dates, visibly separated and badged as historical or discontinued where appropriate.
- Jumu'ah, Ramadan, Hajj, both Eids, and major Sunni and Shia observances, with a visible crescent-sighting warning.
- A Catholic daily Mass reading plan with season, Sunday cycle A/B/C, weekday cycle I/II, and an official USCCB link for every date.
- Location-aware Christian prayer rhythm, Jewish proportional-hour windows, and Muslim salah times. Methods and local-authority caveats accompany every schedule.
- A day-level astronomical moon model plus separate, explicitly interpretive Kabbalah, tropical astrology, and Pythagorean numerology cards.

The catalog aims at principal shared, denominational, and weekly observances. “All Abrahamic holidays” has no finite universal list: local churches, dioceses, saints' calendars, Jewish minhagim, Islamic legal schools, Sufi orders, and regional calendars contain many additional dates. The app makes its represented scope visible and its data model is designed to accept more entries without UI changes.

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

The Go tests validate known Hebrew and Hijri conversions, Gregorian and Orthodox Easter, the August 2026 full moon, daily numerology, lectionary cycles and links, all three prayer schedules, core observances, Eid duration behavior, and annual three-tradition coverage. The second test step builds the React production bundle.

The Hebrew-month suite also exhaustively round-trips every supported civil date from 1900 through 2100, verifies common/leap-year Adar navigation, checks Hebrew month lengths, and asserts the complete Megillat Ta'anit and wood-offering manifests.

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

Returns a real Hebrew lunisolar month whose 29 or 30 day cells carry canonical Gregorian ISO dates. The response supplies server-calculated `previous` and `next` month references so clients navigate Tishrei-to-Elul order and Adar I/II correctly.

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
- Prayer times are planning aids. A rabbi or trusted luach should determine practical Jewish zmanim, and a local mosque should determine salah conventions and iqamah times.
- Catholic reading links point to the USCCB schedule for the United States. Optional memorials and diocesan calendars can differ.
- Moon illumination is a day-level mean-lunation estimate, not an ephemeris for navigation, ritual rulings, or observatory work.
- Kabbalistic month correspondences vary by lineage. Astrology and numerology are symbolic reflection tools, not scientific forecasts.

The About screen links the USCCB, Orthodox Church in America, Hebcal, U.S. Naval Observatory, and NOAA sources used for calculation design and validation.
# lunarium
