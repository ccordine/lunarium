import { useEffect, useMemo, useRef, useState } from 'react'
import { getAbout, getHebrewMonth, getMonth, getObservances } from './api.js'

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
}

const WEEKDAYS = ['SUN', 'MON', 'TUE', 'WED', 'THU', 'FRI', 'SAT']

function App() {
  const now = new Date()
  const [cursor, setCursor] = useState({ year: now.getFullYear(), month: now.getMonth() + 1 })
	const [calendarSystem, setCalendarSystem] = useState('gregorian')
	const [hebrewCursor, setHebrewCursor] = useState(null)
  const [view, setView] = useState('calendar')
  const [monthData, setMonthData] = useState(null)
	const [hebrewData, setHebrewData] = useState(null)
  const [indexData, setIndexData] = useState(null)
  const [selectedDate, setSelectedDate] = useState(toISODate(now))
  const [filters, setFilters] = useState({ christianity: true, judaism: true, islam: true })
  const [location, setLocation] = useState(readLocation)
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState('')
  const [about, setAbout] = useState(null)
  const [dialog, setDialog] = useState(null)

  useEffect(() => {
    const controller = new AbortController()
    setLoading(true)
    setError('')
    getMonth(cursor.year, cursor.month, location, controller.signal)
      .then((payload) => {
        setMonthData(payload)
        const wanted = payload.days.find((day) => day.date === selectedDate)
        if (!wanted) {
          const today = payload.days.find((day) => day.isToday)
          setSelectedDate((today || payload.days[0]).date)
        }
      })
      .catch((reason) => {
        if (reason.name !== 'AbortError') setError(reason.message)
      })
      .finally(() => setLoading(false))
    return () => controller.abort()
  }, [cursor, location])

  useEffect(() => {
	if (calendarSystem !== 'hebrew' || !hebrewCursor) return
	const controller = new AbortController()
	setLoading(true)
	setError('')
	getHebrewMonth(hebrewCursor.year, hebrewCursor.month, location, controller.signal)
	  .then((payload) => {
		setHebrewData(payload)
		const wanted = payload.days.find((day) => day.date === selectedDate)
		if (!wanted) setSelectedDate(payload.days[0].date)
	  })
	  .catch((reason) => {
		if (reason.name !== 'AbortError') setError(reason.message)
	  })
	  .finally(() => setLoading(false))
	return () => controller.abort()
  }, [calendarSystem, hebrewCursor, location])

  useEffect(() => {
    if (view !== 'atlas') return
    const controller = new AbortController()
    getObservances(cursor.year, controller.signal)
      .then(setIndexData)
      .catch((reason) => {
        if (reason.name !== 'AbortError') setError(reason.message)
      })
    return () => controller.abort()
  }, [cursor.year, view])

	const calendarData = calendarSystem === 'hebrew' ? hebrewData : monthData
	const visibleData = view === 'calendar' ? calendarData : monthData
  const selectedDay = visibleData?.days.find((day) => day.date === selectedDate) || visibleData?.days[0]

  function moveMonth(delta) {
	if (view === 'calendar' && calendarSystem === 'hebrew') {
	  const target = delta < 0 ? hebrewData?.previous : hebrewData?.next
	  if (target) {
		setHebrewData(null)
		setHebrewCursor({ year: target.year, month: target.month })
	  }
	  return
	}
    setCursor((current) => {
      const date = new Date(current.year, current.month - 1 + delta, 1)
      return { year: date.getFullYear(), month: date.getMonth() + 1 }
    })
  }

  function goToday() {
    const today = new Date()
    setCursor({ year: today.getFullYear(), month: today.getMonth() + 1 })
    setSelectedDate(toISODate(today))
    setView('calendar')
	if (calendarSystem === 'hebrew') {
	  setLoading(true)
	  getMonth(today.getFullYear(), today.getMonth() + 1, location)
		.then((payload) => {
		  const day = payload.days.find((item) => item.date === toISODate(today))
		  if (day) {
			setHebrewData(null)
			setHebrewCursor({ year: day.sacredDates.hebrewYear, month: day.sacredDates.hebrewMonthNumber })
		  }
		})
		.catch((reason) => setError(reason.message))
		.finally(() => setLoading(false))
	}
  }

	function changeCalendarSystem(nextSystem) {
	  if (nextSystem === calendarSystem) return
	  if (nextSystem === 'hebrew') {
		const anchor = monthData?.days.find((day) => day.date === selectedDate) || monthData?.days[0]
		if (!anchor) return
		setHebrewData(null)
		setHebrewCursor({ year: anchor.sacredDates.hebrewYear, month: anchor.sacredDates.hebrewMonthNumber })
	  } else {
		const anchor = parseDate(selectedDate)
		setCursor({ year: anchor.getFullYear(), month: anchor.getMonth() + 1 })
	  }
	  setCalendarSystem(nextSystem)
	}

  function openAbout() {
    setDialog('about')
    if (!about) getAbout().then(setAbout).catch((reason) => setError(reason.message))
  }

  return (
    <div className="app-shell">
      <Header
        view={view}
        setView={setView}
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
            setCursor={setCursor}
            view={view}
			calendarSystem={calendarSystem}
			onCalendarSystemChange={changeCalendarSystem}
          />

          {error && <ErrorState message={error} />}
		  {loading && !visibleData ? <CalendarSkeleton /> : null}

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

          {!error && monthData && view === 'readings' && <ReadingsView data={monthData} />}

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
            setLocation(next)
            localStorage.setItem('lunarium-location', JSON.stringify(next))
            setDialog(null)
          }}
        />
      )}
      {dialog === 'about' && <AboutDialog about={about} onClose={() => setDialog(null)} />}
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
        <button className={view === 'calendar' ? 'active' : ''} onClick={() => setView('calendar')}>Calendar</button>
        <button className={view === 'atlas' ? 'active' : ''} onClick={() => setView('atlas')}>Observance atlas</button>
        <button className={view === 'readings' ? 'active' : ''} onClick={() => setView('readings')}>Catholic readings</button>
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
        <h1>Three calendars.<br /><em>One sky.</em></h1>
        <p className="hero-intro">Explore sacred days across Christian, Jewish, and Islamic traditions—held alongside the rhythms of the moon.</p>
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

