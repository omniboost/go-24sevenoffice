package twentyfour

import (
	"encoding/xml"
	"net/http"
	"net/url"

	"github.com/omniboost/go-24sevenoffice/utils"
	"github.com/pkg/errors"
)

func (c *Client) NewGetEntryIDRequest() GetEntryIDRequest {
	r := GetEntryIDRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type GetEntryIDRequest struct {
	client      *Client
	queryParams *GetEntryIDRequestQueryParams
	pathParams  *GetEntryIDRequestPathParams
	method      string
	headers     http.Header
	requestBody GetEntryIDRequestBody
}

func (r GetEntryIDRequest) SOAPAction() string {
	return "http://24sevenOffice.com/webservices/GetEntryId"
}

func (r GetEntryIDRequest) NewQueryParams() *GetEntryIDRequestQueryParams {
	return &GetEntryIDRequestQueryParams{}
}

type GetEntryIDRequestQueryParams struct {
}

func (p GetEntryIDRequestQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetEntryIDRequest) QueryParams() QueryParams {
	return r.queryParams
}

func (r GetEntryIDRequest) NewPathParams() *GetEntryIDRequestPathParams {
	return &GetEntryIDRequestPathParams{}
}

type GetEntryIDRequestPathParams struct {
}

func (p *GetEntryIDRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetEntryIDRequest) PathParams() PathParams {
	return r.pathParams
}

func (r *GetEntryIDRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetEntryIDRequest) Method() string {
	return r.method
}

func (r GetEntryIDRequest) NewRequestBody() GetEntryIDRequestBody {
	return GetEntryIDRequestBody{}
}

type GetEntryIDRequestBody struct {
	XMLName xml.Name `xml:"web:GetEntryId"`

	ArgEntryID struct {
		XMLName xml.Name `xml:"web:argEntryId"`
		Date    Date     `xml:"web:Date"`
		SortNo  int      `xml:"web:SortNo"`
		EntryNo int      `xml:"web:EntryNo"`
	} `xml:"web:argEntryId"`
}

func (r *GetEntryIDRequest) RequestBody() *GetEntryIDRequestBody {
	return &r.requestBody
}

func (r *GetEntryIDRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *GetEntryIDRequest) SetRequestBody(body GetEntryIDRequestBody) {
	r.requestBody = body
}

func (r *GetEntryIDRequest) NewResponseBody() *GetEntryIDRequestResponseBody {
	return &GetEntryIDRequestResponseBody{}
}

type GetEntryIDRequestResponseBody struct {
	XMLName xml.Name `xml:"GetEntryIdResponse"`

	GetEntryIdResult struct {
		Date    string `xml:"Date"`
		SortNo  string `xml:"SortNo"`
		EntryNo string `xml:"EntryNo"`
	} `xml:"GetEntryIdResult"`
}

func (r *GetEntryIDRequest) URL() (*url.URL, error) {
	u, err := r.client.GetEndpointURL("/Economy/Account/V004/Accountservice.asmx", r.PathParams())
	return &u, err
}

func (r *GetEntryIDRequest) Do() (GetEntryIDRequestResponseBody, error) {
	var err error

	// Create http request
	req, err := r.client.NewRequest(nil, r)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	err = r.client.InitSession(req)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	// Process query parameters
	err = utils.AddQueryParamsToRequest(r.QueryParams(), req, false)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	responseBody := r.NewResponseBody()
	_, err = r.client.Do(req, responseBody)
	if err != nil {
		return *responseBody, errors.WithStack(err)
	}

	return *responseBody, nil
}
