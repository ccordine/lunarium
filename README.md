# Lunarium

Lunarium is a Go + React calendar for exploring sacred time across Christianity, Judaism, and Islam. A day brings together civil, Hebrew, and tabular Hijri dates; lunar phase; observance details; Catholic readings; location-aware prayer windows; and clearly labeled Kabbalah, astrology, and numerology lenses.

## What is included

- A responsive month calendar with tradition filters, Hebrew dates, moon phases, numerology scores, and selected-day details.
- An annual observance atlas with community/denomination labels, meaning, common practices, scripture references, date caveats, and source links.
- Western and Orthodox Paschal cycles; principal Catholic, Orthodox, Anglican, Lutheran, Reformed, and Protestant annual observances.
- Shabbat, Rosh Chodesh, High Holy Days, Torah festivals, fasts, selected modern observances, and diaspora/Israel notes.
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

## Docker and Worknet hosting

Lunarium ships as a multi-stage production image. The final container runs as an unprivileged user with a read-only filesystem, no Linux capabilities, an internal health check, and no published host port.

Start the supported Worknet runtime first so its external Docker network exists:

```bash
cd ~/Networking
./worknet-up.sh --skip-deps --no-stack
```

Then build and start Lunarium:

```bash
cd ~/Projects/tools/lunar-calendar
make docker-up
```

Publish it through Worknet's canonical **Published Apps** API:

```bash
make worknet-publish
```

This persists `lunarium.worknet` in Worknet, creates its DNS record and local-CA certificate, attaches nginx-edge to the application network, and routes HTTPS to `lunarium-app:8080`. Open:

```text
https://lunarium.worknet
```

The client must use Worknet DNS and trust the Worknet root CA. No manual edits to Worknet's generated DNS or nginx files are required.

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

### `GET /api/v1/observances?year=2026`

Returns the deduplicated annual atlas and counts by tradition.

### `GET /api/v1/about`

Returns calculation methodology, sources, and pastoral/data cautions rendered in the About dialog.

## Accuracy and respectful use

- Jewish and Islamic observances begin at sunset even though each card is anchored to its following civil daylight date.
- Hebrew dates are arithmetic calendar calculations. Hijri dates are tabular estimates and may differ by a day from local crescent sighting or an official national calendar.
- Prayer times are planning aids. A rabbi or trusted luach should determine practical Jewish zmanim, and a local mosque should determine salah conventions and iqamah times.
- Catholic reading links point to the USCCB schedule for the United States. Optional memorials and diocesan calendars can differ.
- Moon illumination is a day-level mean-lunation estimate, not an ephemeris for navigation, ritual rulings, or observatory work.
- Kabbalistic month correspondences vary by lineage. Astrology and numerology are symbolic reflection tools, not scientific forecasts.

The About screen links the USCCB, Orthodox Church in America, Hebcal, U.S. Naval Observatory, and NOAA sources used for calculation design and validation.
# lunarium
