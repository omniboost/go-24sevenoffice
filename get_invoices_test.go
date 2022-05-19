package twentyfour_test

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	twentyfour "github.com/omniboost/go-24sevenoffice"
)

func TestGetInvoices(t *testing.T) {
	req := client.NewGetInvoicesRequest()
	req.RequestBody().SearchParams.ChangedAfter = twentyfour.DateTime{time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)}
	req.RequestBody().InvoiceReturnProperties = []string{
		"OrderId",
		"CustomerId",
		"CustomerName",
		"CustomerDeliveryName",
		"CustomerDeliveryPhone",
		"DeliveryAlternative",
		"Addresses",
		"OrderStatus",
		"InvoiceId",
		"DateOrdered",
		"DateInvoiced",
		"DateChanged",
		"PaymentTime",
		"CustomerReferenceNo",
		"ProjectId",
		"OurReference",
		"IncludeVAT",
		"YourReference",
		"OrderTotalIncVat",
		"OrderTotalVat",
		"InvoiceTitle",
		"InvoiceText",
		"Paid",
		"OCR",
		"CustomerOrgNo",
		"Currency",
		"PaymentMethodId",
		"PaymentAmount",
		"ProductionManagerId",
		"SalesOpportunityId",
		"TypeOfSaleId",
		"Distributor",
		"DistributionMethod",
		"DepartmentId",
		"ExternalStatus",
		"DeliveryDate",
		"SkipStock",
		"ProductionNumber",
		"ReferenceInvoiceId",
		"ReferenceOrderId",
		"ReferenceNumber",
		"InvoiceEmailAddress",
		"AccrualDate",
		"AccrualLength",
		"RoundFactor",
		"InvoiceTemplateId",
		"VippsNumber",
		"DeliveryMethod",
		"SendToFactoring",
		"Commission",
		"InvoiceRows",
		"APIException",
		"UserDefinedDimensions",
		"GLNNumber",
		"CustomerDeliveryId",
	}
	req.RequestBody().RowReturnProperties = []string{
		"ProductId",
		"RowId",
		"Name",
		"Quantity",
		"Type",
		"Price",
	}
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
