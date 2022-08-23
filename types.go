package twentyfour

import (
	"encoding/xml"

	"github.com/cydev/zero"
	"github.com/omniboost/go-24sevenoffice/omitempty"
)

type BundleList struct {
	XMLName xml.Name `xml:"web:BundleList"`

	AllowDifference   bool    `xml:"web:AllowDifference"`
	DirectLedger      bool    `xml:"web:DirectLedger"`
	SaveOption        int     `xml:"web:SaveOption"`
	DefaultCustomerId int     `xml:"web:DefaultCustomerId"`
	Bundles           Bundles `xml:"web:Bundles>web:Bundle"`
}

type Bundles []Bundle

type Bundle struct {
	XMLName xml.Name `xml:"web:Bundle"`

	YearID                 int      `xml:"web:YearId"`
	Sort                   int      `xml:"web:Sort"`
	Name                   string   `xml:"web:Name"`
	BundleDirectAccounting bool     `xml:"web:BundleDirectAccounting"`
	Vouchers               Vouchers `xml:"web:Vouchers>web:Voucher"`
}

type Vouchers []Voucher

type Voucher struct {
	XMLName xml.Name `xml:"web:Voucher"`

	TransactionNo int            `xml:"web:TransactionNo"`
	Sort          int            `xml:"web:Sort"`
	Entries       VoucherEntries `xml:"web:Entries>web:Entry"`
}

type VoucherEntries []VoucherEntry

type VoucherEntry struct {
	XMLName xml.Name `xml:"web:Entry"`

	AccountNo    string  `xml:"web:AccountNo"`
	CustomerID   int     `xml:"web:CustomerId"`
	Date         Date    `xml:"web:Date"`
	DueDate      Date    `xml:"web:DueDate"`
	Amount       float64 `xml:"web:Amount"`
	CurrencyID   string  `xml:"web:CurrencyId"`
	CurrencyRate float64 `xml:"web:CurrencyRate"`
	CurrencyUnit int     `xml:"web:CurrencyUnit"`
	Comment      string  `xml:"web:Comment"`
	LinkId       string  `xml:"web:LinkId,omitempty"`
	TaxNo        int     `xml:"web:TaxNo,omitempty"` // actually TaxID
}

type Identities []Identity

type Identity struct {
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
}

type Invoices []Invoice

