package handlers

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWhatIsMyIPHandler(t *testing.T) {
	testCases := []struct {
		format     string
		expectedCT string
	}{
		{"json", "application/json"},
		{"text", "text/plain"},
		{"xml", "application/xml"},
		{"invalid", "application/json"},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("format=%s", tc.format), func(t *testing.T) {
			req, err := http.NewRequest("GET", fmt.Sprintf("/api/ip?format=%s", tc.format), nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(WhatIsMyIPHandler)
			handler.ServeHTTP(rr, req)

			assert.Equal(t, http.StatusOK, rr.Code)
			assert.Equal(t, tc.expectedCT, rr.Header().Get("Content-Type"))

			if tc.format != "text" {
				assert.NotEmpty(t, rr.Body.String())
			}
		})
	}
}
