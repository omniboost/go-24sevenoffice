package twentyfour

import (
	"encoding/xml"
	"net/http"
	"net/url"

	"github.com/hashicorp/go-multierror"
	"github.com/omniboost/go-24sevenoffice/utils"
	"github.com/pkg/errors"
)

func (c *Client) NewSaveCompaniesRequest() SaveCompaniesRequest {
	r := SaveCompaniesRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type SaveCompaniesRequest struct {
	client      *Client
	queryParams *SaveCompaniesRequestQueryParams
	pathParams  *SaveCompaniesRequestPathParams
	method      string
	headers     http.Header
	requestBody SaveCompaniesRequestBody
}

func (r SaveCompaniesRequest) SOAPAction() string {
	return "http://24sevenOffice.com/webservices/SaveCompanies"
}

func (r SaveCompaniesRequest) NewQueryParams() *SaveCompaniesRequestQueryParams {
	return &SaveCompaniesRequestQueryParams{}
}

type SaveCompaniesRequestQueryParams struct {
}

func (p SaveCompaniesRequestQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *SaveCompaniesRequest) QueryParams() QueryParams {
	return r.queryParams
}

func (r SaveCompaniesRequest) NewPathParams() *SaveCompaniesRequestPathParams {
	return &SaveCompaniesRequestPathParams{}
}

type SaveCompaniesRequestPathParams struct {
}

func (p *SaveCompaniesRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *SaveCompaniesRequest) PathParams() PathParams {
	return r.pathParams
}

func (r *SaveCompaniesRequest) SetMethod(method string) {
	r.method = method
}

func (r *SaveCompaniesRequest) Method() string {
	return r.method
}

func (r SaveCompaniesRequest) NewRequestBody() SaveCompaniesRequestBody {
	return SaveCompaniesRequestBody{}
}

type SaveCompaniesRequestBody struct {
	XMLName xml.Name `xml:"web:SaveCompanies"`

	Companies Companies `xml:"web:companies>web:Company"`
}

func (rb SaveCompaniesRequestBody) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name = xml.Name{Local: "web:SaveCompanies"}

	namespaces := []xml.Attr{
		{Name: xml.Name{Space: "", Local: "xmlns"}, Value: "http://24sevenOffice.com/webservices"},
	}
	for _, ns := range namespaces {
		start.Attr = append(start.Attr, ns)
	}

	type alias SaveCompaniesRequestBody
	a := alias(rb)
	return e.EncodeElement(a, start)
}

func (r *SaveCompaniesRequest) RequestBody() *SaveCompaniesRequestBody {
	return &r.requestBody
}

func (r *SaveCompaniesRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *SaveCompaniesRequest) SetRequestBody(body SaveCompaniesRequestBody) {
	r.requestBody = body
}

func (r *SaveCompaniesRequest) NewResponseBody() *SaveCompaniesRequestResponseBody {
	return &SaveCompaniesRequestResponseBody{}
}

type SaveCompaniesRequestResponseBody struct {
	XMLName xml.Name `xml:"SaveCompaniesResponse"`

	Companies Companies `xml:"SaveCompaniesResult>Company"`
}

func (r *SaveCompaniesRequest) URL() (*url.URL, error) {
	u, err := r.client.GetEndpointURL("/CRM/Company/V001/CompanyService.asmx", r.PathParams())
	return &u, err
}

func (r *SaveCompaniesRequest) Do() (SaveCompaniesRequestResponseBody, error) {
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

	var errs error
	for _, c := range responseBody.Companies {
		if c.APIException.Message != "" {
			errs = multierror.Append(errs, err)
		}
	}

	return *responseBody, errs
}
