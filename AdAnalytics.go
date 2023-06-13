package pinterest

import (
	"cloud.google.com/go/civil"
	"encoding/json"
	"fmt"
	errortools "github.com/leapforce-libraries/go_errortools"
)

type AdAnalytics struct {
	data map[string]json.RawMessage
}

type GetAdAnalyticsConfig struct {
	AdAccountId          string
	StartDate            civil.Date
	EndDate              civil.Date
	AdIds                []string
	Columns              []string
	Granularity          Granularity
	ClickWindowDays      *WindowDays
	EngagementWindowDays *WindowDays
	ViewWindowDays       *WindowDays
	ConversionReportTime *ConversionReportTime
}

func (service *Service) GetAdAnalytics(cfg *GetAdAnalyticsConfig) (*[]Analytics, *errortools.Error) {
	if cfg == nil {
		return nil, errortools.ErrorMessage("GetAdAnalyticsConfig must not be nil")
	}

	var url = fmt.Sprintf("ad_accounts/%s/ads/analytics", cfg.AdAccountId)

	return service.getAnalytics(url, &getAnalyticsConfig{
		StartDate:            cfg.StartDate,
		EndDate:              cfg.EndDate,
		AdIds:                &cfg.AdIds,
		Columns:              cfg.Columns,
		Granularity:          cfg.Granularity,
		ClickWindowDays:      cfg.ClickWindowDays,
		EngagementWindowDays: cfg.EngagementWindowDays,
		ViewWindowDays:       cfg.ViewWindowDays,
		ConversionReportTime: cfg.ConversionReportTime,
	})
}
