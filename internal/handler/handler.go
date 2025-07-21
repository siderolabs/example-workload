// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// Package handler provides HTTP handlers for the example workload.
package handler

import (
	_ "embed"
	"log"
	"net/http"
	"strings"
)

//go:embed index.html
var htmlContent []byte

//go:embed index.txt
var textContent []byte

// Example serves different content based on the User-Agent header.
func Example(w http.ResponseWriter, r *http.Request) {
	userAgent := strings.ToLower(r.Header.Get("User-Agent"))

	// Check for terminal clients (curl, wget, HTTPie)
	isTerminalClient := strings.Contains(userAgent, "curl") ||
		strings.Contains(userAgent, "wget") ||
		strings.Contains(userAgent, "httpie")

	if isTerminalClient {
		// Serve plain text response for terminal clients
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")

		w.WriteHeader(http.StatusOK)
		w.Write(textContent) //nolint:errcheck
	} else {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		w.WriteHeader(http.StatusOK)
		w.Write(htmlContent) //nolint:errcheck
	}

	// Log request
	log.Printf("%s %s %s - %s", r.RemoteAddr, r.Method, r.URL.Path, userAgent)
}
