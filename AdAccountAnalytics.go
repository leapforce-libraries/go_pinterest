package pinterest

import (
	"cloud.google.com/go/civil"
	"fmt"
	errortools "github.com/leapforce-libraries/go_errortools"
)

type GetAdAccountAnalyticsConfig struct {
	AdAccountId          string
	StartDate            civil.Date
	EndDate              civil.Date
	Columns              []string
	Granularity          Granularity
	ClickWindowDays      *WindowDays
	EngagementWindowDays *WindowDays
	ViewWindowDays       *WindowDays
	ConversionReportTime *ConversionReportTime
}

func (service *Service) GetAdAccountAnalytics(cfg *GetAdAccountAnalyticsConfig) (*[]Analytics, *errortools.Error) {
	if cfg == nil {
		return nil, errortools.ErrorMessage("GetAdAccountAnalyticsConfig must not be bil")
	}

	var url = fmt.Sprintf("ad_accounts/%s/analytics", cfg.AdAccountId)

	return service.getAnalytics(url, &getAnalyticsConfig{
		StartDate:            cfg.StartDate,
		EndDate:              cfg.EndDate,
		CampaignIds:          nil,
		Columns:              cfg.Columns,
		Granularity:          cfg.Granularity,
		ClickWindowDays:      cfg.ClickWindowDays,
		EngagementWindowDays: cfg.EngagementWindowDays,
		ViewWindowDays:       cfg.ViewWindowDays,
		ConversionReportTime: cfg.ConversionReportTime,
	})
}
