package twentyfour_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestGetCompanies(t *testing.T) {
	req := client.NewGetCompaniesRequest()
	// req.RequestBody().SearchParams.CompanyID = 1
	req.RequestBody().SearchParams.ExternalID = "M00000059"
	req.RequestBody().ReturnProperties = []string{
		"APIException",
		"Id",
		"ExternalId",
		"OrganizationNumber",
		"Name",
		"FirstName",
		"NickName",
		"Addresses",
		"PhoneNumbers",
		"EmailAddresses",
		"Url",
		"Country",
		"Note",
		"InvoiceLanguage",
		"Type",
		"Username",
		"Password",
		"IncorporationDate",
		"DateCreated",
		"Status",
		"PriceList",
		"Owner",
		"BankAccountNo",
		"BankAccountType",
		"BankAccountCountry",
		"BankAccountBic",
		"TermsOfDeliveryId",
		"AccountDebit",
		"AccountCredit",
		"Discount",
		"TypeGroup",
		"ShareCapital",
		"NumberOfEmployees",
		"Turnover",
		"Profit",
		"IndustryId",
		"MemberNo",
		"DateChanged",
		"BlockInvoice",
		"Relations",
		"Maps",
		"DistributionMethod",
		"CurrencyId",
		"PaymentTime",
		"GLNNumber",
		"Factoring",
		"LedgerCustomerAccount",
		"LedgerSupplierAccount",
		"VatNumber",
		"Private",
		"ExplicitlySpecifyNewCompanyId",
	}
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