type Invoice struct {
	OrderId               int                   `xml:"OrderId,omitempty"`                // Can be specified in requests to SaveInvoices, otherwise set by the system. If an existing orderId is set the request is treated as an update for an InvoiceOrder. Default value: int.MinValue
	CustomerId            int                   `xml:"CustomerId,omitempty"`             // Required. Must exist in system. Default value: int.MinValue
	CustomerName          string                `xml:"CustomerName,omitempty"`           // Default value: “”. Max length 100 characters.
	CustomerDeliveryName  string                `xml:"CustomerDeliveryName,omitempty"`   // Default value: “”. Use this property when getting the Customer Delivery Name. the Delivery Name property under ‘Addresses’ can also be used for setting the name, but only this property can be used for both getting and setting the Delivery Name. Max length 250 characters
	CustomerDeliveryPhone string                `xml:"CustomerDeliveryPhone,omitempty"`  // Customer Delivery Phone. Max length 25 characters
	DeliveryAlternative   string                `xml:"DeliveryAlternative,omitempty"`    // Default value: “”. Max length 250 characters
	Addresses             Addresses             `xml:"Addresses,omitempty"`              // Default value: null
	OrderStatus           OrderSlipStateType    `xml:"OrderStatus,omitempty"`            // Default value: Web
	InvoiceId             int                   `xml:"InvoiceId,omitempty"`              // Cannot be specified in requests to SaveInvoices, this is always defined by the system. Default value: int.MinValue
	DateOrdered           DateTime              `xml:"DateOrdered,omitempty"`            // Default value: DateTime.MinValue
	DateInvoiced          DateTime              `xml:"DateInvoiced,omitempty"`           // Default value: DateTime.MinValue
	DateChanged           DateTime              `xml:"DateChanged,omitempty"`            // Default value: DateTime.MinValue
	PaymentTime           int                   `xml:"PaymentTime,omitempty"`            // Default (no change): int.MinValue
	CustomerReferenceNo   string                `xml:"CustomerReferenceNo,omitempty"`    // Deprecated
	ProjectId             int                   `xml:"ProjectId,omitempty"`              // If set, must exist in system. Default value: int.MinValue
	OurReference          int                   `xml:"OurReference,omitempty"`           // EmployeeId. If set, must exist in system. Default value: int.MinValue
	IncludeVAT            bool                  `xml:"IncludeVAT,omitempty"`             // Default value: null
	YourReference         string                `xml:"YourReference,omitempty"`          // Default value: “”. Max length 50 characters
	OrderTotalIncVat      float64               `xml:"OrderTotalIncVat,omitempty"`       // Default value: Decimal.MinValue. Read only
	OrderTotalVat         float64               `xml:"OrderTotalVat,omitempty"`          // Default value: Decimal.MinValue. Read only
	InvoiceTitle          string                `xml:"InvoiceTitle,omitempty"`           // Default value: “”. Max length 300 characters. This is the "comment" field of an invoice.
	InvoiceText           string                `xml:"InvoiceText,omitempty"`            // Default value: “”. Max length 750 characters
	Paid                  DateTime              `xml:"Paid,omitempty"`                   // Default value: DateTime.MinValue
	OCR                   string                `xml:"OCR,omitempty"`                    // Default value: “”. This is the invoice KID number. Max length 32 characters
	CustomerOrgNo         string                `xml:"CustomerOrgNo,omitempty"`          // Default value: “”. Max length 20 characters
	Currency              Currency              `xml:"Currency,omitempty"`               // Default value: null
	PaymentMethodId       int                   `xml:"PaymentMethodId,omitempty"`        // Default value: int.MinValue
	PaymentAmount         float64               `xml:"PaymentAmount,omitempty"`          // Write only property. Used for registering payments. Default value: Decimal.MinValue
	ProductionManagerId   int                   `xml:"ProductionManagerId,omitempty"`    // If set, must exist in system. Default value: int.MinValue
	SalesOpportunityId    int                   `xml:"SalesOpportunityId,omitempty"`     // If set, must exist in system. Default value: int.MinValue
	TypeOfSaleId          int                   `xml:"TypeOfSaleId,omitempty"`           // Default value: int.MinValue. You can get the values from GetTypeGroupList in the ClientService
	Distributor           Distributor           `xml:"Distributor,omitempty"`            // Default value: Default
	DistributionMethod    DistributionMethod    `xml:"DistributionMethod,omitempty"`     // Default value: Unchanged
	DepartmentId          int                   `xml:"DepartmentId,omitempty"`           // If set, must exist in system. Default value: int.MinValue
	ExternalStatus        int                   `xml:"ExternalStatus,omitempty"`         // Default value: int.MinValue
	DeliveryDate          DateTime              `xml:"DeliveryDate,omitempty"`           // Default value: DateTime.MinValue
	SkipStock             bool                  `xml:"SkipStock,omitempty"`              // Default value: false
	ProductionNumber      string                `xml:"ProductionNumber,omitempty"`       // Default value: “”. Max length 20 characters
	ReferenceInvoiceId    int                   `xml:"ReferenceInvoiceId,omitempty"`     // Default value: int.MinValue. Used for reference to original invoice when making a credit note.
	ReferenceOrderId      int                   `xml:"ReferenceOrderId,omitempty"`       // Default value: int.MinValue. Used for reference to original order when making a credit note.
	ReferenceNumber       string                `xml:"ReferenceNumber,omitempty"`        // Default value: “”. Max length 50 characters
	InvoiceEmailAddress   string                `xml:"InvoiceEmailAddress,omitempty"`    // Default value: “”. Max length 250 characters
	AccrualDate           DateTime              `xml:"AccrualDate,omitempty"`            // Default value: DateTime.MinValue. Determines the start date for the accrual period(s)
	AccrualLength         int                   `xml:"AccrualLength,omitempty"`          // Default value: int.MinValue. Sets the number of accrual months
	RoundFactor           float64               `xml:"RoundFactor,omitempty"`            // Default value: Decimal.MinValue. May be set to 0.1, 0.5 or 1.0
	InvoiceTemplateId     Guid                  `xml:"InvoiceTemplateId,omitempty"`      // Default value: 00000000-0000-0000-0000-000000000000
	VippsNumber           string                `xml:"VippsNumber,omitempty"`            // Deprecated
	DeliveryMethod        DeliveryMethod        `xml:"DeliveryMethod,omitempty"`         // Default value: null
	SendToFactoring       bool                  `xml:"SendToFactoring,omitempty"`        // Default value: true
	Commission            float64               `xml:"Commission,omitempty"`             // Default value: Decimal.MinValue
	InvoiceRows           InvoiceRows           `xml:"InvoiceRows>InvoiceRow,omitempty"` // Default value: null
	APIException          APIException          `xml:"APIException,omitempty"`           // Default value: null
	UserDefinedDimensions UserDefinedDimensions `xml:"UserDefinedDimensions,omitempty"`  // Dimensions defined by the user
	GLNNumber             string                `xml:"GLNNumber,omitempty"`              // Default value: “”. Uses Customer GLNNumber if default value is used and customer card has GLNNumber. To override this logic and set GLNNumber to the value NULL, the value in your request should be set to the string value NULL. Max length 13 characters
	CustomerDeliveryId    int                   `xml:"CustomerDeliveryId,omitempty"`     // Default value: int.MinValue. CustomerId of recipient, in cases where it differs from CustomerId
}

