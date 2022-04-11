package twentyfour

import (
	"encoding/xml"
	"net/http"
	"net/url"

	"github.com/omniboost/go-24sevenoffice/utils"
	"github.com/pkg/errors"
)

func (c *Client) NewGetTypeListRequest() GetTypeListRequest {
	r := GetTypeListRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type GetTypeListRequest struct {
	client      *Client
	queryParams *GetTypeListRequestQueryParams
	pathParams  *GetTypeListRequestPathParams
	method      string
	headers     http.Header
	requestBody GetTypeListRequestBody
}

func (r GetTypeListRequest) SOAPAction() string {
	return "http://24sevenOffice.com/webservices/GetTypeList"
}

func (r GetTypeListRequest) NewQueryParams() *GetTypeListRequestQueryParams {
	return &GetTypeListRequestQueryParams{}
}

type GetTypeListRequestQueryParams struct {
}

func (p GetTypeListRequestQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetTypeListRequest) QueryParams() QueryParams {
	return r.queryParams
}

func (r GetTypeListRequest) NewPathParams() *GetTypeListRequestPathParams {
	return &GetTypeListRequestPathParams{}
}

type GetTypeListRequestPathParams struct {
}

func (p *GetTypeListRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetTypeListRequest) PathParams() PathParams {
	return r.pathParams
}

func (r *GetTypeListRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetTypeListRequest) Method() string {
	return r.method
}

func (r GetTypeListRequest) NewRequestBody() GetTypeListRequestBody {
	return GetTypeListRequestBody{}
}

type GetTypeListRequestBody struct {
	XMLName xml.Name `xml:"GetTypeList"`
}

func (r *GetTypeListRequest) RequestBody() *GetTypeListRequestBody {
	return &r.requestBody
}

func (r *GetTypeListRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *GetTypeListRequest) SetRequestBody(body GetTypeListRequestBody) {
	r.requestBody = body
}

func (r *GetTypeListRequest) NewResponseBody() *GetTypeListRequestResponseBody {
	return &GetTypeListRequestResponseBody{}
}

type GetTypeListRequestResponseBody struct {
	XMLName xml.Name `xml:"GetTypeListResponse"`

	GetTypeListResult struct {
		Text     string `xml:",chardata"`
		TypeData []struct {
			Text          string `xml:",chardata"`
			TypeId        string `xml:"TypeId"`
			Title         string `xml:"Title"`
			EntrySeriesId string `xml:"EntrySeriesId"`
			TypeNo        string `xml:"TypeNo"`
		} `xml:"TypeData"`
	} `xml:"GetTypeListResult"`
}

func (r *GetTypeListRequest) URL() (*url.URL, error) {
	u, err := r.client.GetEndpointURL("/Economy/Account/V004/Accountservice.asmx", r.PathParams())
	return &u, err
}

func (r *GetTypeListRequest) Do() (GetTypeListRequestResponseBody, error) {
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
