package pinterest

import (
	"fmt"
	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	"net/http"
	"net/url"
	"strings"
)

type Campaign struct {
	Id                           string       `json:"id"`
	AdAccountId                  string       `json:"ad_account_id"`
	Name                         string       `json:"name"`
	Status                       string       `json:"status"`
	LifetimeSpendCap             int64        `json:"lifetime_spend_cap"`
	DailySpendCap                int64        `json:"daily_spend_cap"`
	OrderLineId                  string       `json:"order_line_id"`
	TrackingUrls                 TrackingUrls `json:"tracking_urls"`
	StartTime                    int64        `json:"start_time"`
	EndTime                      int64        `json:"end_time"`
	SummaryStatus                string       `json:"summary_status"`
	ObjectiveType                string       `json:"objective_type"`
	CreatedTime                  int64        `json:"created_time"`
	UpdatedTime                  int64        `json:"updated_time"`
	Type                         string       `json:"type"`
	IsFlexibleDailyBudgets       bool         `json:"is_flexible_daily_budgets"`
	IsCampaignBudgetOptimization bool         `json:"is_campaign_budget_optimization"`
}

type ListCampaignsConfig struct {
	AdAccountId    string
	CampaignIds    *[]string
	EntityStatuses *[]EntityStatus
	PageSize       *int64
	Order          *Order
}

type ListCampaignsResponse struct {
	Items    []Campaign
	Bookmark string
}

func (service *Service) ListCampaigns(cfg *ListCampaignsConfig) (*[]Campaign, *errortools.Error) {
	if cfg == nil {
		return nil, errortools.ErrorMessage("ListCampaignsConfig must not be bil")
	}

	var campaigns []Campaign

	var values = url.Values{}
	if cfg.CampaignIds != nil {
		values.Set("campaign_ids", strings.Join(*cfg.CampaignIds, ","))
	}
	if cfg.EntityStatuses != nil {
		var s []string
		for _, e := range *cfg.EntityStatuses {
			s = append(s, string(e))
		}
		values.Set("entity_statuses", strings.Join(s, ","))
	}
	if cfg.PageSize != nil {
		values.Set("page_size", fmt.Sprintf("%v", *cfg.PageSize))
	}
	if cfg.Order != nil {
		values.Set("order", string(*cfg.Order))
	}

	for {
		var response ListCampaignsResponse

		requestConfig := go_http.RequestConfig{
			Method:        http.MethodGet,
			Url:           service.url(fmt.Sprintf("ad_accounts/%s/campaigns?%s", cfg.AdAccountId, values.Encode())),
			ResponseModel: &response,
		}

		_, _, e := service.httpRequest(&requestConfig)
		if e != nil {
			return nil, e
		}

		campaigns = append(campaigns, response.Items...)

		if response.Bookmark == "" {
			break
		}

		values.Set("bookmark", response.Bookmark)
	}

	return &campaigns, nil
}