func (i Invoice) IsEmpty() bool {
	return zero.IsZero(i)
}

func (i Invoice) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return omitempty.MarshalXML(i, e, start)
}

type InvoiceOrders []InvoiceOrder

type InvoiceOrder struct {
	OrderID              string    `xml:"OrderId"`
	CustomerID           string    `xml:"CustomerId"`
	CustomerName         string    `xml:"CustomerName"`
	CustomerDeliveryName string    `xml:"CustomerDeliveryName"`
	Addresses            Addresses `xml:"Addresses"`
	OrderStatus          string    `xml:"OrderStatus"`
	InvoiceID            string    `xml:"InvoiceId"`
	DateOrdered          Date      `xml:"DateOrdered"`
	DateInvoiced         Date      `xml:"DateInvoiced"`
	DateChanged          Date      `xml:"DateChanged"`
	PaymentTime          string    `xml:"PaymentTime"`
	OurReference         string    `xml:"OurReference"`
	IncludeVAT           struct {
		Nil string `xml:"nil,attr"`
	} `xml:"IncludeVAT"`
	OrderTotalIncVat string `xml:"OrderTotalIncVat"`
	OrderTotalVat    string `xml:"OrderTotalVat"`
	CustomerOrgNo    string `xml:"CustomerOrgNo"`
	Currency         struct {
		Symbol string `xml:"Symbol"`
	} `xml:"Currency"`
	TypeOfSaleID        string      `xml:"TypeOfSaleId"`
	InvoiceEmailAddress string      `xml:"InvoiceEmailAddress"`
	InvoiceRows         InvoiceRows `xml:"InvoiceRows>InvoiceRow"`
	AccrualLength       string      `xml:"AccrualLength"`
	RoundFactor         string      `xml:"RoundFactor"`
	InvoiceTemplateID   string      `xml:"InvoiceTemplateId"`
	DeliveryMethod      string      `xml:"DeliveryMethod"`
	SendToFactoring     string      `xml:"SendToFactoring"`
	Commission          string      `xml:"Commission"`
}

type Companies []Company

type Company struct {
	ID                    int          `xml:"Id,omitempty"`
	ExternalID            string       `xml:"ExternalId,omitempty"`
	Name                  string       `xml:"Name"`
	FirstName             string       `xml:"FirstName,omitempty"`
	Addresses             Addresses    `xml:"Addresses,omitempty"`
	Country               string       `xml:"Country,omitempty"`
	InvoiceLanguage       string       `xml:"InvoiceLanguage,omitempty"`
	Type                  string       `xml:"Type,omitempty"`
	DateCreated           Date         `xml:"DateCreated,omitempty"`
	PriceList             string       `xml:"PriceList,omitempty"`
	Owner                 string       `xml:"Owner,omitempty"`
	TypeGroup             string       `xml:"TypeGroup,omitempty"`
	DateChanged           Date         `xml:"DateChanged,omitempty"`
	Maps                  string       `xml:"Maps,omitempty"`
	PaymentTime           string       `xml:"PaymentTime,omitempty"`
	Factoring             string       `xml:"Factoring,omitempty"`
	LedgerCustomerAccount string       `xml:"LedgerCustomerAccount,omitempty"`
	LedgerSupplierAccount string       `xml:"LedgerSupplierAccount,omitempty"`
	Private               string       `xml:"Private,omitempty"`
	APIException          APIException `xml:"APIException,omitempty"`
}

// func (c Company) IsEmpty() bool {
// 	return zero.IsZero(c)
// }

// func (c Company) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
// 	return omitempty.MarshalXML(c, e, start)
// }

type APIException struct {
	Type    string `xml:"web:Type"`
	Message string `xml:"web:Message"`
}

type Addresses struct {
	Post     Address `xml:"Post,omitempty"`
	Delivery Address `xml:"Delivery,omitempty"`
	Visit    Address `xml:"Visit,omitempty"`
	Invoice  Address `xml:"Invoice,omitempty"`
}

