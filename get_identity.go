package twentyfour

import (
	"encoding/xml"
	"net/http"
	"net/url"

	"github.com/omniboost/go-24sevenoffice/utils"
	"github.com/pkg/errors"
)

func (c *Client) NewGetIdentityRequest() GetIdentityRequest {
	r := GetIdentityRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type GetIdentityRequest struct {
	client      *Client
	queryParams *GetIdentityRequestQueryParams
	pathParams  *GetIdentityRequestPathParams
	method      string
	headers     http.Header
	requestBody GetIdentityRequestBody
}

func (r GetIdentityRequest) SOAPAction() string {
	return "http://24sevenOffice.com/webservices/GetIdentity"
}

func (r GetIdentityRequest) NewQueryParams() *GetIdentityRequestQueryParams {
	return &GetIdentityRequestQueryParams{}
}

type GetIdentityRequestQueryParams struct {
}

func (p GetIdentityRequestQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetIdentityRequest) QueryParams() QueryParams {
	return r.queryParams
}

func (r GetIdentityRequest) NewPathParams() *GetIdentityRequestPathParams {
	return &GetIdentityRequestPathParams{}
}

type GetIdentityRequestPathParams struct {
}

func (p *GetIdentityRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetIdentityRequest) PathParams() PathParams {
	return r.pathParams
}

func (r *GetIdentityRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetIdentityRequest) Method() string {
	return r.method
}

func (r GetIdentityRequest) NewRequestBody() GetIdentityRequestBody {
	return GetIdentityRequestBody{}
}

type GetIdentityRequestBody struct {
	XMLName xml.Name `xml:"GetIdentity"`
}

func (r *GetIdentityRequest) RequestBody() *GetIdentityRequestBody {
	return &r.requestBody
}

func (r *GetIdentityRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *GetIdentityRequest) SetRequestBody(body GetIdentityRequestBody) {
	r.requestBody = body
}

func (r *GetIdentityRequest) NewResponseBody() *GetIdentityRequestResponseBody {
	return &GetIdentityRequestResponseBody{}
}

type GetIdentityRequestResponseBody struct {
	XMLName xml.Name `xml:"GetIdentityResponse"`

	GetIdentityWithCredentialResult struct {
		Identities Identities `xml:"Identity"`
	} `xml:"GetIdentityResult"`
}

func (r *GetIdentityRequest) URL() (*url.URL, error) {
	u, err := r.client.GetEndpointURL("/authenticate/v001/authenticate.asmx", r.PathParams())
	return &u, err
}

func (r *GetIdentityRequest) Do() (GetIdentityRequestResponseBody, error) {
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
