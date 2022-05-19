package twentyfour

import (
	"encoding/xml"
	"net/http"
	"net/url"

	"github.com/omniboost/go-24sevenoffice/utils"
	"github.com/pkg/errors"
)

func (c *Client) NewGetCompaniesRequest() GetCompaniesRequest {
	r := GetCompaniesRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type GetCompaniesRequest struct {
	client      *Client
	queryParams *GetCompaniesRequestQueryParams
	pathParams  *GetCompaniesRequestPathParams
	method      string
	headers     http.Header
	requestBody GetCompaniesRequestBody
}

func (r GetCompaniesRequest) SOAPAction() string {
	return "http://24sevenOffice.com/webservices/GetCompanies"
}

func (r GetCompaniesRequest) NewQueryParams() *GetCompaniesRequestQueryParams {
	return &GetCompaniesRequestQueryParams{}
}

type GetCompaniesRequestQueryParams struct {
}

func (p GetCompaniesRequestQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetCompaniesRequest) QueryParams() QueryParams {
	return r.queryParams
}

func (r GetCompaniesRequest) NewPathParams() *GetCompaniesRequestPathParams {
	return &GetCompaniesRequestPathParams{}
}

type GetCompaniesRequestPathParams struct {
}

func (p *GetCompaniesRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetCompaniesRequest) PathParams() PathParams {
	return r.pathParams
}

func (r *GetCompaniesRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetCompaniesRequest) Method() string {
	return r.method
}

func (r GetCompaniesRequest) NewRequestBody() GetCompaniesRequestBody {
	return GetCompaniesRequestBody{}
}

type GetCompaniesRequestBody struct {
	XMLName xml.Name `xml:"web:GetCompanies"`

	SearchParams struct {
		ExternalID string `xml:"web:ExternalId,omitempty"`
		CompanyID  int    `xml:"web:CompanyId,omitempty"`
		// CompanyIDs         []int    `xml:"web:CompanyIds>web:int,omitempty"`
		CompanyName        string   `xml:"web:CompanyName,omitempty"`
		ChangedAfter       DateTime `xml:"web:ChangedAfter,omitempty"`
		CompanyEmail       string   `xml:"web:CompanyEmail,omitempty"`
		CompanyPhone       string   `xml:"web:CompanyPhone,omitempty"`
		OrganizationNumber string   `xml:"web:OrganizationNumber,omitempty"`
	} `xml:"web:searchParams"`
	// CompanyReturnProperties struct {
	// 	String []string `xml:"string"`
	// } `xml:"invoiceReturnProperties"`
	ReturnProperties []string `xml:"web:returnProperties>web:string"`
}

func (r *GetCompaniesRequest) RequestBody() *GetCompaniesRequestBody {
	return &r.requestBody
}

func (r *GetCompaniesRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *GetCompaniesRequest) SetRequestBody(body GetCompaniesRequestBody) {
	r.requestBody = body
}

func (r *GetCompaniesRequest) NewResponseBody() *GetCompaniesRequestResponseBody {
	return &GetCompaniesRequestResponseBody{}
}

type GetCompaniesRequestResponseBody struct {
	XMLName xml.Name `xml:"GetCompaniesResponse"`

	Companies Companies `xml:"GetCompaniesResult>Company"`
}

func (r *GetCompaniesRequest) URL() (*url.URL, error) {
	u, err := r.client.GetEndpointURL("/CRM/Company/V001/CompanyService.asmx", r.PathParams())
	return &u, err
}

func (r *GetCompaniesRequest) Do() (GetCompaniesRequestResponseBody, error) {
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
