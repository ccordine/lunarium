import { useEffect, useMemo, useRef, useState } from 'react'
import { getAbout, getHebrewMonth, getMonth, getObservances } from './api.js'
import { buildHoroscope, buildSkyMonth, buildSkySnapshot, SKY_YEAR_MAX, SKY_YEAR_MIN, ZODIAC } from './astronomy.js'
import { hebrewAnchorDay, hebrewBootstrapAnchor, hebrewMonthTarget, monthNavigationState } from './navigation.js'

const DEFAULT_LOCATION = {
  name: 'New York, NY',
  latitude: 40.7128,
  longitude: -74.006,
  timezone: 'America/New_York',
}

const TRADITIONS = {
  christianity: { label: 'Christian', short: 'C', icon: '✦' },
  judaism: { label: 'Jewish', short: 'J', icon: '✡' },
  islam: { label: 'Islamic', short: 'I', icon: '☾' },
  polytheist: { label: 'Pagan & polytheist', short: 'P', icon: '☉' },
  ancient: { label: 'Ancient world', short: 'A', icon: '𓂀' },
}

const WEEKDAYS = ['SUN', 'MON', 'TUE', 'WED', 'THU', 'FRI', 'SAT']
const APP_VIEWS = new Set(['calendar', 'sky', 'atlas', 'readings'])

function App() {
	const [location, setLocation] = useState(readLocation)
	const initialDate = dateKeyInTimeZone(new Date(), location.timezone)
	const [cursor, setCursor] = useState(() => cursorForDateKey(initialDate))
	const [calendarSystem, setCalendarSystem] = useState('hebrew')
	const [hebrewCursor, setHebrewCursor] = useState(null)
  const [view, setView] = useState(() => {
    const requested = new URLSearchParams(window.location.search).get('view')
    return APP_VIEWS.has(requested) ? requested : 'calendar'
  })
  const [monthData, setMonthData] = useState(null)
	const [hebrewData, setHebrewData] = useState(null)
  const [indexData, setIndexData] = useState(null)
  const [selectedDate, setSelectedDate] = useState(initialDate)
  const [filters, setFilters] = useState({ christianity: true, judaism: true, islam: true, polytheist: true, ancient: true })
  const [monthLoading, setMonthLoading] = useState(true)
  const [hebrewLoading, setHebrewLoading] = useState(true)
  const [indexLoading, setIndexLoading] = useState(false)
  const [monthError, setMonthError] = useState('')
  const [hebrewError, setHebrewError] = useState('')
  const [indexError, setIndexError] = useState('')
  const [about, setAbout] = useState(null)
  const [aboutError, setAboutError] = useState('')
  const [dialog, setDialog] = useState(null)

  useEffect(() => {
    const controller = new AbortController()
    let active = true
    setMonthLoading(true)
    setMonthError('')
    getMonth(cursor.year, cursor.month, location, controller.signal)
      .then((payload) => {
        if (!active) return
        setMonthData(payload)
      })
      .catch((reason) => {
        if (active && reason.name !== 'AbortError') setMonthError(reason.message)
      })
      .finally(() => {
        if (active) setMonthLoading(false)
      })
    return () => {
      active = false
      controller.abort()
    }
  }, [cursor, location])

	useEffect(() => {
	if (view !== 'calendar' || calendarSystem !== 'hebrew' || hebrewCursor || !monthData || monthData.year !== cursor.year || monthData.month !== cursor.month) return
	const anchor = hebrewBootstrapAnchor(monthData, selectedDate, SKY_YEAR_MIN, SKY_YEAR_MAX)
	if (!anchor) return
	if (anchor.date !== selectedDate) setSelectedDate(anchor.date)
	setHebrewCursor({ year: anchor.sacredDates.hebrewYear, month: anchor.sacredDates.hebrewMonthNumber })
  }, [view, calendarSystem, hebrewCursor, monthData, selectedDate, cursor.year, cursor.month])

  useEffect(() => {
	if (view !== 'calendar' || calendarSystem !== 'hebrew' || !hebrewCursor) return
	const controller = new AbortController()
	let active = true
	setHebrewLoading(true)
	setHebrewError('')
	getHebrewMonth(hebrewCursor.year, hebrewCursor.month, location, controller.signal)
	  .then((payload) => {
		if (!active) return
		setHebrewData(payload)
		setSelectedDate((current) => payload.days.some((day) => day.date === current) ? current : payload.days[0].date)
	  })
	  .catch((reason) => {
		if (active && reason.name !== 'AbortError') setHebrewError(reason.message)
	  })
	  .finally(() => {
		if (active) setHebrewLoading(false)
	  })
	return () => {
	  active = false
	  controller.abort()
	}
  }, [view, calendarSystem, hebrewCursor, location])

  useEffect(() => {
	if (view !== 'atlas') return
	const controller = new AbortController()
	let active = true
	setIndexLoading(true)
	setIndexError('')
	getObservances(cursor.year, controller.signal)
	  .then((payload) => {
		if (active) setIndexData(payload)
	  })
	  .catch((reason) => {
		if (active && reason.name !== 'AbortError') setIndexError(reason.message)
	  })
	  .finally(() => {
		if (active) setIndexLoading(false)
	  })
	return () => {
	  active = false
	  controller.abort()
	}
  }, [cursor.year, view])

	const gregorianData = monthData?.year === cursor.year && monthData?.month === cursor.month ? monthData : null
	const calendarData = calendarSystem === 'hebrew' ? hebrewData : gregorianData
	const visibleData = view === 'calendar' ? calendarData : gregorianData
	const error = view === 'atlas' ? indexError : view === 'calendar' && calendarSystem === 'hebrew' ? (hebrewCursor ? hebrewError : monthError) : monthError
	const loading = view === 'atlas' ? indexLoading : view === 'calendar' && calendarSystem === 'hebrew' ? (hebrewCursor ? hebrewLoading : monthLoading) : monthLoading
  const skyData = useMemo(
    () => view === 'sky' ? buildSkyMonth(cursor.year, cursor.month, location) : null,
    [view, cursor.year, cursor.month, location],
  )
  const selectedDay = visibleData?.days.find((day) => day.date === selectedDate) || visibleData?.days[0]

	useEffect(() => {
	  if (!gregorianData || (view === 'calendar' && calendarSystem === 'hebrew')) return
	  if (gregorianData.days.some((day) => day.date === selectedDate)) return
	  const todayKey = dateKeyInTimeZone(new Date(), location.timezone)
	  const anchor = gregorianData.days.find((day) => day.date === todayKey) || gregorianData.days[0]
	  setSelectedDate(anchor.date)
	}, [gregorianData, view, calendarSystem, selectedDate, location.timezone])

  function moveMonth(delta) {
	if (view === 'atlas') {
	  setCursor((current) => {
		const year = current.year + delta
		return year < SKY_YEAR_MIN || year > SKY_YEAR_MAX ? current : { ...current, year }
	  })
	  return
	}
		if (view === 'calendar' && calendarSystem === 'hebrew') {
	  const target = hebrewMonthTarget(hebrewData, delta)
	  if (target) {
		setHebrewLoading(true)
		setHebrewError('')
		setHebrewCursor(target)
	  }
	  return
	}
	setMonthLoading(true)
    setCursor((current) => {
      const date = new Date(current.year, current.month - 1 + delta, 1)
	  const next = { year: date.getFullYear(), month: date.getMonth() + 1 }
	  if (next.year < SKY_YEAR_MIN || next.year > SKY_YEAR_MAX) return current
	  return next
    })
  }

  function goToday() {
	const todayKey = dateKeyInTimeZone(new Date(), location.timezone)
	setCursor(cursorForDateKey(todayKey))
	setSelectedDate(todayKey)
    setView('calendar')
	if (calendarSystem === 'hebrew') {
	  setHebrewLoading(true)
	  setHebrewData(null)
	  setHebrewCursor(null)
	}
  }

	function changeCalendarSystem(nextSystem) {
	  if (nextSystem === calendarSystem) return
	  if (nextSystem === 'hebrew') {
		const anchor = hebrewAnchorDay(gregorianData, selectedDate, SKY_YEAR_MIN, SKY_YEAR_MAX)
		if (!anchor) return
		setHebrewLoading(true)
		setHebrewData(null)
		setHebrewCursor({ year: anchor.sacredDates.hebrewYear, month: anchor.sacredDates.hebrewMonthNumber })
	  } else {
		setMonthLoading(true)
		setMonthData(null)
		setCursor(cursorForDateKey(selectedDate))
	  }
	  setCalendarSystem(nextSystem)
	}

  function changeView(nextView) {
    if (nextView !== 'calendar' && selectedDate) {
	  const viewDate = nextView === 'sky' ? clampSkyDateKey(selectedDate) : selectedDate
      setMonthLoading(true)
      setMonthData(null)
	  setCursor(cursorForDateKey(viewDate))
	  if (viewDate !== selectedDate) setSelectedDate(viewDate)
    }
    if (nextView === 'calendar' && calendarSystem === 'hebrew') {
	  const anchor = hebrewAnchorDay(gregorianData, selectedDate, SKY_YEAR_MIN, SKY_YEAR_MAX)
      if (anchor) {
		setHebrewLoading(true)
        setHebrewData(null)
        setHebrewCursor({ year: anchor.sacredDates.hebrewYear, month: anchor.sacredDates.hebrewMonthNumber })
	  } else {
		setHebrewLoading(true)
		setHebrewData(null)
		setHebrewCursor(null)
      }
    }
    setView(nextView)
  }

  function openAbout() {
    setDialog('about')
    if (!about) {
      setAboutError('')
      getAbout().then(setAbout).catch((reason) => setAboutError(reason.message))
    }
  }

  return (
    <div className="app-shell">
      <Header
        view={view}
        setView={changeView}
        location={location}
        onLocation={() => setDialog('location')}
        onAbout={openAbout}
      />

      <main>
        <Hero day={selectedDay} location={location} onToday={goToday} />

        <section className="workspace" aria-live="polite">
		  <Toolbar
			data={visibleData}
            cursor={cursor}
            filters={filters}
            setFilters={setFilters}
            moveMonth={moveMonth}
            view={view}
			calendarSystem={calendarSystem}
			onCalendarSystemChange={changeCalendarSystem}
			navigationLoading={loading}
          />

          {error && <ErrorState message={error} />}
		  {loading && !visibleData && view !== 'atlas' ? <CalendarSkeleton /> : null}

		  {!error && calendarData && view === 'calendar' && (
            <CalendarView
			  data={calendarData}
              selectedDate={selectedDate}
              setSelectedDate={setSelectedDate}
              filters={filters}
              selectedDay={selectedDay}
			  calendarSystem={calendarSystem}
            />
          )}

          {!error && gregorianData && view === 'readings' && <ReadingsView data={gregorianData} />}

          {!error && gregorianData && skyData && view === 'sky' && (
            <SkyView
              data={skyData}
              monthData={gregorianData}
              selectedDate={selectedDate}
              setSelectedDate={setSelectedDate}
              location={location}
            />
          )}

          {!error && view === 'atlas' && (
            <AtlasView data={indexData} filters={filters} year={cursor.year} />
          )}
        </section>
      </main>

      <Footer onAbout={openAbout} />

      {dialog === 'location' && (
        <LocationDialog
          location={location}
          onClose={() => setDialog(null)}
          onSave={(next) => {
			setMonthData(null)
			setHebrewData(null)
			setMonthLoading(true)
			if (view === 'calendar' && calendarSystem === 'hebrew') setHebrewLoading(true)
            setLocation(next)
            localStorage.setItem('lunarium-location', JSON.stringify(next))
            setDialog(null)
          }}
        />
      )}
		{dialog === 'about' && <AboutDialog about={about} error={aboutError} onClose={() => setDialog(null)} />}
    </div>
  )
}

