package twentyfour

import (
	"encoding/xml"
	"net/http"
	"net/url"

	"github.com/omniboost/go-24sevenoffice/utils"
	"github.com/pkg/errors"
)

func (c *Client) NewSaveBundleListRequest() SaveBundleListRequest {
	r := SaveBundleListRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type SaveBundleListRequest struct {
	client      *Client
	queryParams *SaveBundleListRequestQueryParams
	pathParams  *SaveBundleListRequestPathParams
	method      string
	headers     http.Header
	requestBody SaveBundleListRequestBody
}

func (r SaveBundleListRequest) SOAPAction() string {
	return "http://24sevenOffice.com/webservices/SaveBundleList"
}

func (r SaveBundleListRequest) NewQueryParams() *SaveBundleListRequestQueryParams {
	return &SaveBundleListRequestQueryParams{}
}

type SaveBundleListRequestQueryParams struct {
}

func (p SaveBundleListRequestQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *SaveBundleListRequest) QueryParams() QueryParams {
	return r.queryParams
}

func (r SaveBundleListRequest) NewPathParams() *SaveBundleListRequestPathParams {
	return &SaveBundleListRequestPathParams{}
}

type SaveBundleListRequestPathParams struct {
}

func (p *SaveBundleListRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *SaveBundleListRequest) PathParams() PathParams {
	return r.pathParams
}

func (r *SaveBundleListRequest) SetMethod(method string) {
	r.method = method
}

func (r *SaveBundleListRequest) Method() string {
	return r.method
}

func (r SaveBundleListRequest) NewRequestBody() SaveBundleListRequestBody {
	return SaveBundleListRequestBody{}
}

type SaveBundleListRequestBody struct {
	XMLName xml.Name `xml:"web:SaveBundleList"`

	BundleList BundleList `xml:"web:BundleList"`
}

func (r *SaveBundleListRequest) RequestBody() *SaveBundleListRequestBody {
	return &r.requestBody
}

func (r *SaveBundleListRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *SaveBundleListRequest) SetRequestBody(body SaveBundleListRequestBody) {
	r.requestBody = body
}

func (r *SaveBundleListRequest) NewResponseBody() *SaveBundleListRequestResponseBody {
	return &SaveBundleListRequestResponseBody{}
}

type SaveBundleListRequestResponseBody struct {
	XMLName xml.Name `xml:"SaveBundleListResponse"`

	SaveBundleListResult struct {
		Type        string `xml:"Type"`
		Description string `xml:"Description"`
	} `xml:"SaveBundleListResult"`
}

func (r *SaveBundleListRequest) URL() (*url.URL, error) {
	u, err := r.client.GetEndpointURL("/Economy/Account/V004/Accountservice.asmx", r.PathParams())
	return &u, err
}

func (r *SaveBundleListRequest) Do() (SaveBundleListRequestResponseBody, error) {
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

	if responseBody.SaveBundleListResult.Type != "Ok" {
		return *responseBody, errors.Errorf("%s: %s", responseBody.SaveBundleListResult.Type, responseBody.SaveBundleListResult.Description)
	}

	return *responseBody, nil
}
