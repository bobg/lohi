package schema

// History is the top-level structure of a Google Location History JSON file.
type History struct {
	RawSignals          []*RawSignal       `json:"rawSignals,omitempty"`
	SemanticSegments    []*SemanticSegment `json:"semanticSegments,omitempty"`
	UserLocationProfile *LocationProfile   `json:"userLocationProfile,omitempty"`
}

type (
	Activity struct {
		DistanceMeters float64            `json:"distanceMeters,omitempty"`
		End            *Location          `json:"end,omitempty"`
		Parking        *Parking           `json:"parking,omitempty"`
		Probability    float64            `json:"probability,omitempty"`
		Start          *Location          `json:"start,omitempty"`
		TopCandidate   *ActivityCandidate `json:"topCandidate,omitempty"`
	}

	ActivityCandidate struct {
		Probability float64 `json:"probability,omitempty"`
		Type        string  `json:"type,omitempty"`
	}

	ActivityRecord struct {
		ProbableActivities []*ProbableActivity `json:"probableActivities,omitempty"`
		Timestamp          string              `json:"timestamp,omitempty"`
	}

	DeviceRecord struct {
		Mac     int64 `json:"mac,omitempty"`
		RawRssi int64 `json:"rawRssi,omitempty"`
	}

	FrequentPlace struct {
		Label         string `json:"label,omitempty"`
		PlaceID       string `json:"placeId,omitempty"`
		PlaceLocation string `json:"placeLocation,omitempty"`
	}

	IdentifiedPlace struct {
		Identifier *PlaceID `json:"identifier,omitempty"`
	}

	Location struct {
		LatLng string `json:"latLng,omitempty"`
	}

	LocationProfile struct {
		FrequentPlaces []*FrequentPlace `json:"frequentPlaces,omitempty"`
	}

	Parking struct {
		Location  *Location `json:"location,omitempty"`
		StartTime string    `json:"startTime,omitempty"`
	}

	PlaceCandidate struct {
		PlaceID       string    `json:"placeId,omitempty"`
		PlaceLocation *Location `json:"placeLocation,omitempty"`
		Probability   float64   `json:"probability,omitempty"`
		SemanticType  string    `json:"semanticType,omitempty"`
	}

	PlaceID struct {
		PlaceID string `json:"placeId,omitempty"`
	}

	Position struct {
		AccuracyMeters       int64   `json:"accuracyMeters,omitempty"`
		AltitudeMeters       float64 `json:"altitudeMeters,omitempty"`
		LatLng               string  `json:"LatLng,omitempty"`
		Source               string  `json:"source,omitempty"`
		SpeedMetersPerSecond float64 `json:"speedMetersPerSecond,omitempty"`
		Timestamp            string  `json:"timestamp,omitempty"`
	}

	ProbableActivity struct {
		Confidence float64 `json:"confidence,omitempty"`
		Type       string  `json:"type,omitempty"`
	}

	RawSignal struct {
		ActivityRecord *ActivityRecord `json:"activityRecord,omitempty"`
		Position       *Position       `json:"position,omitempty"`
		WifiScan       *WifiScan       `json:"wifiScan,omitempty"`
	}

	SemanticSegment struct {
		Activity                          *Activity       `json:"activity,omitempty"`
		EndTime                           string          `json:"endTime,omitempty"`
		EndTimeTimezoneUtcOffsetMinutes   int64           `json:"endTimeTimezoneUtcOffsetMinutes,omitempty"`
		StartTime                         string          `json:"startTime,omitempty"`
		StartTimeTimezoneUtcOffsetMinutes int64           `json:"startTimeTimezoneUtcOffsetMinutes,omitempty"`
		TimelineMemory                    *TimelineMemory `json:"timelineMemory,omitempty"`
		TimelinePath                      []*TimelinePath `json:"timelinePath,omitempty"`
		Visit                             *Visit          `json:"visit,omitempty"`
	}

	TimelineMemory struct {
		Trip *Trip `json:"trip,omitempty"`
	}

	TimelinePath struct {
		Point string `json:"point,omitempty"`
		Time  string `json:"time,omitempty"`
	}

	Visit struct {
		HierarchyLevel  int64           `json:"hierarchyLevel,omitempty"`
		IsTimelessVisit bool            `json:"isTimelessVisit,omitempty"`
		Probability     float64         `json:"probability,omitempty"`
		TopCandidate    *PlaceCandidate `json:"topCandidate,omitempty"`
	}

	WifiScan struct {
		DeliveryTime   string          `json:"deliveryTime,omitempty"`
		DevicesRecords []*DeviceRecord `json:"devicesRecords,omitempty"`
	}

	Trip struct {
		Destinations          []*IdentifiedPlace `json:"destinations,omitempty"`
		DistanceFromOriginKms int64              `json:"distanceFromOriginKms,omitempty"`
	}
)