function Header({ view, setView, location, onLocation, onAbout }) {
  return (
    <header className="topbar">
      <button className="brand" onClick={() => setView('calendar')} aria-label="Lunarium home">
        <span className="brand-mark" aria-hidden="true"><span /></span>
        <span>
          <strong>LUNARIUM</strong>
          <small>SACRED TIME, MAPPED</small>
        </span>
      </button>
      <nav className="primary-nav" aria-label="Primary navigation">
        <button className={view === 'calendar' ? 'active' : ''} aria-current={view === 'calendar' ? 'page' : undefined} onClick={() => setView('calendar')}>Calendar</button>
        <button className={view === 'sky' ? 'active' : ''} aria-current={view === 'sky' ? 'page' : undefined} onClick={() => setView('sky')}>Sky</button>
        <button className={view === 'atlas' ? 'active' : ''} aria-current={view === 'atlas' ? 'page' : undefined} onClick={() => setView('atlas')}>Observance atlas</button>
        <button className={view === 'readings' ? 'active' : ''} aria-current={view === 'readings' ? 'page' : undefined} onClick={() => setView('readings')}>Catholic readings</button>
      </nav>
      <div className="header-actions">
        <button className="location-button" onClick={onLocation} title="Set calculation location">
          <PinIcon /> <span>{location.name}</span>
        </button>
        <button className="round-button" onClick={onAbout} aria-label="About methodology">i</button>
      </div>
    </header>
  )
}

function Hero({ day, location, onToday }) {
  if (!day) return <section className="hero hero-empty" />
  const date = parseDate(day.date)
  return (
    <section className="hero">
      <div className="constellation constellation-one" aria-hidden="true" />
      <div className="constellation constellation-two" aria-hidden="true" />
      <div className="hero-copy">
        <p className="eyebrow">{formatDate(date, { weekday: 'long', month: 'long', day: 'numeric' }).toUpperCase()}</p>
        <h1>Many calendars.<br /><em>One sky.</em></h1>
        <p className="hero-intro">Explore living sacred days and carefully sourced historical calendars—held alongside lunar and planetary rhythms.</p>
        <div className="calendar-coordinates">
          <span><small>HEBREW</small>{day.sacredDates.hebrew}</span>
          <i aria-hidden="true" />
          <span><small>HIJRI · TABULAR</small>{day.sacredDates.islamic}</span>
        </div>
        <button className="text-button light" onClick={onToday}>Return to today <ArrowIcon /></button>
      </div>

      <div className="hero-moon" aria-label={`${day.moon.phase}, ${day.moon.illumination}% illuminated`}>
        <div className={`moon-disc ${day.moon.waxing ? 'waxing' : 'waning'}`} style={{ '--phase-angle': `${day.moon.angle}deg` }}>
          <span>{day.moon.emoji}</span>
        </div>
        <div className="orbit-line" aria-hidden="true"><b /></div>
      </div>

      <div className="hero-facts">
        <div className="hero-fact featured">
          <span className="fact-icon">{day.moon.emoji}</span>
          <p><small>MOON PHASE</small><strong>{day.moon.phase}</strong><span>{day.moon.illumination}% illuminated</span></p>
        </div>
        <div className="hero-fact">
          <span className="fact-icon">{day.astrology.symbol}</span>
          <p><small>TROPICAL SUN</small><strong>{day.astrology.sunSign}</strong><span>{day.astrology.element} · {day.astrology.mode}</span></p>
        </div>
        <div className="hero-fact">
          <span className="number-glyph">{day.numerology.score}</span>
          <p><small>UNIVERSAL DAY</small><strong>{day.numerology.title}</strong><span>{day.numerology.theme.split(' · ')[0]}</span></p>
        </div>
        <p className="hero-location"><PinIcon /> Calculated for {location.name}</p>
      </div>
    </section>
  )
}

