package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
	_ "time/tzdata"

	"lunarcalendar/internal/calendar"
)

func main() {
	addr := flag.String("addr", ":8080", "HTTP listen address")
	assets := flag.String("assets", "web/dist", "path to the built React application")
	flag.Parse()

	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/v1/health", func(w http.ResponseWriter, _ *http.Request) {
		writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
	})
	mux.HandleFunc("GET /api/v1/calendar", func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		year, err := intQuery(r, "year", now.Year())
		if err != nil || year < 1900 || year > 2100 {
			writeError(w, http.StatusBadRequest, "year must be between 1900 and 2100")
			return
		}
		month, err := intQuery(r, "month", int(now.Month()))
		if err != nil || month < 1 || month > 12 {
			writeError(w, http.StatusBadRequest, "month must be between 1 and 12")
			return
		}
		location, err := locationQuery(r)
		if err != nil {
			writeError(w, http.StatusBadRequest, err.Error())
			return
		}
		writeJSON(w, http.StatusOK, calendar.BuildMonth(year, time.Month(month), location))
	})
	mux.HandleFunc("GET /api/v1/calendar/hebrew", hebrewCalendarHandler)
	mux.HandleFunc("GET /api/v1/observances", func(w http.ResponseWriter, r *http.Request) {
		year, err := intQuery(r, "year", time.Now().Year())
		if err != nil || year < 1900 || year > 2100 {
			writeError(w, http.StatusBadRequest, "year must be between 1900 and 2100")
			return
		}
		writeJSON(w, http.StatusOK, calendar.BuildObservanceIndex(year))
	})
	mux.HandleFunc("GET /api/v1/about", func(w http.ResponseWriter, _ *http.Request) {
		writeJSON(w, http.StatusOK, calendar.About())
	})

	if info, err := os.Stat(*assets); err == nil && info.IsDir() {
		mux.Handle("/", spaHandler(*assets))
	} else {
		mux.HandleFunc("GET /", func(w http.ResponseWriter, _ *http.Request) {
			w.Header().Set("Content-Type", "text/plain; charset=utf-8")
			w.WriteHeader(http.StatusServiceUnavailable)
			_, _ = w.Write([]byte("React build not found. Run `npm install && npm run build` in web/, or use `npm run dev`.\n"))
		})
	}

	server := &http.Server{
		Addr:              *addr,
		Handler:           securityHeaders(mux),
		ReadHeaderTimeout: 5 * time.Second,
		IdleTimeout:       60 * time.Second,
	}
	log.Printf("Lunarium listening on %s", *addr)
	log.Fatal(server.ListenAndServe())
}

func hebrewCalendarHandler(w http.ResponseWriter, r *http.Request) {
	year, err := intQuery(r, "year", 0)
	if err != nil || year < 1 || year > 9999 {
		writeError(w, http.StatusBadRequest, "year must be a Hebrew year between 1 and 9999")
		return
	}
	month, err := intQuery(r, "month", 0)
	if err != nil || month < 1 || month > 13 {
		writeError(w, http.StatusBadRequest, "month must be between 1 and 13")
		return
	}
	location, err := locationQuery(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	response, err := calendar.BuildHebrewMonth(year, month, location)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, response)
}

func intQuery(r *http.Request, key string, fallback int) (int, error) {
	raw := r.URL.Query().Get(key)
	if raw == "" {
		return fallback, nil
	}
	return strconv.Atoi(raw)
}

func locationQuery(r *http.Request) (calendar.Location, error) {
	location := calendar.DefaultLocation
	query := r.URL.Query()
	if value := query.Get("latitude"); value != "" {
		parsed, err := strconv.ParseFloat(value, 64)
		if err != nil || parsed < -66 || parsed > 66 {
			return location, fmt.Errorf("latitude must be a number between -66 and 66")
		}
		location.Latitude = parsed
	}
	if value := query.Get("longitude"); value != "" {
		parsed, err := strconv.ParseFloat(value, 64)
		if err != nil || parsed < -180 || parsed > 180 {
			return location, fmt.Errorf("longitude must be a number between -180 and 180")
		}
		location.Longitude = parsed
	}
	if value := query.Get("timezone"); value != "" {
		if _, err := time.LoadLocation(value); err != nil {
			return location, fmt.Errorf("timezone must be a valid IANA timezone")
		}
		location.Timezone = value
	}
	if value := strings.TrimSpace(query.Get("locationName")); value != "" {
		location.Name = value
	}
	return location, nil
}

func writeJSON(w http.ResponseWriter, status int, value any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(value)
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]string{"error": message})
}

func securityHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("Referrer-Policy", "strict-origin-when-cross-origin")
		next.ServeHTTP(w, r)
	})
}

func spaHandler(root string) http.Handler {
	files := http.FileServer(http.Dir(root))
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		clean := filepath.Clean(strings.TrimPrefix(r.URL.Path, "/"))
		if clean == "." {
			clean = "index.html"
		}
		if _, err := os.Stat(filepath.Join(root, clean)); err != nil {
			r.URL.Path = "/"
		}
		files.ServeHTTP(w, r)
	})
}
