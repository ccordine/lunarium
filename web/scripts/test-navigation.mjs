import assert from 'node:assert/strict'
import { hebrewAnchorDay, hebrewBootstrapAnchor, hebrewMonthTarget, monthNavigationState } from '../src/navigation.js'

const common = {
  view: 'calendar',
  calendarSystem: 'hebrew',
  cursor: { year: 2026, month: 8 },
  navigationLoading: false,
  minYear: 1900,
  maxYear: 2100,
}
const normal = {
  previous: { year: 5786, month: 5 },
  next: { year: 5787, month: 7 },
}
assert.deepEqual(monthNavigationState({ ...common, data: normal }), {
  atSupportedStart: false,
  atSupportedEnd: false,
  waitingForHebrewMonth: false,
})
assert.deepEqual(hebrewMonthTarget(normal, -1), { year: 5786, month: 5 })
assert.deepEqual(hebrewMonthTarget(normal, 1), { year: 5787, month: 7 })

const lowerBoundary = { previous: null, next: { year: 5660, month: 11 } }
assert.equal(monthNavigationState({ ...common, data: lowerBoundary }).atSupportedStart, true)
assert.equal(hebrewMonthTarget(lowerBoundary, -1), null)

const upperBoundary = { previous: { year: 5861, month: 8 }, next: null }
assert.equal(monthNavigationState({ ...common, data: upperBoundary }).atSupportedEnd, true)
assert.equal(hebrewMonthTarget(upperBoundary, 1), null)
assert.equal(monthNavigationState({ ...common, data: null, navigationLoading: true }).waitingForHebrewMonth, true)

const january1900 = {
  year: 1900,
  month: 1,
  days: [{ date: '1900-01-01' }, { date: '1900-01-31' }],
}
assert.equal(hebrewAnchorDay(january1900, '1900-01-01', 1900, 2100).date, '1900-01-31')
assert.equal(hebrewBootstrapAnchor(january1900, '1900-01-01', 1900, 2100).date, '1900-01-31', 'delayed Hebrew bootstrap must skip the range-crossing January 1 month')
const december2100 = {
  year: 2100,
  month: 12,
  days: [{ date: '2100-12-01' }, { date: '2100-12-31' }],
}
assert.equal(hebrewAnchorDay(december2100, '2100-12-31', 1900, 2100).date, '2100-12-01')
assert.equal(hebrewBootstrapAnchor(december2100, '2100-12-31', 1900, 2100).date, '2100-12-01', 'delayed Hebrew bootstrap must skip the range-crossing December 31 month')

console.log('Calendar navigation boundary checks passed.')
