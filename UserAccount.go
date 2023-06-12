package pinterest

import (
	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	"net/http"
)

type UserAccount struct {
	FollowingCount int64  `json:"following_count"`
	BusinessName   string `json:"business_name"`
	FollowerCount  int64  `json:"follower_count"`
	PinCount       int64  `json:"pin_count"`
	WebsiteUrl     string `json:"website_url"`
	Username       string `json:"username"`
	AccountType    string `json:"account_type"`
	BoardCount     int64  `json:"board_count"`
	MonthlyViews   int64  `json:"monthly_views"`
	Id             string `json:"id"`
	ProfileImage   string `json:"profile_image"`
}

func (service *Service) GetUserAccount() (*UserAccount, *errortools.Error) {
	var userAccount UserAccount

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodGet,
		Url:           service.url("user_account"),
		ResponseModel: &userAccount,
	}

	_, _, e := service.httpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &userAccount, nil
}
