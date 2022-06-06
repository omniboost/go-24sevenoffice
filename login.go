package twentyfour

import (
	"encoding/xml"
	"net/http"
	"net/url"

	"github.com/omniboost/go-24sevenoffice/utils"
	"github.com/pkg/errors"
)

func (c *Client) NewLoginRequest() LoginRequest {
	r := LoginRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type LoginRequest struct {
	client      *Client
	queryParams *LoginRequestQueryParams
	pathParams  *LoginRequestPathParams
	method      string
	headers     http.Header
	requestBody LoginRequestBody
}

func (r LoginRequest) SOAPAction() string {
	return "http://24sevenOffice.com/webservices/Login"
}

func (r LoginRequest) NewQueryParams() *LoginRequestQueryParams {
	return &LoginRequestQueryParams{}
}

type LoginRequestQueryParams struct {
}

func (p LoginRequestQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *LoginRequest) QueryParams() QueryParams {
	return r.queryParams
}

func (r LoginRequest) NewPathParams() *LoginRequestPathParams {
	return &LoginRequestPathParams{}
}

type LoginRequestPathParams struct {
}

func (p *LoginRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *LoginRequest) PathParams() PathParams {
	return r.pathParams
}

func (r *LoginRequest) SetMethod(method string) {
	r.method = method
}

func (r *LoginRequest) Method() string {
	return r.method
}

func (r LoginRequest) NewRequestBody() LoginRequestBody {
	return LoginRequestBody{}
}

type LoginRequestBody struct {
	XMLName xml.Name `xml:"web:Login"`

	Credential struct {
		ApplicationID string `xml:"web:ApplicationId"`
		IdentityID    string `xml:"web:IdentityId,omitempty"`
		Username      string `xml:"web:Username"`
		Password      string `xml:"web:Password"`
	} `xml:"web:credential"`
}

func (r *LoginRequest) RequestBody() *LoginRequestBody {
	return &r.requestBody
}

func (r *LoginRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *LoginRequest) SetRequestBody(body LoginRequestBody) {
	r.requestBody = body
}

func (r *LoginRequest) NewResponseBody() *LoginRequestResponseBody {
	return &LoginRequestResponseBody{}
}

type LoginRequestResponseBody struct {
	XMLName xml.Name `xml:"LoginResponse"`

	LoginResult string `xml:"LoginResult"`
}

func (r *LoginRequest) URL() (*url.URL, error) {
	u, err := r.client.GetEndpointURL("/authenticate/v001/authenticate.asmx", r.PathParams())
	return &u, err
}

func (r *LoginRequest) Do() (LoginRequestResponseBody, error) {
	var err error

	// Create http request
	req, err := r.client.NewRequest(nil, r)
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
