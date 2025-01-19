package main

import (
	"net/http"
	"testing"
)

func TestFormatDate(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"2025-01-15", "15 January 2025"},
		{"2024-12-01", "1 December 2024"},
		{"invalid-date", "invalid-date"},
	}

	for _, test := range tests {
		result := FormatDate(test.input)
		if result != test.expected {
			t.Errorf("FormatDate(%q) = %q; want %q", test.input, result, test.expected)
		}
	}
}

func TestGetCountry(t *testing.T) {
	countries := map[string]string{
		"1": "France",
		"2": "Netherlands",
		"3": "Ireland",
	}

	for key, expected := range countries {
		result := GetCountry(key)
		if result != expected {
			t.Errorf("GetCountry(%q) = %q; want %q", key, result, expected)
		}
	}
}

func TestCheckAppointments_NoData(t *testing.T) {
	mockResponse := `[]`
	server := startMockServer(mockResponse)
	defer func(server *http.Server) {
		err := server.Close()
		if err != nil {
			return
		}
	}(server)

	CheckAppointments("France", "Istanbul")
}

func startMockServer(response string) *http.Server {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(response))
		if err != nil {
			return
		}
	})
	server := &http.Server{Addr: ":8081", Handler: handler}
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			return
		}
	}()
	return server
}

func TestCheckAppointments_WithAppointments(t *testing.T) {
	mockResponse := `[{
		"mission_country": "France",
		"appointment_date": "2025-01-15",
		"source_country": "Turkiye",
		"center_name": "VFS Global - Istanbul",
		"visa_category": "Tourist Visa",
		"book_now_link": "https://example.com/appointment"
	}]`

	server := startMockServer(mockResponse)
	defer func(server *http.Server) {
		err := server.Close()
		if err != nil {
			return
		}
	}(server)

	CheckAppointments("France", "Istanbul")
}
