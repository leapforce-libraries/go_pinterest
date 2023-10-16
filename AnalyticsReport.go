package pinterest

import (
	"cloud.google.com/go/civil"
	"fmt"
	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	"net/http"
)

type ReportStatus string

const (
	ReportStatusDoesNotExist ReportStatus = "DOES_NOT_EXIST"
	ReportStatusFinished     ReportStatus = "FINISHED"
	ReportStatusInProgress   ReportStatus = "IN_PROGRESS"
	ReportStatusExpired      ReportStatus = "EXPIRED"
	ReportStatusFailed       ReportStatus = "FAILED"
	ReportStatusCancelled    ReportStatus = "CANCELLED"
)

type CreateAnalyticsReportConfig struct {
	StartDate              civil.Date            `json:"start_date"`
	EndDate                civil.Date            `json:"end_date"`
	Granularity            Granularity           `json:"granularity"`
	ClickWindowDays        *WindowDays           `json:"click_window_days,omitempty"`
	EngagementWindowDays   *WindowDays           `json:"engagement_window_days,omitempty"`
	ViewWindowDays         *WindowDays           `json:"view_window_days,omitempty"`
	ConversionReportTime   *ConversionReportTime `json:"conversion_report_time,omitempty"`
	AttributionTypes       *[]string             `json:"attribution_types,omitempty"`
	CampaignIds            *[]string             `json:"campaign_ids,omitempty"`
	CampaignStatuses       *[]string             `json:"campaign_statuses,omitempty"`
	CampaignObjectiveTypes *[]string             `json:"campaign_objective_types,omitempty"`
	AdGroupIds             *[]string             `json:"ad_group_ids,omitempty"`
	AdGroupStatuses        *[]string             `json:"ad_group_statuses,omitempty"`
	AdIds                  *[]string             `json:"ad_ids,omitempty"`
	AdStatuses             *[]string             `json:"ad_statuses,omitempty"`
	ProductGroupIds        *[]string             `json:"product_group_ids,omitempty"`
	ProductGroupStatuses   *[]string             `json:"product_group_statuses,omitempty"`
	ProductItemIds         *[]string             `json:"product_item_ids,omitempty"`
	TargetingTypes         *[]string             `json:"targeting_types,omitempty"`
	MetricFilters          *[]struct {
		Field    string   `json:"field"`
		Operator string   `json:"operator"`
		Values   []string `json:"values"`
	} `json:"metrics_filters,omitempty"`
	Columns      []string `json:"columns"`
	Level        string   `json:"level"`
	ReportFormat *string  `json:"report_format,omitempty"`
}

type CreateAnalyticsReportResponse struct {
	Status  ReportStatus `json:"report_status"`
	Token   string       `json:"token"`
	Message string       `json:"message"`
}

func (service *Service) CreateAnalyticsReport(adAccountId string, cfg *CreateAnalyticsReportConfig) (*CreateAnalyticsReportResponse, *errortools.Error) {
	if cfg == nil {
		return nil, errortools.ErrorMessage("CreateAnalyticsReportConfig must not be nil")
	}

	var response CreateAnalyticsReportResponse

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodPost,
		Url:           service.url(fmt.Sprintf("ad_accounts/%s/reports", adAccountId)),
		BodyModel:     cfg,
		ResponseModel: &response,
	}

	_, _, e := service.httpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &response, nil
}

type GetAnalyticsReportResponse struct {
	Status ReportStatus `json:"report_status"`
	Url    string       `json:"url"`
	Size   int64        `json:"size"`
}

func (service *Service) GetAnalyticsReport(adAccountId string, token string) (*GetAnalyticsReportResponse, *errortools.Error) {
	var response GetAnalyticsReportResponse

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodGet,
		Url:           service.url(fmt.Sprintf("ad_accounts/%s/reports?token=%s", adAccountId, token)),
		ResponseModel: &response,
	}

	_, _, e := service.httpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &response, nil
}
