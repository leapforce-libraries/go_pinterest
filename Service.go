package pinterest

import (
	"fmt"
	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	oauth2 "github.com/leapforce-libraries/go_oauth2"
	"github.com/leapforce-libraries/go_oauth2/tokensource"
	"net/http"
)

const (
	apiName            string = "Pinterest"
	apiUrl             string = "https://api.pinterest.com/v5"
	defaultRedirectUrl string = "http://localhost:8080/oauth/redirect"
	authUrl            string = "https://www.pinterest.com/oauth"
	tokenUrl           string = "https://api.pinterest.com/v5/oauth/token"
	tokenHttpMethod    string = http.MethodPost
)

type Service struct {
	clientId      string
	oAuth2Service *oauth2.Service
	redirectUrl   *string
	errorResponse *ErrorResponse
}

type ServiceConfig struct {
	ClientId     string
	ClientSecret string
	TokenSource  tokensource.TokenSource
	RedirectUrl  *string
}

func NewService(cfg *ServiceConfig) (*Service, *errortools.Error) {
	if cfg == nil {
		return nil, errortools.ErrorMessage("ServiceConfig must not be a nil pointer")
	}

	if cfg.ClientId == "" {
		return nil, errortools.ErrorMessage("ClientId not provided")
	}

	redirectUrl := defaultRedirectUrl
	if cfg.RedirectUrl != nil {
		redirectUrl = *cfg.RedirectUrl
	}

	oauth2ServiceConfig := oauth2.ServiceConfig{
		ClientId:        cfg.ClientId,
		ClientSecret:    cfg.ClientSecret,
		RedirectUrl:     redirectUrl,
		AuthUrl:         authUrl,
		TokenUrl:        tokenUrl,
		TokenHttpMethod: tokenHttpMethod,
		TokenSource:     cfg.TokenSource,
	}
	oauth2Service, e := oauth2.NewService(&oauth2ServiceConfig)
	if e != nil {
		return nil, e
	}

	return &Service{
		clientId:      cfg.ClientId,
		oAuth2Service: oauth2Service,
		redirectUrl:   cfg.RedirectUrl,
	}, nil
}

func (service *Service) httpRequest(requestConfig *go_http.RequestConfig) (*http.Request, *http.Response, *errortools.Error) {
	// add error model
	service.errorResponse = &ErrorResponse{}
	requestConfig.ErrorModel = service.errorResponse

	request, response, e := service.oAuth2Service.HttpRequest(requestConfig)
	if e != nil {
		if service.errorResponse.Message != "" {
			e.SetMessage(service.errorResponse.Message)
		}
	}

	if e != nil {
		return request, response, e
	}

	return request, response, nil
}

func (service *Service) AuthorizeUrl(scope string) string {
	if service.redirectUrl == nil {
		return ""
	}
	return fmt.Sprintf("%s?client_id=%s&redirect_uri=%s&response_type=code&scope=%s", authUrl, service.clientId, *service.redirectUrl, scope)
}

func (service *Service) GetTokenFromCode(r *http.Request) *errortools.Error {
	return service.oAuth2Service.GetTokenFromCode(r, nil)
}

func (service *Service) url(path string) string {
	return fmt.Sprintf("%s/%s", apiUrl, path)
}

func (service *Service) ApiName() string {
	return apiName
}

func (service *Service) ApiKey() string {
	return service.clientId
}

func (service *Service) ApiCallCount() int64 {
	return service.oAuth2Service.ApiCallCount()
}

func (service *Service) ApiReset() {
	service.oAuth2Service.ApiReset()
}

func (service *Service) ErrorResponse() *ErrorResponse {
	return service.errorResponse
}
