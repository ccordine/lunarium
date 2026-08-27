export async function getMonth(year, month, location, signal) {
  const query = new URLSearchParams({
    year: String(year),
    month: String(month),
    latitude: String(location.latitude),
    longitude: String(location.longitude),
    timezone: location.timezone,
    locationName: location.name,
  })
  return getJSON(`/api/v1/calendar?${query}`, signal)
}

export async function getHebrewMonth(year, month, location, signal) {
  const query = new URLSearchParams({
    year: String(year),
    month: String(month),
    latitude: String(location.latitude),
    longitude: String(location.longitude),
    timezone: location.timezone,
    locationName: location.name,
  })
  return getJSON(`/api/v1/calendar/hebrew?${query}`, signal)
}

export async function getObservances(year, signal) {
  return getJSON(`/api/v1/observances?year=${year}`, signal)
}

export async function getAbout(signal) {
  return getJSON('/api/v1/about', signal)
}

async function getJSON(url, signal) {
  const response = await fetch(url, { signal })
  if (!response.ok) {
    let message = `Request failed (${response.status})`
    try {
      const body = await response.json()
      message = body.error || message
    } catch {
      // Keep the HTTP status fallback.
    }
    throw new Error(message)
  }
  return response.json()
}
