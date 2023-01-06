package kDrive

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"time"
)

const (
	apiURL        = "https://api.infomaniak.com/"
	apiVersion    = "2"
	kDriveVersion = "2022-06-28"
	maxRetries    = 3
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
	apiVersion    string
	kDriveVersion string
	driveId       DriveId

	maxRetries int

	Token Token

	Activity ActivityService
	//Files       FilesService
	//HtmlPage    HtmlPageService
	//Invitations InvitationsService
	//SharedLink  SharedLinkService
	//Settings    SettingsService
	//Statistics  StatisticsService
	//Users       UsersService
}

func NewClient(token Token, opts ...ClientOption) *Client {
	u, err := url.Parse(apiURL)
	if err != nil {
		panic(err)
	}
	c := &Client{
		httpClient:    http.DefaultClient,
		Token:         token,
		baseUrl:       u,
		apiVersion:    apiVersion,
		kDriveVersion: kDriveVersion,
		maxRetries:    maxRetries,
	}

	c.Activity = &ActivityClient{apiClient: c}
	//c.Files = &FilesClient{apiClient: c}
	//c.HtmlPage = &HtmlPageClient{apiClient: c}
	//c.Invitations = &InvitationsClient{apiClient: c}
	//c.SharedLink = &SharedLinkClient{apiClient: c}
	//c.Settings = &SettingsClient{apiClient: c}
	//c.Statistics = &StatisticsClient{apiClient: c}
	//c.Users = &UsersClient{apiClient: c}

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

// WithRetry overrides the default number of max retry attempts on 429 errors
func WithRetry(retries int) ClientOption {
	return func(c *Client) {
		c.maxRetries = retries
	}
}

func (c *Client) request(ctx context.Context, method string, urlStr string, queryParams map[string]string, requestBody interface{}) (*http.Response, error) {
	u, err := c.baseUrl.Parse(fmt.Sprintf("%s/%s", c.apiVersion, urlStr))
	if err != nil {
		return nil, err
	}

	var buf io.ReadWriter
	if requestBody != nil && !reflect.ValueOf(requestBody).IsNil() {
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

	failedAttempts := 0
	var res *http.Response
	for {
		var err error
		res, err = c.httpClient.Do(req.WithContext(ctx))
		if err != nil {
			return nil, err
		}

		if res.StatusCode != http.StatusTooManyRequests {
			break
		}

		failedAttempts++
		if failedAttempts == c.maxRetries {
			return nil, &RateLimitedError{Message: fmt.Sprintf("Retry request with 429 response failed after %d retries", failedAttempts)}
		}

		retryAfterHeader := res.Header["Retry-After"]
		if len(retryAfterHeader) == 0 {
			return nil, &RateLimitedError{Message: "Retry-After header missing from kDrive API response headers for 429 response"}
		}
		retryAfter := retryAfterHeader[0]

		waitSeconds, err := strconv.Atoi(retryAfter)
		if err != nil {
			break // should not happen
		}
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-time.After(time.Duration(waitSeconds) * time.Second):
		}
	}

	if res.StatusCode != http.StatusOK {
		var apiErr Error
		err = json.NewDecoder(res.Body).Decode(&apiErr)
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
