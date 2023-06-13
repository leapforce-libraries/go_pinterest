package pinterest

import (
	"fmt"
	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	"github.com/leapforce-libraries/go_pinterest/types"
	"net/http"
	"net/url"
)

type Board struct {
	Id                  string               `json:"id"`
	CreatedAt           types.DateTimeString `json:"created_at"`
	BoardPinsModifiedAt types.DateTimeString `json:"board_pins_modified_at"`
	Name                string               `json:"name"`
	Description         string               `json:"description"`
	CollaboratorCount   int64                `json:"collaborator_count"`
	PinCount            int64                `json:"pin_count"`
	FollowerCount       int64                `json:"follower_count"`
	Media               BoardMedia           `json:"media"`
	Owner               Owner                `json:"owner"`
	Privacy             string               `json:"privacy"`
}

type BoardMedia struct {
	ImageCoverUrl    string   `json:"image_cover_url"`
	PinThumbnailUrls []string `json:"pin_thumbnail_urls"`
}
type ListBoardsConfig struct {
	AdAccountId *string
	PageSize    *int64
	Privacy     *string
}

type ListBoardsResponse struct {
	Items    []Board
	Bookmark string
}

func (service *Service) ListBoards(cfg *ListBoardsConfig) (*[]Board, *errortools.Error) {
	var boards []Board

	var values = url.Values{}
	if cfg != nil {
		if cfg.AdAccountId != nil {
			values.Set("ad_account_id", *cfg.AdAccountId)
		}
		if cfg.PageSize != nil {
			values.Set("page_size", fmt.Sprintf("%v", *cfg.PageSize))
		}
		if cfg.Privacy != nil {
			values.Set("privacy", *cfg.Privacy)
		}
	}

	for {
		var response ListBoardsResponse

		requestConfig := go_http.RequestConfig{
			Method:        http.MethodGet,
			Url:           service.url(fmt.Sprintf("boards?%s", values.Encode())),
			ResponseModel: &response,
		}

		_, _, e := service.httpRequest(&requestConfig)
		if e != nil {
			return nil, e
		}

		boards = append(boards, response.Items...)

		if response.Bookmark == "" {
			break
		}

		values.Set("bookmark", response.Bookmark)
	}

	return &boards, nil
}
