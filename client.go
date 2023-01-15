package kDrive

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
)

const (
	apiURL        = "https://api.infomaniak.com/"
	appURL        = "https://kdrive.infomaniak.com/"
	apiVersion    = "2"
	kDriveVersion = "2022-06-28"
)

type Token string

func (it Token) String() string {
	return string(it)
}

type DriveId string

func (it DriveId) String() string {
	return string(it)
}

// ClientOption to configure API client
type ClientOption func(*Client)

type Client struct {
	httpClient    *http.Client
	baseUrl       *url.URL
	appUrl        *url.URL
	apiVersion    string
	kDriveVersion string
	DriveId       DriveId

	Token Token

	//Activity ActivityService
	File FileService
	//HtmlPage    HtmlPageService
	//Invitation  InvitationsService
	Link LinkService
	//Setting     SettingsService
	//Statistic   StatisticsService
	//User        UsersService
}

func NewClient(driveId DriveId, token Token, opts ...ClientOption) *Client {
	apiUrl, err := url.Parse(apiURL)
	if err != nil {
		panic(err)
	}
	appUrl, err := url.Parse(appURL)
	if err != nil {
		panic(err)
	}
	c := &Client{
		httpClient:    http.DefaultClient,
		Token:         token,
		baseUrl:       apiUrl,
		appUrl:        appUrl,
		DriveId:       driveId,
		apiVersion:    apiVersion,
		kDriveVersion: kDriveVersion,
	}

	//c.Activity = &ActivityClient{apiClient: c}
	c.File = &FileClient{apiClient: c}
	//c.HtmlPage = &HtmlPageClient{apiClient: c}
	//c.Invitations = &InvitationClient{apiClient: c}
	c.Link = &LinkClient{apiClient: c}
	//c.Setting = &SettingClient{apiClient: c}
	//c.Statistic = &StatisticClient{apiClient: c}
	//c.User = &UserClient{apiClient: c}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

// WithHTTPClient overrides the default http.Client
func WithHTTPClient(client *http.Client) ClientOption {
	return func(c *Client) {
		c.httpClient = client
	}
}

// WithVersion overrides the kDrive API version
func WithVersion(version string) ClientOption {
	return func(c *Client) {
		c.kDriveVersion = version
	}
}

func (c *Client) request(ctx context.Context, method string, urlStr string, queryParams map[string]string, requestBody interface{}) (*http.Response, error) {

	if queryParams["per_page"] != "" {
		perPage, err := strconv.Atoi(queryParams["per_page"])
		if err != err {
			return nil, err
		}
		if perPage < 1 || perPage > 1000 {
			return nil, errors.New("query 'per_page' must be between 1 and 1000")
		}
	}

	u, err := c.baseUrl.Parse(fmt.Sprintf("%s/drive/%s/%s/", c.apiVersion, c.DriveId, urlStr))
	if err != nil {
		return nil, err
	}

	var buf io.ReadWriter
	if requestBody != nil && !reflect.ValueOf(&requestBody).IsNil() {
		body, err := json.Marshal(requestBody)
		if err != nil {
			return nil, err
		}
		buf = bytes.NewBuffer(body)
	}

	if len(queryParams) > 0 {
		q := u.Query()
		for k, v := range queryParams {
			q.Add(k, v)
		}
		u.RawQuery = q.Encode()
	}
	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.Token.String()))
	req.Header.Add("kDrive-Version", c.kDriveVersion)
	req.Header.Add("Content-Type", "application/json")

	var res *http.Response
	res, err = c.httpClient.Do(req.WithContext(ctx))
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		var apiErr Error
		err = json.NewDecoder(res.Body).Decode(&apiErr)
		apiErr.ErrorResult.Code = res.Status
		apiErr.ErrorResult.Description += fmt.Sprintf(" -> %s", req.URL.String())
		if err != nil {
			return nil, err
		}
		return nil, &apiErr
	}
	return res, nil
}

type Pagination struct {
	StartCursor Cursor
	PageSize    int
}

func (p *Pagination) ToQuery() map[string]string {
	if p == nil {
		return nil
	}
	r := map[string]string{}
	if p.StartCursor != "" {
		r["start_cursor"] = p.StartCursor.String()
	}

	if p.PageSize != 0 {
		r["page_size"] = strconv.Itoa(p.PageSize)
	}

	return r
}
