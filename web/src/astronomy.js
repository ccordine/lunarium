import {
  AngleBetween,
  ApsisKind,
  Body,
  Constellation,
  Ecliptic,
  EclipticGeoMoon,
  Equator,
  GeoVector,
  Horizon,
  Illumination,
  MoonPhase,
	NextGlobalSolarEclipse,
  NextLunarApsis,
	NextLunarEclipse,
  NextMoonNode,
  NextMoonQuarter,
  NodeEventKind,
  Observer,
  SearchAltitude,
  SearchGlobalSolarEclipse,
  SearchHourAngle,
  SearchLocalSolarEclipse,
  SearchLunarApsis,
  SearchLunarEclipse,
  SearchMaxElongation,
  SearchMoonNode,
  SearchMoonQuarter,
  SearchPeakMagnitude,
  SearchPlanetApsis,
  SearchRiseSet,
  SearchTransit,
  Seasons,
  SunPosition,
} from 'astronomy-engine'

const ENGINE_SOURCE = 'https://github.com/cosinekitty/astronomy'
const AU_KM = 149_597_870.7
const HOUR = 60 * 60 * 1000
const DAY = 24 * HOUR
const SKY_MONTH_CACHE = new Map()

export const SKY_YEAR_MIN = 1900
export const SKY_YEAR_MAX = 2100

export const ZODIAC = [
  { name: 'Aries', symbol: '♈', element: 'Fire' },
  { name: 'Taurus', symbol: '♉', element: 'Earth' },
  { name: 'Gemini', symbol: '♊', element: 'Air' },
  { name: 'Cancer', symbol: '♋', element: 'Water' },
  { name: 'Leo', symbol: '♌', element: 'Fire' },
  { name: 'Virgo', symbol: '♍', element: 'Earth' },
  { name: 'Libra', symbol: '♎', element: 'Air' },
  { name: 'Scorpio', symbol: '♏', element: 'Water' },
  { name: 'Sagittarius', symbol: '♐', element: 'Fire' },
  { name: 'Capricorn', symbol: '♑', element: 'Earth' },
  { name: 'Aquarius', symbol: '♒', element: 'Air' },
  { name: 'Pisces', symbol: '♓', element: 'Water' },
]

export const SKY_BODIES = [
  { body: Body.Sun, name: 'Sun', symbol: '☉', color: '#d7a94d' },
  { body: Body.Moon, name: 'Moon', symbol: '☽', color: '#d8dbe7' },
  { body: Body.Mercury, name: 'Mercury', symbol: '☿', color: '#baa58e' },
  { body: Body.Venus, name: 'Venus', symbol: '♀', color: '#dfad87' },
  { body: Body.Mars, name: 'Mars', symbol: '♂', color: '#c66c5c' },
  { body: Body.Jupiter, name: 'Jupiter', symbol: '♃', color: '#d4a371' },
  { body: Body.Saturn, name: 'Saturn', symbol: '♄', color: '#c9bd86' },
  { body: Body.Uranus, name: 'Uranus', symbol: '♅', color: '#7bbdcc' },
  { body: Body.Neptune, name: 'Neptune', symbol: '♆', color: '#657fca' },
  { body: Body.Pluto, name: 'Pluto', symbol: '♇', color: '#9d82a9' },
]

const PLANETS = SKY_BODIES.slice(2)
const ASPECTS = [
  { angle: 0, name: 'conjunction', symbol: '☌' },
  { angle: 60, name: 'sextile', symbol: '⚹' },
  { angle: 90, name: 'square', symbol: '□' },
  { angle: 120, name: 'trine', symbol: '△' },
  { angle: 180, name: 'opposition', symbol: '☍' },
]
const QUARTERS = [
  { name: 'New Moon', symbol: '●', summary: 'The Moon and Sun share the same apparent ecliptic longitude.' },
  { name: 'First Quarter Moon', symbol: '◐', summary: 'The Moon is 90° east of the Sun in apparent ecliptic longitude.' },
  { name: 'Full Moon', symbol: '○', summary: 'The Moon is opposite the Sun in apparent ecliptic longitude.' },
  { name: 'Last Quarter Moon', symbol: '◑', summary: 'The Moon is 90° west of the Sun in apparent ecliptic longitude.' },
]
const SEASON_EVENTS = [
  ['mar_equinox', 'March equinox', '♈'],
  ['jun_solstice', 'June solstice', '☀'],
  ['sep_equinox', 'September equinox', '♎'],
  ['dec_solstice', 'December solstice', '❄'],
]

/**
 * Calculates a bounded set of notable sky events for one Gregorian month.
 * All calculations run locally through Astronomy Engine; no network call is made.
 */
export function buildSkyMonth(year, month, location) {
	assertSupportedYear(year)
	if (!Number.isInteger(month) || month < 1 || month > 12) throw new RangeError('Sky month must be an integer from 1 through 12.')
	const timeZone = validTimeZone(location.timezone)
	const effectiveLocation = { ...location, timezone: timeZone }
	const cacheKey = [year, month, effectiveLocation.latitude, effectiveLocation.longitude, effectiveLocation.elevation || 0, timeZone].join('|')
	if (SKY_MONTH_CACHE.has(cacheKey)) return SKY_MONTH_CACHE.get(cacheKey)
  const monthKey = `${year}-${String(month).padStart(2, '0')}`
  const nextMonthDate = new Date(Date.UTC(year, month, 1))
  const nextMonthKey = `${nextMonthDate.getUTCFullYear()}-${String(nextMonthDate.getUTCMonth() + 1).padStart(2, '0')}-01`
	const start = zonedDateTimeToUTC(`${monthKey}-01`, 0, timeZone)
	const end = zonedBoundaryDateTimeToUTC(nextMonthKey, 0, timeZone)
  const searchStart = new Date(start.getTime() - HOUR)
	const observer = makeObserver(effectiveLocation)
  const context = createContext(start, end)
  const events = []
  const warnings = []

  collectSafely(events, warnings, 'lunar phases', () => collectMoonQuarters(searchStart, end))
  collectSafely(events, warnings, 'seasons', () => collectSeasons(year, start, end))
  collectSafely(events, warnings, 'lunar apsides', () => collectLunarApsides(searchStart, end))
  collectSafely(events, warnings, 'lunar nodes', () => collectLunarNodes(searchStart, end))
  collectSafely(events, warnings, 'eclipses', () => collectEclipses(searchStart, start, end, observer))
  collectSafely(events, warnings, 'inner-planet visibility', () => collectInnerPlanetEvents(searchStart, start, end, observer))
  collectSafely(events, warnings, 'planetary apsides', () => collectPlanetApsides(searchStart, start, end))
  collectSafely(events, warnings, 'retrograde stations', () => collectStations(context))
  collectSafely(events, warnings, 'zodiac ingresses', () => collectIngresses(context))
  collectSafely(events, warnings, 'ecliptic aspects', () => collectAspects(context))
  collectSafely(events, warnings, 'close approaches', () => collectApproaches(context))
  collectSafely(events, warnings, 'planet gatherings', () => collectPlanetGathering(context))

  const unique = new Map()
  events
	.filter((item) => item.time >= start && item.time < end && localDateKey(item.time, timeZone).startsWith(monthKey))
    .forEach((item) => {
      const key = `${item.category}|${item.name}|${Math.round(item.time.getTime() / 60000)}`
      if (!unique.has(key)) unique.set(key, item)
    })

  const sorted = [...unique.values()]
    .sort((a, b) => a.time - b.time || b.priority - a.priority)
	.map((item) => finalizeEvent(item, timeZone))
  const eventsByDate = Object.groupBy
    ? Object.groupBy(sorted, (item) => item.date)
    : sorted.reduce((groups, item) => ({ ...groups, [item.date]: [...(groups[item.date] || []), item] }), {})

	const result = {
    year,
    month,
    label: new Intl.DateTimeFormat('en-US', { month: 'long', year: 'numeric', timeZone: 'UTC' }).format(new Date(Date.UTC(year, month - 1, 1))),
    events: sorted,
    eventsByDate,
    eventDays: Object.keys(eventsByDate).sort(),
    sourceName: 'Astronomy Engine',
    sourceUrl: ENGINE_SOURCE,
    methodology: `Ephemeris events are calculated locally for the supported ${SKY_YEAR_MIN}–${SKY_YEAR_MAX} range. Close approaches use true three-dimensional angular separation; aspects and ingresses use apparent geocentric ecliptic longitude.`,
    warnings,
  }
	SKY_MONTH_CACHE.set(cacheKey, result)
	if (SKY_MONTH_CACHE.size > 24) SKY_MONTH_CACHE.delete(SKY_MONTH_CACHE.keys().next().value)
	return result
}

