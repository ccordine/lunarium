import assert from 'node:assert/strict'
import {
  buildHoroscope,
  buildSkyMonth,
  buildSkySnapshot,
  SKY_YEAR_MAX,
  SKY_YEAR_MIN,
  zonedDateTimeToUTC,
} from '../src/astronomy.js'

const newYork = {
  name: 'New York, NY',
  latitude: 40.7128,
  longitude: -74.006,
  timezone: 'America/New_York',
}

const april2024 = buildSkyMonth(2024, 4, newYork)
const greatAmericanEclipse = april2024.events.find((event) => event.globalKind === 'total')
assert.ok(greatAmericanEclipse, 'April 2024 should contain the total solar eclipse')
assert.equal(greatAmericanEclipse.date, '2024-04-08')
assert.equal(greatAmericanEclipse.local, true, 'the eclipse should be locally visible from New York')
assert.equal(greatAmericanEclipse.localKind, 'partial', 'New York experienced a partial, not total, eclipse')
assert.match(greatAmericanEclipse.localPeakTime, /^2024-04-08T19:25:/)
assert.match(greatAmericanEclipse.globalPeakTime, /^2024-04-08T18:17:/)
assert.equal(greatAmericanEclipse.time, greatAmericanEclipse.localPeakTime, 'a location-specific event should be timed at its local peak')
assert.deepEqual(greatAmericanEclipse.contacts.map((contact) => contact.key), ['partial-begin', 'peak', 'partial-end'])
assert.match(greatAmericanEclipse.contacts[0].time, /^2024-04-08T18:10:/)
assert.match(greatAmericanEclipse.contacts[2].time, /^2024-04-08T20:36:/)
assert.ok(greatAmericanEclipse.contacts.every((contact) => contact.visible), 'all New York eclipse contacts should be above the horizon')
assert.ok(Math.abs(greatAmericanEclipse.durationMinutes - 145.7) < 0.2)

const march2026 = buildSkyMonth(2026, 3, newYork)
assert.equal(march2026.events.find((event) => event.name === 'March equinox')?.date, '2026-03-20')

const august2026 = buildSkyMonth(2026, 8, newYork)
assert.ok(august2026.events.some((event) => event.globalKind === 'total' && event.date === '2026-08-12'))
assert.ok(august2026.events.some((event) => event.name === 'Partial lunar eclipse' && event.date === '2026-08-28'))
assert.ok(august2026.events.some((event) => event.category === 'alignment'), 'monthly timeline should include exact ecliptic alignments')
assert.ok(august2026.events.some((event) => event.category === 'alignment' && event.bodies.includes('Moon')), 'monthly timeline should include bounded lunar aspects')
assert.ok(august2026.events.some((event) => event.category === 'close approach'), 'monthly timeline should include true angular close approaches')
const sunMoonOctants = august2026.events.filter((event) => event.category === 'alignment' && event.bodies.join('|') === 'Sun|Moon')
assert.deepEqual(sunMoonOctants.map((event) => event.name), ['Sun trine Moon', 'Sun sextile Moon', 'Sun sextile Moon', 'Sun trine Moon'], 'Sun–Moon sextiles and trines should remain after richer quarter rows are deduplicated')
assert.ok(!sunMoonOctants.some((event) => /conjunction|square|opposition/.test(event.name)), 'quarter-equivalent Sun–Moon aspects should not duplicate lunar phase rows')
assert.equal(new Set(august2026.events.map((event) => event.id)).size, august2026.events.length, 'event IDs should be unique')
assert.ok(august2026.events.every((event) => event.date.startsWith('2026-08-')), 'all events should map into the requested local month')
assert.deepEqual(august2026.warnings, [], 'a fully calculated month should disclose no skipped families')
const augustLunarEclipse = august2026.events.find((event) => event.name === 'Partial lunar eclipse')
assert.deepEqual(augustLunarEclipse.contacts.map((contact) => contact.key), ['penumbral-begin', 'partial-begin', 'peak', 'partial-end', 'penumbral-end'])
assert.ok(augustLunarEclipse.phaseDurations.penumbralMinutes > augustLunarEclipse.phaseDurations.partialMinutes)

const february2026 = buildSkyMonth(2026, 2, newYork)
assert.ok(february2026.events.some((event) => event.name === 'Mercury stations retrograde' && event.date === '2026-02-26'), 'retrograde station finder should materialize a known 2026 station')

const kiritimati = buildSkyMonth(2026, 8, { name: 'Kiritimati', latitude: 1.8721, longitude: -157.4278, timezone: 'Pacific/Kiritimati' })
assert.ok(kiritimati.events.every((event) => event.date.startsWith('2026-08-')), 'UTC+14 boundary events should stay inside the requested local month')

const july2000 = buildSkyMonth(2000, 7, newYork)
const julySolarEclipses = july2000.events.filter((event) => event.globalKind && event.bodies.includes('Sun'))
assert.equal(julySolarEclipses.length, 2, 'July 2000 should retain both global solar eclipses')
assert.ok(julySolarEclipses.some((event) => event.globalPeakTime.startsWith('2000-07-31T02:13:')))