func (a Addresses) IsEmpty() bool {
	return zero.IsZero(a)
}

func (a Addresses) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return omitempty.MarshalXML(a, e, start)
}

type Address struct {
	Street     string `xml:"web:Street,omitempty"`
	State      string `xml:"web:State,omitempty"`
	PostalCode string `xml:"web:PostalCode,omitempty"`
	PostalArea string `xml:"web:PostalArea,omitempty"`
	Name       string `xml:"web:Name,omitempty"`
	City       string `xml:"web:City,omitempty"`
	Country    string `xml:"web:Country,omitempty"`
}

func (a Address) IsEmpty() bool {
	return zero.IsZero(a)
}

type InvoiceRows []InvoiceRow

type InvoiceRow struct {
	ProductId             int                   `xml:"ProductId,omitempty"`             // Required for Row Type Normal. Must exist in system. Default value: int.MinValue
	ProductNo             string                `xml:"ProductNo,omitempty"`             // Read only. When creating orders with SaveInvoices you must use ProductId
	RowId                 int                   `xml:"RowId,omitempty"`                 // Used when editing an existing order. Default value: int.MinValue
	VatRate               float64               `xml:"VatRate,omitempty"`               // Default value: Decimal.MinValue. Read only
	Price                 float64               `xml:"Price,omitempty"`                 // Default value: Decimal.MinValue
	Name                  string                `xml:"Name,omitempty"`                  // Default value: “”. Max length 300 characters
	DiscountRate          float64               `xml:"DiscountRate,omitempty"`          // Default value: Decimal.MinValue
	Quantity              float64               `xml:"Quantity,omitempty"`              // Default value: Decimal.MinValue
	QuantityDelivered     float64               `xml:"QuantityDelivered,omitempty"`     // Default value: Decimal.MinValue
	QuantityOrdered       float64               `xml:"QuantityOrdered,omitempty"`       // Default value: Decimal.MinValue
	QuantityRest          float64               `xml:"QuantityRest,omitempty"`          // -1. This property is used for creating rest orders
	Cost                  float64               `xml:"Cost,omitempty"`                  // Default value: Decimal.MinValue
	InPrice               float64               `xml:"InPrice,omitempty"`               // Default value: Decimal.MinValue
	SequenceNumber        int                   `xml:"SequenceNumber,omitempty"`        // Default value: Int16.MinValue
	Hidden                bool                  `xml:"Hidden,omitempty"`                // Default value: false. Makes the row hidden on the actual invoice statement.
	Type                  RowType               `xml:"Type,omitempty"`                  // Normal or Text. Default value: Normal. Please note that TextBold is deprecated
	AccrualDate           DateTime              `xml:"AccrualDate,omitempty"`           // Default value: DateTime.MinValue
	AccrualLength         int                   `xml:"AccrualLength,omitempty"`         // Default value: int.MinValue
	ChangeState           ChangeState           `xml:"ChangeState,omitempty"`           // This property must be used when editing an already exisiting order. Default value: ChangeState.None
	TypeGroupId           int                   `xml:"TypeGroupId,omitempty"`           //
	AccountProject        bool                  `xml:"AccountProject,omitempty"`        // N/A
	DepartmentId          int                   `xml:"DepartmentId,omitempty"`          //
	ProjectId             int                   `xml:"ProjectId,omitempty"`             //
	UserDefinedDimensions UserDefinedDimensions `xml:"UserDefinedDimensions,omitempty"` // Dimensions defined by the user
	TaxSettings           TaxSettings           `xml:"TaxSettings,omitempty"`           // Override the typegroup tax rate for this invoice row
}

type RowType string

type ChangeState string

type UserDefinedDimensions []interface{}

type TaxSettings struct {
	TaxAccount int     `xml:"TaxAccount,omitempty"` // The account
	TaxCode    int     `xml:"TaxCode,omitempty"`    // Get this from GetTaxCodeList in the AccountService
	TaxRate    float64 `xml:"TaxRate,omitempty"`    // The tax rate you wish to set
}

type DeliveryMethod struct {
	ID          int    `xml:"Id,omitempty"`          // DeliveryId
	Description string `xml:"Description,omitempty"` // Description
}

type OrderSlipStateType string

type Currency string

type Distributor interface{}

type DistributionMethod interface{}

type Guid string

type TaxCodes []TaxCode

type TaxCode struct {
	TaxID     int     `xml:"TaxId"`
	TaxNo     string  `xml:"TaxNo"`
	TaxName   string  `xml:"TaxName"`
	TaxRate   float64 `xml:"TaxRate"`
	AccountNo string  `xml:"AccountNo"`
}