/** Calculates a location-aware evening snapshot for a civil date. */
export function buildSkySnapshot(dateString, location) {
	assertSupportedDate(dateString)
	const timeZone = validTimeZone(location.timezone)
	const time = zonedDateTimeToUTC(dateString, 21, timeZone)
	const observer = makeObserver({ ...location, timezone: timeZone })
  const almanac = buildDailyAlmanac(dateString, observer, timeZone)
  const positions = SKY_BODIES.map((meta) => positionFor(meta, time, observer))
  const moonAngle = MoonPhase(time)
  const moonLight = safeIllumination(Body.Moon, time)
  const aspects = currentAspects(time)
  const retrograde = positions.filter((item) => item.retrograde).map((item) => item.name)
	const solarCenter = horizontalPosition(Body.Sun, time, observer, null)

  return {
    date: dateString,
    referenceTime: time.toISOString(),
	  referenceLabel: `21:00 ${timeZone}`,
    positions,
    aspects,
    retrograde,
    almanac,
    warnings: almanac.warnings,
    moon: {
      phase: phaseName(moonAngle),
      angle: round(moonAngle, 1),
      illumination: moonLight ? round(moonLight.phase_fraction * 100, 1) : null,
      distanceKm: moonLight ? Math.round(moonLight.geo_dist * AU_KM) : null,
    },
	  skyState: solarCenter.altitude < -18 ? 'Astronomical night' : solarCenter.altitude < -12 ? 'Astronomical twilight' : solarCenter.altitude < -6 ? 'Nautical twilight' : solarCenter.altitude < 0 ? 'Civil twilight' : 'Daylight',
	  caveat: 'Right ascension, declination, altitude, and azimuth are topocentric at 21:00 local civil time. Ecliptic longitude and tropical sign are apparent geocentric values; IAU constellation lookup uses J2000 coordinates.',
  }
}

/**
 * Calculates observational circumstances that belong to one selected civil day.
 * These repeat daily, so they stay in the inspector instead of flooding the
 * monthly event timeline.
 */
function buildDailyAlmanac(dateString, observer, timeZone) {
  const start = zonedDateTimeToUTC(dateString, 0, timeZone)
  const end = zonedBoundaryDateTimeToUTC(addCivilDays(dateString, 1), 0, timeZone)
  const limitDays = Math.max(1.1, (end.getTime() - start.getTime()) / DAY + 0.1)
  const warnings = []
  const withinDay = (value) => {
    const date = value?.date || value?.time?.date
    return date instanceof Date && date >= start && date < end ? date.toISOString() : null
  }
  const safely = (family, calculation) => {
    try {
      return calculation()
    } catch (error) {
      pushWarning(warnings, family, error)
      return null
    }
  }
  const altitudeCrossing = (key, label, direction, altitude) => ({
    key,
    label,
    time: withinDay(safely(`daily ${label.toLowerCase()}`, () => SearchAltitude(Body.Sun, observer, direction, start, limitDays, altitude))),
  })
  const riseSet = (body, direction, family) => withinDay(safely(family, () => SearchRiseSet(body, observer, direction, start, limitDays)))
  const culmination = (body, family) => {
	const event = safely(family, () => SearchHourAngle(body, observer, 0, start, limitDays))
    if (!event || !withinDay(event)) return { time: null, altitude: null }
    return { time: event.time.date.toISOString(), altitude: round(event.hor.altitude, 1) }
  }

  const solarCulmination = culmination(Body.Sun, 'daily solar culmination')
  const solar = [
    altitudeCrossing('astronomical-dawn', 'Astronomical dawn', +1, -18),
    altitudeCrossing('nautical-dawn', 'Nautical dawn', +1, -12),
    altitudeCrossing('civil-dawn', 'Civil dawn', +1, -6),
    { key: 'sunrise', label: 'Sunrise', time: riseSet(Body.Sun, +1, 'daily sunrise') },
    { key: 'solar-culmination', label: 'Solar culmination', ...solarCulmination },
    { key: 'sunset', label: 'Sunset', time: riseSet(Body.Sun, -1, 'daily sunset') },
    altitudeCrossing('civil-dusk', 'Civil dusk', -1, -6),
    altitudeCrossing('nautical-dusk', 'Nautical dusk', -1, -12),
    altitudeCrossing('astronomical-dusk', 'Astronomical dusk', -1, -18),
  ]

  const bodyRows = [SKY_BODIES[1], ...PLANETS].map((meta) => {
    const upper = culmination(meta.body, `daily ${meta.name.toLowerCase()} culmination`)
    return {
      name: meta.name,
      symbol: meta.symbol,
      rise: riseSet(meta.body, +1, `daily ${meta.name.toLowerCase()} rise`),
      culmination: upper.time,
      culminationAltitude: upper.altitude,
      set: riseSet(meta.body, -1, `daily ${meta.name.toLowerCase()} set`),
    }
  })

  return {
    date: dateString,
    timeZone,
    solar,
    bodies: bodyRows,
    warnings,
    caveat: 'Rise and set include standard refraction and apparent disc radius where applicable. Twilight uses the Sun’s unrefracted center at −6°, −12°, and −18°. A blank time means that crossing does not occur during this local civil date.',
  }
}

