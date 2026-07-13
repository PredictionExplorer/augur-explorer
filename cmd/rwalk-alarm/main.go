// Command rwalk-alarm monitors a list of URLs and notifies people over
// WhatsApp when a URL keeps failing.
//
// Usage: rwalk-alarm [url_list_file]
//
// The URL list file has one "URL<TAB>Message header" entry per line.
// Configuration (environment):
//
//	PHONE_LIST         comma-separated person:phone records to notify
//	WHATSAPP_TOKEN     WhatsApp Cloud API token
//	WHATSAPP_PHONE_ID  WhatsApp Cloud API phone number id
//
// The monitoring engine lives in internal/notify/urlalarm; this binary only
// parses configuration and wires the production WhatsApp client.
package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/PredictionExplorer/augur-explorer/internal/notify/urlalarm"
	"github.com/PredictionExplorer/augur-explorer/internal/notify/wanotif"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	if err := run(ctx, os.Args[1:], os.Getenv, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "rwalk-alarm: %v\n", err)
		os.Exit(1)
	}
}

// run parses configuration and drives the engine until ctx is cancelled.
func run(ctx context.Context, args []string, getenv func(string) string, out io.Writer) error {
	if len(args) != 1 {
		return fmt.Errorf("usage: rwalk-alarm [url_list_file]")
	}

	logger := log.New(out, "INFO: ", log.Ldate|log.Ltime)

	data, err := os.ReadFile(filepath.Clean(args[0]))
	if err != nil {
		return fmt.Errorf("reading url list: %w", err)
	}
	urls, err := urlalarm.ParseURLList(data)
	if err != nil {
		return fmt.Errorf("parsing url list: %w", err)
	}
	logger.Printf("Loaded %v URLs", len(urls))

	people, err := urlalarm.ParsePhoneList(getenv("PHONE_LIST"))
	if err != nil {
		return fmt.Errorf("parsing PHONE_LIST: %w", err)
	}
	logger.Printf("Loaded %v phones for notification", len(people))

	token := getenv("WHATSAPP_TOKEN")
	if token == "" {
		return fmt.Errorf("WHATSAPP_TOKEN environment variable is empty")
	}
	phoneID := getenv("WHATSAPP_PHONE_ID")
	if phoneID == "" {
		return fmt.Errorf("WHATSAPP_PHONE_ID environment variable is empty")
	}

	engine := urlalarm.New(urlalarm.Config{
		URLs:   urls,
		People: people,
	}, wanotif.NewWhatsapp(token, phoneID), nil, logger)

	if err := engine.Run(ctx); !errors.Is(err, context.Canceled) {
		return err
	}
	logger.Print("Exiting upon user request")
	return nil
}
