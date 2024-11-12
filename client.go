package odesli

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const (
	APIPath   = "https://api.song.link/v1-alpha.1"
	LinksPath = APIPath + "/links"
)

const DefaultHTTPTimeout = 60 * time.Second

type GetLinksRequest struct {

	// The unique identifier of the streaming entity, e.g.`1443109064` which is an
	// iTunes ID.If `url` is not provided, you must provide `platform`, `type` and
	// `id`.
	ID string

	// The URL of a valid song or album from any of our supported platforms. It is
	// safest to encode the URL, e.g. with `encodeURIComponent()`.
	URL string

	// Two-letter country code. Specifies the country/location we use when searching
	// streaming catalogs. Optional. Defaults to `US`.
	UserCountry string

	// The platform of the entity you'd like to match. See above section for
	// supported platforms. If `url` is not provided, you must provide `platform`,
	// `type` and `id`.
	Platform Platform

	// The type of streaming entity. We support `song` and `album`. If `url` is not
	// provided, you must provide `platform`, `type` and `id`.
	Type EntityType
}

func (r GetLinksRequest) GetURLValues() url.Values {
	v := url.Values{}
	if r.ID != "" {
		v.Set("id", r.ID)
	}
	if r.URL != "" {
		v.Set("url", r.URL)
	}
	if r.UserCountry != "" {
		v.Set("userCountry", r.UserCountry)
	}
	if r.Platform != "" {
		v.Set("platform", string(r.Platform))
	}
	if r.Type != "" {
		v.Set("type", string(r.Type))
	}
	return v
}

type GetLinksResponse struct {

	// The unique ID for the input entity that was supplied in the request. The data
	// for this entity, such as title, artistName, etc. will be found in an object at
	// `nodesByUniqueId[entityUniqueId]`
	EntityUniqueID string `json:"entityUniqueId,omitempty"`

	// The userCountry query param that was supplied in the request. It signals
	// the country/availability we use to query the streaming platforms. Defaults
	// to 'US' if no userCountry supplied in the request.
	//
	// NOTE: As a fallback, our service may respond with matches that were found in a
	// locale other than the userCountry supplied
	UserCountry string `json:"userCountry,omitempty"`

	// A URL that will render the Songlink page for this entity
	PageUrl string `json:"pageUrl,omitempty"`

	// A collection of objects. Each key is a platform, and each value is an
	// object that contains data for linking to the match
	LinksByPlatform map[Platform]LinkByPlatform

	// A collection of objects. Each key is a unique identifier for a streaming
	// entity, and each value is an object that contains data for that entity, such
	// as `title`, `artistName`, `thumbnailUrl`, etc.
	EntitiesByUniqueId map[string]Entity `json:"entitiesByUniqueId,omitempty"`
}

type Entity struct {

	// This is the unique identifier on the streaming platform/API provider
	Id string `json:"id,omitempty"`

	Type EntityType `json:"type,omitempty"`

	Title string `json:"title,omitempty"`

	ArtistName string `json:"artistName,omitempty"`

	ThumbnailUrl string `json:"thumbnailUrl,omitempty"`

	ThumbnailWidth int `json:"thumbnailWidth,omitempty"`

	ThumbnailHeight int `json:"thumbnailHeight,omitempty"`

	// The API provider that powered this match. Useful if you'd like to use
	// this entity's data to query the API directly
	ApiProvider APIProvider `json:"apiProvider,omitempty"`

	// An array of platforms that are "powered" by this entity. E.g. an entity
	// from Apple Music will generally have a `platforms` array of
	// `["appleMusic", "itunes"]` since both those platforms/links are derived
	// from this single entity
	Platforms []Platform `json:"platforms,omitempty"`
}

func (e *Entity) UnmarshalJSON(data []byte) error {
	type Alias Entity
	aux := &struct {
		Id interface{} `json:"id"` // unmarshal as string or number as priority field for id
		*Alias
	}{
		Alias: (*Alias)(e),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	switch id := aux.Id.(type) {
	case string:
		e.Id = id
	case int:
		e.Id = strconv.Itoa(id)
	case int64:
		e.Id = strconv.FormatInt(id, 10)
	case float32:
		e.Id = strconv.FormatFloat(float64(id), 'f', -1, 64)
	case float64:
		e.Id = strconv.FormatFloat(id, 'f', -1, 64)
	default:
		return fmt.Errorf("unexpected type for id: %T", id)
	}

	return nil
}

type LinkByPlatform struct {

	// The unique ID for this entity. Use it to look up data about this entity
	// at `entitiesByUniqueId[entityUniqueId]`
	EntityUniqueId string `json:"entityUniqueId"`

	// The URL for this match
	Url string `json:"url"`

	// The native app URI that can be used on mobile devices to open this
	// entity directly in the native app
	NativeAppUriMobile string `json:"nativeAppUriMobile"`

	// The native app URI that can be used on desktop devices to open this
	// entity directly in the native app
	NativeAppUriDesktop string `json:"nativeAppUriDesktop"`
}

type API interface {
	GetLinks(ctx context.Context, req GetLinksRequest) (GetLinksResponse, error)
}

type ClientOption struct {
	APIToken string
	Debug    bool
}

type Client struct {
	client *http.Client
	debug  bool
}

func NewClient(opt ClientOption) (*Client, error) {
	var rt http.RoundTripper
	t := &http.Transport{}
	if opt.APIToken != "" {
		rt = TransportWithAPIToken(t, opt.APIToken)
	} else {
		rt = t
	}
	return &Client{
		client: &http.Client{
			Transport: rt,
			Timeout:   DefaultHTTPTimeout,
		},
		debug: opt.Debug,
	}, nil
}

func (c *Client) GetLinks(ctx context.Context, r GetLinksRequest) (GetLinksResponse, error) {
	path := fmt.Sprint(LinksPath, "?", r.GetURLValues().Encode())
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		path,
		nil,
	)
	if err != nil {
		return GetLinksResponse{}, err
	}
	if c.debug {
		log.Printf("request: %s", req.URL.String())
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return GetLinksResponse{}, err
	}
	defer func() { _ = resp.Body.Close() }()
	err = checkResponse(resp)
	if err != nil {
		return GetLinksResponse{}, err
	}
	res := GetLinksResponse{}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return GetLinksResponse{}, err
	}
	if c.debug {
		log.Printf("body: %s", string(body))
	}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return GetLinksResponse{}, err
	}
	return res, nil
}

func checkResponse(resp *http.Response) error {
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return nil
	}
	return fmt.Errorf("unexpected response: %s", resp.Status)
}
