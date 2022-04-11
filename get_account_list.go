package twentyfour

import (
	"encoding/xml"
	"net/http"
	"net/url"

	"github.com/omniboost/go-24sevenoffice/utils"
	"github.com/pkg/errors"
)

func (c *Client) NewGetAccountListRequest() GetAccountListRequest {
	r := GetAccountListRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type GetAccountListRequest struct {
	client      *Client
	queryParams *GetAccountListRequestQueryParams
	pathParams  *GetAccountListRequestPathParams
	method      string
	headers     http.Header
	requestBody GetAccountListRequestBody
}

func (r GetAccountListRequest) SOAPAction() string {
	return "http://24sevenOffice.com/webservices/GetAccountList"
}

func (r GetAccountListRequest) NewQueryParams() *GetAccountListRequestQueryParams {
	return &GetAccountListRequestQueryParams{}
}

type GetAccountListRequestQueryParams struct {
}

func (p GetAccountListRequestQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetAccountListRequest) QueryParams() QueryParams {
	return r.queryParams
}

func (r GetAccountListRequest) NewPathParams() *GetAccountListRequestPathParams {
	return &GetAccountListRequestPathParams{}
}

type GetAccountListRequestPathParams struct {
}

func (p *GetAccountListRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetAccountListRequest) PathParams() PathParams {
	return r.pathParams
}

func (r *GetAccountListRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetAccountListRequest) Method() string {
	return r.method
}

func (r GetAccountListRequest) NewRequestBody() GetAccountListRequestBody {
	return GetAccountListRequestBody{}
}

type GetAccountListRequestBody struct {
	XMLName xml.Name `xml:"GetAccountList"`
}

func (r *GetAccountListRequest) RequestBody() *GetAccountListRequestBody {
	return &r.requestBody
}

func (r *GetAccountListRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *GetAccountListRequest) SetRequestBody(body GetAccountListRequestBody) {
	r.requestBody = body
}

func (r *GetAccountListRequest) NewResponseBody() *GetAccountListRequestResponseBody {
	return &GetAccountListRequestResponseBody{}
}

type GetAccountListRequestResponseBody struct {
	XMLName xml.Name `xml:"GetAccountListResponse"`

	GetAccountListResult struct {
		AccountData []struct {
			AccountId   string `xml:"AccountId"`
			AccountNo   string `xml:"AccountNo"`
			AccountName string `xml:"AccountName"`
			AccountTax  string `xml:"AccountTax"`
			TaxNo       string `xml:"TaxNo"`
		} `xml:"AccountData"`
	} `xml:"GetAccountListResult"`
}

func (r *GetAccountListRequest) URL() (*url.URL, error) {
	u, err := r.client.GetEndpointURL("/Economy/Account/V004/Accountservice.asmx", r.PathParams())
	return &u, err
}

func (r *GetAccountListRequest) Do() (GetAccountListRequestResponseBody, error) {
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