/** A deterministic symbolic prompt, deliberately separated from computed astronomy. */
export function buildHoroscope(snapshot, signName) {
  const sign = ZODIAC.find((item) => item.name === signName) || ZODIAC[0]
  const moon = snapshot.positions.find((item) => item.name === 'Moon')
  const sun = snapshot.positions.find((item) => item.name === 'Sun')
  const relation = zodiacRelationship(sign.name, moon.sign)
  const retrogradeText = snapshot.retrograde.length
    ? `${snapshot.retrograde.slice(0, 3).join(', ')} ${snapshot.retrograde.length === 1 ? 'is' : 'are'} retrograde in the geocentric model.`
    : 'No tracked planet is retrograde in the geocentric model.'
  const closest = snapshot.aspects[0]
  const aspectText = closest
    ? `${closest.a} ${closest.symbol} ${closest.b} is the tightest major aspect in this evening snapshot.`
    : 'No major aspect is within the three-degree display orb this evening.'
  const prompts = {
    same: 'Name what is already asking for your attention, then give it one concrete form.',
    supportive: 'Use familiar strengths generously, while leaving room for a different method.',
    dynamic: 'Treat friction as information: pause, choose the real priority, and act cleanly.',
    neutral: 'Notice what changes when you replace prediction with a careful question.',
  }

  return {
    sign,
    title: `${sign.symbol} ${sign.name} reflection`,
    theme: `${moon.sign} Moon · ${sun.sign} Sun · ${relation.label}`,
    prompt: prompts[relation.kind],
    rationale: `${aspectText} ${retrogradeText}`,
    caveat: 'This horoscope is a symbolic, entertainment-oriented reflection generated from the computed sky. Astrology is not a scientific method and this is not advice or a factual prediction.',
  }
}

function createContext(start, end) {
  const longitudeCache = new Map()
  const vectorCache = new Map()
  return {
    start,
    end,
    longitude(body, time) {
      const key = `${body}|${time.getTime()}`
      if (!longitudeCache.has(key)) longitudeCache.set(key, apparentLongitude(body, time))
      return longitudeCache.get(key)
    },
    vector(body, time) {
      const key = `${body}|${time.getTime()}`
      if (!vectorCache.has(key)) vectorCache.set(key, GeoVector(body, time, true))
      return vectorCache.get(key)
    },
  }
}

function collectMoonQuarters(start, end) {
  const found = []
  let quarter = SearchMoonQuarter(start)
  while (quarter.time.date < end) {
    const meta = QUARTERS[quarter.quarter]
    found.push(skyEvent({
      slug: `moon-quarter-${quarter.quarter}`,
      name: meta.name,
      symbol: meta.symbol,
      category: 'lunar phase',
      time: quarter.time.date,
      summary: meta.summary,
      details: 'An exact geocentric quarter-phase instant, which may occur on a different civil date in another timezone.',
      bodies: ['Moon', 'Sun'],
      priority: 10,
    }))
    quarter = NextMoonQuarter(quarter)
  }
  return found
}

function collectSeasons(year, start, end) {
  const seasons = Seasons(year)
  return SEASON_EVENTS.flatMap(([key, name, symbol]) => {
    const time = seasons[key].date
    if (time < start || time >= end) return []
    return [skyEvent({
      slug: key,
      name,
      symbol,
      category: 'season',
      time,
      summary: name.includes('equinox')
        ? 'The apparent Sun crosses the celestial equator.'
        : 'The apparent Sun reaches its northernmost or southernmost declination.',
      details: 'This is the exact astronomical season boundary, independent of weather or civil convention.',
      bodies: ['Sun', 'Earth'],
      priority: 10,
    })]
  })
}

function collectLunarApsides(start, end) {
  const found = []
  let apsis = SearchLunarApsis(start)
  while (apsis.time.date < end) {
    const perigee = apsis.kind === ApsisKind.Pericenter
    found.push(skyEvent({
      slug: perigee ? 'lunar-perigee' : 'lunar-apogee',
      name: perigee ? 'Lunar perigee' : 'Lunar apogee',
      symbol: perigee ? '↘' : '↗',
      category: 'distance',
      time: apsis.time.date,
      summary: `The Moon reaches its ${perigee ? 'minimum' : 'maximum'} orbital distance for this cycle: ${Math.round(apsis.dist_km).toLocaleString('en-US')} km center-to-center.`,
      details: 'Perigee and apogee are orbital distance extrema, not by themselves a “supermoon” or “micromoon.”',
      bodies: ['Moon', 'Earth'],
      value: `${Math.round(apsis.dist_km).toLocaleString('en-US')} km`,
      priority: 7,
    }))
    apsis = NextLunarApsis(apsis)
  }
  return found
}

function collectLunarNodes(start, end) {
  const found = []
  let node = SearchMoonNode(start)
  while (node.time.date < end) {
    const ascending = node.kind === NodeEventKind.Ascending
    found.push(skyEvent({
      slug: ascending ? 'ascending-node' : 'descending-node',
      name: `Moon at ${ascending ? 'ascending' : 'descending'} node`,
      symbol: ascending ? '☊' : '☋',
      category: 'orbit crossing',
      time: node.time.date,
      summary: `The Moon crosses the ecliptic from ${ascending ? 'south to north' : 'north to south'}.`,
      details: 'A node crossing can contribute to an eclipse only when it also occurs close enough to new or full Moon.',
      bodies: ['Moon'],
      priority: 5,
    }))
    node = NextMoonNode(node)
  }
  return found
}