function Toolbar({ data, cursor, filters, setFilters, moveMonth, view, calendarSystem, onCalendarSystemChange, navigationLoading }) {
  const title = view === 'atlas' ? `Observances of ${cursor.year}` : view === 'readings' ? `Reading plan · ${data?.label || ''}` : view === 'sky' ? `The sky in ${data?.label || ''}` : data?.label
	const range = view === 'calendar' && calendarSystem === 'hebrew' && data
	  ? `${formatDate(parseDate(data.startDate), { month: 'short', day: 'numeric' })}–${formatDate(parseDate(data.endDate), { month: 'short', day: 'numeric', year: 'numeric' })}`
		  : ''
	const { atSupportedStart, atSupportedEnd, waitingForHebrewMonth } = monthNavigationState({
	  view,
	  calendarSystem,
	  data,
	  cursor,
	  navigationLoading,
	  minYear: SKY_YEAR_MIN,
	  maxYear: SKY_YEAR_MAX,
	})
  return (
    <div className="toolbar">
      <div className="month-control">
		<button onClick={() => moveMonth(-1)} disabled={atSupportedStart || waitingForHebrewMonth} aria-label={view === 'atlas' ? 'Previous year' : 'Previous month'}><ChevronIcon direction="left" /></button>
        <div>
          <p className="eyebrow dark">{view === 'atlas' ? 'ANNUAL LIBRARY' : view === 'readings' ? 'CATHOLIC LECTIONARY' : view === 'sky' ? 'ASTRONOMICAL ALTERNATE VIEW' : 'MONTHLY CALENDAR'}</p>
          <h2>{title || 'Loading…'}</h2>
		  {range && <small className="period-range">{range} · Gregorian span</small>}
        </div>
		<button onClick={() => moveMonth(1)} disabled={atSupportedEnd || waitingForHebrewMonth} aria-label={view === 'atlas' ? 'Next year' : 'Next month'}><ChevronIcon /></button>
      </div>
		  <div className="toolbar-actions">
		{view === 'calendar' && (
		  <div className="calendar-system-switch" role="group" aria-label="Calendar view">
			<button className={calendarSystem === 'hebrew' ? 'active' : ''} onClick={() => onCalendarSystemChange('hebrew')} aria-pressed={calendarSystem === 'hebrew'}>Hebrew lunisolar</button>
			<button className={calendarSystem === 'gregorian' ? 'active' : ''} onClick={() => onCalendarSystemChange('gregorian')} aria-pressed={calendarSystem === 'gregorian'}>Gregorian alternate</button>
		  </div>
		)}
		{(view === 'calendar' || view === 'atlas') && <TraditionFilters filters={filters} setFilters={setFilters} />}
	  </div>
    </div>
  )
}

function TraditionFilters({ filters, setFilters }) {
  return (
    <div className="tradition-filters" role="group" aria-label="Filter traditions">
      {Object.entries(TRADITIONS).map(([key, meta]) => (
        <button
          key={key}
          className={`${key} ${filters[key] ? 'enabled' : ''}`}
          onClick={() => setFilters((current) => ({ ...current, [key]: !current[key] }))}
          aria-pressed={filters[key]}
        >
          <span>{meta.icon}</span>{meta.label}
        </button>
      ))}
    </div>
  )
}

function CalendarView({ data, selectedDate, setSelectedDate, filters, selectedDay, calendarSystem }) {
  const guideRef = useRef(null)
  function select(date) {
    setSelectedDate(date)
    if (window.innerWidth < 880) setTimeout(() => guideRef.current?.scrollIntoView({ behavior: 'smooth', block: 'start' }), 40)
  }
  return (
    <div className="calendar-layout">
      <div className="calendar-card">
        <div className="weekday-row">
          {WEEKDAYS.map((day) => <span key={day}>{day}</span>)}
        </div>
        <div className="calendar-grid">
          {Array.from({ length: data.firstWeekday }, (_, index) => <div className="empty-day" key={`empty-${index}`} />)}
          {data.days.map((day) => (
            <DayCell
              key={day.date}
              day={day}
              selected={selectedDate === day.date}
              filters={filters}
			  calendarSystem={calendarSystem}
              onSelect={() => select(day.date)}
            />
          ))}
        </div>
        <div className="calendar-footer">
          <p><MoonMini /> Phase is estimated at local noon · sacred dates begin at sunset where noted</p>
          <p>{data.observanceCount} distinct observances this month</p>
        </div>
      </div>
      <aside className="day-guide" ref={guideRef}>
		{selectedDay && <DayGuide day={selectedDay} filters={filters} />}
      </aside>
    </div>
  )
}

function DayCell({ day, selected, filters, onSelect, calendarSystem }) {
  const events = day.observances.filter((event) => filters[event.tradition])
	const hebrewMode = calendarSystem === 'hebrew'
	const secondaryDate = hebrewMode
	  ? formatDate(parseDate(day.date), { month: 'short', day: 'numeric' })
	  : `${day.sacredDates.hebrewDay} ${shortMonth(day.sacredDates.hebrewMonth)}`
  return (
	<button
	  className={`day-cell ${hebrewMode ? 'hebrew-view' : ''} ${selected ? 'selected' : ''} ${day.isToday ? 'today' : ''}`}
	  onClick={onSelect}
	  aria-label={`${formatDate(parseDate(day.date), { weekday: 'long', month: 'long', day: 'numeric', year: 'numeric' })}; ${day.sacredDates.hebrew}; ${events.length} ${events.length === 1 ? 'observance' : 'observances'}`}
	  aria-pressed={selected}
	>
      <span className="day-top">
		<b>{hebrewMode ? day.sacredDates.hebrewDay : day.day}</b>
        <span title={`${day.moon.phase} · ${day.moon.illumination}%`}>{day.moon.emoji}</span>
      </span>
	  <small className="hebrew-mini">{secondaryDate}</small>
      <span className="event-stack">
        {events.slice(0, 2).map((event) => (
          <span className={`event-chip ${event.tradition}`} key={event.id} title={event.name}>
            <i /> {event.name}
          </span>
        ))}
        {events.length > 2 && <span className="more-events">+{events.length - 2} more</span>}
      </span>
      <span className="day-score" title={`Numerology score ${day.numerology.score}`}>{day.numerology.score}</span>
      {day.isToday && <span className="today-label">TODAY</span>}
    </button>
  )
}

function DayGuide({ day, filters }) {
	const observances = day.observances.filter((event) => filters[event.tradition])
	const filteredDay = { ...day, observances }
	const [tab, setTab] = useState(observances.length ? 'observances' : 'prayer')
  useEffect(() => {
	setTab(observances.length ? 'observances' : 'prayer')
  }, [day.date])
  const date = parseDate(day.date)
	const tabs = [
	  ['observances', 'Sacred days'],
	  ['prayer', 'Prayer'],
	  ['lenses', 'Lenses'],
	  ['readings', 'Readings'],
	]
  function handleTabKeyDown(event, index) {
	let nextIndex
	if (event.key === 'ArrowRight') nextIndex = (index + 1) % tabs.length
	else if (event.key === 'ArrowLeft') nextIndex = (index - 1 + tabs.length) % tabs.length
	else if (event.key === 'Home') nextIndex = 0
	else if (event.key === 'End') nextIndex = tabs.length - 1
	else return
	event.preventDefault()
	const nextKey = tabs[nextIndex][0]
	setTab(nextKey)
	requestAnimationFrame(() => document.getElementById(`guide-tab-${nextKey}`)?.focus())
  }
  return (
    <div>
      <div className="guide-header">
        <p className="eyebrow dark">DAY GUIDE</p>
        <h3>{formatDate(date, { weekday: 'long' })}<br /><em>{formatDate(date, { month: 'long', day: 'numeric' })}</em></h3>
        <div className="sacred-date-lines">
          <span><i>א</i>{day.sacredDates.hebrew}</span>
          <span><i>ق</i>{day.sacredDates.islamic}</span>
        </div>
      </div>

      <div className="guide-tabs" role="tablist">
		{tabs.map(([key, label], index) => (
		  <button
			id={`guide-tab-${key}`}
			role="tab"
			aria-selected={tab === key}
			aria-controls={`guide-panel-${key}`}
			tabIndex={tab === key ? 0 : -1}
			className={tab === key ? 'active' : ''}
			onClick={() => setTab(key)}
			onKeyDown={(event) => handleTabKeyDown(event, index)}
			key={key}
		  >
			{label}{key === 'observances' && <span>{observances.length}</span>}
		  </button>
		))}
      </div>

	  <div className="guide-body" id={`guide-panel-${tab}`} role="tabpanel" aria-labelledby={`guide-tab-${tab}`}>
		{tab === 'observances' && <ObservancePanel day={filteredDay} />}
        {tab === 'prayer' && <PrayerPanel schedules={day.prayers} />}
        {tab === 'lenses' && <LensesPanel day={day} />}
        {tab === 'readings' && <ReadingCard reading={day.reading} expanded />}
      </div>
    </div>
  )
}

