package twentyfour

import (
	"encoding/xml"
	"net/http"
	"net/url"

	"github.com/omniboost/go-24sevenoffice/utils"
	"github.com/pkg/errors"
)

func (c *Client) NewGetIdentitiesRequest() GetIdentitiesRequest {
	r := GetIdentitiesRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type GetIdentitiesRequest struct {
	client      *Client
	queryParams *GetIdentitiesRequestQueryParams
	pathParams  *GetIdentitiesRequestPathParams
	method      string
	headers     http.Header
	requestBody GetIdentitiesRequestBody
}

func (r GetIdentitiesRequest) SOAPAction() string {
	return "http://24sevenOffice.com/webservices/GetIdentities"
}

func (r GetIdentitiesRequest) NewQueryParams() *GetIdentitiesRequestQueryParams {
	return &GetIdentitiesRequestQueryParams{}
}

type GetIdentitiesRequestQueryParams struct {
}

func (p GetIdentitiesRequestQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetIdentitiesRequest) QueryParams() QueryParams {
	return r.queryParams
}

func (r GetIdentitiesRequest) NewPathParams() *GetIdentitiesRequestPathParams {
	return &GetIdentitiesRequestPathParams{}
}

type GetIdentitiesRequestPathParams struct {
}

func (p *GetIdentitiesRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetIdentitiesRequest) PathParams() PathParams {
	return r.pathParams
}

func (r *GetIdentitiesRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetIdentitiesRequest) Method() string {
	return r.method
}

func (r GetIdentitiesRequest) NewRequestBody() GetIdentitiesRequestBody {
	return GetIdentitiesRequestBody{}
}

type GetIdentitiesRequestBody struct {
	XMLName xml.Name `xml:"GetIdentities"`
}

func (r *GetIdentitiesRequest) RequestBody() *GetIdentitiesRequestBody {
	return &r.requestBody
}

func (r *GetIdentitiesRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *GetIdentitiesRequest) SetRequestBody(body GetIdentitiesRequestBody) {
	r.requestBody = body
}

func (r *GetIdentitiesRequest) NewResponseBody() *GetIdentitiesRequestResponseBody {
	return &GetIdentitiesRequestResponseBody{}
}

type GetIdentitiesRequestResponseBody struct {
	XMLName xml.Name `xml:"GetIdentitiesResponse"`

	GetIdentitiesWithCredentialResult struct {
		Identities Identities `xml:"Identity"`
	} `xml:"GetIdentitiesResult"`
}

func (r *GetIdentitiesRequest) URL() (*url.URL, error) {
	u, err := r.client.GetEndpointURL("/authenticate/v001/authenticate.asmx", r.PathParams())
	return &u, err
}

func (r *GetIdentitiesRequest) Do() (GetIdentitiesRequestResponseBody, error) {
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
