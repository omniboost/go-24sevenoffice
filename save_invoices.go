package twentyfour

import (
	"encoding/xml"
	"net/http"
	"net/url"

	"github.com/hashicorp/go-multierror"
	"github.com/omniboost/go-24sevenoffice/utils"
	"github.com/pkg/errors"
)

func (c *Client) NewSaveInvoicesRequest() SaveInvoicesRequest {
	r := SaveInvoicesRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type SaveInvoicesRequest struct {
	client      *Client
	queryParams *SaveInvoicesRequestQueryParams
	pathParams  *SaveInvoicesRequestPathParams
	method      string
	headers     http.Header
	requestBody SaveInvoicesRequestBody
}

func (r SaveInvoicesRequest) SOAPAction() string {
	return "http://24sevenOffice.com/webservices/SaveInvoices"
}

func (r SaveInvoicesRequest) NewQueryParams() *SaveInvoicesRequestQueryParams {
	return &SaveInvoicesRequestQueryParams{}
}

type SaveInvoicesRequestQueryParams struct {
}

func (p SaveInvoicesRequestQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *SaveInvoicesRequest) QueryParams() QueryParams {
	return r.queryParams
}

func (r SaveInvoicesRequest) NewPathParams() *SaveInvoicesRequestPathParams {
	return &SaveInvoicesRequestPathParams{}
}

type SaveInvoicesRequestPathParams struct {
}

func (p *SaveInvoicesRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *SaveInvoicesRequest) PathParams() PathParams {
	return r.pathParams
}

func (r *SaveInvoicesRequest) SetMethod(method string) {
	r.method = method
}

func (r *SaveInvoicesRequest) Method() string {
	return r.method
}

func (r SaveInvoicesRequest) NewRequestBody() SaveInvoicesRequestBody {
	return SaveInvoicesRequestBody{}
}

type SaveInvoicesRequestBody struct {
	XMLName xml.Name `xml:"web:SaveInvoices"`

	Invoices Invoices `xml:"web:invoices>web:InvoiceOrder"`
}

func (rb SaveInvoicesRequestBody) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name = xml.Name{Local: "web:SaveInvoices"}

	namespaces := []xml.Attr{
		{Name: xml.Name{Space: "", Local: "xmlns"}, Value: "http://24sevenOffice.com/webservices"},
	}
	for _, ns := range namespaces {
		start.Attr = append(start.Attr, ns)
	}

	type alias SaveInvoicesRequestBody
	a := alias(rb)
	return e.EncodeElement(a, start)
}

func (r *SaveInvoicesRequest) RequestBody() *SaveInvoicesRequestBody {
	return &r.requestBody
}

func (r *SaveInvoicesRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *SaveInvoicesRequest) SetRequestBody(body SaveInvoicesRequestBody) {
	r.requestBody = body
}

func (r *SaveInvoicesRequest) NewResponseBody() *SaveInvoicesRequestResponseBody {
	return &SaveInvoicesRequestResponseBody{}
}

type SaveInvoicesRequestResponseBody struct {
	XMLName xml.Name `xml:"SaveInvoicesResponse"`

	Invoices Invoices `xml:"SaveInvoicesResult>Invoice"`
}

func (r *SaveInvoicesRequest) URL() (*url.URL, error) {
	u, err := r.client.GetEndpointURL("/Economy/InvoiceOrder/V001/InvoiceService.asmx", r.PathParams())
	return &u, err
}

func (r *SaveInvoicesRequest) Do() (SaveInvoicesRequestResponseBody, error) {
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
	for _, c := range responseBody.Invoices {
		if c.APIException.Message != "" {
			errs = multierror.Append(errs, errors.Errorf("%s: %s", c.APIException.Type, c.APIException.Message))
		}
	}

	return *responseBody, errs
}