function ObservancePanel({ day }) {
  if (!day.observances.length) {
    return (
      <div className="quiet-day">
        <span>{day.moon.emoji}</span>
        <h4>A spacious day</h4>
        <p>No cataloged festival begins today. The weekly prayer rhythms and lunar cycle still offer a way to mark the day.</p>
      </div>
    )
  }
	const current = day.observances.filter((event) => !event.historical)
	const historical = day.observances.filter((event) => event.historical)
	return (
	  <div className="observance-groups">
		{current.length > 0 && <section><h4 className="observance-group-title">CURRENT & LIVING TRADITIONS</h4><div className="observance-list">{current.map((event) => <ObservanceCard event={event} key={`${event.id}-${event.dayIndex}`} />)}</div></section>}
		{historical.length > 0 && <section><h4 className="observance-group-title historical">HISTORICAL CALENDAR RECORDS</h4><p className="group-caveat">Preserved for study; these notices are not presented as modern observances.</p><div className="observance-list">{historical.map((event) => <ObservanceCard event={event} compact key={`${event.id}-${event.dayIndex}`} />)}</div></section>}
	  </div>
	)
}

function ObservanceCard({ event, compact = false }) {
  const meta = TRADITIONS[event.tradition] || { label: event.tradition, icon: '•' }
  return (
    <details className={`observance-card ${event.tradition}`} open={!compact}>
      <summary>
        <span className="tradition-symbol">{meta.icon}</span>
        <span className="event-heading">
          <small>{meta.label.toUpperCase()} · {event.category.toUpperCase()}</small>
          <strong>{event.name}</strong>
          <span>{event.communities.join(' · ')}</span>
		  {(event.calendarCorpus || event.origin || event.observanceStatus) && <span className="provenance-badges">{event.calendarCorpus && <i>{event.calendarCorpus}</i>}{event.origin && <i>{event.origin}</i>}{event.observanceStatus && <i className={event.historical ? 'historical' : ''}>{event.observanceStatus}</i>}</span>}
        </span>
        <ChevronIcon />
      </summary>
      <div className="event-content">
        {event.durationDays > 1 && <p className="range-note">Day {event.dayIndex} of {event.durationDays} · {formatRange(event.date, event.endDate)}</p>}
        <p>{event.summary}</p>
        <blockquote>{event.meaning}</blockquote>
        {event.practices?.length > 0 && (
          <div><small className="field-label">COMMON PRACTICES</small><ul>{event.practices.map((practice) => <li key={practice}>{practice}</li>)}</ul></div>
        )}
        {event.scripture?.length > 0 && <p className="scripture"><BookIcon /> {event.scripture.join(' · ')}</p>}
		{event.historicalNote && <p className="historical-note"><small className="field-label">HISTORICAL CONTEXT</small>{event.historicalNote}</p>}
		{event.nativeDateLabel && <p className="native-date-note"><small className="field-label">NATIVE CALENDAR DATE</small>{event.nativeDateLabel}</p>}
		{(event.era || event.site || event.attestationLayer) && <p className="corpus-line">{[event.era, event.site, event.attestationLayer].filter(Boolean).join(' · ')}</p>}
		{event.dateCertainty && <p className="caveat"><InfoIcon /> {event.dateCertainty}</p>}
        {event.dateNote && <p className="caveat"><InfoIcon /> {event.dateNote}</p>}
		<a href={event.sourceUrl} target="_blank" rel="noreferrer">{event.sourceName || 'Calendar source'} <ArrowIcon /></a>
      </div>
    </details>
  )
}

function PrayerPanel({ schedules }) {
  const [active, setActive] = useState('christianity')
  const schedule = schedules.find((item) => item.tradition === active) || schedules[0]
  return (
    <div className="prayer-panel">
      <div className="prayer-selector">
        {schedules.map((item) => (
          <button key={item.tradition} className={`${item.tradition} ${active === item.tradition ? 'active' : ''}`} onClick={() => setActive(item.tradition)}>
            {TRADITIONS[item.tradition].icon} {TRADITIONS[item.tradition].label}
          </button>
        ))}
      </div>
      <p className="schedule-name">{schedule.name}</p>
      <div className="time-list">
        {schedule.times.map((item) => (
          <div key={item.name}><span><strong>{item.name}</strong>{item.note && <small>{item.note}</small>}</span><time>{item.time}</time></div>
        ))}
      </div>
      <p className="method-note"><InfoIcon /> {schedule.method}</p>
      <p className="caveat">{schedule.caveat}</p>
    </div>
  )
}

function LensesPanel({ day }) {
  return (
    <div className="lens-list">
      <article className="lens-card moon-lens">
        <span className="lens-icon">{day.moon.emoji}</span>
        <div><small>MOON · ASTRONOMICAL ESTIMATE</small><h4>{day.moon.phase}</h4><p>{day.moon.ageDays} days old · {day.moon.illumination}% illuminated · {day.moon.waxing ? 'Waxing' : 'Waning'}</p></div>
        <p className="caveat">{day.moon.accuracyNote}</p>
      </article>
      <article className="lens-card kabbalah-lens">
        <span className="lens-icon">{day.kabbalah.letter.split(' ')[0]}</span>
        <div><small>KABBALAH · {day.kabbalah.month.toUpperCase()}</small><h4>{day.kabbalah.theme}</h4><p>{day.kabbalah.sign} · {day.kabbalah.tribe} · sense of {day.kabbalah.sense.toLowerCase()}</p></div>
        <blockquote>{day.kabbalah.practice}</blockquote>
        <p className="caveat">{day.kabbalah.caveat}</p>
      </article>
      <article className="lens-card astrology-lens">
        <span className="lens-icon">{day.astrology.symbol}</span>
        <div><small>ASTROLOGY · TROPICAL SUN</small><h4>{day.astrology.sunSign}</h4><p>{day.astrology.element} · {day.astrology.mode} · {day.astrology.theme}</p></div>
        <p className="caveat">{day.astrology.caveat}</p>
      </article>
      <article className="lens-card numerology-lens">
        <span className="lens-icon score">{day.numerology.score}</span>
        <div><small>NUMEROLOGY · UNIVERSAL DAY</small><h4>{day.numerology.title}</h4><p>{day.numerology.theme}</p></div>
        <blockquote>{day.numerology.prompt}</blockquote>
        <p className="formula">{day.numerology.method}</p>
        <p className="caveat">{day.numerology.caveat}</p>
      </article>
    </div>
  )
}

