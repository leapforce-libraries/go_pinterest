package pinterest

type Owner struct {
	Username string `json:"username"`
}

type Granularity string

const (
	GranularityTotal Granularity = "TOTAL"
	GranularityDay   Granularity = "DAY"
	GranularityHour  Granularity = "HOUR"
	GranularityWeek  Granularity = "WEEK"
	GranularityMonth Granularity = "MONTH"
)

type ConversionReportTime string

const (
	ConversionReportTimeAction     ConversionReportTime = "TIME_OF_AD_ACTION"
	ConversionReportTimeConversion ConversionReportTime = "TIME_OF_CONVERSION"
)

type WindowDays int64

const (
	WindowDays0  WindowDays = 0
	WindowDays1  WindowDays = 1
	WindowDays7  WindowDays = 7
	WindowDays14 WindowDays = 14
	WindowDays30 WindowDays = 30
	WindowDays60 WindowDays = 60
)

type EntityStatus string

const (
	EntityStatusActive   EntityStatus = "ACTIVE"
	EntityStatusPaused   EntityStatus = "PAUSED"
	EntityStatusArchived EntityStatus = "ARCHIVED"
)

type Order string

const (
	OrderAscending  Order = "ASCENDING"
	OrderDescending Order = "DESCENDING"
)

type TrackingUrls struct {
	Impression           []string `json:"impression"`
	Click                []string `json:"click"`
	Engagement           []string `json:"engagement"`
	BuyableButton        []string `json:"buyable_button"`
	AudienceVerification []string `json:"audience_verification"`
}
