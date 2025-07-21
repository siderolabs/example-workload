// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package handler_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/siderolabs/example-workload/internal/handler"
)

// TestExample tests the Example handler for different User-Agent headers.
func TestExample(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		userAgent string
		wantText  bool
		wantHTML  bool
	}{
		{
			name:      "curl client",
			userAgent: "curl/7.64.1",
			wantText:  true,
			wantHTML:  false,
		},
		{
			name:      "wget client",
			userAgent: "Wget/1.20.3",
			wantText:  true,
			wantHTML:  false,
		},
		{
			name:      "HTTPie client",
			userAgent: "HTTPie/2.6.0",
			wantText:  true,
			wantHTML:  false,
		},
		{
			name:      "browser client",
			userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3",
			wantText:  false,
			wantHTML:  true,
		},
	}

	srv := httptest.NewServer(http.HandlerFunc(handler.Example))
	t.Cleanup(srv.Close)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			req, err := http.NewRequestWithContext(t.Context(), http.MethodGet, srv.URL, nil)
			require.NoError(t, err)

			req.Header.Set("User-Agent", tt.userAgent)

			resp, err := http.DefaultClient.Do(req)
			require.NoError(t, err)

			t.Cleanup(func() {
				require.NoError(t, resp.Body.Close())
			})

			assert.Equal(t, http.StatusOK, resp.StatusCode)

			body, err := io.ReadAll(resp.Body)
			require.NoError(t, err)

			switch {
			case tt.wantText:
				assert.Equal(t, "text/plain; charset=utf-8", resp.Header.Get("Content-Type"))
				assert.Contains(t, string(body), "🎉 CONGRATULATIONS! 🎉")
			case tt.wantHTML:
				assert.Equal(t, "text/html; charset=utf-8", resp.Header.Get("Content-Type"))
				assert.Contains(t, string(body), "<html")
			default:
				t.Fatal("unexpected test case: neither wantText nor wantHTML is true")
			}
		})
	}
}