function SkyView({ data, monthData, selectedDate, setSelectedDate, location }) {
  const [eventFilter, setEventFilter] = useState('all')
  const effectiveDate = monthData.days.some((day) => day.date === selectedDate) ? selectedDate : monthData.days[0].date
  const snapshot = useMemo(() => buildSkySnapshot(effectiveDate, location), [effectiveDate, location])
  const filteredEvents = useMemo(() => data.events.filter((event) => skyEventGroup(event) === eventFilter || eventFilter === 'all'), [data, eventFilter])
	const selectedEvents = filteredEvents.filter((event) => event.date === effectiveDate)
	const calculationWarnings = [...(data.warnings || []), ...(snapshot.warnings || [])]
  const eventDays = useMemo(() => {
    const groups = new Map()
    filteredEvents.forEach((event) => groups.set(event.date, [...(groups.get(event.date) || []), event]))
    return [...groups.entries()]
  }, [filteredEvents])

  useEffect(() => {
    if (effectiveDate !== selectedDate) setSelectedDate(effectiveDate)
  }, [effectiveDate, selectedDate, setSelectedDate])

  return (
    <div className="sky-layout">
      <section className="sky-timeline">
        <header className="sky-intro">
          <div>
            <p className="eyebrow dark">EPHEMERIS · {data.year}</p>
            <h3>A month of<br /><em>moving light</em></h3>
          </div>
          <p>{data.methodology}</p>
          <div className="sky-stat-row">
            <span><strong>{data.events.length}</strong><small>calculated events</small></span>
            <span><strong>{data.eventDays.length}</strong><small>active dates</small></span>
            <span><strong>{snapshot.retrograde.length}</strong><small>retrograde tonight</small></span>
          </div>
        </header>
		<CalculationWarnings warnings={calculationWarnings} />

        <div className="sky-filter-row" aria-label="Filter astronomical events">
          {[
            ['all', 'All events'],
            ['essential', 'Phases & seasons'],
            ['motion', 'Motion & orbits'],
            ['alignment', 'Alignments & proximity'],
          ].map(([key, label]) => (
            <button key={key} className={eventFilter === key ? 'active' : ''} onClick={() => setEventFilter(key)} aria-pressed={eventFilter === key}>{label}</button>
          ))}
        </div>

        <div className="sky-event-days">
          {eventDays.map(([date, events]) => (
            <article className={`sky-day-group ${date === effectiveDate ? 'selected' : ''}`} key={date}>
              <button className="sky-day-heading" onClick={() => setSelectedDate(date)}>
                <span><small>{formatDate(parseDate(date), { weekday: 'short' }).toUpperCase()}</small><strong>{formatDate(parseDate(date), { month: 'short', day: 'numeric' })}</strong></span>
                <i>{events.length} {events.length === 1 ? 'event' : 'events'}</i>
                <ChevronIcon />
              </button>
              <div className="sky-event-stack">
                {events.map((event) => <SkyEventRow event={event} timeZone={location.timezone} key={event.id} />)}
              </div>
            </article>
          ))}
          {!eventDays.length && <div className="empty-search">No events match this sky filter.</div>}
        </div>
      </section>

      <aside className="sky-inspector">
        <header className="sky-inspector-header">
          <p className="eyebrow">EVENING SKY · {snapshot.skyState.toUpperCase()}</p>
          <h3>{formatDate(parseDate(effectiveDate), { weekday: 'long' })}<br /><em>{formatDate(parseDate(effectiveDate), { month: 'long', day: 'numeric' })}</em></h3>
          <p><PinIcon /> {location.name} · {snapshot.referenceLabel}</p>
        </header>

        <div className="sky-wheel-wrap">
          <SkyWheel snapshot={snapshot} />
          <div className="moon-snapshot">
            <span>☽</span><p><small>MOON</small><strong>{snapshot.moon.phase}</strong><i>{snapshot.moon.illumination}% · {snapshot.moon.distanceKm?.toLocaleString('en-US')} km</i></p>
          </div>
        </div>

        {selectedEvents.length > 0 && (
          <section className="tonight-events">
            <h4>EVENTS ON THIS DATE <span>{selectedEvents.length}</span></h4>
            {selectedEvents.map((event) => <SkyEventDetail event={event} timeZone={location.timezone} key={event.id} />)}
          </section>
		)}

		<SkyAlmanac almanac={snapshot.almanac} timeZone={location.timezone} />
        <SkyPositionTable snapshot={snapshot} />
        <HoroscopePanel snapshot={snapshot} />

        <p className="sky-method-note"><InfoIcon /> {snapshot.caveat} Calculations use <a href={data.sourceUrl} target="_blank" rel="noreferrer">{data.sourceName}</a>.</p>
      </aside>
    </div>
  )
}

function SkyEventRow({ event, timeZone }) {
  return (
    <div className={`sky-event-row ${categoryClass(event.category)}`}>
      <span className="sky-event-symbol">{event.symbol}</span>
      <span className="sky-event-copy">
        <small>{event.category.toUpperCase()} · {formatSkyTime(event.time, timeZone)}</small>
        <strong>{event.name}</strong>
        <p>{event.summary}</p>
      </span>
      {event.value && <b>{event.value}</b>}
    </div>
  )
}

function SkyEventDetail({ event, timeZone }) {
  return (
    <details className={`sky-event-detail ${categoryClass(event.category)}`}>
      <summary><span>{event.symbol}</span><p><small>{event.category.toUpperCase()} · {formatSkyTime(event.time, timeZone)}</small><strong>{event.name}</strong></p><ChevronIcon /></summary>
      <div>
        <p>{event.summary}</p>
		{event.globalPeakTime && event.localPeakTime && (
		  <p className="eclipse-circumstances">
			<strong>Selected location:</strong> {event.localKind} · {formatSkyTime(event.localPeakTime, timeZone)}<br />
			<strong>Worldwide event:</strong> {event.globalKind} · {formatSkyTime(event.globalPeakTime, timeZone)}
		  </p>
		)}
		{event.contacts?.length > 0 && (
		  <div className="sky-contact-list">
			<h5>CALCULATED CONTACTS{event.durationMinutes ? ` · ${formatDurationMinutes(event.durationMinutes)}` : ''}</h5>
			<ul>
			  {event.contacts.map((contact) => (
				<li key={contact.key}>
				  <span>{contact.label}</span>
				  <time dateTime={contact.time}>{formatSkyTime(contact.time, timeZone)}</time>
				  {Number.isFinite(contact.altitude) && <small>{Math.abs(contact.altitude)}° {contact.altitude >= 0 ? 'above' : 'below'} horizon</small>}
				</li>
			  ))}
			</ul>
		  </div>
		)}
        <p className="caveat">{event.details}</p>
        {event.bodies.length > 0 && <p className="sky-body-tags">{event.bodies.map((body) => <span key={body}>{body}</span>)}</p>}
        {event.local !== null && <p className={`visibility-note ${event.local ? 'visible' : ''}`}>{event.local ? '✓ Above the selected horizon at peak' : 'Local visibility is limited or absent at peak'}</p>}
        <a href={event.sourceUrl} target="_blank" rel="noreferrer">Calculation source <ArrowIcon /></a>
      </div>
    </details>
  )
}

function CalculationWarnings({ warnings }) {
	if (!warnings.length) return null
	const unique = warnings.filter((warning, index) => warnings.findIndex((item) => item.family === warning.family && item.message === warning.message) === index)
	return (
	  <aside className="sky-calculation-warning" role="status" aria-live="polite">
		<strong>Some calculation families were unavailable</strong>
		<ul>{unique.map((warning) => <li key={`${warning.family}-${warning.message}`}><b>{warning.family}:</b> {warning.message}</li>)}</ul>
	  </aside>
	)
}

function SkyAlmanac({ almanac, timeZone }) {
	const titleId = `sky-almanac-${almanac.date}`
	return (
	  <section className="sky-almanac" aria-labelledby={titleId}>
		<header>
		  <div><p className="eyebrow" id={titleId}>DAILY OBSERVING ALMANAC</p><h4>From first light to last</h4></div>
		  <span>{timeZone}</span>
		</header>
		<dl className="solar-almanac-grid">
		  {almanac.solar.map((event) => (
			<div key={event.key}>
			  <dt>{event.label}</dt>
			  <dd><AlmanacTime value={event.time} timeZone={timeZone} />{Number.isFinite(event.altitude) && <small>{event.altitude}° altitude</small>}</dd>
			</div>
		  ))}
		</dl>
		<div className="body-almanac-scroll" tabIndex="0" role="region" aria-label="Moon and planet rise, culmination, and set times">
		  <table className="body-almanac-table">
			<caption>Moon and planet observing times for {almanac.date}</caption>
			<thead><tr><th scope="col">Body</th><th scope="col">Rise</th><th scope="col">Culminate</th><th scope="col">Set</th></tr></thead>
			<tbody>
			  {almanac.bodies.map((body) => (
				<tr key={body.name}>
				  <th scope="row"><span>{body.symbol}</span>{body.name}</th>
				  <td><AlmanacTime value={body.rise} timeZone={timeZone} /></td>
				  <td><AlmanacTime value={body.culmination} timeZone={timeZone} />{Number.isFinite(body.culminationAltitude) && <small>{body.culminationAltitude}°</small>}</td>
				  <td><AlmanacTime value={body.set} timeZone={timeZone} /></td>
				</tr>
			  ))}
			</tbody>
		  </table>
		</div>
		<p className="almanac-caveat"><InfoIcon /> {almanac.caveat}</p>
	  </section>
	)
}

