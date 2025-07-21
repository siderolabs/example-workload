// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package main

import (
	"context"
	_ "embed"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"golang.org/x/sync/errgroup"
	"golang.org/x/sys/unix"

	"github.com/siderolabs/example-workload/internal/handler"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}

func run() error {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, unix.SIGTERM)
	defer cancel()

	eg, ctx := errgroup.WithContext(ctx)

	server := &http.Server{
		Addr:    ":8080",
		Handler: http.HandlerFunc(handler.Example),
	}

	log.Printf("Starting server on %s", server.Addr)

	eg.Go(func() error {
		// Start the HTTP server
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			return err
		}

		return nil
	})

	eg.Go(func() error {
		<-ctx.Done()

		gracefulCtx, gracefulCancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer gracefulCancel()

		return server.Shutdown(gracefulCtx) //nolint:contextcheck
	})

	return eg.Wait()
}
