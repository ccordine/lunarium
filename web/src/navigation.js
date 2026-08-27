export function hasMonthReference(reference) {
  return Number.isInteger(reference?.year) && reference.year > 0 && Number.isInteger(reference?.month) && reference.month > 0
}

export function hebrewMonthTarget(data, delta) {
  const target = delta < 0 ? data?.previous : data?.next
  return hasMonthReference(target) ? { year: target.year, month: target.month } : null
}

export function monthNavigationState({ view, calendarSystem, data, cursor, navigationLoading, minYear, maxYear }) {
  const hebrewCalendar = view === 'calendar' && calendarSystem === 'hebrew'
  if (hebrewCalendar) {
    return {
      atSupportedStart: Boolean(data) && !hasMonthReference(data.previous),
      atSupportedEnd: Boolean(data) && !hasMonthReference(data.next),
      waitingForHebrewMonth: navigationLoading || !data,
    }
  }
  return {
    atSupportedStart: cursor.year === minYear && (view === 'atlas' || cursor.month === 1),
    atSupportedEnd: cursor.year === maxYear && (view === 'atlas' || cursor.month === 12),
    waitingForHebrewMonth: false,
  }
}

// A Hebrew month containing January 1, 1900 or December 31, 2100 crosses the
// supported civil range. When switching from the Gregorian edge months, use a
// day guaranteed to belong to the nearest fully supported Hebrew month.
export function hebrewAnchorDay(data, selectedDate, minYear, maxYear) {
  if (!data?.days?.length) return null
  if (data.year === minYear && data.month === 1) return data.days[data.days.length - 1]
  if (data.year === maxYear && data.month === 12) return data.days[0]
  return data.days.find((day) => day.date === selectedDate) || data.days[0]
}

export function hebrewBootstrapAnchor(data, selectedDate, minYear, maxYear) {
  const preferred = data?.days?.find((day) => day.date === selectedDate) || data?.days?.find((day) => day.isToday)
  return hebrewAnchorDay(data, preferred?.date, minYear, maxYear)
}