function AlmanacTime({ value, timeZone }) {
	if (!value) return <span className="no-crossing" aria-label="Does not occur on this civil date">—</span>
	return <time dateTime={value}>{formatSkyClock(value, timeZone)}</time>
}

function SkyWheel({ snapshot }) {
  const point = (longitude, radius = 112) => {
    const angle = (longitude - 90) * Math.PI / 180
    return { x: 150 + Math.cos(angle) * radius, y: 150 + Math.sin(angle) * radius }
  }
  const bodyPoints = Object.fromEntries(snapshot.positions.map((body) => [body.name, point(body.longitude)]))
  return (
    <svg className="sky-wheel" viewBox="0 0 300 300" role="img" aria-label="Tropical ecliptic positions at the selected time">
      <circle cx="150" cy="150" r="137" className="wheel-edge" />
      <circle cx="150" cy="150" r="112" className="wheel-orbit" />
      <circle cx="150" cy="150" r="56" className="wheel-core" />
      {ZODIAC.map((sign, index) => {
        const outer = point(index * 30, 137)
        const inner = point(index * 30, 112)
        const label = point(index * 30 + 15, 126)
        return <g key={sign.name}><line x1={inner.x} y1={inner.y} x2={outer.x} y2={outer.y} /><text x={label.x} y={label.y}>{sign.symbol}</text></g>
      })}
      {snapshot.aspects.slice(0, 7).map((aspect) => {
        const a = bodyPoints[aspect.a]
        const b = bodyPoints[aspect.b]
        return a && b ? <line className={`wheel-aspect aspect-${aspect.name}`} x1={a.x} y1={a.y} x2={b.x} y2={b.y} key={`${aspect.a}-${aspect.b}-${aspect.name}`} /> : null
      })}
      {snapshot.positions.map((body, index) => {
        const marker = point(body.longitude, 105 + (index % 2) * 14)
        return <g className={`wheel-body ${body.retrograde ? 'retrograde' : ''}`} transform={`translate(${marker.x} ${marker.y})`} key={body.name}><circle r="10" style={{ fill: body.color }} /><text y="1">{body.symbol}</text>{body.retrograde && <text className="wheel-rx" x="8" y="-8">℞</text>}<title>{body.name}: {body.degree}° {body.sign}</title></g>
      })}
      <text className="wheel-center-label" x="150" y="146">TROPICAL</text>
      <text className="wheel-center-value" x="150" y="163">ECLIPTIC</text>
    </svg>
  )
}

function SkyPositionTable({ snapshot }) {
  const [expanded, setExpanded] = useState(false)
  const positions = expanded ? snapshot.positions : snapshot.positions.slice(0, 7)
  return (
    <section className="position-panel">
      <div className="position-heading"><h4>BODY POSITIONS</h4><span>ALT / AZ</span></div>
      <div className="position-table">
        {positions.map((body) => (
          <div key={body.name}>
            <span className="position-symbol" style={{ '--body-color': body.color }}>{body.symbol}</span>
            <p><strong>{body.name}{body.retrograde && <i>℞</i>}</strong><small>{body.degree}° {body.sign} · {body.constellation}</small></p>
            <span className={body.aboveHorizon ? 'above' : ''}><strong>{body.altitude}°</strong><small>{body.direction} {body.azimuth}°</small></span>
          </div>
        ))}
      </div>
      <button className="position-more" onClick={() => setExpanded((value) => !value)}>{expanded ? 'Show primary bodies' : 'Include Uranus, Neptune & Pluto'} <ChevronIcon /></button>
      <p className="position-key"><i /> Above horizon at 21:00 local · altitude includes standard refraction</p>
    </section>
  )
}

function HoroscopePanel({ snapshot }) {
  const sunSign = snapshot.positions.find((body) => body.name === 'Sun')?.sign || 'Aries'
  const [sign, setSign] = useState(sunSign)
  const horoscope = buildHoroscope(snapshot, sign)
  return (
    <section className="horoscope-panel">
      <div className="horoscope-heading"><p className="eyebrow">SYMBOLIC LENS</p><h4>Choose a sun sign</h4></div>
      <div className="sign-picker" aria-label="Choose horoscope sun sign">
        {ZODIAC.map((item) => <button className={item.name === sign ? 'active' : ''} onClick={() => setSign(item.name)} title={item.name} aria-label={item.name} aria-pressed={item.name === sign} key={item.name}>{item.symbol}</button>)}
      </div>
      <article className="horoscope-card">
        <span>{horoscope.sign.symbol}</span>
        <div><small>{horoscope.theme.toUpperCase()}</small><h4>{horoscope.title}</h4></div>
        <blockquote>{horoscope.prompt}</blockquote>
        <p>{horoscope.rationale}</p>
        <p className="caveat"><InfoIcon /> {horoscope.caveat}</p>
      </article>
    </section>
  )
}

function skyEventGroup(event) {
  if (['lunar phase', 'season', 'eclipse'].includes(event.category)) return 'essential'
  if (['alignment', 'close approach'].includes(event.category)) return 'alignment'
  return 'motion'
}

function categoryClass(category) {
  return category.replaceAll(' ', '-')
}

function ReadingsView({ data }) {
  const grouped = useMemo(() => {
    const groups = []
    data.days.forEach((day) => {
      const last = groups.at(-1)
      if (last?.season === day.reading.season) last.days.push(day)
      else groups.push({ season: day.reading.season, days: [day] })
    })
    return groups
  }, [data])
  return (
    <div className="readings-layout">
      <aside className="reading-intro">
        <p className="eyebrow dark">LECTIO · {data.year}</p>
        <h3>A month in<br /><em>the Word</em></h3>
        <p>Follow the Roman Catholic daily Mass schedule. Every row opens the official USCCB readings for that exact date.</p>
        <div className="cycle-card">
          <span><small>SUNDAY CYCLE</small><strong>{data.days[0].reading.sundayCycle}</strong></span>
          <span><small>WEEKDAY CYCLE</small><strong>{data.days[0].reading.weekdayCycle}</strong></span>
        </div>
        <p className="caveat"><InfoIcon /> Optional memorials and diocesan calendars can alter the selection. The linked USCCB page is authoritative for the U.S. schedule.</p>
      </aside>
      <div className="reading-schedule">
        {grouped.map((group) => (
          <section key={`${group.season}-${group.days[0].date}`}>
            <h4><span />{group.season}</h4>
            {group.days.map((day) => <ReadingRow day={day} key={day.date} />)}
          </section>
        ))}
      </div>
    </div>
  )
}

function ReadingRow({ day }) {
  const date = parseDate(day.date)
  return (
    <a className={`reading-row ${day.isToday ? 'today' : ''}`} href={day.reading.sourceUrl} target="_blank" rel="noreferrer">
      <span className="reading-date"><small>{formatDate(date, { weekday: 'short' }).toUpperCase()}</small><strong>{day.day}</strong></span>
      <span className="reading-title"><strong>{day.reading.liturgicalDay}</strong><small>{day.observances.find((event) => event.tradition === 'christianity')?.category || 'Daily liturgy'}</small></span>
      <span className="reading-cycle">{date.getDay() === 0 ? `YEAR ${day.reading.sundayCycle}` : `CYCLE ${day.reading.weekdayCycle}`}</span>
      <ArrowIcon />
    </a>
  )
}

function ReadingCard({ reading }) {
  return (
    <article className="reading-card">
      <span className="reading-book"><BookIcon /></span>
      <p className="eyebrow dark">CATHOLIC MASS · {reading.season.toUpperCase()}</p>
      <h4>{reading.liturgicalDay}</h4>
      <div className="reading-cycles"><span>Sunday Year {reading.sundayCycle}</span><span>Weekday Cycle {reading.weekdayCycle}</span></div>
      <p>{reading.scheduleNote}</p>
      <a className="primary-button" href={reading.sourceUrl} target="_blank" rel="noreferrer">Open today’s readings <ArrowIcon /></a>
    </article>
  )
}

