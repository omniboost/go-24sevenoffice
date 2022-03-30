package netsuite_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/omniboost/go-24sevenoffice"
)

func TestTransactionSearch(t *testing.T) {
	client.SearchPreferences.PageSize = 10
	// client.SearchPreferences.ReturnSearchColumns = true
	client.SearchPreferences.BodyFieldsOnly = false
	req := client.NewTransactionSearchRequest()
	// req.RequestBody().SearchRecord.Basic.Name.Operator = "contains"
	// req.RequestBody().SearchRecord.Basic.Name.SearchValue = "e"
	req.RequestBody().SearchRecord.Basic.InternalID.Operator = "anyOf"
	req.RequestBody().SearchRecord.Basic.InternalID.SearchValue = []netsuite.RecordRef{
		{InternalID: "1302826"},
	}
	// req.RequestBody().SearchRecord.Basic.Type.Operator = "anyOf"
	// req.RequestBody().SearchRecord.Basic.Type.SearchValue = "_invoice"
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
