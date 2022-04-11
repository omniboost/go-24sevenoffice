package twentyfour

import (
	"encoding/xml"
	"net/http"
	"net/url"

	"github.com/omniboost/go-24sevenoffice/utils"
	"github.com/pkg/errors"
)

func (c *Client) NewGetTaxCodeListRequest() GetTaxCodeListRequest {
	r := GetTaxCodeListRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type GetTaxCodeListRequest struct {
	client      *Client
	queryParams *GetTaxCodeListRequestQueryParams
	pathParams  *GetTaxCodeListRequestPathParams
	method      string
	headers     http.Header
	requestBody GetTaxCodeListRequestBody
}

func (r GetTaxCodeListRequest) SOAPAction() string {
	return "http://24sevenOffice.com/webservices/GetTaxCodeList"
}

func (r GetTaxCodeListRequest) NewQueryParams() *GetTaxCodeListRequestQueryParams {
	return &GetTaxCodeListRequestQueryParams{}
}

type GetTaxCodeListRequestQueryParams struct {
}

func (p GetTaxCodeListRequestQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetTaxCodeListRequest) QueryParams() QueryParams {
	return r.queryParams
}

func (r GetTaxCodeListRequest) NewPathParams() *GetTaxCodeListRequestPathParams {
	return &GetTaxCodeListRequestPathParams{}
}

type GetTaxCodeListRequestPathParams struct {
}

func (p *GetTaxCodeListRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetTaxCodeListRequest) PathParams() PathParams {
	return r.pathParams
}

func (r *GetTaxCodeListRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetTaxCodeListRequest) Method() string {
	return r.method
}

func (r GetTaxCodeListRequest) NewRequestBody() GetTaxCodeListRequestBody {
	return GetTaxCodeListRequestBody{}
}

type GetTaxCodeListRequestBody struct {
	XMLName xml.Name `xml:"GetTaxCodeList"`
}

func (r *GetTaxCodeListRequest) RequestBody() *GetTaxCodeListRequestBody {
	return &r.requestBody
}

func (r *GetTaxCodeListRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *GetTaxCodeListRequest) SetRequestBody(body GetTaxCodeListRequestBody) {
	r.requestBody = body
}

func (r *GetTaxCodeListRequest) NewResponseBody() *GetTaxCodeListRequestResponseBody {
	return &GetTaxCodeListRequestResponseBody{}
}

type GetTaxCodeListRequestResponseBody struct {
	XMLName xml.Name `xml:"GetTaxCodeListResponse"`

	GetTaxCodeListResult struct {
		Text           string `xml:",chardata"`
		TaxCodeElement []struct {
			TaxId     int     `xml:"TaxId"`
			TaxNo     string  `xml:"TaxNo"`
			TaxName   string  `xml:"TaxName"`
			TaxRate   float64 `xml:"TaxRate"`
			AccountNo string  `xml:"AccountNo"`
		} `xml:"TaxCodeElement"`
	} `xml:"GetTaxCodeListResult"`
}

func (r *GetTaxCodeListRequest) URL() (*url.URL, error) {
	u, err := r.client.GetEndpointURL("/Economy/Account/V004/Accountservice.asmx", r.PathParams())
	return &u, err
}

func (r *GetTaxCodeListRequest) Do() (GetTaxCodeListRequestResponseBody, error) {
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