function AtlasView({ data, filters, year }) {
  const [query, setQuery] = useState('')
  const [scope, setScope] = useState(() => {
    const requested = new URLSearchParams(window.location.search).get('scope')
    return ['all', 'dated', 'native'].includes(requested) ? requested : 'all'
  })
  if (!data || data.year !== year) return <CalendarSkeleton />
	const nativeCount = data.observances.filter((event) => event.catalogOnly).length
	const normalizedQuery = query.trim().toLowerCase()
	const events = data.observances.filter((event) => {
	  const inScope = scope === 'all' || (scope === 'native' ? event.catalogOnly : !event.catalogOnly)
	  const searchable = `${event.name} ${event.summary} ${event.communities.join(' ')} ${event.origin || ''} ${event.observanceStatus || ''} ${event.historicalNote || ''} ${event.calendarCorpus || ''} ${event.nativeDateLabel || ''} ${event.era || ''} ${event.site || ''}`.toLowerCase()
	  return inScope && filters[event.tradition] && searchable.includes(normalizedQuery)
	})
  return (
    <div className="atlas-layout">
      <div className="atlas-summary">
        {Object.entries(TRADITIONS).map(([key, meta]) => (
          <div className={key} key={key}><span>{meta.icon}</span><strong>{data.counts[key]}</strong><small>{meta.label} entries</small></div>
        ))}
        <label className="search-box"><SearchIcon /><input value={query} onChange={(event) => setQuery(event.target.value)} placeholder="Search rites, communities, corpora, sites, or native dates" /></label>
        <nav className="atlas-scope" aria-label="Observance date scope">
          {[
            ['all', 'All records', data.observances.length],
            ['dated', 'Projected & living dates', data.observances.length - nativeCount],
            ['native', 'Native-date only', nativeCount],
          ].map(([key, label, count]) => <button className={scope === key ? 'active' : ''} onClick={() => setScope(key)} aria-pressed={scope === key} key={key}>{label}<span>{count}</span></button>)}
          <small>{events.length} shown</small>
        </nav>
      </div>
      <div className="atlas-grid">
        {events.map((event) => (
          <article className={`atlas-card ${event.tradition} ${event.catalogOnly ? 'catalog-only' : ''}`} key={event.id}>
            <div className="atlas-date">{event.catalogOnly ? <><strong>Native date</strong><small>not projected</small></> : <><strong>{formatDate(parseDate(event.date), { month: 'short', day: 'numeric' })}</strong>{event.endDate && <small>through {formatDate(parseDate(event.endDate), { month: 'short', day: 'numeric' })}</small>}</>}</div>
            <span className="tradition-symbol">{TRADITIONS[event.tradition].icon}</span>
            <small>{TRADITIONS[event.tradition].label.toUpperCase()} · {event.category.toUpperCase()}</small>
            <h3>{event.name}</h3>
			{(event.calendarCorpus || event.origin || event.observanceStatus) && <div className="atlas-provenance">{event.calendarCorpus && <span>{event.calendarCorpus}</span>}{event.origin && <span>{event.origin}</span>}{event.observanceStatus && <span className={event.historical ? 'historical' : ''}>{event.observanceStatus}</span>}</div>}
            <p>{event.summary}</p>
            <div className="community-tags">{event.communities.slice(0, 3).map((community) => <span key={community}>{community}</span>)}</div>
			<details>
			  <summary>Meaning, evidence & practice <ChevronIcon /></summary>
			  <blockquote>{event.meaning}</blockquote>
			  {event.nativeDateLabel && <p className="native-date-note"><small className="field-label">NATIVE DATE</small>{event.nativeDateLabel}</p>}
			  {(event.era || event.site || event.attestationLayer) && <p className="corpus-line">{[event.era, event.site, event.attestationLayer].filter(Boolean).join(' · ')}</p>}
			  {event.scripture?.length > 0 && <p className="scripture"><BookIcon /> {event.scripture.join(' · ')}</p>}
			  {event.practices?.length > 0 && <ul>{event.practices.map((practice) => <li key={practice}>{practice}</li>)}</ul>}
			  {event.historicalNote && <p>{event.historicalNote}</p>}
			  {event.projectionKind && <p className="caveat"><InfoIcon /> Projection: {event.projectionKind}</p>}
			  {event.projectionStatus && <p className="caveat"><InfoIcon /> Status: {event.projectionStatus}</p>}
			  {event.anchorLocation && <p className="caveat"><InfoIcon /> Anchor: {event.anchorLocation}</p>}
			  {event.dayBoundary && <p className="caveat"><InfoIcon /> Day boundary: {event.dayBoundary}</p>}
			  {event.dateCertainty && <p className="caveat">{event.dateCertainty}</p>}
			  {event.dateNote && <p className="caveat">{event.dateNote}</p>}
			  <a className="atlas-source" href={event.sourceUrl} target="_blank" rel="noreferrer">{event.sourceName}<ArrowIcon /></a>
			</details>
          </article>
        ))}
		{events.length === 0 && (
		  <div className="atlas-empty" role="status">
			<strong>No observances match this view.</strong>
			<span>Try another date scope, tradition filter, or search term.</span>
		  </div>
		)}
      </div>
      {!events.length && <div className="empty-search">No observances match this search and filter combination.</div>}
    </div>
  )
}

function LocationDialog({ location, onClose, onSave }) {
  const [draft, setDraft] = useState(location)
  const [locating, setLocating] = useState(false)
  const [message, setMessage] = useState('')
  function detectLocation() {
    if (!navigator.geolocation) {
      setMessage('Geolocation is not available in this browser.')
      return
    }
    setLocating(true)
    navigator.geolocation.getCurrentPosition(
      ({ coords }) => {
        setDraft({ name: 'Device location', latitude: roundCoordinate(coords.latitude), longitude: roundCoordinate(coords.longitude), timezone: Intl.DateTimeFormat().resolvedOptions().timeZone || 'UTC' })
        setLocating(false)
      },
      () => {
        setMessage('Location permission was not granted. You can enter coordinates manually.')
        setLocating(false)
      },
      { timeout: 10000 },
    )
  }
  function saveLocation() {
    try {
      onSave(normalizeLocation(draft))
    } catch (reason) {
      setMessage(reason.message)
    }
  }
  return (
    <Modal title="Calculation location" onClose={onClose}>
      <p>Coordinates and timezone set prayer windows, civil-day boundaries, and local sky circumstances. They are sent only to this app’s Go API for calendar calculations; astronomical calculations run locally in your browser, and the app does not forward your location to third parties.</p>
      <button className="detect-button" onClick={detectLocation} disabled={locating}><PinIcon /> {locating ? 'Finding location…' : 'Use device location'}</button>
      {message && <p className="form-message">{message}</p>}
      <div className="form-grid">
        <label>Name<input value={draft.name} onChange={(event) => setDraft({ ...draft, name: event.target.value })} /></label>
        <label>Timezone<input value={draft.timezone} onChange={(event) => setDraft({ ...draft, timezone: event.target.value })} placeholder="America/New_York" /></label>
        <label>Latitude<input type="number" min="-66" max="66" step="0.0001" value={draft.latitude} onChange={(event) => setDraft({ ...draft, latitude: event.target.value })} /></label>
        <label>Longitude<input type="number" min="-180" max="180" step="0.0001" value={draft.longitude} onChange={(event) => setDraft({ ...draft, longitude: event.target.value })} /></label>
      </div>
      <div className="dialog-actions"><button onClick={onClose}>Cancel</button><button className="primary-button" onClick={saveLocation}>Save location</button></div>
    </Modal>
  )
}

function AboutDialog({ about, error, onClose }) {
  return (
    <Modal title="About Lunarium" onClose={onClose} wide>
      {error ? <p className="form-message">{error}</p> : !about ? <p>Loading methodology…</p> : (
        <div className="about-content">
          <p className="about-lede">Lunarium is a learning calendar: it brings traditions into respectful proximity without pretending they are interchangeable.</p>
          <h3>How dates are made</h3>
          <ol>{about.methodology.map((item) => <li key={item}>{item}</li>)}</ol>
          <h3>Sources & validation</h3>
          <div className="source-list">{about.sources.map((source) => <a href={source.url} target="_blank" rel="noreferrer" key={source.name}><strong>{source.name}</strong><span>{source.use}</span><ArrowIcon /></a>)}</div>
          <h3>Read this with care</h3>
          {about.disclaimers.map((item) => <p className="caveat" key={item}><InfoIcon /> {item}</p>)}
        </div>
      )}
    </Modal>
  )
}

