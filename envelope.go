package twentyfour

import (
	"encoding/xml"

	"github.com/cydev/zero"
	"github.com/omniboost/go-24sevenoffice/omitempty"
)

type RequestEnvelope struct {
	XMLName xml.Name

	Header Header `xml:"env:Header"`
	Body   Body   `xml:"env:Body"`
}

func NewRequestEnvelope() RequestEnvelope {
	return RequestEnvelope{
		Header: NewHeader(),
	}
}

type ResponseEnvelope struct {
	XMLName xml.Name

	Header Header `xml:"Header"`
	Body   Body   `xml:"Body"`
}

func (env RequestEnvelope) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name = xml.Name{Local: "env:Envelope"}

	namespaces := []xml.Attr{
		{Name: xml.Name{Space: "", Local: "xmlns:xsi"}, Value: "http://www.w3.org/2001/XMLSchema-instance"},
		{Name: xml.Name{Space: "", Local: "xmlns:xsd"}, Value: "http://www.w3.org/2001/XMLSchema"},
		{Name: xml.Name{Space: "", Local: "xmlns:env"}, Value: "http://schemas.xmlsoap.org/soap/envelope/"},
		{Name: xml.Name{Space: "", Local: "xmlns:web"}, Value: "http://24sevenOffice.com/webservices"},
	}
	for _, ns := range namespaces {
		start.Attr = append(start.Attr, ns)
	}

	type alias RequestEnvelope
	a := alias(env)
	return e.EncodeElement(a, start)
}

type Body struct {
	ActionBody interface{} `xml:",any"`
}

type Header struct {
}

func (h Header) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return omitempty.MarshalXML(h, e, start)
}

func (h Header) IsEmpty() bool {
	return zero.IsZero(h)
}

func NewHeader() Header {
	return Header{}
}

type ActionBody interface{}
