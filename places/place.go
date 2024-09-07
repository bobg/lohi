package places

type Place struct {
	AccessibilityOptions         *AccessibilityOptions           `json:"accessibilityOptions,omitempty"`
	AddressComponents            []*AddressComponents            `json:"addressComponents,omitempty"`
	AdrFormatAddress             string                          `json:"adrFormatAddress,omitempty"`
	AreaSummary                  *AreaSummary                    `json:"areaSummary,omitempty"`
	BusinessStatus               string                          `json:"businessStatus,omitempty"`
	CurbsidePickup               bool                            `json:"curbsidePickup,omitempty"`
	CurrentOpeningHours          *CurrentOpeningHours            `json:"currentOpeningHours,omitempty"`
	CurrentSecondaryOpeningHours []*CurrentSecondaryOpeningHours `json:"currentSecondaryOpeningHours,omitempty"`
	Delivery                     bool                            `json:"delivery,omitempty"`
	DineIn                       bool                            `json:"dineIn,omitempty"`
	DisplayName                  *LocalizedText                  `json:"displayName,omitempty"`
	EditorialSummary             *LocalizedText                  `json:"editorialSummary,omitempty"`
	FormattedAddress             string                          `json:"formattedAddress,omitempty"`
	GenerativeSummary            *GenerativeSummary              `json:"generativeSummary,omitempty"`
	GoodForChildren              bool                            `json:"goodForChildren,omitempty"`
	GoodForGroups                bool                            `json:"goodForGroups,omitempty"`
	GoodForWatchingSports        bool                            `json:"goodForWatchingSports,omitempty"`
	GoogleMapsUri                string                          `json:"googleMapsUri,omitempty"`
	IconBackgroundColor          string                          `json:"iconBackgroundColor,omitempty"`
	IconMaskBaseUri              string                          `json:"iconMaskBaseUri,omitempty"`
	Id                           string                          `json:"id,omitempty"`
	InternationalPhoneNumber     string                          `json:"internationalPhoneNumber,omitempty"`
	LiveMusic                    bool                            `json:"liveMusic,omitempty"`
	Location                     *LatLong                        `json:"location,omitempty"`
	MenuForChildren              bool                            `json:"menuForChildren,omitempty"`
	Name                         string                          `json:"name,omitempty"`
	NationalPhoneNumber          string                          `json:"nationalPhoneNumber,omitempty"`
	OutdoorSeating               bool                            `json:"outdoorSeating,omitempty"`
	ParkingOptions               *ParkingOptions                 `json:"parkingOptions,omitempty"`
	PaymentOptions               *PaymentOptions                 `json:"paymentOptions,omitempty"`
	Photos                       []*Photo                        `json:"photos,omitempty"`
	PlusCode                     *PlusCode                       `json:"plusCode,omitempty"`
	PriceLevel                   string                          `json:"priceLevel,omitempty"`
	PrimaryType                  string                          `json:"primaryType,omitempty"`
	PrimaryTypeDisplayName       *LocalizedText                  `json:"primaryTypeDisplayName,omitempty"`
	Rating                       float64                         `json:"rating,omitempty"`
	RegularOpeningHours          *RegularOpeningHours            `json:"regularOpeningHours,omitempty"`
	RegularSecondaryOpeningHours []*RegularSecondaryOpeningHours `json:"regularSecondaryOpeningHours,omitempty"`
	Reservable                   bool                            `json:"reservable,omitempty"`
	Restroom                     bool                            `json:"restroom,omitempty"`
	Reviews                      []*Review                       `json:"reviews,omitempty"`
	ServesBeer                   bool                            `json:"servesBeer,omitempty"`
	ServesBreakfast              bool                            `json:"servesBreakfast,omitempty"`
	ServesBrunch                 bool                            `json:"servesBrunch,omitempty"`
	ServesCocktails              bool                            `json:"servesCocktails,omitempty"`
	ServesCoffee                 bool                            `json:"servesCoffee,omitempty"`
	ServesDessert                bool                            `json:"servesDessert,omitempty"`
	ServesDinner                 bool                            `json:"servesDinner,omitempty"`
	ServesLunch                  bool                            `json:"servesLunch,omitempty"`
	ServesVegetarianFood         bool                            `json:"servesVegetarianFood,omitempty"`
	ServesWine                   bool                            `json:"servesWine,omitempty"`
	ShortFormattedAddress        string                          `json:"shortFormattedAddress,omitempty"`
	Takeout                      bool                            `json:"takeout,omitempty"`
	Types                        []string                        `json:"types,omitempty"`
	UserRatingCount              int64                           `json:"userRatingCount,omitempty"`
	UtcOffsetMinutes             int64                           `json:"utcOffsetMinutes,omitempty"`
	Viewport                     *Viewport                       `json:"viewport,omitempty"`
	WebsiteUri                   string                          `json:"websiteUri,omitempty"`
}

type AccessibilityOptions struct {
	WheelchairAccessibleEntrance bool `json:"wheelchairAccessibleEntrance,omitempty"`
	WheelchairAccessibleParking  bool `json:"wheelchairAccessibleParking,omitempty"`
	WheelchairAccessibleRestroom bool `json:"wheelchairAccessibleRestroom,omitempty"`
	WheelchairAccessibleSeating  bool `json:"wheelchairAccessibleSeating,omitempty"`
}

