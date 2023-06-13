package pinterest

import (
	"cloud.google.com/go/civil"
	"encoding/json"
	"fmt"
	errortools "github.com/leapforce-libraries/go_errortools"
)

type CampaignAnalytics struct {
	data map[string]json.RawMessage
}

type GetCampaignAnalyticsConfig struct {
	AdAccountId          string
	StartDate            civil.Date
	EndDate              civil.Date
	CampaignIds          []string
	Columns              []string
	Granularity          Granularity
	ClickWindowDays      *WindowDays
	EngagementWindowDays *WindowDays
	ViewWindowDays       *WindowDays
	ConversionReportTime *ConversionReportTime
}

func (service *Service) GetCampaignAnalytics(cfg *GetCampaignAnalyticsConfig) (*[]Analytics, *errortools.Error) {
	if cfg == nil {
		return nil, errortools.ErrorMessage("GetCampaignAnalyticsConfig must not be nil")
	}

	var url = fmt.Sprintf("ad_accounts/%s/campaigns/analytics", cfg.AdAccountId)

	return service.getAnalytics(url, &getAnalyticsConfig{
		StartDate:            cfg.StartDate,
		EndDate:              cfg.EndDate,
		CampaignIds:          &cfg.CampaignIds,
		Columns:              cfg.Columns,
		Granularity:          cfg.Granularity,
		ClickWindowDays:      cfg.ClickWindowDays,
		EngagementWindowDays: cfg.EngagementWindowDays,
		ViewWindowDays:       cfg.ViewWindowDays,
		ConversionReportTime: cfg.ConversionReportTime,
	})
}