function Modal({ title, onClose, children, wide = false }) {
	const dialogRef = useRef(null)
  useEffect(() => {
	const previouslyFocused = document.activeElement
	const focusable = () => [...(dialogRef.current?.querySelectorAll('button:not([disabled]), a[href], input:not([disabled]), select:not([disabled]), textarea:not([disabled]), [tabindex]:not([tabindex="-1"])') || [])]
	requestAnimationFrame(() => focusable()[0]?.focus())
	const handler = (event) => {
	  if (event.key === 'Escape') {
		onClose()
		return
	  }
	  if (event.key !== 'Tab') return
	  const items = focusable()
	  if (!items.length) return
	  const first = items[0]
	  const last = items[items.length - 1]
	  if (event.shiftKey && document.activeElement === first) {
		event.preventDefault()
		last.focus()
	  } else if (!event.shiftKey && document.activeElement === last) {
		event.preventDefault()
		first.focus()
	  }
	}
    window.addEventListener('keydown', handler)
	return () => {
	  window.removeEventListener('keydown', handler)
	  if (previouslyFocused instanceof HTMLElement) previouslyFocused.focus()
	}
  }, [onClose])
  return (
    <div className="modal-backdrop" role="presentation" onMouseDown={(event) => event.target === event.currentTarget && onClose()}>
	  <section ref={dialogRef} className={`modal ${wide ? 'wide' : ''}`} role="dialog" aria-modal="true" aria-labelledby="lunarium-dialog-title">
		<div className="modal-header"><p className="eyebrow dark">LUNARIUM</p><h2 id="lunarium-dialog-title">{title}</h2><button onClick={onClose} aria-label="Close dialog">×</button></div>
        <div className="modal-body">{children}</div>
      </section>
    </div>
  )
}

function Footer({ onAbout }) {
  return (
    <footer>
      <span className="brand-mark small" aria-hidden="true"><span /></span>
      <p><strong>Lunarium</strong><br />A contemplative calendar for an interconnected world.</p>
      <button onClick={onAbout}>Methodology & sources</button>
      <p className="footer-note">Dates and prayer times are guides. Confirm observance with your local community.</p>
    </footer>
  )
}

function ErrorState({ message }) {
  return <div className="error-state"><strong>The calendar could not be loaded.</strong><span>{message}</span><small>Make sure the Go API is running on port 8080.</small></div>
}

function CalendarSkeleton() {
  return <div className="skeleton"><div /><div /><div /><div /><div /><div /></div>
}

function formatDate(date, options) {
  return new Intl.DateTimeFormat('en-US', options).format(date)
}

function formatSkyTime(value, timeZone) {
  return new Intl.DateTimeFormat('en-US', {
    timeZone: safeTimeZone(timeZone),
    hour: 'numeric',
    minute: '2-digit',
    timeZoneName: 'short',
  }).format(new Date(value))
}

function formatSkyClock(value, timeZone) {
	return new Intl.DateTimeFormat('en-US', {
	  timeZone: safeTimeZone(timeZone),
	  hour: 'numeric',
	  minute: '2-digit',
	}).format(new Date(value))
}

function formatDurationMinutes(value) {
	const minutes = Math.round(value)
	const hours = Math.floor(minutes / 60)
	const remainder = minutes % 60
	if (!hours) return `${remainder} min`
	return `${hours}h${remainder ? ` ${remainder}m` : ''}`
}

function parseDate(value) {
  return new Date(`${value}T12:00:00`)
}

function dateKeyInTimeZone(date, timeZone) {
	const parts = Object.fromEntries(new Intl.DateTimeFormat('en-US', {
	  timeZone: safeTimeZone(timeZone), year: 'numeric', month: '2-digit', day: '2-digit',
	}).formatToParts(date).map((part) => [part.type, part.value]))
	return `${parts.year}-${parts.month}-${parts.day}`
}

function cursorForDateKey(value) {
	const [year, month] = String(value).split('-').map(Number)
	return { year, month }
}

function clampSkyDateKey(value) {
	const [year, month, day] = String(value).split('-').map(Number)
	const targetYear = Math.min(SKY_YEAR_MAX, Math.max(SKY_YEAR_MIN, year))
	if (targetYear === year) return value
	const targetMonth = Math.min(12, Math.max(1, month || 1))
	const finalDay = Math.min(day || 1, new Date(targetYear, targetMonth, 0).getDate())
	return `${targetYear}-${String(targetMonth).padStart(2, '0')}-${String(finalDay).padStart(2, '0')}`
}

function formatRange(start, end) {
  if (!end) return formatDate(parseDate(start), { month: 'short', day: 'numeric' })
  return `${formatDate(parseDate(start), { month: 'short', day: 'numeric' })}–${formatDate(parseDate(end), { month: 'short', day: 'numeric' })}`
}

function shortMonth(value) {
  return value.length > 5 ? value.slice(0, 3) : value
}

function readLocation() {
  try {
    const saved = JSON.parse(localStorage.getItem('lunarium-location'))
    return normalizeLocation(saved)
  } catch {
    // Use a transparent, editable default.
  }
  return DEFAULT_LOCATION
}

function normalizeLocation(value) {
	if (value?.latitude === null || value?.latitude === undefined || String(value.latitude).trim() === '') throw new Error('Latitude is required.')
	if (value?.longitude === null || value?.longitude === undefined || String(value.longitude).trim() === '') throw new Error('Longitude is required.')
  const latitude = Number(value?.latitude)
  const longitude = Number(value?.longitude)
  const timezone = String(value?.timezone || '').trim()
  if (!Number.isFinite(latitude) || latitude < -66 || latitude > 66) throw new Error('Latitude must be between −66 and 66 degrees for the supported solar calculations.')
  if (!Number.isFinite(longitude) || longitude < -180 || longitude > 180) throw new Error('Longitude must be between −180 and 180 degrees.')
  if (!timezone || safeTimeZone(timezone) !== timezone) throw new Error('Timezone must be a valid IANA name such as America/New_York.')
  return {
    name: String(value?.name || '').trim() || 'Selected location',
    latitude: roundCoordinate(latitude),
    longitude: roundCoordinate(longitude),
    timezone,
  }
}

function safeTimeZone(value) {
  try {
    new Intl.DateTimeFormat('en-US', { timeZone: value }).format()
    return value
  } catch {
    return 'UTC'
  }
}

function roundCoordinate(value) {
  return Math.round(value * 10000) / 10000
}

function PinIcon() {
  return <svg viewBox="0 0 24 24" aria-hidden="true"><path d="M20 10c0 5-8 11-8 11S4 15 4 10a8 8 0 1 1 16 0Z" /><circle cx="12" cy="10" r="2.5" /></svg>
}
function ArrowIcon() {
  return <svg viewBox="0 0 24 24" aria-hidden="true"><path d="M5 12h14M14 7l5 5-5 5" /></svg>
}
function ChevronIcon({ direction = 'right' }) {
  return <svg className={direction === 'left' ? 'flip' : ''} viewBox="0 0 24 24" aria-hidden="true"><path d="m9 5 7 7-7 7" /></svg>
}
function InfoIcon() {
  return <svg viewBox="0 0 24 24" aria-hidden="true"><circle cx="12" cy="12" r="9" /><path d="M12 11v6M12 7.5v.2" /></svg>
}
function BookIcon() {
  return <svg viewBox="0 0 24 24" aria-hidden="true"><path d="M4 5.5A2.5 2.5 0 0 1 6.5 3H11v16H6.5A2.5 2.5 0 0 0 4 21.5v-16ZM20 5.5A2.5 2.5 0 0 0 17.5 3H13v16h4.5a2.5 2.5 0 0 1 2.5 2.5v-16Z" /></svg>
}
function SearchIcon() {
  return <svg viewBox="0 0 24 24" aria-hidden="true"><circle cx="11" cy="11" r="7" /><path d="m16 16 5 5" /></svg>
}
function MoonMini() {
  return <span className="moon-mini" aria-hidden="true">◐</span>
}

export default App
