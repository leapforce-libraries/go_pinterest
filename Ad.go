package pinterest

import (
	"fmt"
	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	"net/http"
	"net/url"
	"strings"
)

type Ad struct {
	AdGroupId                             string       `json:"ad_group_id"`
	AndroidDeepLink                       string       `json:"android_deep_link"`
	CarouselAndroidDeepLinks              []string     `json:"carousel_android_deep_links"`
	CarouselDestinationUrls               []string     `json:"carousel_destination_urls"`
	CarouselIosDeepLinks                  []string     `json:"carousel_ios_deep_links"`
	ClickTrackingUrl                      string       `json:"click_tracking_url"`
	CreativeType                          string       `json:"creative_type"`
	DestinationUrl                        string       `json:"destination_url"`
	IosDeepLink                           string       `json:"ios_deep_link"`
	IsPinDeleted                          bool         `json:"is_pin_deleted"`
	IsRemovable                           bool         `json:"is_removable"`
	Name                                  string       `json:"name"`
	Status                                string       `json:"status"`
	TrackingUrls                          TrackingUrls `json:"tracking_urls"`
	ViewTrackingUrl                       string       `json:"view_tracking_url"`
	PinId                                 string       `json:"pin_id"`
	AdAccountId                           string       `json:"ad_account_id"`
	CampaignId                            string       `json:"campaign_id"`
	CollectionItemsDestinationUrlTemplate string       `json:"collection_items_destination_url_template"`
	CreatedTime                           int64        `json:"created_time"`
	Id                                    string       `json:"id"`
	RejectedReasons                       []string     `json:"rejected_reasons"`
	RejectionLabels                       []string     `json:"rejection_labels"`
	ReviewStatus                          string       `json:"review_status"`
	Type                                  string       `json:"type"`
	UpdatedTime                           int64        `json:"updated_time"`
	SummaryStatus                         string       `json:"summary_status"`
}

type ListAdsConfig struct {
	AdAccountId    string
	CampaignIds    *[]string
	AdGroupIds     *[]string
	AdIds          *[]string
	EntityStatuses *[]EntityStatus
	PageSize       *int64
	Order          *Order
}

type ListAdsResponse struct {
	Items    []Ad
	Bookmark string
}

func (service *Service) ListAds(cfg *ListAdsConfig) (*[]Ad, *errortools.Error) {
	if cfg == nil {
		return nil, errortools.ErrorMessage("ListAdsConfig must not be bil")
	}

	var campaigns []Ad

	var values = url.Values{}
	if cfg.CampaignIds != nil {
		values.Set("campaign_ids", strings.Join(*cfg.CampaignIds, ","))
	}
	if cfg.AdGroupIds != nil {
		values.Set("ad_group_ids", strings.Join(*cfg.AdGroupIds, ","))
	}
	if cfg.AdIds != nil {
		values.Set("ad_ids", strings.Join(*cfg.AdIds, ","))
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
		var response ListAdsResponse

		requestConfig := go_http.RequestConfig{
			Method:        http.MethodGet,
			Url:           service.url(fmt.Sprintf("ad_accounts/%s/ads?%s", cfg.AdAccountId, values.Encode())),
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
