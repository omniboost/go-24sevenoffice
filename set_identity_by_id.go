package twentyfour

import (
	"encoding/xml"
	"net/http"
	"net/url"

	"github.com/omniboost/go-24sevenoffice/utils"
	"github.com/pkg/errors"
)

func (c *Client) NewSetIdentityByIDRequest() SetIdentityByIDRequest {
	r := SetIdentityByIDRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type SetIdentityByIDRequest struct {
	client      *Client
	queryParams *SetIdentityByIDRequestQueryParams
	pathParams  *SetIdentityByIDRequestPathParams
	method      string
	headers     http.Header
	requestBody SetIdentityByIDRequestBody
}

func (r SetIdentityByIDRequest) SOAPAction() string {
	return "http://24sevenOffice.com/webservices/SetIdentityById"
}

func (r SetIdentityByIDRequest) NewQueryParams() *SetIdentityByIDRequestQueryParams {
	return &SetIdentityByIDRequestQueryParams{}
}

type SetIdentityByIDRequestQueryParams struct {
}

func (p SetIdentityByIDRequestQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *SetIdentityByIDRequest) QueryParams() QueryParams {
	return r.queryParams
}

func (r SetIdentityByIDRequest) NewPathParams() *SetIdentityByIDRequestPathParams {
	return &SetIdentityByIDRequestPathParams{}
}

type SetIdentityByIDRequestPathParams struct {
}

func (p *SetIdentityByIDRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *SetIdentityByIDRequest) PathParams() PathParams {
	return r.pathParams
}

func (r *SetIdentityByIDRequest) SetMethod(method string) {
	r.method = method
}

func (r *SetIdentityByIDRequest) Method() string {
	return r.method
}

func (r SetIdentityByIDRequest) NewRequestBody() SetIdentityByIDRequestBody {
	return SetIdentityByIDRequestBody{}
}

type SetIdentityByIDRequestBody struct {
	XMLName xml.Name `xml:"web:SetIdentityById"`

	IdentityID string `xml:"web:identityId"`
}

func (r *SetIdentityByIDRequest) RequestBody() *SetIdentityByIDRequestBody {
	return &r.requestBody
}

func (r *SetIdentityByIDRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *SetIdentityByIDRequest) SetRequestBody(body SetIdentityByIDRequestBody) {
	r.requestBody = body
}

func (r *SetIdentityByIDRequest) NewResponseBody() *SetIdentityByIDRequestResponseBody {
	return &SetIdentityByIDRequestResponseBody{}
}

type SetIdentityByIDRequestResponseBody struct {
	XMLName xml.Name `xml:"SetIdentityByIdResponse"`

	SetIdentityByIDResult bool `xml:"SetIdentityByIdResult"`
}

func (r *SetIdentityByIDRequest) URL() (*url.URL, error) {
	u, err := r.client.GetEndpointURL("/authenticate/v001/authenticate.asmx", r.PathParams())
	return &u, err
}

func (r *SetIdentityByIDRequest) Do() (SetIdentityByIDRequestResponseBody, error) {
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
