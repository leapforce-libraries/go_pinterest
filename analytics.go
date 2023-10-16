package pinterest

import (
	"cloud.google.com/go/civil"
	"encoding/json"
	"fmt"
	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	"net/http"
	"net/url"
	"strings"
)

type Analytics map[string]json.RawMessage

type getAnalyticsConfig struct {
	StartDate            civil.Date
	EndDate              civil.Date
	CampaignIds          *[]string
	AdGroupIds           *[]string
	AdIds                *[]string
	Columns              []string
	Granularity          Granularity
	ClickWindowDays      *WindowDays
	EngagementWindowDays *WindowDays
	ViewWindowDays       *WindowDays
	ConversionReportTime *ConversionReportTime
}

func (a *Analytics) value(key string) json.RawMessage {
	if a == nil {
		return nil
	}

	j, ok := (*a)[key]
	if !ok {
		return nil
	}

	return j
}

func (a *Analytics) valueString(key string) string {
	j := a.value(key)
	if j == nil {
		return ""
	}

	var s string
	err := json.Unmarshal(j, &s)
	if err != nil {
		return ""
	}

	return s
}

func (a *Analytics) valueDate(key string) *civil.Date {
	j := a.value(key)
	if j == nil {
		return nil
	}

	var d civil.Date
	err := json.Unmarshal(j, &d)
	if err != nil {
		return nil
	}

	return &d
}

func (a *Analytics) valueFloat64(key string) *float64 {
	j := a.value(key)
	if j == nil {
		return nil
	}

	var f float64
	err := json.Unmarshal(j, &f)
	if err != nil {
		return nil
	}

	return &f
}

func (a *Analytics) valueInt64(key string) *int64 {
	j := a.value(key)
	if j == nil {
		return nil
	}

	var i int64
	err := json.Unmarshal(j, &i)
	if err != nil {
		return nil
	}

	return &i
}

func (a *Analytics) AdAccountId() string {
	return a.id("AD_ACCOUNT_ID")
}

func (a *Analytics) CampaignId() string {
	return a.id("CAMPAIGN_ID")
}

func (a *Analytics) AdGroupId() string {
	return a.id("AD_GROUP_ID")
}

func (a *Analytics) AdId() string {
	return a.id("AD_ID")
}

func (a *Analytics) id(key string) string {
	id := a.valueString(key)
	if id == "" {
		i := a.valueInt64(key)
		if i != nil {
			id = fmt.Sprintf("%v", *i)
		}
	}
	return id
}

func (a *Analytics) Date() *civil.Date {
	return a.valueDate("DATE")
}

func (a *Analytics) Float64(key string) *float64 {
	return a.valueFloat64(key)
}

func (a *Analytics) Int64(key string) *int64 {
	return a.valueInt64(key)
}

func (service *Service) getAnalytics(url_ string, cfg *getAnalyticsConfig) (*[]Analytics, *errortools.Error) {
	if cfg == nil {
		return nil, errortools.ErrorMessage("CreateAnalyticsReportConfig must not be nil")
	}

	var values = url.Values{}
	values.Set("start_date", cfg.StartDate.String())
	values.Set("end_date", cfg.EndDate.String())
	values.Set("columns", strings.Join(cfg.Columns, ","))
	values.Set("granularity", string(cfg.Granularity))
	if cfg.CampaignIds != nil {
		values.Set("campaign_ids", strings.Join(*cfg.CampaignIds, ","))
	}
	if cfg.AdGroupIds != nil {
		values.Set("ad_group_ids", strings.Join(*cfg.AdGroupIds, ","))
	}
	if cfg.AdIds != nil {
		values.Set("ad_ids", strings.Join(*cfg.AdIds, ","))
	}
	if cfg.ClickWindowDays != nil {
		values.Set("click_window_days", fmt.Sprintf("%v", *cfg.ClickWindowDays))
	}
	if cfg.EngagementWindowDays != nil {
		values.Set("engagement_window_days", fmt.Sprintf("%v", *cfg.EngagementWindowDays))
	}
	if cfg.ViewWindowDays != nil {
		values.Set("view_window_days", fmt.Sprintf("%v", *cfg.ViewWindowDays))
	}
	if cfg.ConversionReportTime != nil {
		values.Set("conversion_report_time", string(*cfg.ConversionReportTime))
	}

	var analytics []Analytics

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodGet,
		Url:           service.url(fmt.Sprintf("%s?%s", url_, values.Encode())),
		ResponseModel: &analytics,
	}

	_, _, e := service.httpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &analytics, nil
}