function collectEclipses(searchStart, start, end, observer) {
  const found = []
	let lunar = SearchLunarEclipse(searchStart)
	while (lunar.peak.date < end) {
	  if (lunar.peak.date >= start) {
		const horizon = horizontalPosition(Body.Moon, lunar.peak.date, observer)
		const contacts = lunarEclipseContacts(lunar, observer)
		const coverage = lunar.kind === 'penumbral'
		  ? 'Earth’s penumbra shades the lunar disc; no portion enters the umbra.'
		  : `${Math.round(lunar.obscuration * 100)}% of the lunar disc is calculated to be inside Earth’s umbra at peak.`
		found.push(skyEvent({
		  slug: 'lunar-eclipse',
		  name: `${titleCase(lunar.kind)} lunar eclipse`,
		  symbol: '◉',
		  category: 'eclipse',
		  time: lunar.peak.date,
		  summary: `The eclipse reaches its global peak. ${coverage}`,
		  details: `At the selected location, the Moon is ${round(Math.abs(horizon.altitude), 1)}° ${horizon.altitude >= 0 ? 'above' : 'below'} the horizon at peak. Visibility for the full event also depends on rise/set times and weather.`,
		  bodies: ['Moon', 'Earth', 'Sun'],
		  local: horizon.altitude >= 0,
		  contacts,
		  durationMinutes: round(2 * lunar.sd_penum, 1),
		  phaseDurations: {
			penumbralMinutes: round(2 * lunar.sd_penum, 1),
			partialMinutes: lunar.sd_partial ? round(2 * lunar.sd_partial, 1) : 0,
			totalMinutes: lunar.sd_total ? round(2 * lunar.sd_total, 1) : 0,
		  },
		  priority: 12,
		}))
	  }
	  lunar = NextLunarEclipse(lunar.peak)
	}

	// A local solar-eclipse peak can occur on the other side of a civil-month
	// boundary from the worldwide peak. Search both neighboring days and let the
	// caller's final event-time filter decide which month owns the occurrence.
	const solarSearchStart = new Date(start.getTime() - 2 * DAY)
	const solarSearchEnd = new Date(end.getTime() + 2 * DAY)
	let solar = SearchGlobalSolarEclipse(solarSearchStart)
	while (solar.peak.date < solarSearchEnd) {
		const localResult = SearchLocalSolarEclipse(new Date(solar.peak.date.getTime() - 2 * DAY), observer)
		const localMatch = Math.abs(localResult.peak.time.date.getTime() - solar.peak.date.getTime()) < 2 * DAY
		const coordinates = Number.isFinite(solar.latitude)
		  ? ` The global shadow peaks near ${round(solar.latitude, 1)}°, ${round(solar.longitude, 1)}°.`
		  : ''
		if (localMatch) {
		  const altitude = localResult.peak.altitude
		  const contacts = localSolarEclipseContacts(localResult)
		  found.push(skyEvent({
			slug: 'solar-eclipse-local',
			name: `${titleCase(localResult.kind)} solar eclipse at selected location`,
			symbol: '◍',
			category: 'eclipse',
			time: localResult.peak.time.date,
			summary: `${round(localResult.obscuration * 100, 1)}% maximum area obscuration is calculated at the selected location. The worldwide event is classified as ${solar.kind}.${coordinates}`,
			details: `The selected-location peak is separate from the worldwide peak. The Sun is ${round(Math.abs(altitude), 1)}° ${altitude >= 0 ? 'above' : 'below'} the selected horizon at local peak. Never view the Sun without certified eclipse protection.`,
			bodies: ['Sun', 'Moon', 'Earth'],
			local: altitude > 0,
			localKind: localResult.kind,
			localPeakTime: localResult.peak.time.date,
			localObscuration: round(localResult.obscuration * 100, 1),
			contacts,
			durationMinutes: round((localResult.partial_end.time.date.getTime() - localResult.partial_begin.time.date.getTime()) / 60000, 1),
			globalKind: solar.kind,
			globalPeakTime: solar.peak.date,
			priority: 12,
		  }))
		} else {
		  found.push(skyEvent({
			slug: 'solar-eclipse-global',
			name: `${titleCase(solar.kind)} solar eclipse · global event`,
			symbol: '◍',
			category: 'eclipse',
			time: solar.peak.date,
			summary: `A ${solar.kind} solar eclipse is visible somewhere on Earth.${coordinates}`,
			details: 'Astronomy Engine does not find this eclipse intersecting the selected location. The displayed time is the worldwide peak, not a local viewing time. Never view the Sun without certified eclipse protection.',
			bodies: ['Sun', 'Moon', 'Earth'],
			local: false,
			globalKind: solar.kind,
			globalPeakTime: solar.peak.date,
			priority: 12,
		  }))
		}
	  solar = NextGlobalSolarEclipse(solar.peak)
	}
  return found
}

function collectInnerPlanetEvents(searchStart, start, end, observer) {
  const found = []
  for (const meta of PLANETS.slice(0, 2)) {
    const elongation = SearchMaxElongation(meta.body, searchStart)
    if (elongation.time.date >= start && elongation.time.date < end) {
      found.push(skyEvent({
        slug: `${meta.name.toLowerCase()}-elongation`,
        name: `${meta.name} greatest elongation`,
        symbol: meta.symbol,
        category: 'visibility',
        time: elongation.time.date,
        summary: `${meta.name} reaches ${round(elongation.elongation, 1)}° from the Sun in the ${elongation.visibility} sky.`,
        details: 'Greatest elongation is often the most favorable part of an apparition, but local horizon and weather still govern visibility.',
        bodies: [meta.name, 'Sun'],
        value: `${round(elongation.elongation, 1)}°`,
        priority: 7,
      }))
    }

    const transit = SearchTransit(meta.body, searchStart)
    if (transit.peak.date >= start && transit.peak.date < end) {
	  const contacts = transitContacts(transit, observer)
	  const peakAltitude = contacts.find((contact) => contact.key === 'peak')?.altitude ?? null
      found.push(skyEvent({
        slug: `${meta.name.toLowerCase()}-transit`,
        name: `Transit of ${meta.name}`, symbol: '⊙', category: 'transit', time: transit.peak.date,
        summary: `${meta.name} crosses the Sun’s disc as seen geocentrically; minimum center separation is ${round(transit.separation, 2)} arcminutes.`,
        details: `The geocentric transit lasts ${formatMinutes((transit.finish.date.getTime() - transit.start.date.getTime()) / 60000)} from first to last contact. Contact altitudes indicate whether the Sun is above the selected horizon. Never look at the Sun without certified solar viewing equipment.`,
        bodies: [meta.name, 'Sun'],
		contacts,
		durationMinutes: round((transit.finish.date.getTime() - transit.start.date.getTime()) / 60000, 1),
		local: peakAltitude === null ? null : peakAltitude >= 0,
		priority: 12,
      }))
    }
  }

  const brightest = SearchPeakMagnitude(Body.Venus, searchStart)
  if (brightest.time.date >= start && brightest.time.date < end) {
    found.push(skyEvent({
      slug: 'venus-peak-brightness', name: 'Venus at peak brightness', symbol: '♀', category: 'visibility', time: brightest.time.date,
      summary: `Venus reaches calculated visual magnitude ${round(brightest.mag, 2)}.`,
      details: 'Smaller or more negative magnitude numbers indicate a brighter apparent object.',
      bodies: ['Venus'], value: `${round(brightest.mag, 2)} mag`, priority: 7,
    }))
  }
  return found
}