const march1904 = buildSkyMonth(1904, 3, newYork)
const marchLunarEclipses = march1904.events.filter((event) => event.category === 'eclipse' && event.bodies[0] === 'Moon')
assert.equal(marchLunarEclipses.length, 2, 'March 1904 should retain both lunar eclipses')
assert.ok(marchLunarEclipses.some((event) => event.time.startsWith('1904-03-31T12:32:')))

const march2022 = buildSkyMonth(2022, 3, newYork)
assert.ok(march2022.events.some((event) => event.name === 'Moon–Mercury close approach' && event.time.startsWith('2022-04-01T02:53:') && event.date === '2022-03-31'), 'DST-short months should retain refined events in their final local hours')

const november2019 = buildSkyMonth(2019, 11, newYork)
const mercuryTransit = november2019.events.find((event) => event.name === 'Transit of Mercury')
assert.ok(mercuryTransit, 'November 2019 should include Mercury’s transit')
assert.deepEqual(mercuryTransit.contacts.map((contact) => contact.key), ['start', 'peak', 'finish'])
assert.match(mercuryTransit.contacts[0].time, /^2019-11-11T12:34:/)
assert.match(mercuryTransit.contacts[2].time, /^2019-11-11T18:03:/)
assert.ok(Math.abs(mercuryTransit.durationMinutes - 328.5) < 0.2)

const crossBoundaryEclipse = buildSkyMonth(2022, 5, {
  name: 'Custom observer',
  latitude: -15,
  longitude: -70,
  timezone: 'Africa/Johannesburg',
})
assert.ok(crossBoundaryEclipse.events.some((event) => event.globalPeakTime?.startsWith('2022-04-30T20:41:') && event.date === '2022-05-01'), 'a selected-location eclipse peak must survive when its worldwide peak belongs to the preceding civil month')

assert.equal(zonedDateTimeToUTC('2008-06-01', 0, 'Africa/Casablanca').toISOString(), '2008-06-01T00:00:00.000Z', 'a midnight DST gap must not truncate the preceding month')

const snapshot = buildSkySnapshot('2026-08-27', newYork)
assert.equal(snapshot.positions.length, 10)
assert.equal(snapshot.positions[0].name, 'Sun')
assert.equal(snapshot.positions[1].name, 'Moon')
assert.ok(snapshot.positions.every((body) => Number.isFinite(body.longitude)))
assert.ok(snapshot.retrograde.includes('Saturn'))
assert.deepEqual(snapshot.warnings, [])
assert.equal(snapshot.almanac.timeZone, newYork.timezone)
assert.equal(snapshot.almanac.bodies.length, 9, 'daily almanac should include the Moon and eight tracked planets')
const solarTimes = Object.fromEntries(snapshot.almanac.solar.map((event) => [event.key, event.time]))
assert.match(solarTimes['astronomical-dawn'], /^2026-08-27T08:39:/)
assert.match(solarTimes.sunrise, /^2026-08-27T10:18:/)
assert.match(solarTimes.sunset, /^2026-08-27T23:36:/)
assert.match(solarTimes['astronomical-dusk'], /^2026-08-28T01:14:/)
const moonAlmanac = snapshot.almanac.bodies.find((body) => body.name === 'Moon')
assert.match(moonAlmanac.rise, /^2026-08-27T23:25:/)
assert.match(moonAlmanac.culmination, /^2026-08-27T04:13:/)
assert.match(moonAlmanac.set, /^2026-08-27T09:34:/)
assert.ok(Number.isFinite(moonAlmanac.culminationAltitude))

const historicalSnapshot = buildSkySnapshot('2020-01-06', newYork)
assert.equal(historicalSnapshot.positions.find((body) => body.name === 'Mars')?.constellation, 'Libra', 'IAU constellation lookup requires J2000 coordinates')

const invalidZoneSnapshot = buildSkySnapshot('2026-08-27', { ...newYork, timezone: 'Not/A_Timezone' })
assert.equal(invalidZoneSnapshot.referenceLabel, '21:00 UTC')

const horoscope = buildHoroscope(snapshot, 'Leo')
assert.equal(horoscope.sign.name, 'Leo')
assert.match(horoscope.caveat, /not a scientific method/i)

assert.equal(SKY_YEAR_MIN, 1900)
assert.equal(SKY_YEAR_MAX, 2100)
assert.throws(() => buildSkyMonth(1899, 12, newYork), /1900–2100/)
assert.throws(() => buildSkySnapshot('2101-01-01', newYork), /1900–2100/)
assert.throws(() => zonedDateTimeToUTC('0050-01-01', 0, 'UTC'), /1900–2100/, 'two-digit-era years must not silently map into 19xx')
assert.equal(buildSkyMonth(2100, 12, newYork).year, 2100, 'the complete upper-bound year should remain calculable')
assert.equal(buildSkySnapshot('2100-12-31', newYork).date, '2100-12-31', 'the final supported civil day should remain calculable')

console.log(`Astronomy checks passed: ${august2026.events.length} August 2026 events, ${snapshot.positions.length} bodies.`)