function Toolbar({ data, cursor, filters, setFilters, moveMonth, setCursor, view, calendarSystem, onCalendarSystemChange }) {
  const title = view === 'atlas' ? `Observances of ${cursor.year}` : view === 'readings' ? `Reading plan · ${data?.label || ''}` : data?.label
	const range = view === 'calendar' && calendarSystem === 'hebrew' && data
	  ? `${formatDate(parseDate(data.startDate), { month: 'short', day: 'numeric' })}–${formatDate(parseDate(data.endDate), { month: 'short', day: 'numeric', year: 'numeric' })}`
	  : ''
  return (
    <div className="toolbar">
      <div className="month-control">
        <button onClick={() => moveMonth(-1)} aria-label="Previous month"><ChevronIcon direction="left" /></button>
        <div>
          <p className="eyebrow dark">{view === 'atlas' ? 'ANNUAL LIBRARY' : view === 'readings' ? 'CATHOLIC LECTIONARY' : 'MONTHLY CALENDAR'}</p>
          <h2>{title || 'Loading…'}</h2>
		  {range && <small className="period-range">{range} · Gregorian span</small>}
        </div>
        <button onClick={() => moveMonth(1)} aria-label="Next month"><ChevronIcon /></button>
      </div>
      {view === 'atlas' && (
        <div className="year-stepper">
          <button onClick={() => setCursor((c) => ({ ...c, year: c.year - 1 }))}>−</button>
          <span>{cursor.year}</span>
          <button onClick={() => setCursor((c) => ({ ...c, year: c.year + 1 }))}>+</button>
        </div>
      )}
	  <div className="toolbar-actions">
		{view === 'calendar' && (
		  <div className="calendar-system-switch" aria-label="Calendar view">
			<button className={calendarSystem === 'gregorian' ? 'active' : ''} onClick={() => onCalendarSystemChange('gregorian')} aria-pressed={calendarSystem === 'gregorian'}>Gregorian</button>
			<button className={calendarSystem === 'hebrew' ? 'active' : ''} onClick={() => onCalendarSystemChange('hebrew')} aria-pressed={calendarSystem === 'hebrew'}>Hebrew lunar</button>
		  </div>
		)}
		{view !== 'readings' && <TraditionFilters filters={filters} setFilters={setFilters} />}
	  </div>
    </div>
  )
}

