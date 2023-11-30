package twentyfour

import (
	"encoding/xml"
	"net/http"
	"net/url"

	"github.com/omniboost/go-24sevenoffice/utils"
	"github.com/pkg/errors"
)

func (c *Client) NewGetTransactionsRequest() GetTransactionsRequest {
	r := GetTransactionsRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type GetTransactionsRequest struct {
	client      *Client
	queryParams *GetTransactionsRequestQueryParams
	pathParams  *GetTransactionsRequestPathParams
	method      string
	headers     http.Header
	requestBody GetTransactionsRequestBody
}

func (r GetTransactionsRequest) SOAPAction() string {
	return "http://24sevenoffice.com/webservices/economy/accounting/GetTransactions"
}

func (r GetTransactionsRequest) NewQueryParams() *GetTransactionsRequestQueryParams {
	return &GetTransactionsRequestQueryParams{}
}

type GetTransactionsRequestQueryParams struct {
}

func (p GetTransactionsRequestQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetTransactionsRequest) QueryParams() QueryParams {
	return r.queryParams
}

func (r GetTransactionsRequest) NewPathParams() *GetTransactionsRequestPathParams {
	return &GetTransactionsRequestPathParams{}
}

type GetTransactionsRequestPathParams struct {
}

func (p *GetTransactionsRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetTransactionsRequest) PathParams() PathParams {
	return r.pathParams
}

func (r *GetTransactionsRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetTransactionsRequest) Method() string {
	return r.method
}

func (r GetTransactionsRequest) NewRequestBody() GetTransactionsRequestBody {
	return GetTransactionsRequestBody{
		Xmlns: "http://24sevenoffice.com/webservices/economy/accounting/",
	}
}

type GetTransactionsRequestBody struct {
	XMLName xml.Name `xml:"GetTransactions"`
	Xmlns   string   `xml:"xmlns,attr"`

	SearchParams struct {
		DateStart Date `xml:"DateStart,omitempty"`
		DateEnd   Date `xml:"DateEnd,omitempty"`
	} `xml:"searchParams"`
}

func (r *GetTransactionsRequest) RequestBody() *GetTransactionsRequestBody {
	return &r.requestBody
}

func (r *GetTransactionsRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *GetTransactionsRequest) SetRequestBody(body GetTransactionsRequestBody) {
	r.requestBody = body
}

func (r *GetTransactionsRequest) NewResponseBody() *GetTransactionsRequestResponseBody {
	return &GetTransactionsRequestResponseBody{}
}

type GetTransactionsRequestResponseBody struct {
	XMLName xml.Name `xml:"GetTransactionsResponse"`

	GetTransactionsResult struct {
	} `xml:"GetTransactionsResult"`
}

func (r *GetTransactionsRequest) URL() (*url.URL, error) {
	u, err := r.client.GetEndpointURL("/Economy/Accounting/V001/TransactionService.asmx", r.PathParams())
	return &u, err
}

func (r *GetTransactionsRequest) Do() (GetTransactionsRequestResponseBody, error) {
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
