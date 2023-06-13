package pinterest

import (
	"fmt"
	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	"net/http"
	"net/url"
	"strings"
)

type AdGroup struct {
	Name                       string                          `json:"name"`
	Status                     string                          `json:"status"`
	BudgetInMicroCurrency      int64                           `json:"budget_in_micro_currency"`
	BidInMicroCurrency         int64                           `json:"bid_in_micro_currency"`
	OptimizationGoalMetadata   AdGroupOptimizationGoalMetadata `json:"optimization_goal_metadata"`
	BudgetType                 string                          `json:"budget_type"`
	StartTime                  int64                           `json:"start_time"`
	EndTime                    int64                           `json:"end_time"`
	TargetingSpec              AdGroupTargetingSpec            `json:"targeting_spec"`
	LifetimeFrequencyCap       int64                           `json:"lifetime_frequency_cap"`
	TrackingUrls               TrackingUrls                    `json:"tracking_urls"`
	AutoTargetingEnabled       bool                            `json:"auto_targeting_enabled"`
	PlacementGroup             string                          `json:"placement_group"`
	PacingDeliveryType         string                          `json:"pacing_delivery_type"`
	CampaignId                 string                          `json:"campaign_id"`
	BillableEvent              string                          `json:"billable_event"`
	BidStrategyType            string                          `json:"bid_strategy_type"`
	Id                         string                          `json:"id"`
	AdAccountId                string                          `json:"ad_account_id"`
	CreatedTime                int64                           `json:"created_time"`
	UpdatedTime                int64                           `json:"updated_time"`
	Type                       string                          `json:"type"`
	ConversionLearningModeType string                          `json:"conversion_learning_mode_type"`
	SummaryStatus              string                          `json:"summary_status"`
	FeedProfileId              string                          `json:"feed_profile_id"`
	DcaAssets                  interface{}                     `json:"dca_assets"`
}

type AdGroupOptimizationGoalMetadata struct {
	ConversionTagV3GoalMetadata struct {
		AttributionWindows struct {
			ClickWindowDays      int64 `json:"click_window_days"`
			EngagementWindowDays int64 `json:"engagement_window_days"`
			ViewWindowDays       int64 `json:"view_window_days"`
		} `json:"attribution_windows"`
		ConversionEvent             string `json:"conversion_event"`
		ConversionTagId             string `json:"conversion_tag_id"`
		CpaGoalValueInMicroCurrency string `json:"cpa_goal_value_in_micro_currency"`
		IsRoasOptimized             bool   `json:"is_roas_optimized"`
		LearningModeType            string `json:"learning_mode_type"`
	} `json:"conversion_tag_v3_goal_metadata"`
	FrequencyGoalMetadata struct {
		Frequency int64  `json:"frequency"`
		Timerange string `json:"timerange"`
	} `json:"frequency_goal_metadata"`
	ScrollupGoalMetadata struct {
		ScrollupGoalValueInMicroCurrency string `json:"scrollup_goal_value_in_micro_currency"`
	} `json:"scrollup_goal_metadata"`
}

type AdGroupTargetingSpec struct {
	AGEBUCKET           []string `json:"AGE_BUCKET"`
	APPTYPE             []string `json:"APPTYPE"`
	AUDIENCEEXCLUDE     []string `json:"AUDIENCE_EXCLUDE"`
	AUDIENCEINCLUDE     []string `json:"AUDIENCE_INCLUDE'"`
	GENDER              []string `json:"GENDER"`
	GEO                 []string `json:"GEO"`
	INTEREST            []string `json:"INTEREST"`
	LOCALE              []string `json:"LOCALE"`
	LOCATION            []string `json:"LOCATION"`
	SHOPPINGRETARGETING []struct {
		LookbackWindow  int64   `json:"lookback_window"`
		TagTypes        []int64 `json:"tag_types"`
		ExclusionWindow int64   `json:"exclusion_window"`
	} `json:"SHOPPING_RETARGETING"`
	TARGETINGSTRATEGY []string `json:"TARGETING_STRATEGY"`
}

type ListAdGroupsConfig struct {
	AdAccountId               string
	CampaignIds               *[]string
	AdGroupIds                *[]string
	EntityStatuses            *[]EntityStatus
	PageSize                  *int64
	Order                     *Order
	TranslateInterestsToNames *bool
}

type ListAdGroupsResponse struct {
	Items    []AdGroup
	Bookmark string
}

func (service *Service) ListAdGroups(cfg *ListAdGroupsConfig) (*[]AdGroup, *errortools.Error) {
	if cfg == nil {
		return nil, errortools.ErrorMessage("ListAdGroupsConfig must not be bil")
	}

	var campaigns []AdGroup

	var values = url.Values{}
	if cfg.CampaignIds != nil {
		values.Set("campaign_ids", strings.Join(*cfg.CampaignIds, ","))
	}
	if cfg.AdGroupIds != nil {
		values.Set("ad_group_ids", strings.Join(*cfg.AdGroupIds, ","))
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
	if cfg.TranslateInterestsToNames != nil {
		values.Set("translate_interests_to_names", fmt.Sprintf("%v", *cfg.TranslateInterestsToNames))
	}

	for {
		var response ListAdGroupsResponse

		requestConfig := go_http.RequestConfig{
			Method:        http.MethodGet,
			Url:           service.url(fmt.Sprintf("ad_accounts/%s/ad_groups?%s", cfg.AdAccountId, values.Encode())),
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