function lunarEclipseContacts(eclipse, observer) {
	const peak = eclipse.peak.date
	const contacts = []
	const add = (key, label, offsetMinutes) => {
	  const time = new Date(peak.getTime() + offsetMinutes * 60000)
	  const horizon = horizontalPosition(Body.Moon, time, observer)
	  contacts.push({ key, label, time, altitude: round(horizon.altitude, 1), visible: horizon.altitude >= 0 })
	}
	if (eclipse.sd_penum) add('penumbral-begin', 'Penumbral phase begins', -eclipse.sd_penum)
	if (eclipse.sd_partial) add('partial-begin', 'Partial phase begins', -eclipse.sd_partial)
	if (eclipse.sd_total) add('total-begin', 'Total phase begins', -eclipse.sd_total)
	add('peak', 'Maximum eclipse', 0)
	if (eclipse.sd_total) add('total-end', 'Total phase ends', eclipse.sd_total)
	if (eclipse.sd_partial) add('partial-end', 'Partial phase ends', eclipse.sd_partial)
	if (eclipse.sd_penum) add('penumbral-end', 'Penumbral phase ends', eclipse.sd_penum)
	return contacts
}

function localSolarEclipseContacts(eclipse) {
	return [
	  ['partial-begin', 'Partial eclipse begins', eclipse.partial_begin],
	  ['central-begin', `${titleCase(eclipse.kind)} phase begins`, eclipse.total_begin],
	  ['peak', 'Maximum eclipse', eclipse.peak],
	  ['central-end', `${titleCase(eclipse.kind)} phase ends`, eclipse.total_end],
	  ['partial-end', 'Partial eclipse ends', eclipse.partial_end],
	].flatMap(([key, label, event]) => event ? [{ key, label, time: event.time.date, altitude: round(event.altitude, 1), visible: event.altitude >= 0 }] : [])
}

function transitContacts(transit, observer) {
	return [
	  ['start', 'Transit begins', transit.start.date],
	  ['peak', 'Maximum transit', transit.peak.date],
	  ['finish', 'Transit ends', transit.finish.date],
	].map(([key, label, time]) => {
	  const horizon = horizontalPosition(Body.Sun, time, observer)
	  return { key, label, time, altitude: round(horizon.altitude, 1), visible: horizon.altitude >= 0 }
	})
}

function collectPlanetApsides(searchStart, start, end) {
  const found = []
  const bodies = [{ body: Body.Earth, name: 'Earth', symbol: '⊕' }, ...PLANETS]
  for (const meta of bodies) {
    const apsis = SearchPlanetApsis(meta.body, searchStart)
    if (apsis.time.date < start || apsis.time.date >= end) continue
    const perihelion = apsis.kind === ApsisKind.Pericenter
    found.push(skyEvent({
      slug: `${meta.name.toLowerCase()}-${perihelion ? 'perihelion' : 'aphelion'}`,
      name: `${meta.name} at ${perihelion ? 'perihelion' : 'aphelion'}`,
      symbol: meta.symbol,
      category: 'orbit',
      time: apsis.time.date,
      summary: `${meta.name} reaches its ${perihelion ? 'closest' : 'farthest'} point from the Sun in this orbit: ${round(apsis.dist_au, 4)} AU.`,
      details: 'This heliocentric distance extremum is not the same as the planet’s closest approach to Earth.',
      bodies: [meta.name, 'Sun'],
      value: `${round(apsis.dist_au, 4)} AU`,
      priority: meta.body === Body.Earth ? 7 : 3,
    }))
  }
  return found
}

function collectStations(context) {
  const found = []
  const scanStart = new Date(context.start.getTime() - DAY)
  const scanEnd = new Date(context.end.getTime() + DAY)
  for (const meta of PLANETS) {
    let left = scanStart
    let leftRate = motionRate(context, meta.body, left)
    for (let cursor = new Date(left.getTime() + 12 * HOUR); cursor <= scanEnd; cursor = new Date(cursor.getTime() + 12 * HOUR)) {
      const rightRate = motionRate(context, meta.body, cursor)
      if (Math.sign(leftRate) !== Math.sign(rightRate)) {
        const time = bisectRoot((date) => motionRate(context, meta.body, date), left, cursor)
        const afterward = motionRate(context, meta.body, new Date(time.getTime() + 6 * HOUR))
        const retrograde = afterward < 0
        found.push(skyEvent({
          slug: `${meta.name.toLowerCase()}-station-${retrograde ? 'retrograde' : 'direct'}`,
          name: `${meta.name} stations ${retrograde ? 'retrograde' : 'direct'}`,
          symbol: retrograde ? '℞' : 'D',
          category: 'station',
          time,
          summary: `${meta.name}’s apparent geocentric ecliptic motion changes from ${retrograde ? 'direct to retrograde' : 'retrograde to direct'}.`,
          details: 'Retrograde is an apparent reversal against the stars caused by the changing geometry of Earth and the planet; the planet does not reverse its orbit.',
          bodies: [meta.name, 'Earth'],
          priority: 8,
        }))
      }
      left = cursor
      leftRate = rightRate
    }
  }
  return found
}

function collectIngresses(context) {
  const found = []
  const scanStart = new Date(context.start.getTime() - 6 * HOUR)
  const scanEnd = new Date(context.end.getTime() + 6 * HOUR)
  for (const meta of SKY_BODIES) {
    let left = scanStart
    let before = zodiacIndex(context.longitude(meta.body, left))
    for (let cursor = new Date(left.getTime() + 6 * HOUR); cursor <= scanEnd; cursor = new Date(cursor.getTime() + 6 * HOUR)) {
      const after = zodiacIndex(context.longitude(meta.body, cursor))
      if (after !== before) {
        const time = bisectTransition((date) => zodiacIndex(context.longitude(meta.body, date)), before, left, cursor)
        const rate = meta.body === Body.Sun || meta.body === Body.Moon ? 1 : motionRate(context, meta.body, time)
        const target = ZODIAC[after]
        found.push(skyEvent({
          slug: `${meta.name.toLowerCase()}-ingress-${target.name.toLowerCase()}`,
          name: `${meta.name} ${rate < 0 ? 'retrogrades into' : 'enters'} ${target.name}`,
          symbol: target.symbol,
          category: 'zodiac ingress',
          time,
          summary: `${meta.name} crosses into the tropical 30° sector named ${target.name}.`,
          details: 'A tropical zodiac ingress is defined by ecliptic longitude from the moving March equinox; it is not a crossing of the unequal IAU constellation boundary.',
          bodies: [meta.name],
          frame: 'tropical ecliptic',
          priority: meta.body === Body.Moon ? 3 : 6,
        }))
        before = after
      }
      left = cursor
    }
  }
  return found
}