function TraditionFilters({ filters, setFilters }) {
  return (
    <div className="tradition-filters" aria-label="Filter traditions">
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
	<button className={`day-cell ${hebrewMode ? 'hebrew-view' : ''} ${selected ? 'selected' : ''} ${day.isToday ? 'today' : ''}`} onClick={onSelect}>
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
		<button className={tab === 'observances' ? 'active' : ''} onClick={() => setTab('observances')}>Sacred days <span>{observances.length}</span></button>
        <button className={tab === 'prayer' ? 'active' : ''} onClick={() => setTab('prayer')}>Prayer</button>
        <button className={tab === 'lenses' ? 'active' : ''} onClick={() => setTab('lenses')}>Lenses</button>
        <button className={tab === 'readings' ? 'active' : ''} onClick={() => setTab('readings')}>Readings</button>
      </div>

      <div className="guide-body">
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
  const meta = TRADITIONS[event.tradition]
  return (
    <details className={`observance-card ${event.tradition}`} open={!compact}>
      <summary>
        <span className="tradition-symbol">{meta.icon}</span>
        <span className="event-heading">
          <small>{meta.label.toUpperCase()} · {event.category.toUpperCase()}</small>
          <strong>{event.name}</strong>
          <span>{event.communities.join(' · ')}</span>
		  {(event.origin || event.observanceStatus) && <span className="provenance-badges">{event.origin && <i>{event.origin}</i>}{event.observanceStatus && <i className={event.historical ? 'historical' : ''}>{event.observanceStatus}</i>}</span>}
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
  if (!data || data.year !== year) return <CalendarSkeleton />
	const events = data.observances.filter((event) => filters[event.tradition] && `${event.name} ${event.summary} ${event.communities.join(' ')} ${event.origin || ''} ${event.observanceStatus || ''} ${event.historicalNote || ''}`.toLowerCase().includes(query.toLowerCase()))
  return (
    <div className="atlas-layout">
      <div className="atlas-summary">
        {Object.entries(TRADITIONS).map(([key, meta]) => (
          <div className={key} key={key}><span>{meta.icon}</span><strong>{data.counts[key]}</strong><small>{meta.label} entries</small></div>
        ))}
        <label className="search-box"><SearchIcon /><input value={query} onChange={(event) => setQuery(event.target.value)} placeholder="Search observances or communities" /></label>
      </div>
      <div className="atlas-grid">
        {events.map((event) => (
          <article className={`atlas-card ${event.tradition}`} key={event.id}>
            <div className="atlas-date"><strong>{formatDate(parseDate(event.date), { month: 'short', day: 'numeric' })}</strong>{event.endDate && <small>through {formatDate(parseDate(event.endDate), { month: 'short', day: 'numeric' })}</small>}</div>
            <span className="tradition-symbol">{TRADITIONS[event.tradition].icon}</span>
            <small>{TRADITIONS[event.tradition].label.toUpperCase()} · {event.category.toUpperCase()}</small>
            <h3>{event.name}</h3>
			{(event.origin || event.observanceStatus) && <div className="atlas-provenance">{event.origin && <span>{event.origin}</span>}{event.observanceStatus && <span className={event.historical ? 'historical' : ''}>{event.observanceStatus}</span>}</div>}
            <p>{event.summary}</p>
            <div className="community-tags">{event.communities.slice(0, 3).map((community) => <span key={community}>{community}</span>)}</div>
			<details><summary>Meaning & practice <ChevronIcon /></summary><blockquote>{event.meaning}</blockquote>{event.practices?.length > 0 && <ul>{event.practices.map((practice) => <li key={practice}>{practice}</li>)}</ul>}{event.historicalNote && <p>{event.historicalNote}</p>}{event.dateNote && <p className="caveat">{event.dateNote}</p>}<a className="atlas-source" href={event.sourceUrl} target="_blank" rel="noreferrer">{event.sourceName}<ArrowIcon /></a></details>
          </article>
        ))}
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
  return (
    <Modal title="Prayer-time location" onClose={onClose}>
      <p>Coordinates and timezone are used only to calculate solar prayer windows in your browser session. They are sent to this app’s Go API and never forwarded elsewhere.</p>
      <button className="detect-button" onClick={detectLocation} disabled={locating}><PinIcon /> {locating ? 'Finding location…' : 'Use device location'}</button>
      {message && <p className="form-message">{message}</p>}
      <div className="form-grid">
        <label>Name<input value={draft.name} onChange={(event) => setDraft({ ...draft, name: event.target.value })} /></label>
        <label>Timezone<input value={draft.timezone} onChange={(event) => setDraft({ ...draft, timezone: event.target.value })} placeholder="America/New_York" /></label>
        <label>Latitude<input type="number" min="-66" max="66" step="0.0001" value={draft.latitude} onChange={(event) => setDraft({ ...draft, latitude: Number(event.target.value) })} /></label>
        <label>Longitude<input type="number" min="-180" max="180" step="0.0001" value={draft.longitude} onChange={(event) => setDraft({ ...draft, longitude: Number(event.target.value) })} /></label>
      </div>
      <div className="dialog-actions"><button onClick={onClose}>Cancel</button><button className="primary-button" onClick={() => onSave(draft)}>Save location</button></div>
    </Modal>
  )
}

function AboutDialog({ about, onClose }) {
  return (
    <Modal title="About Lunarium" onClose={onClose} wide>
      {!about ? <p>Loading methodology…</p> : (
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
  useEffect(() => {
    const handler = (event) => event.key === 'Escape' && onClose()
    window.addEventListener('keydown', handler)
    return () => window.removeEventListener('keydown', handler)
  }, [onClose])
  return (
    <div className="modal-backdrop" role="presentation" onMouseDown={(event) => event.target === event.currentTarget && onClose()}>
      <section className={`modal ${wide ? 'wide' : ''}`} role="dialog" aria-modal="true" aria-label={title}>
        <div className="modal-header"><p className="eyebrow dark">LUNARIUM</p><h2>{title}</h2><button onClick={onClose} aria-label="Close dialog">×</button></div>
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

function parseDate(value) {
  return new Date(`${value}T12:00:00`)
}

function toISODate(date) {
  const year = date.getFullYear()
  const month = String(date.getMonth() + 1).padStart(2, '0')
  const day = String(date.getDate()).padStart(2, '0')
  return `${year}-${month}-${day}`
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
    if (saved?.timezone && Number.isFinite(saved?.latitude) && Number.isFinite(saved?.longitude)) return saved
  } catch {
    // Use a transparent, editable default.
  }
  return DEFAULT_LOCATION
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
