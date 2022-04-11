package twentyfour

import (
	"encoding/xml"
	"net/http"
	"net/url"

	"github.com/omniboost/go-24sevenoffice/utils"
	"github.com/pkg/errors"
)

func (c *Client) NewHasSessionRequest() HasSessionRequest {
	r := HasSessionRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type HasSessionRequest struct {
	client      *Client
	queryParams *HasSessionRequestQueryParams
	pathParams  *HasSessionRequestPathParams
	method      string
	headers     http.Header
	requestBody HasSessionRequestBody
}

func (r HasSessionRequest) SOAPAction() string {
	return "http://24sevenOffice.com/webservices/HasSession"
}

func (r HasSessionRequest) NewQueryParams() *HasSessionRequestQueryParams {
	return &HasSessionRequestQueryParams{}
}

type HasSessionRequestQueryParams struct {
}

func (p HasSessionRequestQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *HasSessionRequest) QueryParams() QueryParams {
	return r.queryParams
}

func (r HasSessionRequest) NewPathParams() *HasSessionRequestPathParams {
	return &HasSessionRequestPathParams{}
}

type HasSessionRequestPathParams struct {
}

func (p *HasSessionRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *HasSessionRequest) PathParams() PathParams {
	return r.pathParams
}

func (r *HasSessionRequest) SetMethod(method string) {
	r.method = method
}

func (r *HasSessionRequest) Method() string {
	return r.method
}

func (r HasSessionRequest) NewRequestBody() HasSessionRequestBody {
	return HasSessionRequestBody{}
}

type HasSessionRequestBody struct {
	XMLName xml.Name `xml:"web:HasSession"`
}

func (r *HasSessionRequest) RequestBody() *HasSessionRequestBody {
	return &r.requestBody
}

func (r *HasSessionRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *HasSessionRequest) SetRequestBody(body HasSessionRequestBody) {
	r.requestBody = body
}

func (r *HasSessionRequest) NewResponseBody() *HasSessionRequestResponseBody {
	return &HasSessionRequestResponseBody{}
}

type HasSessionRequestResponseBody struct {
	XMLName xml.Name `xml:"HasSessionResponse"`

	HasSessionResult bool `xml:"HasSessionResult"`
}

func (r *HasSessionRequest) URL() (*url.URL, error) {
	u, err := r.client.GetEndpointURL("/authenticate/v001/authenticate.asmx", r.PathParams())
	return &u, err
}

func (r *HasSessionRequest) Do() (HasSessionRequestResponseBody, error) {
	var err error

	// Create http request
	req, err := r.client.NewRequest(nil, r)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	r.client.SetSessionCookie(req, r.client.SessionID())

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