function collectAspects(context) {
  const found = []
	const bodies = [SKY_BODIES[0], SKY_BODIES[1], ...PLANETS]
  // Pad the sampling window so an exact aspect in the first or last six hours
  // of the requested local month still has samples on both sides of its
  // minimum. The final month filter removes any refined event outside scope.
  const times = sampleTimes(
    new Date(context.start.getTime() - 6 * HOUR),
    new Date(context.end.getTime() + 6 * HOUR),
    6,
  )
  for (let a = 0; a < bodies.length; a += 1) {
    for (let b = a + 1; b < bodies.length; b += 1) {
	  const first = bodies[a]
	  const second = bodies[b]
	  const sunMoonPair = first.body === Body.Sun && second.body === Body.Moon
      const separations = times.map((time) => longitudeSeparation(context.longitude(first.body, time), context.longitude(second.body, time)))
      for (const aspect of ASPECTS) {
		// New/full Moon and the two quarter squares already have richer lunar-phase
		// rows. Sextiles and trines are distinct octant relationships and must not
		// disappear with those duplicates.
		if (sunMoonPair && [0, 90, 180].includes(aspect.angle)) continue
        const orbs = separations.map((value) => Math.abs(value - aspect.angle))
        for (let i = 1; i < times.length - 1; i += 1) {
          if (orbs[i] > 4 || orbs[i] > orbs[i - 1] || orbs[i] >= orbs[i + 1]) continue
          const refined = minimizeBetween(
            (time) => Math.abs(longitudeSeparation(context.longitude(first.body, time), context.longitude(second.body, time)) - aspect.angle),
            times[i - 1],
            times[i + 1],
          )
          if (refined.value > 0.08) continue
          found.push(skyEvent({
            slug: `aspect-${first.name}-${second.name}-${aspect.name}`.toLowerCase(),
            name: `${first.name} ${aspect.name} ${second.name}`,
            symbol: aspect.symbol,
            category: 'alignment',
            time: refined.time,
            summary: `${first.name} and ${second.name} reach an apparent ecliptic-longitude separation of ${aspect.angle}° (orb ${round(refined.value, 3)}°).`,
            details: 'This is a geocentric projection onto the ecliptic. It does not mean the bodies are physically close or perfectly collinear in three-dimensional space.',
            bodies: [first.name, second.name],
            value: `${aspect.angle}°`,
            frame: 'geocentric ecliptic',
            priority: aspect.angle === 0 || aspect.angle === 180 ? 6 : 4,
          }))
          i += 1
        }
      }
    }
  }
  return found
}

function collectApproaches(context) {
  const found = []
  const bodies = [SKY_BODIES[1], ...PLANETS]
  const times = sampleTimes(
    new Date(context.start.getTime() - 6 * HOUR),
    new Date(context.end.getTime() + 6 * HOUR),
    6,
  )
  for (let a = 0; a < bodies.length; a += 1) {
    for (let b = a + 1; b < bodies.length; b += 1) {
      const first = bodies[a]
      const second = bodies[b]
      const threshold = first.body === Body.Moon ? 4 : 3
      const distances = times.map((time) => AngleBetween(context.vector(first.body, time), context.vector(second.body, time)))
      for (let i = 1; i < times.length - 1; i += 1) {
        if (distances[i] > threshold + 2 || distances[i] > distances[i - 1] || distances[i] >= distances[i + 1]) continue
        const refined = minimizeBetween(
          (time) => AngleBetween(context.vector(first.body, time), context.vector(second.body, time)),
          times[i - 1],
          times[i + 1],
        )
        if (refined.value > threshold) continue
        found.push(skyEvent({
          slug: `approach-${first.name}-${second.name}`.toLowerCase(),
          name: `${first.name}–${second.name} close approach`,
          symbol: '↔',
          category: 'close approach',
          time: refined.time,
          summary: `The two bodies reach a true apparent angular separation of ${round(refined.value, 2)}° on the sky.`,
          details: 'This angle is measured in three-dimensional sky coordinates as seen from Earth; their physical distances remain very different.',
          bodies: [first.name, second.name],
          value: `${round(refined.value, 2)}°`,
          frame: 'geocentric sky angle',
          priority: refined.value < 1 ? 8 : 5,
        }))
        i += 1
      }
    }
  }
  return found
}

function collectPlanetGathering(context) {
  const visible = PLANETS.slice(0, 5)
  let best = null
  for (const time of sampleTimes(context.start, context.end, 12)) {
    const points = visible.map((meta) => ({ ...meta, longitude: context.longitude(meta.body, time) })).sort((a, b) => a.longitude - b.longitude)
    const doubled = [...points, ...points.map((item) => ({ ...item, longitude: item.longitude + 360 }))]
    for (let left = 0; left < points.length; left += 1) {
      let right = left
      while (right + 1 < left + points.length && doubled[right + 1].longitude - doubled[left].longitude <= 15) right += 1
      const count = right - left + 1
      if (count < 3) continue
      const span = doubled[right].longitude - doubled[left].longitude
      if (!best || count > best.count || (count === best.count && span < best.span)) {
        best = { time, count, span, bodies: doubled.slice(left, right + 1).map((item) => item.name) }
      }
    }
  }
  if (!best) return []
  return [skyEvent({
    slug: `planet-gathering-${best.bodies.join('-')}`.toLowerCase(),
	name: `${best.count}-planet sampled gathering`,
    symbol: '⋮',
    category: 'alignment',
    time: best.time,
	summary: `At this 12-hour sample, ${best.bodies.join(', ')} fit within a ${round(best.span, 1)}° span of geocentric ecliptic longitude.`,
	details: 'This is the tightest sampled snapshot in the month, not a minute-precise extremum. “Gathering” is a display rule (three or more naked-eye planets within 15°), not a claim of perfect physical alignment, equal visibility, or a discrete beginning and end.',
    bodies: best.bodies,
    value: `${round(best.span, 1)}° span`,
    frame: 'geocentric ecliptic',
    priority: 7,
  })]
}

