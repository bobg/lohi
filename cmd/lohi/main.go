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
		placesCache string
		placesQPS   int
		tokenFile   string
	)

	flag.StringVar(&credsFile, "creds", "creds.json", "Google API credentials file")
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
		prevDayNum, err = showSemanticSegment(ctx, seg, prevDayNum, placeService)
		if err != nil {
			return errors.Wrap(err, "showing semantic segment")
		}
	}

	return nil
}

func showSemanticSegment(ctx context.Context, seg *schema.SemanticSegment, prevDayNum int, placeService places.Service) (int, error) {
	if seg == nil {
		return prevDayNum, nil
	}

	startTime, err := time.Parse(time.RFC3339, seg.StartTime)
	if err != nil {
		return 0, errors.Wrapf(err, "parsing start time %q", seg.StartTime)
	}
	startLoc := time.FixedZone("start", int(seg.StartTimeTimezoneUtcOffsetMinutes)*60)
	startTime = startTime.In(startLoc)

	endTime, err := time.Parse(time.RFC3339, seg.EndTime)
	if err != nil {
		return 0, errors.Wrapf(err, "parsing end time %q", seg.EndTime)
	}

	// Probably not necessary...
	endLoc := time.FixedZone("end", int(seg.EndTimeTimezoneUtcOffsetMinutes)*60)
	endTime = endTime.In(endLoc)

	dur := strings.TrimSuffix(endTime.Sub(startTime).Round(time.Minute).String(), "0s")

	if a := seg.Activity; a != nil {
		return showActivity(ctx, a, startTime, dur, prevDayNum)
	}
	if v := seg.Visit; v != nil {
		return showVisit(ctx, v, startTime, dur, prevDayNum, placeService)
	}

	return prevDayNum, nil
}

func showActivity(ctx context.Context, a *schema.Activity, startTime time.Time, dur string, prevDayNum int) (int, error) {
	if dayNum := startTime.Year()*1000 + startTime.YearDay(); dayNum != prevDayNum {
		fmt.Printf("\n%s\n\n", startTime.Format("2006-01-02"))
		prevDayNum = dayNum
	}

	fmt.Printf("    %s [%s]", startTime.Format("15:04"), dur)
	if candidate := a.TopCandidate; candidate != nil {
		fmt.Printf(" [%s]", candidate.Type)
	}
	fmt.Print("\n")

	return prevDayNum, nil
}

func showVisit(ctx context.Context, v *schema.Visit, startTime time.Time, dur string, prevDayNum int, placeService places.Service) (int, error) {
	if dayNum := startTime.Year()*1000 + startTime.YearDay(); dayNum != prevDayNum {
		fmt.Printf("\n%s\n\n", startTime.Format("2006-01-02"))
		prevDayNum = dayNum
	}

	fmt.Printf("  %s [%s]", startTime.Format("15:04"), dur)
	if candidate := v.TopCandidate; candidate != nil && candidate.PlaceID != "" {
		place, err := placeService.GetPlace(ctx, candidate.PlaceID)
		if err != nil {
			return 0, errors.Wrapf(err, "getting place %s", candidate.PlaceID)
		}
		fmt.Print(" ")
		showPlace(place)
	}
	fmt.Print("\n")

	return prevDayNum, nil
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