type AddressComponents struct {
	LanguageCode string   `json:"languageCode,omitempty"`
	LongText     string   `json:"longText,omitempty"`
	ShortText    string   `json:"shortText,omitempty"`
	Types        []string `json:"types,omitempty"`
}

type AreaSummary struct {
	ContentBlocks []*ContentBlock `json:"contentBlocks,omitempty"`
}

type CurrentOpeningHours struct {
	OpenNow             bool              `json:"openNow,omitempty"`
	Periods             []*DateTimePeriod `json:"periods,omitempty"`
	WeekdayDescriptions []string          `json:"weekdayDescriptions,omitempty"`
}

type CurrentSecondaryOpeningHours struct {
	OpenNow             bool              `json:"openNow,omitempty"`
	Periods             []*DateTimePeriod `json:"periods,omitempty"`
	SecondaryHoursType  string            `json:"secondaryHoursType,omitempty"`
	WeekdayDescriptions []string          `json:"weekdayDescriptions,omitempty"`
}

type LocalizedText struct {
	LanguageCode string `json:"languageCode,omitempty"`
	Text         string `json:"text,omitempty"`
}

type GenerativeSummary struct {
	Description *LocalizedText `json:"description,omitempty"`
	Overview    *LocalizedText `json:"overview,omitempty"`
	References  *Reviews       `json:"references,omitempty"`
}

type LatLong struct {
	Latitude  float64 `json:"latitude,omitempty"`
	Longitude float64 `json:"longitude,omitempty"`
}

type ParkingOptions struct {
	FreeParkingLot    bool `json:"freeParkingLot,omitempty"`
	FreeStreetParking bool `json:"freeStreetParking,omitempty"`
	PaidStreetParking bool `json:"paidStreetParking,omitempty"`
	ValetParking      bool `json:"valetParking,omitempty"`
}

type PaymentOptions struct {
	AcceptsCashOnly    bool `json:"acceptsCashOnly,omitempty"`
	AcceptsCreditCards bool `json:"acceptsCreditCards,omitempty"`
	AcceptsDebitCards  bool `json:"acceptsDebitCards,omitempty"`
	AcceptsNfc         bool `json:"acceptsNfc,omitempty"`
}

type Photo struct {
	AuthorAttributions []*AuthorAttribution `json:"authorAttributions,omitempty"`
	HeightPx           int64                `json:"heightPx,omitempty"`
	Name               string               `json:"name,omitempty"`
	WidthPx            int64                `json:"widthPx,omitempty"`
}

type PlusCode struct {
	CompoundCode string `json:"compoundCode,omitempty"`
	GlobalCode   string `json:"globalCode,omitempty"`
}

type RegularOpeningHours struct {
	OpenNow             bool          `json:"openNow,omitempty"`
	Periods             []*TimePeriod `json:"periods,omitempty"`
	WeekdayDescriptions []string      `json:"weekdayDescriptions,omitempty"`
}

type RegularSecondaryOpeningHours struct {
	OpenNow             bool          `json:"openNow,omitempty"`
	Periods             []*TimePeriod `json:"periods,omitempty"`
	SecondaryHoursType  string        `json:"secondaryHoursType,omitempty"`
	WeekdayDescriptions []string      `json:"weekdayDescriptions,omitempty"`
}

type Review struct {
	AuthorAttribution              *AuthorAttribution `json:"authorAttribution,omitempty"`
	Name                           string             `json:"name,omitempty"`
	OriginalText                   *LocalizedText     `json:"originalText,omitempty"`
	PublishTime                    string             `json:"publishTime,omitempty"`
	Rating                         int64              `json:"rating,omitempty"`
	RelativePublishTimeDescription string             `json:"relativePublishTimeDescription,omitempty"`
	Text                           *LocalizedText     `json:"text,omitempty"`
}

type Viewport struct {
	High *LatLong `json:"high,omitempty"`
	Low  *LatLong `json:"low,omitempty"`
}

type ContentBlock struct {
	Content    *LocalizedText `json:"content,omitempty"`
	References *Places        `json:"references,omitempty"`
	Topic      string         `json:"topic,omitempty"`
}

type DateTimePeriod struct {
	Close *DateTime `json:"close,omitempty"`
	Open  *DateTime `json:"open,omitempty"`
}

type Reviews struct {
	Reviews []*Review `json:"reviews,omitempty"`
}

type AuthorAttribution struct {
	DisplayName string `json:"displayName,omitempty"`
	PhotoUri    string `json:"photoUri,omitempty"`
	Uri         string `json:"uri,omitempty"`
}

type TimePeriod struct {
	Close *Time `json:"close,omitempty"`
	Open  *Time `json:"open,omitempty"`
}

type Places struct {
	Places []string `json:"places,omitempty"`
}

type DateTime struct {
	Date   *Date `json:"date,omitempty"`
	Day    int64 `json:"day,omitempty"`
	Hour   int64 `json:"hour,omitempty"`
	Minute int64 `json:"minute,omitempty"`
}

type Time struct {
	Day    int64 `json:"day,omitempty"`
	Hour   int64 `json:"hour,omitempty"`
	Minute int64 `json:"minute,omitempty"`
}

type Date struct {
	Day   int64 `json:"day,omitempty"`
	Month int64 `json:"month,omitempty"`
	Year  int64 `json:"year,omitempty"`
}
