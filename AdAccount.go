package pinterest

import (
	"fmt"
	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	"net/http"
	"net/url"
)

type AdAccount struct {
	Id          string   `json:"id"`
	Name        string   `json:"name"`
	Owner       Owner    `json:"owner"`
	Country     string   `json:"country"`
	Currency    string   `json:"currency"`
	Permissions []string `json:"permissions"`
	CreatedTime int64    `json:"created_time"`
	UpdatedTime int64    `json:"updated_time"`
}

type ListAdAccountsConfig struct {
	PageSize *int64
}

type ListAdAccountsResponse struct {
	Items    []AdAccount
	Bookmark string
}

func (service *Service) ListAdAccounts(cfg *ListAdAccountsConfig) (*[]AdAccount, *errortools.Error) {
	var adAccounts []AdAccount

	var values = url.Values{}
	if cfg != nil {
		if cfg.PageSize != nil {
			values.Set("page_size", fmt.Sprintf("%v", *cfg.PageSize))
		}
	}

	for {
		var response ListAdAccountsResponse

		requestConfig := go_http.RequestConfig{
			Method:        http.MethodGet,
			Url:           service.url(fmt.Sprintf("ad_accounts?%s", values.Encode())),
			ResponseModel: &response,
		}

		_, _, e := service.httpRequest(&requestConfig)
		if e != nil {
			return nil, e
		}

		adAccounts = append(adAccounts, response.Items...)

		if response.Bookmark == "" {
			break
		}

		values.Set("bookmark", response.Bookmark)
	}

	return &adAccounts, nil
}

func (service *Service) GetAdAccount(adAccountId string) (*AdAccount, *errortools.Error) {
	var adAccount AdAccount

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodGet,
		Url:           service.url(fmt.Sprintf("ad_accounts/%s", adAccountId)),
		ResponseModel: &adAccount,
	}

	_, _, e := service.httpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &adAccount, nil
}