function positionFor(meta, time, observer) {
  const longitude = apparentLongitude(meta.body, time)
  const eq = Equator(meta.body, time, observer, true, true)
	const j2000 = Equator(meta.body, time, observer, false, true)
  const horizon = Horizon(time, observer, eq.ra, eq.dec, 'normal')
  const illumination = safeIllumination(meta.body, time)
	const constellation = safeConstellation(j2000.ra, j2000.dec)
  const rate = meta.body === Body.Sun || meta.body === Body.Moon ? null : apparentMotionRate(meta.body, time)
  const zodiac = ZODIAC[zodiacIndex(longitude)]
  return {
    ...meta,
    longitude: round(longitude, 2),
    sign: zodiac.name,
    signSymbol: zodiac.symbol,
    degree: round(longitude % 30, 1),
    constellation: constellation?.name || '—',
    constellationSymbol: constellation?.symbol || '',
    rightAscension: round(eq.ra, 3),
    declination: round(eq.dec, 2),
    altitude: round(horizon.altitude, 1),
    azimuth: round(horizon.azimuth, 1),
    direction: compassPoint(horizon.azimuth),
    aboveHorizon: horizon.altitude >= 0,
    retrograde: rate !== null && rate < 0,
    motionRate: rate === null ? null : round(rate, 3),
    magnitude: illumination && Number.isFinite(illumination.mag) ? round(illumination.mag, 2) : null,
    distance: illumination ? formatDistance(meta.body, illumination.geo_dist) : null,
  }
}

function currentAspects(time) {
  const bodies = [SKY_BODIES[0], SKY_BODIES[1], ...PLANETS]
  const positions = new Map(bodies.map((meta) => [meta.body, apparentLongitude(meta.body, time)]))
  const found = []
  for (let a = 0; a < bodies.length; a += 1) {
    for (let b = a + 1; b < bodies.length; b += 1) {
      const separation = longitudeSeparation(positions.get(bodies[a].body), positions.get(bodies[b].body))
      const aspect = ASPECTS.map((item) => ({ ...item, orb: Math.abs(separation - item.angle) })).sort((x, y) => x.orb - y.orb)[0]
      if (aspect.orb <= 3) found.push({ a: bodies[a].name, b: bodies[b].name, ...aspect, separation: round(separation, 2), orb: round(aspect.orb, 2) })
    }
  }
  return found.sort((a, b) => a.orb - b.orb)
}

function skyEvent({ slug, name, symbol, category, time, summary, details, bodies = [], value = '', local = null, frame = '', priority = 1, ...metadata }) {
	return { slug, name, symbol, category, time: new Date(time), summary, details, bodies, value, local, frame, priority, ...metadata }
}

function finalizeEvent(item, timeZone) {
	const event = {
    ...item,
    id: `${item.slug}-${item.time.toISOString()}`,
    time: item.time.toISOString(),
    date: localDateKey(item.time, timeZone),
    sourceName: 'Astronomy Engine',
    sourceUrl: ENGINE_SOURCE,
	}
	if (item.localPeakTime) event.localPeakTime = new Date(item.localPeakTime).toISOString()
	if (item.globalPeakTime) event.globalPeakTime = new Date(item.globalPeakTime).toISOString()
	if (item.contacts) event.contacts = item.contacts.map((contact) => ({ ...contact, time: new Date(contact.time).toISOString() }))
	return event
}

function collectSafely(target, warnings, family, collector) {
  try {
    target.push(...collector())
  } catch (error) {
    // One optional event family should not suppress the rest of the calculated
    // month, but a partial result must be disclosed to the reader.
	pushWarning(warnings, family, error)
	console.warn(`Lunarium astronomy calculation skipped ${family}:`, error)
  }
}

function pushWarning(warnings, family, error) {
	const message = error instanceof Error ? error.message : String(error)
	if (!warnings.some((warning) => warning.family === family && warning.message === message)) warnings.push({ family, message })
}

function apparentLongitude(body, time) {
  if (body === Body.Sun) return normalize360(SunPosition(time).elon)
  if (body === Body.Moon) return normalize360(EclipticGeoMoon(time).lon)
  return normalize360(Ecliptic(GeoVector(body, time, true)).elon)
}

function motionRate(context, body, time) {
  const before = context.longitude(body, new Date(time.getTime() - 3 * HOUR))
  const after = context.longitude(body, new Date(time.getTime() + 3 * HOUR))
  return signedAngle(after - before) * 4
}

function apparentMotionRate(body, time) {
  const before = apparentLongitude(body, new Date(time.getTime() - 3 * HOUR))
  const after = apparentLongitude(body, new Date(time.getTime() + 3 * HOUR))
  return signedAngle(after - before) * 4
}

function horizontalPosition(body, time, observer, refraction = 'normal') {
  const eq = Equator(body, time, observer, true, true)
	return Horizon(time, observer, eq.ra, eq.dec, refraction)
}

function makeObserver(location) {
  return new Observer(Number(location.latitude), Number(location.longitude), Number(location.elevation || 0))
}

function sampleTimes(start, end, hours) {
  const times = []
  for (let value = start.getTime(); value <= end.getTime(); value += hours * HOUR) times.push(new Date(value))
	if (!times.length || times[times.length - 1].getTime() !== end.getTime()) times.push(new Date(end))
  return times
}

function minimizeBetween(fn, start, end) {
  let left = start.getTime()
  let right = end.getTime()
  for (let i = 0; i < 34; i += 1) {
    const third = (right - left) / 3
    const a = left + third
    const b = right - third
    if (fn(new Date(a)) <= fn(new Date(b))) right = b
    else left = a
  }
  const time = new Date((left + right) / 2)
  return { time, value: fn(time) }
}

function bisectRoot(fn, start, end) {
  let left = start.getTime()
  let right = end.getTime()
  let leftValue = fn(new Date(left))
  for (let i = 0; i < 42; i += 1) {
    const middle = (left + right) / 2
    const middleValue = fn(new Date(middle))
    if (Math.sign(leftValue) === Math.sign(middleValue)) {
      left = middle
      leftValue = middleValue
    } else right = middle
  }
  return new Date((left + right) / 2)
}

function bisectTransition(fn, initialValue, start, end) {
  let left = start.getTime()
  let right = end.getTime()
  for (let i = 0; i < 38; i += 1) {
    const middle = (left + right) / 2
    if (fn(new Date(middle)) === initialValue) left = middle
    else right = middle
  }
  return new Date((left + right) / 2)
}

function localDateKey(date, timeZone) {
  const parts = new Intl.DateTimeFormat('en-US', {
    timeZone: validTimeZone(timeZone), year: 'numeric', month: '2-digit', day: '2-digit',
  }).formatToParts(date)
  const value = Object.fromEntries(parts.map((part) => [part.type, part.value]))
  return `${value.year}-${value.month}-${value.day}`
}

export function zonedDateTimeToUTC(dateString, hour, timeZone) {
	const parts = assertSupportedDate(dateString)
	return resolveZonedDateTimeToUTC(parts, dateString, hour, timeZone)
}

