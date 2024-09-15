// Command lohi parses a Google location-history file and prints a human-readable summary.
package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/bobg/errors"

	"github.com/bobg/lohi/places"
	"github.com/bobg/lohi/schema"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}
}

func run() error {
	var (
		credsFile   string
		doTrips     bool
		placesCache string
		placesQPS   int
		tokenFile   string
	)

	flag.StringVar(&credsFile, "creds", "creds.json", "Google API credentials file")
	flag.BoolVar(&doTrips, "trips", false, "show trips instead of places")
	flag.StringVar(&placesCache, "cache", "places.db", "places cache file")
	flag.IntVar(&placesQPS, "qps", 1, "places API queries per second")
	flag.StringVar(&tokenFile, "token", "token.json", "OAuth token file")
	flag.Parse()

	ctx := context.Background()

	creds, err := os.ReadFile(credsFile)
	if err != nil {
		return errors.Wrapf(err, "reading %s", credsFile)
	}

	var placeService places.Service
	placeService, err = places.NewRealService(ctx, creds, tokenFile)
	if err != nil {
		return errors.Wrap(err, "creating real place service")
	}
	placeService = places.NewRateLimitedService(1, placeService)
	placeService, err = places.NewCachingService(placesCache, placeService)
	if err != nil {
		return errors.Wrap(err, "creating caching place service")
	}

	var (
		dec = json.NewDecoder(os.Stdin)
		h   *schema.History
	)
	if err := dec.Decode(&h); err != nil {
		return errors.Wrap(err, "decoding history")
	}

	var prevDayNum int

	for _, seg := range h.SemanticSegments {
		if seg == nil {
			continue
		}

		var (
			activity = seg.Activity
			trip     *schema.Trip
			visit    = seg.Visit
		)

		if doTrips {
			if seg.TimelineMemory == nil {
				continue
			}
			trip = seg.TimelineMemory.Trip
			if trip == nil {
				continue
			}
		} else if activity == nil && visit == nil {
			continue
		}

		startTime, err := time.Parse(time.RFC3339, seg.StartTime)
		if err != nil {
			return errors.Wrapf(err, "parsing start time %q", seg.StartTime)
		}
		startLoc := time.FixedZone("start", int(seg.StartTimeTimezoneUtcOffsetMinutes)*60)
		startTime = startTime.In(startLoc)

		endTime, err := time.Parse(time.RFC3339, seg.EndTime)
		if err != nil {
			return errors.Wrapf(err, "parsing end time %q", seg.EndTime)
		}

		dur := endTime.Sub(startTime)

		if dayNum := startTime.Year()*1000 + startTime.YearDay(); dayNum != prevDayNum {
			fmt.Printf("\n%s\n\n", startTime.Format("2006-01-02"))
			prevDayNum = dayNum
		}

		if doTrips {
			dur = dur.Round(24 * time.Hour)
			days := int(dur.Hours() / 24)
			if days > 0 {
				if days == 1 {
					fmt.Printf("  1 day\n\n")
				} else {
					fmt.Printf("  %d days\n\n", days)
				}
			}
			for _, dest := range trip.Destinations {
				if dest == nil {
					continue
				}
				if dest.Identifier == nil {
					continue
				}
				if dest.Identifier.PlaceID == "" {
					continue
				}
				place, err := placeService.GetPlace(ctx, dest.Identifier.PlaceID)
				if err != nil {
					return errors.Wrapf(err, "getting place %s", dest.Identifier.PlaceID)
				}
				fmt.Print("    ")
				showPlace(place)
				fmt.Print("\n")
			}
			continue
		}

		dur = dur.Round(time.Minute)
		durStr := strings.TrimSuffix(dur.String(), "0s")

		if activity != nil {
			fmt.Printf("    %s [%s]", startTime.Format("15:04"), durStr)
			if candidate := activity.TopCandidate; candidate != nil {
				fmt.Printf(" [%s]", candidate.Type)
			}
			fmt.Print("\n")
			continue
		}

		if visit != nil {
			fmt.Printf("  %s [%s]", startTime.Format("15:04"), durStr)
			if candidate := visit.TopCandidate; candidate != nil && candidate.PlaceID != "" {
				place, err := placeService.GetPlace(ctx, candidate.PlaceID)
				if err != nil {
					return errors.Wrapf(err, "getting place %s", candidate.PlaceID)
				}
				fmt.Print(" ")
				showPlace(place)
			}
			fmt.Print("\n")
		}
	}

	return nil
}

func showPlace(place *places.Place) {
	if place == nil {
		return
	}

	var (
		addr  = place.FormattedAddress
		loc   = place.Location
		dname = place.DisplayName
	)

	if addr == "" && loc == nil && dname == nil {
		return
	}

	if addr != "" {
		lines := strings.Split(addr, "\n")
		addr = strings.Join(lines, ", ")
	}

	if d := place.DisplayName; d != nil {
		fmt.Print(d.Text)
		if addr != "" {
			fmt.Printf(" (%s)", addr)
		}
	} else if addr != "" {
		fmt.Print(addr)
	} else {
		fmt.Printf("https://maps.google.com/?q=%.7f,%.7f", loc.Latitude, loc.Longitude)
	}
}
