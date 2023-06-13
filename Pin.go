package pinterest

import (
	"fmt"
	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	"github.com/leapforce-libraries/go_pinterest/types"
	"net/http"
	"net/url"
	"strings"
)

type Pin struct {
	Id              string               `json:"id"`
	CreatedAt       types.DateTimeString `json:"created_at"`
	Link            string               `json:"link"`
	Title           string               `json:"title"`
	Description     string               `json:"description"`
	DominantColor   string               `json:"dominant_color"`
	AltText         string               `json:"alt_text"`
	CreativeType    string               `json:"creative_type"`
	BoardId         string               `json:"board_id"`
	BoardSectionId  string               `json:"board_section_id"`
	BoardOwner      Owner                `json:"board_owner"`
	IsOwner         bool                 `json:"is_owner"`
	Media           PinMedia             `json:"media"`
	ParentPinId     string               `json:"parent_pin_id"`
	IsStandard      bool                 `json:"is_standard"`
	HasBeenPromoted bool                 `json:"has_been_promoted"`
	Note            string               `json:"note"`
}

type PinMedia struct {
	MediaType string `json:"media_type"`
}

type ListPinsConfig struct {
	BoardId       string
	AdAccountId   *string
	PageSize      *int64
	CreativeTypes *[]string
}

type ListPinsResponse struct {
	Items    []Pin
	Bookmark string
}

func (service *Service) ListPins(cfg *ListPinsConfig) (*[]Pin, *errortools.Error) {
	if cfg == nil {
		return nil, errortools.ErrorMessage("ListPinsConfig must not be bil")
	}

	var pins []Pin

	var values = url.Values{}
	if cfg.AdAccountId != nil {
		values.Set("ad_account_id", *cfg.AdAccountId)
	}
	if cfg.PageSize != nil {
		values.Set("page_size", fmt.Sprintf("%v", *cfg.PageSize))
	}
	if cfg.CreativeTypes != nil {
		values.Set("creative_types", strings.Join(*cfg.CreativeTypes, ","))
	}

	for {
		var response ListPinsResponse

		requestConfig := go_http.RequestConfig{
			Method:        http.MethodGet,
			Url:           service.url(fmt.Sprintf("boards/%s/pins?%s", cfg.BoardId, values.Encode())),
			ResponseModel: &response,
		}

		_, _, e := service.httpRequest(&requestConfig)
		if e != nil {
			return nil, e
		}

		pins = append(pins, response.Items...)

		if response.Bookmark == "" {
			break
		}

		values.Set("bookmark", response.Bookmark)
	}

	return &pins, nil
}
