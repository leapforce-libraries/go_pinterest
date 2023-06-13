package pinterest

import (
	"cloud.google.com/go/civil"
	"encoding/json"
	"fmt"
	errortools "github.com/leapforce-libraries/go_errortools"
)

type AdGroupAnalytics struct {
	data map[string]json.RawMessage
}

type GetAdGroupAnalyticsConfig struct {
	AdAccountId          string
	StartDate            civil.Date
	EndDate              civil.Date
	AdGroupIds           []string
	Columns              []string
	Granularity          Granularity
	ClickWindowDays      *WindowDays
	EngagementWindowDays *WindowDays
	ViewWindowDays       *WindowDays
	ConversionReportTime *ConversionReportTime
}

func (service *Service) GetAdGroupAnalytics(cfg *GetAdGroupAnalyticsConfig) (*[]Analytics, *errortools.Error) {
	if cfg == nil {
		return nil, errortools.ErrorMessage("GetAdGroupAnalyticsConfig must not be nil")
	}

	var url = fmt.Sprintf("ad_accounts/%s/ad_groups/analytics", cfg.AdAccountId)

	return service.getAnalytics(url, &getAnalyticsConfig{
		StartDate:            cfg.StartDate,
		EndDate:              cfg.EndDate,
		AdGroupIds:           &cfg.AdGroupIds,
		Columns:              cfg.Columns,
		Granularity:          cfg.Granularity,
		ClickWindowDays:      cfg.ClickWindowDays,
		EngagementWindowDays: cfg.EngagementWindowDays,
		ViewWindowDays:       cfg.ViewWindowDays,
		ConversionReportTime: cfg.ConversionReportTime,
	})
}