function zonedBoundaryDateTimeToUTC(dateString, hour, timeZone) {
	const parts = parseCivilDate(dateString)
	const upperBoundary = parts.year === SKY_YEAR_MAX + 1 && parts.month === 1 && parts.day === 1
	if (!upperBoundary) assertSupportedYear(parts.year)
	return resolveZonedDateTimeToUTC(parts, dateString, hour, timeZone)
}

function resolveZonedDateTimeToUTC({ year, month, day }, dateString, hour, timeZone) {
	if (!Number.isInteger(hour) || hour < 0 || hour > 23) throw new RangeError('Civil hour must be an integer from 0 through 23.')
  const wanted = Date.UTC(year, month - 1, day, hour)
  let guess = wanted
  const formatter = new Intl.DateTimeFormat('en-US', {
    timeZone: validTimeZone(timeZone), hourCycle: 'h23', year: 'numeric', month: '2-digit', day: '2-digit', hour: '2-digit', minute: '2-digit', second: '2-digit',
  })
	for (let i = 0; i < 6; i += 1) {
    const parts = Object.fromEntries(formatter.formatToParts(new Date(guess)).map((part) => [part.type, part.value]))
    const represented = Date.UTC(Number(parts.year), Number(parts.month) - 1, Number(parts.day), Number(parts.hour), Number(parts.minute), Number(parts.second))
    guess += wanted - represented
  }
	const representedAtGuess = wallTimeValue(formatter, new Date(guess))
	if (representedAtGuess === wanted) return new Date(guess)

	// Some zones advance their clocks at midnight, so the requested wall time
	// may not exist. Match Temporal's compatible/later behavior by returning the
	// first real instant whose local wall clock is at or beyond the target.
	const scanStart = wanted - 36 * HOUR
	const scanEnd = wanted + 36 * HOUR
	for (let value = scanStart; value <= scanEnd; value += 60_000) {
	  if (wallTimeValue(formatter, new Date(value)) >= wanted) return new Date(value)
	}
	throw new RangeError(`Unable to resolve ${dateString} ${hour}:00 in ${validTimeZone(timeZone)}`)
}

function assertSupportedYear(year) {
	if (!Number.isInteger(year) || year < SKY_YEAR_MIN || year > SKY_YEAR_MAX) {
	  throw new RangeError(`Sky calculations support civil years ${SKY_YEAR_MIN}–${SKY_YEAR_MAX}.`)
	}
	return year
}

function assertSupportedDate(dateString) {
	const parts = parseCivilDate(dateString)
	assertSupportedYear(parts.year)
	return parts
}

function parseCivilDate(dateString) {
	const match = /^(\d{4})-(\d{2})-(\d{2})$/.exec(String(dateString))
	if (!match) throw new RangeError('Sky date must use YYYY-MM-DD.')
	const year = Number(match[1])
	const month = Number(match[2])
	const day = Number(match[3])
	const probe = new Date(0)
	probe.setUTCHours(0, 0, 0, 0)
	probe.setUTCFullYear(year, month - 1, day)
	if (probe.getUTCFullYear() !== year || probe.getUTCMonth() !== month - 1 || probe.getUTCDate() !== day) throw new RangeError(`Invalid civil date: ${dateString}.`)
	return { year, month, day }
}

function addCivilDays(dateString, amount) {
	const { year, month, day } = assertSupportedDate(dateString)
	const date = new Date(Date.UTC(year, month - 1, day + amount))
	const nextYear = date.getUTCFullYear()
	if (nextYear > SKY_YEAR_MAX && !(nextYear === SKY_YEAR_MAX + 1 && date.getUTCMonth() === 0 && date.getUTCDate() === 1)) assertSupportedYear(nextYear)
	return `${String(nextYear).padStart(4, '0')}-${String(date.getUTCMonth() + 1).padStart(2, '0')}-${String(date.getUTCDate()).padStart(2, '0')}`
}

function wallTimeValue(formatter, date) {
	const parts = Object.fromEntries(formatter.formatToParts(date).map((part) => [part.type, part.value]))
	return Date.UTC(Number(parts.year), Number(parts.month) - 1, Number(parts.day), Number(parts.hour), Number(parts.minute), Number(parts.second))
}

function validTimeZone(value) {
  try {
    new Intl.DateTimeFormat('en-US', { timeZone: value }).format()
    return value
  } catch {
    return 'UTC'
  }
}

function safeIllumination(body, time) {
  try { return Illumination(body, time) } catch { return null }
}

function safeConstellation(ra, dec) {
  try { return Constellation(ra, dec) } catch { return null }
}

function zodiacRelationship(first, second) {
  const a = ZODIAC.findIndex((item) => item.name === first)
  const b = ZODIAC.findIndex((item) => item.name === second)
  const distance = Math.min((a - b + 12) % 12, (b - a + 12) % 12)
  if (distance === 0) return { kind: 'same', label: 'concentrated emphasis' }
  if ([2, 4].includes(distance)) return { kind: 'supportive', label: 'flowing relationship' }
  if ([3, 6].includes(distance)) return { kind: 'dynamic', label: 'dynamic relationship' }
  return { kind: 'neutral', label: 'open-ended relationship' }
}

function formatDistance(body, au) {
  if (body === Body.Moon) return `${Math.round(au * AU_KM).toLocaleString('en-US')} km`
  return `${round(au, body === Body.Sun ? 4 : 3)} AU`
}

function formatMinutes(value) {
	const minutes = Math.round(value)
	const hours = Math.floor(minutes / 60)
	const remainder = minutes % 60
	if (!hours) return `${remainder} min`
	return `${hours} hr${hours === 1 ? '' : 's'}${remainder ? ` ${remainder} min` : ''}`
}

function phaseName(angle) {
  if (angle < 11.25 || angle >= 348.75) return 'New Moon'
  if (angle < 78.75) return 'Waxing crescent'
  if (angle < 101.25) return 'First quarter'
  if (angle < 168.75) return 'Waxing gibbous'
  if (angle < 191.25) return 'Full Moon'
  if (angle < 258.75) return 'Waning gibbous'
  if (angle < 281.25) return 'Last quarter'
  return 'Waning crescent'
}

function compassPoint(azimuth) {
  return ['N', 'NE', 'E', 'SE', 'S', 'SW', 'W', 'NW'][Math.round(normalize360(azimuth) / 45) % 8]
}

function zodiacIndex(longitude) {
  return Math.floor(normalize360(longitude) / 30) % 12
}

function longitudeSeparation(a, b) {
  return Math.abs(signedAngle(a - b))
}

function normalize360(value) {
  return ((value % 360) + 360) % 360
}

function signedAngle(value) {
  return ((value + 540) % 360) - 180
}

function round(value, digits) {
  const scale = 10 ** digits
  return Math.round(value * scale) / scale
}

function titleCase(value) {
  return value.charAt(0).toUpperCase() + value.slice(1)
}
