package twentyfour

import (
	"encoding/xml"
	"net/http"
	"net/url"

	"github.com/omniboost/go-24sevenoffice/utils"
	"github.com/pkg/errors"
)

func (c *Client) NewGetTypeGroupListRequest() GetTypeGroupListRequest {
	r := GetTypeGroupListRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type GetTypeGroupListRequest struct {
	client      *Client
	queryParams *GetTypeGroupListRequestQueryParams
	pathParams  *GetTypeGroupListRequestPathParams
	method      string
	headers     http.Header
	requestBody GetTypeGroupListRequestBody
}

func (r GetTypeGroupListRequest) SOAPAction() string {
	return "http://24sevenOffice.com/webservices/GetTypeGroupList"
}

func (r GetTypeGroupListRequest) NewQueryParams() *GetTypeGroupListRequestQueryParams {
	return &GetTypeGroupListRequestQueryParams{}
}

type GetTypeGroupListRequestQueryParams struct {
}

func (p GetTypeGroupListRequestQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetTypeGroupListRequest) QueryParams() QueryParams {
	return r.queryParams
}

func (r GetTypeGroupListRequest) NewPathParams() *GetTypeGroupListRequestPathParams {
	return &GetTypeGroupListRequestPathParams{}
}

type GetTypeGroupListRequestPathParams struct {
}

func (p *GetTypeGroupListRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetTypeGroupListRequest) PathParams() PathParams {
	return r.pathParams
}

func (r *GetTypeGroupListRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetTypeGroupListRequest) Method() string {
	return r.method
}

func (r GetTypeGroupListRequest) NewRequestBody() GetTypeGroupListRequestBody {
	return GetTypeGroupListRequestBody{}
}

type GetTypeGroupListRequestBody struct {
	XMLName xml.Name `xml:"GetTypeGroupList"`
}

func (r *GetTypeGroupListRequest) RequestBody() *GetTypeGroupListRequestBody {
	return &r.requestBody
}

func (r *GetTypeGroupListRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *GetTypeGroupListRequest) SetRequestBody(body GetTypeGroupListRequestBody) {
	r.requestBody = body
}

func (r *GetTypeGroupListRequest) NewResponseBody() *GetTypeGroupListRequestResponseBody {
	return &GetTypeGroupListRequestResponseBody{}
}

type GetTypeGroupListRequestResponseBody struct {
	XMLName xml.Name `xml:"GetTypeGroupListResponse"`

	GetTypeGroupListResult struct {
		AccountingGroup []struct {
			AccountId string `xml:"AccountId"`
			ID        string `xml:"Id"`
			Name      string `xml:"Name"`
		} `xml:"AccountingGroup"`
	} `xml:"GetTypeGroupListResult"`
}

func (r *GetTypeGroupListRequest) URL() (*url.URL, error) {
	u, err := r.client.GetEndpointURL("/Client/V001/ClientService.asmx", r.PathParams())
	return &u, err
}

func (r *GetTypeGroupListRequest) Do() (GetTypeGroupListRequestResponseBody, error) {
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
