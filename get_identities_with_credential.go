package twentyfour

import (
	"encoding/xml"
	"net/http"
	"net/url"

	"github.com/omniboost/go-24sevenoffice/utils"
	"github.com/pkg/errors"
)

func (c *Client) NewGetIdentitiesWithCredentialRequest() GetIdentitiesWithCredentialRequest {
	r := GetIdentitiesWithCredentialRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type GetIdentitiesWithCredentialRequest struct {
	client      *Client
	queryParams *GetIdentitiesWithCredentialRequestQueryParams
	pathParams  *GetIdentitiesWithCredentialRequestPathParams
	method      string
	headers     http.Header
	requestBody GetIdentitiesWithCredentialRequestBody
}

func (r GetIdentitiesWithCredentialRequest) SOAPAction() string {
	return "http://24sevenOffice.com/webservices/GetIdentitiesWithCredential"
}

func (r GetIdentitiesWithCredentialRequest) NewQueryParams() *GetIdentitiesWithCredentialRequestQueryParams {
	return &GetIdentitiesWithCredentialRequestQueryParams{}
}

type GetIdentitiesWithCredentialRequestQueryParams struct {
}

func (p GetIdentitiesWithCredentialRequestQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetIdentitiesWithCredentialRequest) QueryParams() QueryParams {
	return r.queryParams
}

func (r GetIdentitiesWithCredentialRequest) NewPathParams() *GetIdentitiesWithCredentialRequestPathParams {
	return &GetIdentitiesWithCredentialRequestPathParams{}
}

type GetIdentitiesWithCredentialRequestPathParams struct {
}

func (p *GetIdentitiesWithCredentialRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetIdentitiesWithCredentialRequest) PathParams() PathParams {
	return r.pathParams
}

func (r *GetIdentitiesWithCredentialRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetIdentitiesWithCredentialRequest) Method() string {
	return r.method
}

func (r GetIdentitiesWithCredentialRequest) NewRequestBody() GetIdentitiesWithCredentialRequestBody {
	return GetIdentitiesWithCredentialRequestBody{}
}

type GetIdentitiesWithCredentialRequestBody struct {
	XMLName xml.Name `xml:"web:GetIdentitiesWithCredential"`

	Credential struct {
		ApplicationID string `xml:"web:ApplicationId"`
		IdentidyID    string `xml:"web:IdentidyId,omitempty"`
		Username      string `xml:"web:Username"`
		Password      string `xml:"web:Password"`
	} `xml:"web:credential"`
}

func (r *GetIdentitiesWithCredentialRequest) RequestBody() *GetIdentitiesWithCredentialRequestBody {
	return &r.requestBody
}

func (r *GetIdentitiesWithCredentialRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *GetIdentitiesWithCredentialRequest) SetRequestBody(body GetIdentitiesWithCredentialRequestBody) {
	r.requestBody = body
}

func (r *GetIdentitiesWithCredentialRequest) NewResponseBody() *GetIdentitiesWithCredentialRequestResponseBody {
	return &GetIdentitiesWithCredentialRequestResponseBody{}
}

type GetIdentitiesWithCredentialRequestResponseBody struct {
	XMLName xml.Name `xml:"GetIdentitiesWithCredentialResponse"`

	GetIdentitiesWithCredentialResult struct {
		Identity struct {
			ID   string `xml:"Id"`
			User struct {
				ContactId  string `xml:"ContactId"`
				ID         string `xml:"Id"`
				Name       string `xml:"Name"`
				EmployeeId string `xml:"EmployeeId"`
			} `xml:"User"`
			Client struct {
				ID   string `xml:"Id"`
				Name string `xml:"Name"`
			} `xml:"Client"`
			IsCurrent   string `xml:"IsCurrent"`
			IsDefault   string `xml:"IsDefault"`
			IsProtected string `xml:"IsProtected"`
			Servers     struct {
				Server struct {
					ID   string `xml:"Id"`
					Type string `xml:"Type"`
				} `xml:"Server"`
			} `xml:"Servers"`
		} `xml:"Identity"`
	} `xml:"GetIdentitiesWithCredentialResult"`
}

func (r *GetIdentitiesWithCredentialRequest) URL() (*url.URL, error) {
	u, err := r.client.GetEndpointURL("/authenticate/v001/authenticate.asmx", r.PathParams())
	return &u, err
}

func (r *GetIdentitiesWithCredentialRequest) Do() (GetIdentitiesWithCredentialRequestResponseBody, error) {
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
