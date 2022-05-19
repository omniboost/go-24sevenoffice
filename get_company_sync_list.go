package twentyfour

import (
	"encoding/xml"
	"net/http"
	"net/url"

	"github.com/omniboost/go-24sevenoffice/utils"
	"github.com/pkg/errors"
)

func (c *Client) NewGetCompanySyncListRequest() GetCompanySyncListRequest {
	r := GetCompanySyncListRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type GetCompanySyncListRequest struct {
	client      *Client
	queryParams *GetCompanySyncListRequestQueryParams
	pathParams  *GetCompanySyncListRequestPathParams
	method      string
	headers     http.Header
	requestBody GetCompanySyncListRequestBody
}

func (r GetCompanySyncListRequest) SOAPAction() string {
	return "http://24sevenOffice.com/webservices/GetCompanySyncList"
}

func (r GetCompanySyncListRequest) NewQueryParams() *GetCompanySyncListRequestQueryParams {
	return &GetCompanySyncListRequestQueryParams{}
}

type GetCompanySyncListRequestQueryParams struct {
}

func (p GetCompanySyncListRequestQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetCompanySyncListRequest) QueryParams() QueryParams {
	return r.queryParams
}

func (r GetCompanySyncListRequest) NewPathParams() *GetCompanySyncListRequestPathParams {
	return &GetCompanySyncListRequestPathParams{}
}

type GetCompanySyncListRequestPathParams struct {
}

func (p *GetCompanySyncListRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetCompanySyncListRequest) PathParams() PathParams {
	return r.pathParams
}

func (r *GetCompanySyncListRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetCompanySyncListRequest) Method() string {
	return r.method
}

func (r GetCompanySyncListRequest) NewRequestBody() GetCompanySyncListRequestBody {
	return GetCompanySyncListRequestBody{}
}

type GetCompanySyncListRequestBody struct {
	XMLName xml.Name `xml:"web:GetCompanySyncList"`

	SyncSearchParameters struct {
		Page         int      `xml:"web:Page,omitempty"`
		ChangedAfter DateTime `xml:"web:ChangedAfter,omitempty"`
	} `xml:"web:syncSearchParameters"`
}

func (r *GetCompanySyncListRequest) RequestBody() *GetCompanySyncListRequestBody {
	return &r.requestBody
}

func (r *GetCompanySyncListRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *GetCompanySyncListRequest) SetRequestBody(body GetCompanySyncListRequestBody) {
	r.requestBody = body
}

func (r *GetCompanySyncListRequest) NewResponseBody() *GetCompanySyncListRequestResponseBody {
	return &GetCompanySyncListRequestResponseBody{}
}

type GetCompanySyncListRequestResponseBody struct {
	XMLName xml.Name `xml:"GetCompanySyncListResponse"`

	GetCompanySyncListResult struct {
		Text         string `xml:",chardata"`
		CurrentPage  string `xml:"CurrentPage"`
		TotalPages   string `xml:"TotalPages"`
		TotalItems   string `xml:"TotalItems"`
		ItemsPerPage string `xml:"ItemsPerPage"`
		Items        struct {
			Text        string `xml:",chardata"`
			SyncCompany struct {
				Text           string `xml:",chardata"`
				CompanyId      string `xml:"CompanyId"`
				DateChanged    string `xml:"DateChanged"`
				DateRegistered string `xml:"DateRegistered"`
				Active         string `xml:"Active"`
			} `xml:"SyncCompany"`
		} `xml:"Items"`
	} `xml:"GetCompanySyncListResult"`
}

func (r *GetCompanySyncListRequest) URL() (*url.URL, error) {
	u, err := r.client.GetEndpointURL("/CRM/Company/V001/CompanyService.asmx", r.PathParams())
	return &u, err
}

func (r *GetCompanySyncListRequest) Do() (GetCompanySyncListRequestResponseBody, error) {
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
