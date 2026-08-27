package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"lunarcalendar/internal/calendar"
)

func TestHebrewCalendarHandler(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/api/v1/calendar/hebrew?year=5786&month=6&timezone=America%2FNew_York", nil)
	recorder := httptest.NewRecorder()

	hebrewCalendarHandler(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Fatalf("status = %d, body = %s", recorder.Code, recorder.Body.String())
	}
	if contentType := recorder.Header().Get("Content-Type"); !strings.HasPrefix(contentType, "application/json") {
		t.Fatalf("content type = %q", contentType)
	}
	var response calendar.MonthResponse
	if err := json.NewDecoder(recorder.Body).Decode(&response); err != nil {
		t.Fatal(err)
	}
	if response.CalendarSystem != calendar.HebrewCalendar || response.Label != "Elul 5786" {
		t.Fatalf("response period = %#v", response)
	}
	if response.StartDate != "2026-08-14" || response.EndDate != "2026-09-11" || len(response.Days) != 29 {
		t.Fatalf("response range = %s through %s with %d days", response.StartDate, response.EndDate, len(response.Days))
	}
	if response.Next != (calendar.MonthReference{Year: 5787, Month: 7, Label: "Tishrei 5787"}) {
		t.Fatalf("next = %#v", response.Next)
	}
}

func TestHebrewCalendarHandlerValidation(t *testing.T) {
	tests := []struct {
		name  string
		query string
		want  string
	}{
		{name: "missing year", query: "month=6", want: "year must be a Hebrew year"},
		{name: "bad month", query: "year=5786&month=14", want: "month must be between 1 and 13"},
		{name: "Adar II in common year", query: "year=5785&month=13", want: "month must be between 1 and 12 for Hebrew year 5785"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodGet, "/api/v1/calendar/hebrew?"+test.query, nil)
			recorder := httptest.NewRecorder()

			hebrewCalendarHandler(recorder, request)

			if recorder.Code != http.StatusBadRequest {
				t.Fatalf("status = %d, body = %s", recorder.Code, recorder.Body.String())
			}
			var response map[string]string
			if err := json.NewDecoder(recorder.Body).Decode(&response); err != nil {
				t.Fatal(err)
			}
			if !strings.Contains(response["error"], test.want) {
				t.Fatalf("error = %q, want fragment %q", response["error"], test.want)
			}
		})
	}
}
